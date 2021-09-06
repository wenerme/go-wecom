package wecom

import "encoding/xml"

// ModifyCalendarPushEvent 当用户修改了API创建的日历时，触发该事件。
type ModifyCalendarPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// CalID 日历ID
	CalID string `xml:"CalId" json:"CalId"`
	// CreateTime 消息创建时间，unix时间戳
	CreateTime int64 `xml:"CreateTime" json:"CreateTime"`
	// Event 事件类型，此时固定为：modify_calendar
	Event string `xml:"Event" json:"Event"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName" json:"FromUserName"`
	// MsgType 消息类型，此时固定为：event
	MsgType string `xml:"MsgType" json:"MsgType"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName" json:"ToUserName"`
}

func (ModifyCalendarPushEvent) EventType() string {
	return "modify_calendar" //nolint:goconst
}

func (ModifyCalendarPushEvent) MessageType() string {
	return "event" //nolint:goconst
}

// DeleteCalendarPushEvent 当组织者删除了API创建的日历时，触发该事件。
type DeleteCalendarPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// CalID 日历ID
	CalID string `xml:"CalId" json:"CalId"`
	// CreateTime 消息创建时间，unix时间戳
	CreateTime int64 `xml:"CreateTime" json:"CreateTime"`
	// Event 事件类型，此时固定为：delete_calendar
	Event string `xml:"Event" json:"Event"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName" json:"FromUserName"`
	// MsgType 消息类型，此时固定为：event
	MsgType string `xml:"MsgType" json:"MsgType"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName" json:"ToUserName"`
}

func (DeleteCalendarPushEvent) EventType() string {
	return "delete_calendar" //nolint:goconst
}

func (DeleteCalendarPushEvent) MessageType() string {
	return "event" //nolint:goconst
}

// AddSchedulePushEvent 当用户在API创建的日历上添加了日程后，触发该事件。
type AddSchedulePushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// CalID 日历ID
	CalID string `xml:"CalId" json:"CalId"`
	// CreateTime 消息创建时间，unix时间戳
	CreateTime int64 `xml:"CreateTime" json:"CreateTime"`
	// Event 事件类型，此时固定为：add_schedule
	Event string `xml:"Event" json:"Event"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName" json:"FromUserName"`
	// MsgType 消息类型，此时固定为：event
	MsgType string `xml:"MsgType" json:"MsgType"`
	// ScheduleID 日程ID
	ScheduleID string `xml:"ScheduleId" json:"ScheduleId"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName" json:"ToUserName"`
}

func (AddSchedulePushEvent) EventType() string {
	return "add_schedule" //nolint:goconst
}

func (AddSchedulePushEvent) MessageType() string {
	return "event" //nolint:goconst
}

// ModifySchedulePushEvent 当用户在API创建的日历上修改了日程后，触发该事件。
type ModifySchedulePushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// CalID 日历ID
	CalID string `xml:"CalId" json:"CalId"`
	// CreateTime 消息创建时间，unix时间戳
	CreateTime int64 `xml:"CreateTime" json:"CreateTime"`
	// Event 事件类型，此时固定为：modify_schedule
	Event string `xml:"Event" json:"Event"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName" json:"FromUserName"`
	// MsgType 消息类型，此时固定为：event
	MsgType string `xml:"MsgType" json:"MsgType"`
	// ScheduleID 日程ID
	ScheduleID string `xml:"ScheduleId" json:"ScheduleId"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName" json:"ToUserName"`
}

func (ModifySchedulePushEvent) EventType() string {
	return "modify_schedule" //nolint:goconst
}

func (ModifySchedulePushEvent) MessageType() string {
	return "event" //nolint:goconst
}

// DeleteSchedulePushEvent 当用户在API创建的日历上删除了日程后，触发该事件。
type DeleteSchedulePushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// CalID 日历ID
	CalID string `xml:"CalId" json:"CalId"`
	// CreateTime 消息创建时间，unix时间戳
	CreateTime int64 `xml:"CreateTime" json:"CreateTime"`
	// Event 事件类型，此时固定为：delete_schedule
	Event string `xml:"Event" json:"Event"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName" json:"FromUserName"`
	// MsgType 消息类型，此时固定为：event
	MsgType string `xml:"MsgType" json:"MsgType"`
	// ScheduleID 日程ID
	ScheduleID string `xml:"ScheduleId" json:"ScheduleId"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName" json:"ToUserName"`
}

func (DeleteSchedulePushEvent) EventType() string {
	return "delete_schedule" //nolint:goconst
}

func (DeleteSchedulePushEvent) MessageType() string {
	return "event" //nolint:goconst
}

func init() {
	RegisterEventModel(
		ModifyCalendarPushEvent{},
		DeleteCalendarPushEvent{},
		AddSchedulePushEvent{},
		ModifySchedulePushEvent{},
		DeleteSchedulePushEvent{},
	)
}
