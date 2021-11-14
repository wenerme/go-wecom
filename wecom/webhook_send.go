package wecom

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
	"github.com/wenerme/go-req"
)

type WebhookSendRequest struct {
	Key     string
	Content MessageContent
	Context context.Context
	Request req.Request
}

type WebhookSendResponse struct{}

// WebhookSend request webhook to send robot messages
func WebhookSend(r *WebhookSendRequest) (err error) {
	b := &SendPayload{}
	err = b.SetContent(r.Content)
	if err != nil {
		return err
	}
	er := GenericResponse{}
	err = r.Request.With(req.Request{
		URL:    "https://qyapi.weixin.qq.com/cgi-bin/webhook/send",
		Method: http.MethodPost,
		Query: map[string]string{
			"key": r.Key,
		},
		Context: r.Context,
		Body:    b,
		Options: []interface{}{req.JSONEncode, req.JSONDecode},
	}).Fetch(&er)
	if err == nil {
		err = er.AsError()
	}
	return
}

type MessageContent interface {
	MessageType() string
}

const (
	MessageTypeText         = "text"
	MessageTypeMarkdown     = "markdown"
	MessageTypeImage        = "image"
	MessageTypeNews         = "news"
	MessageTypeFile         = "file"
	MessageTypeTemplateCare = "template_card"
)

type SendImageContent struct {
	Base64 string `json:"base64,omitempty"`
	MD5    string `json:"md5,omitempty"`
}

func (SendImageContent) MessageType() string {
	return MessageTypeImage
}

type SendTextContent struct {
	Content             string   `json:"content,omitempty"`
	MentionedList       []string `json:"mentioned_list,omitempty"`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

func (SendTextContent) MessageType() string {
	return MessageTypeText
}

type SendNewsContent struct {
	Articles []SendNewsArticle `json:"articles,omitempty"`
}

type SendNewsArticle struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	URL         string `json:"url,omitempty"`
	PictureURL  string `json:"picurl,omitempty"`
}

func (SendNewsContent) MessageType() string {
	return MessageTypeNews
}

type SendMarkdownContent struct {
	Content string `json:"content,omitempty"`
}

func (SendMarkdownContent) MessageType() string {
	return MessageTypeMarkdown
}

type SendPayload struct {
	MessageType string               `json:"msgtype"`
	Text        *SendTextContent     `json:"text,omitempty"`
	Markdown    *SendMarkdownContent `json:"markdown,omitempty"`
	Image       *SendImageContent    `json:"image,omitempty"`
	News        *SendNewsContent     `json:"news,omitempty"`
}

func (p *SendPayload) SetContent(c MessageContent) (err error) {
	if c == nil {
		err = errors.Errorf("nil message content")
		return
	}
	p.MessageType = c.MessageType()
	switch v := c.(type) {
	case SendTextContent:
		p.Text = &v
	case *SendTextContent:
		p.Text = v
	case SendMarkdownContent:
		p.Markdown = &v
	case *SendMarkdownContent:
		p.Markdown = v
	case SendImageContent:
		p.Image = &v
	case *SendImageContent:
		p.Image = v
	case SendNewsContent:
		p.News = &v
	case *SendNewsContent:
		p.News = v
	default:
		err = errors.Errorf("unsupported message content %T", v)
	}
	return
}
