package wecom

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
)

var nonceEncoder = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567")

func createNonce() string {
	return nonceEncoder.EncodeToString(randBytes(15))
}

func randBytes(n int) []byte {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return b
}

// RandAES256Key generate a random base64 aes key without padding
func RandAES256Key() string {
	return base64.RawStdEncoding.EncodeToString(randBytes(32))
}
