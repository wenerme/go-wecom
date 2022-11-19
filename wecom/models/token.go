package models

import (
	"github.com/wenerme/go-wecom/commons/gorms"
	"time"
)

type Token struct {
	gorms.Model
	Type      string
	CorpID    string
	ExpiredAt *time.Time
}
