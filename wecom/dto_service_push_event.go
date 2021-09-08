package wecom

import "encoding/xml"

// SuiteTicketPushEvent 企业微信服务器会定时（每十分钟）推送ticket。ticket会实时变更，并用于后续接口的调用。
type SuiteTicketPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// InfoType suite_ticket
	InfoType string `xml:"InfoType" json:"InfoType"`
	// SuiteID 第三方应用的SuiteId
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// SuiteTicket Ticket内容，最长为512字节
	SuiteTicket string `xml:"SuiteTicket" json:"SuiteTicket"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
}

// EventType impl EventModel
func (SuiteTicketPushEvent) EventType() string {
	return "suite_ticket" //nolint:goconst
}

// CreateAuthPushEvent 从企业微信应用市场发起授权时，企业微信后台会推送授权成功通知。
type CreateAuthPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCode 授权的auth_code,最长为512字节。用于获取企业的永久授权码。5分钟内有效
	AuthCode string `xml:"AuthCode" json:"AuthCode"`
	// InfoType create_auth
	InfoType string `xml:"InfoType" json:"InfoType"`
	// SuiteID 第三方应用的SuiteId
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
}

// EventType impl EventModel
func (CreateAuthPushEvent) EventType() string {
	return "create_auth" //nolint:goconst
}

// ChangeAuthPushEvent 当授权方（即授权企业）在企业微信管理端的授权管理中，修改了对应用的授权后，企业微信服务器推送变更授权通知。服务商接收到变更通知之后，需自行调用获取企业授权信息进行授权内容变更比对。
type ChangeAuthPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权方的corpid
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// InfoType change_auth
	InfoType string `xml:"InfoType" json:"InfoType"`
	// SuiteID 第三方应用的SuiteId
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
}

// EventType impl EventModel
func (ChangeAuthPushEvent) EventType() string {
	return "change_auth" //nolint:goconst
}

// CancelAuthPushEvent 当授权方（即授权企业）在企业微信管理端的授权管理中，取消了对应用的授权托管后，企业微信后台会推送取消授权通知。
type CancelAuthPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权方企业的corpid
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// InfoType cancel_auth
	InfoType string `xml:"InfoType" json:"InfoType"`
	// SuiteID 第三方应用的SuiteId
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
}

// EventType impl EventModel
func (CancelAuthPushEvent) EventType() string {
	return "cancel_auth" //nolint:goconst
}

type ChangeContactCreateUserPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// Avatar 头像url。注：如果要获取小图将url最后的”/0”改成”/100”即可，仅通讯录管理应用可获取
	Avatar string `xml:"Avatar" json:"Avatar"`
	// ChangeType 固定为create_user
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// Department 更新后成员所在部门列表
	Department string `xml:"Department" json:"Department"`
	// Email 邮箱，仅通讯录管理应用可获取
	Email string `xml:"Email" json:"Email"`
	// EnglishName 英文名
	EnglishName string `xml:"EnglishName" json:"EnglishName"`
	// ExtAttr 扩展属性，仅通讯录管理应用可获取
	ExtAttr string `xml:"ExtAttr" json:"ExtAttr"`
	// Gender 性别。1表示男性，2表示女性
	Gender string `xml:"Gender" json:"Gender"`
	// InfoType 固定为change_contact
	InfoType string `xml:"InfoType" json:"InfoType"`
	// IsLeader 上级字段，标识是否为上级。0表示普通成员，1表示上级。仅通讯录管理应用可获取
	IsLeader string `xml:"IsLeader" json:"IsLeader"`
	// Mobile 手机号码，仅通讯录管理应用可获取
	Mobile string `xml:"Mobile" json:"Mobile"`
	// Name 成员名称
	Name string `xml:"Name" json:"Name"`
	// Position 职位信息。长度为0~64个字节，仅通讯录管理应用可获取
	Position string `xml:"Position" json:"Position"`
	// Status 激活状态：1&#x3D;激活或关注， 2&#x3D;禁用， 4&#x3D;未激活 已激活代表已激活企业微信或已关注微工作台（原企业号）。未激活代表既未激活企业微信又未关注微工作台（原企业号）
	Status string `xml:"Status" json:"Status"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Telephone 座机，仅通讯录管理应用可获取
	Telephone string `xml:"Telephone" json:"Telephone"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
	// UserID 成员UserID
	UserID string `xml:"UserID" json:"UserID"`
}

// EventType impl EventModel
func (ChangeContactCreateUserPushEvent) EventType() string {
	return "change_contact" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeContactCreateUserPushEvent) EventChangeType() string {
	return "create_user" //nolint:goconst
}

type ChangeContactUpdateUserPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// Avatar 头像url。注：如果要获取小图将url最后的”/0”改成”/100”即可。变更时推送，仅通讯录管理应用可获取
	Avatar string `xml:"Avatar" json:"Avatar"`
	// ChangeType 固定为update_user
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// Department 更新后成员所在部门列表
	Department string `xml:"Department" json:"Department"`
	// Email 邮箱，变更时推送 ，仅通讯录应用可获取
	Email string `xml:"Email" json:"Email"`
	// EnglishName 英文名
	EnglishName string `xml:"EnglishName" json:"EnglishName"`
	// ExtAttr 扩展属性，变更时推送，仅通讯录应用可获取
	ExtAttr string `xml:"ExtAttr" json:"ExtAttr"`
	// Gender 性别，变更时推送。1表示男性，2表示女性
	Gender string `xml:"Gender" json:"Gender"`
	// InfoType 固定为change_contact
	InfoType string `xml:"InfoType" json:"InfoType"`
	// IsLeader 上级字段，标识是否为上级。0表示普通成员，1表示上级。仅通讯录管理应用可获取
	IsLeader string `xml:"IsLeader" json:"IsLeader"`
	// Mobile 手机号码，变更时推送，仅通讯录应用可获取
	Mobile string `xml:"Mobile" json:"Mobile"`
	// Name 成员名称，变更时推送
	Name string `xml:"Name" json:"Name"`
	// NewUserID 新的UserID，变更时推送（userid由系统生成时可更改一次）
	NewUserID string `xml:"NewUserID" json:"NewUserID"`
	// Position 职位信息。长度为0~64个字节，仅通讯录管理应用可获取
	Position string `xml:"Position" json:"Position"`
	// Status 激活状态：1&#x3D;激活或关注， 2&#x3D;禁用， 4&#x3D;未激活（重新启用未激活用户或者退出企业并且取消关注时触发）
	Status string `xml:"Status" json:"Status"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Telephone 座机，仅通讯录应用可获取
	Telephone string `xml:"Telephone" json:"Telephone"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
	// UserID 变更信息的成员UserID
	UserID string `xml:"UserID" json:"UserID"`
}

// EventType impl EventModel
func (ChangeContactUpdateUserPushEvent) EventType() string {
	return "change_contact" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeContactUpdateUserPushEvent) EventChangeType() string {
	return "update_user" //nolint:goconst
}

type ChangeContactDeleteUserPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 固定为delete_user
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// InfoType 固定为change_contact
	InfoType string `xml:"InfoType" json:"InfoType"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
	// UserID 变更信息的成员UserID
	UserID string `xml:"UserID" json:"UserID"`
}

