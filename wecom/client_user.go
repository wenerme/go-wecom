package wecom

import (
	"net/http"

	"github.com/wenerme/go-req"
)

type GetUserInfoRequest struct {
	AccessToken string `json:"access_token,omitempty"`
	Code        string `json:"code,omitempty"`
}

type GetUserInfoResponse struct {
	UserID         string `json:"UserId,omitempty"`          // 成员UserID
	OpenID         string `json:"OpenId,omitempty"`          // 非企业成员的标识，对当前企业唯一。不超过64字节
	DeviceID       string `json:"DeviceId,omitempty"`        // 手机设备号(由企业微信在安装时随机生成，删除重装会改变，升级不受影响)
	ExternalUserID string `json:"external_userid,omitempty"` // 外部联系人id，当且仅当用户是企业的客户，且跟进人在应用的可见范围内时返回。如果是第三方应用调用，针对同一个客户，同一个服务商不同应用获取到的id相同
}

// GetUserInfo
//
// see https://work.weixin.qq.com/api/doc/90000/90135/91023
func (c *Client) GetUserInfo(r *GetUserInfoRequest) (out GetUserInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  http.MethodGet,
		URL:     "/cgi-bin/user/getuserinfo",
		Query:   r,
		Options: []interface{}{},
	}).Fetch(&out)
	return
}
