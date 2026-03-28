package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// GetChatMsg 获取会话记录数据
// 通过SDK接口拉取企业一段时间内的会话记录，支持分页和代理配置。
//
// see https://developer.work.weixin.qq.com/document/path/91774
func (c *Client) GetChatMsg(r *GetChatMsgRequest, opts ...any) (out GetChatMsgResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/external-contact/getchatmsg",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetChatMsgRequest is request of Client.GetChatMsg
type GetChatMsgRequest struct {
	// Seq 本次请求获取消息记录开始的seq值。首次访问填写0，非首次使用上次企业微信返回的最大seq。Uint64类型
	Seq string `json:"seq" validate:"required"`
	// Limit 一次调用限制的limit值，不能超过1000。uint32类型
	Limit int `json:"limit" validate:"required"`
	// Proxy 使用代理的请求，需要传入代理的链接。如：socks5://10.0.0.1:8081 或者 http://10.0.0.1:8081
	Proxy string `json:"proxy"`
	// Passwd 代理账号密码，需要传入代理的账号密码。如 user_name:passwd_123
	Passwd string `json:"passwd"`
	// Timeout 超时时长，单位 秒
	Timeout int `json:"timeout" validate:"required"`
}

// GetChatMsgResponse is response of Client.GetChatMsg
type GetChatMsgResponse struct {
	// Chatdata 聊天记录数据内容数组
	Chatdata []map[string]any `json:"chatdata"`
}

