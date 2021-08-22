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
