package wecom

import "encoding/xml"

// GetAgentAllowUserInfos is model of GetAgentResponse.AllowUserInfos
type GetAgentAllowUserInfos struct {
	Users []struct {
		UserID string `json:"userid"`
	} `json:"user"`
}

// GetAgentAllowParties is model of GetAgentResponse.AllowParties
type GetAgentAllowParties struct {
	PartyIDs []int `json:"partyid"`
}

// GetAgentAllowTags is model of GetAgentResponse.AllowTags
type GetAgentAllowTags struct {
	TagIDs []int `json:"tagid"`
}

// ListAgentResponseItem is model of ListAgentResponse.AgentList
type ListAgentResponseItem struct {
	AgentID       int    `json:"agentid"`
	Name          string `json:"name"`
	SquareLogoURL string `json:"square_logo_url"`
}

type WorkbenchTemplateItemKeyData struct {
	Items []WorkbenchTemplateItemKeyDataItem `json:"items"`
}

type WorkbenchTemplateItemKeyDataItem struct {
	Key      string `json:"key"`
	Data     string `json:"data"`
	JumpURL  string `json:"jump_url"`
	PagePath string `json:"pagepath"`
}

type WorkbenchTemplateItemImage struct {
	URL      string `json:"url"`
	JumpURL  string `json:"jump_url"`
	PagePath string `json:"pagepath"`
}

type WorkbenchTemplateItemWebView struct {
	URL      string `json:"url"`
	JumpURL  string `json:"jump_url"`
	PagePath string `json:"pagepath"`
}

type WorkbenchTemplateItemList struct {
	Title    string `json:"title"`
	JumpURL  string `json:"jump_url"`
	PagePath string `json:"pagepath"`
}

// SwitchWorkbenchModePushEvent 修改设置工作台自定义开关事件推送
//
// see https://work.weixin.qq.com/api/doc/90001/90143/94620
type SwitchWorkbenchModePushEvent struct {
	XMLName xml.Name `xml:"xml"`
	// AgentID 企业应用的id，整型。可在应用的设置页面查看
	AgentID int `xml:"AgentID"`
	// CreateTime 消息创建时间（整型）
	CreateTime int64 `xml:"CreateTime"`
	// Event 事件类型：template_card_event，点击任务卡片按钮
	Event string `xml:"Event"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName"`
	// MsgType 消息类型，此时固定为：event
	MsgType string `xml:"MsgType"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName"`
	// Mode 1表示开启工作台自定义模式，0表示关闭工作台自定义模式
	Mode int `xml:"Mode"`
}

func (SwitchWorkbenchModePushEvent) EventType() string {
	return "switch_workbench_mode"
}

func init() {
	RegisterEventModel(SwitchWorkbenchModePushEvent{})
}
