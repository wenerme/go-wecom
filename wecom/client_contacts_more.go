package wecom

// RootDepartmentID is 1
const RootDepartmentID = 1

// SimpleListUserResponseItem is item model of SimpleListUserResponse.UserList
type SimpleListUserResponseItem struct {
	// UserID 成员UserID。对应管理端的帐号
	UserID string `json:"userid"`
	// Name 成员名称，代开发自建应用需要管理员授权才返回；此字段从2019年12月30日起，对新创建第三方应用不再返回真实name，使用userid代替name，2020年6月30日起，对所有历史第三方应用不再返回真实name，使用userid代替name，后续第三方仅通讯录应用可获取，未返回名称的情况需要通过通讯录展示组件来展示名字
	Name string `json:"name"`
	// Department 成员所属部门列表。列表项为部门ID，32位整型
	Department []int `json:"department"`
	// OpenUserID 全局唯一。对于同一个服务商，不同应用获取到企业内同一个成员的open_userid是相同的，最多64个字节。仅第三方应用可获取
	OpenUserID string `json:"open_userid"`
}

// ListUserResponseItem is item model of ListUserResponse.UserList
type ListUserResponseItem struct {
	// UserID 成员UserID。对应管理端的帐号
	UserID string `json:"userid"`
	// Name 成员名称；第三方不可获取，调用时返回userid以代替name；代开发自建应用需要管理员授权才返回；对于非第三方创建的成员，第三方通讯录应用也不可获取；未返回名称的情况需要通过通讯录展示组件来展示名字
	Name string `json:"name"`
	// Mobile 手机号码，代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	Mobile string `json:"mobile"`
	// Department 成员所属部门id列表，仅返回该应用有查看权限的部门id
	Department []int `json:"department"`
	// Order 部门内的排序值，默认为0。数量必须和department一致，数值越大排序越前面。值范围是[0, 2^32)
	Order []int `json:"order"`
	// Position 职务信息；代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	Position string `json:"position"`
	// Gender 性别。0表示未定义，1表示男性，2表示女性
	Gender string `json:"gender"`
	// Email 邮箱，代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	Email string `json:"email"`
	// IsLeaderInDept 表示在所在的部门内是否为上级。0-否；1-是。是一个列表，数量必须与department一致。第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	IsLeaderInDept []int `json:"is_leader_in_dept"`
	// Avatar 头像url。 第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	Avatar string `json:"avatar"`
	// ThumbAvatar 头像缩略图url。第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	ThumbAvatar string `json:"thumb_avatar"`
	// Telephone 座机。代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	Telephone string `json:"telephone"`
	// Alias 别名；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	Alias string `json:"alias"`
	// ExtAttr 扩展属性，代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	ExtAttr ExtAttrs `json:"extattr"`
	// Status 激活状态: 1&#x3D;已激活，2&#x3D;已禁用，4&#x3D;未激活，5&#x3D;退出企业。已激活代表已激活企业微信或已关注微工作台（原企业号）。未激活代表既未激活企业微信又未关注微工作台（原企业号）。
	Status int `json:"status"`
	// QrCode 员工个人二维码，扫描可添加为外部联系人(注意返回的是一个url，可在浏览器上打开该url以展示二维码)；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	QrCode string `json:"qr_code"`
	// ExternalProfile 成员对外属性，字段详情见对外属性；代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	ExternalProfile UserExternalProfile `json:"external_profile"`
	// ExternalPosition 对外职务，如果设置了该值，则以此作为对外展示的职务，否则以position来展示。代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	ExternalPosition string `json:"external_position"`
	// Address 地址。代开发自建应用需要管理员授权才返回；第三方仅通讯录应用可获取；对于非第三方创建的成员，第三方通讯录应用也不可获取
	Address string `json:"address"`
	// OpenUserID 全局唯一。对于同一个服务商，不同应用获取到企业内同一个成员的open_userid是相同的，最多64个字节。仅第三方应用可获取
	OpenUserID string `json:"open_userid"`
	// MainDepartment 主部门
	MainDepartment int    `json:"main_department"`
	EnglishName    string `json:"english_name"`
	HideMobile     int    `json:"hide_mobile"`
}

type UserExternalProfile struct {
	// ExternalCorpName 企业对外简称，需从已认证的企业简称中选填。可在“我的企业”页中查看企业简称认证状态。
	ExternalCorpName string `json:"external_corp_name"`
	// WechatChannels 视频号属性
	WechatChannels struct {
		// Nickname 视频号名字（设置后，成员将对外展示该视频号）
		Nickname string `json:"nickname"`
		// Status 对外展示视频号状态。0表示企业视频号已被确认，可正常使用，1表示企业视频号待确认
		Status int `json:"status"`
	} `json:"wechat_channels"`
	ExternalAttr []ExtAttr `json:"external_attr"`
}

