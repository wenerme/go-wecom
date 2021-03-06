package wecom

import (
	"github.com/wenerme/go-req"
)

// ExportSimpleUser 导出成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/94849
func (c *Client) ExportSimpleUser(r *ExportSimpleUserRequest, opts ...interface{}) (out ExportSimpleUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/export/simple_user",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ExportUser 导出成员详情
//
// see https://work.weixin.qq.com/api/doc/90000/90135/94851
func (c *Client) ExportUser(r *ExportUserRequest, opts ...interface{}) (out ExportUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/export/user",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ExportDepartment 导出部门
//
// see https://work.weixin.qq.com/api/doc/90000/90135/94852
func (c *Client) ExportDepartment(r *ExportDepartmentRequest, opts ...interface{}) (out ExportDepartmentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/export/department",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ExportTagUser 导出标签成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/94853
func (c *Client) ExportTagUser(r *ExportTagUserRequest, opts ...interface{}) (out ExportTagUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/export/taguser",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ExportGetResult 接口定义
//
// see https://work.weixin.qq.com/api/doc/90000/90135/94854
func (c *Client) ExportGetResult(r *ExportGetResultRequest, opts ...interface{}) (out ExportGetResultResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/export/get_result",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ExportSimpleUserRequest is request of Client.ExportSimpleUser
type ExportSimpleUserRequest struct {
	// EncodingAesKey base64encode的加密密钥，长度固定为43，加密方式采用aes-256-cbc方式
	EncodingAesKey string `json:"encoding_aeskey"  validate:"required"`
	// BlockSize 每块数据的人员数，支持范围[10^4,10^6]，默认值为10^6
	BlockSize int `json:"block_size"  `
}

// ExportSimpleUserResponse is response of Client.ExportSimpleUser
type ExportSimpleUserResponse struct {
	// JobID 任务ID，可通过获取导出结果接口查询任务结果
	JobID string `json:"jobid"  `
}

// ExportUserRequest is request of Client.ExportUser
type ExportUserRequest struct {
	// EncodingAesKey base64encode的加密密钥，长度固定为43，加密方式采用aes-256-cbc方式
	EncodingAesKey string `json:"encoding_aeskey"  validate:"required"`
	// BlockSize 每块数据的人员数，支持范围[10^4,10^6]，默认值为10^6
	BlockSize int `json:"block_size"  `
}

// ExportUserResponse is response of Client.ExportUser
type ExportUserResponse struct {
	// JobID 任务ID，可通过获取导出结果接口查询任务结果
	JobID string `json:"jobid"  `
}

// ExportDepartmentRequest is request of Client.ExportDepartment
type ExportDepartmentRequest struct {
	// EncodingAesKey base64encode的加密密钥，长度固定为43，加密方式采用aes-256-cbc方式
	EncodingAesKey string `json:"encoding_aeskey"  validate:"required"`
	// BlockSize 每块数据的部门数，支持范围[10^4,10^6]，默认值为10^6
	BlockSize int `json:"block_size"  `
}

// ExportDepartmentResponse is response of Client.ExportDepartment
type ExportDepartmentResponse struct {
	// JobID 任务ID，可通过获取导出结果接口查询任务结果
	JobID string `json:"jobid"  `
}

// ExportTagUserRequest is request of Client.ExportTagUser
type ExportTagUserRequest struct {
	// TagID 需要导出的标签
	TagID int `json:"tagid"  validate:"required"`
	// EncodingAesKey base64encode的加密密钥，长度固定为43，加密方式采用aes-256-cbc方式
	EncodingAesKey string `json:"encoding_aeskey"  validate:"required"`
	// BlockSize 每块数据的人员数和部门数之和，支持范围[10^4,10^6]，默认值为10^6
	BlockSize int `json:"block_size"  `
}

// ExportTagUserResponse is response of Client.ExportTagUser
type ExportTagUserResponse struct {
	// JobID 任务ID，可通过获取导出结果接口查询任务结果
	JobID string `json:"jobid"  `
}

// ExportGetResultRequest is request of Client.ExportGetResult
type ExportGetResultRequest struct {
	// JobID 导出任务接口成功后返回
	JobID string `json:"jobid"  validate:"required"`
}

// ExportGetResultResponse is response of Client.ExportGetResult
type ExportGetResultResponse struct {
	// Status 任务状态:0-未处理，1-处理中，2-完成，3-异常失败
	Status string `json:"status"  `
	// DataList 数据文件列表
	DataList ExportGetResultResponseDataList `json:"data_list"  `
}

// ExportGetResultResponseDataList is model of ExportGetResultResponse.DataList
type ExportGetResultResponseDataList struct {
	// URL 数据下载链接,支持指定Range头部分段下载。有效期2个小时
	URL string `json:"url"  `
	// Size 密文数据大小
	Size int `json:"size"  `
	// Md5 密文数据md5
	Md5 string `json:"md5"  `
}
