package main

import (
	"crypto/rand"
	"database/sql"
	"database/sql/driver"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/wenerme/go-wecom/WeWorkFinanceSDK/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	_ "github.com/glebarez/go-sqlite"
	gsqlite "github.com/glebarez/go-sqlite"

	"github.com/glebarez/sqlite"
	dotenv "github.com/joho/godotenv"
	"github.com/oklog/ulid/v2"
	"github.com/sirupsen/logrus"
	"github.com/wenerme/go-wecom/WeWorkFinanceSDK"
)

var defaultEntropySource = ulid.Monotonic(rand.Reader, 0)

var fileID string
var keepPolling bool
var pollingInterval = time.Minute * 5
var printMessage = true

var _db *gorm.DB
var _client WeWorkFinanceSDK.Client

func MustNewULID() string {
	n := ulid.MustNew(ulid.Timestamp(time.Now()), defaultEntropySource)
	return strings.ToLower(n.String())
}

func main() {
	_ = dotenv.Load()
	flag.StringVar(&fileID, "file-id", "", "file id to pull")
	flag.BoolVar(&keepPolling, "keep-polling", false, "keep polling")
	flag.Parse()

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	client, err := WeWorkFinanceSDK.NewClientFromEnv()
	if err != nil {
		panic(err)
	}

	if fileID != "" {
		data, err := client.ReadMediaData(WeWorkFinanceSDK.GetMediaDataOptions{
			FileID: fileID,
		})
		if err != nil {
			panic(err)
		}
		sum := WeWorkFinanceSDK.MD5Sum(data)
		fmt.Printf(`file id: %s
size: %v
md5sum: %s
`, fileID, len(data), sum)
		return
	}

	cwd, _ := os.Getwd()
	mediaPath := path.Join(cwd, "media")
	// avoid cross device link
	tmpDir := path.Join(cwd, "tmp")
	_ = os.MkdirAll(mediaPath, 0o755)
	_ = os.MkdirAll(tmpDir, 0o755)

	db := mustInitDB()
	_db = db
	_client = client

	if keepPolling {
		printMessage = false
		for {
			func() {
				defer func() {
					if err := recover(); err != nil {
						logrus.WithError(fmt.Errorf("%v", err)).Error("polling")
					}
				}()
				poll()
			}()
			time.Sleep(pollingInterval)
		}
	} else {
		poll()
	}
}

func poll() {
	db := _db
	client := _client

	seq := sql.NullInt64{}
	db.Model(&models.Message{}).Select("max(sequence)").Scan(&seq)
	lastSequence := uint64(seq.Int64)
	logrus.WithFields(logrus.Fields{
		"last_sequence": lastSequence,
	}).Info("polling started")
	total := 0
	for {
		data, err := client.GetChatData(WeWorkFinanceSDK.GetChatDataOptions{
			Sequence: lastSequence,
			Limit:    1000,
			Timeout:  5,
		})
		if err != nil {
			panic(err)
		}
		if len(data) == 0 {
			break
		}
		total += len(data)
		lastSequence = data[len(data)-1].Sequence

		if printMessage {
			for _, v := range data {
				fmt.Println(v.Message)
				if v.MessageID != v.Message.GetID() {
					panic("message id not match")
				}
			}
		}

		var messages []*models.Message
		var allMedias []*models.Media
		var files []*models.File
		for _, v := range data {
			message := models.FromMessage(v.Message)
			medias := WeWorkFinanceSDK.GetMedias(v.Message)
			for _, media := range medias {
				if media.MD5Sum != "" {
					var last *models.File
					db.Where("md5_sum = ?", media.MD5Sum).Limit(1).Find(&last)
					if last.ID != "" {
						logrus.Infof("skip file %s already exists", media.MD5Sum)
						continue
					}
				}
			RETRY:
				data, err := client.ReadMediaData(WeWorkFinanceSDK.GetMediaDataOptions{
					FileID: media.ID,
				})
				if err != nil {
					panic(err)
				}
				if media.MD5Sum != "" {
					sum := WeWorkFinanceSDK.MD5Sum(data)
					if sum != media.MD5Sum {
						logrus.Warnf("md5sum not match, retrying %s", media.ID)
						time.Sleep(2 * time.Second)
						goto RETRY
					}
				}
				if err = media.VerifyData(data, nil); err != nil {
					panic(err)
				}

				file := models.FileFromMedia(media)
				files = append(files, file)
				allMedias = append(allMedias, models.FromMedia(media))
			}
			message.HasMedia = len(medias) > 0
			messages = append(messages, message)
		}

		if err = db.Clauses(clause.OnConflict{
			DoNothing: true,
		}).CreateInBatches(files, 10).Error; err != nil {
			panic(err)
		}
		if err = db.Clauses(clause.OnConflict{
			DoNothing: true,
		}).CreateInBatches(allMedias, 100).Error; err != nil {
			panic(err)
		}
		if err = db.Clauses(clause.OnConflict{
			DoNothing: true,
		}).CreateInBatches(messages, 100).Error; err != nil {
			panic(err)
		}
	}
	logrus.WithFields(logrus.Fields{
		"last_sequence": lastSequence,
		"pulled":        total,
	}).Info("polling done")
}

func mustInitDB() *gorm.DB {
	gsqlite.MustRegisterScalarFunction("gen_ulid", 0, func(ctx *gsqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
		return MustNewULID(), nil
	})
	db, err := gorm.Open(sqlite.Open("wwfinance.db?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	logrus.Info("db migration")
	err = db.AutoMigrate(&models.Message{}, &models.Media{}, &models.File{})
	if err != nil {
		panic(err)
	}
	return db
}
