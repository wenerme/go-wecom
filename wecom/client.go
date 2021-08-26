package wecom

import (
	"context"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/wenerme/go-req"
)

const DefaultAPI = "https://qyapi.weixin.qq.com"

type Client struct {
	Conf               Conf
	Request            req.Request
	AccessTokenCache   TokenResponse
	JsAPITicketCache   TicketResponse
	AgentTicketCache   TicketResponse
	ProviderTokenCache ProviderTokenResponse
	SuiteTokenCache    SuiteTokenResponse
	OnTokenUpdate      func(c *Client)

	updatingToken         int32
	updatingProviderToken int32
	updatingSuiteToken    int32
	updatingTicket        int32
	updatingAgentTicket   int32
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
	return c
}

type ClientCache struct {
	AccessToken   *TokenResponse
	JsAPITicket   *TicketResponse
	AgentTicket   *TicketResponse
	ProviderToken *ProviderTokenResponse
	SuiteToken    *SuiteTokenResponse
}

func (c *Client) DumpCache() ClientCache {
	a := &c.AccessTokenCache
	b := &c.JsAPITicketCache
	d := &c.AgentTicketCache
	e := &c.ProviderTokenCache
	f := &c.SuiteTokenCache
	now := time.Now().Unix()
	if a.ExpireAt < now {
		a = nil
	}
	if b.ExpireAt < now {
		b = nil
	}
	if d.ExpireAt < now {
		d = nil
	}
	if e.ExpireAt < now {
		e = nil
	}
	if f.ExpireAt < now {
		f = nil
	}
	cc := ClientCache{
		AccessToken:   a,
		JsAPITicket:   b,
		AgentTicket:   d,
		ProviderToken: e,
		SuiteToken:    f,
	}
	return cc
}

func (c *Client) LoadCache(cc *ClientCache) {
	if cc == nil {
		return
	}
	now := time.Now().Unix()
	if cc.AccessToken != nil && cc.AccessToken.ExpireAt > now {
		c.AccessTokenCache = *cc.AccessToken
	}
	if cc.JsAPITicket != nil && cc.JsAPITicket.ExpireAt > now {
		c.JsAPITicketCache = *cc.JsAPITicket
	}
	if cc.AgentTicket != nil && cc.AgentTicket.ExpireAt > now {
		c.AgentTicketCache = *cc.AgentTicket
	}
	if cc.ProviderToken != nil && cc.ProviderToken.ExpireAt > now {
		c.ProviderTokenCache = *cc.ProviderToken
	}
	if cc.SuiteToken != nil && cc.SuiteToken.ExpireAt > now {
		c.SuiteTokenCache = *cc.SuiteToken
	}
}

func (c *Client) shouldRefresh(ts int64) bool {
	// 默认 2 小时有效期
	// 提前 30 分钟换取
	return ts-time.Now().Unix() < 30*60
}

func (c *Client) JsAPITicket() (string, error) {
	var lastErr error
	if c.shouldRefresh(c.JsAPITicketCache.ExpireAt) {
		if atomic.CompareAndSwapInt32(&c.updatingTicket, 0, 1) {
			defer func() {
				atomic.CompareAndSwapInt32(&c.updatingTicket, 1, 0)
			}()
			var next TicketResponse
			next, lastErr = c.GetJsAPITicket()
			if lastErr != nil {
				logrus.WithError(lastErr).Error("get js api ticket failed")
			} else {
				c.JsAPITicketCache = next
				if c.OnTokenUpdate != nil {
					c.OnTokenUpdate(c)
				}
			}
		}
	}

	return validToken(c.JsAPITicketCache, "js api ticket", lastErr)
}

