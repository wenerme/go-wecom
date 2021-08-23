package wecom

import "github.com/wenerme/go-req"

func (c *Client) ListExternalContact(r *ListExternalContactRequest) (out ListExternalContactResponse, err error) {
	err = c.Request.With(req.Request{
		URL:   "/cgi-bin/externalcontact/list",
		Query: r,
	}).Fetch(&out)
	return
}
