package WeWorkFinanceSDK

import (
	"fmt"
)

type MediaData struct {
	OutIndexBuf string `json:"outindexbuf,omitempty"`
	Finished    bool   `json:"is_finish,omitempty"`
	Data        []byte `json:"data,omitempty"`
}
type HasMedias interface {
	GetMedias() []Media
}

type Media struct {
	ID             string // fileid
	Name           string // filename
	Ext            string // png, gif
	Width          int
	Height         int
	Size           int    // file size
	Length         int    // audio, video
	MD5Sum         string // 消息中的 MD5
	OriginalMD5Sum string // 写入文件的 MD5 - 可能不一样，暂不知道为什么

	Message   Message `json:"-"` //  track back
	MessageID string
	Index     int
	Data      []byte
}

type MediaVerifyOptions struct {
	SkipChecksum bool
}

func (m *Media) VerifyData(data []byte, o *MediaVerifyOptions) error {
	if o == nil {
		o = &MediaVerifyOptions{}
	}
	if m.Size != 0 && len(data) != m.Size {
		return fmt.Errorf("size not match: %d != %d", len(data), m.Size)
	}
	md5sum := MD5Sum(data)
	if m.MD5Sum != "" && m.MD5Sum != md5sum {
		if !o.SkipChecksum {
			return fmt.Errorf("md5sum not match: %s != %s", m.MD5Sum, md5sum)
		}
		m.OriginalMD5Sum = m.MD5Sum
	}

	m.Size = len(data)
	m.MD5Sum = md5sum
	m.Data = data
	return nil
}

func GetMedias(messages ...Message) (o []*Media) {
	for _, message := range messages {
		iface, _ := message.(HasMedias)
		if iface == nil {
			continue
		}
		for _, media := range iface.GetMedias() {
			if media.ID != "" {
				media.MessageID = message.GetID()
				media.Message = message
				m := media
				o = append(o, &m)
			}
		}
		if len(o) > 1 {
			for n := range o {
				o[n].Index = n + 1
			}
		}
	}
	return
}
