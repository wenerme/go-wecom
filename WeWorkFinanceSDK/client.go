package WeWorkFinanceSDK

// #cgo LDFLAGS: -L${SRCDIR}/libs -lWeWorkFinanceSdk_C -lm
// #cgo CFLAGS: -I${SRCDIR}/libs
// #include <stdlib.h>
// #include "WeWorkFinanceSdk_C.h"
import "C"

import (
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"sync"
	"unsafe"
)

type ClientOptions struct {
	Proxy           string
	ProxyCredential string
	Timeout         int
	PrivateKey      *rsa.PrivateKey
	PrivateKeyFn    func(ver int) (*rsa.PrivateKey, error)
}
type ClientOptionBuilder struct {
	Options ClientOptions
	Errors  []error
	Client  *Client
}

func (b *ClientOptionBuilder) PrivateKeys(m map[int]string) *ClientOptionBuilder {
	b.PrivateKeyFn(func(ver int) (string, error) {
		return m[ver], nil
	})
	return b
}

func (b *ClientOptionBuilder) PrivateKeyFn(fn func(ver int) (string, error)) *ClientOptionBuilder {
	keys := sync.Map{}
	b.Options.PrivateKeyFn = func(ver int) (key *rsa.PrivateKey, err error) {
		val, _ := keys.Load(ver)
		if key, _ = val.(*rsa.PrivateKey); key != nil {
			return
		}
		s, err := fn(ver)
		if err != nil || s == "" {
			return nil, err
		}
		key, err = ParsePrivateKey(s)
		if err == nil {
			keys.Store(ver, key)
		}
		return
	}
	return b
}

func (b *ClientOptionBuilder) Proxy(value string) *ClientOptionBuilder {
	b.Options.Proxy = value
	return b
}

func (b *ClientOptionBuilder) ProxyCredential(value string) *ClientOptionBuilder {
	b.Options.ProxyCredential = value
	return b
}

func (b *ClientOptionBuilder) ParseEnv() *ClientOptionBuilder {
	if b.Options.PrivateKey == nil {
		pkFile := os.Getenv("WWF_PRIVATE_KEY_FILE")
		pk := os.Getenv("WWF_PRIVATE_KEY")

		if pkFile != "" && pk == "" {
			file, err := os.ReadFile(pkFile)
			if err != nil {
				b.Errors = append(b.Errors, err)
				return b
			}
			pk = string(file)
		}
		if pk != "" {
			b.PrivateKey(pk)
		}
	}
	if b.Options.Proxy == "" {
		b.Proxy(os.Getenv("WWF_PROXY"))
	}
	if b.Options.ProxyCredential == "" {
		b.ProxyCredential(os.Getenv("WWF_PROXY_CREDENTIAL"))
	}
	if b.Options.Timeout == 0 {
		b.Options.Timeout, _ = strconv.Atoi(os.Getenv("WWF_TIMEOUT"))
	}
	return b
}

func (b *ClientOptionBuilder) PrivateKey(v string) *ClientOptionBuilder {
	key, err := ParsePrivateKey(v)
	if err != nil {
		b.Errors = append(b.Errors, err)
	} else {
		b.Options.PrivateKey = key
	}
	return b
}

func (b *ClientOptionBuilder) Apply() (*Client, error) {
	options, err := b.Build()
	if err != nil {
		return nil, err
	}
	b.Client.options = options
	return b.Client, nil
}

func (b *ClientOptionBuilder) MustApply() *Client {
	client, err := b.Apply()
	if err != nil {
		panic(err)
	}
	return client
}

func (b *ClientOptionBuilder) MustBuild() ClientOptions {
	options, err := b.Build()
	if err != nil {
		panic(err)
	}
	return options
}

func (b *ClientOptionBuilder) Build() (ClientOptions, error) {
	if len(b.Errors) > 0 {
		return ClientOptions{}, b.Errors[0]
	}
	o := b.Options
	if o.PrivateKey == nil && o.PrivateKeyFn == nil {
		b.Errors = append(b.Errors, ErrorOfCode(10000, "no private key"))
	}
	return o, nil
}

func NewClientFromEnv() (*Client, error) {
	corpID, _ := os.LookupEnv("WWF_CORP_ID")
	corpSecret, _ := os.LookupEnv("WWF_CORP_SECRET")
	corpSecretFile, _ := os.LookupEnv("WWF_CORP_SECRET_FILE")
	if corpSecretFile != "" && corpSecret == "" {
		file, err := os.ReadFile(corpSecretFile)
		if err != nil {
			panic(err)
		}
		corpSecret = string(file)
	}
	if corpID == "" {
		return nil, fmt.Errorf("corpId not founed from env")
	}
	if corpSecret == "" {
		return nil, fmt.Errorf("corpSecret not founed from env")
	}
	client, err := NewClient(corpID, corpSecret)
	if err == nil {
		client, err = client.Options().ParseEnv().Apply()
	}
	return client, err
}

