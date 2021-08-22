package wecom

import (
	"github.com/wenerme/go-req"
)

func (c *Client) GetAPIDomainIP() (out GetAPIDomainIPResponse, err error) {
	var er GenericResponse
	err = c.Request.With(req.Request{
		URL: "/cgi-bin/get_api_domain_ip",
	}).Fetch(&er, &out)
	return out, err
}
