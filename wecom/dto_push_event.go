package wecom

import "encoding/xml"

// SubscribePushEvent 小程序在管理端开启接收消息配置后，也可收到关注/取消关注事件本事件触发时机为：
//
// 成员已经加入企业，管理员添加成员到应用可见范围(或移除可见范围)时成员已经在应用可见范围，成员加入(或退出)企业时
type SubscribePushEvent struct {
	XMLName xml.Name `xml:"xml"`
	// AgentID 企业应用的id，整型。可在应用的设置页面查看
	AgentID string `xml:"AgentID"`
	// CreateTime 消息创建时间（整型）
	CreateTime int64 `xml:"CreateTime"`
	// Event 事件类型，subscribe(关注)、unsubscribe(取消关注)
	Event string `xml:"Event"`
	// EventKey 事件KEY值，此事件该值为空
	EventKey string `xml:"EventKey"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName"`
	// MsgType 消息类型，此时固定为：event
	MsgType string `xml:"MsgType"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName"`
}

func (SubscribePushEvent) EventType() string {
	return "subscribe" //nolint:goconst
}

func (SubscribePushEvent) MessageType() string {
	return "event" //nolint:goconst
}

// EnterAgentPushEvent 本事件在成员进入企业微信的应用时触发
type EnterAgentPushEvent struct {
	XMLName xml.Name `xml:"xml"`
	// AgentID 企业应用的id，整型。可在应用的设置页面查看
	AgentID string `xml:"AgentID"`
	// CreateTime 消息创建时间（整型）
	CreateTime int64 `xml:"CreateTime"`
	// Event 事件类型：enter_agent
	Event string `xml:"Event"`
	// EventKey 事件KEY值，此事件该值为空
	EventKey string `xml:"EventKey"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName"`
	// MsgType 消息类型，此时固定为：event
	MsgType string `xml:"MsgType"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName"`
}

func (EnterAgentPushEvent) EventType() string {
	return "enter_agent" //nolint:goconst
}

func (EnterAgentPushEvent) MessageType() string {
	return "event" //nolint:goconst
}

// LocationPushEvent 成员同意上报地理位置后，每次在进入应用会话时都会上报一次地理位置。企业可以在管理端修改应用是否需要获取地理位置权限。
type LocationPushEvent struct {
	XMLName xml.Name `xml:"xml"`
	// AgentID 企业应用的id，整型。可在应用的设置页面查看
	AgentID string `xml:"AgentID"`
	// AppType app类型，在企业微信固定返回wxwork，在微信不返回该字段
	AppType string `xml:"AppType"`
	// CreateTime 消息创建时间（整型）
	CreateTime int64 `xml:"CreateTime"`
	// Event 事件类型：LOCATION
	Event string `xml:"Event"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName"`
	// Latitude 地理位置纬度
	Latitude string `xml:"Latitude"`
	// Longitude 地理位置经度
	Longitude string `xml:"Longitude"`
	// MsgType 消息类型，此时固定为：event
	MsgType string `xml:"MsgType"`
	// Precision 地理位置精度
	Precision string `xml:"Precision"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName"`
}

func (LocationPushEvent) EventType() string {
	return "LOCATION" //nolint:goconst
}

func (LocationPushEvent) MessageType() string {
	return "event" //nolint:goconst
}

// BatchJobResultPushEvent 本事件是成员在使用异步任务接口时，用于接收任务执行完毕的结果通知。
type BatchJobResultPushEvent struct {
	XMLName xml.Name `xml:"xml"`
	// CreateTime 消息创建时间（整型）
	CreateTime int64 `xml:"CreateTime"`
	// ErrCode 返回码
	ErrCode string `xml:"ErrCode"`
	// ErrMsg 对返回码的文本描述内容
	ErrMsg string `xml:"ErrMsg"`
	// Event 事件类型：batch_job_result
	Event string `xml:"Event"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName"`
	// JobID 异步任务id，最大长度为64字符
	JobID string `xml:"JobId"`
	// JobType 操作类型，字符串，目前分别有：sync_user(增量更新成员)、 replace_user(全量覆盖成员）、invite_user(邀请成员关注）、replace_party(全量覆盖部门)
	JobType string `xml:"JobType"`
	// MsgType 消息类型，此时固定为：event
	MsgType string `xml:"MsgType"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName"`
}

func (BatchJobResultPushEvent) EventType() string {
	return "batch_job_result" //nolint:goconst
}

func (BatchJobResultPushEvent) MessageType() string {
	return "event" //nolint:goconst
}

