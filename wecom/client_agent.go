package wecom

import "github.com/wenerme/go-req"

// GetAgent 获取指定的应用详情
// 对于互联企业的应用，如果需要获取应用可见范围内其他互联企业的部门与成员，请调用互联企业-获取应用可见范围接口
//
// see https://work.weixin.qq.com/api/doc/90001/90143/90363
func (c *Client) GetAgent(r *GetAgentRequest, opts ...interface{}) (out GetAgentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/agent/get",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListAgent 获取access_token对应的应用列表
//
// see https://work.weixin.qq.com/api/doc/90001/90143/90363
func (c *Client) ListAgent(opts ...interface{}) (out ListAgentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/agent/list",
		Options: opts,
	}).Fetch(&out)
	return
}

// SetWorkbenchTemplate 设置应用在工作台展示的模版
// 请求说明：该接口指定应用自定义模版类型。同时也支持设置企业默认模版数据。若type指定为 “normal” 则为取消自定义模式，改为普通展示模式
//
// see https://work.weixin.qq.com/api/doc/90001/90143/94620
func (c *Client) SetWorkbenchTemplate(r *SetWorkbenchTemplateRequest, opts ...interface{}) (out SetWorkbenchTemplateResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/agent/set_workbench_template",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetWorkbenchTemplate 获取应用在工作台展示的模版
//
// see https://work.weixin.qq.com/api/doc/90001/90143/94620
func (c *Client) GetWorkbenchTemplate(r *GetWorkbenchTemplateRequest, opts ...interface{}) (out GetWorkbenchTemplateResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/agent/get_workbench_template",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetWorkbenchData 设置应用在用户工作台展示的数据
//
// see https://work.weixin.qq.com/api/doc/90001/90143/94620
func (c *Client) SetWorkbenchData(r *SetWorkbenchDataRequest, opts ...interface{}) (out SetWorkbenchDataResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/agent/set_workbench_data",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

type GetAgentRequest struct {
	// AgentID 应用id
	AgentID int `json:"agentid"  validate:"required"`
}

type GetAgentResponse struct {
	// AgentID 企业应用id
	AgentID int `json:"agentid"  `
	// Name 企业应用名称
	Name string `json:"name"  `
	// SquareLogoURL 企业应用方形头像
	SquareLogoURL string `json:"square_logo_url"  `
	// Description 企业应用详情
	Description string `json:"description"  `
	// AllowUserInfos 企业应用可见范围（人员），其中包括userid
	AllowUserInfos GetAgentAllowUserInfos `json:"allow_userinfos"  `
	// AllowParties 企业应用可见范围（部门）
	AllowParties GetAgentAllowParties `json:"allow_partys"  `
	// AllowTags 企业应用可见范围（标签）
	AllowTags GetAgentAllowTags `json:"allow_tags"  `
	// Close 企业应用是否被停用
	Close int `json:"close"  `
	// RedirectDomain 企业应用可信域名
	RedirectDomain string `json:"redirect_domain"  `
	// ReportLocationFlag 企业应用是否打开地理位置上报 0：不上报；1：进入会话上报；
	ReportLocationFlag int `json:"report_location_flag"  `
	// IsReportEnter 是否上报用户进入应用事件。0：不接收；1：接收
	IsReportEnter int `json:"isreportenter"  `
	// HomeURL 应用主页url
	HomeURL string `json:"home_url"  `
}

type ListAgentResponse struct {
	// AgentList AgentItemArray
	AgentList []ListAgent `json:"agentlist"  `
}

type SetWorkbenchTemplateRequest struct {
	// Type 模版类型，目前支持的自定义类型包括 “keydata”、 “image”、 “list”、 “webview” 。若设置的type为 “normal”,则相当于从自定义模式切换为普通宫格或者列表展示模式
	Type string `json:"type"  validate:"required"`
	// AgentID 应用id
	AgentID int `json:"agentid"  validate:"required"`
	// KeyData 若type指定为 “keydata”，且需要设置企业级别默认数据，则需要设置关键数据型模版数据,数据结构参考“关键数据型”
	KeyData WorkbenchTemplateItemKeyData `json:"keydata"  `
	// Image 若type指定为 “image”，且需要设置企业级别默认数据，则需要设置图片型模版数据,数据结构参考“图片型”
	Image WorkbenchTemplateItemImage `json:"image"  `
	// List 若type指定为 “list”，且需要设置企业级别默认数据，则需要设置列表型模版数据,数据结构参考“列表型”
	List []WorkbenchTemplateItemList `json:"list"  `
	// Webview 若type指定为 “webview”，且需要设置企业级别默认数据，则需要设置webview型模版数据,数据结构参考“webview型”
	Webview []WorkbenchTemplateItemWebView `json:"webview"  `
	// ReplaceUserData 是否覆盖用户工作台的数据。设置为true的时候，会覆盖企业所有用户当前设置的数据。若设置为false,则不会覆盖用户当前设置的所有数据。默认为false
	ReplaceUserData bool `json:"replace_user_data"  `
}

type SetWorkbenchTemplateResponse struct {
	// AccessToken 是
	AccessToken string `json:"access_token"  `
	// AgentID 是
	AgentID string `json:"agentid"  `
}

type GetWorkbenchTemplateRequest struct {
	// AgentID 应用id
	AgentID int `json:"agentid"  validate:"required"`
}

type GetWorkbenchTemplateResponse struct {
	// AccessToken 是
	AccessToken string `json:"access_token"  `
	// AgentID 是
	AgentID string `json:"agentid"  `
	// UserID 是
	UserID string `json:"userid"  `
	// Type 是
	Type string `json:"type"  `
	// KeyData 否
	KeyData WorkbenchTemplateItemKeyData `json:"keydata"  `
	// Image 否
	Image WorkbenchTemplateItemImage `json:"image"  `
	// List 否
	List []WorkbenchTemplateItemList `json:"list"  `
	// Webview 否
	Webview []WorkbenchTemplateItemWebView `json:"webview"  `
}

type SetWorkbenchDataRequest struct {
	// AgentID 应用id
	AgentID int `json:"agentid"  validate:"required"`
	// UserID 需要设置的用户的userid
	UserID string `json:"userid"  validate:"required"`
	// Type 目前支持  “keydata”、 “image”、 “list” 、”webview”
	Type string `json:"type"  validate:"required"`
	// KeyData 若type指定为 “keydata”，则需要设置关键数据型模版数据,数据结构参考“关键数据型”
	KeyData WorkbenchTemplateItemKeyData `json:"keydata"  `
	// Image 若type指定为 “image”，则需要设置图片型模版数据，数据结构参考“图片型”
	Image WorkbenchTemplateItemImage `json:"image"  `
	// List 若type指定为 “list”，则需要设置列表型模版数据，数据结构参考“列表型”
	List []WorkbenchTemplateItemList `json:"list"  `
	// Webview 若type指定为 “webview”，则需要设置webview型模版数据，数据结构参考“webview数据型”
	Webview []WorkbenchTemplateItemWebView `json:"webview"  `
}

type SetWorkbenchDataResponse struct {
	// AccessToken 是
	AccessToken string `json:"access_token"  `
	// AgentID 是
	AgentID string `json:"agentid"  `
	// UserID 是
	UserID string `json:"userid"  `
	// Type 是
	Type string `json:"type"  `
	// KeyData 否
	KeyData string `json:"keydata"  `
	// Image 否
	Image string `json:"image"  `
	// List 否
	List string `json:"list"  `
	// Webview 否
	Webview string `json:"webview"  `
}
