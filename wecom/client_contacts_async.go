package wecom

import "github.com/wenerme/go-req"

// BatchSyncUser 增量更新成员
// 本接口以userid（帐号）为主键，增量更新企业微信通讯录成员。请先下载CSV模板(下载增量更新成员模版)，根据需求填写文件内容。
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90980
func (c *Client) BatchSyncUser(r *BatchSyncUserRequest) (out BatchSyncUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/batch/syncuser",
		Body:   r,
	}).Fetch(&out)
	return
}

// BatchReplaceUser 全量覆盖成员
// 本接口以userid为主键，全量覆盖企业的通讯录成员，任务完成后企业的通讯录成员与提交的文件完全保持一致。请先下载CSV文件(下载全量覆盖成员模版)，根据需求填写文件内容。
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90981
func (c *Client) BatchReplaceUser(r *BatchReplaceUserRequest) (out BatchReplaceUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/batch/replaceuser",
		Body:   r,
	}).Fetch(&out)
	return
}

// BatchReplaceParty 全量覆盖部门
// 本接口以partyid为键，全量覆盖企业的通讯录组织架构，任务完成后企业的通讯录组织架构与提交的文件完全保持一致。请先下载CSV文件(下载全量覆盖部门模版)，根据需求填写文件内容。
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90982
func (c *Client) BatchReplaceParty(r *BatchReplacePartyRequest) (out BatchReplacePartyResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/batch/replaceparty",
		Body:   r,
	}).Fetch(&out)
	return
}

// BatchGetResult 获取异步任务结果
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90983
func (c *Client) BatchGetResult(r *BatchGetResultRequest) (out BatchGetResultResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "GET",
		URL:    "/cgi-bin/batch/getresult",
		Query:  r,
	}).Fetch(&out)
	return
}

type BatchSyncUserRequest struct {
	// MediaID 上传的csv文件的media_id
	MediaID string `json:"media_id"`
	// ToInvite 是否邀请新建的成员使用企业微信（将通过微信服务通知或短信或邮件下发邀请，每天自动下发一次，最多持续3个工作日），默认值为true。
	ToInvite bool `json:"to_invite"`
	// Callback 回调信息。如填写该项则任务完成后，通过callback推送事件给企业。具体请参考应用回调模式中的相应选项
	Callback BatchSyncCallback `json:"callback"`
}

type BatchSyncUserResponse struct {
	// JobID 异步任务id，最大长度为64字节
	JobID string `json:"jobid"`
}

type BatchReplaceUserRequest struct {
	// MediaID 上传的csv文件的media_id
	MediaID string `json:"media_id"`
	// ToInvite 是否邀请新建的成员使用企业微信（将通过微信服务通知或短信或邮件下发邀请，每天自动下发一次，最多持续3个工作日），默认值为true。
	ToInvite bool `json:"to_invite"`
	// Callback 回调信息。如填写该项则任务完成后，通过callback推送事件给企业。具体请参考应用回调模式中的相应选项
	Callback BatchSyncCallback `json:"callback"`
}

type BatchReplaceUserResponse struct {
	// JobID 异步任务id，最大长度为64字节
	JobID string `json:"jobid"`
}

type BatchReplacePartyRequest struct {
	// MediaID 上传的csv文件的media_id
	MediaID string `json:"media_id"`
	// Callback 回调信息。如填写该项则任务完成后，通过callback推送事件给企业。具体请参考应用回调模式中的相应选项
	Callback BatchSyncCallback `json:"callback"`
}

type BatchReplacePartyResponse struct {
	// JobID 异步任务id，最大长度为64字节
	JobID string `json:"jobid"`
}

type BatchGetResultRequest struct {
	// JobID 异步任务id，最大长度为64字节
	JobID string `json:"jobid"`
}

type BatchGetResultResponse struct {
	// Status 任务状态，整型，1表示任务开始，2表示任务进行中，3表示任务已完成
	Status int `json:"status"`
	// Type 操作类型，字节串，目前分别有：1. sync_user(增量更新成员)  2. replace_user(全量覆盖成员)3. replace_party(全量覆盖部门)
	Type string `json:"type"`
	// Total 任务运行总条数
	Total int `json:"total"`
	// Percentage 目前运行百分比，当任务完成时为100
	Percentage int `json:"percentage"`
	// Result 详细的处理结果，具体格式参考下面说明。当任务完成后此字段有效
	Result []BatchGetResult `json:"result"`
}
