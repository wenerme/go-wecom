package wecom

import "encoding/xml"

type SuiteTicketEvent struct {
	XMLName     xml.Name `xml:"xml"`
	SuiteID     string   `xml:"SuiteId"`
	InfoType    string   `xml:"InfoType"`
	TimeStamp   int64    `xml:"TimeStamp"`
	SuiteTicket string   `xml:"SuiteTicket"`
}

// ServiceGenericCallbackRequest 服务商回调请求
type ServiceGenericCallbackRequest struct {
	MessageSignature string `json:"msg_signature"`
	Timestamp        string `json:"timestamp"`
	Nonce            string `json:"nonce"`
	EchoString       string `json:"echostr,omitempty"`
}
