package wecom_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/wenerme/go-req"
	"github.com/wenerme/go-wecom/wecom"
)

func ExampleNewClient() {
	client := wecom.NewClient(wecom.Conf{
		CorpID:     "",
		AgentID:    "",
		CorpSecret: "",
	})
	// 加载缓存 - 复用之前的 Token
	cache := wecom.ClientCache{}
	if bytes, err := os.ReadFile("wecom-cache.json"); err == nil {
		if json.Unmarshal(bytes, &cache) == nil {
			client.LoadCache(&cache)
		}
	}
	// 当 Token 变化时生成缓存
	client.OnTokenUpdate = func(c *wecom.Client) {
		cache := c.DumpCache()
		bytes, _ := json.Marshal(cache)
		_ = os.WriteFile("wecom-cache.json", bytes, 0o600)
	}

	// 访问接口会自动获取或使用当前缓存
	token, err := client.AccessToken()
	if err != nil {
		panic(err)
	}
	fmt.Println("Token", token)
	ticket, err := client.JsAPITicket()
	if err != nil {
		panic(err)
	}
	fmt.Println("Ticket", ticket)

	// 访问没有实现的接口
	dto := wecom.GetAPIDomainIPResponse{}
	err = client.Request.With(req.Request{
		URL:     "/cgi-bin/get_api_domain_ip",
		Options: []interface{}{
			// 如果不需要 access_token
			// wecom.WithoutAccessToken,
		},
	}).Fetch(&dto)
	if err != nil {
		panic(err)
	}
	fmt.Println("response", dto)
}
