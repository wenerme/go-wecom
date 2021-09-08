package wwcrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	//nolint:gosec
	"encoding/base64"
	"encoding/binary"
	"errors"
	"sort"
	"strings"
)

type Crypto struct {
	ReceiveID      string
	Token          string
	EncodingAESKey string

	aesKey    []byte
	decrypter cipher.BlockMode
	encrypter cipher.BlockMode
}

func (c *Crypto) Reset() {
	c.aesKey = nil
	c.decrypter = nil
}

func (c *Crypto) EncryptMessage(dec []byte) (enc []byte, err error) {
	buf := &bytes.Buffer{}
	r := make([]byte, 16)
	_, err = rand.Read(r)
	if err != nil {
		return nil, err
	}
	_, _ = buf.Write(r)

	binary.BigEndian.PutUint32(r, uint32(len(dec)))
	_, _ = buf.Write(r[:4])
	_, _ = buf.Write(dec)
	_, _ = buf.Write([]byte(c.ReceiveID))

	return c.Encrypt(buf.Bytes())
}

func (c *Crypto) DecryptMessage(enc []byte) ([]byte, error) {
	dec, err := c.Decrypt(enc)
	if err != nil {
		return nil, err
	}
	// random
	dec = dec[16:]
	size := binary.BigEndian.Uint32(dec)
	dec = dec[4:]
	receiveID := string(dec[size:])
	if c.ReceiveID != receiveID {
		return nil, errors.New("receive id not equal")
	}
	return dec[:size], nil
}

func (c *Crypto) Verify(signature, timestamp, nonce, echo string) ([]byte, error) {
	if c.Signature(timestamp, nonce, echo) != signature {
		return nil, errors.New("signature not equal")
	}
	enc, err := base64.StdEncoding.DecodeString(echo)
	if err != nil {
		return nil, err
	}
	return c.DecryptMessage(enc)
}

func (c *Crypto) Signature(timestamp, nonce, enc string) string {
	a := []string{c.Token, timestamp, nonce, enc}
	sort.Strings(a)
	s := strings.Join(a, "")
	return sha1sum(s)
}

func (c *Crypto) Encrypt(in []byte) (out []byte, err error) {
	encrypter, err := c.GetEncrypter()
	if err == nil {
		in, err = pkcs7pad(in, encrypter.BlockSize())
	}
	if err == nil {
		out = make([]byte, len(in))
		encrypter.CryptBlocks(out, in)
	}
	return
}

func (c *Crypto) Decrypt(in []byte) (out []byte, err error) {
	block, err := c.GetDecrypter()
	if err != nil {
		return nil, err
	}
	out = make([]byte, len(in))
	block.CryptBlocks(out, in)
	out, err = pkcs7strip(out, block.BlockSize())
	return
}

func (c *Crypto) GetAESKey() (key []byte, err error) {
	if len(c.aesKey) == 0 {
		c.aesKey, err = base64.RawStdEncoding.DecodeString(c.EncodingAESKey)
	}
	return c.aesKey, err
}

func (c *Crypto) initCipher() error {
	key, err := c.GetAESKey()
	if err == nil {
		var block cipher.Block
		block, err = aes.NewCipher(key)
		if err == nil {
			c.encrypter = cipher.NewCBCEncrypter(block, key[:aes.BlockSize])
			c.decrypter = cipher.NewCBCDecrypter(block, key[:aes.BlockSize])
		}
	}
	return err
}

func (c *Crypto) GetDecrypter() (cipher.BlockMode, error) {
	if c.decrypter != nil {
		return c.decrypter, nil
	}
	err := c.initCipher()
	return c.decrypter, err
}

func (c *Crypto) GetEncrypter() (cipher.BlockMode, error) {
	if c.encrypter != nil {
		return c.encrypter, nil
	}
	err := c.initCipher()
	return c.encrypter, err
}
