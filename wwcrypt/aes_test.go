package wwcrypt_test

import (
	"crypto/aes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wenerme/go-wecom/wecom"
	"github.com/wenerme/go-wecom/wwcrypt"
)

func TestAES(t *testing.T) {
	k := wecom.RandAES256Key()
	helper, err := wwcrypt.NewAESFromEncodeKey(k)
	assert.NoError(t, err)
	hello := []byte("1234567890123456")
	raw := dup(hello)
	enc := helper.Encrypt(raw)
	assert.NotEqual(t, hello, raw)
	assert.Equal(t, hello, helper.Decrypt(enc))
	assert.Equal(t, aes.BlockSize, len(enc))

	raw = dup(hello)
	helper.Duplicate = true
	helper.Encrypt(raw)
	assert.Equal(t, hello, raw)
}

func dup(p []byte) []byte {
	q := make([]byte, len(p))
	copy(q, p)
	return q
}
