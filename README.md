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

- 支持自建应用开发 - AccessToken
- 支持第三方应用开发 - AuthCorpAccessToken
- 支持缓存所有带时效的信息 - AccessToken, JsTicket, AgentTicket, SuiteToken, AuthCorpAccessToken, PreAuthCode, ProviderAccessToken
  - 缓存支持自定义存储 - 默认内存存储
- 支持从自定义的存储获取 密钥 信息 - SuiteTicket, PermanentCode
- 没有内部状态和 goroutine
- 自动尝试提前获取相应的 Token 和 Ticket - 有效期的 80%
- 实现逻辑清晰 - 没有实现的接口可直接调用
- wwcrypt - 企业微信回调加密实现 - 作用同 sbzhu/weworkapi_golang
- 数据模型大多基于官方接口文档生成 - 包含注释说明
- 包含 API+Event Mock 测试

```go
package wecom_test

import (
  "fmt"
  "os"

  "github.com/wenerme/go-req"
  "github.com/wenerme/go-wecom/wecom"
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
```

### 第三方应用开发配置
- 根据使用的接口不同，用到的信息也会不同

```go
client := wecom.NewClient(wecom.Conf{
  CorpID:   "",
  ProviderSecret: "",
  AuthCorpID:   "",
  AuthCorpPermanentCode: "",
  SuiteID:      "",
  SuiteSecret:  "",
  SuiteTicket:  "",
})
```

## 接口支持情况

* [x] 通讯录管理
* [x] 客户联系
* [ ] 微信客服
* [x] 身份验证
* [x] 应用管理
* [x] 消息推送
* [ ] 素材管理
* [ ] OA
* [-] 效率工具
* [ ] 企业支付
* [-] 会话内容存档
* [ ] 企业互联
* [ ] 电子发票

<details>
<summary>通讯录管理</summary>

* [x] 成员管理
  - [x] 创建成员
  - [x] 读取成员
  - [x] 更新成员
  - [x] 删除成员
  - [x] 批量删除成员
  - [x] 获取部门成员
  - [x] 获取部门成员详情
  - [x] userid与openid互换
  - [x] 二次验证
  - [x] 邀请成员
* [x] 部门管理
  - [x] 创建部门
  - [x] 更新部门
  - [x] 删除部门
  - [x] 获取部门列表
* [x] 标签管理
  - [x] 创建标签
  - [x] 更新标签名字
  - [x] 删除标签
  - [x] 获取标签成员
  - [x] 增加标签成员
  - [x] 删除标签成员
  - [x] 获取标签列表
* [x] 异步批量接口
  - [x] 增量更新成员
  - [x] 全量覆盖成员
  - [x] 全量覆盖部门
  - [x] 获取异步任务结果
* [x] 通讯录回调通知
  - [x] 成员变更通知
  - [x] 部门变更通知
  - [x] 标签变更通知
  - [x] 异步任务完成通知

</details>

<details>
<summary>客户联系 API</summary>

* [x] 成员对外信息
* [x] 客户管理
  - [x] 获取客户列表
  - [x] 获取客户详情
  - [x] 批量获取客户详情
  - [x] 修改客户备注信息
* [x] 客户标签管理
  - [x] 管理企业标签
  - [x] 编辑客户企业标签
* [x] 客户分配
  - [x] 获取离职成员列表
  - [x] 分配在职或离职成员的客户
  - [x] 查询客户接替结果
  - [x] 分配离职成员的客户群
* [x] 变更回调通知
  - [x] 添加企业客户事件
  - [x] 编辑企业客户事件
  - [x] 外部联系人免验证添加成员事件
  - [x] 删除企业客户事件
  - [x] 删除跟进成员事件
  - [x] 客户接替失败事件
  - [x] 客户群变更事件

</details>

<details>
<summary>身份验证 API</summary>

* [x] 获取访问用户身份

</details>

<details>
<summary>应用管理</summary>

* [x] 获取应用
* [x] 设置应用
* [x] 自定义菜单
  - [x] 创建菜单
  - [x] 获取菜单
  - [x] 删除菜单

