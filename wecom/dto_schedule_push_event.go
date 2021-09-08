package wecom

import "encoding/xml"

// ModifyCalendarPushEvent 修改日历事件
// 当用户修改了API创建的日历时，触发该事件。
//
// see https://open.work.weixin.qq.com/api/doc/90000/90135/93651
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

// EventType impl EventModel
func (ModifyCalendarPushEvent) EventType() string {
	return "modify_calendar" //nolint:goconst
}

// MessageType impl MessageModel
func (ModifyCalendarPushEvent) MessageType() string {
	return "event" //nolint:goconst
}

// DeleteCalendarPushEvent 删除日历事件
// 当组织者删除了API创建的日历时，触发该事件。
//
// see https://open.work.weixin.qq.com/api/doc/90000/90135/93651
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

// EventType impl EventModel
func (DeleteCalendarPushEvent) EventType() string {
	return "delete_calendar" //nolint:goconst
}

// MessageType impl MessageModel
func (DeleteCalendarPushEvent) MessageType() string {
	return "event" //nolint:goconst
}

// AddSchedulePushEvent 添加日程事件
// 当用户在API创建的日历上添加了日程后，触发该事件。
//
// see https://open.work.weixin.qq.com/api/doc/90000/90135/93651
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

// EventType impl EventModel
func (AddSchedulePushEvent) EventType() string {
	return "add_schedule" //nolint:goconst
}

// MessageType impl MessageModel
func (AddSchedulePushEvent) MessageType() string {
	return "event" //nolint:goconst
}

// ModifySchedulePushEvent 修改日程事件
// 当用户在API创建的日历上修改了日程后，触发该事件。
//
// see https://open.work.weixin.qq.com/api/doc/90000/90135/93651
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

// EventType impl EventModel
func (ModifySchedulePushEvent) EventType() string {
	return "modify_schedule" //nolint:goconst
}

// MessageType impl MessageModel
func (ModifySchedulePushEvent) MessageType() string {
	return "event" //nolint:goconst
}

// DeleteSchedulePushEvent 删除日程事件
// 当用户在API创建的日历上删除了日程后，触发该事件。
//
// see https://open.work.weixin.qq.com/api/doc/90000/90135/93651
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

// EventType impl EventModel
func (DeleteSchedulePushEvent) EventType() string {
	return "delete_schedule" //nolint:goconst
}

// MessageType impl MessageModel
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
