package wecom

import "encoding/xml"

// MessageTextPushEvent 消息示例：
type MessageTextPushEvent struct {
	XMLName xml.Name `xml:"xml"`
	// AgentID 企业应用的id，整型。可在应用的设置页面查看
	AgentID string `xml:"AgentID"`
	// Content 文本消息内容
	Content string `xml:"Content"`
	// CreateTime 消息创建时间（整型）
	CreateTime int64 `xml:"CreateTime"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName"`
	// MsgID 消息id，64位整型
	MsgID string `xml:"MsgId"`
	// MsgType 消息类型，此时固定为：text
	MsgType string `xml:"MsgType"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName"`
}

func (e MessageTextPushEvent) EventType() string {
	return e.MessageType()
}

func (MessageTextPushEvent) MessageType() string {
	return "text" //nolint:goconst
}

// MessageImagePushEvent 消息示例：
type MessageImagePushEvent struct {
	XMLName xml.Name `xml:"xml"`
	// AgentID 企业应用的id，整型。可在应用的设置页面查看
	AgentID string `xml:"AgentID"`
	// CreateTime 消息创建时间（整型）
	CreateTime int64 `xml:"CreateTime"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName"`
	// MediaID 图片媒体文件id，可以调用获取媒体文件接口拉取，仅三天内有效
	MediaID string `xml:"MediaId"`
	// MsgID 消息id，64位整型
	MsgID string `xml:"MsgId"`
	// MsgType 消息类型，此时固定为：image
	MsgType string `xml:"MsgType"`
	// PicURL 图片链接
	PicURL string `xml:"PicUrl"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName"`
}

func (e MessageImagePushEvent) EventType() string {
	return e.MessageType()
}

func (MessageImagePushEvent) MessageType() string {
	return "image" //nolint:goconst
}

// MessageVoicePushEvent 消息示例：
type MessageVoicePushEvent struct {
	XMLName xml.Name `xml:"xml"`
	// AgentID 企业应用的id，整型。可在应用的设置页面查看
	AgentID string `xml:"AgentID"`
	// CreateTime 消息创建时间（整型）
	CreateTime int64 `xml:"CreateTime"`
	// Format 语音格式，如amr，speex等
	Format string `xml:"Format"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName"`
	// MediaID 语音媒体文件id，可以调用获取媒体文件接口拉取数据，仅三天内有效
	MediaID string `xml:"MediaId"`
	// MsgID 消息id，64位整型
	MsgID string `xml:"MsgId"`
	// MsgType 消息类型，此时固定为：voice
	MsgType string `xml:"MsgType"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName"`
}

func (e MessageVoicePushEvent) EventType() string {
	return e.MessageType()
}

func (MessageVoicePushEvent) MessageType() string {
	return "voice" //nolint:goconst
}

// MessageVideoPushEvent 消息示例：
type MessageVideoPushEvent struct {
	XMLName xml.Name `xml:"xml"`
	// AgentID 企业应用的id，整型。可在应用的设置页面查看
	AgentID string `xml:"AgentID"`
	// CreateTime 消息创建时间（整型）
	CreateTime int64 `xml:"CreateTime"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName"`
	// MediaID 视频媒体文件id，可以调用获取媒体文件接口拉取数据，仅三天内有效
	MediaID string `xml:"MediaId"`
	// MsgID 消息id，64位整型
	MsgID string `xml:"MsgId"`
	// MsgType 消息类型，此时固定为：video
	MsgType string `xml:"MsgType"`
	// ThumbMediaID 视频消息缩略图的媒体id，可以调用获取媒体文件接口拉取数据，仅三天内有效
	ThumbMediaID string `xml:"ThumbMediaId"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName"`
}

func (e MessageVideoPushEvent) EventType() string {
	return e.MessageType()
}

func (MessageVideoPushEvent) MessageType() string {
	return "video" //nolint:goconst
}

// MessageLocationPushEvent 消息示例：
type MessageLocationPushEvent struct {
	XMLName xml.Name `xml:"xml"`
	// AgentID 企业应用的id，整型。可在应用的设置页面查看
	AgentID string `xml:"AgentID"`
	// AppType app类型，在企业微信固定返回wxwork，在微信不返回该字段
	AppType string `xml:"AppType"`
	// CreateTime 消息创建时间（整型）
	CreateTime int64 `xml:"CreateTime"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName"`
	// Label 地理位置信息
	Label string `xml:"Label"`
	// LocationX 地理位置纬度
	LocationX string `xml:"Location_X"`
	// LocationY 地理位置经度
	LocationY string `xml:"Location_Y"`
	// MsgID 消息id，64位整型
	MsgID string `xml:"MsgId"`
	// MsgType 消息类型，此时固定为： location
	MsgType string `xml:"MsgType"`
	// Scale 地图缩放大小
	Scale string `xml:"Scale"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName"`
}

func (e MessageLocationPushEvent) EventType() string {
	return e.MessageType()
}

func (MessageLocationPushEvent) MessageType() string {
	return "location" //nolint:goconst
}

// MessageLinkPushEvent 消息示例：
type MessageLinkPushEvent struct {
	XMLName xml.Name `xml:"xml"`
	// AgentID 企业应用的id，整型。可在应用的设置页面查看
	AgentID string `xml:"AgentID"`
	// CreateTime 消息创建时间（整型）
	CreateTime int64 `xml:"CreateTime"`
	// Description 描述
	Description string `xml:"Description"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName"`
	// MsgID 消息id，64位整型
	MsgID string `xml:"MsgId"`
	// MsgType 消息类型，此时固定为：link
	MsgType string `xml:"MsgType"`
	// PicURL 封面缩略图的url
	PicURL string `xml:"PicUrl"`
	// Title 标题
	Title string `xml:"Title"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName"`
	// URL 链接跳转的url
	URL string `xml:"Url"`
}

func (e MessageLinkPushEvent) EventType() string {
	return e.MessageType()
}

func (MessageLinkPushEvent) MessageType() string {
	return "link" //nolint:goconst
}

func init() {
	RegisterEventModel(
		MessageTextPushEvent{},
		MessageImagePushEvent{},
		MessageVoicePushEvent{},
		MessageVideoPushEvent{},
		MessageLocationPushEvent{},
		MessageLinkPushEvent{},
	)
}
