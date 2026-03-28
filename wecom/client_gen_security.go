package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// ListAdminOperLog 获取管理端操作日志
// 读取企业微信管理端操作日志，支持增量分页查询
//
// see https://developer.work.weixin.qq.com/document/path/100179
func (c *Client) ListAdminOperLog(r *ListAdminOperLogRequest, opts ...any) (out ListAdminOperLogResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/security/admin_oper_log/list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListAdminOperLogRequest is request of Client.ListAdminOperLog
type ListAdminOperLogRequest struct {
	// StartTime 开始时间，取值范围：不早于180天前
	StartTime int `json:"start_time" validate:"required"`
	// EndTime 结束时间，需大于start_time且跨度不超过7天
	EndTime int `json:"end_time" validate:"required"`
	// OperType 操作类型过滤，不填表示全部
	OperType int `json:"oper_type"`
	// UserID 操作者userid
	UserID string `json:"userid"`
	// Cursor 分页游标，不填表示首页
	Cursor string `json:"cursor"`
	// Limit 最大记录数，1 ~ 400，默认最多400个
	Limit int `json:"limit"`
}

// ListAdminOperLogResponse is response of Client.ListAdminOperLog
type ListAdminOperLogResponse struct {
	// HasMore 是否还有下一页
	HasMore bool `json:"has_more"`
	// NextCursor 下一页的分页游标
	NextCursor string `json:"next_cursor"`
	// RecordList 记录列表
	RecordList []map[string]any `json:"record_list"`
}

// GetFileOperRecord 文件防泄漏
// 查询文件上传、下载、转发等操作记录
//
// see https://developer.work.weixin.qq.com/document/path/98883
func (c *Client) GetFileOperRecord(r *GetFileOperRecordRequest, opts ...any) (out GetFileOperRecordResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/security/get_file_oper_record",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetFileOperRecordRequest is request of Client.GetFileOperRecord
type GetFileOperRecordRequest struct {
	// StartTime 开始时间
	StartTime int `json:"start_time" validate:"required"`
	// EndTime 结束时间，开始时间到结束时间的范围不能超过14天
	EndTime int `json:"end_time" validate:"required"`
	// UserIDList 需要查询的文件操作者的userid，单次最多可以传100个用户
	UserIDList []string `json:"userid_list"`
	// Operation 操作类型和来源结构
	Operation any `json:"operation"`
	// Cursor 由企业微信后台返回，第一次调用可不填
	Cursor string `json:"cursor"`
	// Limit 限制返回的条数，最多设置为1000
	Limit int `json:"limit"`
}

// GetFileOperRecordResponse is response of Client.GetFileOperRecord
type GetFileOperRecordResponse struct {
	// HasMore 是否还有更多数据
	HasMore bool `json:"has_more"`
	// NextCursor 仅has_more值为true时返回该字段，下一次调用将该值填到cursor字段，以实现分页查询
	NextCursor string `json:"next_cursor"`
	// RecordList 操作记录列表
	RecordList []map[string]any `json:"record_list"`
}

// GetScreenOperRecord 截屏/录屏管理
// 查询截屏操作记录
//
// see https://developer.work.weixin.qq.com/document/path/100128
func (c *Client) GetScreenOperRecord(r *GetScreenOperRecordRequest, opts ...any) (out GetScreenOperRecordResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/security/get_screen_oper_record",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetScreenOperRecordRequest is request of Client.GetScreenOperRecord
type GetScreenOperRecordRequest struct {
	// StartTime 开始时间
	StartTime int `json:"start_time" validate:"required"`
	// EndTime 结束时间，开始时间到结束时间的范围不能超过14天
	EndTime int `json:"end_time" validate:"required"`
	// UserIDList 需要查询的截屏操作者的userid，单次最多可以传100个用户。设置的userid需要在应用的可见范围内
	UserIDList []string `json:"userid_list"`
	// DepartmentIDList 需要查询的截屏操作者部门的department_id，单次最多可以传100个部门id。设置的department_id需要在应用的可见范围内
	DepartmentIDList []int `json:"department_id_list"`
	// ScreenShotType 截屏内容的类型，不设置默认为全部。1: 聊天, 2: 通讯录, 3: 邮件, 4: 文件, 5: 日程, 6: 其他
	ScreenShotType int `json:"screen_shot_type"`
	// Cursor 由企业微信后台返回，第一次调用可不填
	Cursor string `json:"cursor"`
	// Limit 限制返回的条数，最多设置为1000
	Limit int `json:"limit"`
}

// GetScreenOperRecordResponse is response of Client.GetScreenOperRecord
type GetScreenOperRecordResponse struct {
	// HasMore 是否还有更多数据
	HasMore bool `json:"has_more"`
	// NextCursor 仅has_more值为true时返回该字段，下一次调用将该值填到cursor字段，以实现分页查询
	NextCursor string `json:"next_cursor"`
	// RecordList 操作记录列表
	RecordList []map[string]any `json:"record_list"`
}

// GetServerDomainIp 获取企业微信域名IP信息
// 获取最新的域名和IP信息，用于应对SaaS服务中域名和IP的变动
//
// see https://developer.work.weixin.qq.com/document/path/100079
func (c *Client) GetServerDomainIp(r *GetServerDomainIpRequest, opts ...any) (out GetServerDomainIpResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/security/get_server_domain_ip",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetServerDomainIpRequest is request of Client.GetServerDomainIp
type GetServerDomainIpRequest struct {}

// GetServerDomainIpResponse is response of Client.GetServerDomainIp
type GetServerDomainIpResponse struct {
	// DomainList 域名列表
	DomainList []map[string]any `json:"domain_list"`
	// IPList IP列表
	IPList []map[string]any `json:"ip_list"`
}

// ListMemberOperLog 获取成员操作记录
// 读取成员操作记录，用于定时安全预警
//
// see https://developer.work.weixin.qq.com/document/path/100178
func (c *Client) ListMemberOperLog(r *ListMemberOperLogRequest, opts ...any) (out ListMemberOperLogResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/security/member_oper_log/list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMemberOperLogRequest is request of Client.ListMemberOperLog
type ListMemberOperLogRequest struct {
	// StartTime 开始时间，不早于180天前
	StartTime int `json:"start_time" validate:"required"`
	// EndTime 结束时间，大于start_time且跨度不超过7天
	EndTime int `json:"end_time" validate:"required"`
	// OperType 操作类型过滤，不填表示全部
	OperType int `json:"oper_type"`
	// UserID 操作者userid过滤
	UserID string `json:"userid"`
	// Cursor 分页游标，不填表示首页
	Cursor string `json:"cursor"`
	// Limit 最大记录数，1~400
	Limit int `json:"limit"`
}

// ListMemberOperLogResponse is response of Client.ListMemberOperLog
type ListMemberOperLogResponse struct {
	// HasMore 是否还有下一页
	HasMore bool `json:"has_more"`
	// NextCursor 下一页的分页游标
	NextCursor string `json:"next_cursor"`
	// RecordList 记录列表
	RecordList []map[string]any `json:"record_list"`
}

// ImportTrustDevice 导入可信企业设备
// 通过相关接口导入可信企业设备，获取/删除可信企业设备、可信个人设备、未知设备，并对未知设备的归属申请进行确认或驳回。
//
// see https://developer.work.weixin.qq.com/document/path/98920
func (c *Client) ImportTrustDevice(r *ImportTrustDeviceRequest, opts ...any) (out ImportTrustDeviceResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/security/trustdevice/import",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ImportTrustDeviceRequest is request of Client.ImportTrustDevice
type ImportTrustDeviceRequest struct {
	// DeviceList 设备列表
	DeviceList []map[string]any `json:"device_list" validate:"required"`
}

// ImportTrustDeviceResponse is response of Client.ImportTrustDevice
type ImportTrustDeviceResponse struct {
	// Result 导入结果列表
	Result []map[string]any `json:"result"`
}

// ListSecurityVip 获取高级功能账号列表
// 查询企业已分配高级功能且在应用可见范围的账号列表
//
// see https://developer.work.weixin.qq.com/document/path/99506
func (c *Client) ListSecurityVip(r *ListSecurityVipRequest, opts ...any) (out ListSecurityVipResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/security/vip/list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListSecurityVipRequest is request of Client.ListSecurityVip
type ListSecurityVipRequest struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"`
	// Limit 用于分页查询，每次请求返回的数据上限。默认100，最大200
	Limit int `json:"limit"`
}

// ListSecurityVipResponse is response of Client.ListSecurityVip
type ListSecurityVipResponse struct {
	// HasMore 是否还有更多数据未获取
	HasMore bool `json:"has_more"`
	// NextCursor 下一次请求的cursor值
	NextCursor string `json:"next_cursor"`
	// UserIDList 符合条件的企业成员userid列表
	UserIDList []string `json:"userid_list"`
}

// SubmitBatchAddJob 分配高级功能账号
// 该接口用于分配应用可见范围内企业成员的高级功能。
//
// see https://developer.work.weixin.qq.com/document/path/99503
func (c *Client) SubmitBatchAddJob(r *SubmitBatchAddJobRequest, opts ...any) (out SubmitBatchAddJobResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/security/vip/submit_batch_add_job",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SubmitBatchAddJobRequest is request of Client.SubmitBatchAddJob
type SubmitBatchAddJobRequest struct {
	// UserIDList 要分配高级功能的企业成员userid列表，单次操作最大限制100个
	UserIDList []string `json:"userid_list" validate:"required"`
}

// SubmitBatchAddJobResponse is response of Client.SubmitBatchAddJob
type SubmitBatchAddJobResponse struct {
	// JobID 批量分配高级功能的任务id
	JobID string `json:"jobid"`
	// InvalidUserIDList 非法的userid 列表，不在应用可见范围的useri以及无法识别的userid
	InvalidUserIDList []string `json:"invalid_userid_list"`
}

// BatchDelSecurityVip 取消高级功能账号
// 该接口用于撤销分配应用可见范围企业成员的高级功能。
//
// see https://developer.work.weixin.qq.com/document/path/99505
func (c *Client) BatchDelSecurityVip(r *BatchDelSecurityVipRequest, opts ...any) (out BatchDelSecurityVipResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/security/vip/submit_batch_del_job",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchDelSecurityVipRequest is request of Client.BatchDelSecurityVip
type BatchDelSecurityVipRequest struct {
	// UserIDList 要撤销分配高级功能的企业成员userid列表，单次操作最多限制100个
	UserIDList []string `json:"userid_list" validate:"required"`
}

// BatchDelSecurityVipResponse is response of Client.BatchDelSecurityVip
type BatchDelSecurityVipResponse struct {
	// JobID 批量取消高级功能的任务id
	JobID string `json:"jobid"`
	// InvalidUserIDList 非法的userid 列表，不在应用可见范围的useri以及无法识别的userid
	InvalidUserIDList []string `json:"invalid_userid_list"`
}

