package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// DownloadWeDriveFile 下载微盘文件
// 该接口仅可用于下载汇报中关联的微盘文件。
//
// see https://developer.work.weixin.qq.com/document/path/98021
func (c *Client) DownloadWeDriveFile(r *DownloadWeDriveFileRequest, opts ...any) (out DownloadWeDriveFileResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/journal/download_wedrive_file",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DownloadWeDriveFileRequest is request of Client.DownloadWeDriveFile
type DownloadWeDriveFileRequest struct {
	// Journaluuid 汇报记录id
	Journaluuid string `json:"journaluuid" validate:"required"`
	// Fileid 微盘fileid。必须是journaluuid对应的汇报关联的wedrive_files
	Fileid string `json:"fileid" validate:"required"`
}

// DownloadWeDriveFileResponse is response of Client.DownloadWeDriveFile
type DownloadWeDriveFileResponse struct {
	// DownloadURL 下载请求url (有效期2个小时)
	DownloadURL string `json:"download_url"`
	// CookieName 下载请求带cookie的key
	CookieName string `json:"cookie_name"`
	// CookieValue 下载请求带cookie的value
	CookieValue string `json:"cookie_value"`
}

// GetJournalDetail 获取汇报记录详情
// 根据汇报记录单号查询企业微信“汇报应用”的汇报详情。
//
// see https://developer.work.weixin.qq.com/document/path/93394
func (c *Client) GetJournalDetail(r *GetJournalDetailRequest, opts ...any) (out GetJournalDetailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/journal/get_record_detail",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetJournalDetailRequest is request of Client.GetJournalDetail
type GetJournalDetailRequest struct {
	// Journaluuid 汇报记录单号，不多于256字节
	Journaluuid string `json:"journaluuid" validate:"required"`
}

// GetJournalDetailResponse is response of Client.GetJournalDetail
type GetJournalDetailResponse struct {
	// Info 汇报详情对象
	Info any `json:"info"`
}

// ListJournalRecord 批量获取汇报记录单号
// 企业可通过本接口获取企业一段时间内企业微信“汇报应用”汇报记录编号，支持按汇报表单ID、申请人、部门等条件筛选。
//
// see https://developer.work.weixin.qq.com/document/path/93474
func (c *Client) ListJournalRecord(r *ListJournalRecordRequest, opts ...any) (out ListJournalRecordResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/journal/get_record_list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListJournalRecordRequest is request of Client.ListJournalRecord
type ListJournalRecordRequest struct {
	// Starttime 开始时间
	Starttime int `json:"starttime" validate:"required"`
	// Endtime 结束时间，开始时间和结束时间间隔不能超过一个月
	Endtime int `json:"endtime" validate:"required"`
	// Cursor 游标首次请求传0，非首次请求携带上一次请求返回的next_cursor
	Cursor int `json:"cursor" validate:"required"`
	// Limit 拉取条数
	Limit int `json:"limit" validate:"required"`
	// Filters 过滤条件列表
	Filters []map[string]any `json:"filters"`
}

// ListJournalRecordResponse is response of Client.ListJournalRecord
type ListJournalRecordResponse struct {
	// JournaluuidList 汇报记录id列表
	JournaluuidList []string `json:"journaluuid_list"`
	// NextCursor 下一次拉取游标
	NextCursor int `json:"next_cursor"`
	// Endflag 0代表还有数据，1代表已无数据
	Endflag int `json:"endflag"`
}

// ListJournalStat 获取汇报统计数据
// 企业可通过调用本接口，根据汇报表单id查询企业微信“汇报应用”的汇报统计汇总信息。该接口只能拉取到已经汇总的统计数据，对于尚未完成汇总的周期不会返回。
//
// see https://developer.work.weixin.qq.com/document/path/93475
func (c *Client) ListJournalStat(r *ListJournalStatRequest, opts ...any) (out ListJournalStatResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/journal/get_stat_list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListJournalStatRequest is request of Client.ListJournalStat
type ListJournalStatRequest struct {
	// TemplateID 汇报表单id，不多于256字节
	TemplateID string `json:"template_id" validate:"required"`
	// Starttime 开始时间
	Starttime int `json:"starttime" validate:"required"`
	// Endtime 结束时间，时间区间最大长度为一年
	Endtime int `json:"endtime" validate:"required"`
}

// ListJournalStatResponse is response of Client.ListJournalStat
type ListJournalStatResponse struct {
	// StatList 统计数据列表
	StatList []map[string]any `json:"stat_list"`
	// TemplateID 汇报表单id
	TemplateID string `json:"template_id"`
	// TemplateName 汇报表单名称
	TemplateName string `json:"template_name"`
	// ReportRange 汇报人员范围
	ReportRange any `json:"report_range"`
	// UserList 指定人集合
	UserList []any `json:"user_list"`
	// UserID 用户id
	UserID string `json:"userid"`
	// PartyList 指定部门集合
	PartyList []any `json:"party_list"`
	// OpenPartyid 部门id
	OpenPartyid string `json:"open_partyid"`
	// TagList 指定标签集合
	TagList []any `json:"tag_list"`
	// OpenTagID 标签id
	OpenTagID string `json:"open_tagid"`
	// WhiteRange 白名单集合
	WhiteRange any `json:"white_range"`
	// Receivers 固定汇报对象
	Receivers any `json:"receivers"`
	// LeaderList 指定上级集合
	LeaderList []any `json:"leader_list"`
	// Level 上级级别从1开始
	Level string `json:"level"`
	// CycleBeginTime 周期开始时间
	CycleBeginTime string `json:"cycle_begin_time"`
	// CycleEndTime 周期结束时间
	CycleEndTime string `json:"cycle_end_time"`
	// StatBeginTime 统计开始时间
	StatBeginTime string `json:"stat_begin_time"`
	// StatEndTime 统计结束时间
	StatEndTime string `json:"stat_end_time"`
	// ReportList 已汇报用户列表
	ReportList []any `json:"report_list"`
	// User 汇报用户/未汇报用户
	User any `json:"user"`
	// Itemlist 汇报记录列表
	Itemlist []any `json:"itemlist"`
	// Journaluuid 汇报记录id
	Journaluuid string `json:"journaluuid"`
	// Reporttime 汇报时间
	Reporttime string `json:"reporttime"`
	// Flag 是否迟交，1迟交;0非迟交
	Flag string `json:"flag"`
	// UnreportList 未汇报用户列表
	UnreportList []any `json:"unreport_list"`
	// ReportType 汇报方式：2按日汇报; 3按周汇报; 4按月汇报
	ReportType string `json:"report_type"`
}
