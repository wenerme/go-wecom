package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// CreateDoc 新建文档
// 该接口用于新建文档、表格及智能表格
//
// see https://developer.work.weixin.qq.com/document/path/97658
func (c *Client) CreateDoc(r *CreateDocRequest, opts ...any) (out CreateDocResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/create_doc",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateDocRequest is request of Client.CreateDoc
type CreateDocRequest struct {
	// Spaceid 空间spaceid。若指定spaceid，则fatherid也要同时指定
	Spaceid string `json:"spaceid"`
	// Fatherid 父目录fileid, 在根目录时为空间spaceid
	Fatherid string `json:"fatherid"`
	// DocType 文档类型，3:文档 4:表格 10:智能表格
	DocType int `json:"doc_type" validate:"required"`
	// DocName 文档名字（注意：文件名最多填255个字符，超过255个字符会被截断）
	DocName string `json:"doc_name" validate:"required"`
	// AdminUsers 文档管理员userid
	AdminUsers []string `json:"admin_users"`
}

// CreateDocResponse is response of Client.CreateDoc
type CreateDocResponse struct {
	// URL 新建文档的访问链接
	URL string `json:"url"`
	// Docid 新建文档的docid。docid仅在创建时返回，需要开发者妥善保存
	Docid string `json:"docid"`
}

// CreateForm 创建收集表
// 该接口用于创建收集表。
//
// see https://developer.work.weixin.qq.com/document/path/97668
func (c *Client) CreateForm(r *CreateFormRequest, opts ...any) (out CreateFormResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/create_form",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateFormRequest is request of Client.CreateForm
type CreateFormRequest struct {
	// Spaceid 空间spaceid
	Spaceid string `json:"spaceid"`
	// Fatherid 父目录fileid，在根目录时为空间spaceid
	Fatherid string `json:"fatherid"`
	// FormInfo 收集表信息
	FormInfo any `json:"form_info" validate:"required"`
	// FormTitle 收集表标题
	FormTitle string `json:"form_title" validate:"required"`
	// FormDesc 收集表描述
	FormDesc string `json:"form_desc"`
	// FormHeader 收集表表头背景图链接
	FormHeader string `json:"form_header"`
	// FormQuestion 收集表的问题列表
	FormQuestion any `json:"form_question" validate:"required"`
	// Items 问题数组。不超过200个。
	Items []any `json:"items" validate:"required"`
	// QuestionID 问题id，从1开始。如果是家校范围收集表，id从2开始。
	QuestionID int `json:"question_id" validate:"required"`
	// Title 问题描述
	Title string `json:"title" validate:"required"`
	// Pos 问题序号，从1开始。
	Pos int `json:"pos" validate:"required"`
	// Status 问题状态。1：正常；2：被删除
	Status int `json:"status" validate:"required"`
	// ReplyType 问题类型。1：文本；2：单选；3：多选；5：位置；9：图片；10：文件；11：日期；14：时间；15：下拉列表；16：体温；17：签名；18：部门；19：成员 22：时长
	ReplyType int `json:"reply_type" validate:"required"`
	// MustReply 是否必答
	MustReply bool `json:"must_reply" validate:"required"`
	// Note 问题备注
	Note string `json:"note"`
	// Placeholder 编辑提示
	Placeholder string `json:"placeholder"`
	// QuestionExtendSetting 问题的额外设置。不同问题类型有相应的设置
	QuestionExtendSetting any `json:"question_extend_setting"`
	// OptionItem 单选/多选/下拉列表题的选项列表
	OptionItem []any `json:"option_item" validate:"required"`
	// Key 选项key（1，2，3...）
	Key int `json:"key" validate:"required"`
	// Value 选项内容
	Value string `json:"value" validate:"required"`
	// FormSetting 收集表设置
	FormSetting any `json:"form_setting"`
	// FillOutAuth 填写权限。0：所有人；1：企业内指定人/部门；4:家校所有范围。默认为0，所有人可填写。
	FillOutAuth int `json:"fill_out_auth"`
	// FillInRange 指定的可填写的人/部门。当timed_repeat_info.enable为true时必填
	FillInRange any `json:"fill_in_range"`
	// Userids 企业成员userid列表
	Userids []any `json:"userids"`
	// Departmentids 部门id列表
	Departmentids []any `json:"departmentids"`
	// SettingManagerRange 收集表管理员
	SettingManagerRange any `json:"setting_manager_range"`
	// TimedRepeatInfo 定时重复设置项
	TimedRepeatInfo any `json:"timed_repeat_info"`
	// AllowMultiFill 是否允许每人提交多份。默认false
	AllowMultiFill bool `json:"allow_multi_fill"`
	// TimedFinish 定时关闭。定时重复与定时结束互斥，若都填，优先定时重复
	TimedFinish int `json:"timed_finish"`
	// CanAnonymous 是否支持匿名填写。默认false
	CanAnonymous bool `json:"can_anonymous"`
	// CanNotifySubmit 是否有回复时提醒。默认false
	CanNotifySubmit bool `json:"can_notify_submit"`
}

// CreateFormResponse is response of Client.CreateForm
type CreateFormResponse struct {
	// Formid 收集表id
	Formid string `json:"formid"`
}

// DeleteDoc 删除文档
// 该接口用于删除指定文档、表格、智能表格及收集表。
//
// see https://developer.work.weixin.qq.com/document/path/99930
func (c *Client) DeleteDoc(r *DeleteDocRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/del_doc",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteDocRequest is request of Client.DeleteDoc
type DeleteDocRequest struct {
	// Docid 文档docid（docid、formid只能填其中一个），仅可删除应用自己创建的文档
	Docid string `json:"docid"`
	// Formid 收集表id（docid、formid只能填其中一个），仅可删除应用自己创建的收集表
	Formid string `json:"formid"`
}

// GetDocAuth 获取文档权限信息
// 该接口用于获取文档、表格、智能表格的查看规则、文档通知范围及权限、安全设置信息
//
// see https://developer.work.weixin.qq.com/document/path/97811
func (c *Client) GetDocAuth(r *GetDocAuthRequest, opts ...any) (out GetDocAuthResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/doc_get_auth",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetDocAuthRequest is request of Client.GetDocAuth
type GetDocAuthRequest struct {
	// Docid 文档id
	Docid string `json:"docid" validate:"required"`
}

// GetDocAuthResponse is response of Client.GetDocAuth
type GetDocAuthResponse struct {
	// AccessRule 文档的查看规则
	AccessRule any `json:"access_rule"`
	// EnableCorpInternal 是否允许企业内成员浏览文档
	EnableCorpInternal bool `json:"enable_corp_internal"`
	// CorpInternalAuth 企业内成员主动查看文档后获得的权限类型 1:只读 2:读写
	CorpInternalAuth int `json:"corp_internal_auth"`
	// EnableCorpExternal 是否允许企业外成员浏览文档
	EnableCorpExternal bool `json:"enable_corp_external"`
	// CorpExternalAuth 企业内成员主动查看文档后获得的权限类型 1:只读 2:读写
	CorpExternalAuth int `json:"corp_external_auth"`
	// CorpInternalApproveOnlyByAdmin 企业内成员浏览文档是否必须由管理员审批
	CorpInternalApproveOnlyByAdmin bool `json:"corp_internal_approve_only_by_admin"`
	// CorpExternalApproveOnlyByAdmin 企业外成员浏览文档是否必须由管理员审批
	CorpExternalApproveOnlyByAdmin bool `json:"corp_external_approve_only_by_admin"`
	// BanShareExternal 是否允许企业外成员浏览文档
	BanShareExternal bool `json:"ban_share_external"`
	// SecureSetting 安全设置信息
	SecureSetting any `json:"secure_setting"`
	// EnableReadonlyCopy 仅浏览权限的成员是否允许导出、复制、打印
	EnableReadonlyCopy bool `json:"enable_readonly_copy"`
	// Watermark 文档水印设置
	Watermark any `json:"watermark"`
	// MarginType 水印密度 1:稀疏 2:紧密
	MarginType int `json:"margin_type"`
	// ShowVisitorName 是否展示访问者名字
	ShowVisitorName bool `json:"show_visitor_name"`
	// ShowText 是否展示水印文字
	ShowText bool `json:"show_text"`
	// Text 水印文字
	Text string `json:"text"`
	// DocMemberList 文档通知范围及权限列表
	DocMemberList []map[string]any `json:"doc_member_list"`
	// Type 文档通知范围成员种类 1:user, 只支持成员
	Type int `json:"type"`
	// UserID 企业成员的userid
	UserID string `json:"userid"`
	// TmpExternalUserID 外部用户临时id
	TmpExternalUserID string `json:"tmp_external_userid"`
	// Auth 该文档通知范围成员的权限 1:只读 2:读写 7:管理员
	Auth int `json:"auth"`
	// CoAuthList 文档查看权限特定部门列表
	CoAuthList []map[string]any `json:"co_auth_list"`
	// Departmentid 特定部门id
	Departmentid int `json:"departmentid"`
}

// ShareDoc 分享文档
// 该接口用于获取文档、表格、智能表格及收集表的分享链接。
//
// see https://developer.work.weixin.qq.com/document/path/97733
func (c *Client) ShareDoc(r *ShareDocRequest, opts ...any) (out ShareDocResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/doc_share",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ShareDocRequest is request of Client.ShareDoc
type ShareDocRequest struct {
	// Docid 文档id（docid、formid只能填其中一个）
	Docid string `json:"docid"`
	// Formid 表单id（docid、formid只能填其中一个）
	Formid string `json:"formid"`
}

// ShareDocResponse is response of Client.ShareDoc
type ShareDocResponse struct {
	// ShareURL 文档分享链接
	ShareURL string `json:"share_url"`
}

// BatchUpdateDocument 编辑文档内容
// 该接口可以对一个在线文档批量执行多个更新操作。
//
// see https://developer.work.weixin.qq.com/document/path/97626
func (c *Client) BatchUpdateDocument(r *BatchUpdateDocumentRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/document/batch_update",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchUpdateDocumentRequest is request of Client.BatchUpdateDocument
type BatchUpdateDocumentRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// Version 操作的文档版本，该参数可以通过获取文档内容接口获得。操作后文档版本将更新一版。要更新的文档版本与最新文档版本相差不能超过100个。
	Version int `json:"version"`
	// Requests 更新操作列表，详见 UpdateRequest
	Requests []map[string]any `json:"requests" validate:"required"`
}

// GetWedocDocument 获取文档数据
// 该接口用于获取文档数据
//
// see https://developer.work.weixin.qq.com/document/path/101161
func (c *Client) GetWedocDocument(r *GetWedocDocumentRequest, opts ...any) (out GetWedocDocumentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/document/get",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetWedocDocumentRequest is request of Client.GetWedocDocument
type GetWedocDocumentRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
}

// GetWedocDocumentResponse is response of Client.GetWedocDocument
type GetWedocDocumentResponse struct {
	// Version 文档版本
	Version int `json:"version"`
	// Document 文档内容根节点
	Document any `json:"document"`
}

// GetDocBaseInfo 获取文档基础信息
// 该接口用于获取指定文档、表格、智能表格及收集表的基础信息。
//
// see https://developer.work.weixin.qq.com/document/path/97734
func (c *Client) GetDocBaseInfo(r *GetDocBaseInfoRequest, opts ...any) (out GetDocBaseInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/get_doc_base_info",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetDocBaseInfoRequest is request of Client.GetDocBaseInfo
type GetDocBaseInfoRequest struct {
	// Docid 文档docid
	Docid string `json:"docid" validate:"required"`
}

// GetDocBaseInfoResponse is response of Client.GetDocBaseInfo
type GetDocBaseInfoResponse struct {
	// Docid 文档docid
	Docid string `json:"docid"`
	// DocName 文档名字
	DocName string `json:"doc_name"`
	// CreateTime 文档创建时间
	CreateTime int `json:"create_time"`
	// ModifyTime 文档最后修改时间
	ModifyTime int `json:"modify_time"`
	// DocType 文档类型：3: 文档 4: 表格 10:智能表格
	DocType int `json:"doc_type"`
}

// GetWedocFormAnswer 读取收集表答案
// 该接口用于读取收集表的答案
//
// see https://developer.work.weixin.qq.com/document/path/97819
func (c *Client) GetWedocFormAnswer(r *GetWedocFormAnswerRequest, opts ...any) (out GetWedocFormAnswerResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/get_form_answer",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetWedocFormAnswerRequest is request of Client.GetWedocFormAnswer
type GetWedocFormAnswerRequest struct {
	// RepeatedID 操作的收集表周期id
	RepeatedID string `json:"repeated_id" validate:"required"`
	// AnswerIds 需要拉取的答案列表，批次大小最大100
	AnswerIds []int `json:"answer_ids" validate:"required"`
}

// GetWedocFormAnswerResponse is response of Client.GetWedocFormAnswer
type GetWedocFormAnswerResponse struct {
	// Answer 答案对象
	Answer any `json:"answer"`
	// AnswerList 答案列表
	AnswerList []any `json:"answer_list"`
	// AnswerID 答案id
	AnswerID int `json:"answer_id"`
	// UserName 用户名
	UserName string `json:"user_name"`
	// Ctime 创建时间
	Ctime int `json:"ctime"`
	// Mtime 修改时间
	Mtime int `json:"mtime"`
	// Reply 该用户的答案明细
	Reply any `json:"reply"`
	// Items 每个问题的答案
	Items []any `json:"items"`
	// QuestionID 问题id
	QuestionID int `json:"question_id"`
	// TextReply 答案
	TextReply string `json:"text_reply"`
	// OptionReply 选择题答案，多选题有多个答案
	OptionReply []any `json:"option_reply"`
	// OptionExtendReply 选择题，其他选项列表
	OptionExtendReply []any `json:"option_extend_reply"`
	// FileExtendReply 文件题答案列表
	FileExtendReply []any `json:"file_extend_reply"`
	// DepartmentReply 部门题答案
	DepartmentReply any `json:"department_reply"`
	// MemberReply 成员题答案
	MemberReply any `json:"member_reply"`
	// DurationReply 时长题答案
	DurationReply any `json:"duration_reply"`
	// AnswerStatus 答案状态 1:正常 3:统计者移除此答案或删除
	AnswerStatus int `json:"answer_status"`
	// TmpExternalUserID 外部用户临时id
	TmpExternalUserID string `json:"tmp_external_userid"`
}

// GetFormInfo 获取收集表信息
// 该接口用于读取收集表的信息
//
// see https://developer.work.weixin.qq.com/document/path/97817
func (c *Client) GetFormInfo(r *GetFormInfoRequest, opts ...any) (out GetFormInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/get_form_info",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetFormInfoRequest is request of Client.GetFormInfo
type GetFormInfoRequest struct {
	// Formid 操作的收集表ID
	Formid string `json:"formid" validate:"required"`
}

// GetFormInfoResponse is response of Client.GetFormInfo
type GetFormInfoResponse struct {
	// FormInfo 收集表信息
	FormInfo any `json:"form_info"`
}

// GetFormStatistic 收集表的统计信息查询
// 该接口用于获取收集表的统计信息、已回答成员列表和未回答成员列表
//
// see https://developer.work.weixin.qq.com/document/path/97818
func (c *Client) GetFormStatistic(r *GetFormStatisticRequest, opts ...any) (out GetFormStatisticResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/get_form_statistic",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetFormStatisticRequest is request of Client.GetFormStatistic
type GetFormStatisticRequest struct {
	// RepeatedID 操作的收集表的repeated_id，来源于get_form_info的返回
	RepeatedID string `json:"repeated_id" validate:"required"`
	// ReqType 请求类型 1:只获取统计结果 2:获取已提交列表 3:获取未提交列表
	ReqType int `json:"req_type" validate:"required"`
	// StartTime 拉取已提交列表时必填，其余type不填。筛选开始时间，以当天的00:00:00开始筛选
	StartTime string `json:"start_time"`
	// EndTime 拉取已提交列表时必填，其余type不填。筛选结束时间，以当天的23:59:59结束筛选
	EndTime string `json:"end_time"`
	// Limit 分页拉取时批次大小，最大10000
	Limit string `json:"limit"`
	// Cursor 分页拉取的游标，首次不传
	Cursor string `json:"cursor"`
}

// GetFormStatisticResponse is response of Client.GetFormStatistic
type GetFormStatisticResponse struct {
	// FillCnt 已填写次数
	FillCnt int `json:"fill_cnt"`
	// FillUserCnt 已填写人数
	FillUserCnt int `json:"fill_user_cnt"`
	// UnfillUserCnt 未填写人数
	UnfillUserCnt int `json:"unfill_user_cnt"`
	// SubmitUsers 已填写人列表
	SubmitUsers []map[string]any `json:"submit_users"`
	// TmpExternalUserID 外部用户临时id，匿名填写不返回
	TmpExternalUserID string `json:"tmp_external_userid"`
	// UserID 企业内成员的id，匿名填写不返回
	UserID string `json:"userid"`
	// SubmitTime 提交时间
	SubmitTime string `json:"submit_time"`
	// AnswerID 答案id
	AnswerID string `json:"answer_id"`
	// UserName 名字，匿名填写不返回
	UserName string `json:"user_name"`
	// UnfillUsers 未填写人列表
	UnfillUsers []any `json:"unfill_users"`
	// HasMore 是否还有更多
	HasMore bool `json:"has_more"`
	// Cursor 上次分页拉取返回的cursor
	Cursor int `json:"cursor"`
}

// UploadWedocImage 上传文档图片
// 该接口用于上传图片
//
// see https://developer.work.weixin.qq.com/document/path/99933
func (c *Client) UploadWedocImage(r *UploadWedocImageRequest, opts ...any) (out UploadWedocImageResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/image_upload",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UploadWedocImageRequest is request of Client.UploadWedocImage
type UploadWedocImageRequest struct {
	// Docid 文档ID，通过新建文档接口创建后获得
	Docid string `json:"docid" validate:"required"`
	// Base64Content base64之后的图片内容
	Base64Content string `json:"base64_content" validate:"required"`
}

// UploadWedocImageResponse is response of Client.UploadWedocImage
type UploadWedocImageResponse struct {
	// URL 图片的url
	URL string `json:"url"`
	// Height 图片的高
	Height int `json:"height"`
	// Width 图片的宽
	Width int `json:"width"`
	// Size 图片的大小
	Size int `json:"size"`
}

// UpdateDocJoinRule 修改文档加入规则
// 该接口用于修改文档、表格、智能表格加入规则。
//
// see https://developer.work.weixin.qq.com/document/path/101477
func (c *Client) UpdateDocJoinRule(r *UpdateDocJoinRuleRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/mod_doc_join_rule",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateDocJoinRuleRequest is request of Client.UpdateDocJoinRule
type UpdateDocJoinRuleRequest struct {
	// Docid 操作的docid
	Docid string `json:"docid" validate:"required"`
	// EnableCorpInternal 是否允许企业内成员浏览文档，有值则覆盖
	EnableCorpInternal bool `json:"enable_corp_internal"`
	// CorpInternalAuth 企业内成员主动查看文档后获得的权限类型 1:只读 2:读写，有值则覆盖
	CorpInternalAuth int `json:"corp_internal_auth"`
	// EnableCorpExternal 是否允许企业外成员浏览文档，有值则覆盖
	EnableCorpExternal bool `json:"enable_corp_external"`
	// CorpExternalAuth 企业外成员主浏览文档后获得的权限类型 1:只读 2:读写，有值则覆盖
	CorpExternalAuth int `json:"corp_external_auth"`
	// CorpInternalApproveOnlyByAdmin 企业内成员加入文档是否必须由管理员审批，enable_corp_internal为false时，只能为true，有值则覆盖。设置为true之前，文档需要有至少一个管理员。
	CorpInternalApproveOnlyByAdmin bool `json:"corp_internal_approve_only_by_admin"`
	// CorpExternalApproveOnlyByAdmin 企业外成员加入文档是否必须由管理员审批，enable_corp_external和ban_share_external均为false时，该参数只能为true，有值则覆盖。设置为true之前，文档...
	CorpExternalApproveOnlyByAdmin bool `json:"corp_external_approve_only_by_admin"`
	// BanShareExternal 是否禁止文档分享到企业外，有值则覆盖
	BanShareExternal bool `json:"ban_share_external"`
	// UpdateCoAuthList 是否更新文档查看权限的特定部门，true时更新特定部门列表
	UpdateCoAuthList bool `json:"update_co_auth_list"`
	// CoAuthList 需要更新文档查看权限特定部门时，覆盖之前部门，特别的：列表为空则清空
	CoAuthList []map[string]any `json:"co_auth_list"`
}

// UpdateDocMember 修改文档成员与权限
// 该接口用于管理文档、表格、智能表格的文档成员，支持增加或删除成员、修改成员的权限。
//
// see https://developer.work.weixin.qq.com/document/path/101476
func (c *Client) UpdateDocMember(r *UpdateDocMemberRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/mod_doc_member",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateDocMemberRequest is request of Client.UpdateDocMember
type UpdateDocMemberRequest struct {
	// Docid 操作的文档id
	Docid string `json:"docid" validate:"required"`
	// UpdateFileMemberList 更新文档成员的列表，批次大小最大100
	UpdateFileMemberList []map[string]any `json:"update_file_member_list"`
	// DelFileMemberList 删除的文档成员列表，批次大小最大一百
	DelFileMemberList []map[string]any `json:"del_file_member_list"`
}

// UpdateDocSafetySetting 修改文档安全设置
// 该接口用于修改文档、表格、智能表格的安全设置。
//
// see https://developer.work.weixin.qq.com/document/path/97782
func (c *Client) UpdateDocSafetySetting(r *UpdateDocSafetySettingRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/mod_doc_safty_setting",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateDocSafetySettingRequest is request of Client.UpdateDocSafetySetting
type UpdateDocSafetySettingRequest struct {
	// Docid 操作的文档id
	Docid string `json:"docid" validate:"required"`
	// EnableReadonlyCopy 是否允许只读成员复制、下载文档，有值则覆盖
	EnableReadonlyCopy bool `json:"enable_readonly_copy"`
	// Watermark 水印设置
	Watermark any `json:"watermark"`
}

// ModifyForm 编辑收集表
// 该接口用于编辑收集表。
//
// see https://developer.work.weixin.qq.com/document/path/97816
func (c *Client) ModifyForm(r *ModifyFormRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/modify_form",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ModifyFormRequest is request of Client.ModifyForm
type ModifyFormRequest struct {
	// Oper 操作类型。1：全量修改问题；2：全量修改设置
	Oper int `json:"oper" validate:"required"`
	// Formid 收集表id
	Formid string `json:"formid" validate:"required"`
	// FormTitle 收集表标题（操作1修改）
	FormTitle string `json:"form_title"`
	// FormDesc 收集表描述（操作1修改）
	FormDesc string `json:"form_desc"`
	// FormHeader 收集表表头背景图链接（操作1修改）
	FormHeader string `json:"form_header"`
	// FormQuestion 收集表的问题列表（操作1修改）
	FormQuestion any `json:"form_question"`
	// Items 问题数组
	Items []any `json:"items" validate:"required"`
	// QuestionID 问题id，从1开始。如果是家校范围收集表，id从2开始。
	QuestionID int `json:"question_id" validate:"required"`
	// Title 问题描述
	Title string `json:"title" validate:"required"`
	// Pos 问题序号，从1开始。
	Pos int `json:"pos" validate:"required"`
	// Status 问题状态。1：正常；2：被删除
	Status int `json:"status" validate:"required"`
	// ReplyType 问题类型。1：文本；2：单选；3：多选；5：位置；9：图片；10：文件；11：日期；14：时间；15：下拉列表；16：体温；17：签名；18：部门；19：成员 22：时长
	ReplyType int `json:"reply_type" validate:"required"`
	// MustReply 是否必答
	MustReply bool `json:"must_reply" validate:"required"`
	// Note 问题备注
	Note string `json:"note"`
	// Placeholder 编辑提示
	Placeholder string `json:"placeholder"`
	// QuestionExtendSetting 问题的额外设置。不同问题类型有相应的设置
	QuestionExtendSetting any `json:"question_extend_setting"`
	// OptionItem 单选/多选/下拉列表题的选项列表
	OptionItem []any `json:"option_item" validate:"required"`
	// Key 选项key（1，2，3...）
	Key int `json:"key" validate:"required"`
	// Value 选项内容
	Value string `json:"value" validate:"required"`
	// FormSetting 收集表设置（操作2修改）
	FormSetting any `json:"form_setting"`
	// FillOutAuth 填写权限。0：所有人；1：企业内指定人/部门
	FillOutAuth int `json:"fill_out_auth" validate:"required"`
	// FillInRange 指定的可填写的人/部门
	FillInRange any `json:"fill_in_range"`
	// Userids 企业成员userid列表
	Userids []any `json:"userids"`
	// Departmentids 部门id列表
	Departmentids []any `json:"departmentids"`
	// SettingManagerRange 收集表管理员
	SettingManagerRange any `json:"setting_manager_range"`
	// TimedRepeatInfo 定时重复设置项
	TimedRepeatInfo any `json:"timed_repeat_info"`
	// AllowMultiFill 是否允许每人提交多份。默认false
	AllowMultiFill bool `json:"allow_multi_fill"`
	// TimedFinish 定时关闭。定时重复与定时结束互斥，若都填，优先定时重复
	TimedFinish int `json:"timed_finish"`
	// CanAnonymous 是否支持匿名填写。默认false
	CanAnonymous bool `json:"can_anonymous"`
	// CanNotifySubmit 是否有回复时提醒。默认false
	CanNotifySubmit bool `json:"can_notify_submit"`
}

// RenameDoc 重命名文档
// 该接口用于对指定文档、表格、智能表格及收集表进行重命名。
//
// see https://developer.work.weixin.qq.com/document/path/99894
func (c *Client) RenameDoc(r *RenameDocRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/rename_doc",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// RenameDocRequest is request of Client.RenameDoc
type RenameDocRequest struct {
	// Docid 文档docid（docid、formid只能填其中一个），仅可修改应用自己创建的文档
	Docid string `json:"docid"`
	// Formid 收集表id（docid、formid只能填其中一个），仅可修改应用自己创建的收集表
	Formid string `json:"formid"`
	// NewName 重命名后的文档名（注意：文档名最多填255个字符，英文算1个，汉字算2个，超过255个字符会被截断）
	NewName string `json:"new_name" validate:"required"`
}

// AddFieldGroup 添加编组
// 本接口用于在智能表中的某个子表里添加编组。单表最多允许有150个编组。每个编组最多允许有150个字段。字段只能同时存在于一个编组。
//
// see https://developer.work.weixin.qq.com/document/path/101100
func (c *Client) AddFieldGroup(r *AddFieldGroupRequest, opts ...any) (out AddFieldGroupResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/add_field_group",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddFieldGroupRequest is request of Client.AddFieldGroup
type AddFieldGroupRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID 表格ID
	SheetID string `json:"sheet_id" validate:"required"`
	// Name 编组名称，不能和已有名称重复
	Name string `json:"name" validate:"required"`
	// Children 编组内容
	Children []map[string]any `json:"children"`
}

// AddFieldGroupResponse is response of Client.AddFieldGroup
type AddFieldGroupResponse struct {
	// FieldGroup 编组
	FieldGroup any `json:"field_group"`
}

// AddSmartsheetFields 添加字段
// 本接口用于在智能表中的某个子表里添加一列或多列新字段。单表最多允许有150个字段。
//
// see https://developer.work.weixin.qq.com/document/path/99904
func (c *Client) AddSmartsheetFields(r *AddSmartsheetFieldsRequest, opts ...any) (out AddSmartsheetFieldsResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/add_fields",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddSmartsheetFieldsRequest is request of Client.AddSmartsheetFields
type AddSmartsheetFieldsRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID 表格ID
	SheetID string `json:"sheet_id" validate:"required"`
	// Fields 字段详情列表
	Fields []map[string]any `json:"fields" validate:"required"`
}

// AddSmartsheetFieldsResponse is response of Client.AddSmartsheetFields
type AddSmartsheetFieldsResponse struct {
	// Fields 新增的字段详情列表
	Fields []map[string]any `json:"fields"`
}

// AddRecords 添加记录
// 本接口用于在 Smartsheet 中的某个子表里添加一行或多行新记录。单表最多允许有100000行记录，15000000个单元格。单次添加建议在500行内
//
// see https://developer.work.weixin.qq.com/document/path/99907
func (c *Client) AddRecords(r *AddRecordsRequest, opts ...any) (out AddRecordsResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/add_records",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddRecordsRequest is request of Client.AddRecords
type AddRecordsRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID Smartsheet 子表ID
	SheetID string `json:"sheet_id" validate:"required"`
	// KeyType 返回记录中单元格的key类型，默认用标题。枚举值：CELL_VALUE_KEY_TYPE_FIELD_TITLE, CELL_VALUE_KEY_TYPE_FIELD_ID
	KeyType string `json:"key_type"`
	// Records 需要添加的记录的具体内容组成的 JSON 数组
	Records []map[string]any `json:"records" validate:"required"`
}

// AddRecordsResponse is response of Client.AddRecords
type AddRecordsResponse struct {
	// Records 由添加成功的记录的具体内容组成的 JSON 数组
	Records []any `json:"records"`
}

// AddSmartSheet 添加子表
// 在表格的某个位置添加一个智能表
//
// see https://developer.work.weixin.qq.com/document/path/99896
func (c *Client) AddSmartSheet(r *AddSmartSheetRequest, opts ...any) (out AddSmartSheetResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/add_sheet",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddSmartSheetRequest is request of Client.AddSmartSheet
type AddSmartSheetRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// Properties 智能表属性
	Properties any `json:"properties"`
}

// AddSmartSheetResponse is response of Client.AddSmartSheet
type AddSmartSheetResponse struct {
	// Properties 智能表属性
	Properties any `json:"properties"`
}

// AddView 添加视图
// 在 Smartsheet 中的某个子表里添加一个新视图。单表最多允许有200个视图。
//
// see https://developer.work.weixin.qq.com/document/path/99900
func (c *Client) AddView(r *AddViewRequest, opts ...any) (out AddViewResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/add_view",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddViewRequest is request of Client.AddView
type AddViewRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID Smartsheet 子表ID
	SheetID string `json:"sheet_id" validate:"required"`
	// ViewTitle 视图标题
	ViewTitle string `json:"view_title" validate:"required"`
	// ViewType 视图类型。见ViewType
	ViewType string `json:"view_type" validate:"required"`
	// PropertyGantt 甘特视图属性，添加甘特图时必填
	PropertyGantt any `json:"property_gantt"`
	// PropertyCalendar 日历视图属性，添加日历视图时必填
	PropertyCalendar any `json:"property_calendar"`
}

// AddViewResponse is response of Client.AddView
type AddViewResponse struct {
	// View 添加视图响应
	View any `json:"view"`
}

// GetSheetPriv 查询智能表格子表权限
// 查询智能表格子表权限详情
//
// see https://developer.work.weixin.qq.com/document/path/99935
func (c *Client) GetSheetPriv(r *GetSheetPrivRequest, opts ...any) (out GetSheetPrivResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/content_priv/get_sheet_priv",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetSheetPrivRequest is request of Client.GetSheetPriv
type GetSheetPrivRequest struct {
	// Docid 智能表ID，通过新建文档接口创建后获得
	Docid string `json:"docid" validate:"required"`
	// Type 权限规则类型，1-全员权限，2-额外权限
	Type int `json:"type" validate:"required"`
	// RuleIDList 需要查询的规则id列表，查询额外权限时填写
	RuleIDList []string `json:"rule_id_list"`
}

// GetSheetPrivResponse is response of Client.GetSheetPriv
type GetSheetPrivResponse struct {
	// RuleList 权限列表
	RuleList []map[string]any `json:"rule_list"`
}

// DeleteSmartsheetFieldGroups 删除编组
// 本接口用于删除智能表的某个子表中的一个或多个编组。
//
// see https://developer.work.weixin.qq.com/document/path/101102
func (c *Client) DeleteSmartsheetFieldGroups(r *DeleteSmartsheetFieldGroupsRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/delete_field_groups",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteSmartsheetFieldGroupsRequest is request of Client.DeleteSmartsheetFieldGroups
type DeleteSmartsheetFieldGroupsRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID 子表ID
	SheetID string `json:"sheet_id" validate:"required"`
	// FieldGroupIds 要删除的编组 ID
	FieldGroupIds []string `json:"field_group_ids" validate:"required"`
}

// DeleteSmartsheetFields 删除字段
// 本接口用于删除智能表中的某个子表里的一列或多列字段。
//
// see https://developer.work.weixin.qq.com/document/path/99905
func (c *Client) DeleteSmartsheetFields(r *DeleteSmartsheetFieldsRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/delete_fields",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteSmartsheetFieldsRequest is request of Client.DeleteSmartsheetFields
type DeleteSmartsheetFieldsRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID 表格ID
	SheetID string `json:"sheet_id" validate:"required"`
	// FieldIds 需要删除的字段id列表
	FieldIds []string `json:"field_ids" validate:"required"`
}

// DeleteSmartsheetRecords 删除记录
// 本接口用于删除 Smartsheet 的某个子表中的一行或多行记录。单次删除建议在500行内
//
// see https://developer.work.weixin.qq.com/document/path/99908
func (c *Client) DeleteSmartsheetRecords(r *DeleteSmartsheetRecordsRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/delete_records",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteSmartsheetRecordsRequest is request of Client.DeleteSmartsheetRecords
type DeleteSmartsheetRecordsRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID Smartsheet 子表ID
	SheetID string `json:"sheet_id" validate:"required"`
	// RecordIds 要删除的记录 ID
	RecordIds []string `json:"record_ids" validate:"required"`
}

// DeleteSheet 删除子表
// 本接口用于删除在线表格中的某个智能表。
//
// see https://developer.work.weixin.qq.com/document/path/99899
func (c *Client) DeleteSheet(r *DeleteSheetRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/delete_sheet",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteSheetRequest is request of Client.DeleteSheet
type DeleteSheetRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID 删除的Smartsheet 子表 ID
	SheetID string `json:"sheet_id" validate:"required"`
}

// DeleteViews 删除视图
// 本接口用于在 smartsheet 中的某个子表里删除若干个视图。
//
// see https://developer.work.weixin.qq.com/document/path/99901
func (c *Client) DeleteViews(r *DeleteViewsRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/delete_views",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteViewsRequest is request of Client.DeleteViews
type DeleteViewsRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID Smartsheet 子表ID
	SheetID string `json:"sheet_id" validate:"required"`
	// ViewIds 要删除的视图ID列表
	ViewIds []string `json:"view_ids" validate:"required"`
}

// GetFieldGroups 获取编组
// 本接口用于在智能表中的某个子表里获取已有的编组。
//
// see https://developer.work.weixin.qq.com/document/path/101103
func (c *Client) GetFieldGroups(r *GetFieldGroupsRequest, opts ...any) (out GetFieldGroupsResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/get_field_groups",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetFieldGroupsRequest is request of Client.GetFieldGroups
type GetFieldGroupsRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID 表格ID
	SheetID string `json:"sheet_id" validate:"required"`
	// Offset 偏移量，初始值为 0
	Offset int `json:"offset"`
	// Limit 分页大小 , 每页返回多少条数据
	Limit int `json:"limit"`
}

// GetFieldGroupsResponse is response of Client.GetFieldGroups
type GetFieldGroupsResponse struct {
	// Total 编组数量
	Total int `json:"total"`
	// HasMore 是否还有更多数据
	HasMore bool `json:"has_more"`
	// Next 下一偏移位置
	Next int `json:"next"`
	// FieldGroups 编组列表
	FieldGroups []map[string]any `json:"field_groups"`
}

// GetSmartSheetFields 查询字段
// 获取智能表中某个子表下字段信息，支持获取全部字段、依据字段名或字段ID获取对应字段。
//
// see https://developer.work.weixin.qq.com/document/path/101157
func (c *Client) GetSmartSheetFields(r *GetSmartSheetFieldsRequest, opts ...any) (out GetSmartSheetFieldsResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/get_fields",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetSmartSheetFieldsRequest is request of Client.GetSmartSheetFields
type GetSmartSheetFieldsRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID 表格ID
	SheetID string `json:"sheet_id" validate:"required"`
	// ViewID 视图 ID
	ViewID string `json:"view_id"`
	// FieldIds 由字段 ID 组成的 JSON 数组
	FieldIds []any `json:"field_ids"`
	// FieldTitles 由字段标题组成的 JSON 数组
	FieldTitles []any `json:"field_titles"`
	// Offset 偏移量，初始值为 0
	Offset int `json:"offset"`
	// Limit 分页大小，每页返回多少条数据；最大值 1000
	Limit int `json:"limit"`
}

// GetSmartSheetFieldsResponse is response of Client.GetSmartSheetFields
type GetSmartSheetFieldsResponse struct {
	// Total 字段总数
	Total int `json:"total"`
	// Fields 字段详情列表
	Fields []map[string]any `json:"fields"`
}

// GetSmartsheetRecords 查询记录
// 获取 Smartsheet 中某个子表下记录信息，支持获取全部记录、依据字段名和记录 ID 获取对应记录、对记录进行排序。
//
// see https://developer.work.weixin.qq.com/document/path/101158
func (c *Client) GetSmartsheetRecords(r *GetSmartsheetRecordsRequest, opts ...any) (out GetSmartsheetRecordsResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/get_records",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetSmartsheetRecordsRequest is request of Client.GetSmartsheetRecords
type GetSmartsheetRecordsRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID Smartsheet 子表ID
	SheetID string `json:"sheet_id" validate:"required"`
	// ViewID 视图 ID
	ViewID string `json:"view_id"`
	// RecordIds 由记录 ID 组成的 JSON 数组
	RecordIds []any `json:"record_ids"`
	// KeyType 返回记录中单元格的key类型，可选 CELL_VALUE_KEY_TYPE_FIELD_TITLE 或 CELL_VALUE_KEY_TYPE_FIELD_ID
	KeyType string `json:"key_type"`
	// FieldTitles 返回指定列，由字段标题组成的 JSON 数组，key_type 为 CELL_VALUE_KEY_TYPE_FIELD_TITLE 时有效
	FieldTitles []any `json:"field_titles"`
	// FieldIds 返回指定列，由字段 ID 组成的 JSON 数组，key_type 为 CELL_VALUE_KEY_TYPE_FIELD_ID 时有效
	FieldIds []any `json:"field_ids"`
	// Sort 对返回记录进行排序的数组
	Sort []any `json:"sort"`
	// Offset 偏移量，初始值为 0
	Offset int `json:"offset"`
	// Limit 分页大小，每页返回多少条数据；最大值 1000
	Limit int `json:"limit"`
	// Ver 版本号
	Ver int `json:"ver"`
	// FilterSpec 过滤设置，不支持和sort一起使用
	FilterSpec any `json:"filter_spec"`
}

// GetSmartsheetRecordsResponse is response of Client.GetSmartsheetRecords
type GetSmartsheetRecordsResponse struct {
	// Total 符合筛选条件的视图总数
	Total int `json:"total"`
	// HasMore 是否还有更多项
	HasMore bool `json:"has_more"`
	// Next 下次下一个搜索结果的偏移量
	Next int `json:"next"`
	// Records 由查询记录的具体内容组成的 JSON 数组
	Records []any `json:"records"`
	// Ver 版本号
	Ver int `json:"ver"`
}

// GetSheet 查询子表
// 本接口用于查询一篇在线表格中全部智能表信息。
//
// see https://developer.work.weixin.qq.com/document/path/101154
func (c *Client) GetSheet(r *GetSheetRequest, opts ...any) (out GetSheetResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/get_sheet",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetSheetRequest is request of Client.GetSheet
type GetSheetRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID 指定子表ID查询
	SheetID string `json:"sheet_id"`
	// NeedAllTypeSheet 获取所有类型子表。为true时可获取包含仪表盘和说明页在内的所有类型的子表
	NeedAllTypeSheet bool `json:"need_all_type_sheet"`
}

// GetSheetResponse is response of Client.GetSheet
type GetSheetResponse struct {
	// SheetList 智能表信息列表
	SheetList []map[string]any `json:"sheet_list"`
}

// GetSmartsheetViews 查询视图
// 本接口用于获取 Smartsheet 中某个子表里全部视图信息。
//
// see https://developer.work.weixin.qq.com/document/path/101155
func (c *Client) GetSmartsheetViews(r *GetSmartsheetViewsRequest, opts ...any) (out GetSmartsheetViewsResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/get_views",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetSmartsheetViewsRequest is request of Client.GetSmartsheetViews
type GetSmartsheetViewsRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID Smartsheet 子表ID
	SheetID string `json:"sheet_id" validate:"required"`
	// ViewIds 需要查询的视图 ID 数组
	ViewIds []string `json:"view_ids"`
	// Offset 偏移量，初始值为 0
	Offset int `json:"offset"`
	// Limit 分页大小 , 每页返回多少条数据；当不填写该参数或将该参数设置为 0 时，如果总数大于 1000，一次性返回 1000 个视图，当总数小于 1000 时，返回全部视图；limit 最大值为 1000
	Limit int `json:"limit"`
}

// GetSmartsheetViewsResponse is response of Client.GetSmartsheetViews
type GetSmartsheetViewsResponse struct {
	// Total 符合筛选条件的视图总数
	Total int `json:"total"`
	// HasMore 是否还有更多项
	HasMore bool `json:"has_more"`
	// Next 下次下一个搜索结果的偏移量
	Next int `json:"next"`
	// Views 视图数据列表
	Views []any `json:"views"`
}

// UpdateFieldGroup 更新编组
// 本接口用于在智能表中的某个子表里更新已有编组。每个编组最多允许有150个字段。字段只能同时存在于一个编组。
//
// see https://developer.work.weixin.qq.com/document/path/101101
func (c *Client) UpdateFieldGroup(r *UpdateFieldGroupRequest, opts ...any) (out UpdateFieldGroupResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/update_field_group",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateFieldGroupRequest is request of Client.UpdateFieldGroup
type UpdateFieldGroupRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID 表格ID
	SheetID string `json:"sheet_id" validate:"required"`
	// FieldGroupID 编组id
	FieldGroupID string `json:"field_group_id" validate:"required"`
	// Name 编组名称，不能和已有名称重复
	Name string `json:"name"`
	// Children 编组内容
	Children []map[string]any `json:"children"`
}

// UpdateFieldGroupResponse is response of Client.UpdateFieldGroup
type UpdateFieldGroupResponse struct {
	// FieldGroupID 编组id
	FieldGroupID string `json:"field_group_id"`
	// Name 编组名称
	Name string `json:"name"`
	// Children 编组内容
	Children []any `json:"children"`
}

// UpdateSmartSheetFields 更新字段
// 本接口用于更新智能表格中的某个子表里的一个或多个字段的标题和字段属性信息。
//
// see https://developer.work.weixin.qq.com/document/path/99906
func (c *Client) UpdateSmartSheetFields(r *UpdateSmartSheetFieldsRequest, opts ...any) (out UpdateSmartSheetFieldsResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/update_fields",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateSmartSheetFieldsRequest is request of Client.UpdateSmartSheetFields
type UpdateSmartSheetFieldsRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID 表格ID
	SheetID string `json:"sheet_id" validate:"required"`
	// Fields 字段详情列表
	Fields []map[string]any `json:"fields" validate:"required"`
}

// UpdateSmartSheetFieldsResponse is response of Client.UpdateSmartSheetFields
type UpdateSmartSheetFieldsResponse struct {
	// Fields 更新后的字段详情列表
	Fields []map[string]any `json:"fields"`
}

// UpdateSmartsheetRecords 更新记录
// 本接口用于更新 Smartsheet 中的某个子表里的一行或多行记录。单次更新建议在500行内
//
// see https://developer.work.weixin.qq.com/document/path/99909
func (c *Client) UpdateSmartsheetRecords(r *UpdateSmartsheetRecordsRequest, opts ...any) (out UpdateSmartsheetRecordsResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/update_records",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateSmartsheetRecordsRequest is request of Client.UpdateSmartsheetRecords
type UpdateSmartsheetRecordsRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID Smartsheet 子表ID
	SheetID string `json:"sheet_id" validate:"required"`
	// KeyType 返回记录中单元格的key类型
	KeyType string `json:"key_type"`
	// Records 由需要更新的记录组成的 JSON 数组
	Records []any `json:"records" validate:"required"`
}

// UpdateSmartsheetRecordsResponse is response of Client.UpdateSmartsheetRecords
type UpdateSmartsheetRecordsResponse struct {
	// Records 由更新成功的记录的具体内容组成的 JSON 数组
	Records []any `json:"records"`
}

// UpdateSmartsheetSubSheet 更新子表
// 本接口用于修改表格中某个子表的标题。
//
// see https://developer.work.weixin.qq.com/document/path/99898
func (c *Client) UpdateSmartsheetSubSheet(r *UpdateSmartsheetSubSheetRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/update_sheet",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateSmartsheetSubSheetRequest is request of Client.UpdateSmartsheetSubSheet
type UpdateSmartsheetSubSheetRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// Properties 子表属性对象
	Properties any `json:"properties"`
}

// UpdateView 更新视图
// 本接口用于更新 Smartsheet 中的某个视图。
//
// see https://developer.work.weixin.qq.com/document/path/99902
func (c *Client) UpdateView(r *UpdateViewRequest, opts ...any) (out UpdateViewResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/update_view",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateViewRequest is request of Client.UpdateView
type UpdateViewRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// SheetID Smartsheet 子表ID
	SheetID string `json:"sheet_id" validate:"required"`
	// ViewID 视图ID
	ViewID string `json:"view_id" validate:"required"`
	// ViewTitle 视图标题
	ViewTitle string `json:"view_title"`
	// Property 视图的排序/过滤/分组/填色配置，详见ViewProperty
	Property any `json:"property"`
}

// UpdateViewResponse is response of Client.UpdateView
type UpdateViewResponse struct {
	// View 更新成功的视图内容
	View any `json:"view"`
}

// UpdateSmartSheetRecord 更新智能表格记录
// 通过Webhook地址更新一行或多行记录
//
// see https://developer.work.weixin.qq.com/document/path/101260
func (c *Client) UpdateSmartSheetRecord(r *UpdateSmartSheetRecordRequest, opts ...any) (out UpdateSmartSheetRecordResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/webhook",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateSmartSheetRecordRequest is request of Client.UpdateSmartSheetRecord
type UpdateSmartSheetRecordRequest struct {
	// Key Webhook密钥，从开启接收外部数据获取
	Key string `json:"key" validate:"required"`
}

// UpdateSmartSheetRecordResponse is response of Client.UpdateSmartSheetRecord
type UpdateSmartSheetRecordResponse struct {
	// UpdateRecords 由更新成功的记录的具体内容组成的 JSON 数组
	UpdateRecords []any `json:"update_records"`
}

// BatchUpdateSpreadsheet 编辑表格内容
// 该接口可以对一个在线表格批量执行多个更新操作。
//
// see https://developer.work.weixin.qq.com/document/path/101168
func (c *Client) BatchUpdateSpreadsheet(r *BatchUpdateSpreadsheetRequest, opts ...any) (out BatchUpdateSpreadsheetResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/spreadsheet/batch_update",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchUpdateSpreadsheetRequest is request of Client.BatchUpdateSpreadsheet
type BatchUpdateSpreadsheetRequest struct {
	// Docid 文档的docid
	Docid string `json:"docid" validate:"required"`
	// Requests 更新操作列表
	Requests []map[string]any `json:"requests" validate:"required"`
}

// BatchUpdateSpreadsheetResponse is response of Client.BatchUpdateSpreadsheet
type BatchUpdateSpreadsheetResponse struct {
	// Data 返回数据对象
	Data any `json:"data"`
}

// GetSpreadsheetProperties 获取表格行列信息
// 该接口用于获取在线表格的工作表、行数、列数等。
//
// see https://developer.work.weixin.qq.com/document/path/97711
func (c *Client) GetSpreadsheetProperties(r *GetSpreadsheetPropertiesRequest, opts ...any) (out GetSpreadsheetPropertiesResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/spreadsheet/get_sheet_properties",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetSpreadsheetPropertiesRequest is request of Client.GetSpreadsheetProperties
type GetSpreadsheetPropertiesRequest struct {
	// Docid 在线表格的docid
	Docid string `json:"docid" validate:"required"`
}

// GetSpreadsheetPropertiesResponse is response of Client.GetSpreadsheetProperties
type GetSpreadsheetPropertiesResponse struct {
	// Properties 工作表属性列表
	Properties []map[string]any `json:"properties"`
}

// GetSheetRangeData 获取表格数据
// 本接口用于获取指定范围内的在线表格信息，单次查询的范围大小需满足：行数<=1000，列数<=200，总单元格数量<=10000。
//
// see https://developer.work.weixin.qq.com/document/path/97661
func (c *Client) GetSheetRangeData(r *GetSheetRangeDataRequest, opts ...any) (out GetSheetRangeDataResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/spreadsheet/get_sheet_range_data",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetSheetRangeDataRequest is request of Client.GetSheetRangeData
type GetSheetRangeDataRequest struct {
	// Docid 在线表格唯一标识
	Docid string `json:"docid" validate:"required"`
	// SheetID 工作表ID，工作表的唯一标识
	SheetID string `json:"sheet_id" validate:"required"`
	// Range 查询的范围，格式遵循A1表示法
	Range string `json:"range" validate:"required"`
}

// GetSheetRangeDataResponse is response of Client.GetSheetRangeData
type GetSheetRangeDataResponse struct {
	// Data 表格数据对象
	Data any `json:"data"`
}

// BatchAddWedDocVip 分配高级功能账号
// 该接口用于分配应用可见范围内企业成员的高级功能。
//
// see https://developer.work.weixin.qq.com/document/path/99516
func (c *Client) BatchAddWedDocVip(r *BatchAddWedDocVipRequest, opts ...any) (out BatchAddWedDocVipResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/vip/batch_add",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchAddWedDocVipRequest is request of Client.BatchAddWedDocVip
type BatchAddWedDocVipRequest struct {
	// UserIDList 要分配高级功能的企业成员userid列表，单次操作最大限制100个
	UserIDList []string `json:"userid_list" validate:"required"`
}

// BatchAddWedDocVipResponse is response of Client.BatchAddWedDocVip
type BatchAddWedDocVipResponse struct {
	// SuccUserIDList 分配成功的userid列表，包括已经是高级功能账号的userid
	SuccUserIDList []string `json:"succ_userid_list"`
	// FailUserIDList 分配失败的userid列表
	FailUserIDList []string `json:"fail_userid_list"`
}

// BatchDeleteDocVip 取消高级功能账号
// 撤销分配应用可见范围企业成员的高级功能
//
// see https://developer.work.weixin.qq.com/document/path/99517
func (c *Client) BatchDeleteDocVip(r *BatchDeleteDocVipRequest, opts ...any) (out BatchDeleteDocVipResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/vip/batch_del",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchDeleteDocVipRequest is request of Client.BatchDeleteDocVip
type BatchDeleteDocVipRequest struct {
	// UserIDList 要撤销分配高级功能的企业成员userid列表，单次操作最多限制100个
	UserIDList []string `json:"userid_list" validate:"required"`
}

// BatchDeleteDocVipResponse is response of Client.BatchDeleteDocVip
type BatchDeleteDocVipResponse struct {
	// SuccUserIDList 撤销分配成功的userid列表
	SuccUserIDList []string `json:"succ_userid_list"`
	// FailUserIDList 撤销分配失败的userid列表
	FailUserIDList []string `json:"fail_userid_list"`
}

// ListVipAccount 获取高级功能账号列表
// 查询企业已分配高级功能且在应用可见范围的账号列表
//
// see https://developer.work.weixin.qq.com/document/path/99518
func (c *Client) ListVipAccount(r *ListVipAccountRequest, opts ...any) (out ListVipAccountResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/vip/list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListVipAccountRequest is request of Client.ListVipAccount
type ListVipAccountRequest struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"`
	// Limit 用于分页查询，每次请求返回的数据上限。默认100，最大200
	Limit int `json:"limit"`
}

// ListVipAccountResponse is response of Client.ListVipAccount
type ListVipAccountResponse struct {
	// HasMore 是否还有更多数据未获取
	HasMore bool `json:"has_more"`
	// NextCursor 下一次请求的cursor值
	NextCursor string `json:"next_cursor"`
	// UserIDList 符合条件的企业成员userid列表
	UserIDList []string `json:"userid_list"`
}

