package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// GetHealthReportStat 获取健康上报使用统计
// 获取指定日期的健康上报使用情况统计
//
// see https://developer.work.weixin.qq.com/document/path/93676
func (c *Client) GetHealthReportStat(r *GetHealthReportStatRequest, opts ...any) (out GetHealthReportStatResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/health/get_health_report_stat",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetHealthReportStatRequest is request of Client.GetHealthReportStat
type GetHealthReportStatRequest struct {
	// Date 具体某天的使用统计，最长支持获取30天前数据
	Date string `json:"date" validate:"required"`
}

// GetHealthReportStatResponse is response of Client.GetHealthReportStat
type GetHealthReportStatResponse struct {
	// Pv 应用使用次数
	Pv int `json:"pv"`
	// Uv 应用使用成员数
	Uv int `json:"uv"`
}

// GetReportAnswer 获取用户填写答案
// 通过此接口可以获取指定的健康上报任务详情。
//
// see https://developer.work.weixin.qq.com/document/path/93679
func (c *Client) GetReportAnswer(r *GetReportAnswerRequest, opts ...any) (out GetReportAnswerResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/health/get_report_answer",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetReportAnswerRequest is request of Client.GetReportAnswer
type GetReportAnswerRequest struct {
	// JobID 任务ID
	JobID string `json:"jobid" validate:"required"`
	// Date 具体某天任务的填写答案，仅支持获取最近14天数据
	Date string `json:"date" validate:"required"`
	// Offset 数据偏移量
	Offset int `json:"offset"`
	// Limit 拉取的数据量，最大值100
	Limit int `json:"limit"`
}

// GetReportAnswerResponse is response of Client.GetReportAnswer
type GetReportAnswerResponse struct {
	// Answers 答案列表
	Answers []map[string]any `json:"answers"`
}

// GetHealthReportJobInfo 获取健康上报任务详情
// 通过此接口可以获取指定的健康上报任务详情。
//
// see https://developer.work.weixin.qq.com/document/path/93678
func (c *Client) GetHealthReportJobInfo(r *GetHealthReportJobInfoRequest, opts ...any) (out GetHealthReportJobInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/health/get_report_job_info",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetHealthReportJobInfoRequest is request of Client.GetHealthReportJobInfo
type GetHealthReportJobInfoRequest struct {
	// JobID 任务ID
	JobID string `json:"jobid" validate:"required"`
	// Date 具体某天任务详情，仅支持获取最近14天数据
	Date string `json:"date" validate:"required"`
}

// GetHealthReportJobInfoResponse is response of Client.GetHealthReportJobInfo
type GetHealthReportJobInfoResponse struct {
	// JobInfo 任务详情对象
	JobInfo any `json:"job_info"`
	// Title 任务名称
	Title string `json:"title"`
	// Creator 发起人的userid
	Creator string `json:"creator"`
	// Type 任务类型。1：表示企业内部成员，2：表示家长
	Type int `json:"type"`
	// ApplyRange 适用范围
	ApplyRange any `json:"apply_range"`
	// ReportTo 汇报对象信息
	ReportTo any `json:"report_to"`
	// ReportType 上报方式。1表示仅上报一次，2表示每天都上报
	ReportType int `json:"report_type"`
	// SkipWeekend 非工作日是否需要上报。0表示周末需要上报，1表示周末不需要上报
	SkipWeekend int `json:"skip_weekend"`
	// FinishCnt 已填表人数
	FinishCnt int `json:"finish_cnt"`
	// QuestionTemplates 健康上报问题列表
	QuestionTemplates []any `json:"question_templates"`
}

// GetHealthReportJobIds 获取健康上报任务ID列表
// 通过此接口可以获取企业当前正在运行的上报任务ID列表。
//
// see https://developer.work.weixin.qq.com/document/path/93677
func (c *Client) GetHealthReportJobIds(r *GetHealthReportJobIdsRequest, opts ...any) (out GetHealthReportJobIdsResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/health/get_report_jobids",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetHealthReportJobIdsRequest is request of Client.GetHealthReportJobIds
type GetHealthReportJobIdsRequest struct {
	// Offset 分页，偏移量，默认为0
	Offset int `json:"offset"`
	// Limit 分页，预期请求的数据量，默认为100，取值范围 1 ~ 100
	Limit int `json:"limit"`
}

// GetHealthReportJobIdsResponse is response of Client.GetHealthReportJobIds
type GetHealthReportJobIdsResponse struct {
	// Ending 是否结束。0：表示还有更多数据，需要继续拉取，1：表示已经拉取完所有数据
	Ending int `json:"ending"`
	// Jobids 任务id列表
	Jobids []string `json:"jobids"`
}

// DeleteLivingReplayData 删除直播回放
// 删除指定的直播回放数据
//
// see https://developer.work.weixin.qq.com/document/path/93743
func (c *Client) DeleteLivingReplayData(r *DeleteLivingReplayDataRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/living/delete_replay_data",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteLivingReplayDataRequest is request of Client.DeleteLivingReplayData
type DeleteLivingReplayDataRequest struct {
	// Livingid 直播id
	Livingid string `json:"livingid" validate:"required"`
}

// ListUserLivingId 获取老师直播ID列表
// 通过此接口可以获取指定老师的所有直播ID
//
// see https://developer.work.weixin.qq.com/document/path/93739
func (c *Client) ListUserLivingId(r *ListUserLivingIdRequest, opts ...any) (out ListUserLivingIdResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/living/get_user_all_livingid",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListUserLivingIdRequest is request of Client.ListUserLivingId
type ListUserLivingIdRequest struct {
	// UserID 企业成员的userid
	UserID string `json:"userid" validate:"required"`
	// Cursor 上一次调用时返回的next_cursor，第一次拉取可以不填
	Cursor string `json:"cursor"`
	// Limit 每次拉取的数据量，默认值和最大值都为100
	Limit int `json:"limit"`
}

// ListUserLivingIdResponse is response of Client.ListUserLivingId
type ListUserLivingIdResponse struct {
	// NextCursor 当前数据最后一个key值，如果下次调用带上该值则从该key值往后拉，用于实现分页拉取，返回空字符串代表已经是最后一页
	NextCursor string `json:"next_cursor"`
	// LivingidList 直播ID列表
	LivingidList []string `json:"livingid_list"`
}

// GetPaymentResult 获取学生付款结果
// 获取由应用本身创建的收款项目的付款详情
//
// see https://developer.work.weixin.qq.com/document/path/94470
func (c *Client) GetPaymentResult(r *GetPaymentResultRequest, opts ...any) (out GetPaymentResultResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/get_payment_result",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetPaymentResultRequest is request of Client.GetPaymentResult
type GetPaymentResultRequest struct {
	// PaymentID 收款项目id，由jssdk的发起班级收款接口或者小程序的发起班级收款接口返回
	PaymentID string `json:"payment_id" validate:"required"`
}

// GetPaymentResultResponse is response of Client.GetPaymentResult
type GetPaymentResultResponse struct {
	// ProjectName 收款项目名称
	ProjectName string `json:"project_name"`
	// Amount 收款金额
	Amount int `json:"amount"`
	// PaymentResult 学生付款信息列表
	PaymentResult []map[string]any `json:"payment_result"`
}

// GetTradeDetail 获取订单详情
// 获取企业微信班级收款订单详情
//
// see https://developer.work.weixin.qq.com/document/path/94471
func (c *Client) GetTradeDetail(r *GetTradeDetailRequest, opts ...any) (out GetTradeDetailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/get_trade",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetTradeDetailRequest is request of Client.GetTradeDetail
type GetTradeDetailRequest struct {
	// PaymentID 收款项目id，由发起班级收款接口返回
	PaymentID string `json:"payment_id" validate:"required"`
	// TradeNo 订单号，由获取学生付款结果返回
	TradeNo string `json:"trade_no" validate:"required"`
}

// GetTradeDetailResponse is response of Client.GetTradeDetail
type GetTradeDetailResponse struct {
	// TransactionID 微信交易单号
	TransactionID string `json:"transaction_id"`
	// PayTime 交易时间
	PayTime int `json:"pay_time"`
}

// GetLivingInfo 获取直播详情
// 获取企业直播的详细信息
//
// see https://developer.work.weixin.qq.com/document/path/93740
func (c *Client) GetLivingInfo(r *GetLivingInfoRequest, opts ...any) (out GetLivingInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/school/living/get_living_info",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetLivingInfoRequest is request of Client.GetLivingInfo
type GetLivingInfoRequest struct {
	// Livingid 直播ID
	Livingid string `json:"livingid" validate:"required"`
}

// GetLivingInfoResponse is response of Client.GetLivingInfo
type GetLivingInfoResponse struct {
	// LivingInfo 直播信息
	LivingInfo any `json:"living_info"`
}

// GetUnwatchLivingStat 获取未观看直播统计
// 获取未观看直播的学生统计，学生的家长必须是已经关注「学校通知」才会纳入统计范围。
//
// see https://developer.work.weixin.qq.com/document/path/93742
func (c *Client) GetUnwatchLivingStat(r *GetUnwatchLivingStatRequest, opts ...any) (out GetUnwatchLivingStatResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/living/get_unwatch_stat",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetUnwatchLivingStatRequest is request of Client.GetUnwatchLivingStat
type GetUnwatchLivingStatRequest struct {
	// Livingid 直播id
	Livingid string `json:"livingid" validate:"required"`
	// NextKey 上一次调用时返回的next_key，初次调用可以填"0"
	NextKey string `json:"next_key"`
}

// GetUnwatchLivingStatResponse is response of Client.GetUnwatchLivingStat
type GetUnwatchLivingStatResponse struct {
	// Ending 是否结束。0：表示还有更多数据，需要继续拉取，1：表示已经拉取完所有数据
	Ending int `json:"ending"`
	// NextKey 当前数据最后一个key值，用于分页拉取
	NextKey string `json:"next_key"`
	// StatInfo 统计信息列表
	StatInfo any `json:"stat_info"`
}

// GetUnwatchStatV2 获取未观看直播统计V2
// 获取未观看直播的学生和家长统计，学生和学生的家长必须是已经关注「学校通知」才会纳入统计范围。
//
// see https://developer.work.weixin.qq.com/document/path/95795
func (c *Client) GetUnwatchStatV2(r *GetUnwatchStatV2Request, opts ...any) (out GetUnwatchStatV2Response, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/living/get_unwatch_stat_v2",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetUnwatchStatV2Request is request of Client.GetUnwatchStatV2
type GetUnwatchStatV2Request struct {
	// Livingid 直播id
	Livingid string `json:"livingid" validate:"required"`
	// NextCursor 上一次调用时返回的next_cursor，初次调用可以填"0"
	NextCursor string `json:"next_cursor"`
}

// GetUnwatchStatV2Response is response of Client.GetUnwatchStatV2
type GetUnwatchStatV2Response struct {
	// HasMore 是否结束。1：表示还有更多数据，需要继续拉取，0：表示已经拉取完所有数据
	HasMore int `json:"has_more"`
	// NextCursor 当前数据最后一个cursor值
	NextCursor string `json:"next_cursor"`
	// StatInfo 统计信息列表
	StatInfo any `json:"stat_info"`
}

// GetLivingWatchStat 获取观看直播统计
// 通过该接口可以获取所有观看直播的人员统计
//
// see https://developer.work.weixin.qq.com/document/path/93741
func (c *Client) GetLivingWatchStat(r *GetLivingWatchStatRequest, opts ...any) (out GetLivingWatchStatResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/living/get_watch_stat",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetLivingWatchStatRequest is request of Client.GetLivingWatchStat
type GetLivingWatchStatRequest struct {
	// Livingid 直播的id
	Livingid string `json:"livingid" validate:"required"`
	// NextKey 上一次调用时返回的next_key，初次调用可以填0
	NextKey string `json:"next_key"`
}

// GetLivingWatchStatResponse is response of Client.GetLivingWatchStat
type GetLivingWatchStatResponse struct {
	// Ending 是否结束。0：表示还有更多数据，需要继续拉取，1：表示已经拉取完所有数据
	Ending int `json:"ending"`
	// NextKey 当前数据最后一个key值，用于分页拉取
	NextKey string `json:"next_key"`
	// StatInfoes 统计信息列表
	StatInfoes any `json:"stat_infoes"`
}

// GetWatchStatV2 获取观看直播统计V2
// 通过该接口可以获取所有观看直播的人员统计
//
// see https://developer.work.weixin.qq.com/document/path/95793
func (c *Client) GetWatchStatV2(r *GetWatchStatV2Request, opts ...any) (out GetWatchStatV2Response, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/living/get_watch_stat_v2",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetWatchStatV2Request is request of Client.GetWatchStatV2
type GetWatchStatV2Request struct {
	// Livingid 直播的id
	Livingid string `json:"livingid" validate:"required"`
	// NextCursor 上一次调用时返回的next_cursor，初次调用可以填"0"
	NextCursor string `json:"next_cursor"`
}

// GetWatchStatV2Response is response of Client.GetWatchStatV2
type GetWatchStatV2Response struct {
	// HasMore 是否结束。1：表示还有更多数据，需要继续拉取，0：表示已经拉取完所有数据
	HasMore int `json:"has_more"`
	// NextCursor 当前数据最后一个cursor值，用于分页拉取
	NextCursor string `json:"next_cursor"`
	// StatInfo 统计信息列表
	StatInfo any `json:"stat_info"`
}

