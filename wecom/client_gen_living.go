package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// CancelLiving 取消预约直播
// 取消预约状态下的直播
//
// see https://developer.work.weixin.qq.com/document/path/93789
func (c *Client) CancelLiving(r *CancelLivingRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/living/cancel",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CancelLivingRequest is request of Client.CancelLiving
type CancelLivingRequest struct {
	// Livingid 直播id，仅允许取消预约状态下的直播id
	Livingid string `json:"livingid" validate:"required"`
}

// CreateLiving 创建预约直播
// 创建企业预约直播
//
// see https://developer.work.weixin.qq.com/document/path/93788
func (c *Client) CreateLiving(r *CreateLivingRequest, opts ...any) (out CreateLivingResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/living/create",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateLivingRequest is request of Client.CreateLiving
type CreateLivingRequest struct {
	// AnchorUserID 直播发起者的userid
	AnchorUserID string `json:"anchor_userid" validate:"required"`
	// Theme 直播的标题，最多支持20个utf8字符
	Theme string `json:"theme" validate:"required"`
	// LivingStart 直播开始时间的unix时间戳
	LivingStart int `json:"living_start" validate:"required"`
	// LivingDuration 直播持续时长
	LivingDuration int `json:"living_duration" validate:"required"`
	// Description 直播的简介，最多支持100个utf8字符
	Description string `json:"description"`
	// Type 直播的类型，0：通用直播，1：小班课，2：大班课，3：企业培训，4：活动直播
	Type int `json:"type"`
	// RemindTime 指定直播开始前多久提醒用户，相对于living_start前的秒数
	RemindTime int `json:"remind_time"`
	// ActivityCoverMediaID 活动直播特定参数，直播间封面图的mediaId
	ActivityCoverMediaID string `json:"activity_cover_mediaid"`
	// ActivityShareMediaID 活动直播特定参数，直播分享卡片图的mediaId
	ActivityShareMediaID string `json:"activity_share_mediaid"`
	// ActivityDetail 活动直播特定参数，活动直播详情信息
	ActivityDetail any `json:"activity_detail"`
}

// CreateLivingResponse is response of Client.CreateLiving
type CreateLivingResponse struct {
	// Livingid 直播id，通过此id可调用进入直播接口
	Livingid string `json:"livingid"`
}

// GetLivingCode 获取微信观看直播凭证
// 通过微信观看直播的凭证，可在微信中H5或小程序页面唤起企业微信直播小程序，并进入对应直播或直播回放。
//
// see https://developer.work.weixin.qq.com/document/path/93838
func (c *Client) GetLivingCode(r *GetLivingCodeRequest, opts ...any) (out GetLivingCodeResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/living/get_living_code",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetLivingCodeRequest is request of Client.GetLivingCode
type GetLivingCodeRequest struct {
	// Livingid 直播id
	Livingid string `json:"livingid" validate:"required"`
	// OpenID 微信用户的openid
	OpenID string `json:"openid" validate:"required"`
}

// GetLivingCodeResponse is response of Client.GetLivingCode
type GetLivingCodeResponse struct {
	// LivingCode 微信观看直播凭证，5分钟内可以重复使用，且仅能在微信上使用
	LivingCode string `json:"living_code"`
}

// GetLivingShareInfo 获取跳转小程序商城的直播观众信息
// 通过此接口，开发者可获取跳转小程序商城的直播间(“推广产品”直播)观众id、邀请人id及对应直播间id，以打通卖货直播的“人货场”信息闭环。
//
// see https://developer.work.weixin.qq.com/document/path/94487
func (c *Client) GetLivingShareInfo(r *GetLivingShareInfoRequest, opts ...any) (out GetLivingShareInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/living/get_living_share_info",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetLivingShareInfoRequest is request of Client.GetLivingShareInfo
type GetLivingShareInfoRequest struct {
	// WwShareCode “推广产品”直播观众跳转小程序商城时会在小程序path中带上ww_share_code=xxxxx参数，ww_share_code五分钟内有效
	WwShareCode string `json:"ww_share_code" validate:"required"`
}

// GetLivingShareInfoResponse is response of Client.GetLivingShareInfo
type GetLivingShareInfoResponse struct {
	// Livingid 直播id
	Livingid string `json:"livingid"`
	// ViewerUserID 观众的userid，观众为企业内部成员时返回
	ViewerUserID string `json:"viewer_userid"`
	// ViewerExternalUserID 观众的external_userid，观众为非企业内部成员时返回
	ViewerExternalUserID string `json:"viewer_external_userid"`
	// InvitorUserID 邀请人的userid，邀请人为企业内部成员时返回
	InvitorUserID string `json:"invitor_userid"`
	// InvitorExternalUserID 邀请人的external_userid，邀请人为非企业内部成员时返回
	InvitorExternalUserID string `json:"invitor_external_userid"`
}

// ModifyLiving 修改预约直播
// 修改已创建的预约直播信息，仅允许修改预约状态下的直播id。
//
// see https://developer.work.weixin.qq.com/document/path/93790
func (c *Client) ModifyLiving(r *ModifyLivingRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/living/modify",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ModifyLivingRequest is request of Client.ModifyLiving
type ModifyLivingRequest struct {
	// Livingid 直播id，仅允许修改预约状态下的直播id
	Livingid string `json:"livingid" validate:"required"`
	// Theme 直播的标题，最多支持60个字节
	Theme string `json:"theme"`
	// LivingStart 直播开始时间的unix时间戳
	LivingStart int `json:"living_start"`
	// LivingDuration 直播持续时长
	LivingDuration int `json:"living_duration"`
	// Type 直播的类型，0：通用直播，1：小班课，2：大班课，3：企业培训，4：活动直播。其中大班课和小班课仅k12学校和IT行业类型能够发起
	Type int `json:"type"`
	// Description 直播的简介，最多支持300个字节
	Description string `json:"description"`
	// RemindTime 指定直播开始前多久提醒用户，相对于living_start前的秒数，默认为0
	RemindTime int `json:"remind_time"`
}

