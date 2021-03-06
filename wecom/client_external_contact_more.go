package wecom

// ExternalContactResponse 外部联系人
// see https://work.weixin.qq.com/api/doc/90001/90143/92265
type ExternalContactResponse struct {
	// ExternalUserID 外部联系人的userid
	ExternalUserID string `json:"external_userid"`
	// Name 外部联系人的名称[注1]
	Name string `json:"name"`
	// Avatar 外部联系人头像，代开发自建应用需要管理员授权才可以获取，第三方不可获取
	Avatar string `json:"avatar"`
	// Type 外部联系人的类型，1表示该外部联系人是微信用户，2表示该外部联系人是企业微信用户
	Type int `json:"type"`
	// Gender 外部联系人性别 0-未知 1-男性 2-女性
	Gender int `json:"gender"`
	// UnionID 外部联系人在微信开放平台的唯一身份标识（微信unionid），通过此字段企业可将外部联系人与公众号/小程序用户关联起来。仅当联系人类型是微信用户，且企业或第三方服务商绑定了微信开发者ID有此字段。查看绑定方法
	UnionID string `json:"unionid"`
	// Position 外部联系人的职位，如果外部企业或用户选择隐藏职位，则不返回，仅当联系人类型是企业微信用户时有此字段
	Position string `json:"position"`
	// CorpName 外部联系人所在企业的简称，仅当联系人类型是企业微信用户时有此字段
	CorpName string `json:"corp_name"`
	// CorpFullName 外部联系人所在企业的主体名称，仅当联系人类型是企业微信用户时有此字段
	CorpFullName string `json:"corp_full_name"`
	// ExternalProfile 外部联系人的自定义展示信息，可以有多个字段和多种类型，包括文本，网页和小程序，仅当联系人类型是企业微信用户时有此字段，字段详情见对外属性；
	ExternalProfile ExternalProfile `json:"external_profile"`
}

type ExternalContactProfileResponse struct {
	ExternalAttr []ExtAttr `json:"external_attr"`
}

// Conclusions 结束语
// https://work.weixin.qq.com/api/doc/90001/90143/92577#%E7%BB%93%E6%9D%9F%E8%AF%AD%E5%AE%9A%E4%B9%89
type Conclusions struct {
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
	Image struct {
		MediaID string `json:"media_id"`
		PicURL  string `json:"pic_url"`
	} `json:"image"`
	Link struct {
		Title  string `json:"title"`
		PicURL string `json:"picurl"`
		Desc   string `json:"desc"`
		URL    string `json:"url"`
	} `json:"link"`
	MiniProgram struct {
		Title      string `json:"title"`
		PicMediaID string `json:"pic_media_id"`
		AppID      string `json:"appid"`
		Page       string `json:"page"`
	} `json:"miniprogram"`
}
