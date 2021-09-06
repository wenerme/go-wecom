package wecom

import "encoding/xml"

/*
第三方回调协议
https://work.weixin.qq.com/api/doc/10982
*/

type EncryptPushEvent struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string   `xml:"ToUserName,omitempty"`
	Encrypt    string   `xml:"Encrypt,omitempty"`
	AgentID    string   `xml:"AgentID,omitempty"`
}

// PushRequestParams 服务商回调请求 query 参数
type PushRequestParams struct {
	MessageSignature string `json:"msg_signature"`
	Timestamp        string `json:"timestamp"`
	Nonce            string `json:"nonce"`
	EchoString       string `json:"echostr,omitempty"` //  验证 URL 时包含
}

type CommonPushEvent struct {
	XMLName xml.Name `xml:"xml"`

	// 自建应用

	ToUsername   string `xml:"ToUserName,omitempty"`   // 企业微信CorpID
	FromUsername string `xml:"FromUserName,omitempty"` // sys 系统生成
	CreateTime   int64  `xml:"CreateTime,omitempty"`   // same as Timestamp
	MsgType      string `xml:"MsgType,omitempty"`      // event,text,image
	MsgID        int64  `xml:"MsgId,omitempty"`
	AgentID      int    `xml:"AgentID,omitempty"`
	Event        string `xml:"Event,omitempty"` // same as InfoType

	// 第三方应用开发

	SuiteID    string `xml:"SuiteId,omitempty"`    // 第三方应用ID
	AuthCorpID string `xml:"AuthCorpId,omitempty"` // 授权企业的CorpID
	InfoType   string `xml:"InfoType,omitempty"`
	Timestamp  int64  `xml:"TimeStamp,omitempty"`

	// 公共

	ChangeType string `xml:"ChangeType,omitempty"`
}

func (e CommonPushEvent) IsMessage() bool {
	return e.MsgType != "" && e.MsgType != EventMessageType
}

func (e CommonPushEvent) IsEvent() bool {
	return e.InfoType != "" || e.MsgType == "event"
}

func (e CommonPushEvent) GetMessageType() string {
	return e.MsgType
}

func (e CommonPushEvent) GetEventType() string {
	if e.Event != "" {
		return e.Event
	}
	if e.InfoType != "" {
		return e.InfoType
	}
	if e.MsgType != "" {
		return e.MsgType
	}
	return ""
}

func (e CommonPushEvent) GetTimestamp() int64 {
	if e.Timestamp != 0 {
		return e.Timestamp
	}
	return e.CreateTime
}

type ChangeContactExtAttrItem struct {
	Name  string `xml:"Name"`
	Value string `xml:"Value"`
}

type ChangeContactExtAttr struct {
	Item []ChangeContactExtAttrItem `xml:"Item"`
}
