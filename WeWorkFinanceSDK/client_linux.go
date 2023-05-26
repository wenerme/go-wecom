//go:build linux && cgo

package WeWorkFinanceSDK

// #cgo LDFLAGS: -L${SRCDIR}/libs -lWeWorkFinanceSdk_C -lm
// #cgo CFLAGS: -I${SRCDIR}/libs
// #include <stdlib.h>
// #include "WeWorkFinanceSdk_C.h"
import "C"

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"unsafe"

	"github.com/pkg/errors"
)

type client struct {
	ptr     *C.WeWorkFinanceSdk_t
	corpID  string
	options ClientOptions
}

func NewClient(corpID string, corpSecret string) (Client, error) {
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
	return &client{
		ptr:    ptr,
		corpID: corpID,
	}, nil
}

func (c *client) GetChatData(o GetChatDataOptions) ([]*ChatData, error) {
	if o.Proxy == "" {
		o.Proxy = c.options.Proxy
	}
	if o.ProxyCredential == "" {
		o.ProxyCredential = c.options.ProxyCredential
	}
	if o.Timeout == 0 {
		o.Timeout = c.options.Timeout
	}
	if o.Unmarshal == nil {
		o.Unmarshal = Unmarshal
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
			var dec []byte
			dec, err = DecryptData(key, v.EncryptRandomKey, v.EncryptChatMessage)
			if err != nil {
				return nil, errors.Wrapf(err, "DecryptData failed, seq: %d", v.Sequence)
			}
			v.Message, err = o.Unmarshal(dec)
			v.Message.SetSequence(v.Sequence)
			v.Message.SetCorpID(c.corpID)
		}
	}

	return data.ChatDataList, nil
}

func (c *client) ReadMediaData(o GetMediaDataOptions) (b []byte, err error) {
	buf := bytes.Buffer{}
	if o.Proxy == "" {
		o.Proxy = c.options.Proxy
	}
	if o.ProxyCredential == "" {
		o.ProxyCredential = c.options.ProxyCredential
	}
	if o.Timeout == 0 {
		o.Timeout = c.options.Timeout
	}
	if c.ptr == nil {
		return nil, errors.New("client closed")
	}

	var indexBufC *C.char
	sdkFileIDC := C.CString(o.FileID)
	proxyC := C.CString(o.Proxy)
	passwdC := C.CString(o.ProxyCredential)
	mediaDataC := C.NewMediaData()
	defer func() {
		// C.free(unsafe.Pointer(indexBufC))
		C.free(unsafe.Pointer(sdkFileIDC))
		C.free(unsafe.Pointer(proxyC))
		C.free(unsafe.Pointer(passwdC))
		C.FreeMediaData(mediaDataC)
	}()

	for {

		retC := C.GetMediaData(c.ptr, indexBufC, sdkFileIDC, proxyC, passwdC, C.int(o.Timeout), mediaDataC)
		indexBufC = C.GetOutIndexBuf(mediaDataC)
		ret := int(retC)
		if ret != 0 {
			return nil, ErrorOfCode(ret, fmt.Sprintf("GetMediaData fileId=%s", o.FileID))
		}
		// 单次最大 512K
		_, _ = buf.Write(C.GoBytes(unsafe.Pointer(C.GetData(mediaDataC)), C.int(C.GetDataLen(mediaDataC))))
		if int(C.IsMediaDataFinish(mediaDataC)) == 1 {
			break
		}
	}
	return buf.Bytes(), err
}

func (c *client) CopyMediaData(o GetMediaDataOptions, w io.Writer) (sum int, err error) {
	o.Index = ""
	var data *MediaData
	var n int
	for {
		data, err = c.GetMediaData(o)
		if err != nil {
			return
		}
		n, err = w.Write(data.Data)
		sum += n
		if data.Finished {
			return
		}
		o.Index = data.OutIndexBuf
	}
}

