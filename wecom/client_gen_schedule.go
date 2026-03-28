package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// CreateCalendar 创建日历
// 该接口用于通过应用在企业内创建一个日历。
//
// see https://developer.work.weixin.qq.com/document/path/97719
func (c *Client) CreateCalendar(r *CreateCalendarRequest, opts ...any) (out CreateCalendarResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/calendar/add",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateCalendarRequest is request of Client.CreateCalendar
type CreateCalendarRequest struct {
	// Calendar 日历信息
	Calendar any `json:"calendar" validate:"required"`
	// AgentID 授权方安装的应用agentid。仅旧的第三方多应用套件需要填此参数
	AgentID int `json:"agentid"`
}

// CreateCalendarResponse is response of Client.CreateCalendar
type CreateCalendarResponse struct {
	// CalID 日历ID
	CalID string `json:"cal_id"`
	// FailResult 无效的输入内容
	FailResult any `json:"fail_result"`
}

// GetCalendarDetail 获取日历详情
// 该接口用于获取应用在企业内创建的日历信息。
//
// see https://developer.work.weixin.qq.com/document/path/97717
func (c *Client) GetCalendarDetail(r *GetCalendarDetailRequest, opts ...any) (out GetCalendarDetailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/calendar/get",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCalendarDetailRequest is request of Client.GetCalendarDetail
type GetCalendarDetailRequest struct {
	// CalIDList 日历ID列表，调用创建日历接口后获得。一次最多可获取1000条
	CalIDList []string `json:"cal_id_list" validate:"required"`
}

// GetCalendarDetailResponse is response of Client.GetCalendarDetail
type GetCalendarDetailResponse struct {
	// CalendarList 日历列表
	CalendarList []map[string]any `json:"calendar_list"`
}

// CreateSchedule 创建日程
// 该接口用于在日历中创建一个日程。
//
// see https://developer.work.weixin.qq.com/document/path/97726
func (c *Client) CreateSchedule(r *CreateScheduleRequest, opts ...any) (out CreateScheduleResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/schedule/add",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateScheduleRequest is request of Client.CreateSchedule
type CreateScheduleRequest struct {
	// Schedule 日程信息
	Schedule any `json:"schedule" validate:"required"`
	// AgentID 授权方安装的应用agentid
	AgentID int `json:"agentid"`
}

// CreateScheduleResponse is response of Client.CreateSchedule
type CreateScheduleResponse struct {
	// ScheduleID 日程ID
	ScheduleID string `json:"schedule_id"`
}

// AddScheduleAttendees 新增日程参与者
// 该接口用于在日历中更新指定的日程参与者列表，支持增量式更新。
//
// see https://developer.work.weixin.qq.com/document/path/97721
func (c *Client) AddScheduleAttendees(r *AddScheduleAttendeesRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/schedule/add_attendees",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddScheduleAttendeesRequest is request of Client.AddScheduleAttendees
type AddScheduleAttendeesRequest struct {
	// ScheduleID 日程ID。创建日程时返回的ID
	ScheduleID string `json:"schedule_id" validate:"required"`
	// Attendees 日程参与者列表。累计最多支持1000人
	Attendees []map[string]any `json:"attendees"`
}

// DeleteScheduleAttendees 删除日程参与者
// 该接口用于在日历中更新指定的日程参与者列表，是增量式操作。
//
// see https://developer.work.weixin.qq.com/document/path/97722
func (c *Client) DeleteScheduleAttendees(r *DeleteScheduleAttendeesRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/schedule/del_attendees",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteScheduleAttendeesRequest is request of Client.DeleteScheduleAttendees
type DeleteScheduleAttendeesRequest struct {
	// ScheduleID 日程ID。创建日程时返回的ID
	ScheduleID string `json:"schedule_id" validate:"required"`
	// Attendees 日程参与者列表，最多可添加1000人
	Attendees []map[string]any `json:"attendees"`
}

// GetScheduleListByCalendar 获取日历下的日程列表
// 获取指定的日历下的日程列表。仅可获取应用自己创建的日历下的日程。
//
// see https://developer.work.weixin.qq.com/document/path/97723
func (c *Client) GetScheduleListByCalendar(r *GetScheduleListByCalendarRequest, opts ...any) (out GetScheduleListByCalendarResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/schedule/get_by_calendar",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetScheduleListByCalendarRequest is request of Client.GetScheduleListByCalendar
type GetScheduleListByCalendarRequest struct {
	// CalID 日历ID
	CalID string `json:"cal_id" validate:"required"`
	// Offset 分页，偏移量，默认为0
	Offset int `json:"offset"`
	// Limit 分页，预期请求的数据量，默认为500，取值范围 1 ~ 1000
	Limit int `json:"limit"`
}

// GetScheduleListByCalendarResponse is response of Client.GetScheduleListByCalendar
type GetScheduleListByCalendarResponse struct {
	// ScheduleList 日程列表
	ScheduleList []map[string]any `json:"schedule_list"`
	// ScheduleID 日程ID
	ScheduleID string `json:"schedule_id"`
	// Sequence 日程编号，是一个自增数字
	Sequence int `json:"sequence"`
	// Admins 管理员userid列表
	Admins []any `json:"admins"`
	// Attendees 日程参与者列表。最多支持300人
	Attendees []any `json:"attendees"`
	// Summary 日程标题
	Summary string `json:"summary"`
	// Description 日程描述
	Description string `json:"description"`
	// Reminders 提醒相关信息
	Reminders any `json:"reminders"`
	// Location 日程地址
	Location string `json:"location"`
	// StartTime 日程开始时间，Unix时间戳
	StartTime int `json:"start_time"`
	// EndTime 日程结束时间，Unix时间戳
	EndTime int `json:"end_time"`
	// Status 日程状态。0-正常；1-已取消
	Status int `json:"status"`
	// CalID 日程所属日历ID
	CalID string `json:"cal_id"`
}
