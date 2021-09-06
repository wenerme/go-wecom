package wecom

import "github.com/wenerme/go-req"

// MessageAuditCheckSingleAgree 获取会话同意情况
// 企业可通过下述接口，获取会话中外部成员的同意情况
//
// see https://work.weixin.qq.com/api/doc/90000/90135/91782
func (c *Client) MessageAuditCheckSingleAgree(r *MessageAuditCheckSingleAgreeRequest, opts ...interface{}) (out MessageAuditCheckSingleAgreeResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/msgaudit/check_single_agree",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// MessageAuditCheckRoomAgree 获取会话同意情况
// 企业可通过下述接口，获取会话中外部成员的同意情况
//
// see https://work.weixin.qq.com/api/doc/90000/90135/91782
func (c *Client) MessageAuditCheckRoomAgree(r *MessageAuditCheckRoomAgreeRequest, opts ...interface{}) (out MessageAuditCheckRoomAgreeResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/msgaudit/check_room_agree",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

type MessageAuditCheckSingleAgreeRequest struct {
	// Info 待查询的会话信息，数组
	Info []struct {
		// UserID 内部成员的userid
		UserID string `json:"userid"  validate:"required"`
		// ExternalOpenID 外部成员的exteranalopenid
		ExternalOpenID string `json:"exteranalopenid"  validate:"required"`
	} `json:"info"  validate:"required"`
}

type MessageAuditCheckSingleAgreeResponse struct {
	// AgreeInfo 同意情况
	AgreeInfo []struct {
		// UserID 内部成员的userid
		UserID string `json:"userid"  validate:"required"`
		// ExternalOpenID 外部成员的exteranalopenid
		ExternalOpenID string `json:"exteranalopenid"  validate:"required"`
		// AgreeStatus 同意:”Agree”，不同意:”Disagree”
		AgreeStatus string `json:"agree_status"  `
		// StatusChangeTime 同意状态改变的具体时间，utc时间
		StatusChangeTime int64 `json:"status_change_time"  `
	} `json:"agreeinfo"  `
}

type MessageAuditCheckRoomAgreeRequest struct {
	// RoomID 待查询的roomid
	RoomID string `json:"roomid"  validate:"required"`
}

type MessageAuditCheckRoomAgreeResponse struct {
	// AgreeInfo 同意情况
	AgreeInfo []struct {
		// ExternalOpenID 外部成员的exteranalopenid
		ExternalOpenID string `json:"exteranalopenid"  validate:"required"`
		// AgreeStatus 同意:”Agree”，不同意:”Disagree”
		AgreeStatus string `json:"agree_status"  `
		// StatusChangeTime 同意状态改变的具体时间，utc时间
		StatusChangeTime int64 `json:"status_change_time"  `
	} `json:"agreeinfo"  `
}

type MessageAuditGetGroupChatMember struct {
	// MemberID roomid群成员的id，userid
	MemberID string `json:"memberid"  `
	// JoinTime roomid群成员的入群时间
	JoinTime int64 `json:"jointime"  `
}
