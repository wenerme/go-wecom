package auth

import (
	"net/url"

	"github.com/wenerme/go-req"
)

// BuildOAuth2URLOptions options for BuildOAuth2URL
type BuildOAuth2URLOptions struct {
	AppID        string `json:"appid,omitempty"` // 企业 CorpID 或 第三方 SuiteID
	RedirectURI  string `json:"redirect_uri,omitempty"`
	ResponseType string `json:"response_type,omitempty"`
	Scope        string `json:"scope,omitempty"` // 默认 snsapi_base, 自建使用 snsapi_userinfo 或 snsapi_privateinfo 需要 AgentID
	State        string `json:"state,omitempty"`
	AgentID      string `json:"agentid,omitempty"`
}

const (
	WecomOAuth2ScopeBase        = "snsapi_base"        // 静默授权，可获取成员的的基础信息（UserId与DeviceId）
	WecomOAuth2ScopeUserinfo    = "snsapi_userinfo"    // 静默授权，可获取成员的详细信息，但不包含手机、邮箱
	WecomOAuth2ScopePrivateInfo = "snsapi_privateinfo" // 手动授权，可获取成员的详细信息，包含手机、邮箱
)

// BuildOAuth2URL build oauth2 redirect url
//
// https://work.weixin.qq.com/api/doc/90001/90143/91120
func BuildOAuth2URL(o BuildOAuth2URLOptions) string {
	return buildURLValues("https://open.weixin.qq.com/connect/oauth2/authorize?response_type=code&scope=snsapi_base#wechat_redirect", o)
}

func buildURLValues(base string, o interface{}) string {
	u, err := url.Parse(base)
	if err != nil {
		panic(err)
	}
	query := u.Query()
	val, err := req.ValuesOf(o)
	if err != nil {
		panic(err)
	}
	mergeValues(query, val)
	u.RawQuery = query.Encode()
	return u.String()
}

func mergeValues(dst url.Values, src url.Values) url.Values {
	for k, v := range src {
		dst.Set(k, v[0])
	}
	return dst
}