func NewClient(corpID string, corpSecret string) (*Client, error) {
	ptr := C.NewSdk()
	corpIDC := C.CString(corpID)
	corpSecretC := C.CString(corpSecret)
	defer func() {
		C.free(unsafe.Pointer(corpIDC))
		C.free(unsafe.Pointer(corpSecretC))
	}()
	retC := C.Init(ptr, corpIDC, corpSecretC)
	ret := int(retC)
	if ret != 0 {
		return nil, ErrorOfCode(ret, "Init")
	}
	return &Client{
		ptr: ptr,
	}, nil
}

type Client struct {
	ptr     *C.WeWorkFinanceSdk_t
	options ClientOptions
}

func (c *Client) Options() *ClientOptionBuilder {
	return &ClientOptionBuilder{Options: c.options, Client: c}
}

type GetChatDataOptions struct {
	Sequence        int64
	Limit           int64
	Proxy           string
	ProxyCredential string
	Timeout         int
	SkipDecrypt     bool
}

func (c *Client) GetChatData(o GetChatDataOptions) ([]*ChatData, error) {
	if o.Proxy == "" {
		o.Proxy = c.options.Proxy
	}
	if o.ProxyCredential == "" {
		o.ProxyCredential = c.options.ProxyCredential
	}
	if o.Timeout == 0 {
		o.Timeout = c.options.Timeout
	}

	proxyC := C.CString(o.Proxy)
	passwdC := C.CString(o.ProxyCredential)
	chatSlice := C.NewSlice()
	defer func() {
		C.free(unsafe.Pointer(proxyC))
		C.free(unsafe.Pointer(passwdC))
		C.FreeSlice(chatSlice)
	}()

	if c.ptr == nil {
		return nil, ErrorOfCode(10002, "Client is closed")
	}

	retC := C.GetChatData(c.ptr, C.ulonglong(o.Sequence), C.uint(o.Limit), proxyC, passwdC, C.int(o.Timeout), chatSlice)
	ret := int(retC)
	if ret != 0 {
		return nil, ErrorOfCode(ret, "GetChatData")
	}
	buf := getContentFromSlice(chatSlice)
	var data ChatRawData
	err := json.Unmarshal(buf, &data)
	if err != nil {
		return nil, err
	}
	if data.Code != 0 {
		return nil, data.Error
	}

	if !o.SkipDecrypt {
		for _, v := range data.ChatDataList {
			key := c.options.PrivateKey
			if c.options.PrivateKeyFn != nil {
				key, err = c.options.PrivateKeyFn(v.PublicKeyVersion)
				if err != nil {
					return nil, err
				}
			}
			if key == nil {
				return nil, ErrorOfCode(10003, "PrivateKey not found")
			}
			v.Message, err = DecryptData(key, v.EncryptRandomKey, v.EncryptChatMessage)
			if err != nil {
				return nil, err
			}
		}
	}

	return data.ChatDataList, nil
}

func (c *Client) Close() {
	if c.ptr == nil {
		return
	}
	C.DestroySdk(c.ptr)
	c.ptr = nil
}

func getContentFromSlice(slice *C.struct_Slice_t) []byte {
	return C.GoBytes(unsafe.Pointer(C.GetContentFromSlice(slice)), C.GetSliceLen(slice))
}

func DecryptData(privateKey *rsa.PrivateKey, encryptRandomKey string, encryptMsg string) (msg Message, err error) {
	encryptKey, err := RSADecryptBase64(privateKey, encryptRandomKey)
	if err != nil {
		return msg, err
	}
	encryptKeyC := C.CString(string(encryptKey))
	encryptMsgC := C.CString(encryptMsg)
	msgSlice := C.NewSlice()
	defer func() {
		C.free(unsafe.Pointer(encryptKeyC))
		C.free(unsafe.Pointer(encryptMsgC))
		C.FreeSlice(msgSlice)
	}()

	retC := C.DecryptData(encryptKeyC, encryptMsgC, msgSlice)
	ret := int(retC)
	if ret != 0 {
		return msg, ErrorOfCode(ret, "DecryptData")
	}
	buf := getContentFromSlice(msgSlice)

	// handle illegal escape character in text
	for i := 0; i < len(buf); {
		if buf[i] < 0x20 {
			buf = append(buf[:i], buf[i+1:]...)
			continue
		}
		i++
	}

	var getType struct {
		Type   string `json:"msgtype,omitempty"`
		Action string `json:"action,omitempty"`
	}
	err = json.Unmarshal(buf, &getType)
	if err != nil {
		return msg, err
	}
	msg = MessageOfType(getType.Action, getType.Type)
	err = json.Unmarshal(buf, msg)
	msg.SetRaw(buf)
	return msg, err
}
