package wecom

import (
	"github.com/wenerme/go-req"
)

// AddCalendar 创建日历
// 该接口用于通过应用在企业内创建一个日历。
//
// see https://open.work.weixin.qq.com/api/doc/90000/90135/93647
func (c *Client) AddCalendar(r *AddCalendarRequest, opts ...interface{}) (out AddCalendarResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/calendar/add",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateCalendar 更新日历
// 该接口用于修改指定日历的信息。
//
// see https://open.work.weixin.qq.com/api/doc/90000/90135/93647
func (c *Client) UpdateCalendar(r *UpdateCalendarRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/calendar/update",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCalendar 获取日历详情
// 该接口用于获取应用在企业内创建的日历信息。
//
// see https://open.work.weixin.qq.com/api/doc/90000/90135/93647
func (c *Client) GetCalendar(r *GetCalendarRequest, opts ...interface{}) (out GetCalendarResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/calendar/get",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteCalendar 删除日历
// 该接口用于删除指定日历。
//
// see https://open.work.weixin.qq.com/api/doc/90000/90135/93647
func (c *Client) DeleteCalendar(r *DeleteCalendarRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/calendar/del",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddSchedule 创建日程
// 该接口用于在日历中创建一个日程。
//
// see https://open.work.weixin.qq.com/api/doc/90000/90135/93648
func (c *Client) AddSchedule(r *AddScheduleRequest, opts ...interface{}) (out AddScheduleResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/schedule/add",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateSchedule 更新日程
// 该接口用于在日历中更新指定的日程。
//
// see https://open.work.weixin.qq.com/api/doc/90000/90135/93648
func (c *Client) UpdateSchedule(r *UpdateScheduleRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/schedule/update",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetSchedule 获取日程详情
// 该接口用于获取指定的日程详情。
//
// see https://open.work.weixin.qq.com/api/doc/90000/90135/93648
func (c *Client) GetSchedule(r *GetScheduleRequest, opts ...interface{}) (out GetScheduleResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/schedule/get",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteSchedule 取消日程
// 该接口用于取消指定的日程。
//
// see https://open.work.weixin.qq.com/api/doc/90000/90135/93648
func (c *Client) DeleteSchedule(r *DeleteScheduleRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/schedule/del",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ScheduleGetByCalendar 获取日历下的日程列表
// 该接口用于获取指定的日历下的日程列表。仅可获取应用自己创建的日历下的日程。
//
// see https://open.work.weixin.qq.com/api/doc/90000/90135/93648
func (c *Client) ScheduleGetByCalendar(r *ScheduleGetByCalendarRequest, opts ...interface{}) (out ScheduleGetByCalendarResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/schedule/get_by_calendar",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

type AddCalendarRequest struct {
	// Calendar 日历信息
	Calendar AddCalendar `json:"calendar"  validate:"required"`
	// AgentID 授权方安装的应用agentid。仅旧的第三方多应用套件需要填此参数
	AgentID int `json:"agentid"  `
}

type AddCalendarResponse struct {
	// CalenderID 日历ID
	CalenderID string `json:"cal_id"  `
}

type AddCalendarRequestShares struct {
	// UserID 日历共享成员的id
	UserID string `json:"userid"  validate:"required"`
	// Readonly 共享成员对日历是否只读权限（即不可编辑日历，不可在日历上添加日程，仅可以退出日历）。0-否；1-是。默认为1，即只读
	Readonly int `json:"readonly"  `
}

type UpdateCalendarRequest struct {
	// Calendar 日历信息
	Calendar UpdateCalendar `json:"calendar"  validate:"required"`
}

type UpdateCalendarRequestShares struct {
	// UserID 日历共享成员的id
	UserID string `json:"userid"  validate:"required"`
	// Readonly 共享成员对日历是否只读权限（即不可编辑日历，不可在日历上添加日程，仅可以退出日历）。0-否；1-是。默认为1，即只读
	Readonly int `json:"readonly"  `
}

type GetCalendarRequest struct {
	// CalenderIDList 日历ID列表，调用创建日历接口后获得。一次最多可获取1000条
	CalenderIDList []string `json:"cal_id_list"  validate:"required"`
}

type GetCalendarResponse struct {
	// CalendarList 日历列表
	CalendarList []GetCalender `json:"calendar_list"  `
}

type GetCalendarResponseShares struct {
	// UserID 日历共享成员的id
	UserID string `json:"userid"  `
	// Readonly 共享成员对日历是否只读权限。0-否；1-是；
	Readonly int `json:"readonly"  `
}

type DeleteCalendarRequest struct {
	// CalenderID 日历ID
	CalenderID string `json:"cal_id"  validate:"required"`
}

type AddScheduleRequest struct {
	// Schedule obj
	Schedule AddScheduleRequestSchedule `json:"schedule"  validate:"required"`
	// AgentID uint32
	AgentID int `json:"agentid"  `
}

type AddScheduleResponse struct {
	// ScheduleID string
	ScheduleID string `json:"schedule_id"  `
}

type AddScheduleRequestSchedule struct {
	// Attendees obj[]
	Attendees []AddScheduleRequestScheduleAttendees `json:"attendees"  `
	// Summary string
	Summary string `json:"summary"  `
	// Description string
	Description string `json:"description"  `
	// Reminders obj
	Reminders AddScheduleRequestScheduleReminders `json:"reminders"  `
	// Location string
	Location string `json:"location"  `
	// Organizer string
	Organizer string `json:"organizer"  validate:"required"`
	// StartTime uint32
	StartTime int `json:"start_time"  `
	// EndTime uint32
	EndTime int `json:"end_time"  `
	// CalenderID string
	CalenderID string `json:"cal_id"  `
}

type AddScheduleRequestScheduleAttendees struct {
	// UserID string
	UserID string `json:"userid"  validate:"required"`
}

type AddScheduleRequestScheduleReminders struct {
	// IsRemind int32
	IsRemind int `json:"is_remind"  `
	// IsRepeat int32
	IsRepeat int `json:"is_repeat"  `
	// RemindBeforeEventSecs uint32
	RemindBeforeEventSecs int `json:"remind_before_event_secs"  `
	// RepeatType uint32
	RepeatType int `json:"repeat_type"  `
	// RepeatUntil uint32
	RepeatUntil int `json:"repeat_until"  `
	// IsCustomRepeat uint32
	IsCustomRepeat int `json:"is_custom_repeat"  `
	// RepeatInterval uint32
	RepeatInterval int `json:"repeat_interval"  `
	// RepeatDayOfWeek uint32[]
	RepeatDayOfWeek []int `json:"repeat_day_of_week"  `
	// RepeatDayOfMonth uint32[]
	RepeatDayOfMonth []int `json:"repeat_day_of_month"  `
	// Timezone uint32
	Timezone int `json:"timezone"  `
}

type UpdateScheduleRequest struct {
	// Schedule obj
	Schedule UpdateScheduleRequestSchedule `json:"schedule"  validate:"required"`
}

type UpdateScheduleRequestSchedule struct {
	// ScheduleID string
	ScheduleID string `json:"schedule_id"  validate:"required"`
	// Attendees obj[]
	Attendees []UpdateScheduleRequestScheduleAttendees `json:"attendees"  `
	// Summary string
	Summary string `json:"summary"  `
	// Description string
	Description string `json:"description"  `
	// Reminders obj
	Reminders UpdateScheduleRequestScheduleReminders `json:"reminders"  `
	// Location string
	Location string `json:"location"  `
	// Organizer string
	Organizer string `json:"organizer"  `
	// StartTime uint32
	StartTime int `json:"start_time"  `
	// EndTime uint32
	EndTime int `json:"end_time"  `
}

type UpdateScheduleRequestScheduleAttendees struct {
	// UserID string
	UserID string `json:"userid"  validate:"required"`
}

type UpdateScheduleRequestScheduleReminders struct {
	// IsRemind int32
	IsRemind int `json:"is_remind"  `
	// IsRepeat int32
	IsRepeat int `json:"is_repeat"  `
	// RemindBeforeEventSecs uint32
	RemindBeforeEventSecs int `json:"remind_before_event_secs"  `
	// RepeatType uint32
	RepeatType int `json:"repeat_type"  `
	// RepeatUntil uint32
	RepeatUntil int `json:"repeat_until"  `
	// IsCustomRepeat uint32
	IsCustomRepeat int `json:"is_custom_repeat"  `
	// RepeatInterval uint32
	RepeatInterval int `json:"repeat_interval"  `
	// RepeatDayOfWeek uint32[]
	RepeatDayOfWeek []int `json:"repeat_day_of_week"  `
	// RepeatDayOfMonth uint32[]
	RepeatDayOfMonth []int `json:"repeat_day_of_month"  `
	// Timezone uint32
	Timezone int `json:"timezone"  `
}

type GetScheduleRequest struct {
	// ScheduleIDList 日程ID列表。一次最多拉取1000条
	ScheduleIDList []string `json:"schedule_id_list"  validate:"required"`
}

type GetScheduleResponse struct {
	// ScheduleList obj[]
	ScheduleList []GetScheduleResponseScheduleList `json:"schedule_list"  `
}

type GetScheduleResponseScheduleList struct {
	// ScheduleID string
	ScheduleID string `json:"schedule_id"  `
	// Attendees obj[]
	Attendees []GetScheduleResponseScheduleListAttendees `json:"attendees"  `
	// Summary string
	Summary string `json:"summary"  `
	// Description string
	Description string `json:"description"  `
	// Reminders obj
	Reminders GetScheduleResponseScheduleListReminders `json:"reminders"  `
	// Location string
	Location string `json:"location"  `
	// Organizer string
	Organizer string `json:"organizer"  `
	// Status uint32
	Status int `json:"status"  `
	// StartTime uint32
	StartTime int `json:"start_time"  `
	// EndTime uint32
	EndTime int `json:"end_time"  `
	// CalenderID string
	CalenderID string `json:"cal_id"  `
}

type GetScheduleResponseScheduleListAttendees struct {
	// UserID string
	UserID string `json:"userid"  `
	// ResponseStatus uint32
	ResponseStatus int `json:"response_status"  `
}

type GetScheduleResponseScheduleListReminders struct {
	// IsRemind int32
	IsRemind int `json:"is_remind"  `
	// IsRepeat int32
	IsRepeat int `json:"is_repeat"  `
	// RemindBeforeEventSecs uint32
	RemindBeforeEventSecs int `json:"remind_before_event_secs"  `
	// RemindTimeDiffs int32[]
	RemindTimeDiffs []int `json:"remind_time_diffs"  `
	// RepeatType uint32
	RepeatType int `json:"repeat_type"  `
	// RepeatUntil uint32
	RepeatUntil int `json:"repeat_until"  `
	// IsCustomRepeat uint32
	IsCustomRepeat int `json:"is_custom_repeat"  `
	// RepeatInterval uint32
	RepeatInterval int `json:"repeat_interval"  `
	// RepeatDayOfWeek uint32[]
	RepeatDayOfWeek []int `json:"repeat_day_of_week"  `
	// RepeatDayOfMonth uint32[]
	RepeatDayOfMonth []int `json:"repeat_day_of_month"  `
	// Timezone uint32
	Timezone int `json:"timezone"  `
	// ExcludeTimeList obj[]
	ExcludeTimeList []GetScheduleResponseScheduleListRemindersExcludeTimeList `json:"exclude_time_list"  `
}

type GetScheduleResponseScheduleListRemindersExcludeTimeList struct {
	// StartTime uint32
	StartTime int `json:"start_time"  `
}

type DeleteScheduleRequest struct {
	// ScheduleID 日程ID
	ScheduleID string `json:"schedule_id"  validate:"required"`
}

type ScheduleGetByCalendarRequest struct {
	// CalenderID 日历ID
	CalenderID string `json:"cal_id"  validate:"required"`
	// Offset 分页，偏移量, 默认为0
	Offset int `json:"offset"  `
	// Limit 分页，预期请求的数据量，默认为500，取值范围 1 ~ 1000
	Limit int `json:"limit"  `
}

type ScheduleGetByCalendarResponse struct {
	// ScheduleList obj[]
	ScheduleList []ScheduleGetByCalendarResponseScheduleList `json:"schedule_list"  `
}

type ScheduleGetByCalendarResponseScheduleList struct {
	// ScheduleID string
	ScheduleID string `json:"schedule_id"  `
	// Attendees obj[]
	Attendees []ScheduleGetByCalendarResponseScheduleListAttendees `json:"attendees"  `
	// Summary string
	Summary string `json:"summary"  `
	// Description string
	Description string `json:"description"  `
	// Reminders obj
	Reminders ScheduleGetByCalendarResponseScheduleListReminders `json:"reminders"  `
	// Location string
	Location string `json:"location"  `
	// Organizer string
	Organizer string `json:"organizer"  `
	// Status uint32
	Status int `json:"status"  `
	// StartTime uint32
	StartTime int `json:"start_time"  `
	// EndTime uint32
	EndTime int `json:"end_time"  `
	// Sequence uint64
	Sequence int `json:"sequence"  `
	// CalenderID string
	CalenderID string `json:"cal_id"  `
}

type ScheduleGetByCalendarResponseScheduleListAttendees struct {
	// UserID string
	UserID string `json:"userid"  `
	// ResponseStatus uint32
	ResponseStatus int `json:"response_status"  `
}

type ScheduleGetByCalendarResponseScheduleListReminders struct {
	// IsRemind int32
	IsRemind int `json:"is_remind"  `
	// IsRepeat int32
	IsRepeat int `json:"is_repeat"  `
	// RemindBeforeEventSecs uint32
	RemindBeforeEventSecs int `json:"remind_before_event_secs"  `
	// RepeatType uint32
	RepeatType int `json:"repeat_type"  `
	// RepeatUntil uint32
	RepeatUntil int `json:"repeat_until"  `
	// IsCustomRepeat uint32
	IsCustomRepeat int `json:"is_custom_repeat"  `
	// RepeatInterval uint32
	RepeatInterval int `json:"repeat_interval"  `
	// RepeatDayOfWeek uint32[]
	RepeatDayOfWeek []int `json:"repeat_day_of_week"  `
	// RepeatDayOfMonth uint32[]
	RepeatDayOfMonth []int `json:"repeat_day_of_month"  `
	// Timezone uint32
	Timezone int `json:"timezone"  `
}
