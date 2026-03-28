package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// GetChainList 获取上下游信息
// 获取企业上下游列表、通讯录分组及企业信息等接口集合
//
// see https://developer.work.weixin.qq.com/document/path/95820
func (c *Client) GetChainList(r *GetChainListRequest, opts ...any) (out GetChainListResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/corpgroup/corp/get_chain_list",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetChainListRequest is request of Client.GetChainList
type GetChainListRequest struct {}

// GetChainListResponse is response of Client.GetChainList
type GetChainListResponse struct {
	// Chains 企业上下游列表
	Chains []map[string]any `json:"chains"`
}

// GetChainUserCustomId 查询成员自定义id
// 上级企业自建应用/代开发应用通过本接口查询成员自定义 id
//
// see https://developer.work.weixin.qq.com/document/path/97441
func (c *Client) GetChainUserCustomId(r *GetChainUserCustomIdRequest, opts ...any) (out GetChainUserCustomIdResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/corpgroup/corp/get_chain_user_custom_id",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetChainUserCustomIdRequest is request of Client.GetChainUserCustomId
type GetChainUserCustomIdRequest struct {
	// ChainID 上下游id
	ChainID string `json:"chain_id" validate:"required"`
	// CorpID 已加入企业id
	CorpID string `json:"corpid" validate:"required"`
	// UserID 企业内的成员
	UserID string `json:"userid" validate:"required"`
}

// GetChainUserCustomIdResponse is response of Client.GetChainUserCustomId
type GetChainUserCustomIdResponse struct {
	// UserCustomID 成员自定义 id
	UserCustomID string `json:"user_custom_id"`
}

// ListCorpGroupAppShare 获取应用共享信息
// 局校互联中的局端或者上下游中的上游企业通过该接口可以获取某个应用分享给的所有企业列表。
//
// see https://developer.work.weixin.qq.com/document/path/95813
func (c *Client) ListCorpGroupAppShare(r *ListCorpGroupAppShareRequest, opts ...any) (out ListCorpGroupAppShareResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/corpgroup/corp/list_app_share_info",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListCorpGroupAppShareRequest is request of Client.ListCorpGroupAppShare
type ListCorpGroupAppShareRequest struct {
	// BusinessType 填0则为企业互联/局校互联，填1则表示上下游企业
	BusinessType int `json:"business_type"`
	// AgentID 上级/上游企业应用agentid
	AgentID int `json:"agentid" validate:"required"`
	// CorpID 下级/下游企业corpid，若指定该参数则表示拉取该下级/下游企业的应用共享信息
	CorpID string `json:"corpid"`
	// Limit 返回的最大记录数，整型，最大值100，默认情况或者值为0表示下拉取全量数据
	Limit int `json:"limit"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回
	Cursor string `json:"cursor"`
}

// ListCorpGroupAppShareResponse is response of Client.ListCorpGroupAppShare
type ListCorpGroupAppShareResponse struct {
	// Ending 1表示拉取完毕，0表示数据没有拉取完
	Ending int `json:"ending"`
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录
	NextCursor string `json:"next_cursor"`
	// CorpList 应用共享信息列表
	CorpList []map[string]any `json:"corp_list"`
}

// RemoveCorpFromGroup 移除企业
// 上级/上游企业通过该接口移除下游企业
//
// see https://developer.work.weixin.qq.com/document/path/95822
func (c *Client) RemoveCorpFromGroup(r *RemoveCorpFromGroupRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/corpgroup/corp/remove_corp",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// RemoveCorpFromGroupRequest is request of Client.RemoveCorpFromGroup
type RemoveCorpFromGroupRequest struct {
	// ChainID 上下游id
	ChainID string `json:"chain_id" validate:"required"`
	// CorpID 需要移除的下游企业corpid
	CorpID string `json:"corpid"`
	// PendingCorpID 需要移除的未加入下游企业corpid，corpid和pending_corpid至少填一个，都填corpid生效
	PendingCorpID string `json:"pending_corpid"`
}

// GetCorpSharedChainList 获取下级企业加入的上下游
// 上级企业自建应用/代开发应用通过本接口查询下级企业所在上下游
//
// see https://developer.work.weixin.qq.com/document/path/97442
func (c *Client) GetCorpSharedChainList(r *GetCorpSharedChainListRequest, opts ...any) (out GetCorpSharedChainListResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/corpgroup/get_corp_shared_chain_list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCorpSharedChainListRequest is request of Client.GetCorpSharedChainList
type GetCorpSharedChainListRequest struct {
	// CorpID 已加入企业id
	CorpID string `json:"corpid"`
}

// GetCorpSharedChainListResponse is response of Client.GetCorpSharedChainList
type GetCorpSharedChainListResponse struct {
	// Chains 上下游列表
	Chains []map[string]any `json:"chains"`
}

// GetAsyncJobResult 获取异步任务结果
// 查询企业微信异步任务的执行结果，支持查询历史任务状态及详细处理信息
//
// see https://developer.work.weixin.qq.com/document/path/95823
func (c *Client) GetAsyncJobResult(r *GetAsyncJobResultRequest, opts ...any) (out GetAsyncJobResultResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/corpgroup/getresult",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetAsyncJobResultRequest is request of Client.GetAsyncJobResult
type GetAsyncJobResultRequest struct {
	// JobID 异步任务id，最大长度为64字节
	JobID string `json:"jobid" validate:"required"`
}

// GetAsyncJobResultResponse is response of Client.GetAsyncJobResult
type GetAsyncJobResultResponse struct {
	// Status 任务状态，1表示任务开始，2表示任务进行中，3表示任务已完成
	Status int `json:"status"`
	// Result 详细的处理结果。当任务完成后此字段有效
	Result any `json:"result"`
}

// ImportChainContact 批量导入上下游联系人
// 提交批量导入上下游联系人任务
//
// see https://developer.work.weixin.qq.com/document/path/95821
func (c *Client) ImportChainContact(r *ImportChainContactRequest, opts ...any) (out ImportChainContactResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/corpgroup/import_chain_contact",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ImportChainContactRequest is request of Client.ImportChainContact
type ImportChainContactRequest struct {
	// ChainID 上下游id。文件中的联系人将会被导入此上下游中
	ChainID string `json:"chain_id" validate:"required"`
	// ContactList 上下游联系人列表。这些联系人将会被导入此上下游中
	ContactList []map[string]any `json:"contact_list" validate:"required"`
}

// ImportChainContactResponse is response of Client.ImportChainContact
type ImportChainContactResponse struct {
	// JobID 异步任务id，最大长度为64字节
	JobID string `json:"jobid"`
}

// AddCorpGroupRule 新增对接规则
// 上下游系统应用可通过该接口新增一条对接规则。注意：新增和更新上下游对接规则的接口每天最多调用1000次。
//
// see https://developer.work.weixin.qq.com/document/path/95664
func (c *Client) AddCorpGroupRule(r *AddCorpGroupRuleRequest, opts ...any) (out AddCorpGroupRuleResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/corpgroup/rule/add_rule",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddCorpGroupRuleRequest is request of Client.AddCorpGroupRule
type AddCorpGroupRuleRequest struct {
	// ChainID 上下游id
	ChainID string `json:"chain_id" validate:"required"`
	// RuleInfo 上下游关系规则的详情
	RuleInfo any `json:"rule_info" validate:"required"`
}

// AddCorpGroupRuleResponse is response of Client.AddCorpGroupRule
type AddCorpGroupRuleResponse struct {
	// RuleID 上下游规则id
	RuleID int `json:"rule_id"`
}

// DeleteCorpGroupRule 删除对接规则
// 上下游系统应用可通过该接口删除企业上下游规则
//
// see https://developer.work.weixin.qq.com/document/path/95663
func (c *Client) DeleteCorpGroupRule(r *DeleteCorpGroupRuleRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/corpgroup/rule/delete_rule",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteCorpGroupRuleRequest is request of Client.DeleteCorpGroupRule
type DeleteCorpGroupRuleRequest struct {
	// ChainID 上下游id
	ChainID string `json:"chain_id" validate:"required"`
	// RuleID 上下游规则id
	RuleID int `json:"rule_id" validate:"required"`
}

// GetRuleInfo 获取对接规则详情
// 上下游系统应用可通过该接口获取企业上下游规则详情
//
// see https://developer.work.weixin.qq.com/document/path/95667
func (c *Client) GetRuleInfo(r *GetRuleInfoRequest, opts ...any) (out GetRuleInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/corpgroup/rule/get_rule_info",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetRuleInfoRequest is request of Client.GetRuleInfo
type GetRuleInfoRequest struct {
	// ChainID 上下游id
	ChainID string `json:"chain_id" validate:"required"`
	// RuleID 上下游规则id
	RuleID int `json:"rule_id" validate:"required"`
}

// GetRuleInfoResponse is response of Client.GetRuleInfo
type GetRuleInfoResponse struct {
	// RuleInfo 上下游关系规则的详情
	RuleInfo any `json:"rule_info"`
}

// ListRuleId 获取对接规则id列表
// 上下游系统应用可通过该接口获取企业上下游规则id列表
//
// see https://developer.work.weixin.qq.com/document/path/95631
func (c *Client) ListRuleId(r *ListRuleIdRequest, opts ...any) (out ListRuleIdResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/corpgroup/rule/list_ids",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListRuleIdRequest is request of Client.ListRuleId
type ListRuleIdRequest struct {
	// ChainID 上下游id
	ChainID string `json:"chain_id" validate:"required"`
}

// ListRuleIdResponse is response of Client.ListRuleId
type ListRuleIdResponse struct {
	// RuleIds 上下游关系规则的id
	RuleIds []int `json:"rule_ids"`
}

// ModifyRule 更新对接规则
// 上下游系统应用可通过该接口修改一条对接规则。
//
// see https://developer.work.weixin.qq.com/document/path/95666
func (c *Client) ModifyRule(r *ModifyRuleRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/corpgroup/rule/modify_rule",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ModifyRuleRequest is request of Client.ModifyRule
type ModifyRuleRequest struct {
	// ChainID 上下游id
	ChainID string `json:"chain_id" validate:"required"`
	// RuleID 上下游规则id
	RuleID int `json:"rule_id" validate:"required"`
	// RuleInfo 上下游关系规则的详情
	RuleInfo any `json:"rule_info" validate:"required"`
}

// LinkUnionidToExternalUserid 上下游关联客户信息-已添加客户
// 通过unionid和openid查询external_userid，用于品牌与经销商之间打通微信客户数据。
//
// see https://developer.work.weixin.qq.com/document/path/95818
func (c *Client) LinkUnionidToExternalUserid(r *LinkUnionidToExternalUseridRequest, opts ...any) (out LinkUnionidToExternalUseridResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/corpgroup/unionid_to_external_userid",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// LinkUnionidToExternalUseridRequest is request of Client.LinkUnionidToExternalUserid
type LinkUnionidToExternalUseridRequest struct {
	// UnionID 微信客户的unionid
	UnionID string `json:"unionid" validate:"required"`
	// OpenID 微信客户的openid
	OpenID string `json:"openid" validate:"required"`
	// CorpID 需要换取的企业corpid，不填则拉取所有企业
	CorpID string `json:"corpid"`
	// MassCallTicket 大批量调用凭据，适用于数据初始化场景
	MassCallTicket string `json:"mass_call_ticket"`
}

// LinkUnionidToExternalUseridResponse is response of Client.LinkUnionidToExternalUserid
type LinkUnionidToExternalUseridResponse struct {
	// ExternalUserIDInfo 该unionid对应的外部联系人信息列表
	ExternalUserIDInfo []map[string]any `json:"external_userid_info"`
}

// UnionidToPendingId unionid查询pending_id
// 上游企业通过此接口将微信unionid转为pending_id，用于关联微信用户与外部联系人
//
// see https://developer.work.weixin.qq.com/document/path/98039
func (c *Client) UnionidToPendingId(r *UnionidToPendingIdRequest, opts ...any) (out UnionidToPendingIdResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/corpgroup/unionid_to_pending_id",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UnionidToPendingIdRequest is request of Client.UnionidToPendingId
type UnionidToPendingIdRequest struct {
	// UnionID 微信客户的unionid
	UnionID string `json:"unionid" validate:"required"`
	// OpenID 微信客户的openid
	OpenID string `json:"openid" validate:"required"`
}

// UnionidToPendingIdResponse is response of Client.UnionidToPendingId
type UnionidToPendingIdResponse struct {
	// PendingID unionid和openid对应的pending_id
	PendingID string `json:"pending_id"`
}

