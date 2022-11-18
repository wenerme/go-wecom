package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/wenerme/go-wecom/WeWorkFinanceSDK"
	"gorm.io/gorm"
)

type Model struct {
	ID        string         `gorm:"primaryKey;default:(gen_ulid())" sql:"DEFAULT:gen_ulid()"` // sqlite must define as (gen_ulid())
	CreatedAt time.Time      `gorm:"index" sql:"type:timestamptz; DEFAULT:current_timestamp"`
	UpdatedAt time.Time      `gorm:"index" sql:"type:timestamptz; DEFAULT:current_timestamp"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	DB        *gorm.DB       `gorm:"-" mapstructure:"-" json:"-" yaml:"-" copier:"-"`
}

func (model *Model) GetDB() *gorm.DB {
	return model.DB
}

func (model *Model) AfterFind(tx *gorm.DB) (err error) {
	model.DB = tx
	return
}

func (model *Model) BeforeSave(tx *gorm.DB) (err error) {
	model.DB = tx
	return
}

type Message struct {
	Model
	CorpID      string    `gorm:"uniqueIndex:messages_corp_id_sequence"`
	MessageID   string    `gorm:"unique"`
	Sequence    uint64    `gorm:"uniqueIndex:messages_corp_id_sequence"`
	MessageTime time.Time `gorm:"index"`
	Type        string
	Action      string
	FromID      string
	ToID        string
	RoomID      string
	Raw         json.RawMessage
	HasMedia    bool
	Medias      []*Media `gorm:"foreignKey:MessageID;references:MessageID"`
}

func FromMessage(mm WeWorkFinanceSDK.Message) *Message {
	x := &Message{}
	x.FromMessage(mm)
	return x
}

func (m *Message) FromMessage(mm WeWorkFinanceSDK.Message) {
	m.MessageID = mm.GetID()
	m.Type = mm.GetType()
	m.Action = mm.GetAction()
	m.Raw = mm.GetRaw()
	m.Sequence = mm.GetSequence()
	m.CorpID = mm.GetCorpID()
	m.MessageTime = mm.GetTime()

	if iface, _ := mm.(WeWorkFinanceSDK.HasBaseMessage); iface != nil {
		base := iface.GetBaseMessage()
		m.FromID = base.From
		m.RoomID = base.RoomID
		if base.RoomID == "" {
			m.ToID = base.ToList[0]
		}
	}
}

func (m *Message) FromMedias(medias []*WeWorkFinanceSDK.Media) {
	m.HasMedia = len(medias) > 0
	for _, media := range medias {
		m.Medias = append(m.Medias, FromMedia(media))
	}
}

type Media struct {
	Model
	CorpID         string `gorm:"uniqueIndex:media_corp_id_sequence_index"`
	MessageID      string `gorm:"index"`
	Sequence       uint64 `gorm:"uniqueIndex:media_corp_id_sequence_index"`
	Index          int    `gorm:"uniqueIndex:media_corp_id_sequence_index"`
	FileName       string
	FileSize       int
	Length         int
	Width          int
	Height         int
	FileID         string
	Ext            string
	MD5Sum         string `gorm:"index"`
	OriginalMD5Sum string `gorm:"index"` // 写入的文件 MD5
}

type File struct {
	Model
	CorpID         string `gorm:"uniqueIndex:files_corp_id_md5sum"`
	MD5Sum         string `gorm:"index;uniqueIndex:files_corp_id_md5sum;not null;default:null"`
	OriginalMD5Sum string `gorm:"index"`
	Name           string
	Ext            string
	Width          int
	Height         int
	Length         int
	Size           int
	ContentType    string
	Content        []byte
}

func FileFromMedia(mm *WeWorkFinanceSDK.Media) *File {
	f := &File{}
	f.FromMedia(mm)
	return f
}

func (m *File) FromMedia(mm *WeWorkFinanceSDK.Media) {
	if mm.Message != nil {
		m.CorpID = mm.Message.GetCorpID()
	}
	m.Name = mm.Name
	m.Ext = mm.Ext
	m.Width = mm.Width
	m.Height = mm.Height
	m.Length = mm.Length
	m.Size = mm.Size
	m.MD5Sum = mm.MD5Sum
	m.OriginalMD5Sum = mm.OriginalMD5Sum
	if m.OriginalMD5Sum == m.MD5Sum {
		m.OriginalMD5Sum = ""
	}
	m.Content = mm.Data
	if len(m.Content) > 0 {
		m.ContentType = http.DetectContentType(m.Content)
	}
}

func FromMedia(mm *WeWorkFinanceSDK.Media) (m *Media) {
	m = &Media{}
	m.FromMedia(mm)
	return
}

func (m *Media) FromMedia(mm *WeWorkFinanceSDK.Media) {
	m.FileID = mm.ID
	m.Ext = mm.Ext
	m.MessageID = mm.MessageID
	m.Index = mm.Index
	m.FileName = mm.Name
	m.FileSize = mm.Size
	m.Length = mm.Length
	m.Width = mm.Width
	m.Height = mm.Height
	m.MD5Sum = mm.MD5Sum
	m.OriginalMD5Sum = mm.OriginalMD5Sum
	if mm.Message != nil {
		m.Sequence = mm.Message.GetSequence()
		m.CorpID = mm.Message.GetCorpID()
	}
}
