package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// AddGrid 添加网格
// 创建新的网格节点
//
// see https://developer.work.weixin.qq.com/document/path/94556
func (c *Client) AddGrid(r *AddGridRequest, opts ...any) (out AddGridResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/grid/add",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddGridRequest is request of Client.AddGrid
type AddGridRequest struct {
	// GridName 网格名称，不能超过30个字
	GridName string `json:"grid_name" validate:"required"`
	// GridParentID 父节点的id
	GridParentID string `json:"grid_parent_id" validate:"required"`
	// GridAdmin 网格负责人userid列表
	GridAdmin []string `json:"grid_admin" validate:"required"`
	// GridMember 该节点的成员userid列表
	GridMember []string `json:"grid_member"`
}

// AddGridResponse is response of Client.AddGrid
type AddGridResponse struct {
	// GridID 网格id
	GridID string `json:"grid_id"`
	// InvalidUserids 不合法的userid列表
	InvalidUserids []string `json:"invalid_userids"`
}

// AddEventCategory 添加事件类别
// 添加一级或二级事件类别，用于网格员在提交巡查上报或居民上报时填入配置的事件类别
//
// see https://developer.work.weixin.qq.com/document/path/94563
func (c *Client) AddEventCategory(r *AddEventCategoryRequest, opts ...any) (out AddEventCategoryResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/grid/add_cata",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddEventCategoryRequest is request of Client.AddEventCategory
type AddEventCategoryRequest struct {
	// CategoryName 分类名称，不能超过30个字
	CategoryName string `json:"category_name" validate:"required"`
	// Level 分类层级，只能传1或者2
	Level int `json:"level" validate:"required"`
	// ParentCategoryID 所属的一级分类的id，level为2时必传
	ParentCategoryID string `json:"parent_category_id"`
}

// AddEventCategoryResponse is response of Client.AddEventCategory
type AddEventCategoryResponse struct {
	// CategoryID 分类id
	CategoryID string `json:"category_id"`
}

// DeleteGrid 删除网格
// 删除指定的网格及其子网格
//
// see https://developer.work.weixin.qq.com/document/path/94559
func (c *Client) DeleteGrid(r *DeleteGridRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/grid/delete",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteGridRequest is request of Client.DeleteGrid
type DeleteGridRequest struct {
	// GridID 网格id，根节点可以填空
	GridID string `json:"grid_id" validate:"required"`
}

// DeleteReportCategory 删除事件类别
// 删除居民联系中的事件类别
//
// see https://developer.work.weixin.qq.com/document/path/94565
func (c *Client) DeleteReportCategory(r *DeleteReportCategoryRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/grid/delete_cata",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteReportCategoryRequest is request of Client.DeleteReportCategory
type DeleteReportCategoryRequest struct {
	// CategoryID 分类id
	CategoryID string `json:"category_id" validate:"required"`
}

// GetUserGridInfo 获取用户负责及参与的网格列表
// 获取指定成员负责及参与的网格列表信息
//
// see https://developer.work.weixin.qq.com/document/path/94567
func (c *Client) GetUserGridInfo(r *GetUserGridInfoRequest, opts ...any) (out GetUserGridInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/grid/get_user_grid_info",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetUserGridInfoRequest is request of Client.GetUserGridInfo
type GetUserGridInfoRequest struct {
	// UserID 需要查询的成员userid
	UserID string `json:"userid" validate:"required"`
}

// GetUserGridInfoResponse is response of Client.GetUserGridInfo
type GetUserGridInfoResponse struct {
	// ManageGrids 作为管理员管理的网格信息
	ManageGrids []map[string]any `json:"manage_grids"`
	// JoinedGrids 作为成员参与的grid信息
	JoinedGrids []map[string]any `json:"joined_grids"`
}

// ListGrid 获取网格列表
// 可以查询目标网格的父网格及全部子网格ID列表
//
// see https://developer.work.weixin.qq.com/document/path/94560
func (c *Client) ListGrid(r *ListGridRequest, opts ...any) (out ListGridResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/grid/list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListGridRequest is request of Client.ListGrid
type ListGridRequest struct {
	// GridID 网格id，不填则表示拉取根节点及其儿子节点
	GridID string `json:"grid_id"`
}

// ListGridResponse is response of Client.ListGrid
type ListGridResponse struct {
	// GridList 网格列表
	GridList []map[string]any `json:"grid_list"`
}

// ListEventCategory 获取事件类别列表
// 可以获取已配置的一级或二级事件类别，网格员在提交「巡查上报」或「居民上报」时将填入配置的事件类别
//
// see https://developer.work.weixin.qq.com/document/path/94566
func (c *Client) ListEventCategory(r *ListEventCategoryRequest, opts ...any) (out ListEventCategoryResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/grid/list_cata",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListEventCategoryRequest is request of Client.ListEventCategory
type ListEventCategoryRequest struct{}

// ListEventCategoryResponse is response of Client.ListEventCategory
type ListEventCategoryResponse struct {
	// CategoryList 分类列表
	CategoryList []map[string]any `json:"category_list"`
	// CategoryID 分类id
	CategoryID string `json:"category_id"`
	// CategoryName 分类名称
	CategoryName string `json:"category_name"`
	// Level 分类层级
	Level int `json:"level"`
	// ParentCategoryID 分类层级为1时，该字段为空
	ParentCategoryID string `json:"parent_category_id"`
}

// UpdateGrid 编辑网格
// 更新网格信息，包括名称、父节点、负责人及成员列表。
//
// see https://developer.work.weixin.qq.com/document/path/94558
func (c *Client) UpdateGrid(r *UpdateGridRequest, opts ...any) (out UpdateGridResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/grid/update",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateGridRequest is request of Client.UpdateGrid
type UpdateGridRequest struct {
	// GridID 网格id
	GridID string `json:"grid_id" validate:"required"`
	// GridName 网格名称，不能超过30个字，同一个目标网格下，不能存在同名的同级子网格
	GridName string `json:"grid_name" validate:"required"`
	// GridParentID 父节点的id，网格结构层级最多支持10层
	GridParentID string `json:"grid_parent_id" validate:"required"`
	// GridAdmin 网格「负责人」userid列表，每个网格至少1个，最多20个负责人
	GridAdmin []string `json:"grid_admin" validate:"required"`
	// GridMember 该节点的成员userid列表，不能超过100个，同一个成员最多能成功10个网格的管理员，为空则表示清空所有的成员
	GridMember []string `json:"grid_member"`
}

// UpdateGridResponse is response of Client.UpdateGrid
type UpdateGridResponse struct {
	// InvalidUserids 不合法的userid
	InvalidUserids []string `json:"invalid_userids"`
}

// UpdateCategory 修改事件类别
// 可以修改已配置的一级或二级事件类别，网格员在提交「巡查上报」或「居民上报」时将填入配置的事件类别
//
// see https://developer.work.weixin.qq.com/document/path/94564
func (c *Client) UpdateCategory(r *UpdateCategoryRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/grid/update_cata",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateCategoryRequest is request of Client.UpdateCategory
type UpdateCategoryRequest struct {
	// CategoryID 分类id
	CategoryID string `json:"category_id" validate:"required"`
	// CategoryName 分类名称，不能超过30个字，同一一级分类下的二级分类名字不能一样
	CategoryName string `json:"category_name" validate:"required"`
	// Level 分类层级，这里只能传1或者2
	Level int `json:"level" validate:"required"`
	// ParentCategoryID 所属的一级分类的id，level为2的话必传
	ParentCategoryID string `json:"parent_category_id"`
}

// GetPatrolCategoryStatistic 获取上报事件分类统计
// 获取上报事件的分类统计数据
//
// see https://developer.work.weixin.qq.com/document/path/93606
func (c *Client) GetPatrolCategoryStatistic(r *GetPatrolCategoryStatisticRequest, opts ...any) (out GetPatrolCategoryStatisticResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/patrol/category_statistic",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetPatrolCategoryStatisticRequest is request of Client.GetPatrolCategoryStatistic
type GetPatrolCategoryStatisticRequest struct {
	// CategoryID 分类ID，不传此字段，能拉取到所有一级分类的数据，然后传一级分类的category_id能拉取到该一级分类下的所有二级分类的数据
	CategoryID string `json:"category_id"`
}

// GetPatrolCategoryStatisticResponse is response of Client.GetPatrolCategoryStatistic
type GetPatrolCategoryStatisticResponse struct {
	// DashboardList 分类统计列表
	DashboardList []map[string]any `json:"dashboard_list"`
	// CategoryID 分类ID
	CategoryID string `json:"category_id"`
	// CategoryName 分类名称
	CategoryName string `json:"category_name"`
	// CategoryLevel 分类等级，1、一级分组；2、二级分类
	CategoryLevel int `json:"category_level"`
	// TotalCase 累计上报
	TotalCase int `json:"total_case"`
	// TotalSolved 累计办结
	TotalSolved int `json:"total_solved"`
	// CategoryType 分类类型，0：其他，1：有效分类
	CategoryType int `json:"category_type"`
}

// GetPatrolCorpStatus 获取单位巡查上报数据统计
// 获取企业或指定网格的巡查上报数据统计信息
//
// see https://developer.work.weixin.qq.com/document/path/93604
func (c *Client) GetPatrolCorpStatus(r *GetPatrolCorpStatusRequest, opts ...any) (out GetPatrolCorpStatusResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/patrol/get_corp_status",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetPatrolCorpStatusRequest is request of Client.GetPatrolCorpStatus
type GetPatrolCorpStatusRequest struct {
	// GridID 网格id，不传的话获取整个企业的概况
	GridID string `json:"grid_id"`
}

// GetPatrolCorpStatusResponse is response of Client.GetPatrolCorpStatus
type GetPatrolCorpStatusResponse struct {
	// ToBeAssigned 待分配数量
	ToBeAssigned int `json:"to_be_assigned"`
	// Processing 办理中数量
	Processing int `json:"processing"`
	// AddedToday 今日上报数量
	AddedToday int `json:"added_today"`
	// SolvedToday 今日办结数量
	SolvedToday int `json:"solved_today"`
	// TotalCase 累计上报数量
	TotalCase int `json:"total_case"`
	// TotalSolved 累计办结数量
	TotalSolved int `json:"total_solved"`
}

// GetGridInfo 获取配置的网格及网格负责人
// 获取企业已配置的网格信息及对应的网格负责人列表
//
// see https://developer.work.weixin.qq.com/document/path/93599
func (c *Client) GetGridInfo(r *GetGridInfoRequest, opts ...any) (out GetGridInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/report/patrol/get_grid_info",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetGridInfoRequest is request of Client.GetGridInfo
type GetGridInfoRequest struct{}

// GetGridInfoResponse is response of Client.GetGridInfo
type GetGridInfoResponse struct {
	// GridList 网格列表
	GridList []map[string]any `json:"grid_list"`
	// GridID 网格 ID
	GridID string `json:"grid_id"`
	// GridName 网格名称
	GridName int `json:"grid_name"`
	// GridAdmin 网格管理员 userId 列表
	GridAdmin []any `json:"grid_admin"`
}

// GetPatrolOrderInfo 获取巡查上报的事件详情信息
// 根据工单ID获取巡查上报事件的详细信息，包括事件描述、处理人、位置信息及处理流程等。
//
// see https://developer.work.weixin.qq.com/document/path/93608
func (c *Client) GetPatrolOrderInfo(r *GetPatrolOrderInfoRequest, opts ...any) (out GetPatrolOrderInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/patrol/get_order_info",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetPatrolOrderInfoRequest is request of Client.GetPatrolOrderInfo
type GetPatrolOrderInfoRequest struct {
	// OrderID 工单id，支持通过工单id直接查询某工单，此参数不为空的话，其他参数无效
	OrderID string `json:"order_id" validate:"required"`
}

// GetPatrolOrderInfoResponse is response of Client.GetPatrolOrderInfo
type GetPatrolOrderInfoResponse struct {
	// OrderInfo 工单信息对象
	OrderInfo any `json:"order_info"`
}

// ListPatrolOrder 获取巡查上报事件列表
// 获取企业内巡查上报的事件列表，支持分页查询和筛选
//
// see https://developer.work.weixin.qq.com/document/path/93607
func (c *Client) ListPatrolOrder(r *ListPatrolOrderRequest, opts ...any) (out ListPatrolOrderResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/patrol/get_order_list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListPatrolOrderRequest is request of Client.ListPatrolOrder
type ListPatrolOrderRequest struct {
	// BeginCreateTime 时间戳，筛选返回begin_create_time之后新创建的上报
	BeginCreateTime int `json:"begin_create_time"`
	// BeginModifyTime 时间戳，筛选返回begin_modify_time之后新修改的上报
	BeginModifyTime int `json:"begin_modify_time"`
	// Cursor 翻页参数，首次查询为空，如果查询条件有变更，需要将这个字段置空
	Cursor string `json:"cursor"`
	// Limit 单页数量，如果不填，默认20条，最大50
	Limit int `json:"limit"`
}

// ListPatrolOrderResponse is response of Client.ListPatrolOrder
type ListPatrolOrderResponse struct {
	// NextCursor 获取下一页数据的凭据，为空字符串代表是最后一页
	NextCursor string `json:"next_cursor"`
	// OrderList 事件列表
	OrderList []map[string]any `json:"order_list"`
}

// GetPatrolUserStatus 获取个人巡查上报数据统计
// 获取指定成员的巡查上报数据统计信息，包括办理中、今日上报和今日办结数量。
//
// see https://developer.work.weixin.qq.com/document/path/93605
func (c *Client) GetPatrolUserStatus(r *GetPatrolUserStatusRequest, opts ...any) (out GetPatrolUserStatusResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/patrol/get_user_status",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetPatrolUserStatusRequest is request of Client.GetPatrolUserStatus
type GetPatrolUserStatusRequest struct {
	// UserID 成员的userId
	UserID string `json:"userid" validate:"required"`
}

// GetPatrolUserStatusResponse is response of Client.GetPatrolUserStatus
type GetPatrolUserStatusResponse struct {
	// Processing 办理中数量
	Processing int `json:"processing"`
	// AddedToday 今日上报数量
	AddedToday int `json:"added_today"`
	// SolvedToday 今日办结数量
	SolvedToday int `json:"solved_today"`
}

// GetReportCategoryStatistic 获取上报事件分类统计
// 获取居民上报事件的分类统计数据
//
// see https://developer.work.weixin.qq.com/document/path/93612
func (c *Client) GetReportCategoryStatistic(r *GetReportCategoryStatisticRequest, opts ...any) (out GetReportCategoryStatisticResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/resident/category_statistic",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetReportCategoryStatisticRequest is request of Client.GetReportCategoryStatistic
type GetReportCategoryStatisticRequest struct {
	// CategoryID 分类ID，不传此字段拉取所有一级分类数据；传一级分类的category_id拉取其下的所有二级分类数据
	CategoryID string `json:"category_id"`
}

// GetReportCategoryStatisticResponse is response of Client.GetReportCategoryStatistic
type GetReportCategoryStatisticResponse struct {
	// DashboardList 分类列表
	DashboardList []map[string]any `json:"dashboard_list"`
}

// GetCorpResidentStatus 获取单位居民上报数据统计
// 获取单位居民上报数据的统计信息，包括办理中、今日上报、待受理等数据。
//
// see https://developer.work.weixin.qq.com/document/path/93610
func (c *Client) GetCorpResidentStatus(r *GetCorpResidentStatusRequest, opts ...any) (out GetCorpResidentStatusResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/resident/get_corp_status",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCorpResidentStatusRequest is request of Client.GetCorpResidentStatus
type GetCorpResidentStatusRequest struct {
	// GridID 网格id，不传的话获取整个企业的概况
	GridID string `json:"grid_id"`
}

// GetCorpResidentStatusResponse is response of Client.GetCorpResidentStatus
type GetCorpResidentStatusResponse struct {
	// Processing 办理中数据
	Processing int `json:"processing"`
	// AddedToday 今日上报数据
	AddedToday int `json:"added_today"`
	// SolvedToday 今日办结数据
	SolvedToday int `json:"solved_today"`
	// Pending 待受理数据
	Pending int `json:"pending"`
	// TotalCase 累计上报数据
	TotalCase int `json:"total_case"`
	// TotalAccepted 累计受理数据
	TotalAccepted int `json:"total_accepted"`
	// TotalSolved 累计办结数据
	TotalSolved int `json:"total_solved"`
}

// GetResidentOrderInfo 获取居民上报的事件详情信息
// 根据工单ID获取居民上报事件的详细信息，包括事件描述、处理人、地点及流程记录等。
//
// see https://developer.work.weixin.qq.com/document/path/93614
func (c *Client) GetResidentOrderInfo(r *GetResidentOrderInfoRequest, opts ...any) (out GetResidentOrderInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/resident/get_order_info",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetResidentOrderInfoRequest is request of Client.GetResidentOrderInfo
type GetResidentOrderInfoRequest struct {
	// OrderID 工单id，支持通过工单id直接查询某工单
	OrderID string `json:"order_id" validate:"required"`
}

// GetResidentOrderInfoResponse is response of Client.GetResidentOrderInfo
type GetResidentOrderInfoResponse struct {
	// OrderInfo 工单详细信息对象
	OrderInfo any `json:"order_info"`
}

// ListResidentOrder 获取居民上报事件列表
// 获取居民上报的事件列表，支持分页和筛选
//
// see https://developer.work.weixin.qq.com/document/path/93613
func (c *Client) ListResidentOrder(r *ListResidentOrderRequest, opts ...any) (out ListResidentOrderResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/resident/get_order_list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListResidentOrderRequest is request of Client.ListResidentOrder
type ListResidentOrderRequest struct {
	// BeginCreateTime 时间戳，筛选返回begin_create_time之后新创建的上报
	BeginCreateTime int `json:"begin_create_time"`
	// BeginModifyTime 时间戳，筛选返回begin_modify_time之后新修改的上报
	BeginModifyTime int `json:"begin_modify_time"`
	// Cursor 翻页参数，首次查询为空，如果查询条件有变更，需要将这个字段置空
	Cursor string `json:"cursor"`
	// Limit 单页数量，如果不填，默认20条，最大50
	Limit int `json:"limit"`
}

// ListResidentOrderResponse is response of Client.ListResidentOrder
type ListResidentOrderResponse struct {
	// NextCursor 获取下一页数据的凭据，为空字符串代表是最后一页
	NextCursor string `json:"next_cursor"`
	// OrderList 事件列表
	OrderList []map[string]any `json:"order_list"`
}

// GetResidentUserStatus 获取个人居民上报数据统计
// 获取指定成员的个人居民上报数据统计信息，包括办理中、今日上报、今日办结及待受理数量。
//
// see https://developer.work.weixin.qq.com/document/path/93611
func (c *Client) GetResidentUserStatus(r *GetResidentUserStatusRequest, opts ...any) (out GetResidentUserStatusResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/report/resident/get_user_status",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetResidentUserStatusRequest is request of Client.GetResidentUserStatus
type GetResidentUserStatusRequest struct {
	// UserID 成员的userId
	UserID string `json:"userid" validate:"required"`
}

// GetResidentUserStatusResponse is response of Client.GetResidentUserStatus
type GetResidentUserStatusResponse struct {
	// Processing 办理中数量
	Processing int `json:"processing"`
	// AddedToday 今日上报数量
	AddedToday int `json:"added_today"`
	// SolvedToday 今日办结数量
	SolvedToday int `json:"solved_today"`
	// Pending 待受理数量
	Pending int `json:"pending"`
}
