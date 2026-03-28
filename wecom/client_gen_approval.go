package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// GetApprovalData 获取审批数据（旧）
// 通过本接口来获取公司一段时间内的审批记录。一次拉取调用最多拉取100个审批记录，可以通过多次拉取的方式来满足需求。
//
// see https://developer.work.weixin.qq.com/document/path/91530
func (c *Client) GetApprovalData(r *GetApprovalDataRequest, opts ...any) (out GetApprovalDataResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/corp/getapprovaldata",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetApprovalDataRequest is request of Client.GetApprovalData
type GetApprovalDataRequest struct {
	// Starttime 获取审批记录的开始时间，Unix时间戳
	Starttime int `json:"starttime" validate:"required"`
	// Endtime 获取审批记录的结束时间，Unix时间戳
	Endtime int `json:"endtime" validate:"required"`
	// NextSpnum 第一个拉取的审批单号，不填从该时间段的第一个审批单拉取
	NextSpnum int `json:"next_spnum"`
}

// GetApprovalDataResponse is response of Client.GetApprovalData
type GetApprovalDataResponse struct {
	// Count 拉取的审批单个数
	Count int `json:"count"`
	// Total 时间段内的总审批单个数
	Total int `json:"total"`
	// NextSpnum 拉取列表的最后一个审批单号
	NextSpnum int `json:"next_spnum"`
	// Data 审批记录列表
	Data []map[string]any `json:"data"`
	// Spname 审批名称
	Spname string `json:"spname"`
	// ApplyName 申请人姓名
	ApplyName string `json:"apply_name"`
	// ApplyOrg 申请人部门
	ApplyOrg string `json:"apply_org"`
	// ApprovalName 审批人姓名列表
	ApprovalName []any `json:"approval_name"`
	// NotifyName 抄送人姓名列表
	NotifyName []any `json:"notify_name"`
	// SpStatus 审批状态：1审批中；2已通过；3已驳回；4已取消；6通过后撤销；10已支付
	SpStatus int `json:"sp_status"`
	// SpNum 审批单号
	SpNum string `json:"sp_num"`
	// ApplyTime 审批单提交时间
	ApplyTime int `json:"apply_time"`
	// ApplyUserID 审批单提交者的userid
	ApplyUserID string `json:"apply_user_id"`
	// Mediaids 审批的附件media_id列表
	Mediaids []any `json:"mediaids"`
	// Expense 报销类型数据项
	Expense any `json:"expense"`
	// ExpenseType 报销类型：1差旅费；2交通费；3招待费；4其他报销
	ExpenseType int `json:"expense_type"`
	// Reason 报销事由
	Reason string `json:"reason"`
	// Item 报销明细列表
	Item []any `json:"item"`
	// ExpenseitemType 费用类型
	ExpenseitemType int `json:"expenseitem_type"`
	// Sums 费用金额，单位元
	Sums int `json:"sums"`
	// Comm 审批模板信息
	Comm any `json:"comm"`
	// ApplyData 审批申请的单据数据（JSON字符串）
	ApplyData string `json:"apply_data"`
	// Leave 请假类型数据项
	Leave any `json:"leave"`
	// Timeunit 请假时间单位：0半天；1小时
	Timeunit int `json:"timeunit"`
	// LeaveType 请假类型：1年假；2事假；3病假；4调休假；5婚假；6产假；7陪产假；8其他
	LeaveType int `json:"leave_type"`
	// StartTime 请假开始时间，unix时间
	StartTime int `json:"start_time"`
	// EndTime 请假结束时间，unix时间
	EndTime int `json:"end_time"`
	// Duration 请假时长，单位小时
	Duration int `json:"duration"`
}

// GetOpenApprovalData 查询自建应用审批单当前状态
// 开发者可主动查询审批单的当前审批状态。
//
// see https://developer.work.weixin.qq.com/document/path/90090
func (c *Client) GetOpenApprovalData(r *GetOpenApprovalDataRequest, opts ...any) (out GetOpenApprovalDataResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/corp/getopenapprovaldata",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetOpenApprovalDataRequest is request of Client.GetOpenApprovalData
type GetOpenApprovalDataRequest struct {
	// ThirdNo 开发者发起申请时定义的审批单号
	ThirdNo string `json:"thirdNo" validate:"required"`
}

// GetOpenApprovalDataResponse is response of Client.GetOpenApprovalData
type GetOpenApprovalDataResponse struct {
	// Data 审批单详情数据
	Data any `json:"data"`
	// ThirdNo 审批单编号
	ThirdNo string `json:"ThirdNo"`
	// OpenTemplateId 审批模板id
	OpenTemplateId string `json:"OpenTemplateId"`
	// OpenSpName 审批模板名称
	OpenSpName string `json:"OpenSpName"`
	// OpenSpstatus 申请单当前审批状态：1-审批中；2-已通过；3-已驳回；4-已撤销
	OpenSpstatus int `json:"OpenSpstatus"`
	// ApplyTime 提交申请时间
	ApplyTime int `json:"ApplyTime"`
	// ApplyUsername 提交者姓名
	ApplyUsername string `json:"ApplyUsername"`
	// ApplyUserParty 提交者所在部门
	ApplyUserParty string `json:"ApplyUserParty"`
	// ApplyUserImage 提交者头像
	ApplyUserImage string `json:"ApplyUserImage"`
	// ApplyUserId 提交者userid
	ApplyUserId string `json:"ApplyUserId"`
	// ApprovalNodes 审批流程信息
	ApprovalNodes any `json:"ApprovalNodes"`
	// NotifyNodes 抄送信息
	NotifyNodes any `json:"NotifyNodes"`
}

// SubmitApprovalEvent 提交审批申请
// 提交审批申请单
//
// see https://developer.work.weixin.qq.com/document/path/91853
func (c *Client) SubmitApprovalEvent(r *SubmitApprovalEventRequest, opts ...any) (out SubmitApprovalEventResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/applyevent",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SubmitApprovalEventRequest is request of Client.SubmitApprovalEvent
type SubmitApprovalEventRequest struct {
	// CreatorUserID 申请人userid
	CreatorUserID string `json:"creator_userid" validate:"required"`
	// TemplateID 模板id
	TemplateID string `json:"template_id" validate:"required"`
	// UseTemplateApprover 审批人模式：0-通过接口指定审批人、抄送人；1-使用此模板在管理后台设置的审批流程
	UseTemplateApprover int `json:"use_template_approver" validate:"required"`
	// ChooseDepartment 提单者提单部门id，不填默认为主部门
	ChooseDepartment int `json:"choose_department"`
	// ApplyData 审批申请数据，包含各控件的值
	ApplyData any `json:"apply_data" validate:"required"`
	// SummaryList 摘要信息，用于显示在审批通知卡片、审批列表的摘要信息，最多3行
	SummaryList []map[string]any `json:"summary_list" validate:"required"`
	// Process 新版流程列表（当use_template_approver=0时必填）
	Process any `json:"process"`
}

// SubmitApprovalEventResponse is response of Client.SubmitApprovalEvent
type SubmitApprovalEventResponse struct {
	// SpNo 表单提交成功后返回的表单编号
	SpNo string `json:"sp_no"`
}

// CreateApprovalTemplate 创建审批模板
// 调用此接口创建审批模板。创建新模板后，管理后台及审批应用内将生成对应模板，并生效默认流程和规则配置。
//
// see https://developer.work.weixin.qq.com/document/path/97437
func (c *Client) CreateApprovalTemplate(r *CreateApprovalTemplateRequest, opts ...any) (out CreateApprovalTemplateResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/approval/create_template",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateApprovalTemplateRequest is request of Client.CreateApprovalTemplate
type CreateApprovalTemplateRequest struct {
	// TemplateName 模版名称数组
	TemplateName []map[string]any `json:"template_name" validate:"required"`
	// TemplateContent 审批模版控件设置，由多个表单控件及其内容组成
	TemplateContent any `json:"template_content" validate:"required"`
}

// CreateApprovalTemplateResponse is response of Client.CreateApprovalTemplate
type CreateApprovalTemplateResponse struct {
	// TemplateID 模版创建成功后返回的模版id
	TemplateID string `json:"template_id"`
}

// UpdateApprovalTemplate 更新审批模板
// 可调用本接口更新审批模板。更新模板后，管理后台及审批应用内将更新原模板的内容，已配置的审批流程和规则不变。
//
// see https://developer.work.weixin.qq.com/document/path/97438
func (c *Client) UpdateApprovalTemplate(r *UpdateApprovalTemplateRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/approval/update_template",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateApprovalTemplateRequest is request of Client.UpdateApprovalTemplate
type UpdateApprovalTemplateRequest struct {
	// TemplateID 模版id
	TemplateID string `json:"template_id" validate:"required"`
	// TemplateName 模版名称数组
	TemplateName []map[string]any `json:"template_name" validate:"required"`
	// TemplateContent 审批模版控件设置
	TemplateContent any `json:"template_content" validate:"required"`
}

// GetApprovalDetail 获取审批申请详情
// 获取指定审批单编号的审批申请详细信息
//
// see https://developer.work.weixin.qq.com/document/path/91983
func (c *Client) GetApprovalDetail(r *GetApprovalDetailRequest, opts ...any) (out GetApprovalDetailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/getapprovaldetail",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetApprovalDetailRequest is request of Client.GetApprovalDetail
type GetApprovalDetailRequest struct {
	// SpNo 审批单编号
	SpNo string `json:"sp_no" validate:"required"`
}

// GetApprovalDetailResponse is response of Client.GetApprovalDetail
type GetApprovalDetailResponse struct {
	// Info 审批申请详情信息
	Info any `json:"info"`
	// SpNo 审批编号
	SpNo string `json:"sp_no"`
	// SpName 审批申请类型名称
	SpName string `json:"sp_name"`
	// SpStatus 申请单状态：1-审批中；2-已通过；3-已驳回等
	SpStatus int `json:"sp_status"`
	// TemplateID 审批模板id
	TemplateID string `json:"template_id"`
	// ApplyTime 审批申请提交时间,Unix时间戳
	ApplyTime int `json:"apply_time"`
	// Applyer 申请人信息
	Applyer any `json:"applyer"`
	// BatchApplyer 批量申请人信息
	BatchApplyer []any `json:"batch_applyer"`
	// SpRecord 审批流程信息
	SpRecord []any `json:"sp_record"`
	// Notifyer 抄送信息
	Notifyer []any `json:"notifyer"`
	// ApplyData 审批申请数据
	ApplyData any `json:"apply_data"`
	// Comments 审批申请备注信息
	Comments []any `json:"comments"`
	// ProcessList 审批流程列表
	ProcessList any `json:"process_list"`
}

// GetApprovalInfo 批量获取审批单号
// 获取企业一段时间内企业微信“审批应用”单据的审批编号，支持按模板类型、申请人、部门、申请单审批状态等条件筛选。
//
// see https://developer.work.weixin.qq.com/document/path/91816
func (c *Client) GetApprovalInfo(r *GetApprovalInfoRequest, opts ...any) (out GetApprovalInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/getapprovalinfo",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetApprovalInfoRequest is request of Client.GetApprovalInfo
type GetApprovalInfoRequest struct {
	// Starttime 审批单提交的时间范围，开始时间，Unix时间戳
	Starttime string `json:"starttime" validate:"required"`
	// Endtime 审批单提交的时间范围，结束时间，Unix时间戳
	Endtime string `json:"endtime" validate:"required"`
	// NewCursor 分页查询游标，默认为空串，后续使用返回的new_next_cursor进行分页拉取
	NewCursor string `json:"new_cursor" validate:"required"`
	// Size 一次请求拉取审批单数量，默认值为100，上限值为100
	Size int `json:"size" validate:"required"`
	// Filters 筛选条件，可对批量拉取的审批申请设置约束条件，支持设置多个条件
	Filters []map[string]any `json:"filters"`
}

// GetApprovalInfoResponse is response of Client.GetApprovalInfo
type GetApprovalInfoResponse struct {
	// SpNoList 审批单号列表，包含满足条件的审批申请
	SpNoList []string `json:"sp_no_list"`
	// NewNextCursor 后续请求查询的游标，当返回结果没有该字段时表示审批单已经拉取完
	NewNextCursor string `json:"new_next_cursor"`
}

// GetApprovalTemplateDetail 获取审批模板详情
// 获取指定审批模板的详细信息，包括模板名称和控件配置。
//
// see https://developer.work.weixin.qq.com/document/path/91982
func (c *Client) GetApprovalTemplateDetail(r *GetApprovalTemplateDetailRequest, opts ...any) (out GetApprovalTemplateDetailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/gettemplatedetail",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetApprovalTemplateDetailRequest is request of Client.GetApprovalTemplateDetail
type GetApprovalTemplateDetailRequest struct {
	// TemplateID 模板的唯一标识id。可在获取审批单据详情、审批状态变化回调通知中获得，也可在审批模板的模板编辑页面浏览器Url链接中获得。
	TemplateID string `json:"template_id" validate:"required"`
}

// GetApprovalTemplateDetailResponse is response of Client.GetApprovalTemplateDetail
type GetApprovalTemplateDetailResponse struct {
	// TemplateNames 模板名称列表，若配置了多语言则包含中英文
	TemplateNames []map[string]any `json:"template_names"`
	// TemplateContent 模板控件详细信息
	TemplateContent any `json:"template_content"`
}

// GetCorpVacationConfig 获取企业假期管理配置
// 获取可见范围内员工的假期管理配置，包括各个假期的id、名称、请假单位、时长计算方式、发放规则等。
//
// see https://developer.work.weixin.qq.com/document/path/93388
func (c *Client) GetCorpVacationConfig(r *GetCorpVacationConfigRequest, opts ...any) (out GetCorpVacationConfigResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/oa/vacation/getcorpconf",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCorpVacationConfigRequest is request of Client.GetCorpVacationConfig
type GetCorpVacationConfigRequest struct {}

// GetCorpVacationConfigResponse is response of Client.GetCorpVacationConfig
type GetCorpVacationConfigResponse struct {
	// Lists 假期列表
	Lists []map[string]any `json:"lists"`
	// ID 假期id
	ID int `json:"id"`
	// Name 假期名称
	Name string `json:"name"`
	// TimeAttr 假期时间刻度：0-按天请假；1-按小时请假
	TimeAttr int `json:"time_attr"`
	// DurationType 时长计算类型：0-自然日；1-工作日
	DurationType int `json:"duration_type"`
	// QuotaAttr 假期发放相关配置
	QuotaAttr any `json:"quota_attr"`
	// PerdayDuration 单位换算值，即1天对应的秒数
	PerdayDuration int `json:"perday_duration"`
	// IsNewovertime 是否关联加班调休，0-不关联，1-关联
	IsNewovertime int `json:"is_newovertime"`
	// EnterCompTimeLimit 入职时间大于n个月可用该假期，单位为月
	EnterCompTimeLimit int `json:"enter_comp_time_limit"`
	// ExpireRule 假期过期规则
	ExpireRule any `json:"expire_rule"`
}

// GetUserVacationQuota 获取成员假期余额
// 通过本接口可获取应用可见范围内各个员工的假期余额数据。
//
// see https://developer.work.weixin.qq.com/document/path/93376
func (c *Client) GetUserVacationQuota(r *GetUserVacationQuotaRequest, opts ...any) (out GetUserVacationQuotaResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/vacation/getuservacationquota",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetUserVacationQuotaRequest is request of Client.GetUserVacationQuota
type GetUserVacationQuotaRequest struct {
	// UserID 需要获取假期余额的成员的userid
	UserID string `json:"userid" validate:"required"`
}

// GetUserVacationQuotaResponse is response of Client.GetUserVacationQuota
type GetUserVacationQuotaResponse struct {
	// Lists 假期列表
	Lists []map[string]any `json:"lists"`
}

// SetUserQuota 修改成员假期余额
// 通过本接口可以修改可见范围内员工的“假期余额”。
//
// see https://developer.work.weixin.qq.com/document/path/93389
func (c *Client) SetUserQuota(r *SetUserQuotaRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/oa/vacation/setoneuserquota",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetUserQuotaRequest is request of Client.SetUserQuota
type SetUserQuotaRequest struct {
	// UserID 需要修改假期余额的成员的userid
	UserID string `json:"userid" validate:"required"`
	// VacationID 假期id
	VacationID int `json:"vacation_id" validate:"required"`
	// Leftduration 设置的假期余额，单位为秒。不能大于1000天或24000小时，当假期时间刻度为按小时请假时，必须为360整倍数；按天请假时，必须为8640整倍数
	Leftduration int `json:"leftduration" validate:"required"`
	// TimeAttr 假期时间刻度：0-按天请假；1-按小时请假。主要用于校验，必须等于企业假期管理配置中设置的假期时间刻度类型
	TimeAttr int `json:"time_attr" validate:"required"`
	// Remarks 修改备注，用于显示在假期余额的修改记录当中，可对修改行为作说明，不超过200字符
	Remarks string `json:"remarks"`
}

