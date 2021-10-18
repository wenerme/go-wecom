package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildURL(t *testing.T) {
	BuildOAuth2URL(BuildOAuth2URLOptions{
		AppID:       "wwabcd",
		Scope:       WecomOAuth2ScopeUserinfo,
		AgentID:     "1000010",
		RedirectURI: "https://example.com/auth/wecom/callback",
	})
	for _, test := range []struct {
		o BuildOAuth2URLOptions
		u string
	}{
		{
			o: BuildOAuth2URLOptions{
				AppID:       "wwabcd",
				Scope:       WecomOAuth2ScopeUserinfo,
				AgentID:     "1000010",
				RedirectURI: "https://example.com/auth/wecom/callback",
				State:       "STATE",
			},
			u: "https://open.weixin.qq.com/connect/oauth2/authorize?agentid=1000010&appid=wwabcd&redirect_uri=https%3A%2F%2Fexample.com%2Fauth%2Fwecom%2Fcallback&response_type=code&scope=snsapi_userinfo&state=STATE#wechat_redirect",
		},
		{
			o: BuildOAuth2URLOptions{
				AppID:       "wwabcd",
				RedirectURI: "https://example.com/auth/wecom/callback",
			},
			u: "https://open.weixin.qq.com/connect/oauth2/authorize?appid=wwabcd&redirect_uri=https%3A%2F%2Fexample.com%2Fauth%2Fwecom%2Fcallback&response_type=code&scope=snsapi_base#wechat_redirect",
		},
	} {
		assert.Equal(t, BuildOAuth2URL(test.o), test.u)
	}
}
