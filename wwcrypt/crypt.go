package wwcrypt

import (
	"crypto/aes"
	"crypto/cipher"

	//nolint:gosec
	"encoding/base64"
	"errors"
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
	m := &ReceiveContent{
		ReceiverID: c.ReceiveID,
		Content:    dec,
	}
	data, err := m.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return c.Encrypt(data)
}

func (c *Crypto) DecryptMessage(enc []byte) ([]byte, error) {
	dec, err := c.Decrypt(enc)
	if err != nil {
		return nil, err
	}
	m := &ReceiveContent{}
	err = m.UnmarshalBinary(dec)
	if err != nil {
		return nil, err
	}
	if c.ReceiveID != m.ReceiverID {
		return nil, errors.New("receive id not equal")
	}
	return m.Content, nil
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
	return Signature(c.Token, timestamp, nonce, enc)
}

func (c *Crypto) Encrypt(in []byte) (out []byte, err error) {
	encrypter, err := c.GetEncrypter()
	if err != nil {
		return nil, err
	}
	return encrypt(encrypter, in)
}

func (c *Crypto) Decrypt(in []byte) (out []byte, err error) {
	block, err := c.GetDecrypter()
	if err != nil {
		return nil, err
	}
	return decrypt(block, in)
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
