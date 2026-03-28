package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// GetDepartment 获取单个部门详情
// 获取指定部门的详细信息，包括部门ID、名称、英文名称、负责人、父部门ID及排序等。
//
// see https://developer.work.weixin.qq.com/document/path/95352
func (c *Client) GetDepartment(r *GetDepartmentRequest, opts ...any) (out GetDepartmentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/department/get",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetDepartmentRequest is request of Client.GetDepartment
type GetDepartmentRequest struct {
	// ID 部门id
	ID int `json:"id" validate:"required"`
}

// GetDepartmentResponse is response of Client.GetDepartment
type GetDepartmentResponse struct {
	// Department 部门详情对象
	Department any `json:"department"`
}

// SimpleListDepartment 获取子部门ID列表
// 获取指定部门及其下的所有子部门ID列表
//
// see https://developer.work.weixin.qq.com/document/path/95350
func (c *Client) SimpleListDepartment(r *SimpleListDepartmentRequest, opts ...any) (out SimpleListDepartmentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/department/simplelist",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SimpleListDepartmentRequest is request of Client.SimpleListDepartment
type SimpleListDepartmentRequest struct {
	// ID 部门id。获取指定部门及其下的子部门，递归。如果不填，默认获取全量组织架构
	ID int `json:"id"`
}

// SimpleListDepartmentResponse is response of Client.SimpleListDepartment
type SimpleListDepartmentResponse struct {
	// DepartmentID 部门列表数据
	DepartmentID []map[string]any `json:"department_id"`
}

// GetExportResult 获取导出结果
// 查询企业微信数据导出任务的结果状态及文件信息
//
// see https://developer.work.weixin.qq.com/document/path/94854
func (c *Client) GetExportResult(r *GetExportResultRequest, opts ...any) (out GetExportResultResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/export/get_result",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetExportResultRequest is request of Client.GetExportResult
type GetExportResultRequest struct {
	// JobID 导出任务ID，由提交导出任务接口返回
	JobID string `json:"jobid" validate:"required"`
}

// GetExportResultResponse is response of Client.GetExportResult
type GetExportResultResponse struct {
	// Status 任务状态:0-未处理，1-处理中，2-完成，3-异常失败
	Status int `json:"status"`
	// DataList 数据文件列表
	DataList []map[string]any `json:"data_list"`
}

// GetUserIdByEmail 邮箱获取userid
// 通过邮箱获取其所对应的userid
//
// see https://developer.work.weixin.qq.com/document/path/95895
func (c *Client) GetUserIdByEmail(r *GetUserIdByEmailRequest, opts ...any) (out GetUserIdByEmailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/user/get_userid_by_email",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetUserIdByEmailRequest is request of Client.GetUserIdByEmail
type GetUserIdByEmailRequest struct {
	// Email 邮箱
	Email string `json:"email" validate:"required"`
	// EmailType 邮箱类型：1-企业邮箱（默认）；2-个人邮箱
	EmailType int `json:"email_type"`
}

// GetUserIdByEmailResponse is response of Client.GetUserIdByEmail
type GetUserIdByEmailResponse struct {
	// UserID 成员UserID。注意：已升级openid的代开发或第三方，获取的是密文userid
	UserID string `json:"userid"`
}

// GetUserIDByMobile 手机号获取userid
// 通过手机号获取其所对应的userid
//
// see https://developer.work.weixin.qq.com/document/path/95402
func (c *Client) GetUserIDByMobile(r *GetUserIDByMobileRequest, opts ...any) (out GetUserIDByMobileResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/user/getuserid",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetUserIDByMobileRequest is request of Client.GetUserIDByMobile
type GetUserIDByMobileRequest struct {
	// Mobile 用户在企业微信通讯录中的手机号码。长度为5~32个字节
	Mobile string `json:"mobile" validate:"required"`
}

// GetUserIDByMobileResponse is response of Client.GetUserIDByMobile
type GetUserIDByMobileResponse struct {
	// UserID 成员UserID。对应管理端的账号，企业内必须唯一。不区分大小写，长度为1~64个字节
	UserID string `json:"userid"`
}

// ListUserId 获取成员ID列表
// 获取企业成员的userid与对应的部门ID列表
//
// see https://developer.work.weixin.qq.com/document/path/96070
func (c *Client) ListUserId(r *ListUserIdRequest, opts ...any) (out ListUserIdResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/user/list_id",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListUserIdRequest is request of Client.ListUserId
type ListUserIdRequest struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用不填
	Cursor string `json:"cursor"`
	// Limit 分页，预期请求的数据量，取值范围 1 ~ 10000
	Limit int `json:"limit"`
}

// ListUserIdResponse is response of Client.ListUserId
type ListUserIdResponse struct {
	// NextCursor 分页游标，下次请求时填写以获取之后分页的记录。如果该字段返回空则表示已没有更多数据
	NextCursor string `json:"next_cursor"`
	// DeptUser 用户-部门关系列表
	DeptUser []map[string]any `json:"dept_user"`
}
