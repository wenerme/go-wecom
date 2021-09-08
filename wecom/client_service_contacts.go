package wecom

import (
	"github.com/wenerme/go-req"
)

// ProviderUploadMedia 上传需要转译的文件
// 素材上传得到media_id，该media_id仅三天内有效media_id在同一企业内应用之间可以共享
//
// see https://work.weixin.qq.com/api/doc/90001/90143/91883
func (c *Client) ProviderUploadMedia(r *ProviderUploadMediaRequest, opts ...interface{}) (out ProviderUploadMediaResponse, err error) {
	opts = append([]interface{}{WithProviderAccessToken}, opts...)
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/service/media/upload",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ProviderIDTranslateContact 异步通讯录id转译
// 通讯录id替换
//
// see https://work.weixin.qq.com/api/doc/90001/90143/91846
func (c *Client) ProviderIDTranslateContact(r *ProviderIDTranslateContactRequest, opts ...interface{}) (out ProviderIDTranslateContactResponse, err error) {
	opts = append([]interface{}{WithProviderAccessToken}, opts...)
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/service/contact/id_translate",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ProviderBatchGetResult 获取异步任务结果
//
// see https://work.weixin.qq.com/api/doc/90001/90143/91882
func (c *Client) ProviderBatchGetResult(r *ProviderBatchGetResultRequest, opts ...interface{}) (out ProviderBatchGetResultResponse, err error) {
	opts = append([]interface{}{WithProviderAccessToken}, opts...)
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/service/batch/getresult",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ProviderSortContact 通讯录userid排序
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92093
func (c *Client) ProviderSortContact(r *ProviderSortContactRequest, opts ...interface{}) (out ProviderSortContactResponse, err error) {
	opts = append([]interface{}{WithProviderAccessToken}, opts...)
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/service/contact/sort",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ProviderUploadMediaRequest is request of Client.ProviderUploadMedia
type ProviderUploadMediaRequest struct {
	Type string `json:"type"  `

	MediaID string `json:"media_id"  `

	CreatedAt string `json:"created_at"  `
}

// ProviderUploadMediaResponse is response of Client.ProviderUploadMedia
type ProviderUploadMediaResponse struct {
	// Type 媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件(file)
	Type string `json:"type"  `
	// MediaID 媒体文件上传后获取的唯一标识，3天内有效
	MediaID string `json:"media_id"  `
	// CreatedAt 媒体文件上传时间戳
	CreatedAt string `json:"created_at"  `
}

// ProviderIDTranslateContactRequest is request of Client.ProviderIDTranslateContact
type ProviderIDTranslateContactRequest struct {
	// ProviderAccessToken 服务商provider_access_token，获取方法参见服务商的凭证
	ProviderAccessToken string `json:"provider_access_token"  `
	// AuthCorpID 授权企业corpid
	AuthCorpID string `json:"auth_corpid"  validate:"required"`
	// MediaIDList 需要转译的文件的media_id列表，只支持xls/xlsx，doc/docx，csv，txt文件。不超过20个文件，获取方式参考 上传需要转译的文件
	MediaIDList []string `json:"media_id_list"  validate:"required"`
	// OutputFileName 转译完打包的文件名，不需带后缀。企业微信后台会打包成zip压缩文件，并自动拼接上.zip后缀。若media_id_list中文件个数大于1，则该字段必填。若media_id_list中文件个数等于1，且未填该字段，则转译完不打包成压缩文件。支持id转译，参见模版语法
	OutputFileName string `json:"output_file_name"  `
	// OutputFileFormat 当前仅支持输出文件为pdf格式。若media_ id_ list中文件存在相同前缀名的文件，则输出文件命名规则为：文件前缀名_ 文件格式后缀.pdf，例如：20200901_ xlsx.pdf
	OutputFileFormat string `json:"output_file_format"  `
}

// ProviderIDTranslateContactResponse is response of Client.ProviderIDTranslateContact
type ProviderIDTranslateContactResponse struct {
	// JobID 异步任务id，最大长度为64字节。jobid用于接口 获取异步任务结果 传递
	JobID string `json:"jobid"  `
}

// ProviderBatchGetResultRequest is request of Client.ProviderBatchGetResult
type ProviderBatchGetResultRequest struct {
	// ProviderAccessToken 服务商provider_access_token，获取方法参见服务商的凭证
	ProviderAccessToken string `json:"provider_access_token"  `
	// JobID 异步任务id，最大长度为64字节
	JobID string `json:"jobid"  validate:"required"`
}

// ProviderBatchGetResultResponse is response of Client.ProviderBatchGetResult
type ProviderBatchGetResultResponse struct {
	// Status 任务状态，整型，1表示任务开始，2表示任务进行中，3表示任务已完成
	Status int `json:"status"  `
	// Type 操作类型，字节串，目前有：id_translate
	Type string `json:"type"  `
	// Result 详细的处理结果。当任务完成后此字段有效
	Result ProviderBatchGetResult `json:"result"  `
}

// ProviderSortContactRequest is request of Client.ProviderSortContact
type ProviderSortContactRequest struct {
	// ProviderAccessToken 应用提供商的provider_access_token，获取方法参见服务商的凭证
	ProviderAccessToken string `json:"provider_access_token"  `
	// AuthCorpID 查询的企业corpid
	AuthCorpID string `json:"auth_corpid"  validate:"required"`
	// SortType 排序方式 0：根据姓名拼音升序排列，返回用户userid列表 1：根据姓名拼音降序排列，返回用户userid列表
	SortType string `json:"sort_type"  `
	// UserIDList 要排序的userid列表，最多支持1000个
	UserIDList string `json:"useridlist"  validate:"required"`
}

// ProviderSortContactResponse is response of Client.ProviderSortContact
type ProviderSortContactResponse struct {
	// UserIDList 排序后的userid列表
	UserIDList string `json:"useridlist"  `
}
