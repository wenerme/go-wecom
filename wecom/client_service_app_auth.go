package wecom

import (
	"github.com/wenerme/go-req"
)

// ProviderGetSuiteToken 获取第三方应用凭证
// 该API用于获取第三方应用凭证（suite_access_token）。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/90600
func (c *Client) ProviderGetSuiteToken(r *ProviderGetSuiteTokenRequest, opts ...interface{}) (out SuiteTokenResponse, err error) {
	opts = append([]interface{}{WithoutAccessToken}, opts...)
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/service/get_suite_token",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ProviderGetPreAuthCode 获取预授权码
// 该API用于获取预授权码。预授权码用于企业授权时的第三方服务商安全验证。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/90601
func (c *Client) ProviderGetPreAuthCode(r *ProviderGetPreAuthCodeRequest, opts ...interface{}) (out PreAuthCodeResponse, err error) {
	opts = append([]interface{}{WithSuiteAccessToken}, opts...)
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/service/get_pre_auth_code",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ProviderSetSessionInfo 设置授权配置
// 该接口可对某次授权进行配置。可支持测试模式（应用未发布时）。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/90602
func (c *Client) ProviderSetSessionInfo(r *ProviderSetSessionInfoRequest, opts ...interface{}) (out GenericResponse, err error) {
	opts = append([]interface{}{WithSuiteAccessToken}, opts...)
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/service/set_session_info",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ProviderGetPermanentCode 获取企业永久授权码
// 该API用于使用临时授权码换取授权方的永久授权码，并换取授权信息、企业access_token，临时授权码一次有效。建议第三方以userid为主键，来建立自己的管理员账号。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/90603
func (c *Client) ProviderGetPermanentCode(r *ProviderGetPermanentCodeRequest, opts ...interface{}) (out ProviderGetPermanentCodeResponse, err error) {
	opts = append([]interface{}{WithSuiteAccessToken}, opts...)
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/service/get_permanent_code",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ProviderGetAuthInfo 获取企业授权信息
// 该API用于通过永久授权码换取企业微信的授权信息。 永久code的获取，是通过临时授权码使用get_permanent_code 接口获取到的permanent_code。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/90604
func (c *Client) ProviderGetAuthInfo(r *ProviderGetAuthInfoRequest, opts ...interface{}) (out ProviderGetAuthInfoResponse, err error) {
	opts = append([]interface{}{WithSuiteAccessToken}, opts...)
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/service/get_auth_info",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ProviderGetCorpToken 获取企业凭证
// 第三方服务商在取得企业的永久授权码后，通过此接口可以获取到企业的access_token。获取后可通过通讯录、应用、消息等企业接口来运营这些应用。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/90605
func (c *Client) ProviderGetCorpToken(r *ProviderGetCorpTokenRequest, opts ...interface{}) (out ProviderGetCorpTokenResponse, err error) {
	opts = append([]interface{}{WithSuiteAccessToken}, opts...)
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/service/get_corp_token",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ProviderGetAdminList 获取应用的管理员列表
// 第三方服务商可以用此接口获取授权企业中某个第三方应用的管理员列表(不包括外部管理员)，以便服务商在用户进入应用主页之后根据是否管理员身份做权限的区分。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/90606
func (c *Client) ProviderGetAdminList(r *ProviderGetAdminListRequest, opts ...interface{}) (out ProviderGetAdminListResponse, err error) {
	opts = append([]interface{}{WithSuiteAccessToken}, opts...)
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/service/get_admin_list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

type ProviderGetSuiteTokenRequest struct {
	// SuiteID 以ww或wx开头应用id（对应于旧的以tj开头的套件id）
	SuiteID string `json:"suite_id"  validate:"required"`
	// SuiteSecret 应用secret
	SuiteSecret string `json:"suite_secret"  validate:"required"`
	// SuiteTicket 企业微信后台推送的ticket
	SuiteTicket string `json:"suite_ticket"  validate:"required"`
}

type SuiteTokenResponse struct {
	// SuiteAccessToken 第三方应用access_token,最长为512字节
	SuiteAccessToken string `json:"suite_access_token"  `
	// ExpiresIn 有效期（秒）
	ExpiresIn int `json:"expires_in"  `
	// ExpireAt 有效期至 - 客户端维护字段
	ExpireAt int64 `json:"expire_at"  `
}

type ProviderGetPreAuthCodeRequest struct {
	// SuiteAccessToken 第三方应用access_token,最长为512字节
	SuiteAccessToken string `json:"suite_access_token"  `
}

type PreAuthCodeResponse struct {
	// PreAuthCode 预授权码,最长为512字节
	PreAuthCode string `json:"pre_auth_code"  `
	// ExpiresIn 有效期（秒）
	ExpiresIn int `json:"expires_in"  `
	// ExpireAt 有效期至 - 客户端维护字段
	ExpireAt int64 `json:"expire_at"  `
}

type ProviderSetSessionInfoRequest struct {
	// SuiteAccessToken 第三方应用access_token
	SuiteAccessToken string `json:"suite_access_token"  `
	// PreAuthCode 预授权码
	PreAuthCode string `json:"pre_auth_code"  validate:"required"`
	// SessionInfo 本次授权过程中需要用到的会话信息
	SessionInfo SetSessionInfo `json:"session_info"  validate:"required"`
}

type ProviderGetPermanentCodeRequest struct {
	// AuthCode 临时授权码，会在授权成功时附加在redirect_uri中跳转回第三方服务商网站，或通过授权成功通知回调推送给服务商。长度为64至512个字节
	AuthCode string `json:"auth_code"  validate:"required"`
}

type ProviderGetPermanentCodeResponse struct {
	// AccessToken 授权方（企业）access_token,最长为512字节
	AccessToken string `json:"access_token"  `
	// ExpiresIn 授权方（企业）access_token超时时间（秒）
	ExpiresIn int `json:"expires_in"  `
	// PermanentCode 企业微信永久授权码,最长为512字节
	PermanentCode string `json:"permanent_code"  `
	// AuthCorpInfo 授权方企业信息
	AuthCorpInfo ProviderGetPermanentCodeResponseAuthCorpInfo `json:"auth_corp_info"  `
	// AuthInfo 授权信息。如果是通讯录应用，且没开启实体应用，是没有该项的。通讯录应用拥有企业通讯录的全部信息读写权限
	AuthInfo ProviderGetPermanentCodeResponseAuthInfo `json:"auth_info"  `
	// AuthUserInfo 授权管理员的信息，可能不返回（企业互联由上级企业共享第三方应用给下级时，不返回授权的管理员信息）
	AuthUserInfo ProviderGetPermanentCodeResponseAuthUserInfo `json:"auth_user_info"  `
	// DealerCorpInfo 代理服务商企业信息。应用被代理后才有该信息
	DealerCorpInfo ProviderGetPermanentCodeResponseDealerCorpInfo `json:"dealer_corp_info"  `
	// RegisterCodeInfo 推广二维码安装相关信息，扫推广二维码安装时返回。（注：无论企业是否新注册，只要通过扫推广二维码安装，都会返回该字段）
	RegisterCodeInfo ProviderGetPermanentCodeResponseRegisterCodeInfo `json:"register_code_info"  `
}

type ProviderGetPermanentCodeResponseAuthCorpInfo struct {
	// CorpID 授权方企业微信id
	CorpID string `json:"corpid"  `
	// CorpName 授权方企业名称，即企业简称
	CorpName string `json:"corp_name"  `
	// CorpType 授权方企业类型，认证号：verified, 注册号：unverified
	CorpType string `json:"corp_type"  `
	// CorpSquareLogoURL 授权方企业方形头像
	CorpSquareLogoURL string `json:"corp_square_logo_url"  `
	// CorpUserMax 授权方企业用户规模
	CorpUserMax int `json:"corp_user_max"  `
	// CorpAgentMax 授权方企业应用数上限
	CorpAgentMax int `json:"corp_agent_max"  `
	// CorpFullName 授权方企业的主体名称(仅认证或验证过的企业有)，即企业全称。
	CorpFullName string `json:"corp_full_name"  `
	// SubjectType 企业类型，1. 企业; 2. 政府以及事业单位; 3. 其他组织, 4.团队号
	SubjectType int `json:"subject_type"  `
	// VerifiedEndTime 认证到期时间
	VerifiedEndTime int `json:"verified_end_time"  `
	// CorpWxQrCode 授权企业在微工作台（原企业号）的二维码，可用于关注微工作台
	CorpWxQrCode string `json:"corp_wxqrcode"  `
	// CorpScale 企业规模。当企业未设置该属性时，值为空
	CorpScale string `json:"corp_scale"  `
	// CorpIndustry 企业所属行业。当企业未设置该属性时，值为空
	CorpIndustry string `json:"corp_industry"  `
	// CorpSubIndustry 企业所属子行业。当企业未设置该属性时，值为空
	CorpSubIndustry string `json:"corp_sub_industry"  `
}

type ProviderGetPermanentCodeResponseAuthInfo struct {
	// Agent 授权的应用信息，注意是一个数组，但仅旧的多应用套件授权时会返回多个agent，对新的单应用授权，永远只返回一个agent
	Agent []ProviderGetPermanentCodeResponseAuthInfoAgent `json:"agent"  `
}

type ProviderGetPermanentCodeResponseAuthInfoAgent struct {
	// AgentID 授权方应用id
	AgentID int `json:"agentid"  `
	// Name 授权方应用名字
	Name string `json:"name"  `
	// SquareLogoURL 授权方应用方形头像
	SquareLogoURL string `json:"square_logo_url"  `
	// RoundLogoURL 授权方应用圆形头像
	RoundLogoURL string `json:"round_logo_url"  `
	// AppID 旧的多应用套件中的对应应用id，新开发者请忽略
	AppID int `json:"appid"  `
	// AuthMode 授权模式，0为管理员授权；1为成员授权
	AuthMode int `json:"auth_mode"  `
	// IsCustomizedApp 是否为代开发自建应用
	IsCustomizedApp bool `json:"is_customized_app"  `
	// Privilege 应用对应的权限
	Privilege ProviderGetPermanentCodeResponseAuthInfoAgentPrivilege `json:"privilege"  `
	// SharedFrom 共享了应用的互联企业信息，仅当由互联的企业共享应用触发的安装时才返回
	SharedFrom ProviderGetPermanentCodeResponseAuthInfoAgentSharedFrom `json:"shared_from"  `
}

type ProviderGetPermanentCodeResponseAuthInfoAgentPrivilege struct {
	// AllowParty 应用可见范围（部门）
	AllowParty []int `json:"allow_party"  `
	// AllowTag 应用可见范围（标签）
	AllowTag []int `json:"allow_tag"  `
	// AllowUser 应用可见范围（成员）
	AllowUser []string `json:"allow_user"  `
	// ExtraParty 额外通讯录（部门）
	ExtraParty []int `json:"extra_party"  `
	// ExtraUser 额外通讯录（成员）
	ExtraUser []string `json:"extra_user"  `
	// ExtraTag 额外通讯录（标签）
	ExtraTag []int `json:"extra_tag"  `
	// Level 权限等级。1:通讯录基本信息只读2:通讯录全部信息只读3:通讯录全部信息读写4:单个基本信息只读5:通讯录全部信息只写
	Level int `json:"level"  `
}

type ProviderGetPermanentCodeResponseAuthInfoAgentSharedFrom struct {
	// CorpID 共享了应用的互联企业信息，仅当由互联的企业共享应用触发的安装时才返回
	CorpID string `json:"corpid"  `
}

type ProviderGetPermanentCodeResponseAuthUserInfo struct {
	// UserID 授权管理员的userid，可能为空（企业互联由上级企业共享第三方应用给下级时，不返回授权的管理员信息）
	UserID string `json:"userid"  `
	// OpenUserID 授权管理员的open_userid，可能为空（企业互联由上级企业共享第三方应用给下级时，不返回授权的管理员信息）
	OpenUserID string `json:"open_userid"  `
	// Name 授权管理员的name，可能为空（企业互联由上级企业共享第三方应用给下级时，不返回授权的管理员信息）
	Name string `json:"name"  `
	// Avatar 授权管理员的头像url，可能为空（企业互联由上级企业共享第三方应用给下级时，不返回授权的管理员信息）
	Avatar string `json:"avatar"  `
}

type ProviderGetPermanentCodeResponseDealerCorpInfo struct {
	// CorpID 代理服务商企业微信id
	CorpID string `json:"corpid"  `
	// CorpName 代理服务商企业微信名称
	CorpName string `json:"corp_name"  `
}

type ProviderGetPermanentCodeResponseRegisterCodeInfo struct {
	// RegisterCode 注册码
	RegisterCode string `json:"register_code"  `
	// TemplateID 推广包ID
	TemplateID string `json:"template_id"  `
	// State 仅当获取注册码指定该字段时才返回
	State string `json:"state"  `
}

type ProviderGetAuthInfoRequest struct {
	// AuthCorpID 授权方corpid
	AuthCorpID string `json:"auth_corpid"  validate:"required"`
	// PermanentCode 永久授权码，通过get_permanent_code获取
	PermanentCode string `json:"permanent_code"  validate:"required"`
}

type ProviderGetAuthInfoResponse struct {
	// AuthCorpInfo 授权方企业信息
	AuthCorpInfo ProviderGetAuthInfoResponseAuthCorpInfo `json:"auth_corp_info"  `
	// AuthInfo 授权信息。如果是通讯录应用，且没开启实体应用，是没有该项的。通讯录应用拥有企业通讯录的全部信息读写权限
	AuthInfo ProviderGetAuthInfoResponseAuthInfo `json:"auth_info"  `
	// DealerCorpInfo 代理服务商企业信息
	DealerCorpInfo ProviderGetAuthInfoResponseDealerCorpInfo `json:"dealer_corp_info"  `
}

type ProviderGetAuthInfoResponseAuthCorpInfo struct {
	// CorpID 授权方企业微信id
	CorpID string `json:"corpid"  `
	// CorpName 授权方企业名称
	CorpName string `json:"corp_name"  `
	// CorpType 授权方企业类型，认证号：verified, 注册号：unverified
	CorpType string `json:"corp_type"  `
	// CorpSquareLogoURL 授权方企业方形头像
	CorpSquareLogoURL string `json:"corp_square_logo_url"  `
	// CorpUserMax 授权方企业用户规模
	CorpUserMax int `json:"corp_user_max"  `
	// CorpFullName 授权方企业的主体名称(仅认证或验证过的企业有)，即企业全称。
	CorpFullName string `json:"corp_full_name"  `
	// SubjectType 企业类型，1. 企业; 2. 政府以及事业单位; 3. 其他组织, 4.团队号
	SubjectType int `json:"subject_type"  `
	// VerifiedEndTime 认证到期时间
	VerifiedEndTime int `json:"verified_end_time"  `
	// CorpWxQrCode 授权企业在微工作台（原企业号）的二维码，可用于关注微工作台，二维码有效期为7天
	CorpWxQrCode string `json:"corp_wxqrcode"  `
	// CorpScale 企业规模。当企业未设置该属性时，值为空
	CorpScale string `json:"corp_scale"  `
	// CorpIndustry 企业所属行业。当企业未设置该属性时，值为空
	CorpIndustry string `json:"corp_industry"  `
	// CorpSubIndustry 企业所属子行业。当企业未设置该属性时，值为空
	CorpSubIndustry string `json:"corp_sub_industry"  `
}

type ProviderGetAuthInfoResponseAuthInfo struct {
	// Agent 授权的应用信息，注意是一个数组，但仅旧的多应用套件授权时会返回多个agent，对新的单应用授权，永远只返回一个agent
	Agent []ProviderGetAuthInfoResponseAuthInfoAgent `json:"agent"  `
}

type ProviderGetAuthInfoResponseAuthInfoAgent struct {
	// AgentID 授权方应用id
	AgentID int `json:"agentid"  `
	// Name 授权方应用名字
	Name string `json:"name"  `
	// SquareLogoURL 授权方应用方形头像
	SquareLogoURL string `json:"square_logo_url"  `
	// RoundLogoURL 授权方应用圆形头像
	RoundLogoURL string `json:"round_logo_url"  `
	// AppID 旧的多应用套件中的对应应用id，新开发者请忽略
	AppID int `json:"appid"  `
	// AuthMode 授权模式，0为管理员授权；1为成员授权
	AuthMode int `json:"auth_mode"  `
	// IsCustomizedApp 是否为代开发自建应用
	IsCustomizedApp bool `json:"is_customized_app"  `
	// Privilege 应用对应的权限
	Privilege ProviderGetAuthInfoResponseAuthInfoAgentPrivilege `json:"privilege"  `
	// SharedFrom 共享了应用的互联企业信息，仅当由互联的企业共享应用触发的安装时才返回
	SharedFrom ProviderGetAuthInfoResponseAuthInfoAgentSharedFrom `json:"shared_from"  `
}

type ProviderGetAuthInfoResponseAuthInfoAgentPrivilege struct {
	// AllowParty 应用可见范围（部门）
	AllowParty []int `json:"allow_party"  `
	// AllowTag 应用可见范围（标签）
	AllowTag []int `json:"allow_tag"  `
	// AllowUser 应用可见范围（成员）
	AllowUser []string `json:"allow_user"  `
	// ExtraParty 额外通讯录（部门）
	ExtraParty []int `json:"extra_party"  `
	// ExtraUser 额外通讯录（成员）
	ExtraUser []string `json:"extra_user"  `
	// ExtraTag 额外通讯录（标签）
	ExtraTag []int `json:"extra_tag"  `
	// Level 权限等级。1:通讯录基本信息只读2:通讯录全部信息只读（已废弃）3:通讯录全部信息读写4:单个基本信息只读5:通讯录全部信息只写（已废弃）
	Level int `json:"level"  `
}

type ProviderGetAuthInfoResponseAuthInfoAgentSharedFrom struct {
	// CorpID 共享了应用的互联企业信息，仅当由互联的企业共享应用触发的安装时才返回
	CorpID string `json:"corpid"  `
}

type ProviderGetAuthInfoResponseDealerCorpInfo struct {
	// CorpID 代理服务商企业微信id
	CorpID string `json:"corpid"  `
	// CorpName 代理服务商企业微信名称
	CorpName string `json:"corp_name"  `
}

type ProviderGetCorpTokenRequest struct {
	// AuthCorpID 授权方corpid
	AuthCorpID string `json:"auth_corpid"  validate:"required"`
	// PermanentCode 永久授权码，通过get_permanent_code获取
	PermanentCode string `json:"permanent_code"  validate:"required"`
}

type ProviderGetCorpTokenResponse struct {
	// AccessToken 授权方（企业）access_token,最长为512字节
	AccessToken string `json:"access_token"  `
	// ExpiresIn 授权方（企业）access_token超时时间
	ExpiresIn int `json:"expires_in"  `
}

type ProviderGetAdminListRequest struct {
	// AuthCorpID 授权方corpid
	AuthCorpID string `json:"auth_corpid"  validate:"required"`
	// AgentID 授权方安装的应用agentid
	AgentID int `json:"agentid"  validate:"required"`
}

type ProviderGetAdminListResponse struct {
	// Admin 应用的管理员列表（不包括外部管理员），成员授权模式下，不返回系统管理员
	Admin []GetAdminListItem `json:"admin"  `
}
