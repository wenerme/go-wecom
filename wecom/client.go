package wecom

import (
	"context"
	"net/http"

	"github.com/wenerme/go-req"
)

// Client of Wecom/WechatWork/企业微信
type Client struct {
	Conf          Conf
	Request       req.Request
	TokenProvider TokenProvider
}

// NewClient create a new Client
func NewClient(conf Conf) *Client {
	ctx := context.Background()
	c := &Client{
		Conf: conf,
		Request: req.Request{
			BaseURL: DefaultAPI,
			Method:  http.MethodGet,
			Options: []interface{}{requestMiddleware, req.JSONEncode, req.JSONDecode},
		},
	}
	c.Request.Context = NewContext(ctx, c)
	c.Request.Extension.With(wecomeMiddleware)
	c.TokenProvider = conf.TokenProvider

	if c.TokenProvider == nil {
		c.TokenProvider = &TokenCache{
			Store: &SyncMapStore{},
		}
	}
	return c
}

// With create new Client with this conf
func (c *Client) With(conf Conf) (neo *Client) {
	cc := *c
	neo = &cc
	neo.Conf = conf
	neo.Request.Context = NewContext(cc.Request.Context, neo)
	return
}

// GetToken request an access_token
func (c *Client) GetToken() (out TokenResponse, err error) {
	err = c.Request.With(req.Request{
		URL: "/cgi-bin/gettoken",
		Query: map[string]interface{}{
			"corpid":     c.Conf.CorpID,
			"corpsecret": c.Conf.CorpSecret,
		},
		Options: []interface{}{WithoutAccessToken},
	}).Fetch(&out)
	if err == nil {
		out.ExpiresAt = int64(out.ExpiresIn) + timeNow().Unix()
	}
	return out, err
}

// GetJsAPITicket request a JsAPITicket
func (c *Client) GetJsAPITicket() (out TicketResponse, err error) {
	err = c.Request.With(req.Request{
		URL: "/cgi-bin/get_jsapi_ticket",
	}).Fetch(&out)
	if err == nil {
		out.ExpiresAt = int64(out.ExpiresIn) + timeNow().Unix()
	}
	return out, err
}

// GetAgentTicket request a AgentTicket
func (c *Client) GetAgentTicket() (out TicketResponse, err error) {
	err = c.Request.With(req.Request{
		URL: "/cgi-bin/ticket/get",
		Query: map[string]interface{}{
			"type": "agent_config",
		},
	}).Fetch(&out)
	if err == nil {
		out.ExpiresAt = int64(out.ExpiresIn) + timeNow().Unix()
	}
	return out, err
}
