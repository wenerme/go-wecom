package wecom

import "github.com/wenerme/go-req"

// LinkSimpleListUser 获取互联企业部门成员
//
// see https://work.weixin.qq.com/api/doc/90000/90135/93168
func (c *Client) LinkSimpleListUser(r *LinkSimpleListUserRequest) (out LinkSimpleListUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/linkedcorp/user/simplelist",
		Body:   r,
	}).Fetch(&out)
	return
}

// LinkListUser 获取互联企业部门成员详情
//
// see https://work.weixin.qq.com/api/doc/90000/90135/93169
func (c *Client) LinkListUser(r *LinkListUserRequest) (out LinkListUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/linkedcorp/user/list",
		Body:   r,
	}).Fetch(&out)
	return
}

// LinkListDepartment 获取互联企业部门列表
//
// see https://work.weixin.qq.com/api/doc/90000/90135/93170
func (c *Client) LinkListDepartment(r *LinkListDepartmentRequest) (out LinkListDepartmentResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/linkedcorp/department/list",
		Body:   r,
	}).Fetch(&out)
	return
}

// LinkGetUser 获取互联企业成员详细信息
//
// see https://work.weixin.qq.com/api/doc/90000/90135/93171
func (c *Client) LinkGetUser(r *LinkGetUserRequest) (out LinkGetUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/linkedcorp/user/get",
		Body:   r,
	}).Fetch(&out)
	return
}

// LinkGetPermList 获取应用的可见范围
// 本接口只返回互联企业中非本企业内的成员和部门的信息，如果要获取本企业的可见范围，请调用“获取应用”接口
//
// see https://work.weixin.qq.com/api/doc/90000/90135/93172
func (c *Client) LinkGetPermList() (out LinkGetPermListResponse, err error) {
	err = c.Request.With(req.Request{
		Method: "POST",
		URL:    "/cgi-bin/linkedcorp/agent/get_perm_list",
	}).Fetch(&out)
	return
}

type LinkSimpleListUserRequest struct {
	// DepartmentID 该字段用的是互联应用可见范围接口返回的department_ids参数，用的是 linkedid + ’/‘ + department_id 拼成的字符串
	DepartmentID string `json:"department_id" validate:"required"`
	// FetchChild 是否递归获取子部门下面的成员：1-递归获取，0-只获取本部门，不传默认只获取本部门成员
	FetchChild bool `json:"fetch_child"`
}

type LinkSimpleListUserResponse struct {
	// UserList 成员列表
	UserList []LinkSimpleListUser `json:"userlist"`
}

type LinkListUserRequest struct {
	// DepartmentID 该字段用的是互联应用可见范围接口返回的department_ids参数，用的是 linkedid + ’/‘ + department_id 拼成的字符串
	DepartmentID string `json:"department_id" validate:"required"`
	// FetchChild 是否递归获取子部门下面的成员：1-递归获取，0-只获取本部门，不传默认只获取本部门成员
	FetchChild bool `json:"fetch_child"`
}

type LinkListUserResponse struct {
	// UserList 成员列表，user包含的属性可在管理端配置
	UserList []LinkListUser `json:"userlist"`
}

type LinkListDepartmentRequest struct {
	// DepartmentID 该字段用的是互联应用可见范围接口返回的department_ids参数，用的是 linkedid + ’/‘ + department_id 拼成的字符串
	DepartmentID string `json:"department_id" validate:"required"`
}

type LinkListDepartmentResponse struct {
	// DepartmentList 部门列表
	DepartmentList []LinkListDepartment `json:"department_list"`
}

type LinkGetUserRequest struct {
	// UserID 该字段用的是互联应用可见范围接口返回的userids参数，用的是 CorpId + ’/‘ + USERID 拼成的字符串
	UserID string `json:"userid" validate:"required"`
}

type LinkGetUserResponse struct {
	// UserInfo 成员的详细信息，user包含的属性可在管理端配置
	UserInfo LinkGetUser `json:"user_info"`
}

type LinkGetPermListResponse struct {
	// UserIds 可见的userids，是用 CorpId + ’/‘ + USERID 拼成的字符串
	UserIds []string `json:"userids"`
	// DepartmentIds 可见的department_ids，是用 linkedid + ’/‘ + department_id 拼成的字符串
	DepartmentIds []string `json:"department_ids"`
}
