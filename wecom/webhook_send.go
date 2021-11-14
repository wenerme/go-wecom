package wecom

import (
	"context"
	"encoding/json"
	"io/fs"
	"net/http"
	"time"

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
	err = req.Request{
		BaseURL: DefaultAPI,
		URL:     "/cgi-bin/webhook/send",
		Method:  http.MethodPost,
		Query: map[string]string{
			"key": r.Key,
		},
		Context: r.Context,
		Body:    b,
		Options: []interface{}{req.JSONEncode, req.JSONDecode},
	}.With(r.Request).Fetch(&er)
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
	MessageTypeTemplateCard = "template_card"

	MentionAll = "@all"
)

type SendImageContent struct {
	Base64 string `json:"base64,omitempty"`
	MD5    string `json:"md5,omitempty"`
}

func (SendImageContent) MessageType() string {
	return MessageTypeImage
}

type SendTextContent struct {
	Content             string   `json:"content,omitempty"`               // <= 2048byte
	MentionedList       []string `json:"mentioned_list,omitempty"`        // userid or MentionAll
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"` // mobile or MentionAll
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

type SendFileContent struct {
	MediaID string `json:"media_id,omitempty"`
}

func (SendFileContent) MessageType() string {
	return MessageTypeFile
}

type SendTemplateCardContent struct {
	CardType string `json:"card_type"`
	Source   struct {
		IconURL     string `json:"icon_url"`
		Description string `json:"desc"`
	} `json:"source"`
	MainTitle struct {
		Title       string `json:"title"`
		Description string `json:"desc"`
	} `json:"main_title"`
	EmphasisContent struct {
		Title       string `json:"title"`
		Description string `json:"desc"`
	} `json:"emphasis_content"`
	SubTitleText          string `json:"sub_title_text"`
	HorizontalContentList []struct {
		KeyName string `json:"keyname"`
		Value   string `json:"value"`
		Type    int    `json:"type,omitempty"`
		URL     string `json:"url,omitempty"`
		MediaID string `json:"media_id,omitempty"`
	} `json:"horizontal_content_list"`
	JumpList []struct {
		Type     int    `json:"type"`
		URL      string `json:"url,omitempty"`
		Title    string `json:"title"`
		AppID    string `json:"appid,omitempty"`
		PagePath string `json:"pagepath,omitempty"`
	} `json:"jump_list"`
	CardAction struct {
		Type     int    `json:"type"`
		URL      string `json:"url"`
		AppID    string `json:"appid"`
		PagePath string `json:"pagepath"`
	} `json:"card_action"`
}

func (SendTemplateCardContent) MessageType() string {
	return MessageTypeTemplateCard
}

type SendPayload struct {
	MessageType  string                   `json:"msgtype"`
	Text         *SendTextContent         `json:"text,omitempty"`
	Markdown     *SendMarkdownContent     `json:"markdown,omitempty"`
	Image        *SendImageContent        `json:"image,omitempty"`
	News         *SendNewsContent         `json:"news,omitempty"`
	File         *SendFileContent         `json:"file,omitempty"`
	TemplateCard *SendTemplateCardContent `json:"template_card,omitempty"`
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
	case SendFileContent:
		p.File = &v
	case *SendFileContent:
		p.File = v
	case SendTemplateCardContent:
		p.TemplateCard = &v
	case *SendTemplateCardContent:
		p.TemplateCard = v
	default:
		err = errors.Errorf("unsupported message content %T", v)
	}
	return
}

type WebhookUploadMediaResponse struct {
	Type      string      `json:"type,omitempty"` // file
	MediaID   string      `json:"media_id,omitempty"`
	CreatedAt json.Number `json:"created_at,omitempty"`
}

func (v WebhookUploadMediaResponse) CreatedAtTime() time.Time {
	i, _ := v.CreatedAt.Int64()
	return time.Unix(i, 0)
}

type WebhookUploadMediaRequest struct {
	Key     string
	File    fs.File
	Context context.Context
	Request req.Request
}

// WebhookUploadMedia upload media to get MediaID
func WebhookUploadMedia(r *WebhookUploadMediaRequest) (out WebhookUploadMediaResponse, err error) {
	er := GenericResponse{}
	err = req.Request{
		BaseURL: DefaultAPI,
		URL:     "/cgi-bin/webhook/upload_media",
		Method:  http.MethodPost,
		Query: map[string]string{
			"key":  r.Key,
			"type": "file",
		},
		Context: r.Context,
		Body:    r.File,
		Options: []interface{}{req.JSONDecode},
	}.With(r.Request).Fetch(&er, &out)
	if err == nil {
		err = er.AsError()
	}
	return
}
