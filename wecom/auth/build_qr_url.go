package auth

// BuildProviderQRAuthURLOptions options for BuildProviderQRAuthURL
type BuildProviderQRAuthURLOptions struct {
	AppID       string `json:"appid,omitempty"` // 企业 CorpID 或 第三方 SuiteID
	RedirectURI string `json:"redirect_uri,omitempty"`
	UserType    string `json:"usertype,omitempty"` // admin代表管理员登录（使用微信扫码）,member代表成员登录（使用企业微信扫码），默认为admin
	State       string `json:"state,omitempty"`
	Lang        string `json:"lang,omitempty"` // 自定义语言，支持zh、en；lang为空则从Headers读取Accept-Language，默认值为zh
}

// BuildProviderQRAuthURL 第三方扫码登陆
//
// 1. 管理员从企业微信管理端单点登录第三方
// 2. 管理员或成员在第三方网站发起登录授权
//
// https://work.weixin.qq.com/api/doc/90001/90143/91124
func BuildProviderQRAuthURL(o BuildProviderQRAuthURLOptions) string {
	return buildURLValues("https://open.work.weixin.qq.com/wwopen/sso/3rd_qrConnect?usertype=admin&lang=", o)
}

// BuildQRAuthURLOptions options for BuildQRAuthURL
type BuildQRAuthURLOptions struct {
	AppID       string `json:"appid,omitempty"` // 企业 CorpID 或 第三方 SuiteID
	AgentID     string `json:"agentid,omitempty"`
	RedirectURI string `json:"redirect_uri,omitempty"`
	State       string `json:"state,omitempty"`
	Lang        string `json:"lang,omitempty"` // 自定义语言，支持zh、en；lang为空则从Headers读取Accept-Language，默认值为zh
}

// BuildQRAuthURL 企业自建应用扫码登陆
//
// https://work.weixin.qq.com/api/doc/90000/90135/91019
func BuildQRAuthURL(o BuildProviderQRAuthURLOptions) string {
	return buildURLValues("https://open.work.weixin.qq.com/wwopen/sso/qrConnect?lang=", o)
}
