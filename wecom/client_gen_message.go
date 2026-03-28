package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// CreateGroupChat 创建群聊会话
// 通过自建应用创建群聊会话
//
// see https://developer.work.weixin.qq.com/document/path/90068
func (c *Client) CreateGroupChat(r *CreateGroupChatRequest, opts ...any) (out CreateGroupChatResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/appchat/create",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateGroupChatRequest is request of Client.CreateGroupChat
type CreateGroupChatRequest struct {
	// Name 群聊名，最多50个utf8字符
	Name string `json:"name"`
	// Owner 指定群主的id
	Owner string `json:"owner"`
	// Userlist 群成员id列表，至少2人，至多2000人
	Userlist []string `json:"userlist" validate:"required"`
	// ChatID 群聊的唯一标志，最长32个字符
	ChatID string `json:"chatid"`
}

// CreateGroupChatResponse is response of Client.CreateGroupChat
type CreateGroupChatResponse struct {
	// ChatID 群聊的唯一标志
	ChatID string `json:"chatid"`
}

// GetAppChat 获取群聊会话
// 获取企业自建应用创建的群聊信息
//
// see https://developer.work.weixin.qq.com/document/path/98914
func (c *Client) GetAppChat(r *GetAppChatRequest, opts ...any) (out GetAppChatResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/appchat/get",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetAppChatRequest is request of Client.GetAppChat
type GetAppChatRequest struct {
	// ChatID 群聊id
	ChatID string `json:"chatid" validate:"required"`
}

// GetAppChatResponse is response of Client.GetAppChat
type GetAppChatResponse struct {
	// ChatInfo 群聊信息
	ChatInfo any `json:"chat_info"`
}

// SendAppChatMessage 应用推送消息
// 企业自建应用向群聊推送文本、图片、视频、文件、图文等类型消息
//
// see https://developer.work.weixin.qq.com/document/path/90071
func (c *Client) SendAppChatMessage(r *SendAppChatMessageRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/appchat/send",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SendAppChatMessageRequest is request of Client.SendAppChatMessage
type SendAppChatMessageRequest struct {
	// ChatID 群聊ID
	ChatID string `json:"chatid" validate:"required"`
	// Msgtype 消息类型 (text, image, voice, video, file, textcard, news, mpnews, markdown等)
	Msgtype string `json:"msgtype" validate:"required"`
	// Safe 是否保密消息，0表示否，1表示是，默认0
	Safe int `json:"safe"`
	// Text 文本消息内容对象 (包含content和mentioned_list)
	Text any `json:"text"`
	// Image 图片消息内容对象 (包含media_id)
	Image any `json:"image"`
	// Voice 语音消息内容对象 (包含media_id)
	Voice any `json:"voice"`
	// Video 视频消息内容对象 (包含media_id, title, description)
	Video any `json:"video"`
	// File 文件消息内容对象 (包含media_id)
	File any `json:"file"`
	// Textcard 文本卡片消息内容对象 (包含title, description, url, btntxt)
	Textcard any `json:"textcard"`
	// News 图文消息内容对象 (包含articles数组)
	News any `json:"news"`
	// Mpnews 图文消息(mpnews)内容对象 (包含articles数组)
	Mpnews any `json:"mpnews"`
	// Markdown Markdown消息内容对象 (包含content)
	Markdown any `json:"markdown"`
}

// UpdateAppChat 修改群聊会话
// 修改企业自建应用创建的群聊信息，包括群名、群主及成员列表。
//
// see https://developer.work.weixin.qq.com/document/path/98913
func (c *Client) UpdateAppChat(r *UpdateAppChatRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/appchat/update",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateAppChatRequest is request of Client.UpdateAppChat
type UpdateAppChatRequest struct {
	// ChatID 群聊id
	ChatID string `json:"chatid" validate:"required"`
	// Name 新的群聊名。若不需更新，请忽略此参数。最多50个utf8字符，超过将截断
	Name string `json:"name"`
	// Owner 新群主的id。若不需更新，请忽略此参数。课程群聊群主必须拥有课程群创建权限，del_user_list包含群主时本字段必填
	Owner string `json:"owner"`
	// AddUserList 添加成员的id列表
	AddUserList []string `json:"add_user_list"`
	// DelUserList 踢出成员的id列表
	DelUserList []string `json:"del_user_list"`
}

// VerifyJoinQrCodeUrl 验证URL有效性
// 企业微信发送验证请求到开发者配置的URL，开发者需校验签名并返回解密后的明文内容以完成验证。
//
// see https://developer.work.weixin.qq.com/document/path/101034
func (c *Client) VerifyJoinQrCodeUrl(r *VerifyJoinQrCodeUrlRequest, opts ...any) (out VerifyJoinQrCodeUrlResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/externalcontact/getjoinqrcode",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// VerifyJoinQrCodeUrlRequest is request of Client.VerifyJoinQrCodeUrl
type VerifyJoinQrCodeUrlRequest struct {
	// MsgSignature 企业微信加密签名
	MsgSignature string `json:"msg_signature" validate:"required"`
	// Timestamp 时间戳
	Timestamp string `json:"timestamp" validate:"required"`
	// Nonce 随机数
	Nonce string `json:"nonce" validate:"required"`
}

// VerifyJoinQrCodeUrlResponse is response of Client.VerifyJoinQrCodeUrl
type VerifyJoinQrCodeUrlResponse struct {
	// EchostrDecrypted 解密后的明文消息内容（响应体）
	EchostrDecrypted string `json:"echostr_decrypted"`
}

// SendExternalContactMessage 发送学校通知
// 学校可以通过此接口来给家长发送不同类型的学校通知，目前支持的消息类型为文本、图片、语音、视频、文件、图文。
//
// see https://developer.work.weixin.qq.com/document/path/91609
func (c *Client) SendExternalContactMessage(r *SendExternalContactMessageRequest, opts ...any) (out SendExternalContactMessageResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/message/send",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SendExternalContactMessageRequest is request of Client.SendExternalContactMessage
type SendExternalContactMessageRequest struct {
	// RecvScope 指定发送对象，0表示发送给家长，1表示发送给学生，2表示发送给家长和学生，默认为0
	RecvScope int `json:"recv_scope"`
	// ToParentUserID 家校通讯录家长列表，recv_scope为0或2表示发送给对应的家长，最多支持1000个
	ToParentUserID []string `json:"to_parent_userid"`
	// ToStudentUserID 家校通讯录学生列表，recv_scope为0表示发送给学生的所有家长，recv_scope为1表示发送给学生，recv_scope为2表示发送给学生和学生的所有家长，最多支持1000个
	ToStudentUserID []string `json:"to_student_userid"`
	// ToParty 家校通讯录部门列表，recv_scope为0表示发送给班级的所有家长，recv_scope为1表示发送给班级的所有学生，recv_scope为2表示发送给班级的所有学生和家长，最多支持100个
	ToParty []string `json:"to_party"`
	// Toall 1表示字段生效，0表示字段无效。recv_scope为0表示发送给学校的所有家长，recv_scope为1表示发送给学校的所有学生，recv_scope为2表示发送给学校的所有学生和家长，默认为0
	Toall int `json:"toall"`
	// Msgtype 消息类型，固定为text/image/voice/video/file/news等
	Msgtype string `json:"msgtype" validate:"required"`
	// AgentID 企业应用的id，整型
	AgentID int `json:"agentid" validate:"required"`
	// Content 消息内容对象（文本消息时包含content字段），最长不超过2048个字节
	Content any `json:"content"`
	// Image 图片消息对象，包含media_id
	Image any `json:"image"`
	// Voice 语音消息对象，包含media_id
	Voice any `json:"voice"`
	// Video 视频消息对象，包含media_id, title, description
	Video any `json:"video"`
	// File 文件消息对象，包含media_id
	File any `json:"file"`
	// News 图文消息对象，包含articles数组
	News any `json:"news"`
	// EnableIDTrans 表示是否开启id转译，0表示否，1表示是，默认0
	EnableIDTrans int `json:"enable_id_trans"`
	// EnableDuplicateCheck 表示是否开启重复消息检查，0表示否，1表示是，默认0
	EnableDuplicateCheck int `json:"enable_duplicate_check"`
	// DuplicateCheckInterval 表示重复消息检查的时间间隔，默认1800s，最大不超过4小时
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

// SendExternalContactMessageResponse is response of Client.SendExternalContactMessage
type SendExternalContactMessageResponse struct {
	// InvalidParentUserID 无效的家长userid列表
	InvalidParentUserID []string `json:"invalid_parent_userid"`
	// InvalidStudentUserID 无效的学生userid列表
	InvalidStudentUserID []string `json:"invalid_student_userid"`
	// InvalidParty 无效的部门partyid列表
	InvalidParty []string `json:"invalid_party"`
}

// RecallAppMessage 撤回应用消息
// 撤回24小时内通过发送应用消息接口推送的消息，仅可撤回企业微信端的数据。
//
// see https://developer.work.weixin.qq.com/document/path/94867
func (c *Client) RecallAppMessage(r *RecallAppMessageRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/message/recall",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// RecallAppMessageRequest is request of Client.RecallAppMessage
type RecallAppMessageRequest struct {
	// MsgID 消息ID。从应用发送消息接口处获得。
	MsgID string `json:"msgid" validate:"required"`
}

// SendMessage 发送应用消息
// 推送文本、图片、视频、文件、图文等类型的企业应用消息
//
// see https://developer.work.weixin.qq.com/document/path/90060
func (c *Client) SendMessage(r *SendMessageRequest, opts ...any) (out SendMessageResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/message/send",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SendMessageRequest is request of Client.SendMessage
type SendMessageRequest struct {
	// Touser 指定接收消息的成员，成员ID列表（多个接收者用'|'分隔，最多支持1000个）。特殊情况：指定为"@all"，则向该企业应用的全部成员发送
	Touser string `json:"touser"`
	// Toparty 指定接收消息的部门，部门ID列表，多个接收者用'|'分隔，最多支持100个。当touser为"@all"时忽略本参数
	Toparty string `json:"toparty"`
	// Totag 指定接收消息的标签，标签ID列表，多个接收者用'|'分隔，最多支持100个。当touser为"@all"时忽略本参数
	Totag string `json:"totag"`
	// Msgtype 消息类型：text、image、voice、video、file、mpnews、news、markdown等
	Msgtype string `json:"msgtype" validate:"required"`
	// AgentID 企业应用的id，整型
	AgentID int `json:"agentid" validate:"required"`
	// Safe 表示是否是保密消息，0表示可对外分享，1表示不能分享且内容显示水印，默认为0
	Safe int `json:"safe"`
	// EnableIDTrans 表示是否开启id转译，0表示否，1表示是，默认0
	EnableIDTrans int `json:"enable_id_trans"`
	// EnableDuplicateCheck 表示是否开启重复消息检查，0表示否，1表示是，默认0
	EnableDuplicateCheck int `json:"enable_duplicate_check"`
	// DuplicateCheckInterval 表示重复消息检查的时间间隔，默认1800s，最大不超过4小时
	DuplicateCheckInterval int `json:"duplicate_check_interval"`
}

// SendMessageResponse is response of Client.SendMessage
type SendMessageResponse struct {
	// Invaliduser 不合法的userid，不区分大小写，统一转为小写
	Invaliduser string `json:"invaliduser"`
	// Invalidparty 不合法的partyid
	Invalidparty string `json:"invalidparty"`
	// Invalidtag 不合法的标签id
	Invalidtag string `json:"invalidtag"`
	// Unlicenseduser 没有基础接口许可(包含已过期)的userid
	Unlicenseduser string `json:"unlicenseduser"`
	// MsgID 消息id，用于撤回应用消息
	MsgID string `json:"msgid"`
	// ResponseCode 仅特定模板卡片消息返回，用于更新模版卡片消息
	ResponseCode string `json:"response_code"`
}

// UpdateTemplateCard 更新模版卡片消息
// 应用可以发送模板卡片消息，发送之后可再通过接口更新可回调的用户任务卡片消息的替换文案信息。仅原卡片为按钮交互型、投票选择型、多项选择型的卡片以及填写了action_menu字段的文本通知型、图文展示型可以调用本接口更新。
//
// see https://developer.work.weixin.qq.com/document/path/94963
func (c *Client) UpdateTemplateCard(r *UpdateTemplateCardRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/message/update_template_card",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateTemplateCardRequest is request of Client.UpdateTemplateCard
type UpdateTemplateCardRequest struct {
	// Userids 企业的成员ID列表（最多支持1000个）
	Userids []string `json:"userids"`
	// Partyids 企业的部门ID列表（最多支持100个）
	Partyids []int `json:"partyids"`
	// Tagids 企业的标签ID列表（最多支持100个）
	Tagids []int `json:"tagids"`
	// Atall 更新整个任务接收人员
	Atall int `json:"atall"`
	// AgentID 应用的agentid
	AgentID int `json:"agentid" validate:"required"`
	// ResponseCode 更新卡片所需要消费的code，可通过发消息接口和回调接口返回值获取，一个code只能调用一次该接口，且只能在72小时内调用
	ResponseCode string `json:"response_code" validate:"required"`
	// Button 按钮配置对象（仅用于更新按钮文案或状态）
	Button any `json:"button"`
	// EnableIDTrans 表示是否开启id转译，0表示否，1表示是，默认0
	EnableIDTrans int `json:"enable_id_trans"`
	// TemplateCard 模板卡片配置对象（用于更新为新的卡片）
	TemplateCard any `json:"template_card"`
}

// UpdateGroupChat 修改群聊会话
// 修改通过智能表格自动化创建的群聊的成员信息
//
// see https://developer.work.weixin.qq.com/document/path/101044
func (c *Client) UpdateGroupChat(r *UpdateGroupChatRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedoc/smartsheet/groupchat/update",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateGroupChatRequest is request of Client.UpdateGroupChat
type UpdateGroupChatRequest struct {
	// Docid 智能表格ID。需为当前应用创建的智能表格
	Docid string `json:"docid" validate:"required"`
	// ChatID 群聊ID，需为对应智能表自动规则创建群聊的ID
	ChatID string `json:"chat_id" validate:"required"`
	// Owner 新群主的id。若不需更新，请忽略此参数。del_user_list包含群主时本字段必填
	Owner string `json:"owner"`
	// AddUserList 添加成员的id列表，一次最多传入500人
	AddUserList []string `json:"add_user_list"`
	// DelUserList 踢出成员的id列表，一次最多传入500人
	DelUserList []string `json:"del_user_list"`
}
