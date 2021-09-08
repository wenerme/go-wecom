package wecom

import (
	"github.com/wenerme/go-req"
)

// CreateDepartment 创建部门
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90205
func (c *Client) CreateDepartment(r *CreateDepartmentRequest, opts ...interface{}) (out CreateDepartmentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/department/create",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateDepartment 更新部门
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90206
func (c *Client) UpdateDepartment(r *UpdateDepartmentRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/department/update",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteDepartment 删除部门
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90207
func (c *Client) DeleteDepartment(r *DeleteDepartmentRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/department/delete",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListDepartment 获取部门列表
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90208
func (c *Client) ListDepartment(r *ListDepartmentRequest, opts ...interface{}) (out ListDepartmentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/department/list",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateDepartmentRequest is request of Client.CreateDepartment
type CreateDepartmentRequest struct {
	// Name 部门名称。同一个层级的部门名称不能重复。长度限制为1~32个字符，字符不能包括\:*?”&lt;&gt;｜
	Name string `json:"name"  validate:"required"`
	// NameEn 英文名称。同一个层级的部门名称不能重复。需要在管理后台开启多语言支持才能生效。长度限制为1~32个字符，字符不能包括\:*?”&lt;&gt;｜
	NameEn string `json:"name_en"  `
	// ParentID 父部门id，32位整型
	ParentID int `json:"parentid"  validate:"required"`
	// Order 在父部门中的次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
	Order int `json:"order"  `
	// ID 部门id，32位整型，指定时必须大于1。若不填该参数，将自动生成id
	ID int `json:"id"  `
}

// CreateDepartmentResponse is response of Client.CreateDepartment
type CreateDepartmentResponse struct {
	// ID 创建的部门id
	ID int `json:"id"  `
}

// UpdateDepartmentRequest is request of Client.UpdateDepartment
type UpdateDepartmentRequest struct {
	// ID 部门id
	ID int `json:"id"  validate:"required"`
	// Name 部门名称。长度限制为1~32个字符，字符不能包括\:*?”&lt;&gt;｜
	Name string `json:"name"  `
	// NameEn 英文名称，需要在管理后台开启多语言支持才能生效。长度限制为1~32个字符，字符不能包括\:*?”&lt;&gt;｜
	NameEn string `json:"name_en"  `
	// ParentID 父部门id
	ParentID int `json:"parentid"  `
	// Order 在父部门中的次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
	Order int `json:"order"  `
}

// DeleteDepartmentRequest is request of Client.DeleteDepartment
type DeleteDepartmentRequest struct {
	// ID 部门id。（注：不能删除根部门；不能删除含有子部门、成员的部门）
	ID string `json:"id"  validate:"required"`
}

// ListDepartmentRequest is request of Client.ListDepartment
type ListDepartmentRequest struct {
	// ID 部门id。获取指定部门及其下的子部门（以及及子部门的子部门等等，递归）。 如果不填，默认获取全量组织架构
	ID int `json:"id"  `
}

// ListDepartmentResponse is response of Client.ListDepartment
type ListDepartmentResponse struct {
	// Department 部门列表数据。
	Department []ListDepartmentResponseItem `json:"department"  `
}
