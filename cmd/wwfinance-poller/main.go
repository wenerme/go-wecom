package main

import (
	"crypto/rand"
	"database/sql"
	"database/sql/driver"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"path"
	"strings"
	"time"

	dotenv "github.com/joho/godotenv"

	"gorm.io/gorm/clause"

	gsqlite "github.com/glebarez/go-sqlite"
	"github.com/wenerme/go-wecom/WeWorkFinanceSDK/models"
	"gorm.io/gorm"

	"github.com/glebarez/sqlite"
	"github.com/oklog/ulid/v2"
	"github.com/wenerme/go-wecom/WeWorkFinanceSDK"
	"gorm.io/driver/postgres"
)

var defaultEntropySource = ulid.Monotonic(rand.Reader, 0)

var (
	fileID          string
	keepPolling     bool
	pollingInterval = time.Minute * 5
	printMessage    = true
	migrate         = true
	migrateOnly     = false
	_db             *gorm.DB
	_client         WeWorkFinanceSDK.Client
	envFile         = ""
)

func MustNewULID() string {
	n := ulid.MustNew(ulid.Timestamp(time.Now()), defaultEntropySource)
	return strings.ToLower(n.String())
}

func main() {
	flag.StringVar(&fileID, "file-id", "", "file id to pull")
	flag.StringVar(&envFile, "env-file", envFile, "load env from file")
	flag.BoolVar(&keepPolling, "polling", keepPolling, "keep polling")
	flag.BoolVar(&migrate, "migrate", migrate, "run migrate and exit")
	flag.BoolVar(&migrateOnly, "migrate-only", migrateOnly, "run migrate and exit")
	flag.Parse()
	if envFile == "" {
		envFile = os.Getenv("ENV_FILE")
	}
	if envFile == "" {
		envFile = ".env"
	}

	slog.Info("load env", "env_file", envFile)
	if err := dotenv.Load(strings.Split(envFile, ",")...); err != nil {
		slog.Warn("load env failed", "error", err)
	}

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
						slog.Error("polling", "error", fmt.Errorf("%v", err))
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
	slog.Info("polling started", "last_sequence", lastSequence)
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
		save(data)
	}
	slog.Info("polling done", "last_sequence", lastSequence, "pulled", total)
}

func save(chatMessages []*WeWorkFinanceSDK.ChatData) {
	db := _db
	client := _client

	var err error
	var messages []*models.Message
	var allMedias []*models.Media
	var files []*models.File
	for _, v := range chatMessages {
		message := models.FromMessage(v.Message)
		medias := WeWorkFinanceSDK.GetMedias(v.Message)
		for _, media := range medias {
			allMedias = append(allMedias, models.FromMedia(media))
			if media.MD5Sum != "" {
				var last *models.File
				db.Select("id").Where("md5_sum = ?", media.MD5Sum).Limit(1).Find(&last)
				if last.ID != "" {
					slog.Info("skip file, already exists", "md5", media.MD5Sum)
					continue
				}
			}
			var data []byte
		RETRY:
			data, err = client.ReadMediaData(WeWorkFinanceSDK.GetMediaDataOptions{
				FileID: media.ID,
			})
			if err != nil {
				panic(err)
			}
			if media.MD5Sum != "" {
				sum := WeWorkFinanceSDK.MD5Sum(data)
				if sum != media.MD5Sum {
					slog.Warn("md5sum not match, retrying", "expect_size", media.Size, "actual_size", len(data), "id", media.ID)
					time.Sleep(2 * time.Second)
					goto RETRY
				}
			}
			if err = media.VerifyData(data, nil); err != nil {
				panic(err)
			}

			file := models.FileFromMedia(media)
			files = append(files, file)
		}
		message.HasMedia = len(medias) > 0
		messages = append(messages, message)
	}

	if err = db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).CreateInBatches(files, 10).Error; err != nil {
		panic(err)
	}
	// works with fk
	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.Clauses(clause.OnConflict{
			DoNothing: true,
		}).CreateInBatches(allMedias, 100).Error; err != nil {
			return err
		}
		if err = tx.Clauses(clause.OnConflict{
			DoNothing: true,
		}).CreateInBatches(messages, 100).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func mustInitDB() *gorm.DB {
	gsqlite.MustRegisterScalarFunction("gen_ulid", 0, func(ctx *gsqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
		return MustNewULID(), nil
	})
	typ := os.Getenv("DB_TYPE")
	dsn := os.Getenv("DB_DSN")
	_ = typ
	var db *gorm.DB
	var err error
	if typ == "" {
		typ = "sqlite"
		if dsn == "" {
			dsn = "data/wwfinance.db?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)"
		}
	}

	switch typ {
	case "sqlite", "sqlite3":
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	case "pg", "postgres", "postgresql":
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		})
	default:
		panic("unsupported db type " + typ)
	}

	if err != nil {
		panic(err)
	}
	slog.Info("db migration", "db_type", typ)
	if migrate {
		err = db.AutoMigrate(&models.Message{}, &models.Media{}, &models.File{})
		if err != nil {
			panic(err)
		}
	}
	if migrateOnly {
		os.Exit(0)
	}
	return db
}
