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
- 支持自建应用开发
- 支持第三方应用开发
- 生成回调事件模型 - 包含文档
  - 客户联系
  - 第三方回调
- 没有内部状态和 goroutine
- 自动尝试提前获取相应的 Token 和 Ticket
- 实现逻辑清晰 - 没有实现的接口可直接调用
- wwcrypt - 企业微信回调加密实现 - 作用同 sbzhu/weworkapi_golang
- 接口支持
  - 通讯录管理 - 成员、部门、标签、异步批量、互联企业
    - [ ] 通讯录回调
    - [ ] 导出任务完成通知 https://work.weixin.qq.com/api/doc/90000/90135/94946
  - 客户联系 - 客户管理、客户标签管理、在职继承、离职继承、客户群管理、消息推送、统计管理、变更回调

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
```

## Reference

- [wenerme/go-req](https://github.com/wenerme/go-req)
  - 接口底层库
- [xen0n/go-workwx](https://github.com/xen0n/go-workwx)
  - 比较成熟的 Golang 企业微信 SDK
  - 没有 第三方接口、服务商接口、会话存档
- [sbzhu/weworkapi_golang](https://github.com/sbzhu/weworkapi_golang)
  - 官方 Golang 加密库
