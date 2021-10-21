package wwcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/xml"
)

func DecryptMessage(encodingAESKey string, post []byte) (msg *ReceiveMessage, err error) {
	msg = &ReceiveMessage{}
	err = xml.Unmarshal(post, msg)
	if err != nil {
		return
	}
	enc, err := base64.StdEncoding.DecodeString(msg.Encrypt)
	if err != nil {
		return
	}
	msg.Content, err = DecryptContent(encodingAESKey, enc)
	return
}

func DecryptContent(encodingAESKey string, enc []byte) (msg *ReceiveContent, err error) {
	var key []byte
	key, err = base64.RawStdEncoding.DecodeString(encodingAESKey)
	if err != nil {
		return nil, err
	}
	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	b := cipher.NewCBCDecrypter(block, key[:aes.BlockSize])
	out, err := decrypt(b, enc)
	if err != nil {
		return nil, err
	}
	msg = &ReceiveContent{}
	return msg, msg.UnmarshalBinary(out)
}