func (c *client) GetMediaData(o GetMediaDataOptions) (*MediaData, error) {
	if o.Proxy == "" {
		o.Proxy = c.options.Proxy
	}
	if o.ProxyCredential == "" {
		o.ProxyCredential = c.options.ProxyCredential
	}
	if o.Timeout == 0 {
		o.Timeout = c.options.Timeout
	}
	indexBufC := C.CString(o.Index)
	sdkFileIDC := C.CString(o.FileID)
	proxyC := C.CString(o.Proxy)
	passwdC := C.CString(o.ProxyCredential)
	mediaDataC := C.NewMediaData()
	defer func() {
		C.free(unsafe.Pointer(indexBufC))
		C.free(unsafe.Pointer(sdkFileIDC))
		C.free(unsafe.Pointer(proxyC))
		C.free(unsafe.Pointer(passwdC))
		C.FreeMediaData(mediaDataC)
	}()

	if c.ptr == nil {
		return nil, errors.New("client closed")
	}

	retC := C.GetMediaData(c.ptr, indexBufC, sdkFileIDC, proxyC, passwdC, C.int(o.Timeout), mediaDataC)
	ret := int(retC)
	if ret != 0 {
		return nil, ErrorOfCode(ret, "GetMediaData")
	}
	return &MediaData{
		OutIndexBuf: C.GoString(C.GetOutIndexBuf(mediaDataC)),
		Data:        C.GoBytes(unsafe.Pointer(C.GetData(mediaDataC)), C.int(C.GetDataLen(mediaDataC))),
		Finished:    int(C.IsMediaDataFinish(mediaDataC)) == 1,
	}, nil
}

// func (c *client) SaveMedia(o SaveMediaOptions) error {
//	if o.TempDir == "" {
//		o.TempDir = c.options.TempDir
//	}
//	if o.TempDir == "" {
//		o.TempDir = os.TempDir()
//	}
//	var name string
//	media := o.Media
//	if media.MD5Sum != "" {
//		name = media.MD5Sum + "." + media.Ext
//	}
//	if name != "" {
//		if stat, _ := os.Stat(path.Join(o.Dir, name)); stat != nil && !o.Force {
//			if !o.KeepData && o.Writer == nil && o.Media.MD5Sum != "" {
//				return nil
//			}
//
//			file, err := os.ReadFile(stat.Name())
//			if err != nil {
//				return err
//			}
//			if o.KeepData {
//
//			}
//			return nil
//		}
//	}
//
//	temp, err := os.CreateTemp(o.TempDir, "weworkmedia")
//	if err != nil {
//		return err
//	}
//	skipClose := false
//	skipDelete := false
//	defer func() {
//		if !skipClose {
//			_ = temp.Close()
//		}
//		if !skipDelete {
//			//_ = os.Remove(temp.Name())
//		}
//	}()
//
//	hash := md5.New()
//	counter := &counterWriter{}
//
//	var writers = []io.Writer{hash, counter}
//	if o.Writer != nil {
//		writers = append(writers, o.Writer)
//	}
//	var buf *bytes.Buffer
//	if o.KeepData {
//		buf = new(bytes.Buffer)
//		writers = append(writers, buf)
//	}
//
//	w := io.MultiWriter(writers...)
//	_, err = c.CopyMediaData(GetMediaDataOptions{
//		FileID: media.ID,
//	}, w)
//	if err != nil {
//		return err
//	}
//
//	skipClose = true
//	_ = temp.Close()
//
//	md5sum := hex.EncodeToString(hash.Sum(nil))
//	if media.MD5Sum != "" {
//		if media.MD5Sum != md5sum && o.CheckSum {
//			println("temp", temp.Name())
//			return fmt.Errorf("media md5sum not match: %s != %s #%v <-> #%v id=%s", media.MD5Sum, md5sum, media.Size, counter.Count, media.ID)
//		}
//	} else {
//		media.MD5Sum = md5sum
//	}
//
//	if media.Size != 0 {
//		if media.Size != counter.Count {
//			return fmt.Errorf("media size not match: %d != %d id=%s", media.Size, counter.Count, media.ID)
//		}
//	}
//	if media.Ext == "" {
//		media.Ext = "data"
//	}
//
//	media.FileMD5Sum = md5sum
//	name = md5sum + "." + media.Ext
//	if buf != nil {
//		media.Data = buf.Bytes()
//	}
//	//
//	err = os.Rename(temp.Name(), path.Join(o.Dir, name))
//	skipDelete = err == nil
//
//	return err
// }

func (c *client) Close() {
	if c.ptr == nil {
		return
	}
	C.DestroySdk(c.ptr)
	c.ptr = nil
}

func getContentFromSlice(slice *C.struct_Slice_t) []byte {
	return C.GoBytes(unsafe.Pointer(C.GetContentFromSlice(slice)), C.GetSliceLen(slice))
}

func DecryptData(privateKey *rsa.PrivateKey, encryptRandomKey string, encryptMsg string) (msg []byte, err error) {
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

	return buf, err
}

// type counterWriter struct {
// 	Count int
// }
//
// func (c *counterWriter) Write(p []byte) (n int, err error) {
// 	n = len(p)
// 	c.Count += n
// 	return n, nil
// }
