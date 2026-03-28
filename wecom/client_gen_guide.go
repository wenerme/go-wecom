package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// GetApiDomainIp 获取企业微信接口IP段
// 获取企业微信服务器的IP段，用于防火墙配置。建议每天定时拉取更新。
//
// see https://developer.work.weixin.qq.com/document/path/92520
func (c *Client) GetApiDomainIp(r *GetApiDomainIpRequest, opts ...any) (out GetApiDomainIpResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/get_api_domain_ip",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetApiDomainIpRequest is request of Client.GetApiDomainIp
type GetApiDomainIpRequest struct {}

// GetApiDomainIpResponse is response of Client.GetApiDomainIp
type GetApiDomainIpResponse struct {
	// IPList 企业微信服务器IP段
	IPList []string `json:"ip_list"`
}

// GetCallbackIp 获取企业微信服务器IP段
// 获取企业微信在回调时使用的IP网段，用于配置防火墙白名单。
//
// see https://developer.work.weixin.qq.com/document/path/90465
func (c *Client) GetCallbackIp(r *GetCallbackIpRequest, opts ...any) (out GetCallbackIpResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/getcallbackip",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCallbackIpRequest is request of Client.GetCallbackIp
type GetCallbackIpRequest struct {}

// GetCallbackIpResponse is response of Client.GetCallbackIp
type GetCallbackIpResponse struct {
	// IPList 企业微信服务器IP段列表
	IPList []string `json:"ip_list"`
}