// ChangeContactCreatePartyPushEvent 当企业通过通讯录助手开通通讯录权限后，成员的变更会通知给企业。变更的事件，将推送到企业微信管理端通讯录助手中的‘接收事件服务器’。由通讯录同步助手调用接口触发的变更事件不回调通讯录同步助手本身。管理员在管理端更改组织架构或者成员信息以及企业微信的成员在客户端变更自己的个人信息将推送给通讯录同步助手。第三方通讯录变更事件参见第三方回调协议
type ChangeContactCreatePartyPushEvent struct {
	XMLName xml.Name `xml:"xml"`
	// ChangeType 此时固定为create_party
	ChangeType string `xml:"ChangeType"`
	// CreateTime 消息创建时间 （整型）
	CreateTime int64 `xml:"CreateTime"`
	// Event 事件的类型，此时固定为change_contact
	Event string `xml:"Event"`
	// FromUsername 此事件该值固定为sys，表示该消息由系统生成
	FromUsername string `xml:"FromUserName"`
	// ID 部门Id
	ID string `xml:"Id"`
	// MsgType 消息的类型，此时固定为event
	MsgType string `xml:"MsgType"`
	// Name 部门名称
	Name string `xml:"Name"`
	// Order 部门排序
	Order string `xml:"Order"`
	// ParentID 父部门id
	ParentID string `xml:"ParentId"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName"`
}

func (ChangeContactCreatePartyPushEvent) EventType() string {
	return "change_contact" //nolint:goconst
}

func (ChangeContactCreatePartyPushEvent) MessageType() string {
	return "event" //nolint:goconst
}

func (ChangeContactCreatePartyPushEvent) EventChangeType() string {
	return "create_party" //nolint:goconst
}

// ClickPushEvent 成员点击自定义菜单后，企业微信会把点击事件推送给应用。点击菜单弹出子菜单，不会产生上报。企业微信iPhone1.2.2/Android1.2.2版本开始支持菜单事件，旧版本企业微信成员点击后将没有回应，应用不能正常接收到事件推送。自定义菜单可以在管理后台的应用设置界面配置。
type ClickPushEvent struct {
	XMLName xml.Name `xml:"xml"`
	// AgentID 企业应用的id，整型。可在应用的设置页面查看
	AgentID string `xml:"AgentID"`
	// CreateTime 消息创建时间（整型）
	CreateTime int64 `xml:"CreateTime"`
	// Event 事件类型：click
	Event string `xml:"Event"`
	// EventKey 事件KEY值，与自定义菜单接口中KEY值对应
	EventKey string `xml:"EventKey"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName"`
	// MsgType 消息类型，此时固定为：event
	MsgType string `xml:"MsgType"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName"`
}

func (ClickPushEvent) EventType() string {
	return "click" //nolint:goconst
}

func (ClickPushEvent) MessageType() string {
	return "event" //nolint:goconst
}

// OpenApprovalChangePushEvent 本事件触发时机为：1.自建/第三方应用调用审批流程引擎发起申请之后，审批状态发生变化时2.自建/第三方应用调用审批流程引擎发起申请之后，在“审批中”状态，有任意审批人进行审批操作时
type OpenApprovalChangePushEvent struct {
	XMLName xml.Name `xml:"xml"`
	// AgentID 企业应用的id，整型。可在应用的设置页面查看
	AgentID string `xml:"AgentID"`
	// ApplyTime 提交申请时间
	ApplyTime string `xml:"ApplyTime"`
	// ApplyUserID 提交者userid
	ApplyUserID string `xml:"ApplyUserId"`
	// ApplyUserImage 提交者头像
	ApplyUserImage string `xml:"ApplyUserImage"`
	// ApplyUserName 提交者姓名
	ApplyUserName string `xml:"ApplyUserName"`
	// ApplyUserParty 提交者所在部门
	ApplyUserParty string `xml:"ApplyUserParty"`
	// ApprovalInfo 审批信息
	ApprovalInfo string `xml:"ApprovalInfo"`
	// ApprovalNode 审批流程信息，可以有多个审批节点
	ApprovalNode string `xml:"ApprovalNode"`
	// ApprovalNodes 审批流程信息
	ApprovalNodes string `xml:"ApprovalNodes"`
	// Approverstep 当前审批节点：0-第一个审批节点；1-第二个审批节点…以此类推
	Approverstep string `xml:"approverstep"`
	// CreateTime 消息发送时间
	CreateTime int64 `xml:"CreateTime"`
	// Event 事件名称：open_approval_change
	Event string `xml:"Event"`
	// FromUsername 发送方：企业微信
	FromUsername string `xml:"FromUserName"`
	// Item 审批节点分支，当节点为标签或上级时，一个节点可能有多个分支
	Item string `xml:"Item"`
	// ItemImage 抄送人头像
	ItemImage string `xml:"ItemImage"`
	// ItemName 抄送人姓名
	ItemName string `xml:"ItemName"`
	// ItemOpTime 分支审批人操作时间
	ItemOpTime string `xml:"ItemOpTime"`
	// Items 审批节点信息，当节点为标签或上级时，一个节点可能有多个分支
	Items string `xml:"Items"`
	// ItemSpeech 分支审批人审批意见
	ItemSpeech string `xml:"ItemSpeech"`
	// ItemStatus 分支审批审批操作状态：1-审批中；2-已同意；3-已驳回；4-已转审
	ItemStatus string `xml:"ItemStatus"`
	// ItemUserID 抄送人userid
	ItemUserID string `xml:"ItemUserId"`
	// MsgType 消息类型
	MsgType string `xml:"MsgType"`
	// NodeAttr 审批节点属性：1-或签；2-会签
	NodeAttr string `xml:"NodeAttr"`
	// NodeStatus 节点审批操作状态：1-审批中；2-已同意；3-已驳回；4-已转审
	NodeStatus string `xml:"NodeStatus"`
	// NodeType 审批节点类型：1-固定成员；2-标签；3-上级
	NodeType string `xml:"NodeType"`
	// NotifyNode 抄送人信息
	NotifyNode string `xml:"NotifyNode"`
	// NotifyNodes 抄送信息，可能有多个抄送人
	NotifyNodes string `xml:"NotifyNodes"`
	// OpenSpName 审批模板名称
	OpenSpName string `xml:"OpenSpName"`
	// OpenSpStatus 申请单当前审批状态：1-审批中；2-已通过；3-已驳回；4-已取消
	OpenSpStatus string `xml:"OpenSpStatus"`
	// OpenTemplateID 审批模板id
	OpenTemplateID string `xml:"OpenTemplateId"`
	// ThirdNo 审批单编号，由开发者在发起申请时自定义
	ThirdNo string `xml:"ThirdNo"`
	// ToUsername 接收方企业Corpid
	ToUsername string `xml:"ToUserName"`
}

