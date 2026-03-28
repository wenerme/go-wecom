package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// GetTempMedia 获取临时素材
// 下载临时素材文件
//
// see https://developer.work.weixin.qq.com/document/path/90077
func (c *Client) GetTempMedia(r *GetTempMediaRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/media/get",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetTempMediaRequest is request of Client.GetTempMedia
type GetTempMediaRequest struct {
	// MediaID 媒体文件id
	MediaID string `json:"media_id" validate:"required"`
}

// GetMediaVoice 获取高清语音素材
// 获取从JSSDK的uploadVoice接口上传的临时语音素材，格式为speex，16K采样率。该音频比上文的临时素材获取接口（格式为amr，8K采样率）更加清晰，适合用作语音识别等对音质要求较高的业务。
//
// see https://developer.work.weixin.qq.com/document/path/90078
func (c *Client) GetMediaVoice(r *GetMediaVoiceRequest, opts ...any) (out GetMediaVoiceResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/media/get/jssdk",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMediaVoiceRequest is request of Client.GetMediaVoice
type GetMediaVoiceRequest struct {
	// MediaID 通过JSSDK的uploadVoice接口上传的语音文件id
	MediaID string `json:"media_id" validate:"required"`
}

// GetMediaVoiceResponse is response of Client.GetMediaVoice
type GetMediaVoiceResponse struct {
	// Body 音频文件二进制流（speex格式）
	Body []any `json:"body"`
}

// UploadMedia 上传临时素材
// 上传临时素材（图片、语音、视频、普通文件），返回media_id
//
// see https://developer.work.weixin.qq.com/document/path/90076
func (c *Client) UploadMedia(r *UploadMediaRequest, opts ...any) (out UploadMediaResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/media/upload",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UploadMediaRequest is request of Client.UploadMedia
type UploadMediaRequest struct {
	// Type 媒体文件类型：image, voice, video, file
	Type string `json:"type" validate:"required"`
}

// UploadMediaResponse is response of Client.UploadMedia
type UploadMediaResponse struct {
	// Type 媒体文件类型
	Type string `json:"type"`
	// MediaID 媒体文件上传后获取的唯一标识，3天内有效
	MediaID string `json:"media_id"`
	// CreatedAt 媒体文件上传时间戳
	CreatedAt string `json:"created_at"`
}

// UploadMediaByUrl 异步上传临时素材
// 支持指定文件的CDN链接，由企微微信后台异步下载和处理临时素材的大文件（最高支持200M）
//
// see https://developer.work.weixin.qq.com/document/path/97078
func (c *Client) UploadMediaByUrl(r *UploadMediaByUrlRequest, opts ...any) (out UploadMediaByUrlResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/media/upload_by_url",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UploadMediaByUrlRequest is request of Client.UploadMediaByUrl
type UploadMediaByUrlRequest struct {
	// Scene 场景值。1-客户联系入群欢迎语素材
	Scene int `json:"scene" validate:"required"`
	// Type 媒体文件类型。目前仅支持video(视频)和file(普通文件)
	Type string `json:"type" validate:"required"`
	// Filename 文件名，标识文件展示的名称
	Filename string `json:"filename" validate:"required"`
	// URL 文件cdn url。必须支持Range分块下载
	URL string `json:"url" validate:"required"`
	// Md5 文件md5
	Md5 string `json:"md5" validate:"required"`
}

// UploadMediaByUrlResponse is response of Client.UploadMediaByUrl
type UploadMediaByUrlResponse struct {
	// JobID 任务id。可通过此jobid查询结果
	JobID string `json:"jobid"`
}

// UploadImage 上传图片
// 上传图片得到图片URL，该URL永久有效，仅能用于图文消息正文或欢迎语等场景。
//
// see https://developer.work.weixin.qq.com/document/path/90079
func (c *Client) UploadImage(r *UploadImageRequest, opts ...any) (out UploadImageResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/media/uploadimg",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UploadImageRequest is request of Client.UploadImage
type UploadImageRequest struct {
	// Media 媒体文件（multipart/form-data）
	Media string `json:"media" validate:"required"`
}

// UploadImageResponse is response of Client.UploadImage
type UploadImageResponse struct {
	// URL 上传后得到的图片URL，永久有效
	URL string `json:"url"`
}