func (c *Client) AgentTicket() (string, error) {
	var lastErr error
	if c.shouldRefresh(c.AgentTicketCache.ExpireAt) {
		if atomic.CompareAndSwapInt32(&c.updatingAgentTicket, 0, 1) {
			defer func() {
				atomic.CompareAndSwapInt32(&c.updatingAgentTicket, 1, 0)
			}()
			var next TicketResponse
			next, lastErr = c.GetAgentTicket()
			if lastErr != nil {
				logrus.WithError(lastErr).Error("get agent ticket failed")
			} else {
				c.AgentTicketCache = next
				if c.OnTokenUpdate != nil {
					c.OnTokenUpdate(c)
				}
			}
		}
	}

	return validToken(c.AgentTicketCache, "agent ticket", lastErr)
}

func (c *Client) AccessToken() (string, error) {
	var lastErr error
	if c.shouldRefresh(c.AccessTokenCache.ExpireAt) {
		if atomic.CompareAndSwapInt32(&c.updatingToken, 0, 1) {
			defer func() {
				atomic.CompareAndSwapInt32(&c.updatingToken, 1, 0)
			}()
			var next TokenResponse
			next, lastErr = c.GetToken()
			if lastErr != nil {
				logrus.WithError(lastErr).Error("get token failed")
			} else {
				c.AccessTokenCache = next
				if c.OnTokenUpdate != nil {
					c.OnTokenUpdate(c)
				}
			}
		}
	}

	return validToken(c.AccessTokenCache, "access token", lastErr)
}

func (c *Client) ProviderToken() (string, error) {
	var lastErr error
	if c.shouldRefresh(c.ProviderTokenCache.ExpireAt) {
		if atomic.CompareAndSwapInt32(&c.updatingProviderToken, 0, 1) {
			defer func() {
				atomic.CompareAndSwapInt32(&c.updatingProviderToken, 1, 0)
			}()
			var next ProviderTokenResponse
			next, lastErr = c.GetProviderToken()
			if lastErr != nil {
				logrus.WithError(lastErr).Error("get token failed")
			} else {
				c.ProviderTokenCache = next
				if c.OnTokenUpdate != nil {
					c.OnTokenUpdate(c)
				}
			}
		}
	}
	return validToken(c.ProviderTokenCache, "provider token", lastErr)
}

func (c *Client) SuiteAccessToken() (string, error) {
	var lastErr error
	if c.shouldRefresh(c.SuiteTokenCache.ExpireAt) {
		if atomic.CompareAndSwapInt32(&c.updatingSuiteToken, 0, 1) {
			defer func() {
				atomic.CompareAndSwapInt32(&c.updatingSuiteToken, 1, 0)
			}()
			var next SuiteTokenResponse
			next, lastErr = c.GetSuiteToken(&GetSuiteTokenRequest{
				SuiteID:     c.Conf.SuiteID,
				SuiteSecret: c.Conf.SuiteSecret,
				SuiteTicket: c.Conf.SuiteTicket,
			})
			if lastErr != nil {
				logrus.WithError(lastErr).Error("get suite token failed")
			} else {
				c.SuiteTokenCache = next
				if c.OnTokenUpdate != nil {
					c.OnTokenUpdate(c)
				}
			}
		}
	}

	return validToken(c.SuiteTokenCache, "suite token", lastErr)
}

func validToken(cache Token, typ string, lastErr error) (string, error) {
	if cache.GetExpireAt() < time.Now().Unix() {
		if lastErr == nil {
			return "", fmt.Errorf("%v expired or invalid", typ)
		}
		return "", errors.Wrapf(lastErr, "get %v failed", typ)
	}
	return cache.GetToken(), nil
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
		out.ExpireAt = int64(out.ExpiresIn) + time.Now().Unix()
	}
	return out, err
}

func (c *Client) GetJsAPITicket() (out TicketResponse, err error) {
	err = c.Request.With(req.Request{
		URL: "/cgi-bin/get_jsapi_ticket",
	}).Fetch(&out)
	if err == nil {
		out.ExpireAt = int64(out.ExpiresIn) + time.Now().Unix()
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
		out.ExpireAt = int64(out.ExpiresIn) + time.Now().Unix()
	}
	return out, err
}
