package wecom

import "github.com/wenerme/go-req"

func WithoutAccessToken(o *middlewareOptions) {
	o.WithoutAccessToken = true
}

func WithSuiteAccessToken(o *middlewareOptions) {
	o.GetToken = func(c *Client, r *req.Request) (string, string, error) {
		token, err := c.SuiteAccessToken()
		return "suite_access_token", token, err
	}
}
