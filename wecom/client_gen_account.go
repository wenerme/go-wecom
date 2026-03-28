package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// BatchOpenUserIdToUserId userid转换
// 将代开发应用或第三方应用获取的密文open_userid转换为明文userid。
//
// see https://developer.work.weixin.qq.com/document/path/95884
func (c *Client) BatchOpenUserIdToUserId(r *BatchOpenUserIdToUserIdRequest, opts ...any) (out BatchOpenUserIdToUserIdResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/batch/openuserid_to_userid",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchOpenUserIdToUserIdRequest is request of Client.BatchOpenUserIdToUserId
type BatchOpenUserIdToUserIdRequest struct {
	// OpenUserIDList open_userid列表，最多不超过1000个。必须是source_agentid对应的应用所获取
	OpenUserIDList []string `json:"open_userid_list" validate:"required"`
	// SourceAgentID 企业授权的代开发自建应用或第三方应用的agentid
	SourceAgentID int `json:"source_agentid" validate:"required"`
}

// BatchOpenUserIdToUserIdResponse is response of Client.BatchOpenUserIdToUserId
type BatchOpenUserIdToUserIdResponse struct {
	// UserIDList 转换成功的列表，包含open_userid和对应的userid
	UserIDList []map[string]any `json:"userid_list"`
	// InvalidOpenUserIDList 不合法的open_userid列表
	InvalidOpenUserIDList []string `json:"invalid_open_userid_list"`
}

// ConvertTmpExternalUserid tmp_external_userid的转换
// 将应用获取的外部用户临时id转换为external_userid
//
// see https://developer.work.weixin.qq.com/document/path/98729
func (c *Client) ConvertTmpExternalUserid(r *ConvertTmpExternalUseridRequest, opts ...any) (out ConvertTmpExternalUseridResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/idconvert/convert_tmp_external_userid",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ConvertTmpExternalUseridRequest is request of Client.ConvertTmpExternalUserid
type ConvertTmpExternalUseridRequest struct {
	// BusinessType 业务类型。1-会议 2-收集表 3-智能表
	BusinessType int `json:"business_type" validate:"required"`
	// UserType 转换的目标用户类型。1-客户 2-企业互联 3-上下游 4-互联企业（圈子）
	UserType int `json:"user_type" validate:"required"`
	// TmpExternalUserIDList 外部用户临时id列表，最多不超过100个
	TmpExternalUserIDList []string `json:"tmp_external_userid_list" validate:"required"`
}

// ConvertTmpExternalUseridResponse is response of Client.ConvertTmpExternalUserid
type ConvertTmpExternalUseridResponse struct {
	// Results 转换成功的结果列表
	Results []map[string]any `json:"results"`
	// InvalidTmpExternalUserIDList 无法转换的tmp_external_userid。可能非法或没有权限
	InvalidTmpExternalUserIDList []string `json:"invalid_tmp_external_userid_list"`
}
