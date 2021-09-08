package wecom

import "encoding/xml"

// ChangeExternalContactAddExternalContactPushEvent 添加企业客户事件
// 授权企业中配置了客户联系功能的成员添加外部联系人时，企业微信服务器会向应用的“指令回调URL”推送该事件
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92277
type ChangeExternalContactAddExternalContactPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 此时固定为add_external_contact
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `xml:"ExternalUserID" json:"ExternalUserID"`
	// InfoType 固定为change_external_contact
	InfoType string `xml:"InfoType" json:"InfoType"`
	// State 添加此用户的「联系我」方式配置的state参数，可用于识别添加此用户的渠道
	State string `xml:"State" json:"State"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
	// UserID 企业服务人员的UserID
	UserID string `xml:"UserID" json:"UserID"`
	// WelcomeCode 欢迎语code，可用于发送欢迎语
	WelcomeCode string `xml:"WelcomeCode" json:"WelcomeCode"`
}

// EventType impl EventModel
func (ChangeExternalContactAddExternalContactPushEvent) EventType() string {
	return "change_external_contact" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeExternalContactAddExternalContactPushEvent) EventChangeType() string {
	return "add_external_contact" //nolint:goconst
}

// ChangeExternalContactEditExternalContactPushEvent 编辑企业客户事件
// 授权企业中配置了客户联系功能的成员编辑外部联系人的备注信息(不包括备注手机号码)或企业标签时，企业微信服务器会向应用的“指令回调URL”推送该事件，但仅修改外部联系人备注手机号时不会触发回调。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92277
type ChangeExternalContactEditExternalContactPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 此时固定为edit_external_contact
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `xml:"ExternalUserID" json:"ExternalUserID"`
	// InfoType 固定为change_external_contact
	InfoType string `xml:"InfoType" json:"InfoType"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
	// UserID 企业服务人员的UserID
	UserID string `xml:"UserID" json:"UserID"`
}

// EventType impl EventModel
func (ChangeExternalContactEditExternalContactPushEvent) EventType() string {
	return "change_external_contact" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeExternalContactEditExternalContactPushEvent) EventChangeType() string {
	return "edit_external_contact" //nolint:goconst
}

// ChangeExternalContactAddHalfExternalContactPushEvent 外部联系人免验证添加成员事件
// 外部联系人添加了配置了客户联系功能且开启了免验证的成员时（此时成员尚未确认添加对方为好友），企业微信服务器会向应用的“指令回调URL”推送该事件
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92277
type ChangeExternalContactAddHalfExternalContactPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 此时固定为add_half_external_contact
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `xml:"ExternalUserID" json:"ExternalUserID"`
	// InfoType 固定为change_external_contact
	InfoType string `xml:"InfoType" json:"InfoType"`
	// State 添加此用户的「联系我」方式配置的state参数，可用于识别添加此用户的渠道
	State string `xml:"State" json:"State"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
	// UserID 企业服务人员的UserID
	UserID string `xml:"UserID" json:"UserID"`
	// WelcomeCode 欢迎语code，可用于发送欢迎语
	WelcomeCode string `xml:"WelcomeCode" json:"WelcomeCode"`
}

// EventType impl EventModel
func (ChangeExternalContactAddHalfExternalContactPushEvent) EventType() string {
	return "change_external_contact" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeExternalContactAddHalfExternalContactPushEvent) EventChangeType() string {
	return "add_half_external_contact" //nolint:goconst
}

// ChangeExternalContactDelExternalContactPushEvent 删除企业客户事件
// 授权企业中配置了客户联系功能的成员删除外部联系人时，企业微信服务器会向应用的“指令回调URL”推送该事件
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92277
type ChangeExternalContactDelExternalContactPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 此时固定为del_external_contact
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `xml:"ExternalUserID" json:"ExternalUserID"`
	// InfoType 固定为change_external_contact
	InfoType string `xml:"InfoType" json:"InfoType"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
	// UserID 企业服务人员的UserID
	UserID string `xml:"UserID" json:"UserID"`
}

// EventType impl EventModel
func (ChangeExternalContactDelExternalContactPushEvent) EventType() string {
	return "change_external_contact" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeExternalContactDelExternalContactPushEvent) EventChangeType() string {
	return "del_external_contact" //nolint:goconst
}

