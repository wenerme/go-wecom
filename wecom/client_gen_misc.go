package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// SendWebhookMessage 发送消息推送
// 通过webhook向群组发送消息，支持文本、markdown、图片等多种类型
//
// see https://developer.work.weixin.qq.com/document/path/101066
func (c *Client) SendWebhookMessage(r *SendWebhookMessageRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/webhook/send",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SendWebhookMessageRequest is request of Client.SendWebhookMessage
type SendWebhookMessageRequest struct {
	// Key webhook密钥，从消息推送页面获取
	Key string `json:"key" validate:"required"`
	// Msgtype 消息类型，固定为text/markdown/markdown_v2/image/news/file/voice/template_card
	Msgtype string `json:"msgtype" validate:"required"`
	// Text 文本消息内容对象，包含content等字段
	Text any `json:"text"`
	// Markdown markdown消息内容对象，包含content字段
	Markdown any `json:"markdown"`
	// MarkdownV2 markdown_v2消息内容对象，包含content字段
	MarkdownV2 any `json:"markdown_v2"`
	// Image 图片消息内容对象，包含base64和md5字段
	Image any `json:"image"`
	// News 图文消息内容对象，包含articles列表
	News any `json:"news"`
	// File 文件消息内容对象，包含media_id字段
	File any `json:"file"`
	// Voice 语音消息内容对象，包含media_id字段
	Voice any `json:"voice"`
	// TemplateCard 模板卡片消息内容对象
	TemplateCard any `json:"template_card"`
}

