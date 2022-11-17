//go:build linux && cgo

package WeWorkFinanceSDK

// #cgo LDFLAGS: -L${SRCDIR}/libs -lWeWorkFinanceSdk_C -lm
// #cgo CFLAGS: -I${SRCDIR}/libs
// #include <stdlib.h>
// #include "WeWorkFinanceSdk_C.h"
import "C"

import (
	"crypto/rsa"
	"encoding/json"
	"unsafe"
)

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

type Client struct {
	ptr     *C.WeWorkFinanceSdk_t
	options ClientOptions
}