// ChangeExternalContactDelFollowUserPushEvent 删除跟进成员事件
// 授权企业中配置了客户联系功能的成员被外部联系人删除时，企业微信服务器会向应用的“指令回调URL”推送该事件
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92277
type ChangeExternalContactDelFollowUserPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 此时固定为del_follow_user
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `xml:"ExternalUserID" json:"ExternalUserID"`
	// InfoType 固定为change_external_contact
	InfoType string `xml:"InfoType" json:"InfoType"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
	// UserID 企业服务人员的UserID
	UserID string `xml:"UserID" json:"UserID"`
}

// EventType impl EventModel
func (ChangeExternalContactDelFollowUserPushEvent) EventType() string {
	return "change_external_contact" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeExternalContactDelFollowUserPushEvent) EventChangeType() string {
	return "del_follow_user" //nolint:goconst
}

// ChangeExternalContactTransferFailPushEvent 客户接替失败事件
// 企业将客户分配给新的成员接替后，当客户添加失败时，企业微信服务器会向应用的“指令回调URL”推送该事件
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92277
type ChangeExternalContactTransferFailPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 此时固定为transfer_fail
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `xml:"ExternalUserID" json:"ExternalUserID"`
	// FailReason 接替失败的原因, customer_refused-客户拒绝， customer_limit_exceed-接替成员的客户数达到上限
	FailReason string `xml:"FailReason" json:"FailReason"`
	// InfoType 固定为change_external_contact
	InfoType string `xml:"InfoType" json:"InfoType"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
	// UserID 企业服务人员的UserID
	UserID string `xml:"UserID" json:"UserID"`
}

// EventType impl EventModel
func (ChangeExternalContactTransferFailPushEvent) EventType() string {
	return "change_external_contact" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeExternalContactTransferFailPushEvent) EventChangeType() string {
	return "transfer_fail" //nolint:goconst
}

// ChangeExternalChatCreatePushEvent 客户群创建事件
// 有新增客户群时，回调该事件。收到该事件后，企业可以调用获取客户群详情接口获取客户群详情。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92277
type ChangeExternalChatCreatePushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 此时固定为create
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// ChatID 群ID
	ChatID string `xml:"ChatId" json:"ChatId"`
	// InfoType 固定为change_external_chat
	InfoType string `xml:"InfoType" json:"InfoType"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
}

// EventType impl EventModel
func (ChangeExternalChatCreatePushEvent) EventType() string {
	return "change_external_chat" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeExternalChatCreatePushEvent) EventChangeType() string {
	return "create" //nolint:goconst
}

// ChangeExternalChatUpdatePushEvent 客户群变更事件
// 客户群被修改后（群名变更，群成员增加或移除，群公告变更），回调该事件。收到该事件后，企业需要再调用获取客户群详情接口，以获取最新的群详情。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92277
type ChangeExternalChatUpdatePushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 此时固定为update
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// ChatID 群ID
	ChatID string `xml:"ChatId" json:"ChatId"`
	// InfoType 固定为change_external_chat
	InfoType string `xml:"InfoType" json:"InfoType"`
	// JoinScene 当是成员入群时有值。表示成员的入群方式0 - 由成员邀请入群（包括直接邀请入群和通过邀请链接入群）3 - 通过扫描群二维码入群
	JoinScene string `xml:"JoinScene" json:"JoinScene"`
	// MemChangeCnt 当是成员入群或退群时有值。表示成员变更数量
	MemChangeCnt string `xml:"MemChangeCnt" json:"MemChangeCnt"`
	// QuitScene 当是成员退群时有值。表示成员的退群方式0 - 自己退群1 - 群主/群管理员移出
	QuitScene string `xml:"QuitScene" json:"QuitScene"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
	// UpdateDetail 变更详情。目前有以下几种：add_member : 成员入群del_member : 成员退群 change_owner : 群主变更change_name : 群名变更change_notice : 群公告变更
	UpdateDetail string `xml:"UpdateDetail" json:"UpdateDetail"`
}

// EventType impl EventModel
func (ChangeExternalChatUpdatePushEvent) EventType() string {
	return "change_external_chat" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeExternalChatUpdatePushEvent) EventChangeType() string {
	return "update" //nolint:goconst
}

