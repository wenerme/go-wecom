package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// AddKfAccount 添加客服账号
// 添加客服账号，并可设置客服名称和头像。目前一家企业最多可添加5000个客服账号。
//
// see https://developer.work.weixin.qq.com/document/path/94662
func (c *Client) AddKfAccount(r *AddKfAccountRequest, opts ...any) (out AddKfAccountResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/kf/account/add",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddKfAccountRequest is request of Client.AddKfAccount
type AddKfAccountRequest struct {
	// Name 客服名称，不多于16个字符
	Name string `json:"name" validate:"required"`
	// MediaID 客服头像临时素材。可以调用上传临时素材接口获取。不多于128个字节
	MediaID string `json:"media_id" validate:"required"`
}

// AddKfAccountResponse is response of Client.AddKfAccount
type AddKfAccountResponse struct {
	// OpenKfid 新创建的客服账号ID
	OpenKfid string `json:"open_kfid"`
}

// DeleteKfAccount 删除客服账号
// 删除已有的客服账号
//
// see https://developer.work.weixin.qq.com/document/path/94663
func (c *Client) DeleteKfAccount(r *DeleteKfAccountRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/kf/account/del",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteKfAccountRequest is request of Client.DeleteKfAccount
type DeleteKfAccountRequest struct {
	// OpenKfid 客服账号ID。不多于64字节
	OpenKfid string `json:"open_kfid" validate:"required"`
}

// ListKfAccount 获取客服账号列表
// 获取客服账号列表，包括客服账号的客服ID、名称和头像。
//
// see https://developer.work.weixin.qq.com/document/path/94706
func (c *Client) ListKfAccount(r *ListKfAccountRequest, opts ...any) (out ListKfAccountResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/kf/account/list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListKfAccountRequest is request of Client.ListKfAccount
type ListKfAccountRequest struct {
	// Offset 分页，偏移量，默认为0
	Offset int `json:"offset"`
	// Limit 分页，预期请求的数据量，默认为100，取值范围 1 ~ 100
	Limit int `json:"limit"`
}

// ListKfAccountResponse is response of Client.ListKfAccount
type ListKfAccountResponse struct {
	// AccountList 账号信息列表
	AccountList []map[string]any `json:"account_list"`
}

// UpdateKfAccount 修改客服账号
// 修改已有的客服账号，可修改客服名称和头像。
//
// see https://developer.work.weixin.qq.com/document/path/94682
func (c *Client) UpdateKfAccount(r *UpdateKfAccountRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/kf/account/update",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateKfAccountRequest is request of Client.UpdateKfAccount
type UpdateKfAccountRequest struct {
	// OpenKfid 要修改的客服账号ID。不多于64字节
	OpenKfid string `json:"open_kfid" validate:"required"`
	// Name 新的客服名称，如不需要修改可不填。不多于16个字符
	Name string `json:"name"`
	// MediaID 新的客服头像临时素材，如不需要修改可不填。可以调用上传临时素材接口获取。不多于128个字节
	MediaID string `json:"media_id"`
}

// BatchGetCustomer 获取客户基础信息
// 批量获取微信客服客户的详细信息
//
// see https://developer.work.weixin.qq.com/document/path/95171
func (c *Client) BatchGetCustomer(r *BatchGetCustomerRequest, opts ...any) (out BatchGetCustomerResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/kf/customer/batchget",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchGetCustomerRequest is request of Client.BatchGetCustomer
type BatchGetCustomerRequest struct {
	// ExternalUserIDList external_userid列表，可填充个数1~100
	ExternalUserIDList []string `json:"external_userid_list" validate:"required"`
	// NeedEnterSessionContext 是否需要返回客户48小时内最后一次进入会话的上下文信息，0-不返回 1-返回。默认不返回
	NeedEnterSessionContext int `json:"need_enter_session_context"`
}

// BatchGetCustomerResponse is response of Client.BatchGetCustomer
type BatchGetCustomerResponse struct {
	// CustomerList 客户列表，包含详细信息
	CustomerList []map[string]any `json:"customer_list"`
	// InvalidExternalUserID 无效的external_userid列表
	InvalidExternalUserID []string `json:"invalid_external_userid"`
}

// GetUpgradeServiceConfig 获取配置的专员与客户群
// 企业需要在管理后台或移动端中的「微信客服」-「升级服务」中，配置专员和客户群。该接口提供获取配置的专员与客户群列表的能力。
//
// see https://developer.work.weixin.qq.com/document/path/94763
func (c *Client) GetUpgradeServiceConfig(r *GetUpgradeServiceConfigRequest, opts ...any) (out GetUpgradeServiceConfigResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/kf/customer/get_upgrade_service_config",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetUpgradeServiceConfigRequest is request of Client.GetUpgradeServiceConfig
type GetUpgradeServiceConfigRequest struct {}

// GetUpgradeServiceConfigResponse is response of Client.GetUpgradeServiceConfig
type GetUpgradeServiceConfigResponse struct {
	// MemberRange 专员服务配置范围
	MemberRange any `json:"member_range"`
	// GroupchatRange 客户群配置范围
	GroupchatRange any `json:"groupchat_range"`
}

// GetCorpStatistic 获取「客户数据统计」企业汇总数据
// 通过此接口，可以获取咨询会话数、咨询客户数等企业汇总统计数据
//
// see https://developer.work.weixin.qq.com/document/path/95489
func (c *Client) GetCorpStatistic(r *GetCorpStatisticRequest, opts ...any) (out GetCorpStatisticResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/kf/get_corp_statistic",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCorpStatisticRequest is request of Client.GetCorpStatistic
type GetCorpStatisticRequest struct {
	// OpenKfid 客服账号ID
	OpenKfid string `json:"open_kfid" validate:"required"`
	// StartTime 起始日期的时间戳，填这一天的0时0分0秒。取值范围：昨天至前180天
	StartTime int `json:"start_time" validate:"required"`
	// EndTime 结束日期的时间戳，填这一天的0时0分0秒。取值范围：昨天至前180天
	EndTime int `json:"end_time" validate:"required"`
}

// GetCorpStatisticResponse is response of Client.GetCorpStatistic
type GetCorpStatisticResponse struct {
	// StatisticList 统计数据列表
	StatisticList []map[string]any `json:"statistic_list"`
	// StatTime 数据统计日期，为当日0点的时间戳
	StatTime string `json:"stat_time"`
	// SessionCnt 咨询会话数
	SessionCnt string `json:"session_cnt"`
	// CustomerCnt 咨询客户数
	CustomerCnt string `json:"customer_cnt"`
	// CustomerMsgCnt 咨询消息总数
	CustomerMsgCnt string `json:"customer_msg_cnt"`
	// UpgradeServiceCustomerCnt 升级服务客户数
	UpgradeServiceCustomerCnt string `json:"upgrade_service_customer_cnt"`
	// AiSessionReplyCnt 智能回复会话数
	AiSessionReplyCnt string `json:"ai_session_reply_cnt"`
	// AiTransferRate 转人工率
	AiTransferRate float64 `json:"ai_transfer_rate"`
	// AiKnowledgeHitRate 知识命中率
	AiKnowledgeHitRate float64 `json:"ai_knowledge_hit_rate"`
	// MsgRejectedCustomerCnt 被拒收消息的客户数
	MsgRejectedCustomerCnt string `json:"msg_rejected_customer_cnt"`
}

// GetServicerStatistic 获取「客户数据统计」接待人员明细数据
// 通过此接口，可获取接入人工会话数、咨询会话数等与接待人员相关的统计信息
//
// see https://developer.work.weixin.qq.com/document/path/95490
func (c *Client) GetServicerStatistic(r *GetServicerStatisticRequest, opts ...any) (out GetServicerStatisticResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/kf/get_servicer_statistic",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetServicerStatisticRequest is request of Client.GetServicerStatistic
type GetServicerStatisticRequest struct {
	// OpenKfid 客服账号ID
	OpenKfid string `json:"open_kfid" validate:"required"`
	// ServicerUserID 接待人员的userid。第三方应用为密文userid，即open_userid
	ServicerUserID string `json:"servicer_userid"`
	// StartTime 起始日期的时间戳，填当天的0时0分0秒。取值范围：昨天至前180天
	StartTime int `json:"start_time" validate:"required"`
	// EndTime 结束日期的时间戳，填当天的0时0分0秒。取值范围：昨天至前180天
	EndTime int `json:"end_time" validate:"required"`
}

// GetServicerStatisticResponse is response of Client.GetServicerStatistic
type GetServicerStatisticResponse struct {
	// StatisticList 统计数据列表
	StatisticList []map[string]any `json:"statistic_list"`
	// StatTime 数据统计日期，为当日0点的时间戳
	StatTime string `json:"stat_time"`
	// SessionCnt 接入人工会话数。客户发过消息并分配给接待人员的咨询会话数
	SessionCnt string `json:"session_cnt"`
	// CustomerCnt 咨询客户数。在会话中发送过消息且接入了人工会话的客户数量，若客户多次咨询只计算一个客户
	CustomerCnt string `json:"customer_cnt"`
	// CustomerMsgCnt 咨询消息总数。客户在会话中发送的消息的数量
	CustomerMsgCnt string `json:"customer_msg_cnt"`
	// ReplyRate 人工回复率。一个自然日内，客户给接待人员发消息的会话中，接待人员回复了的会话的占比
	ReplyRate float64 `json:"reply_rate"`
	// FirstReplyAverageSec 平均首次响应时长，单位：秒
	FirstReplyAverageSec float64 `json:"first_reply_average_sec"`
	// SatisfactionInvestgateCnt 满意度评价发送数
	SatisfactionInvestgateCnt string `json:"satisfaction_investgate_cnt"`
	// SatisfactionParticipationRate 满意度参评率
	SatisfactionParticipationRate float64 `json:"satisfaction_participation_rate"`
	// SatisfiedRate 满意评价占比
	SatisfiedRate float64 `json:"satisfied_rate"`
	// MiddlingRate 一般评价占比
	MiddlingRate float64 `json:"middling_rate"`
	// DissatisfiedRate 不满意评价占比
	DissatisfiedRate float64 `json:"dissatisfied_rate"`
	// UpgradeServiceCustomerCnt 升级服务客户数
	UpgradeServiceCustomerCnt string `json:"upgrade_service_customer_cnt"`
	// UpgradeServiceMemberInviteCnt 专员服务邀请数
	UpgradeServiceMemberInviteCnt string `json:"upgrade_service_member_invite_cnt"`
	// UpgradeServiceMemberCustomerCnt 添加专员的客户数
	UpgradeServiceMemberCustomerCnt string `json:"upgrade_service_member_customer_cnt"`
	// UpgradeServiceGroupchatInviteCnt 客户群服务邀请数
	UpgradeServiceGroupchatInviteCnt string `json:"upgrade_service_groupchat_invite_cnt"`
	// UpgradeServiceGroupchatCustomerCnt 加入客户群的客户数
	UpgradeServiceGroupchatCustomerCnt string `json:"upgrade_service_groupchat_customer_cnt"`
	// MsgRejectedCustomerCnt 被拒收消息的客户数
	MsgRejectedCustomerCnt string `json:"msg_rejected_customer_cnt"`
}

// AddKnowledgeGroup 添加知识库分组
// 通过此接口创建新的知识库分组。
//
// see https://developer.work.weixin.qq.com/document/path/95971
func (c *Client) AddKnowledgeGroup(r *AddKnowledgeGroupRequest, opts ...any) (out AddKnowledgeGroupResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/kf/knowledge/add_group",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddKnowledgeGroupRequest is request of Client.AddKnowledgeGroup
type AddKnowledgeGroupRequest struct {
	// Name 分组名。不超过12个字
	Name string `json:"name" validate:"required"`
}

// AddKnowledgeGroupResponse is response of Client.AddKnowledgeGroup
type AddKnowledgeGroupResponse struct {
	// GroupID 分组ID
	GroupID string `json:"group_id"`
}

// AddKnowledgeIntent 添加问答
// 创建新的知识库问答
//
// see https://developer.work.weixin.qq.com/document/path/95972
func (c *Client) AddKnowledgeIntent(r *AddKnowledgeIntentRequest, opts ...any) (out AddKnowledgeIntentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/kf/knowledge/add_intent",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddKnowledgeIntentRequest is request of Client.AddKnowledgeIntent
type AddKnowledgeIntentRequest struct {
	// GroupID 分组ID
	GroupID string `json:"group_id" validate:"required"`
	// Question 主问题
	Question any `json:"question" validate:"required"`
	// SimilarQuestions 相似问题
	SimilarQuestions any `json:"similar_questions"`
	// Answers 回答列表。目前仅支持1个
	Answers []map[string]any `json:"answers" validate:"required"`
}

// AddKnowledgeIntentResponse is response of Client.AddKnowledgeIntent
type AddKnowledgeIntentResponse struct {
	// IntentID 问答ID
	IntentID string `json:"intent_id"`
}

// SendKfMsg 发送消息
// 当微信客户处于新接入待处理或智能助手接待状态时，调用该接口给用户发送消息。支持文本、图片、语音、视频、文件、图文、小程序、菜单等类型。
//
// see https://developer.work.weixin.qq.com/document/path/94677
func (c *Client) SendKfMsg(r *SendKfMsgRequest, opts ...any) (out SendKfMsgResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/kf/send_msg",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SendKfMsgRequest is request of Client.SendKfMsg
type SendKfMsgRequest struct {
	// Touser 指定接收消息的客户UserID
	Touser string `json:"touser" validate:"required"`
	// OpenKfid 指定发送消息的客服账号ID
	OpenKfid string `json:"open_kfid" validate:"required"`
	// MsgID 指定消息ID，若不指定则系统自动生成。若指定需确保在客服账号内唯一
	MsgID string `json:"msgid"`
	// Msgtype 消息类型 (text, image, voice, video, file, link, miniprogram, msgmenu)
	Msgtype string `json:"msgtype" validate:"required"`
	// Text 文本消息对象，包含content字段
	Text any `json:"text"`
	// Image 图片消息对象，包含media_id字段
	Image any `json:"image"`
	// Voice 语音消息对象，包含media_id字段
	Voice any `json:"voice"`
	// Video 视频消息对象，包含media_id字段
	Video any `json:"video"`
	// File 文件消息对象，包含media_id字段
	File any `json:"file"`
	// Link 图文链接消息对象，包含title, desc, url, thumb_media_id
	Link any `json:"link"`
	// Miniprogram 小程序消息对象，包含appid, title, thumb_media_id, pagepath
	Miniprogram any `json:"miniprogram"`
	// Msgmenu 菜单消息对象，包含head_content, list, tail_content
	Msgmenu any `json:"msgmenu"`
}

// SendKfMsgResponse is response of Client.SendKfMsg
type SendKfMsgResponse struct {
	// MsgID 消息ID。若请求未指定则系统生成，否则原样返回
	MsgID string `json:"msgid"`
}

// SendKfEventMsg 发送欢迎语等事件响应消息
// 根据事件code向用户发送客服欢迎语、排队提示语或会话结束语等消息。
//
// see https://developer.work.weixin.qq.com/document/path/95122
func (c *Client) SendKfEventMsg(r *SendKfEventMsgRequest, opts ...any) (out SendKfEventMsgResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/kf/send_msg_on_event",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SendKfEventMsgRequest is request of Client.SendKfEventMsg
type SendKfEventMsgRequest struct {
	// Code 事件响应消息对应的code。通过事件回调下发，仅可使用一次。
	Code string `json:"code" validate:"required"`
	// MsgID 消息ID。如果请求参数指定了msgid，则原样返回，否则系统自动生成并返回。不多于32字节。
	MsgID string `json:"msgid"`
	// Msgtype 消息类型。对不同的msgtype，有相应的结构描述。
	Msgtype string `json:"msgtype" validate:"required"`
}

// SendKfEventMsgResponse is response of Client.SendKfEventMsg
type SendKfEventMsgResponse struct {
	// MsgID 消息ID
	MsgID string `json:"msgid"`
}

// TransferServiceState 分配客服会话
// 从微信用户发起咨询到会话结束，企业或第三方可使用API获取和变更会话状态，以实现对会话的分配管理。本文档包含获取会话状态和变更会话状态两个接口。
//
// see https://developer.work.weixin.qq.com/document/path/94669
func (c *Client) TransferServiceState(r *TransferServiceStateRequest, opts ...any) (out TransferServiceStateResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/kf/service_state/trans",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// TransferServiceStateRequest is request of Client.TransferServiceState
type TransferServiceStateRequest struct {
	// OpenKfid 客服账号ID
	OpenKfid string `json:"open_kfid" validate:"required"`
	// ExternalUserID 微信客户的external_userid
	ExternalUserID string `json:"external_userid" validate:"required"`
	// ServiceState 变更的目标状态，状态定义和所允许的变更可参考概述中的流程图和表格
	ServiceState int `json:"service_state" validate:"required"`
	// ServicerUserID 接待人员的userid。第三方应用填密文userid，即open_userid。当state=3时要求必填，接待人员须处于“正在接待”中
	ServicerUserID string `json:"servicer_userid"`
}

// TransferServiceStateResponse is response of Client.TransferServiceState
type TransferServiceStateResponse struct {
	// MsgCode 用于发送响应事件消息的code，将会话初次变更为service_state为2和3时，返回回复语code，service_state为4时，返回结束语code
	MsgCode string `json:"msg_code"`
}

// AddServicer 添加接待人员
// 添加指定客服账号的接待人员，每个客服账号目前最多可添加2000个接待人员，20个部门。
//
// see https://developer.work.weixin.qq.com/document/path/94722
func (c *Client) AddServicer(r *AddServicerRequest, opts ...any) (out AddServicerResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/kf/servicer/add",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddServicerRequest is request of Client.AddServicer
type AddServicerRequest struct {
	// OpenKfid 客服账号ID
	OpenKfid string `json:"open_kfid" validate:"required"`
	// UserIDList 接待人员userid列表。第三方应用填密文userid，即open_userid。可填充个数：0 ~ 100。超过100个需分批调用。
	UserIDList []string `json:"userid_list"`
	// DepartmentIDList 接待人员部门id列表。可填充个数：0 ~ 20。userid_list和department_id_list至少需要填其中一个。
	DepartmentIDList []int `json:"department_id_list"`
}

// AddServicerResponse is response of Client.AddServicer
type AddServicerResponse struct {
	// ResultList 操作结果列表
	ResultList []map[string]any `json:"result_list"`
}

// DeleteServiceUser 删除接待人员
// 从客服账号删除接待人员
//
// see https://developer.work.weixin.qq.com/document/path/94723
func (c *Client) DeleteServiceUser(r *DeleteServiceUserRequest, opts ...any) (out DeleteServiceUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/kf/servicer/del",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteServiceUserRequest is request of Client.DeleteServiceUser
type DeleteServiceUserRequest struct {
	// OpenKfid 客服账号ID
	OpenKfid string `json:"open_kfid" validate:"required"`
	// UserIDList 接待人员userid列表。第三方应用填密文userid，即open_userid；可填充个数：0 ~ 100。超过100个需分批调用。
	UserIDList []string `json:"userid_list"`
	// DepartmentIDList 接待人员部门id列表；可填充个数：0 ~ 100。超过100个需分批调用。
	DepartmentIDList []int `json:"department_id_list"`
}

// DeleteServiceUserResponse is response of Client.DeleteServiceUser
type DeleteServiceUserResponse struct {
	// ResultList 操作结果列表
	ResultList []map[string]any `json:"result_list"`
}

// ListServicer 获取接待人员列表
// 获取某个客服账号的接待人员列表
//
// see https://developer.work.weixin.qq.com/document/path/94724
func (c *Client) ListServicer(r *ListServicerRequest, opts ...any) (out ListServicerResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/kf/servicer/list",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListServicerRequest is request of Client.ListServicer
type ListServicerRequest struct {
	// OpenKfid 客服账号ID
	OpenKfid string `json:"open_kfid" validate:"required"`
}

// ListServicerResponse is response of Client.ListServicer
type ListServicerResponse struct {
	// ServicerList 客服账号的接待人员列表
	ServicerList []map[string]any `json:"servicer_list"`
}

// SyncMsg 拉取会话消息
// 获取微信客服相关的消息列表，支持增量拉取最近3天内的消息
//
// see https://developer.work.weixin.qq.com/document/path/94670
func (c *Client) SyncMsg(r *SyncMsgRequest, opts ...any) (out SyncMsgResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/kf/sync_msg",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SyncMsgRequest is request of Client.SyncMsg
type SyncMsgRequest struct {
	// Cursor 上一次调用时返回的next_cursor，第一次拉取可以不填。若不填，从3天内最早的消息开始返回。不多于64字节
	Cursor string `json:"cursor"`
	// Token 回调事件返回的token字段，10分钟内有效；可不填，如果不填接口有严格的频率限制。不多于128字节
	Token string `json:"token"`
	// Limit 期望请求的数据量，默认值和最大值都为1000
	Limit int `json:"limit"`
	// VoiceFormat 语音消息类型，0-Amr 1-Silk，默认0。可通过该参数控制返回的语音格式
	VoiceFormat int `json:"voice_format"`
	// OpenKfid 指定拉取某个客服账号的消息
	OpenKfid string `json:"open_kfid" validate:"required"`
}

// SyncMsgResponse is response of Client.SyncMsg
type SyncMsgResponse struct {
	// NextCursor 下次调用带上该值，则从当前的位置继续往后拉，以实现增量拉取
	NextCursor string `json:"next_cursor"`
	// HasMore 是否还有更多数据。0-否；1-是
	HasMore int `json:"has_more"`
	// MsgList 消息列表
	MsgList []map[string]any `json:"msg_list"`
}

