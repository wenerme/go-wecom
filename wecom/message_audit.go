package wecom

import "encoding/xml"

// MessageAuditNotifyPushEvent 企业收到或发送新消息时，企业微信可以以事件的形式推送到企业指定的url。回调间隔为15秒，在15秒内若有消息则触发回调，若无消息则不会触发回调。
type MessageAuditNotifyPushEvent struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   string   `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	AgentID      int      `xml:"AgentID"` // 企业应用的id
	Event        string   `xml:"Event"`
}

func (MessageAuditNotifyPushEvent) EventType() string {
	return "msgaudit_notify"
}

type ChangeExternalContactMessageAuditApproved struct {
	XMLName        xml.Name `xml:"xml"`
	ToUserName     string   `xml:"ToUserName"`
	FromUserName   string   `xml:"FromUserName"`
	CreateTime     int64    `xml:"CreateTime"`
	MsgType        string   `xml:"MsgType"` // 固定为event
	Event          string   `xml:"Event"`
	ChangeType     string   `xml:"ChangeType"`
	UserID         string   `xml:"UserID"`         // 企业服务人员的UserID
	ExternalUserID string   `xml:"ExternalUserID"` // 外部联系人的userid，注意不是企业成员的帐号
	WelcomeCode    string   `xml:"WelcomeCode"`    // 欢迎语code，可用于发送欢迎语
}

func (ChangeExternalContactMessageAuditApproved) EventType() string {
	return "change_external_contact"
}

func (ChangeExternalContactMessageAuditApproved) EventChangeType() string {
	return "msg_audit_approved"
}

func init() {
	RegisterEventModel(
		MessageAuditNotifyPushEvent{},
		ChangeExternalContactMessageAuditApproved{},
	)
}
