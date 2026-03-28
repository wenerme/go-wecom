package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// GetDialRecord 获取公费电话拨打记录
// 企业可通过此接口，按时间范围拉取成功接通的公费电话拨打记录。
//
// see https://developer.work.weixin.qq.com/document/path/93662
func (c *Client) GetDialRecord(r *GetDialRecordRequest, opts ...any) (out GetDialRecordResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/dial/get_dial_record",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetDialRecordRequest is request of Client.GetDialRecord
type GetDialRecordRequest struct {
	// StartTime 查询的起始时间戳
	StartTime int `json:"start_time"`
	// EndTime 查询的结束时间戳
	EndTime int `json:"end_time"`
	// Offset 分页查询的偏移量
	Offset int `json:"offset"`
	// Limit 分页查询的每页大小，默认为100条，如该参数大于100则按100处理
	Limit int `json:"limit"`
}

// GetDialRecordResponse is response of Client.GetDialRecord
type GetDialRecordResponse struct {
	// Record 拨打记录列表
	Record []map[string]any `json:"record"`
}
