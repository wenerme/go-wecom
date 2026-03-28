package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// ConvertToOpenid 外部联系人openid转换
// 将微信外部联系人的userid转为微信openid，用于调用支付相关接口。暂不支持企业微信外部联系人的userid转openid。
//
// see https://developer.work.weixin.qq.com/document/path/92323
func (c *Client) ConvertToOpenid(r *ConvertToOpenidRequest, opts ...any) (out ConvertToOpenidResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/convert_to_openid",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ConvertToOpenidRequest is request of Client.ConvertToOpenid
type ConvertToOpenidRequest struct {
	// ExternalUserID 外部联系人的userid，注意不是企业成员的账号
	ExternalUserID string `json:"external_userid" validate:"required"`
}

// ConvertToOpenidResponse is response of Client.ConvertToOpenid
type ConvertToOpenidResponse struct {
	// OpenID 该企业的外部联系人openid
	OpenID string `json:"openid"`
}

// GetSubscribeQrCode 获取「学校通知」二维码
// 学校可通过此接口获取「学校通知」二维码，家长可通过扫描此二维码关注「学校通知」并接收学校推送的消息。
//
// see https://developer.work.weixin.qq.com/document/path/92320
func (c *Client) GetSubscribeQrCode(r *GetSubscribeQrCodeRequest, opts ...any) (out GetSubscribeQrCodeResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/externalcontact/get_subscribe_qr_code",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetSubscribeQrCodeRequest is request of Client.GetSubscribeQrCode
type GetSubscribeQrCodeRequest struct{}

// GetSubscribeQrCodeResponse is response of Client.GetSubscribeQrCode
type GetSubscribeQrCodeResponse struct {
	// QrcodeBig 1200px的大尺寸二维码
	QrcodeBig string `json:"qrcode_big"`
	// QrcodeMiddle 430px的中尺寸二维码
	QrcodeMiddle string `json:"qrcode_middle"`
	// QrcodeThumb 258px的小尺寸二维码
	QrcodeThumb string `json:"qrcode_thumb"`
}

// SetExternalContactSubscribeMode 设置「学校通知」的关注模式
// 修改家长关注「学校通知」的模式：可扫码填写资料加入或禁止扫码填写资料加入
//
// see https://developer.work.weixin.qq.com/document/path/92318
func (c *Client) SetExternalContactSubscribeMode(r *SetExternalContactSubscribeModeRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/set_subscribe_mode",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetExternalContactSubscribeModeRequest is request of Client.SetExternalContactSubscribeMode
type SetExternalContactSubscribeModeRequest struct {
	// SubscribeMode 关注模式，1:可扫码填写资料加入，2:禁止扫码填写资料加入
	SubscribeMode int `json:"subscribe_mode" validate:"required"`
}

// GetSchoolAgentAllowScope 获取可使用的家长范围
// 获取可在微信「学校通知-学校应用」使用该应用的家长范围，以学生或部门列表的形式返回。
//
// see https://developer.work.weixin.qq.com/document/path/94895
func (c *Client) GetSchoolAgentAllowScope(r *GetSchoolAgentAllowScopeRequest, opts ...any) (out GetSchoolAgentAllowScopeResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/school/agent/get_allow_scope",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetSchoolAgentAllowScopeRequest is request of Client.GetSchoolAgentAllowScope
type GetSchoolAgentAllowScopeRequest struct {
	// AgentID 应用id
	AgentID int `json:"agentid" validate:"required"`
}

// GetSchoolAgentAllowScopeResponse is response of Client.GetSchoolAgentAllowScope
type GetSchoolAgentAllowScopeResponse struct {
	// AllowScope 可在微信「学校通知-学校应用」使用该应用的家长范围
	AllowScope any `json:"allow_scope"`
}

// ListSchoolDepartment 获取部门列表
// 获取企业下的部门列表，支持指定获取某部门及其子部门
//
// see https://developer.work.weixin.qq.com/document/path/92343
func (c *Client) ListSchoolDepartment(r *ListSchoolDepartmentRequest, opts ...any) (out ListSchoolDepartmentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/school/department/list",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListSchoolDepartmentRequest is request of Client.ListSchoolDepartment
type ListSchoolDepartmentRequest struct {
	// ID 部门id。获取指定部门及其下的子部门。如果不填，默认获取全量组织架构
	ID int `json:"id"`
}

// ListSchoolDepartmentResponse is response of Client.ListSchoolDepartment
type ListSchoolDepartmentResponse struct {
	// Departments 部门列表数据
	Departments []map[string]any `json:"departments"`
}

// UpdateSchoolDepartment 更新部门
// 更新家校通讯录中的部门信息，支持修改部门名称、ID、年级属性及管理员配置。
//
// see https://developer.work.weixin.qq.com/document/path/92341
func (c *Client) UpdateSchoolDepartment(r *UpdateSchoolDepartmentRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/department/update",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateSchoolDepartmentRequest is request of Client.UpdateSchoolDepartment
type UpdateSchoolDepartmentRequest struct {
	// Name 部门名称，长度限制1~32字符
	Name string `json:"name"`
	// Parentid 父部门id，32位整型
	Parentid int `json:"parentid"`
	// ID 部门id，32位整型，必须大于0
	ID int `json:"id" validate:"required"`
	// NewID 修改为新的id
	NewID int `json:"new_id"`
	// RegisterYear 入学年份，YYYY格式，仅当部门类型为年级时生效
	RegisterYear int `json:"register_year"`
	// StandardGrade 标准年级，0表示转换为非标准年级
	StandardGrade int `json:"standard_grade"`
	// Order 在父部门中的次序值
	Order int `json:"order"`
	// DepartmentAdmins 部门管理员列表
	DepartmentAdmins []map[string]any `json:"department_admins"`
}

// GetChatCreateMode 管理「班级群创建方式」
// 企业或第三方可通过接口获取或修改家校通讯录配置「班级群创建方式」，目前支持“自动创建”或“手动创建”两种方式。
//
// see https://developer.work.weixin.qq.com/document/path/92430
func (c *Client) GetChatCreateMode(r *GetChatCreateModeRequest, opts ...any) (out GetChatCreateModeResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/school/get_chat_create_mode",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetChatCreateModeRequest is request of Client.GetChatCreateMode
type GetChatCreateModeRequest struct{}

// GetChatCreateModeResponse is response of Client.GetChatCreateMode
type GetChatCreateModeResponse struct {
	// CreateMode 创建模式。取值范围：0 - 自动创建，1 - 手动创建
	CreateMode int `json:"create_mode"`
}

// GetSchoolUserInfo 获取家校访问用户身份
// 该接口用于根据code获取家长或者学生信息
//
// see https://developer.work.weixin.qq.com/document/path/95791
func (c *Client) GetSchoolUserInfo(r *GetSchoolUserInfoRequest, opts ...any) (out GetSchoolUserInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/school/getuserinfo",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetSchoolUserInfoRequest is request of Client.GetSchoolUserInfo
type GetSchoolUserInfoRequest struct {
	// Code 通过成员授权获取到的code，最大为512字节。每次成员授权带上的code将不一样，code只能使用一次，5分钟未被使用自动过期。
	Code string `json:"code" validate:"required"`
}

// GetSchoolUserInfoResponse is response of Client.GetSchoolUserInfo
type GetSchoolUserInfoResponse struct {
	// DeviceId 手机设备号(由企业微信在安装时随机生成，删除重装会改变，升级不受影响)
	DeviceId string `json:"DeviceId"`
	// ParentUserID 家校通讯录里家长的userid
	ParentUserID string `json:"parent_userid"`
	// StudentUserID 家校通讯录里学生的userid
	StudentUserID string `json:"student_userid"`
}

// SetArchSyncMode 设置家校通讯录自动同步模式
// 修改家校通讯录与班级标签之间的自动同步模式
//
// see https://developer.work.weixin.qq.com/document/path/92345
func (c *Client) SetArchSyncMode(r *SetArchSyncModeRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/set_arch_sync_mode",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetArchSyncModeRequest is request of Client.SetArchSyncMode
type SetArchSyncModeRequest struct {
	// ArchSyncMode 家校通讯录同步模式：1-禁止将标签同步至家校通讯录，2-禁止将家校通讯录同步至标签，3-禁止家校通讯录和标签相互同步
	ArchSyncMode int `json:"arch_sync_mode" validate:"required"`
}

// SetSchoolUpgradeInfo 修改自动升年级的配置
// 设置或修改学校自动升年级的时间与开关状态
//
// see https://developer.work.weixin.qq.com/document/path/92949
func (c *Client) SetSchoolUpgradeInfo(r *SetSchoolUpgradeInfoRequest, opts ...any) (out SetSchoolUpgradeInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/set_upgrade_info",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetSchoolUpgradeInfoRequest is request of Client.SetSchoolUpgradeInfo
type SetSchoolUpgradeInfoRequest struct {
	// UpgradeTime 自动升年级的时间，该时间戳只有月和日有效，不传则默认为传0，代表的1月1号
	UpgradeTime int `json:"upgrade_time"`
	// UpgradeSwitch 开启或关闭自动升年级。0：表示关闭，1：表示开启，不传默认关闭，传所有非1的值也视为关闭
	UpgradeSwitch int `json:"upgrade_switch"`
}

// SetSchoolUpgradeInfoResponse is response of Client.SetSchoolUpgradeInfo
type SetSchoolUpgradeInfoResponse struct {
	// NextUpgradeTime 下次升级的时间，只有年月日有效，如果该学校今年已经升过年级，那本参数会给明年对应日期的时间戳
	NextUpgradeTime int `json:"next_upgrade_time"`
}

// BatchCreateParent 批量创建家长
// 批量创建企业学校内的家长信息
//
// see https://developer.work.weixin.qq.com/document/path/92334
func (c *Client) BatchCreateParent(r *BatchCreateParentRequest, opts ...any) (out BatchCreateParentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/user/batch_create_parent",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchCreateParentRequest is request of Client.BatchCreateParent
type BatchCreateParentRequest struct {
	// Parents 家长列表，每次最多100个
	Parents []map[string]any `json:"parents" validate:"required"`
}

// BatchCreateParentResponse is response of Client.BatchCreateParent
type BatchCreateParentResponse struct {
	// ResultList 失败列表
	ResultList []map[string]any `json:"result_list"`
}

// BatchCreateStudent 批量创建学生
// 批量创建学校成员（学生）
//
// see https://developer.work.weixin.qq.com/document/path/92328
func (c *Client) BatchCreateStudent(r *BatchCreateStudentRequest, opts ...any) (out BatchCreateStudentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/user/batch_create_student",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchCreateStudentRequest is request of Client.BatchCreateStudent
type BatchCreateStudentRequest struct {
	// Students 学生列表，每次最多100个学生
	Students []map[string]any `json:"students" validate:"required"`
}

// BatchCreateStudentResponse is response of Client.BatchCreateStudent
type BatchCreateStudentResponse struct {
	// ResultList 失败的学生列表
	ResultList []map[string]any `json:"result_list"`
}

// BatchDeleteParent 批量删除家长
// 批量删除家校通讯录中的家长成员
//
// see https://developer.work.weixin.qq.com/document/path/92335
func (c *Client) BatchDeleteParent(r *BatchDeleteParentRequest, opts ...any) (out BatchDeleteParentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/user/batch_delete_parent",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchDeleteParentRequest is request of Client.BatchDeleteParent
type BatchDeleteParentRequest struct {
	// Useridlist 家长的userid列表，每次最多100个
	Useridlist []string `json:"useridlist" validate:"required"`
}

// BatchDeleteParentResponse is response of Client.BatchDeleteParent
type BatchDeleteParentResponse struct {
	// ResultList 批量操作结果列表
	ResultList []map[string]any `json:"result_list"`
}

// BatchDeleteSchoolStudent 批量删除学生
// 批量删除企业下的学生成员
//
// see https://developer.work.weixin.qq.com/document/path/92329
func (c *Client) BatchDeleteSchoolStudent(r *BatchDeleteSchoolStudentRequest, opts ...any) (out BatchDeleteSchoolStudentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/user/batch_delete_student",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchDeleteSchoolStudentRequest is request of Client.BatchDeleteSchoolStudent
type BatchDeleteSchoolStudentRequest struct {
	// Useridlist 学生的userid列表，每次最多100个
	Useridlist []string `json:"useridlist" validate:"required"`
}

// BatchDeleteSchoolStudentResponse is response of Client.BatchDeleteSchoolStudent
type BatchDeleteSchoolStudentResponse struct {
	// ResultList 批量操作结果
	ResultList []map[string]any `json:"result_list"`
}

// BatchUpdateParent 批量更新家长
// 批量更新家校通讯录中的家长信息
//
// see https://developer.work.weixin.qq.com/document/path/92336
func (c *Client) BatchUpdateParent(r *BatchUpdateParentRequest, opts ...any) (out BatchUpdateParentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/user/batch_update_parent",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchUpdateParentRequest is request of Client.BatchUpdateParent
type BatchUpdateParentRequest struct {
	// Parents 家长列表，每次最多100个
	Parents []map[string]any `json:"parents"`
}

// BatchUpdateParentResponse is response of Client.BatchUpdateParent
type BatchUpdateParentResponse struct {
	// ResultList 失败的学生列表
	ResultList []map[string]any `json:"result_list"`
}

// BatchUpdateStudent 批量更新学生
// 批量更新学生信息，包括手机号、姓名、部门等
//
// see https://developer.work.weixin.qq.com/document/path/92330
func (c *Client) BatchUpdateStudent(r *BatchUpdateStudentRequest, opts ...any) (out BatchUpdateStudentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/user/batch_update_student",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchUpdateStudentRequest is request of Client.BatchUpdateStudent
type BatchUpdateStudentRequest struct {
	// Students 学生列表，每次最多100个
	Students []map[string]any `json:"students"`
	// StudentUserID 学生UserID。学校内必须唯一。不区分大小写，长度为1~64个字节。只能由数字、字母和“_-@.”四种字符组成，且第一个字符必须是数字或字母
	StudentUserID string `json:"student_userid" validate:"required"`
	// Mobile 学生手机号
	Mobile string `json:"mobile"`
	// NewStudentUserID 要变更的学生UserID，不能与已存在的UserID相同。每个学生仅能修改一次
	NewStudentUserID string `json:"new_student_userid"`
	// Name 学生姓名，长度为1~32个字符
	Name string `json:"name"`
	// Department 学生所在的班级id列表，不超过20个
	Department []any `json:"department"`
}

// BatchUpdateStudentResponse is response of Client.BatchUpdateStudent
type BatchUpdateStudentResponse struct {
	// ResultList 失败的学生列表
	ResultList []map[string]any `json:"result_list"`
	// StudentUserID 失败的学生userid
	StudentUserID string `json:"student_userid"`
}

// CreateParentUser 创建家长
// 创建学校家长成员
//
// see https://developer.work.weixin.qq.com/document/path/92331
func (c *Client) CreateParentUser(r *CreateParentUserRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/user/create_parent",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateParentUserRequest is request of Client.CreateParentUser
type CreateParentUserRequest struct {
	// ParentUserID 家长UserID。学校内必须唯一，可以与企业通讯录内成员UserID相同。不区分大小写，长度为1~64个字节。只能由数字、字母和"_-@."四种字符组成，且第一个字符必须是数字或字母。
	ParentUserID string `json:"parent_userid" validate:"required"`
	// Mobile 家长手机号
	Mobile string `json:"mobile" validate:"required"`
	// ToInvite 是否发起邀请，默认为true，仅验证的学校才能发起邀请。
	ToInvite bool `json:"to_invite"`
	// Children 家长的孩子列表，最多10
	Children []map[string]any `json:"children" validate:"required"`
}

// CreateStudent 创建学生
// 创建学校的学生成员
//
// see https://developer.work.weixin.qq.com/document/path/92325
func (c *Client) CreateStudent(r *CreateStudentRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/user/create_student",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateStudentRequest is request of Client.CreateStudent
type CreateStudentRequest struct {
	// StudentUserID 学生UserID。学校内必须唯一，不区分大小写，长度为1~64个字节。只能由数字、字母和"_-@."四种字符组成，且第一个字符必须是数字或字母。
	StudentUserID string `json:"student_userid" validate:"required"`
	// Mobile 学生手机号
	Mobile string `json:"mobile"`
	// ToInvite 是否发起邀请，默认为true，仅验证的学校才能发起邀请。
	ToInvite bool `json:"to_invite"`
	// Name 学生姓名，长度为1~32个字符
	Name string `json:"name" validate:"required"`
	// Department 学生所在的班级id列表，不超过20个
	Department []int `json:"department" validate:"required"`
}

// DeleteParent 删除家长
// 删除家校通讯录中的家长信息
//
// see https://developer.work.weixin.qq.com/document/path/92332
func (c *Client) DeleteParent(r *DeleteParentRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/school/user/delete_parent",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteParentRequest is request of Client.DeleteParent
type DeleteParentRequest struct {
	// UserID 家校通信录中家长的userid
	UserID string `json:"userid" validate:"required"`
}

// DeleteStudent 删除学生
// 删除家校通讯录中的学生成员
//
// see https://developer.work.weixin.qq.com/document/path/92326
func (c *Client) DeleteStudent(r *DeleteStudentRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/school/user/delete_student",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteStudentRequest is request of Client.DeleteStudent
type DeleteStudentRequest struct {
	// UserID 家校通讯录中学生的userid
	UserID string `json:"userid" validate:"required"`
}

// GetSchoolUser 读取学生或家长
// 获取家校通讯录中的学生或家长信息
//
// see https://developer.work.weixin.qq.com/document/path/92337
func (c *Client) GetSchoolUser(r *GetSchoolUserRequest, opts ...any) (out GetSchoolUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/school/user/get",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetSchoolUserRequest is request of Client.GetSchoolUser
type GetSchoolUserRequest struct {
	// UserID 家校通讯录的userid，家长或者学生的userid。不区分大小写，长度为1~64个字节
	UserID string `json:"userid" validate:"required"`
}

// GetSchoolUserResponse is response of Client.GetSchoolUser
type GetSchoolUserResponse struct {
	// UserType 用户类型:1表示学生，2表示家长
	UserType int `json:"user_type"`
	// Student 学生字段，user_type为1时返回该字段
	Student any `json:"student"`
	// Parent 学生家长，user_type为2时返回该字段
	Parent any `json:"parent"`
}

// ListSchoolUser 获取部门学生详情
// 获取指定部门下的学生列表及其家长信息
//
// see https://developer.work.weixin.qq.com/document/path/96119
func (c *Client) ListSchoolUser(r *ListSchoolUserRequest, opts ...any) (out ListSchoolUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/school/user/list",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListSchoolUserRequest is request of Client.ListSchoolUser
type ListSchoolUserRequest struct {
	// DepartmentID 获取的部门id
	DepartmentID int `json:"department_id" validate:"required"`
}

// ListSchoolUserResponse is response of Client.ListSchoolUser
type ListSchoolUserResponse struct {
	// Students 学生列表
	Students []map[string]any `json:"students"`
}

// ListSchoolParent 获取部门家长详情
// 获取指定部门的家长详情列表
//
// see https://developer.work.weixin.qq.com/document/path/92446
func (c *Client) ListSchoolParent(r *ListSchoolParentRequest, opts ...any) (out ListSchoolParentResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/school/user/list_parent",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListSchoolParentRequest is request of Client.ListSchoolParent
type ListSchoolParentRequest struct {
	// DepartmentID 获取的部门id
	DepartmentID int `json:"department_id" validate:"required"`
}

// ListSchoolParentResponse is response of Client.ListSchoolParent
type ListSchoolParentResponse struct {
	// Parents 家长列表
	Parents []map[string]any `json:"parents"`
	// ParentUserID 家长的userid
	ParentUserID string `json:"parent_userid"`
	// Mobile 家长手机号，第三方不可获取;代开发应用需要管理员授权该权限才返回
	Mobile string `json:"mobile"`
	// IsSubscribe 家长是否关注了“学校通知”，0-未关注，1-已关注
	IsSubscribe int `json:"is_subscribe"`
	// ExternalUserID 家长的external_userid,仅当家长已关注才返回。对同一个服务商来说，同一个家长微信在不同学校下返回的家长的external_userid是一样的
	ExternalUserID string `json:"external_userid"`
	// Children 家长孩子列表
	Children []any `json:"children"`
	// StudentUserID 学生的userid
	StudentUserID string `json:"student_userid"`
	// Relation 家长与学生关系
	Relation string `json:"relation"`
	// Name 学生姓名
	Name string `json:"name"`
}

// UpdateParent 更新家长
// 更新家长信息，包括家长UserID、手机号及关联的学生列表
//
// see https://developer.work.weixin.qq.com/document/path/92333
func (c *Client) UpdateParent(r *UpdateParentRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/user/update_parent",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateParentRequest is request of Client.UpdateParent
type UpdateParentRequest struct {
	// ParentUserID 家长UserID。学校内必须唯一
	ParentUserID string `json:"parent_userid" validate:"required"`
	// NewParentUserID 更新的家长UserID
	NewParentUserID string `json:"new_parent_userid"`
	// Mobile 家长手机号
	Mobile string `json:"mobile"`
	// Children 家长的孩子列表，全量更新
	Children []map[string]any `json:"children"`
}

// UpdateStudent 更新学生
// 更新学校成员（学生）信息
//
// see https://developer.work.weixin.qq.com/document/path/92327
func (c *Client) UpdateStudent(r *UpdateStudentRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/school/user/update_student",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateStudentRequest is request of Client.UpdateStudent
type UpdateStudentRequest struct {
	// StudentUserID 学生UserID。学校内必须唯一。不区分大小写，长度为1~64个字节。只能由数字、字母和"_-@."四种字符组成，且第一个字符必须是数字或字母。
	StudentUserID string `json:"student_userid" validate:"required"`
	// Mobile 学生手机号
	Mobile string `json:"mobile"`
	// NewStudentUserID 要变更的学生UserID，不能与已存在的UserID相同。每个学生仅能修改一次。
	NewStudentUserID string `json:"new_student_userid"`
	// Name 学生姓名，长度为1~32个字符
	Name string `json:"name"`
	// Department 学生所在的班级id列表，不超过20个
	Department []int `json:"department"`
}
