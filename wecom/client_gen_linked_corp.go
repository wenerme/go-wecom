package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// GetCorpGroupToken 获取下级/下游企业的access_token
// 获取应用可见范围内下级/下游企业的access_token，该access_token可用于调用下级/下游企业通讯录的只读接口
//
// see https://developer.work.weixin.qq.com/document/path/95313
func (c *Client) GetCorpGroupToken(r *GetCorpGroupTokenRequest, opts ...any) (out GetCorpGroupTokenResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/corpgroup/corp/gettoken",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCorpGroupTokenRequest is request of Client.GetCorpGroupToken
type GetCorpGroupTokenRequest struct {
	// CorpID 已授权的下级/下游企业corpid
	CorpID string `json:"corpid" validate:"required"`
	// BusinessType 填0则为企业互联/局校互联，填1则表示上下游企业，默认0
	BusinessType int `json:"business_type"`
	// AgentID 已授权的下级/下游企业应用ID
	AgentID int `json:"agentid" validate:"required"`
}

// GetCorpGroupTokenResponse is response of Client.GetCorpGroupToken
type GetCorpGroupTokenResponse struct {
	// AccessToken 获取到的下级/下游企业调用凭证，最长为512字节
	AccessToken string `json:"access_token"`
	// ExpiresIn 凭证的有效时间（秒）
	ExpiresIn int `json:"expires_in"`
}

// TransferMiniProgramSession 获取下级/下游企业小程序session
// 上级/上游企业通过该接口转换为下级/下游企业的小程序session
//
// see https://developer.work.weixin.qq.com/document/path/95317
func (c *Client) TransferMiniProgramSession(r *TransferMiniProgramSessionRequest, opts ...any) (out TransferMiniProgramSessionResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/miniprogram/transfer_session",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// TransferMiniProgramSessionRequest is request of Client.TransferMiniProgramSession
type TransferMiniProgramSessionRequest struct {
	// UserID 通过code2Session接口获取到的加密的userid，不多于64字节
	UserID string `json:"userid" validate:"required"`
	// SessionKey 通过code2Session接口获取到的属于上级/上游企业的会话密钥，不多于64字节
	SessionKey string `json:"session_key" validate:"required"`
}

// TransferMiniProgramSessionResponse is response of Client.TransferMiniProgramSession
type TransferMiniProgramSessionResponse struct {
	// UserID 下级/下游企业用户的ID
	UserID string `json:"userid"`
	// SessionKey 属于下级/下游企业的会话密钥
	SessionKey string `json:"session_key"`
}
