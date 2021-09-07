package wecom

import (
	"context"
	"net/http"

	"github.com/wenerme/go-req"
)

const DefaultAPI = "https://qyapi.weixin.qq.com"

type Client struct {
	Conf          Conf
	Request       req.Request
	TokenProvider TokenProvider
}

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

type ClientCache struct {
	AccessToken              *TokenResponse
	JsAPITicket              *TicketResponse
	AgentTicket              *TicketResponse
	ProviderAccessTokenCache *ProviderTokenResponse
	SuiteAccessTokenCache    *SuiteTokenResponse
}

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

func (c *Client) GetJsAPITicket() (out TicketResponse, err error) {
	err = c.Request.With(req.Request{
		URL: "/cgi-bin/get_jsapi_ticket",
	}).Fetch(&out)
	if err == nil {
		out.ExpiresAt = int64(out.ExpiresIn) + timeNow().Unix()
	}
	return out, err
}

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
