package wecom

import "github.com/wenerme/go-req"

// GetAgent 获取指定的应用详情
// 对于互联企业的应用，如果需要获取应用可见范围内其他互联企业的部门与成员，请调用互联企业-获取应用可见范围接口
//
// see https://work.weixin.qq.com/api/doc/90001/90143/90363
func (c *Client) GetAgent(r *GetAgentRequest) (out GetAgentResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "GET",
		URL:    "/cgi-bin/agent/get",
		Query:  r,
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
func (c *Client) SetWorkbenchTemplate() (out SetWorkbenchTemplateResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/agent/set_workbench_template",
	}).Fetch(&out)
	return
}

// GetWorkbenchTemplate 获取应用在工作台展示的模版
//
// see https://work.weixin.qq.com/api/doc/90001/90143/94620
func (c *Client) GetWorkbenchTemplate() (out GetWorkbenchTemplateResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/agent/get_workbench_template",
	}).Fetch(&out)
	return
}

// SetWorkbenchData 设置应用在用户工作台展示的数据
//
// see https://work.weixin.qq.com/api/doc/90001/90143/94620
func (c *Client) SetWorkbenchData() (out SetWorkbenchDataResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/agent/set_workbench_data",
	}).Fetch(&out)
	return
}

type GetAgentRequest struct {
	// AgentID 应用id
	AgentID int `json:"agentid" validate:"required"`
}

type GetAgentResponse struct {
	// AgentID 企业应用id
	AgentID int `json:"agentid"`
	// Name 企业应用名称
	Name string `json:"name"`
	// SquareLogoURL 企业应用方形头像
	SquareLogoURL string `json:"square_logo_url"`
	// Description 企业应用详情
	Description string `json:"description"`
	// AllowUserInfos 企业应用可见范围（人员），其中包括userid
	AllowUserInfos string `json:"allow_userinfos"`
	// AllowParties 企业应用可见范围（部门）
	AllowParties string `json:"allow_partys"`
	// AllowTags 企业应用可见范围（标签）
	AllowTags string `json:"allow_tags"`
	// Close 企业应用是否被停用
	Close int `json:"close"`
	// RedirectDomain 企业应用可信域名
	RedirectDomain string `json:"redirect_domain"`
	// ReportLocationFlag 企业应用是否打开地理位置上报 0：不上报；1：进入会话上报；
	ReportLocationFlag int `json:"report_location_flag"`
	// IsReportEnter 是否上报用户进入应用事件。0：不接收；1：接收
	IsReportEnter int `json:"isreportenter"`
	// HomeURL 应用主页url
	HomeURL string `json:"home_url"`
}

type ListAgentResponse struct {
	// AgentList AgentItemArray
	AgentList string `json:"agentlist"`
}

type SetWorkbenchTemplateResponse struct {
	// AccessToken 是
	AccessToken string `json:"access_token"`
	// AgentID 是
	AgentID string `json:"agentid"`
}

type GetWorkbenchTemplateResponse struct {
	// AccessToken 是
	AccessToken string `json:"access_token"`
	// AgentID 是
	AgentID string `json:"agentid"`
	// UserID 是
	UserID string `json:"userid"`
	// Type 是
	Type string `json:"type"`
	// KeyData 否
	KeyData string `json:"keydata"`
	// Image 否
	Image string `json:"image"`
	// List 否
	List string `json:"list"`
	// Webview 否
	Webview string `json:"webview"`
}

type SetWorkbenchDataResponse struct {
	// AccessToken 是
	AccessToken string `json:"access_token"`
	// AgentID 是
	AgentID string `json:"agentid"`
	// UserID 是
	UserID string `json:"userid"`
	// Type 是
	Type string `json:"type"`
	// KeyData 否
	KeyData string `json:"keydata"`
	// Image 否
	Image string `json:"image"`
	// List 否
	List string `json:"list"`
	// Webview 否
	Webview string `json:"webview"`
}
