package models

import (
	"time"

	"github.com/wenerme/go-wecom/commons/gorms"
)

type Token struct {
	gorms.Model
	Type      string
	CorpID    string
	ExpiredAt *time.Time
}
