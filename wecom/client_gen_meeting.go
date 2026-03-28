package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// AddMeetingAdvancedLayout 添加会议高级布局
// 对当前会议添加高级布局，支持批量添加。单个会议最多允许添加20个高级布局。
//
// see https://developer.work.weixin.qq.com/document/path/99302
func (c *Client) AddMeetingAdvancedLayout(r *AddMeetingAdvancedLayoutRequest, opts ...any) (out AddMeetingAdvancedLayoutResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/advanced_layout/add",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddMeetingAdvancedLayoutRequest is request of Client.AddMeetingAdvancedLayout
type AddMeetingAdvancedLayoutRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// LayoutList 布局对象列表
	LayoutList []map[string]any `json:"layout_list" validate:"required"`
}

// AddMeetingAdvancedLayoutResponse is response of Client.AddMeetingAdvancedLayout
type AddMeetingAdvancedLayoutResponse struct {
	// LayoutList 布局对象列表
	LayoutList []map[string]any `json:"layout_list"`
}

// ApplyMeetingAdvancedLayout 设置高级布局
// 将会议中的高级自定义布局应用到指定成员或者整个会议。也可以恢复指定成员或整个会议的默认布局。
//
// see https://developer.work.weixin.qq.com/document/path/99304
func (c *Client) ApplyMeetingAdvancedLayout(r *ApplyMeetingAdvancedLayoutRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/advanced_layout/apply",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ApplyMeetingAdvancedLayoutRequest is request of Client.ApplyMeetingAdvancedLayout
type ApplyMeetingAdvancedLayoutRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// LayoutID 选择应用的布局 ID（若传空""，表示恢复成当前会议的默认布局）
	LayoutID string `json:"layout_id" validate:"required"`
	// UserList 用户列表对象数组。如果该字段为空，为会议设置高级自定义布局；如果该字段携带用户，则只为指定用户设置个性布局。单次最多支持20个用户
	UserList []map[string]any `json:"user_list"`
}

// BatchDeleteLayout 批量删除布局
// 根据布局 ID 批量删除布局，可以删除基础布局和高级布局。正在被应用的布局无法删除，请先设置成其他布局或恢复成默认原始布局后再行删除。接口不做布局是否存在的校验，删除不存在的布局不会有提示。
//
// see https://developer.work.weixin.qq.com/document/path/99261
func (c *Client) BatchDeleteLayout(r *BatchDeleteLayoutRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/advanced_layout/batch_delete",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchDeleteLayoutRequest is request of Client.BatchDeleteLayout
type BatchDeleteLayoutRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// LayoutIDList 布局 ID 列表，要删除的一个或多个布局 ID（最多支持20个）。当前正在应用的布局不能被删除
	LayoutIDList []string `json:"layout_id_list" validate:"required"`
}

// GetUserLayout 获取用户布局
// 根据会议 ID 和用户 ID 返回用户的布局设置信息。
//
// see https://developer.work.weixin.qq.com/document/path/99260
func (c *Client) GetUserLayout(r *GetUserLayoutRequest, opts ...any) (out GetUserLayoutResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/advanced_layout/get_user_layout",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetUserLayoutRequest is request of Client.GetUserLayout
type GetUserLayoutRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// TmpOpenID 被操作用户临时身份 ID
	TmpOpenID string `json:"tmp_openid" validate:"required"`
	// InstanceID 被操作用户终端设备类型 ID
	InstanceID int `json:"instance_id" validate:"required"`
}

// GetUserLayoutResponse is response of Client.GetUserLayout
type GetUserLayoutResponse struct {
	// SelectedLayoutID 会议应用的布局 ID
	SelectedLayoutID string `json:"selected_layout_id"`
	// LayoutName 布局名称
	LayoutName string `json:"layout_name"`
	// LayoutType 布局类型
	LayoutType int `json:"layout_type"`
	// PageList 布局单页对象列表
	PageList []map[string]any `json:"page_list"`
}

// ListMeetingAdvancedLayout 获取会议布局列表
// 根据会议 ID 返回会议的基础和高级自定义布局信息列表。高级布局目前仅支持 H.323/SIP 会议室终端。
//
// see https://developer.work.weixin.qq.com/document/path/99259
func (c *Client) ListMeetingAdvancedLayout(r *ListMeetingAdvancedLayoutRequest, opts ...any) (out ListMeetingAdvancedLayoutResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/advanced_layout/list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMeetingAdvancedLayoutRequest is request of Client.ListMeetingAdvancedLayout
type ListMeetingAdvancedLayoutRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
}

// ListMeetingAdvancedLayoutResponse is response of Client.ListMeetingAdvancedLayout
type ListMeetingAdvancedLayoutResponse struct {
	// SelectedLayoutID 会议应用的布局 ID
	SelectedLayoutID string `json:"selected_layout_id"`
	// LayoutList 布局对象列表
	LayoutList []map[string]any `json:"layout_list"`
}

// UpdateMeetingAdvancedLayout 修改会议高级布局
// 对会议中的高级布局进行修改，接口仅支持全量更新。
//
// see https://developer.work.weixin.qq.com/document/path/99303
func (c *Client) UpdateMeetingAdvancedLayout(r *UpdateMeetingAdvancedLayoutRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/advanced_layout/update",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateMeetingAdvancedLayoutRequest is request of Client.UpdateMeetingAdvancedLayout
type UpdateMeetingAdvancedLayoutRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// LayoutID 布局 ID
	LayoutID string `json:"layout_id" validate:"required"`
	// LayoutName 布局名称
	LayoutName string `json:"layout_name"`
	// PageList 布局单页对象列表
	PageList []map[string]any `json:"page_list" validate:"required"`
}

