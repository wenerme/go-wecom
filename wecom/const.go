package wecom

import (
	"bytes"
	"strconv"
)

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

func (v *UserGenderType) UnmarshalJSON(data []byte) error {
	l := len(data)
	if l == 0 {
		return nil
	}
	switch {
	case l == 0:
		return nil
	case l == 4 && bytes.Equal(data, []byte("null")):
		return nil
	case l >= 3 && data[0] == '"' && data[l-1] == '"':
		data = data[1 : l-1]
	}
	val, err := strconv.Atoi(string(data))
	if err == nil {
		*v = UserGenderType(val)
	}
	return err
}

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

// AgentAuthType type of agent auth
// used by ProviderGetLoginInfoResponse.Agent
//
type AgentAuthType int

const (
	AgentAuthTypeUsage = 0
	AgentAuthTypeAdmin = 1
)

func (v AgentAuthType) String() string {
	switch v {
	case 0:
		return "usage"
	case 1:
		return "admin"
	default:
		return strconv.Itoa(int(v))
	}
}

// ExternalContactAddWay 添加客户的来源
// used by GetExternalContactResponseFollowUser.AddWay
//
// see https://open.work.weixin.qq.com/api/doc/90000/90135/92114
type ExternalContactAddWay int

const (
	ExternalContactAddWayUnknown             ExternalContactAddWay = 0   // 未知来源
	ExternalContactAddWayScanQR              ExternalContactAddWay = 1   //	扫描二维码
	ExternalContactAddWaySearchPhone         ExternalContactAddWay = 2   // 搜索手机号
	ExternalContactAddWayShareContactCard    ExternalContactAddWay = 3   // 名片分享
	ExternalContactAddWayGroup               ExternalContactAddWay = 4   // 群聊
	ExternalContactAddWayPhoneContact        ExternalContactAddWay = 5   //	手机通讯录
	ExternalContactAddWayWeChatContact       ExternalContactAddWay = 6   // 微信联系人
	ExternalContactAddWayWeChatFriendRequest ExternalContactAddWay = 7   // 来自微信的添加好友申请
	ExternalContactAddWayCustomService       ExternalContactAddWay = 8   // 安装第三方应用时自动添加的客服人员
	ExternalContactAddWaySearchEMail         ExternalContactAddWay = 9   // 搜索邮箱
	ExternalContactAddWayChannel             ExternalContactAddWay = 10  // 视频号主页添加
	ExternalContactAddWayInternalShare       ExternalContactAddWay = 201 // 内部成员共享
	ExternalContactAddWayAssign              ExternalContactAddWay = 202 // 管理员/负责人分配
)