</details>

<details>
<summary>消息推送</summary>

* [x] 发送应用消息
* [x] 接收消息
* [x] 发送消息到群聊会话
  - [x] 创建群聊会话
  - [x] 修改群聊会话
  - [x] 获取群聊会话
  - [x] 应用推送消息

### 消息类型

* [x] 文本消息
* [x] 图片消息
* [x] 语音消息
* [x] 视频消息
* [x] 文件消息
* [x] 文本卡片消息
* [x] 图文消息
* [x] 图文消息（mpnews）
* [x] markdown消息
* [x] 任务卡片消息

</details>

<details>
<summary>素材管理 API</summary>

* [ ] 上传临时素材
* [ ] 上传永久图片
* [ ] 获取临时素材
* [ ] 获取高清语音素材

</details>

<details>
<summary>OA</summary>

* [ ] 打卡
  - [ ] 获取企业所有打卡规则
  - [ ] 获取员工打卡规则
  - [ ] 获取打卡记录数据
  - [ ] 获取打卡日报数据
  - [ ] 获取打卡月报数据
  - [ ] 获取打卡人员排班信息
  - [ ] 为打卡人员排班
  - [ ] 录入打卡人员人脸信息
* [ ] 审批
  - [ ] 获取审批模板详情
  - [ ] 提交审批申请
  - [ ] 审批申请状态变化回调通知
  - [ ] 批量获取审批单号
  - [ ] 获取审批申请详情
  - [ ] 获取企业假期管理配置
  - [ ] 修改成员假期余额
* [ ] 汇报
  - [ ] 批量获取汇报记录单号
  - [ ] 获取汇报记录详情
  - [ ] 获取汇报统计数据
* [ ] 自建应用
  - [ ] 审批流程引擎
* [ ] 会议室
  - [ ] 会议室管理
  - [ ] 会议室预定管理
* [ ] 紧急通知应用
  - [ ] 发起语音电话
  - [ ] 获取接听状态

</details>


<details>
<summary>效率工具</summary>

* [x] 日程
  - [x] 日历接口
  - [x] 日程接口
  - [x] 回调事件
* [ ] 会议
  - [ ] 创建预约会议
  - [ ] 修改预约会议
  - [ ] 取消预约会议
  - [ ] 获取成员会议ID列表
  - [ ] 获取会议详情
* [ ] 直播
* [ ] 微盘
  - [ ] 空间管理
  - [ ] 空间权限
  - [ ] 文件管理
  - [ ] 文件权限
* [ ] 公费电话
  - [ ] 获取公费电话拨打记录

</details>

<details>
<summary>企业支付</summary>

* [ ] 企业红包
* [ ] 向员工付款
* [ ] 向员工收款
* [ ] 对外收款
* [ ] 签名算法

</details>


<details>
<summary>企业互联</summary>

* [ ] 获取应用共享信息
* [ ] 获取下级企业的access_token
* [ ] 获取下级企业的小程序session

</details>


<details>
<summary>会话内容存档</summary>

* [x] 获取会话内容存档开启成员列表
* [x] 获取会话同意情况
* [x] 客户同意进行聊天内容存档事件回调
* [x] 获取会话内容存档内部群信息

</details>

<details>
<summary>电子发票</summary>

* [ ] 查询电子发票
* [ ] 更新发票状态
* [ ] 批量更新发票状态
* [ ] 批量查询电子发票

</details>

- TBD
  - 将企业管理员添加为外部联系人 https://work.weixin.qq.com/api/doc/13613
  - change_external_tag shuffle
  - 批量安装应用 https://open.work.weixin.qq.com/api/doc/20990

## Reference

- [wenerme/go-req](https://github.com/wenerme/go-req)
  - 接口底层库
- [xen0n/go-workwx](https://github.com/xen0n/go-workwx)
  - 比较成熟的 Golang 企业微信 SDK
  - 没有 第三方接口、服务商接口、会话存档
- [sbzhu/weworkapi_golang](https://github.com/sbzhu/weworkapi_golang)
  - 官方 Golang 加密库