// ChangeExternalChatDismissPushEvent 客户群解散事件
// 当客户群被群主解散后，回调该事件。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92277
type ChangeExternalChatDismissPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 此时固定为dismiss
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// ChatID 群ID
	ChatID string `xml:"ChatId" json:"ChatId"`
	// InfoType 固定为change_external_chat
	InfoType string `xml:"InfoType" json:"InfoType"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
}

// EventType impl EventModel
func (ChangeExternalChatDismissPushEvent) EventType() string {
	return "change_external_chat" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeExternalChatDismissPushEvent) EventChangeType() string {
	return "dismiss" //nolint:goconst
}

// ChangeExternalTagCreatePushEvent 企业客户标签创建事件
// 企业/管理员创建客户标签/标签组时，企业微信服务器将向具有企业客户权限的第三方应用指令回调URL回调此事件。收到该事件后，第三方服务商需要调用获取企业标签库来获取标签/标签组的详细信息。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92277
type ChangeExternalTagCreatePushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 此时固定为create
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// ID 标签或标签组的ID
	ID string `xml:"Id" json:"Id"`
	// InfoType 固定为change_external_tag
	InfoType string `xml:"InfoType" json:"InfoType"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// TagType 创建标签时，此项为tag，创建标签组时，此项为tag_group
	TagType string `xml:"TagType" json:"TagType"`
	// Timestamp unix时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
}

// EventType impl EventModel
func (ChangeExternalTagCreatePushEvent) EventType() string {
	return "change_external_tag" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeExternalTagCreatePushEvent) EventChangeType() string {
	return "create" //nolint:goconst
}

// ChangeExternalTagUpdatePushEvent 企业客户标签变更事件
// 当企业客户标签/标签组被修改时，企业微信服务器将向具有企业客户权限的第三方应用指令回调URL回调此事件。收到该事件后，第三方服务商需要调用获取企业标签库来获取标签/标签组的详细信息。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92277
type ChangeExternalTagUpdatePushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 此时固定为update
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// ID 标签或标签组的ID
	ID string `xml:"Id" json:"Id"`
	// InfoType 固定为change_external_tag
	InfoType string `xml:"InfoType" json:"InfoType"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// TagType 变更标签时，此项为tag，变更标签组时，此项为tag_group
	TagType string `xml:"TagType" json:"TagType"`
	// Timestamp unix时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
}

// EventType impl EventModel
func (ChangeExternalTagUpdatePushEvent) EventType() string {
	return "change_external_tag" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeExternalTagUpdatePushEvent) EventChangeType() string {
	return "update" //nolint:goconst
}

// ChangeExternalTagDeletePushEvent 企业客户标签删除事件
// 当企业客户标签/标签组被删除改时，企业微信服务器将向具有企业客户权限的第三方应用指令回调URL回调此事件。删除标签组时，该标签组下的所有标签将被同时删除，但不会进行回调。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92277
type ChangeExternalTagDeletePushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 此时固定为delete
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// ID 标签或标签组的ID
	ID string `xml:"Id" json:"Id"`
	// InfoType 固定为change_external_tag
	InfoType string `xml:"InfoType" json:"InfoType"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// TagType 删除标签时，此项为tag，删除标签组时，此项为tag_group
	TagType string `xml:"TagType" json:"TagType"`
	// Timestamp unix时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
}

// EventType impl EventModel
func (ChangeExternalTagDeletePushEvent) EventType() string {
	return "change_external_tag" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeExternalTagDeletePushEvent) EventChangeType() string {
	return "delete" //nolint:goconst
}

func init() {
	RegisterEventModel(
		ChangeExternalContactAddExternalContactPushEvent{},
		ChangeExternalContactEditExternalContactPushEvent{},
		ChangeExternalContactAddHalfExternalContactPushEvent{},
		ChangeExternalContactDelExternalContactPushEvent{},
		ChangeExternalContactDelFollowUserPushEvent{},
		ChangeExternalContactTransferFailPushEvent{},
		ChangeExternalChatCreatePushEvent{},
		ChangeExternalChatUpdatePushEvent{},
		ChangeExternalChatDismissPushEvent{},
		ChangeExternalTagCreatePushEvent{},
		ChangeExternalTagUpdatePushEvent{},
		ChangeExternalTagDeletePushEvent{},
	)
}
