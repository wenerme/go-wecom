package wecom

// DefaultAPI default client base url
const DefaultAPI = "https://qyapi.weixin.qq.com"

// RootDepartmentID is 1
const RootDepartmentID = 1

// UserGenderType gender in wecom
// used by ListUserResponseItem.Gender
// used by ExternalContactResponse.Gender
type UserGenderType int

const (
	UserGenderTypeUnknown UserGenderType = 0
	UserGenderTypeMale    UserGenderType = 1
	UserGenderTypeFemale  UserGenderType = 2
)

// UserStatusType status of user
// used by ListUserResponseItem.Status
type UserStatusType int

const (
	UserStatusTypeEnabled  UserStatusType = 1 // 已激活企业微信或已关注微工作台
	UserStatusTypeForbid   UserStatusType = 2 // 已禁用
	UserStatusTypeDisabled UserStatusType = 4 // 未激活企业微信又未关注微工作台
	UserStatusTypeExit     UserStatusType = 5 // 退出企业
)

// ExternalContactType type of external contact
// used by ExternalContactResponse.Type
type ExternalContactType int

const (
	ExternalContactTypeWechat ExternalContactType = 1 // 外部联系人是微信用户
	ExternalContactTypeWecom  ExternalContactType = 2 // 外部联系人是企业微信用户
)

// ExternalContactSourceType source type of external contact
// used by GetExternalContactResponseFollowUser.AddWay
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92265#%E6%9D%A5%E6%BA%90%E5%AE%9A%E4%B9%89
type ExternalContactSourceType int

const (
	ExternalContactSourceTypeUnknown                 ExternalContactSourceType = 0   // 未知来源
	ExternalContactSourceTypeScanQR                  ExternalContactSourceType = 1   // 扫描二维码
	ExternalContactSourceTypeSearchMobile            ExternalContactSourceType = 2   // 搜索手机号
	ExternalContactSourceTypeShareContact            ExternalContactSourceType = 3   // 名片分享
	ExternalContactSourceTypeGroupChat               ExternalContactSourceType = 4   // 群聊
	ExternalContactSourceTypeMobileContact           ExternalContactSourceType = 5   // 手机通讯录
	ExternalContactSourceTypeWechatContact           ExternalContactSourceType = 6   // 微信联系人
	ExternalContactSourceTypeWechatFriendRequest     ExternalContactSourceType = 7   // 来自微信的添加好友申请
	ExternalContactSourceTypeProviderCustomerService ExternalContactSourceType = 8   // 安装第三方应用时自动添加的客服人员
	ExternalContactSourceTypeSearchEmail             ExternalContactSourceType = 9   // 搜索邮箱
	ExternalContactSourceTypeInternalShare           ExternalContactSourceType = 201 // 内部成员共享
	ExternalContactSourceTypeManagerAssignment       ExternalContactSourceType = 202 // 管理员/负责人分配
)

// ExternalContactGroupChatStatusType status type of external group chat
// used by ListGroupChatResponseGroupChatList.Status
//
type ExternalContactGroupChatStatusType int

const (
	ExternalContactGroupChatStatusTypeNormal       ExternalContactGroupChatStatusType = 0 // 跟进人正常
	ExternalContactGroupChatStatusTypeResign       ExternalContactGroupChatStatusType = 1 // 跟进人离职
	ExternalContactGroupChatStatusTypeTransferring ExternalContactGroupChatStatusType = 2 // 离职继承中
	ExternalContactGroupChatStatusTypeTransferred  ExternalContactGroupChatStatusType = 3 // 离职继承完成
)
