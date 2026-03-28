package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// AddMeetingRoom 添加会议室
// 企业可通过此接口添加一个会议室。
//
// see https://developer.work.weixin.qq.com/document/path/93619
func (c *Client) AddMeetingRoom(r *AddMeetingRoomRequest, opts ...any) (out AddMeetingRoomResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/meetingroom/add",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddMeetingRoomRequest is request of Client.AddMeetingRoom
type AddMeetingRoomRequest struct {
	// Name 会议室名称，最多30个字符
	Name string `json:"name" validate:"required"`
	// Capacity 会议室所能容纳的人数
	Capacity int `json:"capacity" validate:"required"`
	// City 会议室所在城市
	City string `json:"city"`
	// Building 会议室所在楼宇
	Building string `json:"building"`
	// Floor 会议室所在楼层
	Floor string `json:"floor"`
	// Equipment 会议室支持的设备列表
	Equipment []int `json:"equipment"`
	// Coordinate 会议室所在建筑经纬度对象
	Coordinate any `json:"coordinate"`
	// Range 会议室使用范围，包含user_list和department_list
	Range any `json:"range"`
}

// AddMeetingRoomResponse is response of Client.AddMeetingRoom
type AddMeetingRoomResponse struct {
	// MeetingroomID 会议室的id
	MeetingroomID int `json:"meetingroom_id"`
}

// GetMeetingRoomBookingInfo 查询会议室的预定信息
// 查询相关会议室在指定时间段的预定情况，如是否已被预定，预定者的userid等信息
//
// see https://developer.work.weixin.qq.com/document/path/93620
func (c *Client) GetMeetingRoomBookingInfo(r *GetMeetingRoomBookingInfoRequest, opts ...any) (out GetMeetingRoomBookingInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/meetingroom/get_booking_info",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMeetingRoomBookingInfoRequest is request of Client.GetMeetingRoomBookingInfo
type GetMeetingRoomBookingInfoRequest struct {
	// MeetingroomID 会议室id
	MeetingroomID int `json:"meetingroom_id"`
	// StartTime 查询预定的起始时间，默认为当前时间
	StartTime int `json:"start_time"`
	// EndTime 查询预定的结束时间，默认为明日0时
	EndTime int `json:"end_time"`
	// City 会议室所在城市
	City string `json:"city"`
	// Building 会议室所在楼宇
	Building string `json:"building"`
	// Floor 会议室所在楼层
	Floor string `json:"floor"`
}

// GetMeetingRoomBookingInfoResponse is response of Client.GetMeetingRoomBookingInfo
type GetMeetingRoomBookingInfoResponse struct {
	// BookingList 会议室预订信息列表
	BookingList []map[string]any `json:"booking_list"`
}
