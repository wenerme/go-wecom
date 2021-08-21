package wecom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	c := &JsSdkConfig{
		Nonce:     "Wm3WZYTPz0wzccnW",
		Timestamp: 1414587457,
	}
	c.Sign("sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg", "http://mp.weixin.qq.com?params=value")
	sum := "0f9de62fce790f9a083d5c99e95740ceb90c27ed"
	assert.Equal(t, sum, sha1sum("jsapi_ticket=sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg&noncestr=Wm3WZYTPz0wzccnW&timestamp=1414587457&url=http://mp.weixin.qq.com?params=value"))
	assert.Equal(t, sum, c.Signature)
}
