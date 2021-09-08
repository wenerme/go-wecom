package wecom

import (
	"crypto/rand"
	"encoding/base32"
	"encoding/base64"
	mathrand "math/rand"
)

var nonceEncoder = base32.NewEncoding("abcdefghijklmnopqrstuvwxyz234567")

func createNonce() string {
	// no pad
	return nonceEncoder.EncodeToString(randBytes(15))
}

func randBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		// err is always nil
		_, _ = mathrand.Read(b) // nolint:gosec
	}
	return b
}

// RandAES256Key generate a random base64 aes key without padding
func RandAES256Key() string {
	return base64.RawStdEncoding.EncodeToString(randBytes(32))
}
