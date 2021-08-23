package wecom

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"
	"github.com/stretchr/testify/assert"
)

func TestManualDecrypt(t *testing.T) {
	EncodingAESKey := "jWmYm7qr5nMoAUwZRjGtBxmz3KA1tkAj3ykkR6q2B2C"
	// AESKey = Base64_Decode(EncodingAESKey + "=")
	AESKey, err := base64.RawStdEncoding.DecodeString(EncodingAESKey)
	assert.NoError(t, err)
	// dev_msg_signature=sha1(sort(token、timestamp、nonce、msg_encrypt))
	// rand_msg = random(16B) + msg_len(4B) + msg + receiveid
	msgEnc := "RypEvHKD8QQKFhvQ6QleEB4J58tiPdvo+rtK1I9qca6aM/wvqnLSV5zEPeusUiX5L5X/0lWfrf0QADHHhGd3QczcdCUpj911L3vg3W/sYYvuJTs3TUUkSUXxaccAS0qhxchrRYt66wiSpGLYL42aM6A8dTT+6k4aSknmPj48kzJs8qLjvd4Xgpue06DOdnLxAUHzM6+kDZ+HMZfJYuR+LtwGc2hgf5gsijff0ekUNXZiqATP7PF5mZxZ3Izoun1s4zG4LUMnvw2r+KqCKIw+3IQH03v+BCA9nMELNqbSf6tiWSrXJB3LAVGUcallcrw8V2t9EL4EhzJWrQUax5wLVMNS0+rUPA3k22Ncx4XXZS9o0MBH27Bo6BpNelZpS+/uh9KsNlY6bHCmJU9p8g7m3fVKn28H3KDYA5Pl/T8Z1ptDAVe0lXdQ2YoyyH2uyPIGHBZZIs2pDBS8R07+qN+E7Q=="
	// ReceiveId 企业应用的回调，表示 corpid；第三方事件的回调，表示 suiteid
	// Token 为接收消息 token
	token := "1372623149"
	// recvId := "wx5823bf96d3bd56c7"
	r := ServiceGenericCallbackRequest{
		Nonce:            "QDG6eK",
		Timestamp:        "1409659813",
		MessageSignature: "477715d11cdb4164915debcba66cb864d751f3e6",
		EchoString:       msgEnc,
	}

	a := []string{token, r.Timestamp, r.Nonce, msgEnc}
	sort.Strings(a)
	sortStr := "13726231491409659813QDG6eKRypEvHKD8QQKFhvQ6QleEB4J58tiPdvo+rtK1I9qca6aM/wvqnLSV5zEPeusUiX5L5X/0lWfrf0QADHHhGd3QczcdCUpj911L3vg3W/sYYvuJTs3TUUkSUXxaccAS0qhxchrRYt66wiSpGLYL42aM6A8dTT+6k4aSknmPj48kzJs8qLjvd4Xgpue06DOdnLxAUHzM6+kDZ+HMZfJYuR+LtwGc2hgf5gsijff0ekUNXZiqATP7PF5mZxZ3Izoun1s4zG4LUMnvw2r+KqCKIw+3IQH03v+BCA9nMELNqbSf6tiWSrXJB3LAVGUcallcrw8V2t9EL4EhzJWrQUax5wLVMNS0+rUPA3k22Ncx4XXZS9o0MBH27Bo6BpNelZpS+/uh9KsNlY6bHCmJU9p8g7m3fVKn28H3KDYA5Pl/T8Z1ptDAVe0lXdQ2YoyyH2uyPIGHBZZIs2pDBS8R07+qN+E7Q=="
	assert.Equal(t, sortStr, strings.Join(a, ""))
	assert.Equal(t, r.MessageSignature, sha1sum(sortStr))

	// aes_msg = base64_decode(msg_encrypt)
	aesMsg, err := base64.StdEncoding.DecodeString(msgEnc)
	assert.NoError(t, err)
	// rand_msg = aes_decrypt(aes_msg, AESKey)
	aesCipher, err := aes.NewCipher(AESKey)
	assert.NoError(t, err)
	dec := cipher.NewCBCDecrypter(aesCipher, AESKey[:aes.BlockSize])
	assert.NoError(t, err)
	dec.CryptBlocks(aesMsg, aesMsg)
	assert.NoError(t, err)
	// pkcs7 unpadding
	aesMsg = aesMsg[:len(aesMsg)-int(aesMsg[len(aesMsg)-1])]
	//
	content := aesMsg[16:]
	msgLen := binary.BigEndian.Uint32(content)
	msg := string(content[4 : msgLen+4])

	recvID := string(content[msgLen+4:])
	assert.Equal(t, "wx5823bf96d3bd56c7", recvID)

	fmt.Println(msg)
	wxCrypt := wxbizmsgcrypt.NewWXBizMsgCrypt(token, EncodingAESKey, recvID, wxbizmsgcrypt.XmlType)
	bytes, cerr := wxCrypt.VerifyURL(r.MessageSignature, r.Timestamp, r.Nonce, r.EchoString)
	assert.Nil(t, cerr)
	assert.Equal(t, msg, string(bytes))
}
