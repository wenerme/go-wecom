package wecom

import "github.com/wenerme/go-req"

// CreateUser 创建成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90195
func (c *Client) CreateUser(r *CreateUserRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/user/create",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateUser 更新成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90197
func (c *Client) UpdateUser(r *UpdateUserRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/user/update",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteUser 删除成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90198
func (c *Client) DeleteUser(r *DeleteUserRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/user/delete",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchDeleteUser 批量删除成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90199
func (c *Client) BatchDeleteUser(r *BatchDeleteUserRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/user/batchdelete",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SimpleListUser 获取部门成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90200
func (c *Client) SimpleListUser(r *SimpleListUserRequest, opts ...interface{}) (out SimpleListUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/user/simplelist",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListUser 获取部门成员详情
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90201
func (c *Client) ListUser(r *ListUserRequest, opts ...interface{}) (out ListUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/user/list",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ConvertToOpenID userid与openid互换
// userid转openid
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90202
func (c *Client) ConvertToOpenID(r *ConvertToOpenIDRequest, opts ...interface{}) (out ConvertToOpenIDResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/user/convert_to_openid",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AuthSuccess 二次验证
// 此接口可以满足安全性要求高的企业进行成员验证。开启二次验证后，当且仅当成员登录时，需跳转至企业自定义的页面进行验证。验证频率可在设置页面选择。开启二次验证方法如下图：企业在开启二次验证时，必须在管理端填写企业二次验证页面的url。当成员登录企业微信或关注微工作台（原企业号）进入企业时，会自动跳转到企业的验证页面。在跳转到企业的验证页面时，会带上如下参数：code&#x3D;CODE。企业收到code后，使用“通讯录同步助手”调用接口“根据code获取成员信息”获取成员的userid。如果成员是首次加入企业，企业获取到userid，并验证了成员信息后，调用如下接口即可让成员成功加入企业。
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90203
func (c *Client) AuthSuccess(r *AuthSuccessRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/user/authsucc",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchInvite 邀请成员
// 企业可通过接口批量邀请成员使用企业微信，邀请后将通过短信或邮件下发通知。
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90975
func (c *Client) BatchInvite(r *BatchInviteRequest, opts ...interface{}) (out BatchInviteResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/batch/invite",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetJoinQrcode 获取加入企业二维码
// 支持企业用户获取实时成员加入二维码。
//
// see https://work.weixin.qq.com/api/doc/90000/90135/91714
func (c *Client) GetJoinQrcode(r *GetJoinQrcodeRequest, opts ...interface{}) (out GetJoinQrcodeResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/corp/get_join_qrcode",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetActiveStat 获取企业活跃成员数
//
// see https://work.weixin.qq.com/api/doc/90000/90135/92714
func (c *Client) GetActiveStat(r *GetActiveStatRequest, opts ...interface{}) (out GetActiveStatResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/user/get_active_stat",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

type CreateUserRequest struct {
	// UserID 成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节。只能由数字、字母和“_-@.”四种字符组成，且第一个字符必须是数字或字母。
	UserID string `json:"userid"  validate:"required"`
	// Name 成员名称。长度为1~64个utf8字符
	Name string `json:"name"  validate:"required"`
	// Alias 成员别名。长度1~32个utf8字符
	Alias string `json:"alias"  `
	// Mobile 手机号码。企业内必须唯一，mobile/email二者不能同时为空
	Mobile string `json:"mobile"  `
	// Department 成员所属部门id列表,不超过100个
	Department []int `json:"department"  validate:"required"`
	// Order 部门内的排序值，默认为0，成员次序以创建时间从小到大排列。个数必须和参数department的个数一致，数值越大排序越前面。有效的值范围是[0, 2^32)
	Order []int `json:"order"  `
	// Position 职务信息。长度为0~128个字符
	Position string `json:"position"  `
	// Gender 性别。1表示男性，2表示女性
	Gender string `json:"gender"  `
	// Email 邮箱。长度6~64个字节，且为有效的email格式。企业内必须唯一，mobile/email二者不能同时为空
	Email string `json:"email"  `
	// Telephone 座机。32字节以内，由纯数字、“-”、“+”或“,”组成。
	Telephone string `json:"telephone"  `
	// IsLeaderInDept 个数必须和参数department的个数一致，表示在所在的部门内是否为上级。1表示为上级，0表示非上级。在审批等应用里可以用来标识上级审批人
	IsLeaderInDept []int `json:"is_leader_in_dept"  `
	// AvatarMediaID 成员头像的mediaid，通过素材管理接口上传图片获得的mediaid
	AvatarMediaID string `json:"avatar_mediaid"  `
	// Enable 启用/禁用成员。1表示启用成员，0表示禁用成员
	Enable int `json:"enable"  `
	// ExtAttr 自定义字段。自定义字段需要先在WEB管理端添加，见扩展属性添加方法，否则忽略未知属性的赋值。与对外属性一致，不过只支持type&#x3D;0的文本和type&#x3D;1的网页类型，详细描述查看对外属性
	ExtAttr ExtAttrs `json:"extattr"  `
	// ToInvite 是否邀请该成员使用企业微信（将通过微信服务通知或短信或邮件下发邀请，每天自动下发一次，最多持续3个工作日），默认值为true。
	ToInvite bool `json:"to_invite"  `
	// ExternalProfile 成员对外属性，字段详情见对外属性
	ExternalProfile UserExternalProfile `json:"external_profile"  `
	// ExternalPosition 对外职务，如果设置了该值，则以此作为对外展示的职务，否则以position来展示。长度12个汉字内
	ExternalPosition string `json:"external_position"  `
	// Nickname 视频号名字（设置后，成员将对外展示该视频号）。须从企业绑定到企业微信的视频号中选择，可在“我的企业”页中查看绑定的视频号
	Nickname string `json:"nickname"  `
	// Address 地址。长度最大128个字符
	Address string `json:"address"  `
	// MainDepartment 主部门
	MainDepartment int `json:"main_department"  `
}

type UpdateUserRequest struct {
	// UserID 成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节
	UserID string `json:"userid"  validate:"required"`
	// Name 成员名称。长度为1~64个utf8字符
	Name string `json:"name"  `
	// Alias 别名。长度为1-32个utf8字符
	Alias string `json:"alias"  `
	// Mobile 手机号码。企业内必须唯一。若成员已激活企业微信，则需成员自行修改（此情况下该参数被忽略，但不会报错）
	Mobile string `json:"mobile"  `
	// Department 成员所属部门id列表，不超过100个
	Department []int `json:"department"  `
	// Order 部门内的排序值，默认为0。当有传入department时有效。数量必须和department一致，数值越大排序越前面。有效的值范围是[0, 2^32)
	Order []int `json:"order"  `
	// Position 职务信息。长度为0~128个字符
	Position string `json:"position"  `
	// Gender 性别。1表示男性，2表示女性
	Gender string `json:"gender"  `
	// Email 邮箱。长度不超过64个字节，且为有效的email格式。企业内必须唯一。若是绑定了腾讯企业邮箱的企业微信，则需要在腾讯企业邮箱中修改邮箱（此情况下该参数被忽略，但不会报错）
	Email string `json:"email"  `
	// Telephone 座机。由1-32位的纯数字、“-”、“+”或“,”组成
	Telephone string `json:"telephone"  `
	// IsLeaderInDept 上级字段，个数必须和department一致，表示在所在的部门内是否为上级。
	IsLeaderInDept []int `json:"is_leader_in_dept"  `
	// AvatarMediaID 成员头像的mediaid，通过素材管理接口上传图片获得的mediaid
	AvatarMediaID string `json:"avatar_mediaid"  `
	// Enable 启用/禁用成员。1表示启用成员，0表示禁用成员
	Enable int `json:"enable"  `
	// ExtAttr 自定义字段。自定义字段需要先在WEB管理端添加，见扩展属性添加方法，否则忽略未知属性的赋值。与对外属性一致，不过只支持type&#x3D;0的文本和type&#x3D;1的网页类型，详细描述查看对外属性
	ExtAttr ExtAttrs `json:"extattr"  `
	// ExternalProfile 成员对外属性，字段详情见对外属性
	ExternalProfile UserExternalProfile `json:"external_profile"  `
	// ExternalPosition 对外职务，如果设置了该值，则以此作为对外展示的职务，否则以position来展示。不超过12个汉字
	ExternalPosition string `json:"external_position"  `
	// Nickname 视频号名字（设置后，成员将对外展示该视频号）。须从企业绑定到企业微信的视频号中选择，可在“我的企业”页中查看绑定的视频号
	Nickname string `json:"nickname"  `
	// Address 地址。长度最大128个字符
	Address string `json:"address"  `
	// MainDepartment 主部门
	MainDepartment int `json:"main_department"  `
}

type DeleteUserRequest struct {
	// UserID 成员UserID。对应管理端的帐号
	UserID string `json:"userid"  validate:"required"`
}

type BatchDeleteUserRequest struct {
	// UserIDList 成员UserID列表。对应管理端的帐号。最多支持200个。若存在无效UserID，直接返回错误
	UserIDList []string `json:"useridlist"  validate:"required"`
}

type SimpleListUserRequest struct {
	// DepartmentID 获取的部门id
	DepartmentID string `json:"department_id"  validate:"required"`
	// FetchChild 是否递归获取子部门下面的成员：1-递归获取，0-只获取本部门
	FetchChild string `json:"fetch_child"  `
}

type SimpleListUserResponse struct {
	// UserList 成员列表
	UserList []SimpleListUser `json:"userlist"  `
}

type ListUserRequest struct {
	// DepartmentID 获取的部门id
	DepartmentID string `json:"department_id"  validate:"required"`
	// FetchChild 1/0：是否递归获取子部门下面的成员
	FetchChild string `json:"fetch_child"  `
}

type ListUserResponse struct {
	// UserList 成员列表
	UserList []ListUser `json:"userlist"  `
}

type ConvertToOpenIDRequest struct {
	// UserID 企业内的成员id
	UserID string `json:"userid"  validate:"required"`
}

type ConvertToOpenIDResponse struct {
	// OpenID 企业微信成员userid对应的openid
	OpenID string `json:"openid"  `
}

type AuthSuccessRequest struct {
	// UserID 成员UserID。对应管理端的帐号
	UserID string `json:"userid"  validate:"required"`
}

type BatchInviteRequest struct {
	// User 成员ID列表, 最多支持1000个。
	User string `json:"user"  `
	// Party 部门ID列表，最多支持100个。
	Party string `json:"party"  `
	// Tag 标签ID列表，最多支持100个。
	Tag string `json:"tag"  `
}

type BatchInviteResponse struct {
	// InvalidUser 非法成员列表
	InvalidUser string `json:"invaliduser"  `
	// InvalidParty 非法部门列表
	InvalidParty string `json:"invalidparty"  `
	// InvalidTag 非法标签列表
	InvalidTag string `json:"invalidtag"  `
}

type GetJoinQrcodeRequest struct {
	// SizeType qrcode尺寸类型，1: 171 x 171; 2: 399 x 399; 3: 741 x 741; 4: 2052 x 2052
	SizeType string `json:"size_type"  `
}

type GetJoinQrcodeResponse struct {
	// JoinQrcode 二维码链接，有效期7天
	JoinQrcode string `json:"join_qrcode"  `
}

type GetActiveStatRequest struct {
	// Date 具体某天的活跃人数，最长支持获取30天前数据
	Date string `json:"date"  validate:"required"`
}

type GetActiveStatResponse struct {
	// ActiveCount 活跃成员数
	ActiveCount int `json:"active_cnt"  `
}