type BatchSyncCallback struct {
	// URL 企业应用接收企业微信推送请求的访问协议和地址，支持http或https协议
	URL string `json:"url"`
	// Token 用于生成签名
	Token string `json:"token"`
	// EncodingAesKey 用于消息体的加密，是AES密钥的Base64编码
	EncodingAesKey string `json:"encodingaeskey"`
}

type BatchGetResult struct {
	UserID       string `json:"userid"`
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

// ListDepartmentResponseItem is item model of ListDepartmentResponse.Department
type ListDepartmentResponseItem struct {
	// ID 创建的部门id
	ID int `json:"id"`
	// Name 部门名称，代开发自建应用需要管理员授权才返回；此字段从2019年12月30日起，对新创建第三方应用不再返回，2020年6月30日起，对所有历史第三方应用不再返回name，返回的name字段使用id代替，后续第三方仅通讯录应用可获取，未返回名称的情况需要通过通讯录展示组件来展示部门名称
	Name string `json:"name"`
	// NameEn 英文名称，此字段从2019年12月30日起，对新创建第三方应用不再返回，2020年6月30日起，对所有历史第三方应用不再返回该字段
	NameEn string `json:"name_en"`
	// ParentID 父部门id。根部门为1
	ParentID int `json:"parentid"`
	// Order 在父部门中的次序值。order值大的排序靠前。值范围是[0, 2^32)
	Order int `json:"order"`
}

// LinkSimpleListUserResponseItem is item model of LinkSimpleListUserResponse.UserList
type LinkSimpleListUserResponseItem struct {
	// UserID 成员UserID。对应管理端的帐号
	UserID string `json:"userid"`
	// Name 成员真实名称
	Name string `json:"name"`
	// Department 成员所属部门id列表，这个字段只会返回传入的department_id所属的互联企业里的部门id
	Department []string `json:"department"`
	// CorpID 所属企业的corpid
	CorpID string `json:"corpid"`
}

// LinkListUserResponseItem is item model of LinkListUserResponse.UserList
type LinkListUserResponseItem struct {
	// UserID 成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节
	UserID string `json:"userid"`
	// Name 成员真实名称
	Name string `json:"name"`
	// Mobile 手机号码
	Mobile string `json:"mobile"`
	// Department 成员所属部门id列表，这个字段只会返回传入的department_id所属的互联企业里的部门id
	Department []string `json:"department"`
	// Position 职务信息
	Position string `json:"position"`
	// Email 邮箱
	Email string `json:"email"`
	// Telephone 座机
	Telephone string `json:"telephone"`
	// CorpID 所属企业的corpid
	CorpID string `json:"corpid"`
	// ExtAttr 扩展属性
	ExtAttr ExtAttrs `json:"extattr"`
}

// LinkListDepartmentResponseItem is item model of LinkListDepartmentResponse.DepartmentList
type LinkListDepartmentResponseItem struct {
	// DepartmentID 部门id
	DepartmentID string `json:"department_id"`
	// DepartmentName 部门名称
	DepartmentName string `json:"department_name"`
	// ParentID 上级部门的id
	ParentID string `json:"parentid"`
	// Order 排序值
	Order int `json:"order"`
}

// LinkGetUserResponseUserInfo is model of LinkGetUserResponse.UserInfo
type LinkGetUserResponseUserInfo struct {
	// UserID 成员UserID。对应管理端的帐号，企业内必须唯一。不区分大小写，长度为1~64个字节
	UserID string `json:"userid"`
	// Name 成员真实名称
	Name string `json:"name"`
	// Mobile 手机号码
	Mobile string `json:"mobile"`
	// Department 成员所属部门id列表，这个字段会返回在应用可见范围内，该用户所在的所有互联企业的部门
	Department []string `json:"department"`
	// Position 职务信息
	Position string `json:"position"`
	// Email 邮箱
	Email string `json:"email"`
	// Telephone 座机
	Telephone string `json:"telephone"`
	// CorpID 所属企业的corpid
	CorpID string `json:"corpid"`
	// ExtAttr 扩展属性
	ExtAttr ExtAttrs `json:"extattr"`
}

type GetTagResponseUserItem struct {
	// UserID 成员帐号
	UserID string `json:"userid"`
	// Name 成员名称，代开发自建应用需要管理员授权才返回该字段；此字段从2019年12月30日起，对新创建第三方应用不再返回，2020年6月30日起，对所有历史第三方应用不再返回，后续第三方仅通讯录应用可获取，未返回名称的情况需要通过通讯录展示组件来展示名字
	Name string `json:"name"`
}

// ListTagResponseItem is model of ListTagResponse.TagList
type ListTagResponseItem struct {
	// TagID 标签id
	TagID int `json:"tagid"`
	// TagName 标签名
	TagName string `json:"tagname"`
}
