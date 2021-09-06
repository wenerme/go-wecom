package wecom

import (
	"github.com/wenerme/go-req"
)

func (c *Client) GetAPIDomainIP() (out IPListResponse, err error) {
	err = c.Request.With(req.Request{
		URL: "/cgi-bin/get_api_domain_ip",
	}).Fetch(&out)
	return
}

// GetCallbackIP 获取企业微信服务器的ip段
func (c *Client) GetCallbackIP() (out IPListResponse, err error) {
	err = c.Request.With(req.Request{
		URL: "/cgi-bin/getcallbackip",
	}).Fetch(&out)
	return
}
