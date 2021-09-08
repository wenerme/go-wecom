package wwcrypt

import (
	"bytes"
	"crypto/sha1" // nolint:gosec
	"encoding/hex"
	"errors"
	"fmt"
)

// pkcs7strip remove pkcs7 padding
func pkcs7strip(data []byte, blockSize int) ([]byte, error) {
	// https://gist.github.com/nanmu42/b838acc10d393bc51cb861128ce7f89c
	length := len(data)
	if length == 0 {
		return data, errors.New("pkcs7: Data is empty")
	}
	if length%blockSize != 0 {
		return data, errors.New("pkcs7: Data is not block-aligned")
	}
	padLen := int(data[length-1])
	ref := bytes.Repeat([]byte{byte(padLen)}, padLen)
	// no pad
	if padLen == 0 || !bytes.HasSuffix(data, ref) {
		return data, nil
	}
	return data[:length-padLen], nil
}

// pkcs7pad add pkcs7 padding
func pkcs7pad(data []byte, blockSize int) ([]byte, error) {
	if blockSize < 0 || blockSize > 256 {
		return data, fmt.Errorf("pkcs7: Invalid block size %d", blockSize)
	}
	padLen := len(data) % blockSize
	if padLen == 0 {
		return data, nil
	}
	padLen = blockSize - padLen
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)
	return append(data, padding...), nil
}

func sha1sum(s string) string {
	sum := sha1.Sum([]byte(s)) //nolint:gosec
	return hex.EncodeToString(sum[:])
}
