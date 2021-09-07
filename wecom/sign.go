package wecom

import (
	"crypto/sha1" // nolint:gosec
	"encoding/hex"
	"fmt"
)

type JsSdkConfig struct {
	AppID     string `json:"appId"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonceStr"`
	Signature string `json:"signature"`
}

type JsSdkAgentConfig struct {
	CorpID    string `json:"corpid"`
	AgentID   string `json:"agentid"`
	Timestamp int64  `json:"timestamp"`
	Nonce     string `json:"nonceStr"`
	Signature string `json:"signature"`
}

func (c *Client) SignAgentConfig(url string) (*JsSdkAgentConfig, error) {
	ticket, err := c.AgentTicket()
	if err != nil {
		return nil, err
	}
	o := &JsSdkAgentConfig{
		CorpID:  c.Conf.CorpID,
		AgentID: c.Conf.AgentID,
	}
	o.Sign(ticket, url)
	return o, nil
}

func (c *Client) SignConfig(url string) (*JsSdkConfig, error) {
	ticket, err := c.JsAPITicket()
	if err != nil {
		return nil, err
	}
	o := &JsSdkConfig{
		AppID: c.Conf.CorpID,
	}
	o.Sign(ticket, url)
	return o, nil
}

func (o *JsSdkConfig) Sign(ticket string, url string) {
	if o.Nonce == "" {
		o.Nonce = createNonce()
	}
	if o.Timestamp == 0 {
		o.Timestamp = timeNow().Unix()
	}
	raw := fmt.Sprintf("jsapi_ticket=%v&noncestr=%v&timestamp=%v&url=%v", ticket, o.Nonce, o.Timestamp, url)
	o.Signature = sha1sum(raw)
}

func (o *JsSdkAgentConfig) Sign(ticket string, url string) {
	if o.Nonce == "" {
		o.Nonce = createNonce()
	}
	if o.Timestamp == 0 {
		o.Timestamp = timeNow().Unix()
	}
	raw := fmt.Sprintf("jsapi_ticket=%v&noncestr=%v&timestamp=%v&url=%v", ticket, o.Nonce, o.Timestamp, url)
	o.Signature = sha1sum(raw)
}

func sha1sum(s string) string {
	sum := sha1.Sum([]byte(s)) //nolint:gosec
	return hex.EncodeToString(sum[:])
}
