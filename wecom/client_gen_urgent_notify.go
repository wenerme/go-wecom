package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// CallPstnCc 发起语音电话
// 通过此接口发起语音电话，提醒员工查看应用推送的重要消息。
//
// see https://developer.work.weixin.qq.com/document/path/91627
func (c *Client) CallPstnCc(r *CallPstnCcRequest, opts ...any) (out CallPstnCcResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/pstncc/call",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CallPstnCcRequest is request of Client.CallPstnCc
type CallPstnCcRequest struct {
	// CalleeUserID 需要呼叫的成员UserID列表
	CalleeUserID []string `json:"callee_userid" validate:"required"`
}

// CallPstnCcResponse is response of Client.CallPstnCc
type CallPstnCcResponse struct {
	// States 自动语音来电呼叫状态列表
	States []map[string]any `json:"states"`
}

// GetPstnccState 获取接听状态
// 通过此接口，了解员工是否已接听语音电话。
//
// see https://developer.work.weixin.qq.com/document/path/91628
func (c *Client) GetPstnccState(r *GetPstnccStateRequest, opts ...any) (out GetPstnccStateResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/pstncc/getstates",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetPstnccStateRequest is request of Client.GetPstnccState
type GetPstnccStateRequest struct {
	// CalleeUserID 用户id
	CalleeUserID string `json:"callee_userid" validate:"required"`
	// Callid 发起自动语音来电callid
	Callid string `json:"callid" validate:"required"`
}

// GetPstnccStateResponse is response of Client.GetPstnccState
type GetPstnccStateResponse struct {
	// Istalked 0.表示未接听，1.表示接听
	Istalked int `json:"istalked"`
	// Calltime 呼叫发起时间戳
	Calltime int `json:"calltime"`
	// Talktime 通话时长单位（s）
	Talktime int `json:"talktime"`
	// Reason 呼叫结果状态：0正常结束，详见文档说明
	Reason int `json:"reason"`
}
