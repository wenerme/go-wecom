package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// AddInterceptRule 新建敏感词规则
// 企业和第三方应用可以通过此接口新建敏感词规则
//
// see https://developer.work.weixin.qq.com/document/path/95100
func (c *Client) AddInterceptRule(r *AddInterceptRuleRequest, opts ...any) (out AddInterceptRuleResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/add_intercept_rule",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddInterceptRuleRequest is request of Client.AddInterceptRule
type AddInterceptRuleRequest struct {
	// RuleName 规则名称，长度1~20个utf8字符
	RuleName string `json:"rule_name" validate:"required"`
	// WordList 敏感词列表，敏感词长度1~32个utf8字符，列表大小不能超过300个
	WordList []string `json:"word_list" validate:"required"`
	// SemanticsList 额外的拦截语义规则，1：手机号、2：邮箱地址、3：红包
	SemanticsList []int `json:"semantics_list"`
	// InterceptType 拦截方式，1:警告并拦截发送；2:仅发警告
	InterceptType int `json:"intercept_type" validate:"required"`
	// ApplicableRange 敏感词适用范围，userid与department不能同时为不填
	ApplicableRange any `json:"applicable_range" validate:"required"`
}

// AddInterceptRuleResponse is response of Client.AddInterceptRule
type AddInterceptRuleResponse struct {
	// RuleID 规则id
	RuleID string `json:"rule_id"`
}

// AddMomentTask 创建发表任务
// 企业和第三方应用可通过该接口创建客户朋友圈的发表任务。
//
// see https://developer.work.weixin.qq.com/document/path/95173
func (c *Client) AddMomentTask(r *AddMomentTaskRequest, opts ...any) (out AddMomentTaskResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/add_moment_task",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddMomentTaskRequest is request of Client.AddMomentTask
type AddMomentTaskRequest struct {
	// Text 文本消息，不能与附件同时为空
	Text any `json:"text"`
	// Attachments 附件，不能与text.content同时为空，最多支持9个图片、1个视频或1个链接
	Attachments []map[string]any `json:"attachments"`
	// VisibleRange 指定的发表范围；若未指定，则表示执行者为应用可见范围内所有成员
	VisibleRange any `json:"visible_range"`
}

// AddMomentTaskResponse is response of Client.AddMomentTask
type AddMomentTaskResponse struct {
	// JobID 异步任务id，最大长度为64字节，24小时有效
	JobID string `json:"jobid"`
}

// AddProductAlbum 管理商品图册
// 企业与第三方应用可通过该接口管理商品图册（创建、获取、列表、编辑、删除）
//
// see https://developer.work.weixin.qq.com/document/path/95101
func (c *Client) AddProductAlbum(r *AddProductAlbumRequest, opts ...any) (out AddProductAlbumResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/add_product_album",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddProductAlbumRequest is request of Client.AddProductAlbum
type AddProductAlbumRequest struct {
	// Description 商品的名称、特色等;不超过300个字
	Description string `json:"description" validate:"required"`
	// Price 商品的价格，单位为分；最大不超过5万元
	Price int `json:"price" validate:"required"`
	// ProductSn 商品编码；不超过128个字节；只能输入数字和字母
	ProductSn string `json:"product_sn"`
	// Attachments 附件类型，仅支持image，最多不超过9个附件
	Attachments []map[string]any `json:"attachments" validate:"required"`
}

// AddProductAlbumResponse is response of Client.AddProductAlbum
type AddProductAlbumResponse struct {
	// ProductID 商品id
	ProductID string `json:"product_id"`
}

// CancelGroupMsgSend 停止企业群发
// 企业和第三方应用可调用此接口，停止无需成员继续发送的企业群发
//
// see https://developer.work.weixin.qq.com/document/path/97611
func (c *Client) CancelGroupMsgSend(r *CancelGroupMsgSendRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/cancel_groupmsg_send",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CancelGroupMsgSendRequest is request of Client.CancelGroupMsgSend
type CancelGroupMsgSendRequest struct {
	// MsgID 群发消息的id，通过获取群发记录列表接口返回
	MsgID string `json:"msgid" validate:"required"`
}

// CancelMomentTask 停止发表企业朋友圈
// 企业和第三方应用可调用此接口，停止尚未发送的企业朋友圈发送任务。
//
// see https://developer.work.weixin.qq.com/document/path/97612
func (c *Client) CancelMomentTask(r *CancelMomentTaskRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/cancel_moment_task",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CancelMomentTaskRequest is request of Client.CancelMomentTask
type CancelMomentTaskRequest struct {
	// MomentID 朋友圈id，可通过获取客户朋友圈企业发表的列表接口获取朋友圈企业发表的列表
	MomentID string `json:"moment_id" validate:"required"`
}

// GetCustomerAcquisitionQuota 获客助手额度管理与使用统计
// 查询剩余使用量及链接使用详情
//
// see https://developer.work.weixin.qq.com/document/path/99337
func (c *Client) GetCustomerAcquisitionQuota(r *GetCustomerAcquisitionQuotaRequest, opts ...any) (out GetCustomerAcquisitionQuotaResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/externalcontact/customer_acquisition_quota",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCustomerAcquisitionQuotaRequest is request of Client.GetCustomerAcquisitionQuota
type GetCustomerAcquisitionQuotaRequest struct {}

// GetCustomerAcquisitionQuotaResponse is response of Client.GetCustomerAcquisitionQuota
type GetCustomerAcquisitionQuotaResponse struct {
	// Total 历史累计使用量
	Total int `json:"total"`
	// Balance 剩余使用量
	Balance int `json:"balance"`
	// QuotaList 额度列表
	QuotaList []map[string]any `json:"quota_list"`
	// ExpireDate 额度过期时间戳，为过期日的零点
	ExpireDate int `json:"expire_date"`
}

// ListCustomerAcquisition 获取获客客户列表
// 企业可通过此接口获取到由指定的获客链接添加的客户列表。
//
// see https://developer.work.weixin.qq.com/document/path/97305
func (c *Client) ListCustomerAcquisition(r *ListCustomerAcquisitionRequest, opts ...any) (out ListCustomerAcquisitionResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/customer_acquisition/customer",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListCustomerAcquisitionRequest is request of Client.ListCustomerAcquisition
type ListCustomerAcquisitionRequest struct {
	// LinkID 获客链接id。需要是当前应用创建
	LinkID string `json:"link_id" validate:"required"`
	// Limit 返回的最大记录数，整型，最大值1000
	Limit int `json:"limit"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"`
}

// ListCustomerAcquisitionResponse is response of Client.ListCustomerAcquisition
type ListCustomerAcquisitionResponse struct {
	// CustomerList 客户列表，包含external_userid, userid, chat_status, state字段
	CustomerList []map[string]any `json:"customer_list"`
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录
	NextCursor string `json:"next_cursor"`
}

// GetCustomerAcquisitionChatInfo 获取成员多次收消息详情
// 企业和服务商可通过此接口获取成员多次收消息情况，如次数、客户id等信息。
//
// see https://developer.work.weixin.qq.com/document/path/100254
func (c *Client) GetCustomerAcquisitionChatInfo(r *GetCustomerAcquisitionChatInfoRequest, opts ...any) (out GetCustomerAcquisitionChatInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/customer_acquisition/get_chat_info",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCustomerAcquisitionChatInfoRequest is request of Client.GetCustomerAcquisitionChatInfo
type GetCustomerAcquisitionChatInfoRequest struct {
	// ChatKey [成员多次收消息事件](#42924/成员多次收消息事件)中回调的会话信息凭据ChatKey，回调后30分钟内有效
	ChatKey string `json:"chat_key" validate:"required"`
}

// GetCustomerAcquisitionChatInfoResponse is response of Client.GetCustomerAcquisitionChatInfo
type GetCustomerAcquisitionChatInfoResponse struct {
	// UserID 成员的userid
	UserID string `json:"userid"`
	// ExternalUserID 客户id
	ExternalUserID string `json:"external_userid"`
	// ChatInfo 聊天信息对象
	ChatInfo any `json:"chat_info"`
}

// ListCustomerAcquisitionLinks 获客链接管理
// 获取当前仍然有效且是当前应用创建的获客链接列表。
//
// see https://developer.work.weixin.qq.com/document/path/97297
func (c *Client) ListCustomerAcquisitionLinks(r *ListCustomerAcquisitionLinksRequest, opts ...any) (out ListCustomerAcquisitionLinksResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/customer_acquisition/list_link",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListCustomerAcquisitionLinksRequest is request of Client.ListCustomerAcquisitionLinks
type ListCustomerAcquisitionLinksRequest struct {
	// Limit 返回的最大记录数，整型，最大值100
	Limit int `json:"limit"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"`
}

// ListCustomerAcquisitionLinksResponse is response of Client.ListCustomerAcquisitionLinks
type ListCustomerAcquisitionLinksResponse struct {
	// LinkIDList link_id列表
	LinkIDList []string `json:"link_id_list"`
	// NextCursor 分页游标，在下次请求时填写以获取之后分页的记录
	NextCursor string `json:"next_cursor"`
}

// ListCustomerStrategy 获取规则组列表
// 企业可通过此接口获取企业配置的所有客户规则组id列表。
//
// see https://developer.work.weixin.qq.com/document/path/94883
func (c *Client) ListCustomerStrategy(r *ListCustomerStrategyRequest, opts ...any) (out ListCustomerStrategyResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/customer_strategy/list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListCustomerStrategyRequest is request of Client.ListCustomerStrategy
type ListCustomerStrategyRequest struct {
	// Cursor 分页查询游标，首次调用可不填
	Cursor string `json:"cursor"`
	// Limit 分页大小，默认为1000，最大不超过1000
	Limit int `json:"limit"`
}

// ListCustomerStrategyResponse is response of Client.ListCustomerStrategy
type ListCustomerStrategyResponse struct {
	// Strategy 规则组列表，包含strategy_id
	Strategy []map[string]any `json:"strategy"`
	// NextCursor 分页游标，用于查询下一个分页的数据，无更多数据时不返回
	NextCursor string `json:"next_cursor"`
}

// ListGroupMsg 获取企业的全部群发记录
// 企业和第三方应用可通过此接口获取企业与成员的群发记录列表。
//
// see https://developer.work.weixin.qq.com/document/path/93539
func (c *Client) ListGroupMsg(r *ListGroupMsgRequest, opts ...any) (out ListGroupMsgResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/get_groupmsg_list_v2",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListGroupMsgRequest is request of Client.ListGroupMsg
type ListGroupMsgRequest struct {
	// ChatType 群发任务的类型，single表示发送给客户，group表示发送给客户群
	ChatType string `json:"chat_type" validate:"required"`
	// StartTime 群发任务记录开始时间
	StartTime int `json:"start_time" validate:"required"`
	// EndTime 群发任务记录结束时间
	EndTime int `json:"end_time" validate:"required"`
	// Creator 群发任务创建人企业账号id
	Creator string `json:"creator"`
	// FilterType 创建人类型。0：企业发表 1：个人发表 2：所有，默认值为2
	FilterType int `json:"filter_type"`
	// Limit 返回的最大记录数，最大值100，默认值50
	Limit int `json:"limit"`
	// Cursor 用于分页查询的游标，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"`
}

// ListGroupMsgResponse is response of Client.ListGroupMsg
type ListGroupMsgResponse struct {
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录
	NextCursor string `json:"next_cursor"`
	// GroupMsgList 群发记录列表
	GroupMsgList []map[string]any `json:"group_msg_list"`
}

// ListMoment 获取客户朋友圈全部的发表记录
// 企业和第三方应用可通过该接口获取企业全部的发表内容。
//
// see https://developer.work.weixin.qq.com/document/path/93545
func (c *Client) ListMoment(r *ListMomentRequest, opts ...any) (out ListMomentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/get_moment_list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMomentRequest is request of Client.ListMoment
type ListMomentRequest struct {
	// StartTime 朋友圈记录开始时间。Unix时间戳
	StartTime int `json:"start_time" validate:"required"`
	// EndTime 朋友圈记录结束时间。Unix时间戳
	EndTime int `json:"end_time" validate:"required"`
	// Creator 朋友圈创建人的userid
	Creator string `json:"creator"`
	// FilterType 朋友圈类型。0：企业发表 1：个人发表 2：所有，包括个人创建以及企业创建，默认情况下为所有类型
	FilterType int `json:"filter_type"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"`
	// Limit 返回的最大记录数，整型，最大值20，默认值20，超过最大值时取默认值
	Limit int `json:"limit"`
}

// ListMomentResponse is response of Client.ListMoment
type ListMomentResponse struct {
	// NextCursor 分页游标，下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"`
	// MomentList 朋友圈列表
	MomentList []map[string]any `json:"moment_list"`
}

// GetStrategyTagList 获取指定规则组下的企业客户标签
// 企业可通过此接口获取某个规则组内的企业客户标签详情。
//
// see https://developer.work.weixin.qq.com/document/path/94882
func (c *Client) GetStrategyTagList(r *GetStrategyTagListRequest, opts ...any) (out GetStrategyTagListResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/get_strategy_tag_list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetStrategyTagListRequest is request of Client.GetStrategyTagList
type GetStrategyTagListRequest struct {
	// StrategyID 规则组id
	StrategyID int `json:"strategy_id"`
	// TagID 要查询的标签id
	TagID []string `json:"tag_id"`
	// GroupID 要查询的标签组id，返回该标签组以及其下的所有标签信息
	GroupID []string `json:"group_id"`
}

// GetStrategyTagListResponse is response of Client.GetStrategyTagList
type GetStrategyTagListResponse struct {
	// TagGroup 标签组列表
	TagGroup []map[string]any `json:"tag_group"`
}

// AddJoinWay 配置客户群进群方式
// 企业可调用此接口来生成并配置「加入群聊」的二维码或者小程序按钮，客户通过扫描二维码或点击小程序上的按钮，即可加入特定的客户群。如果配置的是小程序按钮，需要开发者的小程序接入小程序插件。
//
// see https://developer.work.weixin.qq.com/document/path/92569
func (c *Client) AddJoinWay(r *AddJoinWayRequest, opts ...any) (out AddJoinWayResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/groupchat/add_join_way",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddJoinWayRequest is request of Client.AddJoinWay
type AddJoinWayRequest struct {
	// Scene 场景。1 - 群的小程序插件; 2 - 群的二维码插件
	Scene int `json:"scene" validate:"required"`
	// Remark 联系方式的备注信息，用于助记，超过30个字符将被截断
	Remark string `json:"remark"`
	// AutoCreateRoom 当群满了后，是否自动新建群。0-否；1-是。默认为1
	AutoCreateRoom int `json:"auto_create_room"`
	// RoomBaseName 自动建群的群名前缀，当auto_create_room为1时有效。最长40个utf8字符
	RoomBaseName string `json:"room_base_name"`
	// RoomBaseID 自动建群的群起始序号，当auto_create_room为1时有效
	RoomBaseID int `json:"room_base_id"`
	// ChatIDList 使用该配置的客户群ID列表，最多支持5个
	ChatIDList []string `json:"chat_id_list" validate:"required"`
	// State 企业自定义的state参数，用于区分不同的入群渠道。不超过30个UTF-8字符
	State string `json:"state"`
	// MarkSource 是否标记客户添加来源为该应用创建的「加入群聊」, 默认值为true; 仅对「营销获客」应用生效
	MarkSource bool `json:"mark_source"`
}

// AddJoinWayResponse is response of Client.AddJoinWay
type AddJoinWayResponse struct {
	// ConfigID 配置id
	ConfigID string `json:"config_id"`
}

// TransferGroupChatOwner 分配在职成员的客户群
// 企业可通过此接口，将在职成员为群主的群，分配给另一个客服成员。
//
// see https://developer.work.weixin.qq.com/document/path/95703
func (c *Client) TransferGroupChatOwner(r *TransferGroupChatOwnerRequest, opts ...any) (out TransferGroupChatOwnerResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/groupchat/onjob_transfer",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// TransferGroupChatOwnerRequest is request of Client.TransferGroupChatOwner
type TransferGroupChatOwnerRequest struct {
	// ChatIDList 需要转群主的客户群ID列表。取值范围：1 ~ 100
	ChatIDList []string `json:"chat_id_list" validate:"required"`
	// NewOwner 新群主ID
	NewOwner string `json:"new_owner" validate:"required"`
}

// TransferGroupChatOwnerResponse is response of Client.TransferGroupChatOwner
type TransferGroupChatOwnerResponse struct {
	// FailedChatList 没能成功继承的群
	FailedChatList []map[string]any `json:"failed_chat_list"`
}

// GetGroupChatStatistic 获取「群聊数据统计」数据
// 获取指定日期的群聊统计数据，支持按群主聚合和按自然日聚合两种模式。
//
// see https://developer.work.weixin.qq.com/document/path/93559
func (c *Client) GetGroupChatStatistic(r *GetGroupChatStatisticRequest, opts ...any) (out GetGroupChatStatisticResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/groupchat/statistic",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetGroupChatStatisticRequest is request of Client.GetGroupChatStatistic
type GetGroupChatStatisticRequest struct {
	// DayBeginTime 起始日期的时间戳，填当天的0时0分0秒。取值范围：昨天至前180天
	DayBeginTime int `json:"day_begin_time" validate:"required"`
	// DayEndTime 结束日期的时间戳，填当天的0时0分0秒。如果不填，默认同 day_begin_time。取值范围：昨天至前180天
	DayEndTime int `json:"day_end_time"`
	// OwnerFilter 群主过滤对象
	OwnerFilter any `json:"owner_filter" validate:"required"`
	// OrderBy 排序方式：1-新增群的数量，2-群总数，3-新增群人数，4-群总人数。默认为1
	OrderBy int `json:"order_by"`
	// OrderAsc 是否升序。0-否；1-是。默认降序
	OrderAsc int `json:"order_asc"`
	// Offset 分页，偏移量，默认为0
	Offset int `json:"offset"`
	// Limit 分页，预期请求的数据量，默认为500，取值范围 1 ~ 1000
	Limit int `json:"limit"`
}

// GetGroupChatStatisticResponse is response of Client.GetGroupChatStatistic
type GetGroupChatStatisticResponse struct {
	// Total 命中过滤条件的记录总个数
	Total int `json:"total"`
	// NextOffset 当前分页的下一个offset。当next_offset和total相等时，说明已经取完所有
	NextOffset int `json:"next_offset"`
	// Items 记录列表。表示某个群主所拥有的客户群的统计数据
	Items []map[string]any `json:"items"`
}

// ListMomentStrategy 获取规则组列表
// 企业可通过此接口获取企业配置的所有客户朋友圈规则组id列表。
//
// see https://developer.work.weixin.qq.com/document/path/94891
func (c *Client) ListMomentStrategy(r *ListMomentStrategyRequest, opts ...any) (out ListMomentStrategyResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/moment_strategy/list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMomentStrategyRequest is request of Client.ListMomentStrategy
type ListMomentStrategyRequest struct {
	// Cursor 分页查询游标，首次调用可不填
	Cursor string `json:"cursor"`
	// Limit 分页大小，默认为1000，最大不超过1000
	Limit int `json:"limit"`
}

// ListMomentStrategyResponse is response of Client.ListMomentStrategy
type ListMomentStrategyResponse struct {
	// Strategy 规则组id列表，每个元素包含strategy_id
	Strategy []map[string]any `json:"strategy"`
	// NextCursor 分页游标，用于查询下一个分页的数据，无更多数据时不返回
	NextCursor string `json:"next_cursor"`
}

// OpenGidToChatId 客户群opengid转换
// 将小程序获取的群opengid转换为企业微信客户群的chat_id
//
// see https://developer.work.weixin.qq.com/document/path/94822
func (c *Client) OpenGidToChatId(r *OpenGidToChatIdRequest, opts ...any) (out OpenGidToChatIdResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/opengid_to_chatid",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// OpenGidToChatIdRequest is request of Client.OpenGidToChatId
type OpenGidToChatIdRequest struct {
	// Opengid 小程序在微信获取到的群ID
	Opengid string `json:"opengid" validate:"required"`
}

// OpenGidToChatIdResponse is response of Client.OpenGidToChatId
type OpenGidToChatIdResponse struct {
	// ChatID 客户群ID
	ChatID string `json:"chat_id"`
}

// RemindGroupMsgSend 提醒成员群发
// 企业和第三方应用可调用此接口，重新触发群发通知，提醒成员完成群发任务
//
// see https://developer.work.weixin.qq.com/document/path/97610
func (c *Client) RemindGroupMsgSend(r *RemindGroupMsgSendRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/remind_groupmsg_send",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// RemindGroupMsgSendRequest is request of Client.RemindGroupMsgSend
type RemindGroupMsgSendRequest struct {
	// MsgID 群发消息的id，通过获取群发记录列表接口返回
	MsgID string `json:"msgid" validate:"required"`
}

// UploadAttachment 上传附件资源
// 上传素材到企业微信，获取media_id。该media_id仅三天内有效。
//
// see https://developer.work.weixin.qq.com/document/path/95105
func (c *Client) UploadAttachment(r *UploadAttachmentRequest, opts ...any) (out UploadAttachmentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/media/upload_attachment",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UploadAttachmentRequest is request of Client.UploadAttachment
type UploadAttachmentRequest struct {
	// MediaType 媒体文件类型，分别有图片（image）、视频（video）、普通文件（file）
	MediaType string `json:"media_type" validate:"required"`
	// AttachmentType 附件类型。1：朋友圈；2:商品图册
	AttachmentType int `json:"attachment_type" validate:"required"`
}

// UploadAttachmentResponse is response of Client.UploadAttachment
type UploadAttachmentResponse struct {
	// Type 媒体文件类型，分别有图片（image）、语音（voice）、视频（video），普通文件(file)
	Type string `json:"type"`
	// MediaID 媒体文件上传后获取的唯一标识，三天有效
	MediaID string `json:"media_id"`
	// CreatedAt 媒体文件上传时间戳
	CreatedAt int `json:"created_at"`
}

