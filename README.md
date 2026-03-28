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

**530+ APIs** across **37 categories**, auto-generated from official WeChat Work API documentation.

## Features

- **530+ API methods** auto-generated from official docs with Chinese comments
- Self-built app support (AccessToken)
- Third-party app support (AuthCorpAccessToken, SuiteToken, ProviderToken)
- WebSocket bot client for 智能机器人长连接 (streaming support)
- Robot webhook support
- Token caching with custom storage (memory, database, etc.)
- Auto pre-fetch tokens at 80% expiry
- wwcrypt - callback encryption (same as sbzhu/weworkapi_golang)
- Session archive (会话存档) with C SDK binding
- API + Event mock testing with serialization verification
- No internal state or goroutines

## Quick Start

```go
client := wecom.NewClient(wecom.Conf{
    CorpID:     "CORPID",
    AgentID:    1000002,
    CorpSecret: "SECRET",
})

// Auto-fetches and caches access_token
user, err := client.GetUser(&wecom.GetUserRequest{
    UserID: "zhangsan",
})
```

### Token Cache

```go
store := &wecom.SyncMapStore{}
// Restore from file
if bytes, err := os.ReadFile("wecom-cache.json"); err == nil {
    _ = store.Restore(bytes)
}
// Auto-save on change
store.OnChange = func(s *wecom.SyncMapStore) {
    _ = os.WriteFile("wecom-cache.json", s.Dump(), 0o600)
}

client := wecom.NewClient(wecom.Conf{
    CorpID:     "CORPID",
    AgentID:    1000002,
    CorpSecret: "SECRET",
    TokenProvider: &wecom.TokenCache{Store: store},
})
```

### Webhook

```go
wecom.WebhookSend(&wecom.WebhookSendRequest{
    Key:     "KEY",
    Content: wecom.SendTextContent{Content: "Hello"},
})
```

### Third-party App

```go
client := wecom.NewClient(wecom.Conf{
    CorpID:                "CORPID",
    ProviderSecret:        "PROVIDER_SECRET",
    AuthCorpID:            "AUTH_CORPID",
    AuthCorpPermanentCode: "PERM_CODE",
    SuiteID:               "SUITE_ID",
    SuiteSecret:           "SUITE_SECRET",
    SuiteTicket:           "SUITE_TICKET",
})
```

## API Coverage

530+ APIs across 37 categories, generated from [official documentation](https://developer.work.weixin.qq.com/document/path/90664):

| Category | APIs | Category | APIs |
|----------|------|----------|------|
| Meeting (会议) | 102 | WeDoc (文档) | 45 |
| Contacts (通讯录) | 36 | WeDrive (微盘) | 27 |
| School (家校沟通) | 25 | ExMail (邮件) | 24 |
| External Contact (客户联系) | 20 | Gov (政民沟通) | 20 |
| KF (微信客服) | 17 | Corp Group (上下游) | 14 |
| School App (家校应用) | 13 | Checkin (打卡) | 12 |
| Approval (审批) | 11 | Data Intelligence | 11 |
| Pay (企业支付) | 11 | Message (消息) | 10 |
| Security (安全管理) | 9 | Schedule (日程) | 6 |
| Agent (应用管理) | 4 | + 18 more categories | ... |

All APIs include:
- Typed request/response structs with `json` tags
- Chinese documentation comments from official docs
- `validate:"required"` tags for required parameters
- Links to official documentation

## Session Archive

> **Note**
>
> 1. Session archive data retained for **5** days
> 2. Pull from Sequence+1 (exclusive)
> 3. Max limit: 1000
> 4. Formats: jpg (image), amr (audio), mp4 (video)
> 5. MediaData MD5 may mismatch, retry is OK
> 6. MediaData max 512K per response

```bash
make bin
cp .env.example .env
./bin/wwfinance-libs         # extract C SDK
LD_LIBRARY_PATH=/tmp/wwf/libs ./bin/wwfinance-poller

# Docker
docker run --rm -it -v $PWD/.env:/app/.env -v $PWD/data:/app/data wener/go-wecom
```

```go
client, err := WeWorkFinanceSDK.NewClientFromEnv()
data, err := client.GetChatData(WeWorkFinanceSDK.GetChatDataOptions{
    Limit:   10,
    Timeout: 5,
})
for _, v := range data {
    fmt.Println(v.Message)
}
```

## Code Generation

API code is auto-generated from official WeChat Work documentation:

```
Official Docs (955 pages) → LLM extraction → JSON specs (531) → Go code (530+ methods)
```

Generated files use `client_gen_*.go` prefix with `// Code generated ... DO NOT EDIT.` header.
Hand-written code in `client_*.go` takes precedence and is never overwritten.

## Reference

- [wenerme/go-req](https://github.com/wenerme/go-req) - HTTP request library
- [xen0n/go-workwx](https://github.com/xen0n/go-workwx) - Another Go WeChat Work SDK
- [sbzhu/weworkapi_golang](https://github.com/sbzhu/weworkapi_golang) - Official crypto library
