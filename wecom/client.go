package wecom

import (
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/wenerme/go-req"
)

const DefaultAPI = "https://qyapi.weixin.qq.com"

type Client struct {
	Conf             Conf
	Request          req.Request
	AccessTokenCache TokenResponse
	JsApiTicketCache TicketResponse
	AgentTicketCache TicketResponse
	OnTokenUpdate    func(c *Client)
	tokenMu          sync.RWMutex
}

func NewClient(conf Conf) *Client {
	c := &Client{
		Conf: conf,
		Request: req.Request{
			BaseURL: DefaultAPI,
			Method:  http.MethodGet,
			Options: []interface{}{req.JSONEncode, req.JSONDecode},
		},
	}
	return c
}

type ClientCache struct {
	AccessToken *TokenResponse
	JsApiTicket *TicketResponse
	AgentTicket *TicketResponse
}

func (c *Client) DumpCache() ClientCache {
	a := c.AccessTokenCache
	b := c.JsApiTicketCache
	d := c.AgentTicketCache
	return ClientCache{
		AccessToken: &a,
		JsApiTicket: &b,
		AgentTicket: &d,
	}
}

func (c *Client) LoadCache(cc *ClientCache) {
	if cc == nil {
		return
	}
	if cc.AccessToken != nil {
		c.AccessTokenCache = *cc.AccessToken
	}
	if cc.JsApiTicket != nil {
		c.JsApiTicketCache = *cc.JsApiTicket
	}
	if cc.AgentTicket != nil {
		c.AgentTicketCache = *cc.AgentTicket
	}
}

func (c *Client) shouldRefresh(ts int64) bool {
	// 默认 2 小时有效期
	// 提前 30 分钟换取
	return ts-time.Now().Unix() < 30*60
}

func (c *Client) JsApiTicket() (string, error) {
	if c.shouldRefresh(c.JsApiTicketCache.ExpireAt) {
		// todo try lock
		c.tokenMu.Lock()
		defer c.tokenMu.Unlock()
		cache := c.JsApiTicketCache
		if c.shouldRefresh(cache.ExpireAt) {
			var err error
			var next TicketResponse
			next, err = c.GetJsApiTicket()
			if err != nil {
				logrus.WithError(err).Error("get js api ticket failed")
				if cache.ExpireAt < time.Now().Unix() {
					return "", errors.Wrap(err, "get js api ticket failed")
				}
			} else {
				cache = next
				c.JsApiTicketCache = next
				if c.OnTokenUpdate != nil {
					c.OnTokenUpdate(c)
				}
			}
		}
	}
	return c.JsApiTicketCache.Ticket, nil
}

func (c *Client) AgentTicket() (string, error) {
	if c.shouldRefresh(c.AgentTicketCache.ExpireAt) {
		// todo try lock
		c.tokenMu.Lock()
		defer c.tokenMu.Unlock()
		cache := c.AgentTicketCache
		if c.shouldRefresh(cache.ExpireAt) {
			var err error
			var next TicketResponse
			next, err = c.GetAgentTicket()
			if err != nil {
				logrus.WithError(err).Error("get js api ticket failed")
				if cache.ExpireAt < time.Now().Unix() {
					return "", errors.Wrap(err, "get js api ticket failed")
				}
			} else {
				cache = next
				c.AgentTicketCache = next
				if c.OnTokenUpdate != nil {
					c.OnTokenUpdate(c)
				}
			}
		}
	}
	return c.AgentTicketCache.Ticket, nil
}

func (c *Client) AccessToken() (string, error) {
	if c.shouldRefresh(c.AccessTokenCache.ExpireAt) {
		// todo try lock
		c.tokenMu.Lock()
		defer c.tokenMu.Unlock()
		cache := c.AccessTokenCache
		if c.shouldRefresh(cache.ExpireAt) {
			var err error
			var next TokenResponse
			next, err = c.GetToken()
			if err != nil {
				logrus.WithError(err).Error("get token failed")
				if cache.ExpireAt < time.Now().Unix() {
					return "", errors.Wrap(err, "get token failed")
				}
			} else {
				cache = next
				c.AccessTokenCache = next
				if c.OnTokenUpdate != nil {
					c.OnTokenUpdate(c)
				}
			}
		}
	}
	return c.AccessTokenCache.AccessToken, nil
}

func (c *Client) GetToken() (out TokenResponse, err error) {
	var er GenericResponse
	err = c.Request.With(req.Request{
		URL: "/cgi-bin/gettoken",
		Query: map[string]interface{}{
			"corpid":     c.Conf.CorpID,
			"corpsecret": c.Conf.CorpSecret,
		},
	}).Fetch(&er, &out)
	if err == nil {
		err = er.AsError()
	}
	if err == nil {
		out.ExpireAt = int64(out.ExpiresIn) + time.Now().Unix()
	}
	return out, err
}

func (c *Client) GetJsApiTicket() (out TicketResponse, err error) {
	var accessToken string
	accessToken, err = c.AccessToken()
	if err != nil {
		return
	}
	var er GenericResponse
	err = c.Request.With(req.Request{
		URL: "/cgi-bin/get_jsapi_ticket",
		Query: map[string]interface{}{
			"access_token": accessToken,
		},
	}).Fetch(&er, &out)
	if err == nil {
		err = er.AsError()
	}
	if err == nil {
		out.ExpireAt = int64(out.ExpiresIn) + time.Now().Unix()
	}
	return out, err
}

func (c *Client) GetAgentTicket() (out TicketResponse, err error) {
	var accessToken string
	accessToken, err = c.AccessToken()
	if err != nil {
		return
	}
	var er GenericResponse
	err = c.Request.With(req.Request{
		URL: "/cgi-bin/ticket/get",
		Query: map[string]interface{}{
			"access_token": accessToken,
			"type":         "agent_config",
		},
	}).Fetch(&er, &out)
	if err == nil {
		err = er.AsError()
	}
	if err == nil {
		out.ExpireAt = int64(out.ExpiresIn) + time.Now().Unix()
	}
	return out, err
}
