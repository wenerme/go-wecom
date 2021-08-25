package wecom

import "encoding/xml"

/*
第三方回调协议
https://work.weixin.qq.com/api/doc/10982
*/

type EncryptPushEvent struct {
	XMLName    xml.Name `xml:"xml"`
	ToUserName string   `xml:"ToUserName"`
	Encrypt    string   `xml:"Encrypt"`
	AgentID    string   `xml:"AgentID"`
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

	ToUserName   string `xml:"ToUserName"`   // 企业微信CorpID
	FromUserName string `xml:"FromUserName"` // sys 系统生成
	CreateTime   int64  `xml:"CreateTime"`   // same as Timestamp
	MsgType      int64  `xml:"MsgType"`      // event
	Event        string `xml:"Event"`        // same as InfoType

	// 第三方应用开发

	SuiteID    string `xml:"SuiteId"`    // 第三方应用ID
	AuthCorpId string `xml:"AuthCorpId"` // 授权企业的CorpID
	InfoType   string `xml:"InfoType"`
	Timestamp  int64  `xml:"TimeStamp"`

	// 公共

	ChangeType string `xml:"ChangeType"`
}

type ChangeContactExtAttrItem struct {
	Name  string `xml:"Name"`
	Value string `xml:"Value"`
}

type ChangeContactExtAttr struct {
	Item []ChangeContactExtAttrItem `xml:"Item"`
}
