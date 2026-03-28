package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// ActEmailAccount 禁用/启用邮箱账号
// 禁用或启用成员邮箱和业务邮箱
//
// see https://developer.work.weixin.qq.com/document/path/97683
func (c *Client) ActEmailAccount(r *ActEmailAccountRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/account/act_email",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ActEmailAccountRequest is request of Client.ActEmailAccount
type ActEmailAccountRequest struct {
	// UserID 成员UserID，与publicemail_id至少传一项
	UserID string `json:"userid" validate:"required"`
	// PublicemailID 业务邮箱ID，与userid至少传一项
	PublicemailID int `json:"publicemail_id" validate:"required"`
	// Type 1启用，2禁用
	Type int `json:"type" validate:"required"`
}

// SendExmail 发送普通邮件
// 应用可以通过该接口发送普通邮件，支持附件能力。
//
// see https://developer.work.weixin.qq.com/document/path/97445
func (c *Client) SendExmail(r *SendExmailRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/app/compose_send",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SendExmailRequest is request of Client.SendExmail
type SendExmailRequest struct {
	// To 收件人，to.emails 和 to.userids 至少传一个
	To any `json:"to" validate:"required"`
	// Cc 抄送
	Cc any `json:"cc"`
	// Bcc 密送
	Bcc any `json:"bcc"`
	// Subject 标题
	Subject string `json:"subject" validate:"required"`
	// Content 内容
	Content string `json:"content" validate:"required"`
	// AttachmentList 附件相关
	AttachmentList []map[string]any `json:"attachment_list"`
	// ContentType 内容类型 html，text（默认是html）
	ContentType string `json:"content_type"`
	// EnableIDTrans 表示是否开启id转译，0表示否，1表示是
	EnableIDTrans int `json:"enable_id_trans"`
}

// GetEmailAlias 查询应用邮箱账号
// 应用调用此接口，可以查询自己的应用邮箱账号及别名邮箱
//
// see https://developer.work.weixin.qq.com/document/path/97991
func (c *Client) GetEmailAlias(r *GetEmailAliasRequest, opts ...any) (out GetEmailAliasResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/app/get_email_alias",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetEmailAliasRequest is request of Client.GetEmailAlias
type GetEmailAliasRequest struct{}

// GetEmailAliasResponse is response of Client.GetEmailAlias
type GetEmailAliasResponse struct {
	// Email 当前发信账号的主邮箱地址，发邮件将会以此邮箱地址为发件人
	Email string `json:"email"`
	// AliasList 别名邮箱地址列表，能作为收件人收邮件
	AliasList []string `json:"alias_list"`
}

// ListMail 获取收件箱邮件列表
// 分页获取收件箱下，指定时间范围内的邮件id列表。
//
// see https://developer.work.weixin.qq.com/document/path/97681
func (c *Client) ListMail(r *ListMailRequest, opts ...any) (out ListMailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/app/get_mail_list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListMailRequest is request of Client.ListMail
type ListMailRequest struct {
	// BeginTime 开始时间，unix时间戳
	BeginTime int `json:"begin_time" validate:"required"`
	// EndTime 结束时间，unix时间戳
	EndTime int `json:"end_time" validate:"required"`
	// Cursor 上一次调用时返回的next_cursor，第一次拉取可以不填
	Cursor string `json:"cursor"`
	// Limit 期望请求的数据量，默认值为100，最大值为1000
	Limit int `json:"limit"`
}

// ListMailResponse is response of Client.ListMail
type ListMailResponse struct {
	// NextCursor 应用邮箱账号中邮件未读数（用于分页）
	NextCursor string `json:"next_cursor"`
	// HasMore 是否还有更多数据。0-没有 1-有
	HasMore int `json:"has_more"`
	// MailList 邮件列表
	MailList []map[string]any `json:"mail_list"`
}

// ReadMail 获取邮件内容
// 支持应用获取邮件内容。指定单个邮件id，获取邮件eml数据。
//
// see https://developer.work.weixin.qq.com/document/path/97979
func (c *Client) ReadMail(r *ReadMailRequest, opts ...any) (out ReadMailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/app/read_mail",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ReadMailRequest is request of Client.ReadMail
type ReadMailRequest struct {
	// MailID 邮件id
	MailID string `json:"mail_id" validate:"required"`
}

// ReadMailResponse is response of Client.ReadMail
type ReadMailResponse struct {
	// MailData 邮件eml内容
	MailData string `json:"mail_data"`
}

// UpdateAppEmailAlias 更新应用邮箱账号
// 更新应用邮箱账号，原有的应用邮箱账号将会作为别名邮箱，具有收信能力
//
// see https://developer.work.weixin.qq.com/document/path/97682
func (c *Client) UpdateAppEmailAlias(r *UpdateAppEmailAliasRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/app/update_email_alias",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateAppEmailAliasRequest is request of Client.UpdateAppEmailAlias
type UpdateAppEmailAliasRequest struct {
	// NewEmail 修改后的应用邮箱账号
	NewEmail string `json:"new_email" validate:"required"`
}

// CreateMailGroup 创建邮件群组
// 该接口用于创建新邮件群组，可以指定群组成员，定义群组使用权限范围。
//
// see https://developer.work.weixin.qq.com/document/path/95510
func (c *Client) CreateMailGroup(r *CreateMailGroupRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/group/create",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateMailGroupRequest is request of Client.CreateMailGroup
type CreateMailGroupRequest struct {
	// GroupID 邮件群组ID，邮箱格式
	GroupID string `json:"groupid" validate:"required"`
	// Groupname 邮件群组名称，不能与其他群组重名，长度限定200字节
	Groupname string `json:"groupname" validate:"required"`
	// EmailList 群组内成员邮箱地址列表
	EmailList []string `json:"email_list"`
	// TagList 群组内包含的标签ID列表
	TagList []int `json:"tag_list"`
	// DepartmentList 群组内包含的部门ID列表
	DepartmentList []int `json:"department_list"`
	// GroupList 群组内包含的群组邮箱列表
	GroupList []string `json:"group_list"`
	// AllowType 群组使用权限。0: 企业成员，1任何人，2:组内成员，3:自定义成员
	AllowType int `json:"allow_type"`
	// AllowEmaillist 允许使用群组群发的成员邮箱地址列表
	AllowEmaillist []string `json:"allow_emaillist"`
	// AllowDepartmentlist 允许使用群组群发的部门ID列表
	AllowDepartmentlist []int `json:"allow_departmentlist"`
	// AllowTaglist 允许使用群组群发的标签ID列表
	AllowTaglist []int `json:"allow_taglist"`
}

// DeleteMailGroup 删除邮件群组
// 该接口用于删除已有的邮件群组。
//
// see https://developer.work.weixin.qq.com/document/path/97996
func (c *Client) DeleteMailGroup(r *DeleteMailGroupRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/group/delete",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteMailGroupRequest is request of Client.DeleteMailGroup
type DeleteMailGroupRequest struct {
	// GroupID 邮件群组ID，邮箱格式
	GroupID string `json:"groupid" validate:"required"`
}

// GetMailGroup 获取邮件群组详情
// 该接口用于获取邮件群组详细信息，包含群组名称、群组成员、群组使用权限等。
//
// see https://developer.work.weixin.qq.com/document/path/97997
func (c *Client) GetMailGroup(r *GetMailGroupRequest, opts ...any) (out GetMailGroupResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/exmail/group/get",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMailGroupRequest is request of Client.GetMailGroup
type GetMailGroupRequest struct {
	// GroupID 邮件群组ID，邮箱格式
	GroupID string `json:"groupid" validate:"required"`
}

// GetMailGroupResponse is response of Client.GetMailGroup
type GetMailGroupResponse struct {
	// GroupID 邮件群组ID，邮箱格式
	GroupID string `json:"groupid"`
	// Groupname 邮件群组名称
	Groupname string `json:"groupname"`
	// EmailList 群组内成员邮箱地址列表
	EmailList any `json:"email_list"`
	// TagList 群组内包含的标签ID列表
	TagList any `json:"tag_list"`
	// DepartmentList 群组内包含的部门ID列表
	DepartmentList any `json:"department_list"`
	// GroupList 群组内包含的群组邮箱ID列表
	GroupList any `json:"group_list"`
	// AllowType 群组使用权限。0: 企业成员, 1任何人，2:组内成员，3:自定义成员
	AllowType int `json:"allow_type"`
	// AllowEmaillist 允许使用群组群发的成员邮箱地址列表
	AllowEmaillist any `json:"allow_emaillist"`
	// AllowDepartmentlist 允许使用群组群发的部门ID列表
	AllowDepartmentlist any `json:"allow_departmentlist"`
	// AllowTaglist 允许使用群组群发的标签ID列表
	AllowTaglist any `json:"allow_taglist"`
}

// SearchMailGroup 模糊搜索邮件群组
// 该接口用于通过群组ID模糊搜索邮件群组。
//
// see https://developer.work.weixin.qq.com/document/path/97998
func (c *Client) SearchMailGroup(r *SearchMailGroupRequest, opts ...any) (out SearchMailGroupResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/exmail/group/search",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SearchMailGroupRequest is request of Client.SearchMailGroup
type SearchMailGroupRequest struct {
	// Fuzzy 1开启模糊搜索，0获取全部邮件群组
	Fuzzy int `json:"fuzzy" validate:"required"`
	// GroupID 邮件群组ID，邮箱格式
	GroupID string `json:"groupid"`
}

// SearchMailGroupResponse is response of Client.SearchMailGroup
type SearchMailGroupResponse struct {
	// Count 返回条数
	Count int `json:"count"`
	// Groups 群组列表
	Groups []map[string]any `json:"groups"`
}

// UpdateMailGroup 更新邮件群组
// 该接口用于更新邮件群组，可以修改群组名称、群组成员、群组使用权限等。需要注意的是Json数组类型传空值将会清空其内容，不传则保持不变。
//
// see https://developer.work.weixin.qq.com/document/path/97995
func (c *Client) UpdateMailGroup(r *UpdateMailGroupRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/group/update",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateMailGroupRequest is request of Client.UpdateMailGroup
type UpdateMailGroupRequest struct {
	// GroupID 邮件群组ID，邮箱格式
	GroupID string `json:"groupid" validate:"required"`
	// Groupname 邮件群组名称，不能与其他群组重名，长度限定200字节
	Groupname string `json:"groupname"`
	// EmailList 群组内成员邮箱地址，读取成员的biz_mail字段，不传则不变，传空则清空。成员由email_list，group_list，department_list，tag_list共同组成，不允许全部清空
	EmailList any `json:"email_list"`
	// TagList 群组内包含的标签ID，不传则不变，传空为清空
	TagList any `json:"tag_list"`
	// DepartmentList 群组内包含的部门ID，不传则不变，传空为清空
	DepartmentList any `json:"department_list"`
	// GroupList 群组内包含的群组邮箱ID，不传则不变，传空为清空
	GroupList any `json:"group_list"`
	// AllowType 群组使用权限。0: 企业成员, 1任何人，2:组内成员，3:自定义成员。当值为0、1、2时，不得传入allow_emaillist，allow_departmentlist，allow_tagl...
	AllowType int `json:"allow_type"`
	// AllowEmaillist 允许使用群组群发的成员邮箱地址，不传则不变，传空为清空
	AllowEmaillist any `json:"allow_emaillist"`
	// AllowDepartmentlist 允许使用群组群发的部门ID，不传则不变，传空为清空
	AllowDepartmentlist any `json:"allow_departmentlist"`
	// AllowTaglist 允许使用群组群发的标签ID，不传则不变，传空为清空
	AllowTaglist any `json:"allow_taglist"`
}

// GetMailUnreadCount 获取邮件未读数
// 获取指定成员邮箱当前未读邮件数量
//
// see https://developer.work.weixin.qq.com/document/path/97685
func (c *Client) GetMailUnreadCount(r *GetMailUnreadCountRequest, opts ...any) (out GetMailUnreadCountResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/mail/get_newcount",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMailUnreadCountRequest is request of Client.GetMailUnreadCount
type GetMailUnreadCountRequest struct {
	// UserID 成员UserID
	UserID string `json:"userid" validate:"required"`
}

// GetMailUnreadCountResponse is response of Client.GetMailUnreadCount
type GetMailUnreadCountResponse struct {
	// Count 成员邮箱中邮件未读数
	Count int `json:"count"`
}

// CreatePublicMail 创建公共邮箱
// 该接口用于创建公共邮箱，指定公共邮箱使用权限。
//
// see https://developer.work.weixin.qq.com/document/path/100185
func (c *Client) CreatePublicMail(r *CreatePublicMailRequest, opts ...any) (out CreatePublicMailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/publicmail/create",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreatePublicMailRequest is request of Client.CreatePublicMail
type CreatePublicMailRequest struct {
	// Email 公共邮箱地址
	Email string `json:"email" validate:"required"`
	// Name 公共邮箱名称，不多于64个字符或32个汉字，不得与其他公共邮箱重名
	Name string `json:"name" validate:"required"`
	// UserIDList 有权限使用公共邮箱的成员UserID列表。userid_list、department_list、taglist不能同时为空
	UserIDList any `json:"userid_list"`
	// DepartmentList 有权限使用公共邮箱的部门ID列表
	DepartmentList any `json:"department_list"`
	// TagList 有权限使用公共邮箱的标签ID列表
	TagList any `json:"tag_list"`
	// CreateAuthCode 是否创建客户端专用密码，0表示否，1表示是，默认0。客户端专用密码只会显示一次，请在客户端中输入该密码进行验证，每次生成的密码皆可使用，请勿告诉他人
	CreateAuthCode int `json:"create_auth_code"`
	// AuthCodeInfo 创建客户端专用密码的备注信息
	AuthCodeInfo any `json:"auth_code_info"`
}

// CreatePublicMailResponse is response of Client.CreatePublicMail
type CreatePublicMailResponse struct {
	// ID 公共邮箱ID
	ID int `json:"id"`
	// AuthCodeID 客户端专用密码ID，仅当设置创建客户端专用密码的时候返回，后续可通过接口删除客户端专用密码
	AuthCodeID int `json:"auth_code_id"`
	// AuthCode 客户端专用密码, 仅当设置创建客户端专用密码的时候返回，该客户端专用密码仅会返回一次，请妥善存储
	AuthCode string `json:"auth_code"`
}

// DeletePublicMail 删除公共邮箱
// 该接口用于删除已有的公共邮箱。
//
// see https://developer.work.weixin.qq.com/document/path/100187
func (c *Client) DeletePublicMail(r *DeletePublicMailRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/publicmail/delete",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeletePublicMailRequest is request of Client.DeletePublicMail
type DeletePublicMailRequest struct {
	// ID 公共邮箱ID
	ID int `json:"id" validate:"required"`
}

// DeletePublicMailAuthCode 删除客户端专用密码
// 该接口用于删除公共邮箱的客户端专用密码
//
// see https://developer.work.weixin.qq.com/document/path/100242
func (c *Client) DeletePublicMailAuthCode(r *DeletePublicMailAuthCodeRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/publicmail/delete_auth_code",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeletePublicMailAuthCodeRequest is request of Client.DeletePublicMailAuthCode
type DeletePublicMailAuthCodeRequest struct {
	// ID 公共邮箱ID
	ID int `json:"id" validate:"required"`
	// AuthCodeID 客户端专用密码ID
	AuthCodeID int `json:"auth_code_id" validate:"required"`
}

// GetPublicMailDetail 获取公共邮箱详情
// 该接口用于获取公共邮箱详细信息，包含公共邮箱名称、权限信息。
//
// see https://developer.work.weixin.qq.com/document/path/100188
func (c *Client) GetPublicMailDetail(r *GetPublicMailDetailRequest, opts ...any) (out GetPublicMailDetailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/publicmail/get",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetPublicMailDetailRequest is request of Client.GetPublicMailDetail
type GetPublicMailDetailRequest struct {
	// IDList 公共邮箱ID列表
	IDList []int `json:"id_list" validate:"required"`
}

// GetPublicMailDetailResponse is response of Client.GetPublicMailDetail
type GetPublicMailDetailResponse struct {
	// List 公共邮箱详情列表
	List []map[string]any `json:"list"`
}

// ListPublicMailAuthCode 获取客户端专用密码列表
// 获取公共邮箱创建的客户端专用密码列表，包括客户端专用密码ID、创建时间、最后使用时间、备注。该接口不返回客户端专用密码本身。
//
// see https://developer.work.weixin.qq.com/document/path/100241
func (c *Client) ListPublicMailAuthCode(r *ListPublicMailAuthCodeRequest, opts ...any) (out ListPublicMailAuthCodeResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/publicmail/get_auth_code_list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListPublicMailAuthCodeRequest is request of Client.ListPublicMailAuthCode
type ListPublicMailAuthCodeRequest struct {
	// ID 公共邮箱ID
	ID int `json:"id" validate:"required"`
}

// ListPublicMailAuthCodeResponse is response of Client.ListPublicMailAuthCode
type ListPublicMailAuthCodeResponse struct {
	// AuthCodeList 客户端专用密码列表
	AuthCodeList []map[string]any `json:"auth_code_list"`
}

// SearchPublicMail 模糊搜索公共邮箱
// 该接口用于模糊搜索公共邮箱，可匹配邮箱名也可匹配邮箱地址。
//
// see https://developer.work.weixin.qq.com/document/path/100189
func (c *Client) SearchPublicMail(r *SearchPublicMailRequest, opts ...any) (out SearchPublicMailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/exmail/publicmail/search",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SearchPublicMailRequest is request of Client.SearchPublicMail
type SearchPublicMailRequest struct {
	// Fuzzy 1开启模糊搜索，0获取全部公共邮箱
	Fuzzy int `json:"fuzzy" validate:"required"`
	// Email 公共邮箱名称或邮箱地址
	Email string `json:"email"`
}

// SearchPublicMailResponse is response of Client.SearchPublicMail
type SearchPublicMailResponse struct {
	// List 公共邮箱列表
	List []map[string]any `json:"list"`
}

// UpdatePublicMail 更新公共邮箱
// 更新公共邮箱的名称、使用权限及客户端专用密码配置
//
// see https://developer.work.weixin.qq.com/document/path/100186
func (c *Client) UpdatePublicMail(r *UpdatePublicMailRequest, opts ...any) (out UpdatePublicMailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/publicmail/update",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdatePublicMailRequest is request of Client.UpdatePublicMail
type UpdatePublicMailRequest struct {
	// ID 公共邮箱ID
	ID int `json:"id" validate:"required"`
	// Name 公共邮箱名称，不多于64个字符或32个汉字
	Name string `json:"name"`
	// UserIDList 有权限使用公共邮箱的成员UserID列表对象，包含list字段。不传则不变，传空为清空
	UserIDList any `json:"userid_list"`
	// DepartmentList 有权限使用公共邮箱的部门列表对象，包含list字段。不传则不变，传空为清空
	DepartmentList any `json:"department_list"`
	// TagList 有权限使用公共邮箱的标签列表对象，包含list字段。不传则不变，传空为清空
	TagList any `json:"tag_list"`
	// AliasList 邮箱别名列表对象，包含list字段。更新时为覆盖式更新，传空结构或数组会清空当前别名
	AliasList any `json:"alias_list"`
	// CreateAuthCode 是否创建客户端专用密码，0表示否，1表示是，默认0
	CreateAuthCode int `json:"create_auth_code"`
	// AuthCodeInfo 创建客户端专用密码的备注信息对象
	AuthCodeInfo any `json:"auth_code_info"`
}

// UpdatePublicMailResponse is response of Client.UpdatePublicMail
type UpdatePublicMailResponse struct {
	// AuthCodeID 客户端专用密码ID，仅当设置创建客户端专用密码时返回
	AuthCodeID int `json:"auth_code_id"`
	// AuthCode 客户端专用密码，仅当设置创建客户端专用密码时返回
	AuthCode string `json:"auth_code"`
}

// GetUserOption 获取用户功能属性
// 该接口用于获取用户的功能属性
//
// see https://developer.work.weixin.qq.com/document/path/97684
func (c *Client) GetUserOption(r *GetUserOptionRequest, opts ...any) (out GetUserOptionResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/useroption/get",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetUserOptionRequest is request of Client.GetUserOption
type GetUserOptionRequest struct {
	// UserID 用户UserID
	UserID string `json:"userid" validate:"required"`
	// Type 功能设置属性类型 1: 强制启用安全登录 2: IMAP/SMTP服务 3: POP/SMTP服务 4: 是否启用安全登录
	Type []int `json:"type" validate:"required"`
}

// GetUserOptionResponse is response of Client.GetUserOption
type GetUserOptionResponse struct {
	// Option 功能属性对象
	Option any `json:"option"`
}

// UpdateUserOption 更改用户功能属性
// 该接口用于更新用户的功能属性
//
// see https://developer.work.weixin.qq.com/document/path/98008
func (c *Client) UpdateUserOption(r *UpdateUserOptionRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/useroption/update",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateUserOptionRequest is request of Client.UpdateUserOption
type UpdateUserOptionRequest struct {
	// UserID 用户UserID
	UserID string `json:"userid" validate:"required"`
	// Option 功能设置属性列表
	Option any `json:"option" validate:"required"`
}

// BatchAddVip 分配高级功能账号
// 该接口可以为在应用可见范围的企业成员分配高级功能。
//
// see https://developer.work.weixin.qq.com/document/path/99320
func (c *Client) BatchAddVip(r *BatchAddVipRequest, opts ...any) (out BatchAddVipResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/vip/batch_add",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchAddVipRequest is request of Client.BatchAddVip
type BatchAddVipRequest struct {
	// UserIDList 要分配高级功能的企业成员userid列表，单次操作最大限制100个
	UserIDList []string `json:"userid_list" validate:"required"`
}

// BatchAddVipResponse is response of Client.BatchAddVip
type BatchAddVipResponse struct {
	// SuccUserIDList 分配成功的用户列表，包括之前已经分配过的用户
	SuccUserIDList []string `json:"succ_userid_list"`
	// FailUserIDList 分配失败的用户列表
	FailUserIDList []string `json:"fail_userid_list"`
}

// BatchDelVip 取消高级功能账号
// 该接口用于撤销分配应用可见范围的企业成员的高级功能。
//
// see https://developer.work.weixin.qq.com/document/path/99321
func (c *Client) BatchDelVip(r *BatchDelVipRequest, opts ...any) (out BatchDelVipResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/vip/batch_del",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchDelVipRequest is request of Client.BatchDelVip
type BatchDelVipRequest struct {
	// UserIDList 要撤销分配高级功能的企业成员userid列表，单次操作最多限制100个
	UserIDList []string `json:"userid_list" validate:"required"`
}

// BatchDelVipResponse is response of Client.BatchDelVip
type BatchDelVipResponse struct {
	// SuccUserIDList 撤销分配成功的用户列表
	SuccUserIDList []string `json:"succ_userid_list"`
	// FailUserIDList 撤销分配失败的用户列表
	FailUserIDList []string `json:"fail_userid_list"`
}

// ListExmailVip 获取高级功能账号列表
// 查询企业已分配高级功能且在应用可见范围的账号列表
//
// see https://developer.work.weixin.qq.com/document/path/99322
func (c *Client) ListExmailVip(r *ListExmailVipRequest, opts ...any) (out ListExmailVipResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/exmail/vip/list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListExmailVipRequest is request of Client.ListExmailVip
type ListExmailVipRequest struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"`
	// Limit 用于分页查询，每次请求返回的数据上限。默认100，最大200
	Limit int `json:"limit"`
}

// ListExmailVipResponse is response of Client.ListExmailVip
type ListExmailVipResponse struct {
	// HasMore 是否还有更多数据未获取
	HasMore bool `json:"has_more"`
	// NextCursor 下一次请求的cursor值
	NextCursor string `json:"next_cursor"`
	// UserIDList 符合条件的企业成员userid列表
	UserIDList []string `json:"userid_list"`
}