// EventType impl EventModel
func (ChangeContactDeleteUserPushEvent) EventType() string {
	return "change_contact" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeContactDeleteUserPushEvent) EventChangeType() string {
	return "delete_user" //nolint:goconst
}

type ProviderChangeContactCreatePartyPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 固定为create_party
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// ID 部门Id
	ID string `xml:"Id" json:"Id"`
	// InfoType 固定为change_contact
	InfoType string `xml:"InfoType" json:"InfoType"`
	// Name 部门名称
	Name string `xml:"Name" json:"Name"`
	// Order 部门排序
	Order string `xml:"Order" json:"Order"`
	// ParentID 父部门id
	ParentID string `xml:"ParentId" json:"ParentId"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
}

// EventType impl EventModel
func (ProviderChangeContactCreatePartyPushEvent) EventType() string {
	return "change_contact" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ProviderChangeContactCreatePartyPushEvent) EventChangeType() string {
	return "create_party" //nolint:goconst
}

type ChangeContactUpdatePartyPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 固定为update_party
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// ID 部门Id
	ID string `xml:"Id" json:"Id"`
	// InfoType 固定为change_contact
	InfoType string `xml:"InfoType" json:"InfoType"`
	// Name 部门名称，仅当该字段发生变更时传递
	Name string `xml:"Name" json:"Name"`
	// ParentID 父部门id，仅当该字段发生变更时传递
	ParentID string `xml:"ParentId" json:"ParentId"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
}

