package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// AddFileMember 新增成员
// 该接口用于对指定文件添加成员。
//
// see https://developer.work.weixin.qq.com/document/path/97893
func (c *Client) AddFileMember(r *AddFileMemberRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/file_acl_add",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddFileMemberRequest is request of Client.AddFileMember
type AddFileMemberRequest struct {
	// Fileid 文件fileid
	Fileid string `json:"fileid" validate:"required"`
	// AuthInfo 添加成员的信息列表
	AuthInfo []map[string]any `json:"auth_info" validate:"required"`
}

// DeleteMemberAcl 删除成员
// 该接口用于删除指定文件的成员。
//
// see https://developer.work.weixin.qq.com/document/path/97888
func (c *Client) DeleteMemberAcl(r *DeleteMemberAclRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/file_acl_del",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteMemberAclRequest is request of Client.DeleteMemberAcl
type DeleteMemberAclRequest struct {
	// Fileid 文件fileid
	Fileid string `json:"fileid" validate:"required"`
	// AuthInfo 被移除的成员信息
	AuthInfo []map[string]any `json:"auth_info" validate:"required"`
}

// CreateFile 新建文件夹/文档
// 该接口用于在微盘指定位置新建文件夹、文档
//
// see https://developer.work.weixin.qq.com/document/path/97882
func (c *Client) CreateFile(r *CreateFileRequest, opts ...any) (out CreateFileResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/file_create",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateFileRequest is request of Client.CreateFile
type CreateFileRequest struct {
	// Spaceid 空间spaceid
	Spaceid string `json:"spaceid" validate:"required"`
	// Fatherid 父目录fileid, 在根目录时为空间spaceid
	Fatherid string `json:"fatherid" validate:"required"`
	// FileType 文件类型, 1:文件夹 3:文档(文档) 4:文档(表格)
	FileType int `json:"file_type" validate:"required"`
	// FileName 文件名字（注意：文件名最多填255个字符, 英文算1个, 汉字算2个）
	FileName string `json:"file_name" validate:"required"`
}

// CreateFileResponse is response of Client.CreateFile
type CreateFileResponse struct {
	// Fileid 新建文件的fileid
	Fileid string `json:"fileid"`
	// URL 文档的访问链接，仅在新建文档时返回
	URL string `json:"url"`
}

// DeleteFile 删除文件
// 该接口用于删除指定文件。
//
// see https://developer.work.weixin.qq.com/document/path/97885
func (c *Client) DeleteFile(r *DeleteFileRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/file_delete",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteFileRequest is request of Client.DeleteFile
type DeleteFileRequest struct {
	// Fileid 文件fileid
	Fileid []string `json:"fileid" validate:"required"`
}

// DownloadFile 下载文件
// 该接口用于下载文件。
//
// see https://developer.work.weixin.qq.com/document/path/97881
func (c *Client) DownloadFile(r *DownloadFileRequest, opts ...any) (out DownloadFileResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/file_download",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DownloadFileRequest is request of Client.DownloadFile
type DownloadFileRequest struct {
	// Fileid 文件fileid（只支持下载普通文件，不支持下载文件夹或微文档）
	Fileid string `json:"fileid"`
	// SelectedTicket 微盘和文件选择器jsapi返回的selectedTicket。若填此参数，则不需要填fileid
	SelectedTicket string `json:"selected_ticket"`
}

// DownloadFileResponse is response of Client.DownloadFile
type DownloadFileResponse struct {
	// DownloadURL 下载请求url (有效期2个小时)
	DownloadURL string `json:"download_url"`
	// CookieName 下载请求带cookie的key
	CookieName string `json:"cookie_name"`
	// CookieValue 下载请求带cookie的value
	CookieValue string `json:"cookie_value"`
}

// GetDriveFileInfo 获取文件信息
// 该接口用于获取指定文件的信息。
//
// see https://developer.work.weixin.qq.com/document/path/97886
func (c *Client) GetDriveFileInfo(r *GetDriveFileInfoRequest, opts ...any) (out GetDriveFileInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/file_info",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetDriveFileInfoRequest is request of Client.GetDriveFileInfo
type GetDriveFileInfoRequest struct {
	// Fileid 文件fileid
	Fileid string `json:"fileid" validate:"required"`
}

// GetDriveFileInfoResponse is response of Client.GetDriveFileInfo
type GetDriveFileInfoResponse struct {
	// FileInfo 文件信息对象
	FileInfo any `json:"file_info"`
	// Fileid 文件fileid
	Fileid string `json:"fileid"`
	// FileName 文件名字
	FileName string `json:"file_name"`
	// Spaceid 文件所在的空间spaceid
	Spaceid string `json:"spaceid"`
	// Fatherid 文件所在的目录fileid, 在根目录时为fileid
	Fatherid string `json:"fatherid"`
	// FileSize 文件大小
	FileSize int `json:"file_size"`
	// Ctime 文件创建时间
	Ctime int `json:"ctime"`
	// Mtime 文件最后修改时间
	Mtime int `json:"mtime"`
	// FileType 1: 文件夹 2:文件 3: 文档(文档) 4: 文档(表格) 5:文档(收集表) 6:文档(幻灯片)
	FileType int `json:"file_type"`
	// FileStatus 文件状态, 1:正常 2:删除
	FileStatus int `json:"file_status"`
	// Sha 文件sha。可用于确认是否跟与上传的文件一致，或避免重复上传相同的文件
	Sha string `json:"sha"`
	// Md5 文件md5。可用于确认是否跟与上传的文件一致，或避免重复上传相同的文件
	Md5 string `json:"md5"`
	// URL 仅微文档类型返回访问链接
	URL string `json:"url"`
}

// ListDriveFile 获取文件列表
// 该接口用于获取指定地址下的文件列表。
//
// see https://developer.work.weixin.qq.com/document/path/97887
func (c *Client) ListDriveFile(r *ListDriveFileRequest, opts ...any) (out ListDriveFileResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/file_list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListDriveFileRequest is request of Client.ListDriveFile
type ListDriveFileRequest struct {
	// Spaceid 空间spaceid
	Spaceid string `json:"spaceid" validate:"required"`
	// Fatherid 当前目录的fileid，根目录时为空间spaceid
	Fatherid string `json:"fatherid" validate:"required"`
	// SortType 列表排序方式 1:名字升序；2:名字降序；3:大小升序；4:大小降序；5:修改时间升序；6:修改时间降序
	SortType int `json:"sort_type" validate:"required"`
	// Start 首次填0，后续填上一次请求返回的next_start
	Start int `json:"start" validate:"required"`
	// Limit 分批拉取最大文件数，不超过1000
	Limit int `json:"limit" validate:"required"`
}

// ListDriveFileResponse is response of Client.ListDriveFile
type ListDriveFileResponse struct {
	// HasMore true为列表还有内容，需要继续分批拉取
	HasMore bool `json:"has_more"`
	// NextStart 下次分批拉取对应的请求参数start值
	NextStart int `json:"next_start"`
	// FileList 文件列表对象
	FileList any `json:"file_list"`
}

// MoveFile 移动文件
// 该接口用于将文件移动到指定位置。
//
// see https://developer.work.weixin.qq.com/document/path/97884
func (c *Client) MoveFile(r *MoveFileRequest, opts ...any) (out MoveFileResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/file_move",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// MoveFileRequest is request of Client.MoveFile
type MoveFileRequest struct {
	// Fatherid 当前目录的fileid，根目录时为空间spaceid
	Fatherid string `json:"fatherid" validate:"required"`
	// Replace 如果移动到的目标目录与需要移动的文件重名时，是否覆盖。true:重名文件覆盖 false:重名文件进行冲突重命名处理
	Replace bool `json:"replace"`
	// Fileid 文件fileid列表
	Fileid []string `json:"fileid" validate:"required"`
}

// MoveFileResponse is response of Client.MoveFile
type MoveFileResponse struct {
	// FileList 移动文件的信息列表
	FileList any `json:"file_list"`
}

// RenameFile 重命名文件
// 该接口用于对指定文件进行重命名。
//
// see https://developer.work.weixin.qq.com/document/path/97883
func (c *Client) RenameFile(r *RenameFileRequest, opts ...any) (out RenameFileResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/file_rename",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// RenameFileRequest is request of Client.RenameFile
type RenameFileRequest struct {
	// Fileid 文件fileid
	Fileid string `json:"fileid" validate:"required"`
	// NewName 重命名后的文件名（注意：文件名最多填255个字符，英文算1个，汉字算2个）
	NewName string `json:"new_name" validate:"required"`
}

// RenameFileResponse is response of Client.RenameFile
type RenameFileResponse struct {
	// File 文件详细信息对象
	File any `json:"file"`
}

// UpdateFileSecureSetting 修改文件安全设置
// 该接口用于修改文件安全设置，水印相关设置。仅支持在线文档类型。
//
// see https://developer.work.weixin.qq.com/document/path/97892
func (c *Client) UpdateFileSecureSetting(r *UpdateFileSecureSettingRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/file_secure_setting",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateFileSecureSettingRequest is request of Client.UpdateFileSecureSetting
type UpdateFileSecureSettingRequest struct {
	// Fileid 文件fileid
	Fileid string `json:"fileid" validate:"required"`
	// Watermark 水印设置对象
	Watermark any `json:"watermark"`
}

// SetFileShareSetting 分享设置
// 该接口用于文件的分享设置。
//
// see https://developer.work.weixin.qq.com/document/path/97889
func (c *Client) SetFileShareSetting(r *SetFileShareSettingRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/file_setting",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetFileShareSettingRequest is request of Client.SetFileShareSetting
type SetFileShareSettingRequest struct {
	// Fileid 文件fileid
	Fileid string `json:"fileid" validate:"required"`
	// AuthScope 权限范围：1:指定人 2:企业内 3:企业外 4: 企业内需管理员审批（仅有管理员时可设置）5: 企业外需管理员审批（仅有管理员时可设置）
	AuthScope int `json:"auth_scope" validate:"required"`
	// Auth 权限信息。普通文档：1:仅浏览（可下载) 4:仅预览（仅专业版企业可设置）；如果不填充此字段为保持原有状态。微文档：1:仅浏览（可下载）；如果不填充此字段为保持原有状态
	Auth int `json:"auth"`
}

// GetFileShareUrl 获取分享链接
// 该接口用于获取文件的分享链接
//
// see https://developer.work.weixin.qq.com/document/path/97890
func (c *Client) GetFileShareUrl(r *GetFileShareUrlRequest, opts ...any) (out GetFileShareUrlResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/file_share",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetFileShareUrlRequest is request of Client.GetFileShareUrl
type GetFileShareUrlRequest struct {
	// Fileid 文件fileid
	Fileid string `json:"fileid" validate:"required"`
}

// GetFileShareUrlResponse is response of Client.GetFileShareUrl
type GetFileShareUrlResponse struct {
	// ShareURL 分享文件的链接
	ShareURL string `json:"share_url"`
}

// UploadFile 上传文件
// 该接口用于向微盘中的指定位置上传文件。
//
// see https://developer.work.weixin.qq.com/document/path/97880
func (c *Client) UploadFile(r *UploadFileRequest, opts ...any) (out UploadFileResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/file_upload",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UploadFileRequest is request of Client.UploadFile
type UploadFileRequest struct {
	// Spaceid 空间spaceid
	Spaceid string `json:"spaceid"`
	// Fatherid 父目录fileid, 在根目录时为空间spaceid
	Fatherid string `json:"fatherid"`
	// SelectedTicket 微盘和文件选择器jsapi返回的selectedTicket。若填此参数，则不需要填spaceid/fatherid。
	SelectedTicket string `json:"selected_ticket"`
	// FileName 文件名字（注意：文件名最多填255个字符, 英文算1个, 汉字算2个）
	FileName string `json:"file_name" validate:"required"`
	// FileBase64Content 文件内容base64（注意：只需要填入文件内容的Base64，不需要添加任何如:data:application/x-javascript;base64的数据类型描述信息），文件大小上限为10M...
	FileBase64Content string `json:"file_base64_content" validate:"required"`
}

// UploadFileResponse is response of Client.UploadFile
type UploadFileResponse struct {
	// Fileid 新建文件的fileid
	Fileid string `json:"fileid"`
}

// InitWedriveFileUpload 文件分块上传初始化
// 请求分块上传初始化接口，如果命中秒传，则流程结束，完成上传。
//
// see https://developer.work.weixin.qq.com/document/path/98004
func (c *Client) InitWedriveFileUpload(r *InitWedriveFileUploadRequest, opts ...any) (out InitWedriveFileUploadResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/file_upload_init",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// InitWedriveFileUploadRequest is request of Client.InitWedriveFileUpload
type InitWedriveFileUploadRequest struct {
	// Spaceid 空间spaceid
	Spaceid string `json:"spaceid"`
	// Fatherid 当前目录的fileid，根目录时为空间spaceid
	Fatherid string `json:"fatherid"`
	// SelectedTicket 微盘和文件选择器jsapi返回的selectedTicket。若填此参数，则不需要填spaceid/fatherid
	SelectedTicket string `json:"selected_ticket"`
	// FileName 文件名字
	FileName string `json:"file_name" validate:"required"`
	// Size 文件大小。最大支持20G
	Size int `json:"size" validate:"required"`
	// BlockSha 文件分块累积sha值，按分块顺序填入数组
	BlockSha []string `json:"block_sha" validate:"required"`
	// SkipPushCard 文件创建完成时是否推送企业微信卡片。默认false
	SkipPushCard bool `json:"skip_push_card"`
}

// InitWedriveFileUploadResponse is response of Client.InitWedriveFileUpload
type InitWedriveFileUploadResponse struct {
	// HitExist 是否命中秒传
	HitExist bool `json:"hit_exist"`
	// UploadKey 文件上传凭证。不命中秒传时返回，作为file_upload_part参数
	UploadKey string `json:"upload_key"`
	// Fileid 文件fileid。命中秒传时返回，此时上传流程完成
	Fileid string `json:"fileid"`
}

// GetFilePermission 获取文件权限信息
// 该接口用于获取文件的权限信息。
//
// see https://developer.work.weixin.qq.com/document/path/97891
func (c *Client) GetFilePermission(r *GetFilePermissionRequest, opts ...any) (out GetFilePermissionResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/get_file_permission",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetFilePermissionRequest is request of Client.GetFilePermission
type GetFilePermissionRequest struct {
	// Fileid 文件fileid
	Fileid string `json:"fileid" validate:"required"`
}

// GetFilePermissionResponse is response of Client.GetFilePermission
type GetFilePermissionResponse struct {
	// ShareRange 文件分享设置
	ShareRange any `json:"share_range"`
	// EnableCorpInternal 是否为企业内可访问
	EnableCorpInternal bool `json:"enable_corp_internal"`
	// CorpInternalAuth 企业内权限信息
	CorpInternalAuth int `json:"corp_internal_auth"`
	// EnableCorpExternal 是否为企业外可访问
	EnableCorpExternal bool `json:"enable_corp_external"`
	// CorpExternalAuth 企业外权限信息
	CorpExternalAuth int `json:"corp_external_auth"`
	// SecureSetting 文件安全配置
	SecureSetting any `json:"secure_setting"`
	// EnableReadonlyCopy 是否开启只读备份
	EnableReadonlyCopy bool `json:"enable_readonly_copy"`
	// ModifyOnlyByAdmin 是否只允许管理员进行修改
	ModifyOnlyByAdmin bool `json:"modify_only_by_admin"`
	// EnableReadonlyComment 是否开启只读评论
	EnableReadonlyComment bool `json:"enable_readonly_comment"`
	// BanShareExternal 是否禁止分享到企业外部
	BanShareExternal bool `json:"ban_share_external"`
	// InheritFatherAuth 从文件父路径继承的权限
	InheritFatherAuth any `json:"inherit_father_auth"`
	// Inherit 文件是否开启父路径权限继承
	Inherit bool `json:"inherit"`
	// FileMemberList 查询fileid为文档时返回，为文档所在目录成员，以及其他授权列表
	FileMemberList []map[string]any `json:"file_member_list"`
	// Watermark 水印相关设置
	Watermark any `json:"watermark"`
	// Text 水印文字
	Text string `json:"text"`
	// MarginType 水印类型
	MarginType int `json:"margin_type"`
	// ShowVisitorName 是否显示访问人名称
	ShowVisitorName bool `json:"show_visitor_name"`
	// ForceByAdmin 管理员是否强制要求使用水印
	ForceByAdmin bool `json:"force_by_admin"`
	// ShowText 是否展示水印文本
	ShowText bool `json:"show_text"`
	// ForceBySpaceAdmin 空间管理员是否强制要求使用水印
	ForceBySpaceAdmin bool `json:"force_by_space_admin"`
}

// GetWeDriveProInfo 获取盘专业版信息
// 该接口用于获取专业版信息。
//
// see https://developer.work.weixin.qq.com/document/path/95856
func (c *Client) GetWeDriveProInfo(r *GetWeDriveProInfoRequest, opts ...any) (out GetWeDriveProInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/mng_pro_info",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetWeDriveProInfoRequest is request of Client.GetWeDriveProInfo
type GetWeDriveProInfoRequest struct {}

// GetWeDriveProInfoResponse is response of Client.GetWeDriveProInfo
type GetWeDriveProInfoResponse struct {
	// IsPro true为专业版，false为不是专业版
	IsPro bool `json:"is_pro"`
	// TotalVipAcctNum 总的vip账号数量
	TotalVipAcctNum int `json:"total_vip_acct_num"`
	// UseVipAcctNum 已使用的vip账号数量
	UseVipAcctNum int `json:"use_vip_acct_num"`
	// ProExpireTime 专业版到期时间，时间戳，精确到秒
	ProExpireTime int `json:"pro_expire_time"`
}

// GetSpaceInfo 获取空间信息
// 该接口用于获取空间信息。包括：空间成员及权限及安全设置。
//
// see https://developer.work.weixin.qq.com/document/path/97878
func (c *Client) GetSpaceInfo(r *GetSpaceInfoRequest, opts ...any) (out GetSpaceInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/new_space_info",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetSpaceInfoRequest is request of Client.GetSpaceInfo
type GetSpaceInfoRequest struct {
	// Spaceid 空间spaceid
	Spaceid string `json:"spaceid" validate:"required"`
}

// GetSpaceInfoResponse is response of Client.GetSpaceInfo
type GetSpaceInfoResponse struct {
	// SpaceInfo 空间信息
	SpaceInfo any `json:"space_info"`
}

// AddSpaceMemberDepartment 添加成员/部门
// 该接口用于对指定空间添加成员/部门，可一次性添加多个。
//
// see https://developer.work.weixin.qq.com/document/path/97879
func (c *Client) AddSpaceMemberDepartment(r *AddSpaceMemberDepartmentRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/space_acl_add",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddSpaceMemberDepartmentRequest is request of Client.AddSpaceMemberDepartment
type AddSpaceMemberDepartmentRequest struct {
	// Spaceid 空间spaceid
	Spaceid string `json:"spaceid" validate:"required"`
	// AuthInfo 被添加的空间成员信息列表
	AuthInfo []map[string]any `json:"auth_info" validate:"required"`
}

// RemoveSpaceMemberOrDept 移除成员/部门
// 该接口用于对指定空间移除成员/部门，操作者为应用本身。
//
// see https://developer.work.weixin.qq.com/document/path/97875
func (c *Client) RemoveSpaceMemberOrDept(r *RemoveSpaceMemberOrDeptRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/space_acl_del",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// RemoveSpaceMemberOrDeptRequest is request of Client.RemoveSpaceMemberOrDept
type RemoveSpaceMemberOrDeptRequest struct {
	// Spaceid 空间spaceid
	Spaceid string `json:"spaceid" validate:"required"`
	// AuthInfo 被移除的空间成员信息
	AuthInfo []map[string]any `json:"auth_info" validate:"required"`
}

// CreateSpace 新建空间
// 该接口用于在微盘内新建空间，创建者为应用本身。
//
// see https://developer.work.weixin.qq.com/document/path/97860
func (c *Client) CreateSpace(r *CreateSpaceRequest, opts ...any) (out CreateSpaceResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/space_create",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateSpaceRequest is request of Client.CreateSpace
type CreateSpaceRequest struct {
	// SpaceName 空间标题
	SpaceName string `json:"space_name" validate:"required"`
	// AuthInfo 空间其他成员信息列表
	AuthInfo []map[string]any `json:"auth_info"`
	// SpaceSubType 区分创建空间类型，0:普通（目前只支持0）
	SpaceSubType int `json:"space_sub_type"`
}

// CreateSpaceResponse is response of Client.CreateSpace
type CreateSpaceResponse struct {
	// Spaceid 空间id
	Spaceid string `json:"spaceid"`
}

// DismissSpace 解散空间
// 该接口用于解散已有空间。
//
// see https://developer.work.weixin.qq.com/document/path/97857
func (c *Client) DismissSpace(r *DismissSpaceRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/space_dismiss",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DismissSpaceRequest is request of Client.DismissSpace
type DismissSpaceRequest struct {
	// Spaceid 空间spaceid
	Spaceid string `json:"spaceid" validate:"required"`
}

// RenameSpace 重命名空间
// 该接口用于重命名已有空间。
//
// see https://developer.work.weixin.qq.com/document/path/97856
func (c *Client) RenameSpace(r *RenameSpaceRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/space_rename",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// RenameSpaceRequest is request of Client.RenameSpace
type RenameSpaceRequest struct {
	// Spaceid 空间spaceid
	Spaceid string `json:"spaceid" validate:"required"`
	// SpaceName 重命名后的空间名
	SpaceName string `json:"space_name" validate:"required"`
}

// UpdateSpaceSetting 安全设置
// 修改空间权限，应用通过api调用仅支持设置由本应用创建的空间
//
// see https://developer.work.weixin.qq.com/document/path/97876
func (c *Client) UpdateSpaceSetting(r *UpdateSpaceSettingRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/space_setting",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateSpaceSettingRequest is request of Client.UpdateSpaceSetting
type UpdateSpaceSettingRequest struct {
	// Spaceid 空间spaceid
	Spaceid string `json:"spaceid" validate:"required"`
	// EnableWatermark 启用水印。false:关 true:开 ;如果不填充此字段为保持原有状态
	EnableWatermark bool `json:"enable_watermark"`
	// EnableConfidentialMode 是否开启保密模式。false:关 true:开 如果不填充此字段为保持原有状态
	EnableConfidentialMode bool `json:"enable_confidential_mode"`
	// ShareURLNoApprove 通过链接加入空间无需审批。false:关；true:开；如果不填充此字段为保持原有状态
	ShareURLNoApprove bool `json:"share_url_no_approve"`
	// ShareURLNoApproveDefaultAuth 邀请链接默认权限。1:仅下载 2:可编辑 4:仅预览 5:可上传下载 200:自定义权限；如果不填充此字段为保持原有状态
	ShareURLNoApproveDefaultAuth int `json:"share_url_no_approve_default_auth"`
	// DefaultFileScope 文件默认可查看范围。1:仅成员；2:企业内。如果不填充此字段为保持原有状态
	DefaultFileScope int `json:"default_file_scope"`
	// BanShareExternal 是否禁止文件分享到企业外｜false:关 true:开 如果不填充此字段为保持原有状态
	BanShareExternal bool `json:"ban_share_external"`
}

// GetSpaceShareLink 获取邀请链接
// 该接口用于获取空间邀请分享链接。
//
// see https://developer.work.weixin.qq.com/document/path/97877
func (c *Client) GetSpaceShareLink(r *GetSpaceShareLinkRequest, opts ...any) (out GetSpaceShareLinkResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/space_share",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetSpaceShareLinkRequest is request of Client.GetSpaceShareLink
type GetSpaceShareLinkRequest struct {
	// Spaceid 空间spaceid
	Spaceid string `json:"spaceid" validate:"required"`
}

// GetSpaceShareLinkResponse is response of Client.GetSpaceShareLink
type GetSpaceShareLinkResponse struct {
	// SpaceShareURL 邀请链接
	SpaceShareURL string `json:"space_share_url"`
}

// BatchAddDriveVip 分配高级功能账号
// 该接口用于分配应用可见范围内企业成员的高级功能。
//
// see https://developer.work.weixin.qq.com/document/path/99512
func (c *Client) BatchAddDriveVip(r *BatchAddDriveVipRequest, opts ...any) (out BatchAddDriveVipResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/vip/batch_add",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchAddDriveVipRequest is request of Client.BatchAddDriveVip
type BatchAddDriveVipRequest struct {
	// UserIDList 要分配高级功能的企业成员userid列表，单次操作最大限制100个
	UserIDList []string `json:"userid_list" validate:"required"`
}

// BatchAddDriveVipResponse is response of Client.BatchAddDriveVip
type BatchAddDriveVipResponse struct {
	// SuccUserIDList 分配成功的userid列表，包括已经是高级功能账号的userid
	SuccUserIDList []string `json:"succ_userid_list"`
	// FailUserIDList 分配失败的userid列表
	FailUserIDList []string `json:"fail_userid_list"`
}

// BatchDelDriveVip 取消高级功能账号
// 该接口用于撤销分配应用可见范围企业成员的高级功能。
//
// see https://developer.work.weixin.qq.com/document/path/99513
func (c *Client) BatchDelDriveVip(r *BatchDelDriveVipRequest, opts ...any) (out BatchDelDriveVipResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/vip/batch_del",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchDelDriveVipRequest is request of Client.BatchDelDriveVip
type BatchDelDriveVipRequest struct {
	// UserIDList 要撤销分配高级功能的企业成员userid列表，单次操作最多限制100个
	UserIDList []string `json:"userid_list" validate:"required"`
}

// BatchDelDriveVipResponse is response of Client.BatchDelDriveVip
type BatchDelDriveVipResponse struct {
	// SuccUserIDList 撤销分配成功的userid列表
	SuccUserIDList []string `json:"succ_userid_list"`
	// FailUserIDList 撤销分配失败的userid列表
	FailUserIDList []string `json:"fail_userid_list"`
}

// ListDriveVip 获取高级功能账号列表
// 查询企业已分配高级功能且在应用可见范围的账号列表
//
// see https://developer.work.weixin.qq.com/document/path/99514
func (c *Client) ListDriveVip(r *ListDriveVipRequest, opts ...any) (out ListDriveVipResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/wedrive/vip/list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListDriveVipRequest is request of Client.ListDriveVip
type ListDriveVipRequest struct {
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"`
	// Limit 用于分页查询，每次请求返回的数据上限。默认100，最大200
	Limit int `json:"limit"`
}

// ListDriveVipResponse is response of Client.ListDriveVip
type ListDriveVipResponse struct {
	// HasMore 是否还有更多数据未获取
	HasMore bool `json:"has_more"`
	// NextCursor 下一次请求的cursor值
	NextCursor string `json:"next_cursor"`
	// UserIDList 符合条件的企业成员userid列表
	UserIDList []string `json:"userid_list"`
}

