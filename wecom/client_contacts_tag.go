package wecom

import "github.com/wenerme/go-req"

// CreateTag 创建标签
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90210
func (c *Client) CreateTag(r *CreateTagRequest, opts ...interface{}) (out CreateTagResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/tag/create",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateTag 更新标签名字
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90211
func (c *Client) UpdateTag(r *UpdateTagRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/tag/update",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteTag 删除标签
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90212
func (c *Client) DeleteTag(r *DeleteTagRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/tag/delete",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetTag 获取标签成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90213
func (c *Client) GetTag(r *GetTagRequest, opts ...interface{}) (out GetTagResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/tag/get",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddTagUsers 增加标签成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90214
func (c *Client) AddTagUsers(r *AddTagUsersRequest, opts ...interface{}) (out AddTagUsersResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/tag/addtagusers",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteTagUsers 删除标签成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90215
func (c *Client) DeleteTagUsers(r *DeleteTagUsersRequest, opts ...interface{}) (out DeleteTagUsersResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/tag/deltagusers",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListTag 获取标签列表
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90216
func (c *Client) ListTag(opts ...interface{}) (out ListTagResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/tag/list",
		Options: opts,
	}).Fetch(&out)
	return
}

type CreateTagRequest struct {
	// TagName 标签名称，长度限制为32个字以内（汉字或英文字母），标签名不可与其他标签重名。
	TagName string `json:"tagname"  validate:"required"`
	// TagID 标签id，非负整型，指定此参数时新增的标签会生成对应的标签id，不指定时则以目前最大的id自增。
	TagID int `json:"tagid"  `
}

type CreateTagResponse struct {
	// TagID 标签id
	TagID string `json:"tagid"  `
}

type UpdateTagRequest struct {
	// TagID 标签ID
	TagID int `json:"tagid"  validate:"required"`
	// TagName 标签名称，长度限制为32个字（汉字或英文字母），标签不可与其他标签重名。
	TagName string `json:"tagname"  validate:"required"`
}

type DeleteTagRequest struct {
	// TagID 标签ID
	TagID string `json:"tagid"  validate:"required"`
}

type GetTagRequest struct {
	// TagID 标签ID
	TagID string `json:"tagid"  validate:"required"`
}

type GetTagResponse struct {
	// TagName 标签名
	TagName string `json:"tagname"  `
	// UserList 标签中包含的成员列表
	UserList []GetTagUser `json:"userlist"  `
	// PartyList 标签中包含的部门id列表
	PartyList []int `json:"partylist"  `
}

type AddTagUsersRequest struct {
	// TagID 标签ID
	TagID int `json:"tagid"  validate:"required"`
	// UserList 企业成员ID列表，注意：userlist、partylist不能同时为空，单次请求个数不超过1000
	UserList []string `json:"userlist"  `
	// PartyList 企业部门ID列表，注意：userlist、partylist不能同时为空，单次请求个数不超过100
	PartyList []int `json:"partylist"  `
}

type AddTagUsersResponse struct {
	// InvalidList 非法的成员帐号列表
	InvalidList string `json:"invalidlist"  `
	// InvalidParty 非法的部门id列表
	InvalidParty string `json:"invalidparty"  `
}

type DeleteTagUsersRequest struct {
	// TagID 标签ID
	TagID int `json:"tagid"  validate:"required"`
	// UserList 企业成员ID列表，注意：userlist、partylist不能同时为空，单次请求长度不超过1000
	UserList []string `json:"userlist"  `
	// PartyList 企业部门ID列表，注意：userlist、partylist不能同时为空，单次请求长度不超过100
	PartyList []int `json:"partylist"  `
}

type DeleteTagUsersResponse struct {
	// InvalidList 非法的成员帐号列表
	InvalidList string `json:"invalidlist"  `
	// InvalidParty 非法的部门id列表
	InvalidParty string `json:"invalidparty"  `
}

type ListTagResponse struct {
	// TagList 标签列表
	TagList []ListTag `json:"taglist"  `
}