// EventType impl EventModel
func (ChangeContactUpdatePartyPushEvent) EventType() string {
	return "change_contact" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeContactUpdatePartyPushEvent) EventChangeType() string {
	return "update_party" //nolint:goconst
}

type ChangeContactDeletePartyPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 固定为delete_party
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// ID 部门Id
	ID string `xml:"Id" json:"Id"`
	// InfoType 固定为change_contact
	InfoType string `xml:"InfoType" json:"InfoType"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
}

// EventType impl EventModel
func (ChangeContactDeletePartyPushEvent) EventType() string {
	return "change_contact" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeContactDeletePartyPushEvent) EventChangeType() string {
	return "delete_party" //nolint:goconst
}

type ChangeContactUpdateTagPushEvent struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	// AddPartyItems 标签中新增的部门id列表，用逗号分隔
	AddPartyItems string `xml:"AddPartyItems" json:"AddPartyItems"`
	// AddUserItems 标签中新增的成员userid列表，用逗号分隔
	AddUserItems string `xml:"AddUserItems" json:"AddUserItems"`
	// AuthCorpID 授权企业的CorpID
	AuthCorpID string `xml:"AuthCorpId" json:"AuthCorpId"`
	// ChangeType 固定为update_tag
	ChangeType string `xml:"ChangeType" json:"ChangeType"`
	// DelPartyItems 标签中删除的部门id列表，用逗号分隔
	DelPartyItems string `xml:"DelPartyItems" json:"DelPartyItems"`
	// DelUserItems 标签中删除的成员userid列表，用逗号分隔
	DelUserItems string `xml:"DelUserItems" json:"DelUserItems"`
	// InfoType 固定为change_contact
	InfoType string `xml:"InfoType" json:"InfoType"`
	// SuiteID 第三方应用ID
	SuiteID string `xml:"SuiteId" json:"SuiteId"`
	// TagID 标签Id
	TagID string `xml:"TagId" json:"TagId"`
	// Timestamp 时间戳
	Timestamp int64 `xml:"TimeStamp" json:"TimeStamp"`
}

// EventType impl EventModel
func (ChangeContactUpdateTagPushEvent) EventType() string {
	return "change_contact" //nolint:goconst
}

// EventChangeType impl EventChangeModel
func (ChangeContactUpdateTagPushEvent) EventChangeType() string {
	return "update_tag" //nolint:goconst
}

func init() {
	RegisterEventModel(
		SuiteTicketPushEvent{},
		CreateAuthPushEvent{},
		ChangeAuthPushEvent{},
		CancelAuthPushEvent{},
		ChangeContactCreateUserPushEvent{},
		ChangeContactUpdateUserPushEvent{},
		ChangeContactDeleteUserPushEvent{},
		ChangeContactCreatePartyPushEvent{},
		ChangeContactUpdatePartyPushEvent{},
		ChangeContactDeletePartyPushEvent{},
		ChangeContactUpdateTagPushEvent{},
	)
}
