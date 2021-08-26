package wecom

import "github.com/wenerme/go-req"

// SimpleUser 导出成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/94849
func (c *Client) SimpleUser(r *SimpleUserRequest) (out SimpleUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/export/simple_user",
		Body:   r,
	}).Fetch(&out)
	return
}

// User 导出成员详情
//
// see https://work.weixin.qq.com/api/doc/90000/90135/94851
func (c *Client) User(r *UserRequest) (out UserResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/export/user",
		Body:   r,
	}).Fetch(&out)
	return
}

// Department 导出部门
//
// see https://work.weixin.qq.com/api/doc/90000/90135/94852
func (c *Client) Department(r *DepartmentRequest) (out DepartmentResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/export/department",
		Body:   r,
	}).Fetch(&out)
	return
}

// Taguser 导出标签成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/94853
func (c *Client) Taguser(r *TaguserRequest) (out TaguserResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/export/taguser",
		Body:   r,
	}).Fetch(&out)
	return
}

// GetResult 接口定义
//
// see https://work.weixin.qq.com/api/doc/90000/90135/94854
func (c *Client) GetResult(r *GetResultRequest) (out GetResultResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "GET",
		URL:    "/cgi-bin/export/get_result",
		Query:  r,
	}).Fetch(&out)
	return
}

type SimpleUserRequest struct {
	// EncodingAeskey base64encode的加密密钥，长度固定为43，加密方式采用aes-256-cbc方式
	EncodingAeskey string `json:"encoding_aeskey"`
	// BlockSize 每块数据的人员数，支持范围[10^4,10^6]，默认值为10^6
	BlockSize int `json:"block_size"`
}

type SimpleUserResponse struct {
	// JobID 任务ID，可通过获取导出结果接口查询任务结果
	JobID string `json:"jobid"`
}

type UserRequest struct {
	// EncodingAeskey base64encode的加密密钥，长度固定为43，加密方式采用aes-256-cbc方式
	EncodingAeskey string `json:"encoding_aeskey"`
	// BlockSize 每块数据的人员数，支持范围[10^4,10^6]，默认值为10^6
	BlockSize int `json:"block_size"`
}

type UserResponse struct {
	// JobID 任务ID，可通过获取导出结果接口查询任务结果
	JobID string `json:"jobid"`
}

type DepartmentRequest struct {
	// EncodingAeskey base64encode的加密密钥，长度固定为43，加密方式采用aes-256-cbc方式
	EncodingAeskey string `json:"encoding_aeskey"`
	// BlockSize 每块数据的部门数，支持范围[10^4,10^6]，默认值为10^6
	BlockSize int `json:"block_size"`
}

type DepartmentResponse struct {
	// JobID 任务ID，可通过获取导出结果接口查询任务结果
	JobID string `json:"jobid"`
}

type TaguserRequest struct {
	// TagID 需要导出的标签
	TagID int `json:"tagid"`
	// EncodingAeskey base64encode的加密密钥，长度固定为43，加密方式采用aes-256-cbc方式
	EncodingAeskey string `json:"encoding_aeskey"`
	// BlockSize 每块数据的人员数和部门数之和，支持范围[10^4,10^6]，默认值为10^6
	BlockSize int `json:"block_size"`
}

type TaguserResponse struct {
	// JobID 任务ID，可通过获取导出结果接口查询任务结果
	JobID string `json:"jobid"`
}

type GetResultRequest struct {
	// JobID 导出任务接口成功后返回
	JobID string `json:"jobid"`
}

type GetResultResponse struct {
	// Status 任务状态:0-未处理，1-处理中，2-完成，3-异常失败
	Status string `json:"status"`
	// DataList 数据文件列表
	DataList GetResultResponseDataList `json:"data_list"`
}

type GetResultResponseDataList struct {
	// URL 数据下载链接,支持指定Range头部分段下载。有效期2个小时
	URL string `json:"url"`
	// Size 密文数据大小
	Size string `json:"size"`
	// Md5 密文数据md5
	Md5 string `json:"md5"`
}
