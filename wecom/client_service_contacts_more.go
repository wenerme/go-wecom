package wecom

import "encoding/xml"

type ProviderBatchGetResult struct {
	ContactIDTranslate struct {
		// URL 通讯录转换成功后的url，需要用户通过oauth2授权登录或者单点登录在用户侧打开
		URL string `json:"url"`
	} `json:"contact_id_translate"`
}

// ProviderBatchJobResultPushEvent
// see https://work.weixin.qq.com/api/doc/90001/90143/91875
type ProviderBatchJobResultPushEvent struct {
	XMLName       xml.Name `xml:"xml" json:"-"`
	ServiceCorpID string   `xml:"ServiceCorpId"`
	InfoType      string   `xml:"InfoType"`
	TimeStamp     string   `xml:"TimeStamp"`
	AuthCorpID    string   `xml:"AuthCorpId"`
	BatchJob      struct {
		// JobId 任务id,最大长度为64字符
		JobID string `xml:"JobId"`
		// JobType 任务类型，目前有contact_id_translate
		JobType string `xml:"JobType"`
		ErrCode int    `xml:"ErrCode"`
		ErrMsg  string `xml:"ErrMsg"`
	} `xml:"BatchJob"`
}

// SuiteBatchJobResultPushEvent
// see https://work.weixin.qq.com/api/doc/90001/90143/94956
type SuiteBatchJobResultPushEvent struct {
	XMLName    xml.Name `xml:"xml" json:"-"`
	SuiteID    string   `xml:"SuiteId"`
	InfoType   string   `xml:"InfoType"`
	CreateTime int64    `xml:"CreateTime"`
	TimeStamp  int64    `xml:"TimeStamp"`
	BatchJob   struct {
		// JobId 任务id,最大长度为64字符
		JobID string `xml:"JobId"`
		// JobType 任务类型，目前有 在异步导出的场景下分别有：export_user(导出成员详情)、 export_simple_user(导出成员）、export_department(导出部门）、export_taguser(导出标签成员)
		JobType string `xml:"JobType"`
		ErrCode int    `xml:"ErrCode"`
		ErrMsg  string `xml:"ErrMsg"`
	} `xml:"BatchJob"`
}

func (ProviderBatchJobResultPushEvent) EventType() string {
	return "batch_job_result" //nolint:goconst
}

func (SuiteBatchJobResultPushEvent) EventType() string {
	return "batch_job_result" //nolint:goconst
}

func init() {
	// fixme 区分 provider 事件
	// RegisterEventModel(ProviderBatchJobResultPushEvent{})
}