// CancelMeeting 取消预约会议
// 该接口用于取消一个指定的预约会议。
//
// see https://developer.work.weixin.qq.com/document/path/98990
func (c *Client) CancelMeeting(r *CancelMeetingRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/cancel",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CancelMeetingRequest is request of Client.CancelMeeting
type CancelMeetingRequest struct {
	// Meetingid 会议id，仅允许取消预约状态下的会议
	Meetingid string `json:"meetingid" validate:"required"`
	// SubMeetingid 周期性子会议 ID。如果取消周期性会议且该字段不传，则会取消该系列的周期性会议。
	SubMeetingid string `json:"sub_meetingid"`
}

// CheckMeetingDevice 获取成员设备是否入会
// 查询企业内成员在当前时间前是否有设备进入指定的会议中
//
// see https://developer.work.weixin.qq.com/document/path/99013
func (c *Client) CheckMeetingDevice(r *CheckMeetingDeviceRequest, opts ...any) (out CheckMeetingDeviceResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/check_device_in_meeting",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CheckMeetingDeviceRequest is request of Client.CheckMeetingDevice
type CheckMeetingDeviceRequest struct {
	// UserID 企业成员的userid
	UserID string `json:"userid" validate:"required"`
	// InstanceIDList 终端设备类型列表，该参数不带，则会查询所有设备上的会议信息，带则表示查询指定设备。0：PSTN, 1：PC, 2：Mac, 3：Android, 4：iOS, 5：Web, 6：iPad, 7：...
	InstanceIDList []int `json:"instance_id_list"`
	// MeetingidList 会议 ID 列表。须为本企业创建的会议
	MeetingidList []string `json:"meetingid_list" validate:"required"`
}

// CheckMeetingDeviceResponse is response of Client.CheckMeetingDevice
type CheckMeetingDeviceResponse struct {
	// ResultList 结果列表，包含meetingid和instance_id
	ResultList []map[string]any `json:"result_list"`
}

// CreateMeeting 创建预约会议
// 该接口用于创建一个预约会议。
//
// see https://developer.work.weixin.qq.com/document/path/99104
func (c *Client) CreateMeeting(r *CreateMeetingRequest, opts ...any) (out CreateMeetingResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/create",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateMeetingRequest is request of Client.CreateMeeting
type CreateMeetingRequest struct {
	// AdminUserID 会议管理员userid
	AdminUserID string `json:"admin_userid" validate:"required"`
	// Title 会议的标题，最多支持40个字节或者20个utf8字符
	Title string `json:"title" validate:"required"`
	// MeetingStart 会议开始时间的unix时间戳。需大于当前时间
	MeetingStart int `json:"meeting_start" validate:"required"`
	// MeetingDuration 会议持续时间（单位秒），最小300秒，最大86399秒
	MeetingDuration int `json:"meeting_duration" validate:"required"`
	// Description 会议的描述，最多支持500个字节或者utf8字符
	Description string `json:"description"`
	// Location 会议地点，最多128个字符
	Location string `json:"location"`
	// AgentID 授权方安装的应用agentid。仅旧的第三方多应用套件需要填此参数
	AgentID int `json:"agentid"`
	// Invitees 邀请参会的成员。包含userid列表
	Invitees any `json:"invitees"`
	// CalID 会议所属日历ID。第三方应用必须指定cal_id
	CalID string `json:"cal_id"`
	// Settings 会议配置，详见Settings表
	Settings any `json:"settings"`
	// Reminders 重复会议相关配置，详见Reminders表
	Reminders any `json:"reminders"`
}

// CreateMeetingResponse is response of Client.CreateMeeting
type CreateMeetingResponse struct {
	// Meetingid 会议id，通过此id可调用进入会议接口
	Meetingid string `json:"meetingid"`
	// ExcessUsers 参会人中包含无效会议账号的userid，仅在购买会议专业版企业由于部分参会人无有效会议账号时返回
	ExcessUsers []string `json:"excess_users"`
}

// CreateCustomerShortUrl 创建用户专属参会链接
// 为会议生成多个专属入会链接，不同链接以 customer_data 进行区分。
//
// see https://developer.work.weixin.qq.com/document/path/98818
func (c *Client) CreateCustomerShortUrl(r *CreateCustomerShortUrlRequest, opts ...any) (out CreateCustomerShortUrlResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/create_customer_short_url",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateCustomerShortUrlRequest is request of Client.CreateCustomerShortUrl
type CreateCustomerShortUrlRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// CustomerData 用户专属字段，长度不超过256字节。需以 {"ver": "1.0", "userData":"自定义字段"} 的结构进行 Base64编码
	CustomerData string `json:"customer_data" validate:"required"`
}

// CreateCustomerShortUrlResponse is response of Client.CreateCustomerShortUrl
type CreateCustomerShortUrlResponse struct {
	// MeetingShortURLCustomerData 用户专属参会链接对象列表
	MeetingShortURLCustomerData []map[string]any `json:"meeting_short_url_customer_data"`
}

// ApproveMeetingEnroll 审批会议报名信息
// 批量审批会议的报名信息
//
// see https://developer.work.weixin.qq.com/document/path/98807
func (c *Client) ApproveMeetingEnroll(r *ApproveMeetingEnrollRequest, opts ...any) (out ApproveMeetingEnrollResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/enroll/approve",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ApproveMeetingEnrollRequest is request of Client.ApproveMeetingEnroll
type ApproveMeetingEnrollRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// Action 审批动作：1：取消批准，2：拒绝，3：批准
	Action int `json:"action" validate:"required"`
	// EnrollIDList 报名ID列表
	EnrollIDList []string `json:"enroll_id_list" validate:"required"`
}

// ApproveMeetingEnrollResponse is response of Client.ApproveMeetingEnroll
type ApproveMeetingEnrollResponse struct {
	// HandledCount 成功处理的数量
	HandledCount int `json:"handled_count"`
}

// DeleteMeetingEnroll 删除会议报名信息
// 删除指定会议的报名信息，支持删除成员手动报名的信息和导入的报名信息。
//
// see https://developer.work.weixin.qq.com/document/path/98817
func (c *Client) DeleteMeetingEnroll(r *DeleteMeetingEnrollRequest, opts ...any) (out DeleteMeetingEnrollResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/enroll/delete",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteMeetingEnrollRequest is request of Client.DeleteMeetingEnroll
type DeleteMeetingEnrollRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// EnrollIDList 报名ID列表
	EnrollIDList []map[string]any `json:"enroll_id_list" validate:"required"`
}

// DeleteMeetingEnrollResponse is response of Client.DeleteMeetingEnroll
type DeleteMeetingEnrollResponse struct {
	// TotalCount 成功删除的报名信息数量
	TotalCount int `json:"total_count"`
}

// GetMeetingEnrollConfig 获取会议报名配置
// 获取会议的报名配置和报名问题。
//
// see https://developer.work.weixin.qq.com/document/path/99055
func (c *Client) GetMeetingEnrollConfig(r *GetMeetingEnrollConfigRequest, opts ...any) (out GetMeetingEnrollConfigResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/enroll/get_config",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMeetingEnrollConfigRequest is request of Client.GetMeetingEnrollConfig
type GetMeetingEnrollConfigRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
}

// GetMeetingEnrollConfigResponse is response of Client.GetMeetingEnrollConfig
type GetMeetingEnrollConfigResponse struct {
	// ApproveType 审批类型：1：自动审批，2：手动审批
	ApproveType int `json:"approve_type"`
	// IsCollectQuestion 是否收集问题：1：不收集，2：收集
	IsCollectQuestion int `json:"is_collect_question"`
	// NoRegistrationNeededForStaff 本企业成员无需报名：true为无需报名
	NoRegistrationNeededForStaff bool `json:"no_registration_needed_for_staff"`
	// QuestionList 报名问题列表
	QuestionList []map[string]any `json:"question_list"`
}

// ImportMeetingEnroll 导入会议报名信息
// 指定会议中导入报名信息。仅配置在“可调用接口的应用”列表中的自建应用可调用，且仅允许修改该应用创建的会议的数据。
//
// see https://developer.work.weixin.qq.com/document/path/98816
func (c *Client) ImportMeetingEnroll(r *ImportMeetingEnrollRequest, opts ...any) (out ImportMeetingEnrollResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/enroll/import",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ImportMeetingEnrollRequest is request of Client.ImportMeetingEnroll
type ImportMeetingEnrollRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// EnrollList 报名成员列表
	EnrollList []map[string]any `json:"enroll_list" validate:"required"`
}

// ImportMeetingEnrollResponse is response of Client.ImportMeetingEnroll
type ImportMeetingEnrollResponse struct {
	// TotalCount 成功导入的报名信息条数
	TotalCount int `json:"total_count"`
	// EnrollList 报名成员列表
	EnrollList []map[string]any `json:"enroll_list"`
}

// ListMeetingEnroll 获取会议报名信息
// 获取已报名成员数量和报名成员答题详情
//
// see https://developer.work.weixin.qq.com/document/path/99054
func (c *Client) ListMeetingEnroll(r *ListMeetingEnrollRequest, opts ...any) (out ListMeetingEnrollResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/enroll/list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMeetingEnrollRequest is request of Client.ListMeetingEnroll
type ListMeetingEnrollRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// Status 审批状态筛选字段，默认返回全部。0：全部。1：待审批。2：已拒绝。3：已批准。
	Status int `json:"status"`
	// Cursor 分页查询用，将上一个请求返回的next_cursor字段传入。第一次查询时可不传值
	Cursor string `json:"cursor"`
	// Limit 分页大小，最大50条。默认值为50
	Limit int `json:"limit"`
}

// ListMeetingEnrollResponse is response of Client.ListMeetingEnroll
type ListMeetingEnrollResponse struct {
	// HasMore 是否还有待拉取的成员列表
	HasMore bool `json:"has_more"`
	// NextCursor 分页用，下一次拉取列表将该字段填入cursor字段
	NextCursor string `json:"next_cursor"`
	// EnrollList 当前页的报名列表
	EnrollList []map[string]any `json:"enroll_list"`
}

// QueryMeetingEnrollIdByTmpOpenid 获取会议成员报名 ID
// 支持查询会议中已报名成员的报名 ID。可通过会中成员的 tmp_openid 查询到对应的报名 ID，成员的 tmp_openid 可通过获取参会成员列表接口获取。
//
// see https://developer.work.weixin.qq.com/document/path/99014
func (c *Client) QueryMeetingEnrollIdByTmpOpenid(r *QueryMeetingEnrollIdByTmpOpenidRequest, opts ...any) (out QueryMeetingEnrollIdByTmpOpenidResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/enroll/query_by_tmp_openid",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// QueryMeetingEnrollIdByTmpOpenidRequest is request of Client.QueryMeetingEnrollIdByTmpOpenid
type QueryMeetingEnrollIdByTmpOpenidRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// SortingRules 查询报名 ID 的排序规则。1：优先查询手机号导入报名，再查询成员手动报名，默认值；2：优先查询成员手动报名，再查手机号导入。
	SortingRules int `json:"sorting_rules"`
	// TmpOpenIDList 当场会议的成员临时 ID 数组，单次最多支持500条。
	TmpOpenIDList []string `json:"tmp_openid_list" validate:"required"`
}

// QueryMeetingEnrollIdByTmpOpenidResponse is response of Client.QueryMeetingEnrollIdByTmpOpenid
type QueryMeetingEnrollIdByTmpOpenidResponse struct {
	// EnrollIDList 成员报名 ID 数组，仅返回已报名成员的报名 ID，若传入的成员无人报名，则无该字段。
	EnrollIDList []map[string]any `json:"enroll_id_list"`
}

// SetMeetingEnrollConfig 修改会议报名配置
// 修改会议的报名配置和报名问题，且需要会议已开启报名。
//
// see https://developer.work.weixin.qq.com/document/path/98797
func (c *Client) SetMeetingEnrollConfig(r *SetMeetingEnrollConfigRequest, opts ...any) (out SetMeetingEnrollConfigResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/enroll/set_config",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetMeetingEnrollConfigRequest is request of Client.SetMeetingEnrollConfig
type SetMeetingEnrollConfigRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// ApproveType 审批类型：1：自动审批，默认自动审批。2：手动审批。
	ApproveType int `json:"approve_type"`
	// IsCollectQuestion 是否收集问题：1：不收集，默认不收集问题。2：收集。
	IsCollectQuestion int `json:"is_collect_question"`
	// NoRegistrationNeededForStaff 本企业成员无需报名。true：本企业成员无需报名。false：默认配置，本企业成员及企业外成员需要报名。
	NoRegistrationNeededForStaff bool `json:"no_registration_needed_for_staff"`
	// QuestionList 报名问题列表，非特殊问题按传入的顺序排序，特殊问题会优先放在最前面，仅开启收集问题时有效。详见Question。
	QuestionList []map[string]any `json:"question_list"`
}

// SetMeetingEnrollConfigResponse is response of Client.SetMeetingEnrollConfig
type SetMeetingEnrollConfigResponse struct {
	// QuestionCount 报名问题数量，不收集问题时该字段返回0。
	QuestionCount int `json:"question_count"`
}

// ListAttendee 获取已参会成员列表
// 支持查询预约会议及网络研讨会已参会成员列表。如果会议还未开始，调用此接口查询会返回空列表。
//
// see https://developer.work.weixin.qq.com/document/path/99295
func (c *Client) ListAttendee(r *ListAttendeeRequest, opts ...any) (out ListAttendeeResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/get_attendee_list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListAttendeeRequest is request of Client.ListAttendee
type ListAttendeeRequest struct {
	// Meetingid 会议id
	Meetingid string `json:"meetingid" validate:"required"`
	// SubMeetingid 周期性会议子会议 ID。如果是周期性会议，此参数必传。
	SubMeetingid string `json:"sub_meetingid"`
	// StartTime 参会时间过滤起始时间（单位秒）。时间区间不允许超过31天。
	StartTime int `json:"start_time"`
	// EndTime 参会时间过滤终止时间（单位秒）。时间区间不允许超过31天。
	EndTime int `json:"end_time"`
	// Cursor 分页查询用，将上一个请求返回的next_cursor字段传入。第一次查询时可不传值。
	Cursor string `json:"cursor"`
	// Limit 拉取参会成员条数，目前每页支持最大100条。必须与首次调用获得cursor时传入的limit一致。
	Limit int `json:"limit"`
}

// ListAttendeeResponse is response of Client.ListAttendee
type ListAttendeeResponse struct {
	// HasMore 是否还有待拉取的成员列表
	HasMore bool `json:"has_more"`
	// NextCursor 分页用，下一次拉取列表将该字段填入cursor字段
	NextCursor string `json:"next_cursor"`
	// Attendees 参会人列表
	Attendees []map[string]any `json:"attendees"`
}

// GetCustomerShortUrl 获取用户专属参会链接
// 可以获取指定会议的所有专属参会链接及 customer_data。该接口不支持个人会议号会议、网络研讨会（Webinar）。
//
// see https://developer.work.weixin.qq.com/document/path/98819
func (c *Client) GetCustomerShortUrl(r *GetCustomerShortUrlRequest, opts ...any) (out GetCustomerShortUrlResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/get_customer_short_url",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCustomerShortUrlRequest is request of Client.GetCustomerShortUrl
type GetCustomerShortUrlRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
}

// GetCustomerShortUrlResponse is response of Client.GetCustomerShortUrl
type GetCustomerShortUrlResponse struct {
	// MeetingShortURLCustomerDataList 用户专属参会链接对象列表
	MeetingShortURLCustomerDataList []map[string]any `json:"meeting_short_url_customer_data_list"`
}

// ListMeetingGuests 获取会议嘉宾列表
// 通过会议ID获取会议嘉宾列表
//
// see https://developer.work.weixin.qq.com/document/path/99077
func (c *Client) ListMeetingGuests(r *ListMeetingGuestsRequest, opts ...any) (out ListMeetingGuestsResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/get_guests",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMeetingGuestsRequest is request of Client.ListMeetingGuests
type ListMeetingGuestsRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
}

// ListMeetingGuestsResponse is response of Client.ListMeetingGuests
type ListMeetingGuestsResponse struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid"`
	// MeetingCode 入会码
	MeetingCode string `json:"meeting_code"`
	// Title 会议主题
	Title string `json:"title"`
	// Guests 嘉宾列表
	Guests []map[string]any `json:"guests"`
}

// GetMeetingInfo 获取会议详情
// 该接口用于获取指定会议的详情内容。
//
// see https://developer.work.weixin.qq.com/document/path/99015
func (c *Client) GetMeetingInfo(r *GetMeetingInfoRequest, opts ...any) (out GetMeetingInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/get_info",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMeetingInfoRequest is request of Client.GetMeetingInfo
type GetMeetingInfoRequest struct {
	// Meetingid 会议id。meetingid和meeting_code必须填一个
	Meetingid string `json:"meetingid"`
	// MeetingCode 入会码。meetingid和meeting_code必须填一个
	MeetingCode string `json:"meeting_code"`
	// SubMeetingid 周期性会议子会议id
	SubMeetingid string `json:"sub_meetingid"`
}

// GetMeetingInfoResponse is response of Client.GetMeetingInfo
type GetMeetingInfoResponse struct {
	// AdminUserID 会议管理员的userId
	AdminUserID string `json:"admin_userid"`
	// Title 会议的标题，最大40个字节
	Title string `json:"title"`
	// MeetingStart 会议开始时间的unix时间戳
	MeetingStart int `json:"meeting_start"`
	// MeetingDuration 会议时长
	MeetingDuration int `json:"meeting_duration"`
	// Description 会议的描述，最大600字节
	Description string `json:"description"`
	// Location 会议地点，最多128个字符
	Location string `json:"location"`
	// MainDepartment 发起人所在部门
	MainDepartment int `json:"main_department"`
	// Status 会议的状态。1：待开始 2：会议中 3：已结束 4：已取消 5：已过期
	Status int `json:"status"`
	// MeetingType 会议类型。0：一次性会议 1：周期性会议 2：微信专属会议 3：Rooms 投屏会议 5：个人会议号会议 6：网络研讨会
	MeetingType int `json:"meeting_type"`
	// Attendees 参会成员
	Attendees any `json:"attendees"`
	// Guests 会议嘉宾列表
	Guests []map[string]any `json:"guests"`
	// Settings 会议配置
	Settings any `json:"settings"`
	// Reminders 重复会议相关配置
	Reminders any `json:"reminders"`
	// MeetingCode 会议号
	MeetingCode string `json:"meeting_code"`
	// MeetingLink 入会链接
	MeetingLink string `json:"meeting_link"`
	// HasVote 是否有投票（会议创建人和主持人才有权限查询）
	HasVote bool `json:"has_vote"`
	// SubMeetings 周期性子会议列表
	SubMeetings []map[string]any `json:"sub_meetings"`
	// HasMoreSubMeeting 是否还有更多子会议特例。0：无更多 1：有更多子会议特例
	HasMoreSubMeeting int `json:"has_more_sub_meeting"`
	// RemainSubMeetings 剩余子会议场数
	RemainSubMeetings int `json:"remain_sub_meetings"`
	// CurrentSubMeetingid 当前子会议ID（进行中/即将开始）
	CurrentSubMeetingid string `json:"current_sub_meetingid"`
	// SubRepeatList 周期性会议分段信息
	SubRepeatList []map[string]any `json:"sub_repeat_list"`
}

// ListMeetingInvitees 获取会议受邀成员列表
// 根据会议 ID 获取预约会议时添加到参会人名单的成员（会议受邀成员）列表，支持分页获取
//
// see https://developer.work.weixin.qq.com/document/path/98160
func (c *Client) ListMeetingInvitees(r *ListMeetingInviteesRequest, opts ...any) (out ListMeetingInviteesResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/get_invitees",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMeetingInviteesRequest is request of Client.ListMeetingInvitees
type ListMeetingInviteesRequest struct {
	// Meetingid 会议 id
	Meetingid string `json:"meetingid" validate:"required"`
	// Cursor 分页查询用，将上一个请求返回的 next_cursor 字段传入。第一次查询时可不传值
	Cursor string `json:"cursor"`
}

// ListMeetingInviteesResponse is response of Client.ListMeetingInvitees
type ListMeetingInviteesResponse struct {
	// HasMore 是否还有尚未拉取的成员列表
	HasMore bool `json:"has_more"`
	// NextCursor 当 has_more 为 true 时，下次查询时输入参数 cursor 的值
	NextCursor string `json:"next_cursor"`
	// Invitees 受邀成员列表
	Invitees []map[string]any `json:"invitees"`
}

// GetMeetingQuality 获取会议健康度
// 获取已结束会议的会议及参会成员的健康度。
//
// see https://developer.work.weixin.qq.com/document/path/99053
func (c *Client) GetMeetingQuality(r *GetMeetingQualityRequest, opts ...any) (out GetMeetingQualityResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/get_quality",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMeetingQualityRequest is request of Client.GetMeetingQuality
type GetMeetingQualityRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// SubMeetingid 周期性子会议ID。如果是周期性会议，此参数必传。
	SubMeetingid string `json:"sub_meetingid"`
	// StartTime 参会时间过滤起始时间，UNIX 时间戳（单位秒），可查询的时间区间为过去7天到现在。
	StartTime int `json:"start_time" validate:"required"`
	// Cursor 分页查询用，将上一个请求返回的next_cursor字段传入。第一次查询时可不传值
	Cursor string `json:"cursor"`
	// Limit 分页大小，默认50，最大50
	Limit int `json:"limit"`
}

// GetMeetingQualityResponse is response of Client.GetMeetingQuality
type GetMeetingQualityResponse struct {
	// Quality 健康度：0：无数据，1：健康，2：告警
	Quality int `json:"quality"`
	// AudioQuality 音频质量：0：无数据，1：好，2：较好，3：中，4：差
	AudioQuality int `json:"audio_quality"`
	// VideoQuality 视频质量：0：无数据，1：好，2：较好，3：中，4：差
	VideoQuality int `json:"video_quality"`
	// ScreenShareQuality 共享屏幕质量：0：无数据，1：好，2：较好，3：中，4：差
	ScreenShareQuality int `json:"screen_share_quality"`
	// NetworkQuality 网络质量：0：无数据，1：好，2：较好，3：中，4：差
	NetworkQuality int `json:"network_quality"`
	// Problems 告警的具体问题列表
	Problems []string `json:"problems"`
	// Attendees 参会人员健康度列表（按成员入会时间正序排列）
	Attendees []map[string]any `json:"attendees"`
	// NextCursor 分页用，下一次拉取列表将该字段填入cursor字段
	NextCursor string `json:"next_cursor"`
	// HasMore 是否还有待拉取的列表
	HasMore bool `json:"has_more"`
}

// ListRealtimeAttendee 获取实时会中成员列表
// 获取当前会中成员列表，仅包括会中的成员。如果成员已离会，则不返回。
//
// see https://developer.work.weixin.qq.com/document/path/99012
func (c *Client) ListRealtimeAttendee(r *ListRealtimeAttendeeRequest, opts ...any) (out ListRealtimeAttendeeResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/get_realtime_attendee_list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListRealtimeAttendeeRequest is request of Client.ListRealtimeAttendee
type ListRealtimeAttendeeRequest struct {
	// Meetingid 会议id
	Meetingid string `json:"meetingid" validate:"required"`
	// SubMeetingid 周期性会议子会议 ID。如果是周期性会议，此参数必传。
	SubMeetingid string `json:"sub_meetingid"`
	// Cursor 分页查询用，将上一个请求返回的next_cursor字段传入。第一次查询时可不传值
	Cursor string `json:"cursor"`
	// Limit 单次查询返回的数量，最大50条。limit参数必须与首次调用获得cursor时传入的limit一致。
	Limit int `json:"limit"`
}

// ListRealtimeAttendeeResponse is response of Client.ListRealtimeAttendee
type ListRealtimeAttendeeResponse struct {
	// HasMore 是否还有更多的成员列表
	HasMore bool `json:"has_more"`
	// NextCursor 分页用，下一次拉取列表将该字段填入cursor字段
	NextCursor string `json:"next_cursor"`
	// Attendees 参会人列表
	Attendees []map[string]any `json:"attendees"`
}

// GetUserMeetingId 获取成员会议ID列表
// 该接口用于获取指定成员指定时间内的会议ID列表。
//
// see https://developer.work.weixin.qq.com/document/path/99050
func (c *Client) GetUserMeetingId(r *GetUserMeetingIdRequest, opts ...any) (out GetUserMeetingIdResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/get_user_meetingid",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetUserMeetingIdRequest is request of Client.GetUserMeetingId
type GetUserMeetingIdRequest struct {
	// UserID 企业成员的userid
	UserID string `json:"userid" validate:"required"`
	// Cursor 上一次调用时返回的cursor，初次调用可以填"0"
	Cursor string `json:"cursor"`
	// BeginTime 开始时间
	BeginTime int `json:"begin_time"`
	// EndTime 结束时间，时间跨度不超过180天。如果begin_time和end_time都没填的话，默认end_time为当前时间。
	EndTime int `json:"end_time"`
	// Limit 每次拉取的数据量，默认值和最大值都为100
	Limit int `json:"limit"`
}

// GetUserMeetingIdResponse is response of Client.GetUserMeetingId
type GetUserMeetingIdResponse struct {
	// NextCursor 当前数据最后一个key值，如果下次调用带上该值则从该key值往后拉，用于实现分页拉取，next_cursor为未返回或者为空字符串表示数据已取完
	NextCursor string `json:"next_cursor"`
	// MeetingidList 会议ID列表，可能为空
	MeetingidList []string `json:"meetingid_list"`
}

// AddMeetingLayout 添加会议基础布局
// 对API成功预定的会议添加会议基础布局，支持多个布局的添加，每个布局支持多页模板，默认选中第一页模板作为该布局的首页进行展示。
//
// see https://developer.work.weixin.qq.com/document/path/99300
func (c *Client) AddMeetingLayout(r *AddMeetingLayoutRequest, opts ...any) (out AddMeetingLayoutResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/layout/add",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddMeetingLayoutRequest is request of Client.AddMeetingLayout
type AddMeetingLayoutRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// LayoutList 布局对象列表，详见LayoutRequest
	LayoutList []map[string]any `json:"layout_list" validate:"required"`
	// DefaultLayoutOrder 布局列表中会议需要应用的布局序号，从1开始计数
	DefaultLayoutOrder int `json:"default_layout_order"`
}

// AddMeetingLayoutResponse is response of Client.AddMeetingLayout
type AddMeetingLayoutResponse struct {
	// SelectedLayoutID 会议应用的布局 ID
	SelectedLayoutID string `json:"selected_layout_id"`
	// LayoutList 布局对象列表，详见LayoutResponse
	LayoutList []map[string]any `json:"layout_list"`
}

// AddMeetingBackground 添加会议背景
// 对成功预定的会议添加会议背景，支持多个背景图片的添加。一场会议最多添加7个背景，且仅支持不超过10MB大小的 PNG 格式图片，分辨率最小为1920x1080。
//
// see https://developer.work.weixin.qq.com/document/path/98851
func (c *Client) AddMeetingBackground(r *AddMeetingBackgroundRequest, opts ...any) (out AddMeetingBackgroundResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/layout/add_background",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddMeetingBackgroundRequest is request of Client.AddMeetingBackground
type AddMeetingBackgroundRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// ImageList 图片对象列表，详见Image
	ImageList []map[string]any `json:"image_list" validate:"required"`
	// DefaultImageOrder 图片列表中会议需要使用的背景图片序号，从1开始计数。不填默认为1
	DefaultImageOrder int `json:"default_image_order"`
}

// AddMeetingBackgroundResponse is response of Client.AddMeetingBackground
type AddMeetingBackgroundResponse struct {
	// SelectedBackgroundID 会议应用的背景 ID
	SelectedBackgroundID string `json:"selected_background_id"`
	// BackgroundList 背景对象列表，详见Background
	BackgroundList []map[string]any `json:"background_list"`
}

// BatchDeleteMeetingBackground 批量删除会议背景
// 根据背景 ID 删除多个会议背景。正在被会议应用的背景无法删除，请先设置成其他背景或恢复成会议的默认黑色背景后再行删除。
//
// see https://developer.work.weixin.qq.com/document/path/98854
func (c *Client) BatchDeleteMeetingBackground(r *BatchDeleteMeetingBackgroundRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/layout/batch_delete_background",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchDeleteMeetingBackgroundRequest is request of Client.BatchDeleteMeetingBackground
type BatchDeleteMeetingBackgroundRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// BackgroundIDList 背景 ID 列表
	BackgroundIDList []string `json:"background_id_list" validate:"required"`
}

// DeleteMeetingBackground 删除会议背景
// 根据背景 ID 删除单个会议背景。正在被会议应用的背景无法删除，请先设置成其他背景或恢复成会议的默认黑色背景后再行删除。
//
// see https://developer.work.weixin.qq.com/document/path/98853
func (c *Client) DeleteMeetingBackground(r *DeleteMeetingBackgroundRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/layout/delete_background",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteMeetingBackgroundRequest is request of Client.DeleteMeetingBackground
type DeleteMeetingBackgroundRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// BackgroundID 背景 ID
	BackgroundID string `json:"background_id" validate:"required"`
}

// ListMeetingBackgrounds 获取会议背景列表
// 根据会议 ID 返回会议背景列表信息
//
// see https://developer.work.weixin.qq.com/document/path/99224
func (c *Client) ListMeetingBackgrounds(r *ListMeetingBackgroundsRequest, opts ...any) (out ListMeetingBackgroundsResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/layout/list_background",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMeetingBackgroundsRequest is request of Client.ListMeetingBackgrounds
type ListMeetingBackgroundsRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
}

// ListMeetingBackgroundsResponse is response of Client.ListMeetingBackgrounds
type ListMeetingBackgroundsResponse struct {
	// SelectedBackgroundID 会议应用的背景 ID
	SelectedBackgroundID string `json:"selected_background_id"`
	// BackgroundList 背景对象列表
	BackgroundList []map[string]any `json:"background_list"`
}

// ListMeetingLayoutTemplate 获取布局模板列表
// 获取企业下所有的布局模板列表。
//
// see https://developer.work.weixin.qq.com/document/path/99299
func (c *Client) ListMeetingLayoutTemplate(r *ListMeetingLayoutTemplateRequest, opts ...any) (out ListMeetingLayoutTemplateResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/meeting/layout/list_template",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMeetingLayoutTemplateRequest is request of Client.ListMeetingLayoutTemplate
type ListMeetingLayoutTemplateRequest struct {}

// ListMeetingLayoutTemplateResponse is response of Client.ListMeetingLayoutTemplate
type ListMeetingLayoutTemplateResponse struct {
	// LayoutTemplateList 布局模板对象列表
	LayoutTemplateList []map[string]any `json:"layout_template_list"`
}

// SetMeetingDefaultLayout 设置会议默认布局
// 对API成功预定的会议设置默认布局。
//
// see https://developer.work.weixin.qq.com/document/path/98847
func (c *Client) SetMeetingDefaultLayout(r *SetMeetingDefaultLayoutRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/layout/set_default",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetMeetingDefaultLayoutRequest is request of Client.SetMeetingDefaultLayout
type SetMeetingDefaultLayoutRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// SelectedLayoutID 会议应用的布局 ID（若送空""，表示恢复成会议自带的默认原始布局）
	SelectedLayoutID string `json:"selected_layout_id" validate:"required"`
}

// SetMeetingDefaultBackground 设置会议默认背景
// 对API成功预定的会议设置默认背景。
//
// see https://developer.work.weixin.qq.com/document/path/98852
func (c *Client) SetMeetingDefaultBackground(r *SetMeetingDefaultBackgroundRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/layout/set_default_background",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetMeetingDefaultBackgroundRequest is request of Client.SetMeetingDefaultBackground
type SetMeetingDefaultBackgroundRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// SelectedBackgroundID 会议应用的背景 ID（若送空""，则表示恢复成会议默认的黑色背景）
	SelectedBackgroundID string `json:"selected_background_id" validate:"required"`
}

// UpdateMeetingLayout 修改会议基础布局
// 根据布局 ID 对设置好的会议基础布局进行修改。用户座次设置区分会前和会中两种方式：会前只允许设置邀请者成员，会中只允许设置参会成员。
//
// see https://developer.work.weixin.qq.com/document/path/99301
func (c *Client) UpdateMeetingLayout(r *UpdateMeetingLayoutRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/layout/update",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateMeetingLayoutRequest is request of Client.UpdateMeetingLayout
type UpdateMeetingLayoutRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// LayoutID 布局 ID
	LayoutID string `json:"layout_id" validate:"required"`
	// PageList 布局对象列表，详见LayoutPage
	PageList []map[string]any `json:"page_list" validate:"required"`
	// EnableSetDefault 是否设置为会议应用的布局，默认不设置
	EnableSetDefault bool `json:"enable_set_default"`
}

// HangupMraCall 挂断 MRA 呼叫
// 会议中对 MRA 的呼叫进行挂断操作。
//
// see https://developer.work.weixin.qq.com/document/path/99036
func (c *Client) HangupMraCall(r *HangupMraCallRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/mra/hangup",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// HangupMraCallRequest is request of Client.HangupMraCall
type HangupMraCallRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// Mra 被操作 mra 设备信息
	Mra any `json:"mra" validate:"required"`
}

// QueryMraStatus 获取 MRA 状态信息
// 获取指定 MRA 设备的当前状态信息，包括名称、静音状态、视频状态、举手状态、默认布局等。
//
// see https://developer.work.weixin.qq.com/document/path/99033
func (c *Client) QueryMraStatus(r *QueryMraStatusRequest, opts ...any) (out QueryMraStatusResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/mra/query_status",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// QueryMraStatusRequest is request of Client.QueryMraStatus
type QueryMraStatusRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// TmpOpenID 被查询成员临时身份 ID
	TmpOpenID string `json:"tmp_openid" validate:"required"`
}

// QueryMraStatusResponse is response of Client.QueryMraStatus
type QueryMraStatusResponse struct {
	// TmpOpenID 被查询成员当场会议的成员临时 ID
	TmpOpenID string `json:"tmp_openid"`
	// InstanceID 被查询成员的终端设备类型：9：voip、sip 设备（即MRA设备）
	InstanceID int `json:"instance_id"`
	// UserRole 成员角色：0：普通成员角色，1：创建者角色，2：主持人等
	UserRole int `json:"user_role"`
	// WebinarMemberRole 网络研讨会成员角色：0：普通参会角色，1：内部嘉宾等
	WebinarMemberRole int `json:"webinar_member_role"`
	// IP 成员的 IP 地址。当成员在会中时才能返回
	IP string `json:"ip"`
	// Name 成员当前显示名称
	Name string `json:"name"`
	// AudioState 麦克风状态：true：开启，false：关闭
	AudioState bool `json:"audio_state"`
	// VideoState 摄像头状态：true：开启，false：关闭
	VideoState bool `json:"video_state"`
	// ScreenSharedState 屏幕共享状态：true：开启，false：关闭
	ScreenSharedState bool `json:"screen_shared_state"`
	// DefaultLayout 当前成员的默认分屏设置：1：等分模式，2：全屏模式，3：1+N
	DefaultLayout int `json:"default_layout"`
	// RaiseHandsState 举手状态：true：举手中，false：手放下
	RaiseHandsState bool `json:"raise_hands_state"`
}

// SetMRAHand 设置 MRA 举手或手放下
// API创建的会议中对MRA进行举手和手放下操作
//
// see https://developer.work.weixin.qq.com/document/path/98788
func (c *Client) SetMRAHand(r *SetMRAHandRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/mra/set_raise_hand",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetMRAHandRequest is request of Client.SetMRAHand
type SetMRAHandRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// RaiseHand MRA设备举手操作：true为举手，false为手放下
	RaiseHand bool `json:"raise_hand" validate:"required"`
	// Mra 被操作mra设备的会中临时ID信息
	Mra any `json:"mra" validate:"required"`
}

// BatchPhoneCallout 批量外呼
// 创建批量电话入会呼叫，支持会议未开始或会中外呼，每次最多50路。
//
// see https://developer.work.weixin.qq.com/document/path/98823
func (c *Client) BatchPhoneCallout(r *BatchPhoneCalloutRequest, opts ...any) (out BatchPhoneCalloutResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/phone/callout",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchPhoneCalloutRequest is request of Client.BatchPhoneCallout
type BatchPhoneCalloutRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// PhoneNumbers 外呼的电话号码对象数组
	PhoneNumbers []map[string]any `json:"phone_numbers" validate:"required"`
}

// BatchPhoneCalloutResponse is response of Client.BatchPhoneCallout
type BatchPhoneCalloutResponse struct {
	// PhoneNumbers 成功外呼的电话号码对象数组
	PhoneNumbers []map[string]any `json:"phone_numbers"`
	// InvalidPhoneNumbers 不合法的外呼电话号码对象数组
	InvalidPhoneNumbers []map[string]any `json:"invalid_phone_numbers"`
}

// GetMeetingCalloutStatus 获取会议的外呼状态
// 此接口用于获取外呼电话入会的数据。无论是从客户端、API 还是 Web 端发起的外呼，在该接口都可以获取到，不针对呼叫人做数据的隔离。需要分页，分页最大条数100。Webinar 暂不支持外呼。
//
// see https://developer.work.weixin.qq.com/document/path/99096
func (c *Client) GetMeetingCalloutStatus(r *GetMeetingCalloutStatusRequest, opts ...any) (out GetMeetingCalloutStatusResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/phone/get_callout_status",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMeetingCalloutStatusRequest is request of Client.GetMeetingCalloutStatus
type GetMeetingCalloutStatusRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// Cursor 分页查询用，将上一个请求返回的next_cursor字段传入。第一次查询时可不传值。
	Cursor string `json:"cursor"`
	// Limit 分页大小，默认20，最大100。
	Limit int `json:"limit"`
}

// GetMeetingCalloutStatusResponse is response of Client.GetMeetingCalloutStatus
type GetMeetingCalloutStatusResponse struct {
	// PhoneNumbers 电话号码对象数组
	PhoneNumbers []map[string]any `json:"phone_numbers"`
	// NextCursor 分页用，下一次拉取列表将该字段填入cursor字段
	NextCursor string `json:"next_cursor"`
	// HasMore 是否还有待拉取的列表
	HasMore bool `json:"has_more"`
}

// GetMeetingPhoneTmpOpenId 获取电话入会的成员ID
// 此接口用于获取电话入会的成员ID，以tmp_openid返回。支持查询会议下PSTN的历史tmp_openid列表。
//
// see https://developer.work.weixin.qq.com/document/path/99097
func (c *Client) GetMeetingPhoneTmpOpenId(r *GetMeetingPhoneTmpOpenIdRequest, opts ...any) (out GetMeetingPhoneTmpOpenIdResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/phone/get_tmp_openid",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMeetingPhoneTmpOpenIdRequest is request of Client.GetMeetingPhoneTmpOpenId
type GetMeetingPhoneTmpOpenIdRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// PhoneNumbers 外呼的电话号码对象数组，上限20个
	PhoneNumbers []map[string]any `json:"phone_numbers" validate:"required"`
}

// GetMeetingPhoneTmpOpenIdResponse is response of Client.GetMeetingPhoneTmpOpenId
type GetMeetingPhoneTmpOpenIdResponse struct {
	// TmpOpenIDList 返回的tmp_openid对象数组
	TmpOpenIDList []map[string]any `json:"tmp_openid_list"`
}

// CreateMeetingPollTheme 创建会议投票主题
// 为指定的会议创建投票。该接口支持多问题投票。仅进行中的会议可以调用该接口，操作者是主持人或者会议管理员。
//
// see https://developer.work.weixin.qq.com/document/path/98834
func (c *Client) CreateMeetingPollTheme(r *CreateMeetingPollThemeRequest, opts ...any) (out CreateMeetingPollThemeResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/poll/create_theme",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateMeetingPollThemeRequest is request of Client.CreateMeetingPollTheme
type CreateMeetingPollThemeRequest struct {
	// OperatorUserID 操作者的openid
	OperatorUserID string `json:"operator_userid" validate:"required"`
	// InstanceID 操作者入会所用的设备id
	InstanceID int `json:"instance_id" validate:"required"`
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// PollTopic 投票主题，最多50个字符
	PollTopic string `json:"poll_topic" validate:"required"`
	// PollDesc 投票主题描述，最多100个字符
	PollDesc string `json:"poll_desc" validate:"required"`
	// IsAnony 是否匿名。0：实名，默认值；1：匿名
	IsAnony int `json:"is_anony"`
	// PollQuestions 投票问题数组，每个投票支持添加10个问题
	PollQuestions []map[string]any `json:"poll_questions" validate:"required"`
}

// CreateMeetingPollThemeResponse is response of Client.CreateMeetingPollTheme
type CreateMeetingPollThemeResponse struct {
	// PollThemeID 投票主题 ID
	PollThemeID string `json:"poll_theme_id"`
}

// DeleteMeetingPoll 删除会议投票
// 删除会议投票，仅进行中的会议可以调用该接口，操作者是主持人或者会议管理员
//
// see https://developer.work.weixin.qq.com/document/path/98839
func (c *Client) DeleteMeetingPoll(r *DeleteMeetingPollRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/poll/delete",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteMeetingPollRequest is request of Client.DeleteMeetingPoll
type DeleteMeetingPollRequest struct {
	// OperatorUserID 操作者的openid
	OperatorUserID string `json:"operator_userid" validate:"required"`
	// InstanceID 操作者入会的设备id
	InstanceID int `json:"instance_id" validate:"required"`
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// PollThemeID 投票主题 ID，传入则代表删除投票主题
	PollThemeID string `json:"poll_theme_id"`
	// PollID 投票 ID，传入则代表删除投票实例
	PollID string `json:"poll_id"`
}

// FinishMeetingPoll 结束会议投票
// 结束指定会议的投票。仅进行中的会议可以调用该接口，操作者是主持人或者会议管理员
//
// see https://developer.work.weixin.qq.com/document/path/98841
func (c *Client) FinishMeetingPoll(r *FinishMeetingPollRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/poll/finish",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// FinishMeetingPollRequest is request of Client.FinishMeetingPoll
type FinishMeetingPollRequest struct {
	// OperatorUserID 操作者openid
	OperatorUserID string `json:"operator_userid" validate:"required"`
	// InstanceID 操作者入会设备id
	InstanceID int `json:"instance_id" validate:"required"`
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// PollThemeID 投票主题ID
	PollThemeID string `json:"poll_theme_id" validate:"required"`
	// PollID 投票 ID
	PollID string `json:"poll_id" validate:"required"`
}

// GetPollDetail 获取会议投票详情
// 获取投票的详情，包括问题、选项、投票结果。
//
// see https://developer.work.weixin.qq.com/document/path/99218
func (c *Client) GetPollDetail(r *GetPollDetailRequest, opts ...any) (out GetPollDetailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/poll/get_poll_detail",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetPollDetailRequest is request of Client.GetPollDetail
type GetPollDetailRequest struct {
	// OperatorUserID 操作者openid
	OperatorUserID string `json:"operator_userid" validate:"required"`
	// InstanceID 操作者入会设备对应的id
	InstanceID int `json:"instance_id" validate:"required"`
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// PollID 投票ID
	PollID string `json:"poll_id" validate:"required"`
}

// GetPollDetailResponse is response of Client.GetPollDetail
type GetPollDetailResponse struct {
	// PollThemeID 投票主题id
	PollThemeID string `json:"poll_theme_id"`
	// PollTopic 投票主题
	PollTopic string `json:"poll_topic"`
	// PollDesc 投票描述
	PollDesc string `json:"poll_desc"`
	// IsAnony 是否匿名。0：实名，1：匿名
	IsAnony int `json:"is_anony"`
	// Status 投票状态。1：投票中，2：已结束
	Status int `json:"status"`
	// IsShared 是否共享
	IsShared int `json:"is_shared"`
	// VoteTotalNum 投票人数
	VoteTotalNum int `json:"vote_total_num"`
	// PollQuestionData 投票结果数组
	PollQuestionData []map[string]any `json:"poll_question_data"`
}

// ListMeetingPoll 获取会议投票列表
// 获取指定会议的投票列表。
//
// see https://developer.work.weixin.qq.com/document/path/99216
func (c *Client) ListMeetingPoll(r *ListMeetingPollRequest, opts ...any) (out ListMeetingPollResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/poll/get_poll_list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMeetingPollRequest is request of Client.ListMeetingPoll
type ListMeetingPollRequest struct {
	// OperatorUserID 操作者openid
	OperatorUserID string `json:"operator_userid" validate:"required"`
	// InstanceID 操作者入会设备对应的id
	InstanceID int `json:"instance_id" validate:"required"`
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
}

// ListMeetingPollResponse is response of Client.ListMeetingPoll
type ListMeetingPollResponse struct {
	// PollsThemeInfo 投票主题信息列表
	PollsThemeInfo []map[string]any `json:"polls_theme_info"`
}

// GetMeetingPollThemeInfo 获取会议投票主题信息
// 获取投票主题的详细信息，包括问题、选项，但不包括投票结果。
//
// see https://developer.work.weixin.qq.com/document/path/99217
func (c *Client) GetMeetingPollThemeInfo(r *GetMeetingPollThemeInfoRequest, opts ...any) (out GetMeetingPollThemeInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/poll/get_theme_info",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMeetingPollThemeInfoRequest is request of Client.GetMeetingPollThemeInfo
type GetMeetingPollThemeInfoRequest struct {
	// OperatorUserID 操作者openid
	OperatorUserID string `json:"operator_userid" validate:"required"`
	// InstanceID 操作者入会设备对应的id
	InstanceID int `json:"instance_id" validate:"required"`
	// Meetingid 会议ID
	Meetingid string `json:"meetingid"`
	// PollThemeID 投票主题id
	PollThemeID string `json:"poll_theme_id" validate:"required"`
}

// GetMeetingPollThemeInfoResponse is response of Client.GetMeetingPollThemeInfo
type GetMeetingPollThemeInfoResponse struct {
	// PollTopic 投票主题
	PollTopic string `json:"poll_topic"`
	// PollDesc 投票描述
	PollDesc string `json:"poll_desc"`
	// IsAnony 是否匿名。0：实名，1：匿名
	IsAnony int `json:"is_anony"`
	// PollQuestionData 投票问题数组
	PollQuestionData []map[string]any `json:"poll_question_data"`
}

// StartMeetingPoll 发起会议投票
// 使用已有的投票主题发起投票。仅进行中的会议可以调用该接口，操作者是主持人或者会议管理员。
//
// see https://developer.work.weixin.qq.com/document/path/98840
func (c *Client) StartMeetingPoll(r *StartMeetingPollRequest, opts ...any) (out StartMeetingPollResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/poll/start",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// StartMeetingPollRequest is request of Client.StartMeetingPoll
type StartMeetingPollRequest struct {
	// OperatorUserID 操作者openid
	OperatorUserID string `json:"operator_userid" validate:"required"`
	// InstanceID 操作者入会的设备id
	InstanceID int `json:"instance_id" validate:"required"`
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// PollThemeID 投票主题 ID
	PollThemeID string `json:"poll_theme_id" validate:"required"`
}

// StartMeetingPollResponse is response of Client.StartMeetingPoll
type StartMeetingPollResponse struct {
	// PollID 投票ID
	PollID string `json:"poll_id"`
}

// UpdateMeetingPollTheme 修改会议投票主题
// 修改会议中的投票主题信息，仅支持全覆盖修改。仅进行中的会议可调用，操作者需为主持人或会议管理员。
//
// see https://developer.work.weixin.qq.com/document/path/98835
func (c *Client) UpdateMeetingPollTheme(r *UpdateMeetingPollThemeRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/poll/update_theme",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateMeetingPollThemeRequest is request of Client.UpdateMeetingPollTheme
type UpdateMeetingPollThemeRequest struct {
	// OperatorUserID 操作者openid
	OperatorUserID string `json:"operator_userid" validate:"required"`
	// InstanceID 操作者入会设备对应的id
	InstanceID int `json:"instance_id" validate:"required"`
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// PollThemeID 投票主题id
	PollThemeID string `json:"poll_theme_id" validate:"required"`
	// PollTopic 投票主题，最多50个字符
	PollTopic string `json:"poll_topic" validate:"required"`
	// PollDesc 投票主题描述，最多100个字符
	PollDesc string `json:"poll_desc" validate:"required"`
	// IsAnony 是否匿名。0：实名，默认值；1：匿名
	IsAnony int `json:"is_anony"`
	// PollQuestions 投票问题数组，每个投票支持添加10个问题
	PollQuestions []map[string]any `json:"poll_questions" validate:"required"`
}

// CloseScreenShare 关闭成员屏幕共享
// 此接口用于停止成员发起的屏幕共享。目前暂不支持 MRA 设备作为被操作者的情况
//
// see https://developer.work.weixin.qq.com/document/path/99094
func (c *Client) CloseScreenShare(r *CloseScreenShareRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/realcontrol/close_screen_share",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CloseScreenShareRequest is request of Client.CloseScreenShare
type CloseScreenShareRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// OperatedUser 被操作对象，详见User
	OperatedUser any `json:"operated_user" validate:"required"`
}

// DismissMeeting 结束会议
// 结束一个进行中的会议。
//
// see https://developer.work.weixin.qq.com/document/path/98187
func (c *Client) DismissMeeting(r *DismissMeetingRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/realcontrol/dismiss",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DismissMeetingRequest is request of Client.DismissMeeting
type DismissMeetingRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// ForceDismiss 是否强制结束会议，默认值为1：0：不强制结束会议，会议中有参会者，则无法强制结束会议；1：强制结束会议，会议中有参会者，也会强制结束会议
	ForceDismiss int `json:"force_dismiss"`
	// RetrieveCode 是否回收会议号，默认值为0：0：不回收会议号，可以重新入会；1：回收会议号，不可重新入会。说明：周期性会议如果还有子会议，需设置为不回收会议号，否则会导致后续子会议无法正常进行。此字段对快速会议...
	RetrieveCode int `json:"retrieve_code"`
}

// KickoutMeetingUsers 移出成员
// 将会议中成员移出会议
//
// see https://developer.work.weixin.qq.com/document/path/99051
func (c *Client) KickoutMeetingUsers(r *KickoutMeetingUsersRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/realcontrol/kickout_users",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// KickoutMeetingUsersRequest is request of Client.KickoutMeetingUsers
type KickoutMeetingUsersRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// AllowRejoin 是否允许再次入会：true:允许再次入会;false：不允许
	AllowRejoin bool `json:"allow_rejoin" validate:"required"`
	// OperatedUsers 被操作对象列表
	OperatedUsers []map[string]any `json:"operated_users" validate:"required"`
}

// ManageWaitingRoomUsers 管理等候室成员
// 会议等候室设置，支持主持人将等候室成员移入会议、将会议成员移入等候室、将等候室成员移出等候室等操作。
//
// see https://developer.work.weixin.qq.com/document/path/99018
func (c *Client) ManageWaitingRoomUsers(r *ManageWaitingRoomUsersRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/realcontrol/manage_waiting_room_users",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ManageWaitingRoomUsersRequest is request of Client.ManageWaitingRoomUsers
type ManageWaitingRoomUsersRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// OperateType 操作类型。1：主持人将等候室成员移入会议；2：主持人将会中成员移入等候室；3：主持人将等候室成员移出等候室
	OperateType int `json:"operate_type" validate:"required"`
	// AllowRejoin 移出成员后是否允许其再次加入会议（该字段对 MRA 设备不生效）。operate_type=3时才允许设置
	AllowRejoin bool `json:"allow_rejoin"`
	// OperatedUsers 被操作成员列表
	OperatedUsers []map[string]any `json:"operated_users" validate:"required"`
}

// MuteMeetingUser 静音成员
// 会议中成员静音操作。
//
// see https://developer.work.weixin.qq.com/document/path/99025
func (c *Client) MuteMeetingUser(r *MuteMeetingUserRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/realcontrol/mute_user",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// MuteMeetingUserRequest is request of Client.MuteMeetingUser
type MuteMeetingUserRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// Option true：静音；false：解除静音
	Option bool `json:"option" validate:"required"`
	// OperatedUser 被操作对象列表
	OperatedUser []map[string]any `json:"operated_user" validate:"required"`
}

// SetMeetingConfig 管理会中设置
// 管理进行中会议的设置项，例如：全体静音、是否允许参会者聊天设置、锁定会议、隐藏会议号和密码、是否开启等候室等。目前暂不支持 MRA 设备作为被操作者的情况
//
// see https://developer.work.weixin.qq.com/document/path/99016
func (c *Client) SetMeetingConfig(r *SetMeetingConfigRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/realcontrol/set",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetMeetingConfigRequest is request of Client.SetMeetingConfig
type SetMeetingConfigRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// MuteAll 是否全体静音：true：全体静音；false：关闭全体静音
	MuteAll bool `json:"mute_all"`
	// AllowUnmuteSelf 是否允许成员自己取消静音，请求参数 mute_all 必传，且 mute_all = true 时设置才生效。true：允许；false：不允许
	AllowUnmuteSelf bool `json:"allow_unmute_self"`
	// EnableEnterMute 成员入会静音：0：关闭静音；1：开启静音；2：超过6人自动开启静音
	EnableEnterMute int `json:"enable_enter_mute"`
	// MeetingLocked 是否锁定会议：true：锁定；false：关闭锁定
	MeetingLocked bool `json:"meeting_locked"`
	// HideMeetingCodePassword 隐藏会议号和密码：true：隐藏；false：不隐藏
	HideMeetingCodePassword bool `json:"hide_meeting_code_password"`
	// AllowChat 允许参会者聊天设置：0：允许参会者自由聊天；1：仅允许参会者公开聊天；2：仅允许私聊主持人
	AllowChat int `json:"allow_chat"`
	// AllowShareScreen 是否允许参会者发起屏幕共享：true：允许；false：不允许
	AllowShareScreen bool `json:"allow_share_screen"`
	// AllowExternalUser 是否仅企业成员可入会：true：仅企业成员可入会；false：不限制
	AllowExternalUser bool `json:"allow_external_user"`
	// PlayIvrOnJoin 成员入会是否播放提示音：true：成员入会播放提示音；false：不播放
	PlayIvrOnJoin bool `json:"play_ivr_on_join"`
	// EnableWaitingRoom 是否开启等候室：true：开启；false：关闭
	EnableWaitingRoom bool `json:"enable_waiting_room"`
}

// SetMeetingCohost 管理联席主持人
// 设置或撤销会议中的参会者联席主持人身份。目前暂不支持 MRA 设备作为被操作者的情况
//
// see https://developer.work.weixin.qq.com/document/path/99017
func (c *Client) SetMeetingCohost(r *SetMeetingCohostRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/realcontrol/set_cohost",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetMeetingCohostRequest is request of Client.SetMeetingCohost
type SetMeetingCohostRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// Action 具体设置动作。true：设置联席主持人，false：撤销联席主持人
	Action bool `json:"action" validate:"required"`
	// OperatedUser 被操作成员信息
	OperatedUser any `json:"operated_user" validate:"required"`
}

// SetMeetingNicknames 修改成员在会中显示的昵称
// 此接口用于修改成员在会中显示的昵称。
//
// see https://developer.work.weixin.qq.com/document/path/99095
func (c *Client) SetMeetingNicknames(r *SetMeetingNicknamesRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/realcontrol/set_nicknames",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetMeetingNicknamesRequest is request of Client.SetMeetingNicknames
type SetMeetingNicknamesRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// OpereatedUsers 被操作对象列表
	OpereatedUsers []map[string]any `json:"opereated_users" validate:"required"`
}

// SwitchUserVideo 关闭或开启成员视频
// 关闭指定成员视频，支持关闭或开启 MRA 设备的视频。
//
// see https://developer.work.weixin.qq.com/document/path/99026
func (c *Client) SwitchUserVideo(r *SwitchUserVideoRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/realcontrol/switch_user_video",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SwitchUserVideoRequest is request of Client.SwitchUserVideo
type SwitchUserVideoRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// Video 操作类型。false：关闭视频（默认值）。true：开启视频，仅支持MRA设备。
	Video bool `json:"video"`
	// OperatedUser 被操作者
	OperatedUser any `json:"operated_user" validate:"required"`
}

// DeleteMeetingRecord 删除会议录制
// 删除会议的所有录制文件，该接口会删除会议录制 ID 里对应的所有云录制文件。
//
// see https://developer.work.weixin.qq.com/document/path/98206
func (c *Client) DeleteMeetingRecord(r *DeleteMeetingRecordRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/record/delete",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteMeetingRecordRequest is request of Client.DeleteMeetingRecord
type DeleteMeetingRecordRequest struct {
	// MeetingRecordID 会议录制ID
	MeetingRecordID string `json:"meeting_record_id" validate:"required"`
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
}

// DeleteMeetingRecordFile 删除单个录制文件
// 从会议中删除指定的某个录制文件
//
// see https://developer.work.weixin.qq.com/document/path/98207
func (c *Client) DeleteMeetingRecordFile(r *DeleteMeetingRecordFileRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/record/delete_file",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteMeetingRecordFileRequest is request of Client.DeleteMeetingRecordFile
type DeleteMeetingRecordFileRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// RecordFileID 录制文件ID
	RecordFileID string `json:"record_file_id" validate:"required"`
}

// GetRecordFileDetail 获取单个录制文件详情
// 获取单个云录制的详情信息，包括录制文件和会议纪要，并可获取播放地址和下载地址
//
// see https://developer.work.weixin.qq.com/document/path/100916
func (c *Client) GetRecordFileDetail(r *GetRecordFileDetailRequest, opts ...any) (out GetRecordFileDetailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/record/get_file",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetRecordFileDetailRequest is request of Client.GetRecordFileDetail
type GetRecordFileDetailRequest struct {
	// RecordFileID 会议录制ID
	RecordFileID string `json:"record_file_id" validate:"required"`
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
}

// GetRecordFileDetailResponse is response of Client.GetRecordFileDetail
type GetRecordFileDetailResponse struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid"`
	// MeetingCode 入会码
	MeetingCode string `json:"meeting_code"`
	// RecordFileID 录制文件 ID
	RecordFileID string `json:"record_file_id"`
	// ViewAddress 播放地址
	ViewAddress string `json:"view_address"`
	// DownloadAddress 下载地址，默认6小时过期
	DownloadAddress string `json:"download_address"`
	// DownloadAddressFileType 下载视频文件格式，例如：mp4
	DownloadAddressFileType string `json:"download_address_file_type"`
	// AudioAddress 音频下载地址，默认6小时过期
	AudioAddress string `json:"audio_address"`
	// AudioAddressFileType 下载音频文件格式，例如：m4a
	AudioAddressFileType string `json:"audio_address_file_type"`
	// RecordName 录制文件名
	RecordName string `json:"record_name"`
	// StartTime 录制开始时间
	StartTime int `json:"start_time"`
	// EndTime 录制结束时间
	EndTime int `json:"end_time"`
	// MeetingRecordName 会议录制名
	MeetingRecordName string `json:"meeting_record_name"`
	// MeetingSummary 会议纪要文件列表
	MeetingSummary []map[string]any `json:"meeting_summary"`
	// AiMeetingTranscripts 录制转写文件（智能优化版）列表
	AiMeetingTranscripts []map[string]any `json:"ai_meeting_transcripts"`
}

// GetMeetingRecordFileList 获取会议录制地址
// 获取会议云录制的播放地址和下载地址
//
// see https://developer.work.weixin.qq.com/document/path/100917
func (c *Client) GetMeetingRecordFileList(r *GetMeetingRecordFileListRequest, opts ...any) (out GetMeetingRecordFileListResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/record/get_file_list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMeetingRecordFileListRequest is request of Client.GetMeetingRecordFileList
type GetMeetingRecordFileListRequest struct {
	// MeetingRecordID 会议录制ID
	MeetingRecordID string `json:"meeting_record_id" validate:"required"`
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
}

// GetMeetingRecordFileListResponse is response of Client.GetMeetingRecordFileList
type GetMeetingRecordFileListResponse struct {
	// MeetingRecordID 会议录制ID
	MeetingRecordID string `json:"meeting_record_id"`
	// Meetingid 会议ID
	Meetingid string `json:"meetingid"`
	// MeetingCode 会议code
	MeetingCode string `json:"meeting_code"`
	// Title 会议主题
	Title string `json:"title"`
	// RecordFiles 直播录制文件列表
	RecordFiles []map[string]any `json:"record_files"`
}

// ListMeetingRecordStatistics 获取录制文件访问统计
// 获取会议录制 ID 对应的访问数据，按照天维度返回
//
// see https://developer.work.weixin.qq.com/document/path/99263
func (c *Client) ListMeetingRecordStatistics(r *ListMeetingRecordStatisticsRequest, opts ...any) (out ListMeetingRecordStatisticsResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/record/get_statistics",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMeetingRecordStatisticsRequest is request of Client.ListMeetingRecordStatistics
type ListMeetingRecordStatisticsRequest struct {
	// MeetingRecordID 会议录制ID
	MeetingRecordID string `json:"meeting_record_id" validate:"required"`
	// StartTime 查询起始时间戳，UNIX 时间戳（单位秒）。默认展示最近31天的数据，时间区间不允许超过31天
	StartTime int `json:"start_time"`
	// EndTime 查询结束时间戳，UNIX 时间戳（单位秒）。默认展示最近31天的数据，时间区间不允许超过31天
	EndTime int `json:"end_time"`
}

// ListMeetingRecordStatisticsResponse is response of Client.ListMeetingRecordStatistics
type ListMeetingRecordStatisticsResponse struct {
	// Summaries 统计结果列表，详见Summary
	Summaries []map[string]any `json:"summaries"`
}

// ListMeetingRecord 获取会议录制列表
// 获取云录制记录，根据成员 ID、会议 ID、会议 code 进行查询，支持根据时间区间分页获取。
//
// see https://developer.work.weixin.qq.com/document/path/99236
func (c *Client) ListMeetingRecord(r *ListMeetingRecordRequest, opts ...any) (out ListMeetingRecordResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/record/list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMeetingRecordRequest is request of Client.ListMeetingRecord
type ListMeetingRecordRequest struct {
	// Meetingid 会议 ID，不为空时优先根据会议 ID 查询
	Meetingid string `json:"meetingid"`
	// MeetingCode 会议 code，当 meeting_id 为空且 meeting_code 不为空时根据会议 code 查询。当 meeting_id 和 meeting_code 均为空时，表示查询成员所有会...
	MeetingCode string `json:"meeting_code"`
	// UserID 待查询成员的 userid。meetingid、meeting_code 和 userid 三者选填其中一项
	UserID string `json:"userid"`
	// StartTime 查询起始时间戳，UNIX 时间戳（单位秒）。查询时间区间跨度不允许超过 31 天
	StartTime int `json:"start_time" validate:"required"`
	// EndTime 查询结束时间戳，UNIX 时间戳（单位秒）。查询时间区间跨度不允许超过 31 天
	EndTime int `json:"end_time" validate:"required"`
	// Cursor 上一次调用时返回的 next_cursor，初次调用可不传
	Cursor string `json:"cursor"`
	// Limit 单次查询返回的最大数量限制，默认为 10，最大 20。limit 参数必须与首次调用获得 cursor 时传入的 limit 一致
	Limit int `json:"limit"`
}

// ListMeetingRecordResponse is response of Client.ListMeetingRecord
type ListMeetingRecordResponse struct {
	// HasMore 是否还有未拉取的会议录制列表
	HasMore bool `json:"has_more"`
	// NextCursor 当 has_more=true 时有值，表示当前数据最后一个 key 值，如果下次调用带上该值则从该 key 值往后拉，用于实现分页拉取
	NextCursor string `json:"next_cursor"`
	// RecordList 会议录制列表，详见 Record
	RecordList []map[string]any `json:"record_list"`
}

// GetTranscriptDetail 获取录制转写详情
// 获取云录制转写的详情，包含时间戳、文本等内容。
//
// see https://developer.work.weixin.qq.com/document/path/100926
func (c *Client) GetTranscriptDetail(r *GetTranscriptDetailRequest, opts ...any) (out GetTranscriptDetailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/record/transcript/get_detail",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetTranscriptDetailRequest is request of Client.GetTranscriptDetail
type GetTranscriptDetailRequest struct {
	// RecordFileID 录制文件 ID
	RecordFileID string `json:"record_file_id" validate:"required"`
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// Pid 查询的起始段落 ID。获取 pid 后（含）的段落，默认从0开始。示例："1111"
	Pid string `json:"pid"`
	// Limit 查询的段落数，默认查询全量数据
	Limit int `json:"limit"`
}

// GetTranscriptDetailResponse is response of Client.GetTranscriptDetail
type GetTranscriptDetailResponse struct {
	// HasMore 如果传入 pid 或 limit 字段，且没有一次完全返回，则为 true；否则为 false
	HasMore bool `json:"has_more"`
	// Transcripts 录制转写详情
	Transcripts any `json:"transcripts"`
}

// GetTranscriptParagraphList 获取录制转写段落信息
// 获取云录制转写的段落信息（段落总数、段落 ID）
//
// see https://developer.work.weixin.qq.com/document/path/100925
func (c *Client) GetTranscriptParagraphList(r *GetTranscriptParagraphListRequest, opts ...any) (out GetTranscriptParagraphListResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/record/transcript/get_paragraph_list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetTranscriptParagraphListRequest is request of Client.GetTranscriptParagraphList
type GetTranscriptParagraphListRequest struct {
	// RecordFileID 录制文件 ID
	RecordFileID string `json:"record_file_id" validate:"required"`
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
}

// GetTranscriptParagraphListResponse is response of Client.GetTranscriptParagraphList
type GetTranscriptParagraphListResponse struct {
	// AudioDetect 声纹识别状态：0-未完成，1-已完成
	AudioDetect int `json:"audio_detect"`
	// Paragraphs 段落列表
	Paragraphs []map[string]any `json:"paragraphs"`
}

// SearchTranscript 获取录制转写搜索结果
// 根据指定内容搜索录制转写
//
// see https://developer.work.weixin.qq.com/document/path/100927
func (c *Client) SearchTranscript(r *SearchTranscriptRequest, opts ...any) (out SearchTranscriptResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/record/transcript/search",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SearchTranscriptRequest is request of Client.SearchTranscript
type SearchTranscriptRequest struct {
	// RecordFileID 会议录制文件ID
	RecordFileID string `json:"record_file_id" validate:"required"`
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// Text 搜索的文本
	Text string `json:"text" validate:"required"`
}

// SearchTranscriptResponse is response of Client.SearchTranscript
type SearchTranscriptResponse struct {
	// Hits 搜索结果列表
	Hits []map[string]any `json:"hits"`
	// Timelines 搜索结果时间戳对象列表
	Timelines []map[string]any `json:"timelines"`
}

// UpdateMeetingRecordSharingConfig 修改会议录制共享设置
// 根据会议录制 ID 修改共享等配置，支持修改共享权限、共享密码、共享有效期等信息
//
// see https://developer.work.weixin.qq.com/document/path/98208
func (c *Client) UpdateMeetingRecordSharingConfig(r *UpdateMeetingRecordSharingConfigRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/record/update_sharing_config",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateMeetingRecordSharingConfigRequest is request of Client.UpdateMeetingRecordSharingConfig
type UpdateMeetingRecordSharingConfigRequest struct {
	// MeetingRecordID 会议录制ID
	MeetingRecordID string `json:"meeting_record_id" validate:"required"`
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// SharingConfig 共享配置，详见SharingConfig
	SharingConfig any `json:"sharing_config"`
}

// BookMeetingRooms 预定Rooms会议室
// 对成功预定的会议添加Rooms会议室，支持为同一个会议预定多个Rooms会议室。
//
// see https://developer.work.weixin.qq.com/document/path/99273
func (c *Client) BookMeetingRooms(r *BookMeetingRoomsRequest, opts ...any) (out BookMeetingRoomsResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/rooms/book",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BookMeetingRoomsRequest is request of Client.BookMeetingRooms
type BookMeetingRoomsRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// MeetingRoomIDList Rooms会议室 ID 列表
	MeetingRoomIDList []string `json:"meeting_room_id_list" validate:"required"`
	// SubjectVisible 在会议开始前的一小时内，是否在 Room 上显示会议主题，默认值为 true
	SubjectVisible bool `json:"subject_visible"`
}

// BookMeetingRoomsResponse is response of Client.BookMeetingRooms
type BookMeetingRoomsResponse struct {
	// MeetingRoomList Rooms会议室对象列表
	MeetingRoomList []map[string]any `json:"meeting_room_list"`
}

// CallMeetingRoom 呼叫Rooms会议室
// 通过Rooms会议室ID呼叫Rooms会议室邀请其入会
//
// see https://developer.work.weixin.qq.com/document/path/99276
func (c *Client) CallMeetingRoom(r *CallMeetingRoomRequest, opts ...any) (out CallMeetingRoomResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/rooms/call",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CallMeetingRoomRequest is request of Client.CallMeetingRoom
type CallMeetingRoomRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// MeetingRoomID Rooms会议室ID，与mra_address二选一
	MeetingRoomID string `json:"meeting_room_id"`
	// MraAddress MRA对象
	MraAddress any `json:"mra_address"`
}

// CallMeetingRoomResponse is response of Client.CallMeetingRoom
type CallMeetingRoomResponse struct {
	// InviteID 呼叫ID
	InviteID string `json:"invite_id"`
}

// CancelRoomCall 取消呼叫Rooms会议室
// 会议可以通过Rooms会议室 ID 进行取消呼叫操作。
//
// see https://developer.work.weixin.qq.com/document/path/99275
func (c *Client) CancelRoomCall(r *CancelRoomCallRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/rooms/cancel_call",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CancelRoomCallRequest is request of Client.CancelRoomCall
type CancelRoomCallRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// InviteID 呼叫 ID
	InviteID string `json:"invite_id" validate:"required"`
	// MeetingRoomID Rooms会议室 ID，与 mra_address 二选一
	MeetingRoomID string `json:"meeting_room_id"`
	// MraAddress MRA 对象
	MraAddress any `json:"mra_address"`
}

// GetMeetingRoomConfig 获取Rooms会议室配置项
// 获取Rooms会议室的配置项
//
// see https://developer.work.weixin.qq.com/document/path/99277
func (c *Client) GetMeetingRoomConfig(r *GetMeetingRoomConfigRequest, opts ...any) (out GetMeetingRoomConfigResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/rooms/get_config",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMeetingRoomConfigRequest is request of Client.GetMeetingRoomConfig
type GetMeetingRoomConfigRequest struct {
	// MeetingRoomID Rooms会议室 ID
	MeetingRoomID string `json:"meeting_room_id" validate:"required"`
}

// GetMeetingRoomConfigResponse is response of Client.GetMeetingRoomConfig
type GetMeetingRoomConfigResponse struct {
	// MeetingSettings Rooms会议室会议配置对象，详见MeetingSettings
	MeetingSettings any `json:"meeting_settings"`
	// RecordSettings Rooms会议室录制配置对象，详见RecordSettings
	RecordSettings any `json:"record_settings"`
}

// GetRoomInfo 获取Rooms会议室详情
// 根据Rooms会议室 ID 获取该Rooms会议室详细信息。
//
// see https://developer.work.weixin.qq.com/document/path/99279
func (c *Client) GetRoomInfo(r *GetRoomInfoRequest, opts ...any) (out GetRoomInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/rooms/get_info",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetRoomInfoRequest is request of Client.GetRoomInfo
type GetRoomInfoRequest struct {
	// MeetingRoomID Rooms会议室 ID
	MeetingRoomID string `json:"meeting_room_id" validate:"required"`
}

// GetRoomInfoResponse is response of Client.GetRoomInfo
type GetRoomInfoResponse struct {
	// BasicInfo Rooms会议室基本信息
	BasicInfo any `json:"basic_info"`
	// AccountInfo Rooms会议室账号信息
	AccountInfo any `json:"account_info"`
	// HardwareInfo Rooms会议室硬件信息
	HardwareInfo any `json:"hardware_info"`
	// PmiInfo Rooms会议室PMI信息
	PmiInfo any `json:"pmi_info"`
	// MonitorStatus 告警通知状态：0-未开启，1-已开启
	MonitorStatus int `json:"monitor_status"`
	// IsAllowCall 是否允许被呼叫
	IsAllowCall bool `json:"is_allow_call"`
	// ScheduledStatus 预定状态：0-未开放预定，1-开放预定
	ScheduledStatus int `json:"scheduled_status"`
}

// GetMeetingRoomInventory 获取Rooms会议室资源
// 获取企业购买的Rooms会议室资源
//
// see https://developer.work.weixin.qq.com/document/path/99298
func (c *Client) GetMeetingRoomInventory(r *GetMeetingRoomInventoryRequest, opts ...any) (out GetMeetingRoomInventoryResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/rooms/get_inventory",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMeetingRoomInventoryRequest is request of Client.GetMeetingRoomInventory
type GetMeetingRoomInventoryRequest struct {}

// GetMeetingRoomInventoryResponse is response of Client.GetMeetingRoomInventory
type GetMeetingRoomInventoryResponse struct {
	// NormalCount 普通设备数
	NormalCount int `json:"normal_count"`
	// SpecialCount 专款设备数
	SpecialCount int `json:"special_count"`
	// NormalUsedCount 普通设备使用数
	NormalUsedCount int `json:"normal_used_count"`
	// SpecialUsedCount 专款设备使用数
	SpecialUsedCount int `json:"special_used_count"`
	// NormalExpiredCount 普通设备过期数
	NormalExpiredCount int `json:"normal_expired_count"`
	// SpecialExpiredCount 专款设备过期数
	SpecialExpiredCount int `json:"special_expired_count"`
}

// GetRoomResponseStatus 获取Rooms会议室应答状态
// 会议获取其呼叫Rooms会议室的应答状态
//
// see https://developer.work.weixin.qq.com/document/path/99274
func (c *Client) GetRoomResponseStatus(r *GetRoomResponseStatusRequest, opts ...any) (out GetRoomResponseStatusResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/rooms/get_response_status",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetRoomResponseStatusRequest is request of Client.GetRoomResponseStatus
type GetRoomResponseStatusRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// MeetingRoomID Rooms会议室 ID，与 mra_address 二选一
	MeetingRoomID string `json:"meeting_room_id"`
	// MraAddress MRA 对象
	MraAddress any `json:"mra_address"`
}

// GetRoomResponseStatusResponse is response of Client.GetRoomResponseStatus
type GetRoomResponseStatusResponse struct {
	// Status 应答状态：0-无应答，1-未呼叫，2-入会中，3-被拒绝，4-呼叫中，5-取消呼叫，6-已离会
	Status int `json:"status"`
	// ResponseTime 最近一次应答时间
	ResponseTime string `json:"response_time"`
}

// ListMeetingRoom 获取Rooms会议室列表
// 获取企业下的Rooms会议室列表。
//
// see https://developer.work.weixin.qq.com/document/path/99280
func (c *Client) ListMeetingRoom(r *ListMeetingRoomRequest, opts ...any) (out ListMeetingRoomResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/rooms/list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMeetingRoomRequest is request of Client.ListMeetingRoom
type ListMeetingRoomRequest struct {
	// MeetingRoomName Rooms会议室名称（支持模糊匹配查找）
	MeetingRoomName string `json:"meeting_room_name"`
	// Cursor 分页查询用，将上一个请求返回的next_cursor字段传入。第一次查询时可不传值
	Cursor string `json:"cursor"`
	// Limit 分页大小，从1开始，最大50，默认20
	Limit int `json:"limit"`
}

// ListMeetingRoomResponse is response of Client.ListMeetingRoom
type ListMeetingRoomResponse struct {
	// HasMore 是否还有更多Rooms会议室列表
	HasMore bool `json:"has_more"`
	// NextCursor 分页用，has_more为true时，下一次拉取列表将该字段填入cursor字段
	NextCursor string `json:"next_cursor"`
	// MeetingRoomList Rooms会议室对象列表
	MeetingRoomList []map[string]any `json:"meeting_room_list"`
}

// ListMeetingRoomControllers 获取控制器列表
// 获取企业下的控制器列表。
//
// see https://developer.work.weixin.qq.com/document/path/99231
func (c *Client) ListMeetingRoomControllers(r *ListMeetingRoomControllersRequest, opts ...any) (out ListMeetingRoomControllersResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/rooms/list_controllers",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMeetingRoomControllersRequest is request of Client.ListMeetingRoomControllers
type ListMeetingRoomControllersRequest struct {
	// ControllerName 需要获取的控制器名称（支持模糊匹配查找）
	ControllerName string `json:"controller_name"`
	// Cursor 分页查询用，将上一个请求返回的next_cursor字段传入。第一次查询时可不传值
	Cursor string `json:"cursor"`
	// Limit 分页大小，从1开始，最大50，默认20
	Limit int `json:"limit"`
}

// ListMeetingRoomControllersResponse is response of Client.ListMeetingRoomControllers
type ListMeetingRoomControllersResponse struct {
	// HasMore 是否还有更多Rooms会议室列表
	HasMore bool `json:"has_more"`
	// NextCursor 分页用，has_more为true时，下一次拉取列表将该字段填入cursor字段
	NextCursor string `json:"next_cursor"`
	// ControllerInfoList Rooms会议室对象列表
	ControllerInfoList []map[string]any `json:"controller_info_list"`
}

// ListMeetingRoomDevices 获取设备列表
// 获取企业下的可用设备列表。
//
// see https://developer.work.weixin.qq.com/document/path/99230
func (c *Client) ListMeetingRoomDevices(r *ListMeetingRoomDevicesRequest, opts ...any) (out ListMeetingRoomDevicesResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/rooms/list_devices",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMeetingRoomDevicesRequest is request of Client.ListMeetingRoomDevices
type ListMeetingRoomDevicesRequest struct {
	// MeetingRoomName Rooms会议室名称（支持模糊匹配查找）
	MeetingRoomName string `json:"meeting_room_name"`
	// Cursor 分页查询用，将上一个请求返回的next_cursor字段传入。第一次查询时可不传值
	Cursor string `json:"cursor"`
	// Limit 分页大小，从1开始，最大50，默认20
	Limit int `json:"limit"`
}

// ListMeetingRoomDevicesResponse is response of Client.ListMeetingRoomDevices
type ListMeetingRoomDevicesResponse struct {
	// HasMore 是否还有更多Rooms会议室列表
	HasMore bool `json:"has_more"`
	// NextCursor 分页用，has_more为true时，下一次拉取列表将该字段填入cursor字段
	NextCursor string `json:"next_cursor"`
	// DeviceInfoList 设备信息对象列表
	DeviceInfoList []map[string]any `json:"device_info_list"`
}

// ListRoomMeetings 获取Rooms会议室下的会议列表
// 获取指定Rooms会议室下的会议列表
//
// see https://developer.work.weixin.qq.com/document/path/99278
func (c *Client) ListRoomMeetings(r *ListRoomMeetingsRequest, opts ...any) (out ListRoomMeetingsResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/rooms/list_meetings",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListRoomMeetingsRequest is request of Client.ListRoomMeetings
type ListRoomMeetingsRequest struct {
	// MeetingRoomID Rooms会议室 ID，与rooms_id二者填其一
	MeetingRoomID string `json:"meeting_room_id"`
	// RoomsID rooms 设备 rooms_id。与meeting_room_id二者填其一
	RoomsID string `json:"rooms_id"`
	// StartTime Unix 时间戳。查询起始时间，时间区间不超过90天
	StartTime int `json:"start_time"`
	// EndTime Unix 时间戳。查询结束时间，时间区间不超过90天
	EndTime int `json:"end_time"`
	// Cursor 分页查询用，将上一个请求返回的next_cursor字段传入。第一次查询时可不传值
	Cursor string `json:"cursor"`
	// Limit 分页大小，默认20条，最大20条
	Limit int `json:"limit"`
}

// ListRoomMeetingsResponse is response of Client.ListRoomMeetings
type ListRoomMeetingsResponse struct {
	// HasMore 是否还有更多Rooms会议室列表
	HasMore bool `json:"has_more"`
	// NextCursor 分页用，has_more为true时，下一次拉取列表将该字段填入cursor字段
	NextCursor string `json:"next_cursor"`
	// MeetingInfoList 会议对象列表
	MeetingInfoList []map[string]any `json:"meeting_info_list"`
}

// ReleaseMeetingRooms 释放Rooms会议室
// 通过会议 ID 释放 Rooms 会议室，支持释放多个 Rooms 会议室。
//
// see https://developer.work.weixin.qq.com/document/path/99281
func (c *Client) ReleaseMeetingRooms(r *ReleaseMeetingRoomsRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/rooms/release",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ReleaseMeetingRoomsRequest is request of Client.ReleaseMeetingRooms
type ReleaseMeetingRoomsRequest struct {
	// Meetingid 会议 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// MeetingRoomIDList Rooms 会议室 ID 列表
	MeetingRoomIDList []string `json:"meeting_room_id_list" validate:"required"`
}

// SetMeetingGuests 更新会议嘉宾列表
// 通过会议ID更新会议嘉宾列表
//
// see https://developer.work.weixin.qq.com/document/path/99042
func (c *Client) SetMeetingGuests(r *SetMeetingGuestsRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/set_guests",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetMeetingGuestsRequest is request of Client.SetMeetingGuests
type SetMeetingGuestsRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// Guests 会议嘉宾列表
	Guests []map[string]any `json:"guests" validate:"required"`
}

// SetMeetingInvitees 更新会议受邀成员列表
// 根据会议 ID 更新会议受邀成员列表，最多可以传入2000名受邀成员。
//
// see https://developer.work.weixin.qq.com/document/path/98997
func (c *Client) SetMeetingInvitees(r *SetMeetingInviteesRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/set_invitees",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetMeetingInviteesRequest is request of Client.SetMeetingInvitees
type SetMeetingInviteesRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// Invitees 受邀成员列表。最大支持2000人。
	Invitees []map[string]any `json:"invitees"`
}

// ListMeetingStart 获取会议发起记录
// 该接口用于获取员工发起的会议记录及其状态，包含员工主动发起快速会议及作为首位参与者进入预约会议的场景。
//
// see https://developer.work.weixin.qq.com/document/path/99651
func (c *Client) ListMeetingStart(r *ListMeetingStartRequest, opts ...any) (out ListMeetingStartResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/statistics/get_start_list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMeetingStartRequest is request of Client.ListMeetingStart
type ListMeetingStartRequest struct {
	// Type 查询类型。1:发起成功的会议记录; 2:发起失败的会议
	Type int `json:"type" validate:"required"`
	// BeginTime 查询范围起始时间戳，单位为秒
	BeginTime int `json:"begin_time" validate:"required"`
	// EndTime 查询范围结束时间戳，单位为秒
	EndTime int `json:"end_time" validate:"required"`
	// Limit 每次拉取的数据量，默认值200，最大值1000
	Limit int `json:"limit"`
	// Cursor 用于分页查询的游标，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"`
}

// ListMeetingStartResponse is response of Client.ListMeetingStart
type ListMeetingStartResponse struct {
	// NextCursor 当前数据最后一个key值，如果下次调用带上该值则从该key值往后拉，用于实现分页拉取
	NextCursor string `json:"next_cursor"`
	// HasMore 是否还有数据待拉取
	HasMore bool `json:"has_more"`
	// MeetingList 发起会议成功或失败的记录信息，取决于请求中的type参数
	MeetingList []map[string]any `json:"meeting_list"`
	// UserID 会议发起者的userid。如果是企业外部成员作为首位参与者进入预约会议，则该字段为此预约会议创建者的userid
	UserID string `json:"userid"`
	// StartTime 会议发起时间
	StartTime int `json:"start_time"`
}

// UpdateMeeting 修改预约会议
// 该接口用于修改一个指定的预约会议。
//
// see https://developer.work.weixin.qq.com/document/path/99047
func (c *Client) UpdateMeeting(r *UpdateMeetingRequest, opts ...any) (out UpdateMeetingResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/update",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateMeetingRequest is request of Client.UpdateMeeting
type UpdateMeetingRequest struct {
	// Meetingid 会议id，仅允许修改预约状态下的会议
	Meetingid string `json:"meetingid" validate:"required"`
	// Title 会议的标题，最多支持40个字节或者20个utf8字符
	Title string `json:"title"`
	// MeetingStart 会议开始时间的unix时间戳。需大于当前时间。注：修改该字段时必须同时指定meeting_duration。
	MeetingStart int `json:"meeting_start"`
	// MeetingDuration 会议持续时间（单位秒），最小300秒，最大86399秒。注：修改该字段时，必须同时指定meeting_start。
	MeetingDuration int `json:"meeting_duration"`
	// Description 会议的描述，最多支持500个字节或者utf8字符
	Description string `json:"description"`
	// Location 会议地点，最多128个字符
	Location string `json:"location"`
	// RemindTime 指定会议开始前多久提醒成员，相对于meeting_start前的秒数，默认为0
	RemindTime int `json:"remind_time"`
	// Invitees 邀请参会的成员。包含userid列表
	Invitees any `json:"invitees"`
	// CalID 会议所属日历ID。不多于64字节
	CalID string `json:"cal_id"`
	// Settings 会议配置，详见Settings说明
	Settings any `json:"settings"`
	// Reminders 重复会议相关配置，详见Reminders说明
	Reminders any `json:"reminders"`
}

// UpdateMeetingResponse is response of Client.UpdateMeeting
type UpdateMeetingResponse struct {
	// ExcessUsers 参会人中包含无效会议账号的userid，仅在购买会议专业版企业由于部分参会人无有效会议账号时返回
	ExcessUsers []string `json:"excess_users"`
}

// ListMeetingVip 获取高级功能账号列表
// 查询企业已分配高级功能且在应用可见范围的账号列表
//
// see https://developer.work.weixin.qq.com/document/path/99510
func (c *Client) ListMeetingVip(r *ListMeetingVipRequest, opts ...any) (out ListMeetingVipResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/vip/list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMeetingVipRequest is request of Client.ListMeetingVip
type ListMeetingVipRequest struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"`
	// Limit 用于分页查询，每次请求返回的数据上限。默认100，最大200
	Limit int `json:"limit"`
}

// ListMeetingVipResponse is response of Client.ListMeetingVip
type ListMeetingVipResponse struct {
	// HasMore 是否还有更多数据未获取
	HasMore bool `json:"has_more"`
	// NextCursor 下一次请求的cursor值
	NextCursor string `json:"next_cursor"`
	// UserIDList 符合条件的企业成员userid列表
	UserIDList []string `json:"userid_list"`
}

// BatchAddMeetingVip 分配高级功能账号
// 该接口用于分配应用可见范围企业成员的高级功能。
//
// see https://developer.work.weixin.qq.com/document/path/99508
func (c *Client) BatchAddMeetingVip(r *BatchAddMeetingVipRequest, opts ...any) (out BatchAddMeetingVipResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/vip/submit_batch_add_job",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchAddMeetingVipRequest is request of Client.BatchAddMeetingVip
type BatchAddMeetingVipRequest struct {
	// UserIDList 要分配高级功能的企业成员userid列表，单次操作最大限制100个
	UserIDList []string `json:"userid_list" validate:"required"`
}

// BatchAddMeetingVipResponse is response of Client.BatchAddMeetingVip
type BatchAddMeetingVipResponse struct {
	// JobID 批量分配高级功能的任务id
	JobID string `json:"jobid"`
	// InvalidUserIDList 非法的userid 列表，不在应用可见范围的userid以及无法识别的userid
	InvalidUserIDList []string `json:"invalid_userid_list"`
}

// BatchDeleteMeetingVipJob 取消高级功能账号
// 该接口用于撤销分配应用可见范围企业成员的高级功能。
//
// see https://developer.work.weixin.qq.com/document/path/99509
func (c *Client) BatchDeleteMeetingVipJob(r *BatchDeleteMeetingVipJobRequest, opts ...any) (out BatchDeleteMeetingVipJobResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/vip/submit_batch_del_job",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchDeleteMeetingVipJobRequest is request of Client.BatchDeleteMeetingVipJob
type BatchDeleteMeetingVipJobRequest struct {
	// UserIDList 要撤销分配高级功能的企业成员userid列表，单次操作最多限制100个
	UserIDList []string `json:"userid_list" validate:"required"`
}

// BatchDeleteMeetingVipJobResponse is response of Client.BatchDeleteMeetingVipJob
type BatchDeleteMeetingVipJobResponse struct {
	// JobID 批量取消高级功能的任务id
	JobID string `json:"jobid"`
	// InvalidUserIDList 非法的userid 列表，不在应用可见范围的useri以及无法识别的userid
	InvalidUserIDList []string `json:"invalid_userid_list"`
}

// ListMeetingWaitingRoomUser 获取实时等候室成员列表
// 获取某指定会议的等候室成员列表，需开启等候室且为“会议进行中”状态。如果会议非进行中，调用此接口查询会返回空列表。
//
// see https://developer.work.weixin.qq.com/document/path/98163
func (c *Client) ListMeetingWaitingRoomUser(r *ListMeetingWaitingRoomUserRequest, opts ...any) (out ListMeetingWaitingRoomUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/waitingroom/get_current_user_list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMeetingWaitingRoomUserRequest is request of Client.ListMeetingWaitingRoomUser
type ListMeetingWaitingRoomUserRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// Limit 分页大小，默认10，最大50
	Limit int `json:"limit"`
	// Cursor 分页查询用，将上一个请求返回的next_cursor字段传入。第一次查询时可不传值
	Cursor string `json:"cursor"`
}

// ListMeetingWaitingRoomUserResponse is response of Client.ListMeetingWaitingRoomUser
type ListMeetingWaitingRoomUserResponse struct {
	// HasMore 是否还有未拉取的成员列表
	HasMore bool `json:"has_more"`
	// NextCursor 当has_more=true时有值，下次请求将该字段赋值给cursor字段
	NextCursor string `json:"next_cursor"`
	// UserList 等候室人员对象数组
	UserList []map[string]any `json:"user_list"`
}

// ListWaitingRoomUser 获取等候室成员记录
// 获取等候室成员列表，包括等候室内所有进出过的成员。会前、会中、会后都可以获取。
//
// see https://developer.work.weixin.qq.com/document/path/99065
func (c *Client) ListWaitingRoomUser(r *ListWaitingRoomUserRequest, opts ...any) (out ListWaitingRoomUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/waitingroom/get_user_list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListWaitingRoomUserRequest is request of Client.ListWaitingRoomUser
type ListWaitingRoomUserRequest struct {
	// Meetingid 会议ID
	Meetingid string `json:"meetingid" validate:"required"`
	// Limit 分页大小，默认20，最大50
	Limit int `json:"limit"`
	// Cursor 分页查询用，将上一个请求返回的next_cursor字段传入。第一次查询时可不传值
	Cursor string `json:"cursor"`
}

// ListWaitingRoomUserResponse is response of Client.ListWaitingRoomUser
type ListWaitingRoomUserResponse struct {
	// HasMore 是否还有更多成员列表
	HasMore bool `json:"has_more"`
	// NextCursor 当has_more=true时有值，下次请求将该字段赋值给cursor字段
	NextCursor string `json:"next_cursor"`
	// UserList 用户列表对象数组
	UserList []map[string]any `json:"user_list"`
	// UserID 企业成员userid
	UserID string `json:"userid"`
	// TmpOpenID 会议中为每个参会成员授予的临时 ID
	TmpOpenID string `json:"tmp_openid"`
	// InstanceID 成员的终端设备类型
	InstanceID int `json:"instance_id"`
	// JoinTime 加入时间（单位：毫秒）
	JoinTime int `json:"join_time"`
	// QuitTime 离开时间（单位：毫秒）
	QuitTime int `json:"quit_time"`
}

// CancelWebinar 取消网络研讨会
// 取消指定的网络研讨会。
//
// see https://developer.work.weixin.qq.com/document/path/98870
func (c *Client) CancelWebinar(r *CancelWebinarRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/webinar/cancel",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CancelWebinarRequest is request of Client.CancelWebinar
type CancelWebinarRequest struct {
	// Meetingid 网络研讨会ID。
	Meetingid string `json:"meetingid" validate:"required"`
}

// CreateWebinar 创建网络研讨会
// 预定一场网络研讨会，适用于沙龙、圆桌会议、行业峰会等场景。
//
// see https://developer.work.weixin.qq.com/document/path/98842
func (c *Client) CreateWebinar(r *CreateWebinarRequest, opts ...any) (out CreateWebinarResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/webinar/create",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateWebinarRequest is request of Client.CreateWebinar
type CreateWebinarRequest struct {
	// AdminUserID 网络研讨会管理员userid
	AdminUserID string `json:"admin_userid" validate:"required"`
	// Title 网络研讨会主题（1~255位字符长度）
	Title string `json:"title" validate:"required"`
	// Sponsor 主办方名称（1~40位字符长度）
	Sponsor string `json:"sponsor"`
	// StartTime 会议开始时间戳（单位秒），不能少于当前时间戳半小时以上
	StartTime string `json:"start_time" validate:"required"`
	// EndTime 会议结束时间戳（单位秒）
	EndTime string `json:"end_time" validate:"required"`
	// AdmissionType 观众观看限制类型：0公开，1报名，2密码
	AdmissionType int `json:"admission_type" validate:"required"`
	// Hosts 主持人的成员 ID 列表
	Hosts []map[string]any `json:"hosts"`
	// Password 观众观看密码（4~6位数字）
	Password string `json:"password"`
	// CoverURL 封面图片 URL
	CoverURL string `json:"cover_url"`
	// Description 网络研讨会描述详情
	Description string `json:"description"`
	// EnableGuestInviteLink 是否开启通过邀请链接自动成为嘉宾
	EnableGuestInviteLink bool `json:"enable_guest_invite_link"`
	// MediaSetting 媒体参数配置
	MediaSetting any `json:"media_setting"`
	// EnableQa 是否开启问答
	EnableQa bool `json:"enable_qa"`
	// SensitiveWords 聊天敏感词列表
	SensitiveWords []string `json:"sensitive_words"`
	// EnableManualCheck 是否开启人工审核
	EnableManualCheck bool `json:"enable_manual_check"`
	// ActivityPage 活动页开启配置
	ActivityPage bool `json:"activity_page"`
	// DisplayNumberOfAttendees 活动页展示已报名或已预约人数
	DisplayNumberOfAttendees int `json:"display_number_of_attendees"`
	// PlaybackForAudience 允许观众观看回放
	PlaybackForAudience bool `json:"playback_for_audience" validate:"required"`
	// PreparationMode 是否开启准备模式
	PreparationMode bool `json:"preparation_mode"`
}

// CreateWebinarResponse is response of Client.CreateWebinar
type CreateWebinarResponse struct {
	// Title 网络研讨会主题
	Title string `json:"title"`
	// Meetingid 网络研讨会ID
	Meetingid string `json:"meetingid"`
	// MeetingCode 网络研讨会的会议号
	MeetingCode string `json:"meeting_code"`
	// StartTime 会议开始时间戳（单位秒）
	StartTime string `json:"start_time"`
	// EndTime 会议结束时间戳（单位秒）
	EndTime string `json:"end_time"`
	// AdmissionType 观众观看限制类型
	AdmissionType int `json:"admission_type"`
	// Password 观众观看密码
	Password string `json:"password"`
	// AudienceJoinLink 观众入会链接
	AudienceJoinLink string `json:"audience_join_link"`
	// GuestJoinLink 嘉宾入会链接
	GuestJoinLink string `json:"guest_join_link"`
	// ManualCheckLink 人工审核链接
	ManualCheckLink string `json:"manual_check_link"`
	// ManualCheckPassword 人工审核密码
	ManualCheckPassword string `json:"manual_check_password"`
}

// ApproveWebinarEnroll 审批网络研讨会报名信息
// 批量审批网络研讨会的报名信息
//
// see https://developer.work.weixin.qq.com/document/path/98877
func (c *Client) ApproveWebinarEnroll(r *ApproveWebinarEnrollRequest, opts ...any) (out ApproveWebinarEnrollResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/webinar/enroll/approve",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ApproveWebinarEnrollRequest is request of Client.ApproveWebinarEnroll
type ApproveWebinarEnrollRequest struct {
	// Meetingid 网络研讨会 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// EnrollIDList 报名 ID 列表
	EnrollIDList []string `json:"enroll_id_list" validate:"required"`
	// Action 审批动作：1：取消批准，2：拒绝，3：批准
	Action int `json:"action" validate:"required"`
}

// ApproveWebinarEnrollResponse is response of Client.ApproveWebinarEnroll
type ApproveWebinarEnrollResponse struct {
	// HandledCount 成功处理的数量
	HandledCount int `json:"handled_count"`
}

// DeleteWebinarEnroll 删除网络研讨会报名信息
// 删除指定网络研讨会（Webinar）的报名信息，支持删除成员手动报名的信息和导入的报名信息。
//
// see https://developer.work.weixin.qq.com/document/path/98881
func (c *Client) DeleteWebinarEnroll(r *DeleteWebinarEnrollRequest, opts ...any) (out DeleteWebinarEnrollResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/webinar/enroll/delete",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteWebinarEnrollRequest is request of Client.DeleteWebinarEnroll
type DeleteWebinarEnrollRequest struct {
	// Meetingid 网络研讨会ID
	Meetingid string `json:"meetingid" validate:"required"`
	// EnrollIDList 报名ID列表
	EnrollIDList []map[string]any `json:"enroll_id_list" validate:"required"`
}

// DeleteWebinarEnrollResponse is response of Client.DeleteWebinarEnroll
type DeleteWebinarEnrollResponse struct {
	// TotalCount 成功删除的报名信息数量
	TotalCount int `json:"total_count"`
}

// GetWebinarEnrollConfig 获取网络研讨会报名配置
// 获取网络研讨会的报名配置和报名问题。会议未开启报名时会返回未开启报名错误。
//
// see https://developer.work.weixin.qq.com/document/path/99297
func (c *Client) GetWebinarEnrollConfig(r *GetWebinarEnrollConfigRequest, opts ...any) (out GetWebinarEnrollConfigResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/webinar/enroll/get_config",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetWebinarEnrollConfigRequest is request of Client.GetWebinarEnrollConfig
type GetWebinarEnrollConfigRequest struct {
	// Meetingid 网络研讨会ID
	Meetingid string `json:"meetingid" validate:"required"`
}

// GetWebinarEnrollConfigResponse is response of Client.GetWebinarEnrollConfig
type GetWebinarEnrollConfigResponse struct {
	// ApproveType 审批类型：1：自动审批，默认自动审批；2：手动审批
	ApproveType int `json:"approve_type"`
	// IsCollectQuestion 是否收集问题：1：不收集，默认值为不收集；2：收集
	IsCollectQuestion int `json:"is_collect_question"`
	// NoRegistrationNeededForStaff 本企业成员无需报名：true：本企业成员无需报名；false：默认配置，本企业成员及企业外成员需要报名
	NoRegistrationNeededForStaff bool `json:"no_registration_needed_for_staff"`
	// QuestionList 报名问题列表
	QuestionList []map[string]any `json:"question_list"`
}

// ImportWebinarEnroll 导入网络研讨会报名信息
// 在指定网络研讨会（Webinar）中导入报名信息。
//
// see https://developer.work.weixin.qq.com/document/path/98880
func (c *Client) ImportWebinarEnroll(r *ImportWebinarEnrollRequest, opts ...any) (out ImportWebinarEnrollResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/webinar/enroll/import",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ImportWebinarEnrollRequest is request of Client.ImportWebinarEnroll
type ImportWebinarEnrollRequest struct {
	// Meetingid 网络研讨会ID
	Meetingid string `json:"meetingid" validate:"required"`
	// EnrollList 报名成员列表
	EnrollList []map[string]any `json:"enroll_list" validate:"required"`
}

// ImportWebinarEnrollResponse is response of Client.ImportWebinarEnroll
type ImportWebinarEnrollResponse struct {
	// TotalCount 成功导入的报名信息条数
	TotalCount int `json:"total_count"`
	// EnrollList 报名成员列表
	EnrollList []map[string]any `json:"enroll_list"`
}

// ListWebinarEnroll 获取网络研讨会报名信息
// 获取已报名观众数量和报名观众答题详情。会议未开启报名时会返回未开启报名错误。
//
// see https://developer.work.weixin.qq.com/document/path/99023
func (c *Client) ListWebinarEnroll(r *ListWebinarEnrollRequest, opts ...any) (out ListWebinarEnrollResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/webinar/enroll/list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListWebinarEnrollRequest is request of Client.ListWebinarEnroll
type ListWebinarEnrollRequest struct {
	// Meetingid 网络研讨会 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// Status 审批状态筛选字段，默认返回全部。0：全部，1：待审批，2：已拒绝，3：已批准
	Status int `json:"status"`
	// Cursor 分页查询用，将上一个请求返回的next_cursor字段传入。第一次查询时可不传值
	Cursor string `json:"cursor"`
	// Limit 分页大小，最大50条。默认值为50
	Limit int `json:"limit"`
}

// ListWebinarEnrollResponse is response of Client.ListWebinarEnroll
type ListWebinarEnrollResponse struct {
	// HasMore 是否还有待拉取的成员列表
	HasMore bool `json:"has_more"`
	// NextCursor 分页用，下一次拉取列表将该字段填入cursor字段
	NextCursor string `json:"next_cursor"`
	// EnrollList 当前页的报名列表
	EnrollList []map[string]any `json:"enroll_list"`
}

// QueryWebinarEnrollIdByTmpOpenid 获取网络研讨会成员报名 ID
// 支持获取网络研讨会中已报名成员的报名 ID。
//
// see https://developer.work.weixin.qq.com/document/path/99021
func (c *Client) QueryWebinarEnrollIdByTmpOpenid(r *QueryWebinarEnrollIdByTmpOpenidRequest, opts ...any) (out QueryWebinarEnrollIdByTmpOpenidResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/webinar/enroll/query_by_tmp_openid",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// QueryWebinarEnrollIdByTmpOpenidRequest is request of Client.QueryWebinarEnrollIdByTmpOpenid
type QueryWebinarEnrollIdByTmpOpenidRequest struct {
	// Meetingid 网络研讨会ID
	Meetingid string `json:"meetingid" validate:"required"`
	// SortingRules 查询报名 ID 的排序规则。1：优先查询手机号导入报名，再查询成员手动报名；2：优先查询成员手动报名，再查手机号导入
	SortingRules int `json:"sorting_rules"`
	// TmpOpenIDList 当场会议的成员临时 ID 数组，单次最多支持500条
	TmpOpenIDList []string `json:"tmp_openid_list" validate:"required"`
}

// QueryWebinarEnrollIdByTmpOpenidResponse is response of Client.QueryWebinarEnrollIdByTmpOpenid
type QueryWebinarEnrollIdByTmpOpenidResponse struct {
	// EnrollIDList 成员报名 ID 数组，仅返回已报名成员的报名 ID
	EnrollIDList []map[string]any `json:"enroll_id_list"`
}

// SetWebinarEnrollConfig 修改网络研讨会报名配置
// 修改网络研讨会的报名配置和报名问题，需要会议已开启报名。
//
// see https://developer.work.weixin.qq.com/document/path/99029
func (c *Client) SetWebinarEnrollConfig(r *SetWebinarEnrollConfigRequest, opts ...any) (out SetWebinarEnrollConfigResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/webinar/enroll/set_config",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetWebinarEnrollConfigRequest is request of Client.SetWebinarEnrollConfig
type SetWebinarEnrollConfigRequest struct {
	// Meetingid 网络研讨会 ID
	Meetingid string `json:"meetingid" validate:"required"`
	// ApproveType 审批类型：1：自动审批，默认自动审批；2：手动审批
	ApproveType int `json:"approve_type"`
	// IsCollectQuestion 是否收集问题：1：不收集，默认值为不收集；2：收集
	IsCollectQuestion int `json:"is_collect_question"`
	// NoRegistrationNeededForStaff 本企业成员无需报名：true：本企业成员无需报名；false：默认配置，本企业成员及企业外成员需要报名
	NoRegistrationNeededForStaff bool `json:"no_registration_needed_for_staff"`
	// QuestionList 报名问题列表
	QuestionList []map[string]any `json:"question_list"`
}

// SetWebinarEnrollConfigResponse is response of Client.SetWebinarEnrollConfig
type SetWebinarEnrollConfigResponse struct {
	// QuestionCount 报名问题数量，不收集问题时该字段返回0
	QuestionCount int `json:"question_count"`
}

// GetWebinarDetail 获取网络研讨会详情
// 获取指定的网络研讨会详细信息，支持通过会议 ID 或会议 Code 方式查询。
//
// see https://developer.work.weixin.qq.com/document/path/99020
func (c *Client) GetWebinarDetail(r *GetWebinarDetailRequest, opts ...any) (out GetWebinarDetailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/webinar/get",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetWebinarDetailRequest is request of Client.GetWebinarDetail
type GetWebinarDetailRequest struct {
	// Meetingid 网络研讨会ID，meetingid 和 meeting_code 二者必须送一个
	Meetingid string `json:"meetingid"`
	// MeetingCode 网络研讨会的会议号，meetingid 和 meeting_code 二者必须送一个
	MeetingCode string `json:"meeting_code"`
}

// GetWebinarDetailResponse is response of Client.GetWebinarDetail
type GetWebinarDetailResponse struct {
	// Subject 网络研讨会主题
	Subject string `json:"subject"`
	// Meetingid 网络研讨会ID
	Meetingid string `json:"meetingid"`
	// MeetingCode 网络研讨会的会议号
	MeetingCode string `json:"meeting_code"`
	// Status 当前会议状态
	Status string `json:"status"`
	// Sponsor 主办方名称
	Sponsor string `json:"sponsor"`
	// StartTime 会议开始时间戳（单位秒）
	StartTime string `json:"start_time"`
	// EndTime 会议结束时间戳（单位秒）
	EndTime string `json:"end_time"`
	// AdmissionType 观众观看限制类型
	AdmissionType int `json:"admission_type"`
	// Hosts 主持人的成员 ID 列表
	Hosts []map[string]any `json:"hosts"`
	// Password 观众观看密码
	Password string `json:"password"`
	// CoverURL 封面图片 URL
	CoverURL string `json:"cover_url"`
	// Description 网络研讨会描述详情
	Description string `json:"description"`
	// EnableGuestInviteLink 是否开启通过邀请链接自动成为嘉宾
	EnableGuestInviteLink bool `json:"enable_guest_invite_link"`
	// AudienceJoinLink 观众入会链接
	AudienceJoinLink string `json:"audience_join_link"`
	// GuestJoinLink 嘉宾入会链接
	GuestJoinLink string `json:"guest_join_link"`
	// MediaSetting 媒体参数配置
	MediaSetting any `json:"media_setting"`
	// EnableQa 是否开启问答
	EnableQa bool `json:"enable_qa"`
	// ManualCheckLink 人工审核链接
	ManualCheckLink string `json:"manual_check_link"`
	// ManualCheckPassword 人工审核密码
	ManualCheckPassword string `json:"manual_check_password"`
	// ActivityPage 活动页开启配置
	ActivityPage bool `json:"activity_page"`
	// DisplayNumberOfAttendees 活动页展示已报名或已预约人数
	DisplayNumberOfAttendees int `json:"display_number_of_attendees"`
	// PlaybackForAudience 允许观众观看回放
	PlaybackForAudience bool `json:"playback_for_audience"`
	// PlaybackURL 回放地址
	PlaybackURL string `json:"playback_url"`
	// PreparationMode 是否开启准备模式
	PreparationMode bool `json:"preparation_mode"`
	// WarmUpPicture 暖场图片地址
	WarmUpPicture string `json:"warm_up_picture"`
	// WarmUpVideo 暖场视频地址
	WarmUpVideo string `json:"warm_up_video"`
	// AllowAttendeesInviteOthers 允许参会者在暖场中邀请成员
	AllowAttendeesInviteOthers bool `json:"allow_attendees_invite_others"`
}

// ListWebinarGuest 获取网络研讨会嘉宾列表
// 获取指定的网络研讨会嘉宾列表，支持通过会议 ID 或会议 Code 方式查询。
//
// see https://developer.work.weixin.qq.com/document/path/99019
func (c *Client) ListWebinarGuest(r *ListWebinarGuestRequest, opts ...any) (out ListWebinarGuestResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/webinar/list_guest",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListWebinarGuestRequest is request of Client.ListWebinarGuest
type ListWebinarGuestRequest struct {
	// Meetingid 网络研讨会ID。meetingid 和 meeting_code 二者必须送一个，二者都送时以 meeing_id 为准。
	Meetingid string `json:"meetingid"`
	// MeetingCode 网络研讨会的会议号。meetingid 和 meeting_code 二者必须送一个，二者都送时以 meeing_id 为准。
	MeetingCode string `json:"meeting_code"`
}

// ListWebinarGuestResponse is response of Client.ListWebinarGuest
type ListWebinarGuestResponse struct {
	// Guests 网络研讨会嘉宾列表
	Guests []map[string]any `json:"guests"`
}

// UpdateWebinar 修改网络研讨会
// 修改指定的网络研讨会信息
//
// see https://developer.work.weixin.qq.com/document/path/98843
func (c *Client) UpdateWebinar(r *UpdateWebinarRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/webinar/update",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateWebinarRequest is request of Client.UpdateWebinar
type UpdateWebinarRequest struct {
	// Meetingid 网络研讨会ID
	Meetingid string `json:"meetingid" validate:"required"`
	// Title 网络研讨会主题（1~255位字符长度）
	Title string `json:"title" validate:"required"`
	// Sponsor 主办方名称（1~40位字符长度）
	Sponsor string `json:"sponsor"`
	// StartTime 会议开始时间戳（单位秒），不能少于当前时间戳半小时以上
	StartTime string `json:"start_time" validate:"required"`
	// EndTime 会议结束时间戳（单位秒）
	EndTime string `json:"end_time" validate:"required"`
	// AdmissionType 观众观看限制类型：0公开，1报名，2密码
	AdmissionType int `json:"admission_type" validate:"required"`
	// Hosts 主持人的成员ID列表
	Hosts []map[string]any `json:"hosts"`
	// Password 观众观看密码（4~6位数字），当admission_type=2时必传
	Password string `json:"password"`
	// CoverURL 封面图片URL，需开启活动页配置
	CoverURL string `json:"cover_url"`
	// Description 网络研讨会描述详情，仅支持纯文本，1~5000位字符长度
	Description string `json:"description"`
	// EnableGuestInviteLink 是否开启通过邀请链接自动成为嘉宾
	EnableGuestInviteLink bool `json:"enable_guest_invite_link"`
	// MediaSetting 媒体参数配置
	MediaSetting any `json:"media_setting"`
	// EnableQa 是否开启问答
	EnableQa bool `json:"enable_qa"`
	// SensitiveWords 聊天敏感词列表
	SensitiveWords []string `json:"sensitive_words"`
	// EnableManualCheck 是否开启人工审核
	EnableManualCheck bool `json:"enable_manual_check"`
	// ActivityPage 活动页开启配置
	ActivityPage bool `json:"activity_page"`
	// DisplayNumberOfAttendees 活动页展示已报名或已预约人数：0不展示，1展示
	DisplayNumberOfAttendees int `json:"display_number_of_attendees"`
	// PlaybackForAudience 允许观众观看回放
	PlaybackForAudience bool `json:"playback_for_audience" validate:"required"`
	// PreparationMode 是否开启准备模式
	PreparationMode bool `json:"preparation_mode"`
}

// UpdateWebinarGuestList 更新网络研讨会嘉宾列表
// 更新指定的网络研讨会嘉宾列表。
//
// see https://developer.work.weixin.qq.com/document/path/99091
func (c *Client) UpdateWebinarGuestList(r *UpdateWebinarGuestListRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/webinar/update_guest_list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateWebinarGuestListRequest is request of Client.UpdateWebinarGuestList
type UpdateWebinarGuestListRequest struct {
	// Meetingid 网络研讨会ID
	Meetingid string `json:"meetingid" validate:"required"`
	// Guests 网络研讨会嘉宾列表。详见GuestInfo
	Guests []map[string]any `json:"guests" validate:"required"`
}

// UpdateWebinarWarmUpConfig 管理网络研讨会暖场配置
// 对网络研讨会进行 Webinar 进行暖场设置。可订阅“网络研讨会暖场上传结果”事件用于获得上传结果。
//
// see https://developer.work.weixin.qq.com/document/path/99030
func (c *Client) UpdateWebinarWarmUpConfig(r *UpdateWebinarWarmUpConfigRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/meeting/webinar/update_warm_up",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateWebinarWarmUpConfigRequest is request of Client.UpdateWebinarWarmUpConfig
type UpdateWebinarWarmUpConfigRequest struct {
	// Meetingid 网络研讨会ID
	Meetingid string `json:"meetingid" validate:"required"`
	// WarmUpPicture 暖场图片地址。暖场图片与暖场视频只能选择一个，如果同时传入图片和视频，以图片为准。图片推荐 1280*720 尺寸，支持 png/jpg 格式，大小不超过 5M。
	WarmUpPicture string `json:"warm_up_picture"`
	// WarmUpVideo 暖场视频地址。暖场图片与暖场视频只能选择一个，如果同时传入图片和视频，以图片为准。推荐 1280*720 尺寸，支持 mp4 格式，大小不超过 1G。
	WarmUpVideo string `json:"warm_up_video"`
	// AllowAttendeesInviteOthers 允许参会者在暖场中邀请成员。true：允许，默认允许；false：不允许
	AllowAttendeesInviteOthers bool `json:"allow_attendees_invite_others"`
}

