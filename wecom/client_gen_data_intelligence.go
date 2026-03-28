package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// CreateProgramTask 创建专区程序调用任务
// 创建应用异步调用专区程序的调用任务
//
// see https://developer.work.weixin.qq.com/document/path/99966
func (c *Client) CreateProgramTask(r *CreateProgramTaskRequest, opts ...any) (out CreateProgramTaskResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/chatdata/async_program_task",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateProgramTaskRequest is request of Client.CreateProgramTask
type CreateProgramTaskRequest struct {
	// ProgramID 应用关联的程序id
	ProgramID string `json:"program_id" validate:"required"`
	// AbilityID 程序关联的能力id
	AbilityID string `json:"ability_id" validate:"required"`
	// RequestData 请求的输入JSON，要求与配置的格式匹配
	RequestData string `json:"request_data" validate:"required"`
}

// CreateProgramTaskResponse is response of Client.CreateProgramTask
type CreateProgramTaskResponse struct {
	// JobID 任务id
	JobID string `json:"jobid"`
}

// CheckDebugMode 获取专区调试模式状态
// 查询应用关联程序的调试模式状态
//
// see https://developer.work.weixin.qq.com/document/path/100113
func (c *Client) CheckDebugMode(r *CheckDebugModeRequest, opts ...any) (out CheckDebugModeResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/chatdata/check_debug_mode",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CheckDebugModeRequest is request of Client.CheckDebugMode
type CheckDebugModeRequest struct {
	// ProgramID 应用关联的程序id
	ProgramID string `json:"program_id" validate:"required"`
}

// CheckDebugModeResponse is response of Client.CheckDebugMode
type CheckDebugModeResponse struct {
	// DebugModeStatus 程序当前的调试模式状态，1为关闭，2为开启
	DebugModeStatus int `json:"debug_mode_status"`
}

// CloseDebugMode 关闭专区调试模式
// 关闭数据与智能专区的调试模式
//
// see https://developer.work.weixin.qq.com/document/path/100088
func (c *Client) CloseDebugMode(r *CloseDebugModeRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/chatdata/close_debug_mode",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CloseDebugModeRequest is request of Client.CloseDebugMode
type CloseDebugModeRequest struct {
	// ProgramID 应用关联的程序id
	ProgramID string `json:"program_id" validate:"required"`
}

// GetAuthUserList 获取授权存档的成员列表
// 可通过该接口获取授权企业所有生效中的成员列表。（企业会话内容中授权的生效成员）
//
// see https://developer.work.weixin.qq.com/document/path/99962
func (c *Client) GetAuthUserList(r *GetAuthUserListRequest, opts ...any) (out GetAuthUserListResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/chatdata/get_auth_user_list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetAuthUserListRequest is request of Client.GetAuthUserList
type GetAuthUserListRequest struct {
	// Cursor 上一次调用时返回的next_cursor，第一次拉取可以不填
	Cursor string `json:"cursor"`
	// Limit 本次查询返回的最大条数。不超过1000，默认200条
	Limit int `json:"limit"`
}

// GetAuthUserListResponse is response of Client.GetAuthUserList
type GetAuthUserListResponse struct {
	// AuthUserList 生效成员列表，详见AuthUser
	AuthUserList []map[string]any `json:"auth_user_list"`
	// NextCursor 下一次查询时使用，将值填到请求包的cursor字段中
	NextCursor string `json:"next_cursor"`
	// HasMore 是否还有更多未拉取的数据，1：是；0：否
	HasMore bool `json:"has_more"`
}

// OpenChatDataDebugMode 开启专区调试模式
// 应用通过接口将指定的程序开启调试模式
//
// see https://developer.work.weixin.qq.com/document/path/100087
func (c *Client) OpenChatDataDebugMode(r *OpenChatDataDebugModeRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/chatdata/open_debug_mode",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// OpenChatDataDebugModeRequest is request of Client.OpenChatDataDebugMode
type OpenChatDataDebugModeRequest struct {
	// ProgramID 应用关联的程序id
	ProgramID string `json:"program_id" validate:"required"`
	// DebugToken 程序的调试凭证
	DebugToken string `json:"debug_token" validate:"required"`
}

// SetChatdataHideSensitiveInfoConfig 设置成员会话组件敏感信息隐藏配置
// 可通过此接口，设置成员使用会话组件时，敏感信息是否打星。敏感信息包括：手机号、银行卡号、身份证号。如果未调用接口开启展示敏感信息，会话组件默认为不打星
//
// see https://developer.work.weixin.qq.com/document/path/100139
func (c *Client) SetChatdataHideSensitiveInfoConfig(r *SetChatdataHideSensitiveInfoConfigRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/chatdata/set_hide_sensitiveinfo_config",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetChatdataHideSensitiveInfoConfigRequest is request of Client.SetChatdataHideSensitiveInfoConfig
type SetChatdataHideSensitiveInfoConfigRequest struct {
	// UserID 成员的userid
	UserID string `json:"userid" validate:"required"`
	// Config 敏感信息隐藏配置
	Config any `json:"config" validate:"required"`
}

// SetLogLevel 设置日志打印级别
// 设置指定模型或程序的日志打印级别。设置之后，专区程序只会存储不高于该级别的日志。
//
// see https://developer.work.weixin.qq.com/document/path/100108
func (c *Client) SetLogLevel(r *SetLogLevelRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/chatdata/set_log_level",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetLogLevelRequest is request of Client.SetLogLevel
type SetLogLevelRequest struct {
	// ProgramID 应用关联的程序id
	ProgramID string `json:"program_id" validate:"required"`
	// LogLevel 日志级别。1-ERR, 2-INFO, 3-DBG。指定后，仅会存储不高于该级别的日志。
	LogLevel int `json:"log_level" validate:"required"`
}

// SetPublicKey 设置公钥
// 授权完成后调用此接口上传公钥，消息存档将从此时开始生效。
//
// see https://developer.work.weixin.qq.com/document/path/99961
func (c *Client) SetPublicKey(r *SetPublicKeyRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/chatdata/set_public_key",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetPublicKeyRequest is request of Client.SetPublicKey
type SetPublicKeyRequest struct {
	// PublicKey 开发者为该企业生成的公钥，需包含BEGIN/END标识及换行符转义
	PublicKey string `json:"public_key" validate:"required"`
	// PublicKeyVer 公钥对应的版本号，更换公钥时需比旧公钥版本号大
	PublicKeyVer int `json:"public_key_ver" validate:"required"`
}

// SetReceiveCallback 设置专区接收回调事件
// 应用可通过该接口设置专区中接收回调的程序id
//
// see https://developer.work.weixin.qq.com/document/path/99963
func (c *Client) SetReceiveCallback(r *SetReceiveCallbackRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/chatdata/set_receive_callback",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetReceiveCallbackRequest is request of Client.SetReceiveCallback
type SetReceiveCallbackRequest struct {
	// ProgramID 应用关联的程序id，同一个应用只能设置一个程序接收。若先设置了程序A接收，再调用该接口设置程序B时，会更改为程序B接收
	ProgramID string `json:"program_id" validate:"required"`
}

// SyncCallProgram 应用同步调用专区程序
// 通过API同步调用企业微信数据与智能专区的程序
//
// see https://developer.work.weixin.qq.com/document/path/99965
func (c *Client) SyncCallProgram(r *SyncCallProgramRequest, opts ...any) (out SyncCallProgramResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/chatdata/sync_call_program",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SyncCallProgramRequest is request of Client.SyncCallProgram
type SyncCallProgramRequest struct {
	// ProgramID 应用关联的程序id
	ProgramID string `json:"program_id" validate:"required"`
	// AbilityID 程序关联的能力id
	AbilityID string `json:"ability_id" validate:"required"`
	// NotifyID 通知id。由专区通知应用返回
	NotifyID string `json:"notify_id"`
	// RequestData 请求的输入JSON，要求与配置的输入协议格式匹配
	RequestData string `json:"request_data" validate:"required"`
}

// SyncCallProgramResponse is response of Client.SyncCallProgram
type SyncCallProgramResponse struct {
	// ResponseData 专区程序的输出结果，为自定义的JSON字符串
	ResponseData string `json:"response_data"`
}

// UploadChatDataMedia 上传临时文件到专区
// 该接口将文件上传到数据专区临时存储，得到media_id。media_id仅三天内有效，不能跨企业或应用使用，仅可用于数据专区接口或SDK。
//
// see https://developer.work.weixin.qq.com/document/path/100174
func (c *Client) UploadChatDataMedia(r *UploadChatDataMediaRequest, opts ...any) (out UploadChatDataMediaResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/chatdata/upload_media",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UploadChatDataMediaRequest is request of Client.UploadChatDataMedia
type UploadChatDataMediaRequest struct {
	// Type 文件类型，目前仅支持普通文件：file
	Type string `json:"type" validate:"required"`
}

// UploadChatDataMediaResponse is response of Client.UploadChatDataMedia
type UploadChatDataMediaResponse struct {
	// Type 文件类型，目前仅支持普通文件：file
	Type string `json:"type"`
	// MediaID 文件上传后获取的唯一标识，3天内有效
	MediaID string `json:"media_id"`
	// CreatedAt 文件上传时间戳
	CreatedAt string `json:"created_at"`
}
