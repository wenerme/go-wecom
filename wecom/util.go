package wecom

import (
	"crypto/rand"
	"encoding/base32"
	"strconv"
	"time"
)

var encoder = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567")

func createNonce() string {
	// no pad
	b := make([]byte, 15)
	_, err := rand.Read(b)
	if err != nil {
		return strconv.Itoa(time.Now().Nanosecond())
	}
	return encoder.EncodeToString(b)
}
