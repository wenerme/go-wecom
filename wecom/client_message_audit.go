package wecom

import (
	"github.com/wenerme/go-req"
)

// MessageAuditGetPermitUserList 获取会话内容存档开启成员列表
// 企业可通过此接口，获取企业开启会话内容存档的成员列表
//
// see https://work.weixin.qq.com/api/doc/90000/90135/91614
func (c *Client) MessageAuditGetPermitUserList(r *MessageAuditGetPermitUserListRequest, opts ...interface{}) (out MessageAuditGetPermitUserListResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/msgaudit/get_permit_user_list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// MessageAuditGetGroupChat 获取会话内容存档内部群信息
// 企业可通过此接口，获取会话内容存档本企业的内部群信息，包括群名称、群主id、公告、群创建时间以及所有群成员的id与加入时间。
//
// see https://work.weixin.qq.com/api/doc/90000/90135/92951
func (c *Client) MessageAuditGetGroupChat(r *MessageAuditGetGroupChatRequest, opts ...interface{}) (out MessageAuditGetGroupChatResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/msgaudit/groupchat/get",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// MessageAuditGetPermitUserListRequest is request of Client.MessageAuditGetPermitUserList
type MessageAuditGetPermitUserListRequest struct {
	// Type 拉取对应版本的开启成员列表。1表示办公版；2表示服务版；3表示企业版。非必填，不填写的时候返回全量成员列表。
	Type int `json:"type"  `
}

// MessageAuditGetPermitUserListResponse is response of Client.MessageAuditGetPermitUserList
type MessageAuditGetPermitUserListResponse struct {
	// Ids 设置在开启范围内的成员的userid列表
	Ids []string `json:"ids"  `
}

// MessageAuditGetGroupChatRequest is request of Client.MessageAuditGetGroupChat
type MessageAuditGetGroupChatRequest struct {
	// RoomID 待查询的群id
	RoomID string `json:"roomid"  validate:"required"`
}

// MessageAuditGetGroupChatResponse is response of Client.MessageAuditGetGroupChat
type MessageAuditGetGroupChatResponse struct {
	// RoomName roomid对应的群名称
	RoomName string `json:"roomname"  `
	// Creator roomid对应的群创建者，userid
	Creator string `json:"creator"  `
	// RoomCreateTime roomid对应的群创建时间
	RoomCreateTime int `json:"room_create_time"  `
	// Notice roomid对应的群公告
	Notice string `json:"notice"  `
	// Members roomid对应的群成员列表
	Members []MessageAuditGetGroupChatMember `json:"members"  `
}
