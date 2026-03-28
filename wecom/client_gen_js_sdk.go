package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// GetLaunchCode 获取唤起企业微信code
// 获取唤起个人聊天窗口的code
//
// see https://developer.work.weixin.qq.com/document/path/94345
func (c *Client) GetLaunchCode(r *GetLaunchCodeRequest, opts ...any) (out GetLaunchCodeResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/get_launch_code",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetLaunchCodeRequest is request of Client.GetLaunchCode
type GetLaunchCodeRequest struct {
	// OperatorUserID 当前操作者的userid
	OperatorUserID string `json:"operator_userid" validate:"required"`
	// SingleChat 参数数据结构体
	SingleChat any `json:"single_chat" validate:"required"`
}

// GetLaunchCodeResponse is response of Client.GetLaunchCode
type GetLaunchCodeResponse struct {
	// LaunchCode 唤起页面的code，5分钟内有效，只能消费一次
	LaunchCode string `json:"launch_code"`
}

// GetInvoiceTicket 获取电子发票ticket
// 获取用于调起电子发票的临时票据
//
// see https://developer.work.weixin.qq.com/document/path/100526
func (c *Client) GetInvoiceTicket(r *GetInvoiceTicketRequest, opts ...any) (out GetInvoiceTicketResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/ticket/get",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetInvoiceTicketRequest is request of Client.GetInvoiceTicket
type GetInvoiceTicketRequest struct {
	// Type 票据类型，固定为wx_card
	Type string `json:"type" validate:"required"`
}

// GetInvoiceTicketResponse is response of Client.GetInvoiceTicket
type GetInvoiceTicketResponse struct {
	// Ticket 发票签名临时票据
	Ticket string `json:"ticket"`
	// ExpiresIn 有效期，以秒为单位
	ExpiresIn int `json:"expires_in"`
}

