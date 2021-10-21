package wwcrypt

import (
	"crypto/cipher"
	"sort"
	"strings"
)

func Signature(token, timestamp, nonce, enc string) string {
	a := []string{token, timestamp, nonce, enc}
	sort.Strings(a)
	s := strings.Join(a, "")
	return sha1sum(s)
}

func encrypt(block cipher.BlockMode, in []byte) (out []byte, err error) {
	in, err = pkcs7pad(in, block.BlockSize())
	if err == nil {
		out = make([]byte, len(in))
		block.CryptBlocks(out, in)
	}
	return
}

func decrypt(block cipher.BlockMode, in []byte) (out []byte, err error) {
	out = make([]byte, len(in))
	block.CryptBlocks(out, in)
	out, err = pkcs7strip(out, block.BlockSize())
	return
}
