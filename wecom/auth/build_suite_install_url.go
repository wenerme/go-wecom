package auth

type BuildSuiteInstallURLOption struct {
	SuiteID     string `json:"suite_id,omitempty"`
	PreAuthCode string `json:"pre_auth_code,omitempty"`
	RedirectURI string `json:"redirect_uri,omitempty"`
	State       string `json:"state,omitempty"` // a-zA-Z0-9 的参数值（不超过128个字节），用于第三方自行校验session，防止跨域攻击。
}

// BuildSuiteInstallURL build url for suite installation
//
// see https://work.weixin.qq.com/api/doc/90001/90143/90597
func BuildSuiteInstallURL(o BuildSuiteInstallURLOption) string {
	return buildURLValues("https://open.work.weixin.qq.com/3rdapp/install", o)
}
