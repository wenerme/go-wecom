package wwcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
)

// AES Helper
type AES struct {
	Duplicate bool
	cipher    cipher.Block
	dec       cipher.BlockMode
	enc       cipher.BlockMode
}

// Decrypt with strip
func (c AES) Decrypt(bytes []byte) []byte {
	bytes = c.dup(bytes)

	c.dec.CryptBlocks(bytes, bytes)
	o, err := pkcs7strip(bytes, aes.BlockSize)
	// strip failed
	if err != nil {
		log.Println("pkcs7 strip decrypt failed:", err.Error())
	}
	return o
}

// Encrypt with padding
func (c AES) Encrypt(bytes []byte) []byte {
	bytes = c.dup(bytes)

	o, err := pkcs7pad(bytes, aes.BlockSize)
	// pad failed
	if err != nil {
		// CryptBlocks will panic
		log.Println("pkcs7 pad encrypt failed:", err.Error())
	}
	c.enc.CryptBlocks(o, o)
	return o
}

func (c AES) dup(bytes []byte) []byte {
	if c.Duplicate {
		bytes = dup(bytes)
	}
	return bytes
}

// NewAESFromEncodeKey create from base64 encode key
func NewAESFromEncodeKey(base64Key string) (c *AES, err error) {
	var key []byte
	key, err = base64.RawStdEncoding.DecodeString(base64Key)
	if err == nil {
		c, err = NewAESFromKey(key)
	}
	return
}

// NewAESFromKey create from raw key
func NewAESFromKey(key []byte) (c *AES, err error) {
	c = &AES{}
	c.cipher, err = aes.NewCipher(key)
	if err == nil {
		c.dec = cipher.NewCBCDecrypter(c.cipher, key[:aes.BlockSize])
		c.enc = cipher.NewCBCEncrypter(c.cipher, key[:aes.BlockSize])
	}
	return
}

func dup(p []byte) []byte {
	q := make([]byte, len(p))
	copy(q, p)
	return q
}
