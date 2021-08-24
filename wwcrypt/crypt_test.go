package wwcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"net/url"
	"sort"
	"strings"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/wenerme/go-wecom/wecom"

	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"
	"github.com/stretchr/testify/assert"
)

func TestManualDecrypt(t *testing.T) {
	// 建立连接 - 测试回调模式
	// https://open.work.weixin.qq.com/wwopen/devtool/interface/combine
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
	r := wecom.ServiceGenericCallbackRequest{
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

func TestCrypto(t *testing.T) {
	urls := []string{
		`http://test?msg_signature=6a27598fb5b01ee02684379a339fef6de65c83c2&timestamp=1629824456&nonce=bqe3yfbqfc7&echostr=22vyvcBQ1xafTSJuhMGQuFyOAh00m2enIATgo4R6SGHY%2FYhofNmfwCksP3OP0Gxnq%2BcqnJcuvUQtjSbEu60IUA%3D%3D`,
		`http://test?msg_signature=e225d26d3bb4d06b4656f6b674e2143d4d35e81e&timestamp=1629822102&nonce=oduw1lq7iwp&echostr=teMozE0aNgDd2iKay6VOldGRKQXJbHsFqFc10Y4dYBE%2FMu%2B34X0Vd%2BKP6GoxWXGqYOC9bg4eSNw7d3d6awhm%2BA%3D%3D`,
		`http://test?msg_signature=c2273fb8ef29b1427a88b76b56036b76128c46d0&timestamp=1629824247&nonce=s8a61euba6m&echostr=fSW%2B704DFd%2BeE1SrOhMQ1kiV%2Bt%2F5V7%2Bcs8QIt6qD9xR44PgSlym1NtMvQYw5S%2BRKhQxdy3RNheLNHJuU%2BHGjHQ%3D%3D`,
	}
	c := &Crypto{
		ReceiveID:      "wx5823bf96d3bd56c7",
		Token:          "1372623149",
		EncodingAESKey: "jWmYm7qr5nMoAUwZRjGtBxmz3KA1tkAj3ykkR6q2B2C",
	}

	for _, test := range urls {
		u, err := url.Parse(test)
		if err != nil {
			panic(err)
		}
		r := wecom.ServiceGenericCallbackRequest{}
		assert.NoError(t, DecodeURLValues(u.Query(), &r))
		assert.NoError(t, err)
		fmt.Println("Echo String", r.EchoString)
		fmt.Println("Signature", r.MessageSignature)
		fmt.Println("Timestamp", r.Timestamp)
		fmt.Println("Nonce", r.Nonce)

		assert.Equal(t, r.MessageSignature, c.Signature(r.Timestamp, r.Nonce, r.EchoString))
		msg, err := c.Verify(r.MessageSignature, r.Timestamp, r.Nonce, r.EchoString)
		assert.NoError(t, err)
		assert.Equal(t, "ECHO", string(msg))

		enc, err := c.EncryptMessage(msg)
		assert.NoError(t, err)
		dec, err := c.DecryptMessage(enc)
		assert.NoError(t, err)
		assert.Equal(t, msg, dec)
	}

	c.Reset()
}

func DecodeURLValues(values url.Values, out interface{}) error {
	m := map[string]interface{}{}
	for k, v := range values {
		if len(v) == 0 {
			continue
		}
		switch len(v) {
		case 0:
		case 1:
			m[k] = v[0]
		default:
			m[k] = v
		}
	}
	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           out,
		WeaklyTypedInput: true,
		TagName:          "json",
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(m)
}
