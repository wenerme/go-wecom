package gorms

import (
	"gorm.io/gorm"
	"time"
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
