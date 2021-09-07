package wecom_test

import (
	"fmt"
	"github.com/wenerme/go-req"
	"github.com/wenerme/go-wecom/wecom"
	"os"
)

func ExampleNewClient() {
	// token store - 默认内存 Map - 可以使用数据库实现
	store := &wecom.SyncMapStore{}
	// 加载缓存 - 复用之前的 Token
	if bytes, err := os.ReadFile("wecom-cache.json"); err == nil {
		_ = store.Restore(bytes)
	}
	// 当 Token 变化时生成缓存文件
	store.OnChange = func(s *wecom.SyncMapStore) {
		_ = os.WriteFile("wecom-cache.json", s.Dump(), 0o600)
	}

	client := wecom.NewClient(wecom.Conf{
		CorpID:     "",
		AgentID:    "",
		CorpSecret: "",
		// 不配置默认使用 内存缓存
		TokenProvider: &wecom.TokenCache{
			Store: store,
		},
	})

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
	dto := wecom.IPListResponse{}
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
