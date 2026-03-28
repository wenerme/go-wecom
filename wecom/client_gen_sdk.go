package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// GetUserInfoByCode 通过code获取用户信息
// 通过成员授权获取到的code换取access_token并获取用户ID
//
// see https://developer.work.weixin.qq.com/document/path/101021
func (c *Client) GetUserInfoByCode(r *GetUserInfoByCodeRequest, opts ...any) (out GetUserInfoByCodeResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/user/getuserinfo",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetUserInfoByCodeRequest is request of Client.GetUserInfoByCode
type GetUserInfoByCodeRequest struct {
	// Code 通过成员授权获取到的code，每次成员授权带上的code将不一样，code只能使用一次，5分钟未被使用自动过期
	Code string `json:"code" validate:"required"`
}

// GetUserInfoByCodeResponse is response of Client.GetUserInfoByCode
type GetUserInfoByCodeResponse struct {
	// UserId 成员UserID
	UserId string `json:"UserId"`
}

