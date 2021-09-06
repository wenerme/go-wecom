package wecom

import (
	"net/http"
	"time"

	"github.com/wenerme/go-req"
)

func (c *Client) GetProviderToken() (out ProviderTokenResponse, err error) {
	err = c.Request.With(req.Request{
		Method: http.MethodPost,
		URL:    "/cgi-bin/service/get_provider_token",
		Body: map[string]string{
			"corpid":          c.Conf.CorpID,
			"provider_secret": c.Conf.ProviderSecret,
		},
		Options: []interface{}{WithoutAccessToken},
	}).Fetch(&out)
	if err == nil {
		out.ExpireAt = int64(out.ExpiresIn) + time.Now().Unix()
	}
	return
}

type ProviderGetLoginInfoRequest struct {
	AuthCode string `json:"auth_code"`
}

// ProviderGetLoginInfo 获取登录用户信息
//
// see https://work.weixin.qq.com/api/doc/90001/90143/91125
func (c *Client) ProviderGetLoginInfo(r *ProviderGetLoginInfoRequest) (out ProviderGetLoginInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method: http.MethodPost,
		URL:    "/cgi-bin/service/get_login_info",
		Body:   r,
		Options: []interface{}{func(m *middlewareOptions) {
			m.GetToken = func(c *Client, r *req.Request) (string, string, error) {
				// 名字是 access_token 但内容为 provider_access_token
				token, err := c.ProviderAccessToken()
				return "access_token", token, err
			}
		}},
	}).Fetch(&out)
	return
}

type ProviderGetLoginInfoResponse struct {
	// 登录用户的类型：1.创建者 2.内部系统管理员 3.外部系统管理员 4.分级管理员 5.成员
	UserType int `json:"usertype"`
	UserInfo struct {
		// 登录用户的userid，登录用户在通讯录中时返回
		UserID string `json:"userid"`
		// 全局唯一。对于同一个服务商，不同应用获取到企业内同一个成员的open_userid是相同的，最多64个字节。 登录用户在通讯录中返回
		OpenUserID string `json:"open_userid"`
		// Name 登录用户的名字，登录用户在通讯录中时返回，此字段从2019年12月30日起，对新创建服务商不再返回，2020年6月30日起，对所有历史服务商不再返回真实name，使用userid代替name返回，第三方页面需要通过通讯录展示组件来展示名字
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	} `json:"user_info"`
	CorpInfo struct {
		CorpID string `json:"corpid"`
	} `json:"corp_info"`
	Agent []struct {
		AgentID  int `json:"agentid"`
		AuthType int `json:"auth_type"` // 该管理员对应用的权限：1.管理权限，0.使用权限
	} `json:"agent"` // 该管理员在该提供商中能使用的应用列表，当登录用户为管理员时返回
	AuthInfo struct {
		Department []struct {
			ID       int  `json:"id"`
			Writable bool `json:"writable"`
		} `json:"department"`
	} `json:"auth_info"` // 该管理员拥有的通讯录权限，当登录用户为管理员时返回
}
