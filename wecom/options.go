package wecom

import "github.com/wenerme/go-req"

// WithoutAccessToken when request do not add query access_token
func WithoutAccessToken(o *middlewareOptions) {
	o.WithoutAccessToken = true
}

// WithSuiteAccessToken when request add query suite_access_token
func WithSuiteAccessToken(o *middlewareOptions) {
	o.GetToken = func(c *Client, r *req.Request) (string, string, error) {
		token, err := c.SuiteAccessToken()
		return "suite_access_token", token, err
	}
}

// WithProviderAccessToken when request add query provider_access_token
func WithProviderAccessToken(o *middlewareOptions) {
	o.GetToken = func(c *Client, r *req.Request) (string, string, error) {
		token, err := c.ProviderAccessToken()
		return "provider_access_token", token, err
	}
}

// Debug request with debug=1
//
// About debug https://work.weixin.qq.com/api/doc/90000/90003/90487#debug%E6%A8%A1%E5%BC%8F%E8%B0%83%E7%94%A8%E6%8E%A5%E5%8F%A3
// Query Hint https://open.work.weixin.qq.com/devtool/query
func Debug(o *middlewareOptions) {
	o.Debug = true
}