func (OpenApprovalChangePushEvent) EventType() string {
	return "open_approval_change" //nolint:goconst
}

func (OpenApprovalChangePushEvent) MessageType() string {
	return "event" //nolint:goconst
}

// ShareAgentChangePushEvent 本事件触发时机为：
//
// 上级企业把自建应用共享给下级企业使用上级企业把下级企业从共享应用中移除
type ShareAgentChangePushEvent struct {
	XMLName xml.Name `xml:"xml"`
	// AgentID 上级企业应用的id，整型。可在应用的设置页面查看
	AgentID string `xml:"AgentID"`
	// CreateTime 消息发送时间
	CreateTime int64 `xml:"CreateTime"`
	// Event 事件名称：share_agent_change
	Event string `xml:"Event"`
	// FromUsername 发送方，此处固定为sys
	FromUsername string `xml:"FromUserName"`
	// MsgType 消息类型
	MsgType string `xml:"MsgType"`
	// ToUsername 上级企业CorpId
	ToUsername string `xml:"ToUserName"`
}

func (ShareAgentChangePushEvent) EventType() string {
	return "share_agent_change" //nolint:goconst
}

func (ShareAgentChangePushEvent) MessageType() string {
	return "event" //nolint:goconst
}

// TemplateCardEventPushEvent 应用下发的模板卡片消息，用户点击按钮之后触发此事件应用收到该事件之后，可以响应回复模板卡片更新消息
type TemplateCardEventPushEvent struct {
	XMLName xml.Name `xml:"xml"`
	// AgentID 企业应用的id，整型。可在应用的设置页面查看
	AgentID string `xml:"AgentID"`
	// CardType 通用模板卡片的类型，类型有”text_notice”, “news_notice”, “button_interaction”, “vote_interaction”, “multiple_interaction”五种
	CardType string `xml:"CardType"`
	// CreateTime 消息创建时间（整型）
	CreateTime int64 `xml:"CreateTime"`
	// Event 事件类型：template_card_event，点击任务卡片按钮
	Event string `xml:"Event"`
	// EventKey 与发送任务卡片消息时指定的按钮btn:key值相同
	EventKey string `xml:"EventKey"`
	// FromUsername 成员UserID
	FromUsername string `xml:"FromUserName"`
	// MsgType 消息类型，此时固定为：event
	MsgType string `xml:"MsgType"`
	// OptionIds 对应问题的选项列表
	OptionIds string `xml:"OptionIds"`
	// QuestionKey 问题的key值
	QuestionKey string `xml:"QuestionKey"`
	// ResponseCode 用于调用更新卡片接口的ResponseCode，24小时内有效，且只能使用一次
	ResponseCode string `xml:"ResponseCode"`
	// TaskID 与发送任务卡片消息时指定的task_id相同
	TaskID string `xml:"TaskId"`
	// ToUsername 企业微信CorpID
	ToUsername string `xml:"ToUserName"`
}

func (TemplateCardEventPushEvent) EventType() string {
	return "template_card_event" //nolint:goconst
}

func (TemplateCardEventPushEvent) MessageType() string {
	return "event" //nolint:goconst
}

func init() {
	RegisterEventModel(
		SubscribePushEvent{},
		EnterAgentPushEvent{},
		LocationPushEvent{},
		BatchJobResultPushEvent{},
		ChangeContactCreatePartyPushEvent{},
		ClickPushEvent{},
		OpenApprovalChangePushEvent{},
		ShareAgentChangePushEvent{},
		TemplateCardEventPushEvent{},
	)
}
