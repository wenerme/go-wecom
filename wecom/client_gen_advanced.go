package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// ListApplyId 批量获取申请单ID
// 自建应用可使用此接口，批量获取企业成员高级功能申请单的id信息。
//
// see https://developer.work.weixin.qq.com/document/path/99883
func (c *Client) ListApplyId(r *ListApplyIdRequest, opts ...any) (out ListApplyIdResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/advanced_feature/get_apply_id_list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListApplyIdRequest is request of Client.ListApplyId
type ListApplyIdRequest struct {
	// BusinessType 申请的高级账号类型，1-邮件 2-文档 3-微盘 4-会议
	BusinessType int `json:"business_type" validate:"required"`
	// UserID 申请的userid
	UserID string `json:"userid" validate:"required"`
	// Limit 分页查询的数据上限。默认100，最大200
	Limit int `json:"limit"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"`
	// ReqType 0-所有 1-仅api单 2-非api 申请单，默认为0
	ReqType int `json:"req_type"`
}

// ListApplyIdResponse is response of Client.ListApplyId
type ListApplyIdResponse struct {
	// NextCursor 分页请求的下一页游标
	NextCursor string `json:"next_cursor"`
	// ApplyIDList 申请id列表
	ApplyIDList []string `json:"apply_id_list"`
	// HasMore 是否还有下一页数据
	HasMore bool `json:"has_more"`
}

// SetApprovalDetail 设置审批单审批信息
// 自建应用可使用此接口，设置企业成员高级功能申请单的审批信息。
//
// see https://developer.work.weixin.qq.com/document/path/99880
func (c *Client) SetApprovalDetail(r *SetApprovalDetailRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/advanced_feature/set_approval_detail",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetApprovalDetailRequest is request of Client.SetApprovalDetail
type SetApprovalDetailRequest struct {
	// ApplyID 申请id
	ApplyID string `json:"apply_id" validate:"required"`
	// ApprovalID 审批id，注意：应用生成审批id后，审批id和申请id是一一对应的，不可改变
	ApprovalID string `json:"approval_id" validate:"required"`
	// ApprovalStatus 审批单状态：1-审批中; 2-已驳回; 3-已同意; 101-已撤销
	ApprovalStatus int `json:"approval_status" validate:"required"`
	// ApprovalURL 审批单跳转链接，须已"http://"或"https://"开头
	ApprovalURL string `json:"approval_url" validate:"required"`
	// ProcessList 审批单审批节点信息对象
	ProcessList any `json:"process_list" validate:"required"`
}
