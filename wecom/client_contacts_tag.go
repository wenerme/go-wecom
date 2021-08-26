package wecom

import "github.com/wenerme/go-req"

// CreateTag 创建标签
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90210
func (c *Client) CreateTag(r *CreateTagRequest) (out CreateTagResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/tag/create",
		Body:   r,
	}).Fetch(&out)
	return
}

// UpdateTag 更新标签名字
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90211
func (c *Client) UpdateTag(r *UpdateTagRequest) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/tag/update",
		Body:   r,
	}).Fetch(&out)
	return
}

// DeleteTag 删除标签
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90212
func (c *Client) DeleteTag(r *DeleteTagRequest) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "GET",
		URL:    "/cgi-bin/tag/delete",
		Query:  r,
	}).Fetch(&out)
	return
}

// GetTag 获取标签成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90213
func (c *Client) GetTag(r *GetTagRequest) (out GetTagResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "GET",
		URL:    "/cgi-bin/tag/get",
		Query:  r,
	}).Fetch(&out)
	return
}

// Addtagusers 增加标签成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90214
func (c *Client) Addtagusers(r *AddtagusersRequest) (out AddtagusersResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/tag/addtagusers",
		Body:   r,
	}).Fetch(&out)
	return
}

// Deltagusers 删除标签成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90215
func (c *Client) Deltagusers(r *DeltagusersRequest) (out DeltagusersResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/tag/deltagusers",
		Body:   r,
	}).Fetch(&out)
	return
}

// ListTag 获取标签列表
//
// see https://work.weixin.qq.com/api/doc/90000/90135/90216
func (c *Client) ListTag() (out ListTagResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "GET",
		URL:    "/cgi-bin/tag/list",
	}).Fetch(&out)
	return
}

type CreateTagRequest struct {
	// TagName 标签名称，长度限制为32个字以内（汉字或英文字母），标签名不可与其他标签重名。
	TagName string `json:"tagname"`
	// TagID 标签id，非负整型，指定此参数时新增的标签会生成对应的标签id，不指定时则以目前最大的id自增。
	TagID int `json:"tagid"`
}

type CreateTagResponse struct {
	// TagID 标签id
	TagID string `json:"tagid"`
}

type UpdateTagRequest struct {
	// TagID 标签ID
	TagID int `json:"tagid"`
	// TagName 标签名称，长度限制为32个字（汉字或英文字母），标签不可与其他标签重名。
	TagName string `json:"tagname"`
}

type DeleteTagRequest struct {
	// TagID 标签ID
	TagID string `json:"tagid"`
}

type GetTagRequest struct {
	// TagID 标签ID
	TagID string `json:"tagid"`
}

type GetTagResponse struct {
	// TagName 标签名
	TagName string `json:"tagname"`
	// UserList 标签中包含的成员列表
	UserList []GetTagUser `json:"userlist"`
	// PartyList 标签中包含的部门id列表
	PartyList []int `json:"partylist"`
}

type AddtagusersRequest struct {
	// TagID 标签ID
	TagID int `json:"tagid"`
	// UserList 企业成员ID列表，注意：userlist、partylist不能同时为空，单次请求个数不超过1000
	UserList []string `json:"userlist"`
	// PartyList 企业部门ID列表，注意：userlist、partylist不能同时为空，单次请求个数不超过100
	PartyList []int `json:"partylist"`
}

type AddtagusersResponse struct {
	// InvalidList 非法的成员帐号列表
	InvalidList string `json:"invalidlist"`
	// InvalidParty 非法的部门id列表
	InvalidParty string `json:"invalidparty"`
}

type DeltagusersRequest struct {
	// TagID 标签ID
	TagID int `json:"tagid"`
	// UserList 企业成员ID列表，注意：userlist、partylist不能同时为空，单次请求长度不超过1000
	UserList []string `json:"userlist"`
	// PartyList 企业部门ID列表，注意：userlist、partylist不能同时为空，单次请求长度不超过100
	PartyList []int `json:"partylist"`
}

type DeltagusersResponse struct {
	// InvalidList 非法的成员帐号列表
	InvalidList string `json:"invalidlist"`
	// InvalidParty 非法的部门id列表
	InvalidParty string `json:"invalidparty"`
}

type ListTagResponse struct {
	// TagList 标签列表
	TagList []ListTag `json:"taglist"`
}
