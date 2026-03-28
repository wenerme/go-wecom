package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// GetTfaInfo 获取用户二次验证信息
// 获取用户二次验证信息，用于解锁企业微信终端
//
// see https://developer.work.weixin.qq.com/document/path/99502
func (c *Client) GetTfaInfo(r *GetTfaInfoRequest, opts ...any) (out GetTfaInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/auth/get_tfa_info",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetTfaInfoRequest is request of Client.GetTfaInfo
type GetTfaInfoRequest struct {
	// Code 用户进入二次验证页面时，企业微信颁发的code。每次成员授权带上的code将不一样，code只能使用一次，5分钟未被使用自动过期
	Code string `json:"code" validate:"required"`
}

// GetTfaInfoResponse is response of Client.GetTfaInfo
type GetTfaInfoResponse struct {
	// UserID 成员UserID。若需要获得用户详情信息，可调用通讯录接口：读取成员
	UserID string `json:"userid"`
	// TfaCode 二次验证授权码，开发者可以调用通过二次验证接口，解锁企业微信终端。tfa_code有效期五分钟，且只能使用一次
	TfaCode string `json:"tfa_code"`
}

// GetAuthUserDetail 获取访问用户敏感信息
// 自建应用与代开发应用可通过该接口获取成员授权的敏感字段
//
// see https://developer.work.weixin.qq.com/document/path/95833
func (c *Client) GetAuthUserDetail(r *GetAuthUserDetailRequest, opts ...any) (out GetAuthUserDetailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/auth/getuserdetail",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetAuthUserDetailRequest is request of Client.GetAuthUserDetail
type GetAuthUserDetailRequest struct {
	// UserTicket 成员票据
	UserTicket string `json:"user_ticket" validate:"required"`
}

// GetAuthUserDetailResponse is response of Client.GetAuthUserDetail
type GetAuthUserDetailResponse struct {
	// UserID 成员UserID
	UserID string `json:"userid"`
	// Gender 性别。0表示未定义，1表示男性，2表示女性。仅在用户同意snsapi_privateinfo授权时返回真实值，否则返回0.
	Gender string `json:"gender"`
	// Avatar 头像url。仅在用户同意snsapi_privateinfo授权时返回真实头像，否则返回默认头像
	Avatar string `json:"avatar"`
	// QrCode 员工个人二维码（扫描可添加为外部联系人），仅在用户同意snsapi_privateinfo授权时返回
	QrCode string `json:"qr_code"`
	// Mobile 手机，仅在用户同意snsapi_privateinfo授权时返回，第三方应用不可获取
	Mobile string `json:"mobile"`
	// Email 邮箱，仅在用户同意snsapi_privateinfo授权时返回，第三方应用不可获取
	Email string `json:"email"`
	// BizMail 企业邮箱，仅在用户同意snsapi_privateinfo授权时返回，第三方应用不可获取
	BizMail string `json:"biz_mail"`
	// Address 地址，仅在用户同意snsapi_privateinfo授权时返回，第三方应用不可获取
	Address string `json:"address"`
}

// SubmitTfaSuccess 使用二次验证
// 提交用户二次验证结果
//
// see https://developer.work.weixin.qq.com/document/path/99523
func (c *Client) SubmitTfaSuccess(r *SubmitTfaSuccessRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/user/tfa_succ",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SubmitTfaSuccessRequest is request of Client.SubmitTfaSuccess
type SubmitTfaSuccessRequest struct {
	// UserID 用户的userid
	UserID string `json:"userid" validate:"required"`
	// TfaCode 获取用户二次验证信息接口返回的tfa_code，五分钟内有效且只能使用一次
	TfaCode string `json:"tfa_code" validate:"required"`
}

