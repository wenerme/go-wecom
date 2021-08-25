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

	ToUserName   string `xml:"ToUserName,omitempty"`   // 企业微信CorpID
	FromUserName string `xml:"FromUserName,omitempty"` // sys 系统生成
	CreateTime   int64  `xml:"CreateTime,omitempty"`   // same as Timestamp
	MsgType      int64  `xml:"MsgType,omitempty"`      // event
	Event        string `xml:"Event,omitempty"`        // same as InfoType

	// 第三方应用开发

	SuiteID    string `xml:"SuiteId,omitempty"`    // 第三方应用ID
	AuthCorpID string `xml:"AuthCorpId,omitempty"` // 授权企业的CorpID
	InfoType   string `xml:"InfoType,omitempty"`
	Timestamp  int64  `xml:"TimeStamp,omitempty"`

	// 公共

	ChangeType string `xml:"ChangeType,omitempty"`
}

type ChangeContactExtAttrItem struct {
	Name  string `xml:"Name"`
	Value string `xml:"Value"`
}

type ChangeContactExtAttr struct {
	Item []ChangeContactExtAttrItem `xml:"Item"`
}
