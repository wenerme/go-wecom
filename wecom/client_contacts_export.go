package wecom

import "github.com/wenerme/go-req"

// ExportSimpleUser 导出成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/94849
func (c *Client) ExportSimpleUser(r *ExportSimpleUserRequest) (out ExportSimpleUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/export/simple_user",
		Body:   r,
	}).Fetch(&out)
	return
}

// ExportUser 导出成员详情
//
// see https://work.weixin.qq.com/api/doc/90000/90135/94851
func (c *Client) ExportUser(r *ExportUserRequest) (out ExportUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/export/user",
		Body:   r,
	}).Fetch(&out)
	return
}

// ExportDepartment 导出部门
//
// see https://work.weixin.qq.com/api/doc/90000/90135/94852
func (c *Client) ExportDepartment(r *ExportDepartmentRequest) (out ExportDepartmentResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/export/department",
		Body:   r,
	}).Fetch(&out)
	return
}

// ExportTagUser 导出标签成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/94853
func (c *Client) ExportTagUser(r *ExportTagUserRequest) (out ExportTagUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/export/taguser",
		Body:   r,
	}).Fetch(&out)
	return
}

// ExportGetResult 接口定义
//
// see https://work.weixin.qq.com/api/doc/90000/90135/94854
func (c *Client) ExportGetResult(r *ExportGetResultRequest) (out ExportGetResultResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "GET",
		URL:    "/cgi-bin/export/get_result",
		Query:  r,
	}).Fetch(&out)
	return
}

type ExportSimpleUserRequest struct {
	// EncodingAeskey base64encode的加密密钥，长度固定为43，加密方式采用aes-256-cbc方式
	EncodingAeskey string `json:"encoding_aeskey" validate:"required"`
	// BlockSize 每块数据的人员数，支持范围[10^4,10^6]，默认值为10^6
	BlockSize int `json:"block_size"`
}

type ExportSimpleUserResponse struct {
	// JobID 任务ID，可通过获取导出结果接口查询任务结果
	JobID string `json:"jobid"`
}

type ExportUserRequest struct {
	// EncodingAeskey base64encode的加密密钥，长度固定为43，加密方式采用aes-256-cbc方式
	EncodingAeskey string `json:"encoding_aeskey" validate:"required"`
	// BlockSize 每块数据的人员数，支持范围[10^4,10^6]，默认值为10^6
	BlockSize int `json:"block_size"`
}

type ExportUserResponse struct {
	// JobID 任务ID，可通过获取导出结果接口查询任务结果
	JobID string `json:"jobid"`
}

type ExportDepartmentRequest struct {
	// EncodingAeskey base64encode的加密密钥，长度固定为43，加密方式采用aes-256-cbc方式
	EncodingAeskey string `json:"encoding_aeskey" validate:"required"`
	// BlockSize 每块数据的部门数，支持范围[10^4,10^6]，默认值为10^6
	BlockSize int `json:"block_size"`
}

type ExportDepartmentResponse struct {
	// JobID 任务ID，可通过获取导出结果接口查询任务结果
	JobID string `json:"jobid"`
}

type ExportTagUserRequest struct {
	// TagID 需要导出的标签
	TagID int `json:"tagid" validate:"required"`
	// EncodingAeskey base64encode的加密密钥，长度固定为43，加密方式采用aes-256-cbc方式
	EncodingAeskey string `json:"encoding_aeskey" validate:"required"`
	// BlockSize 每块数据的人员数和部门数之和，支持范围[10^4,10^6]，默认值为10^6
	BlockSize int `json:"block_size"`
}

type ExportTagUserResponse struct {
	// JobID 任务ID，可通过获取导出结果接口查询任务结果
	JobID string `json:"jobid"`
}

type ExportGetResultRequest struct {
	// JobID 导出任务接口成功后返回
	JobID string `json:"jobid" validate:"required"`
}

type ExportGetResultResponse struct {
	// Status 任务状态:0-未处理，1-处理中，2-完成，3-异常失败
	Status string `json:"status"`
	// DataList 数据文件列表
	DataList ExportGetResultResponseDataList `json:"data_list"`
}

type ExportGetResultResponseDataList struct {
	// URL 数据下载链接,支持指定Range头部分段下载。有效期2个小时
	URL string `json:"url"`
	// Size 密文数据大小
	Size string `json:"size"`
	// Md5 密文数据md5
	Md5 string `json:"md5"`
}
