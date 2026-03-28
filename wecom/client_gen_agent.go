package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// SetAgent 设置应用
// 设置企业应用的配置信息
//
// see https://developer.work.weixin.qq.com/document/path/90053
func (c *Client) SetAgent(r *SetAgentRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/agent/set",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetAgentRequest is request of Client.SetAgent
type SetAgentRequest struct {
	// AgentID 企业应用的id
	AgentID int `json:"agentid" validate:"required"`
	// ReportLocationFlag 企业应用是否打开地理位置上报 0：不上报；1：进入会话上报
	ReportLocationFlag int `json:"report_location_flag"`
	// LogoMediaID 企业应用头像的mediaid，通过素材管理接口上传图片获得mediaid，上传后会自动裁剪成方形和圆形两个头像
	LogoMediaID string `json:"logo_mediaid"`
	// Name 企业应用名称，长度不超过32个utf8字符
	Name string `json:"name"`
	// Description 企业应用详情，长度为4至120个utf8字符
	Description string `json:"description"`
	// RedirectDomain 企业应用可信域名。注意：域名需通过所有权校验，否则jssdk功能将受限，此时返回错误码85005
	RedirectDomain string `json:"redirect_domain"`
	// Isreportenter 是否上报用户进入应用事件。0：不接收；1：接收
	Isreportenter int `json:"isreportenter"`
	// HomeURL 应用主页url。url必须以http或者https开头（为了提高安全性，建议使用https）
	HomeURL string `json:"home_url"`
}

// CreateMenu 创建菜单
// 创建自定义菜单，支持多种类型的按钮交互
//
// see https://developer.work.weixin.qq.com/document/path/90055
func (c *Client) CreateMenu(r *CreateMenuRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/menu/create",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateMenuRequest is request of Client.CreateMenu
type CreateMenuRequest struct {
	// AgentID 企业应用的id，整型
	AgentID int `json:"agentid" validate:"required"`
	// Button 一级菜单数组，个数应为1~3个
	Button []map[string]any `json:"button" validate:"required"`
	// SubButton 二级菜单数组，个数应为1~5个
	SubButton []any `json:"sub_button"`
	// Type 菜单的响应动作类型
	Type string `json:"type" validate:"required"`
	// Name 菜单的名字。不能为空，主菜单不能超过16字节，子菜单不能超过40字节
	Name string `json:"name" validate:"required"`
	// Key 菜单KEY值，用于消息接口推送，不超过128字节（click等点击类型必须）
	Key string `json:"key"`
	// URL 网页链接，成员点击菜单可打开链接，不超过1024字节（view类型必须）
	URL string `json:"url"`
	// Pagepath 小程序的页面路径（view_miniprogram类型必须）
	Pagepath string `json:"pagepath"`
	// AppID 小程序的appid（仅与企业绑定的小程序可配置，view_miniprogram类型必须）
	AppID string `json:"appid"`
}

// DeleteMenu 删除菜单
// 删除企业自定义菜单
//
// see https://developer.work.weixin.qq.com/document/path/90057
func (c *Client) DeleteMenu(r *DeleteMenuRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/menu/delete",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteMenuRequest is request of Client.DeleteMenu
type DeleteMenuRequest struct {
	// AgentID 应用id
	AgentID int `json:"agentid" validate:"required"`
}

// GetMenu 获取菜单
// 获取当前配置的第三方或自建应用的菜单信息
//
// see https://developer.work.weixin.qq.com/document/path/90056
func (c *Client) GetMenu(r *GetMenuRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/menu/get",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMenuRequest is request of Client.GetMenu
type GetMenuRequest struct {
	// AgentID 应用id
	AgentID int `json:"agentid" validate:"required"`
}

