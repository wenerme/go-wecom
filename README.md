# go-wecom

[![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Coverage Status][cov-img]][cov] [![Go Report Card][report-card-img]][report-card]

[doc-img]: https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square
[doc]: https://pkg.go.dev/github.com/wenerme/go-wecom?tab=doc
[ci-img]: https://github.com/wenerme/go-wecom/actions/workflows/ci.yml/badge.svg
[ci]: https://github.com/wenerme/go-wecom/actions/workflows/ci.yml
[cov-img]: https://codecov.io/gh/wenerme/go-wecom/branch/main/graph/badge.svg
[cov]: https://codecov.io/gh/wenerme/go-wecom/branch/main
[report-card-img]: https://goreportcard.com/badge/github.com/wenerme/go-wecom
[report-card]: https://goreportcard.com/report/github.com/wenerme/go-wecom

Wechat Work/Wecom/企业微信 Golang SDK

## 特性

- 支持缓存 Token、Ticket、AgentTicket
- 没有内部状态和 goroutine
- 自动尝试提前获取相应的 Token 和 Ticket
- 实现逻辑清晰 - 没有实现的接口可直接调用

```go
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

  // 加载缓存 - 复用之前的 Token
  cache := wecom.ClientCache{}
  bytes, err := os.ReadFile("wecom-cache.json")
  if err == nil {
    if json.Unmarshal(bytes, &cache) == nil {
      client.LoadCache(&cache)
    }
  }

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
```

## Reference

- [xen0n/go-workwx](https://github.com/xen0n/go-workwx)
  - 比较成熟的 Golang 企业微信 SDK
  - 没有 第三方接口、服务商接口、会话存档
