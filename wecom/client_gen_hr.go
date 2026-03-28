package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// GetEmployeeFieldConfig 获取员工字段配置
// 通过这个接口获取员工字段配置信息
//
// see https://developer.work.weixin.qq.com/document/path/99131
func (c *Client) GetEmployeeFieldConfig(r *GetEmployeeFieldConfigRequest, opts ...any) (out GetEmployeeFieldConfigResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/hr/get_fields",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetEmployeeFieldConfigRequest is request of Client.GetEmployeeFieldConfig
type GetEmployeeFieldConfigRequest struct{}

// GetEmployeeFieldConfigResponse is response of Client.GetEmployeeFieldConfig
type GetEmployeeFieldConfigResponse struct {
	// GroupList 字段组的配置信息
	GroupList []map[string]any `json:"group_list"`
}

// GetStaffInfo 获取员工花名册信息
// 通过这个接口获取指定员工的花名册信息
//
// see https://developer.work.weixin.qq.com/document/path/99132
func (c *Client) GetStaffInfo(r *GetStaffInfoRequest, opts ...any) (out GetStaffInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/hr/get_staff_info",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetStaffInfoRequest is request of Client.GetStaffInfo
type GetStaffInfoRequest struct {
	// UserID 需要获取花名册信息的员工的userid
	UserID string `json:"userid" validate:"required"`
	// GetAll 是否获取全部字段信息，不填时默认为否
	GetAll bool `json:"get_all"`
	// Fieldids 需要获取的字段信息
	Fieldids []map[string]any `json:"fieldids"`
}

// GetStaffInfoResponse is response of Client.GetStaffInfo
type GetStaffInfoResponse struct {
	// FieldInfo 获取到的字段信息
	FieldInfo []map[string]any `json:"field_info"`
}

// UpdateStaffInfo 更新员工花名册信息
// 通过这个接口更新指定员工的花名册信息
//
// see https://developer.work.weixin.qq.com/document/path/99133
func (c *Client) UpdateStaffInfo(r *UpdateStaffInfoRequest, opts ...any) (out UpdateStaffInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/hr/update_staff_info",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateStaffInfoRequest is request of Client.UpdateStaffInfo
type UpdateStaffInfoRequest struct {
	// UserID 需要更新花名册信息的员工的userid
	UserID string `json:"userid" validate:"required"`
	// UpdateItems 需要更新、增加或清空单个字段的内容列表
	UpdateItems []map[string]any `json:"update_items"`
	// RemoveItems 可重复的字段组中需要整组删除的字段组列表
	RemoveItems []map[string]any `json:"remove_items"`
	// InsertItems 可重复的字段组中需要增加一组字段的字段组列表
	InsertItems []map[string]any `json:"insert_items"`
}

// UpdateStaffInfoResponse is response of Client.UpdateStaffInfo
type UpdateStaffInfoResponse struct {
	// UpdateResults 更新字段的结果列表
	UpdateResults []map[string]any `json:"update_results"`
	// RemoveResults 删除字段组的结果列表
	RemoveResults []map[string]any `json:"remove_results"`
	// InsertResults 增加字段组的结果列表
	InsertResults []map[string]any `json:"insert_results"`
}
