package wecom

func init() {
	registerClientAPIPath("/cgi-bin/report/grid/add", "AddGrid", cRef.AddGrid)
	registerClientAPIPath("/cgi-bin/report/grid/add_cata", "AddEventCategory", cRef.AddEventCategory)
	registerClientAPIPath("/cgi-bin/report/grid/delete", "DeleteGrid", cRef.DeleteGrid)
	registerClientAPIPath("/cgi-bin/report/grid/delete_cata", "DeleteReportCategory", cRef.DeleteReportCategory)
	registerClientAPIPath("/cgi-bin/report/grid/get_user_grid_info", "GetUserGridInfo", cRef.GetUserGridInfo)
	registerClientAPIPath("/cgi-bin/report/grid/list", "ListGrid", cRef.ListGrid)
	registerClientAPIPath("/cgi-bin/report/grid/list_cata", "ListEventCategory", cRef.ListEventCategory)
	registerClientAPIPath("/cgi-bin/report/grid/update", "UpdateGrid", cRef.UpdateGrid)
	registerClientAPIPath("/cgi-bin/report/grid/update_cata", "UpdateCategory", cRef.UpdateCategory)
	registerClientAPIPath("/cgi-bin/report/patrol/category_statistic", "GetPatrolCategoryStatistic", cRef.GetPatrolCategoryStatistic)
	registerClientAPIPath("/cgi-bin/report/patrol/get_corp_status", "GetPatrolCorpStatus", cRef.GetPatrolCorpStatus)
	registerClientAPIPath("/cgi-bin/report/patrol/get_grid_info", "GetGridInfo", cRef.GetGridInfo)
	registerClientAPIPath("/cgi-bin/report/patrol/get_order_info", "GetPatrolOrderInfo", cRef.GetPatrolOrderInfo)
	registerClientAPIPath("/cgi-bin/report/patrol/get_order_list", "ListPatrolOrder", cRef.ListPatrolOrder)
	registerClientAPIPath("/cgi-bin/report/patrol/get_user_status", "GetPatrolUserStatus", cRef.GetPatrolUserStatus)
	registerClientAPIPath("/cgi-bin/report/resident/category_statistic", "GetReportCategoryStatistic", cRef.GetReportCategoryStatistic)
	registerClientAPIPath("/cgi-bin/report/resident/get_corp_status", "GetCorpResidentStatus", cRef.GetCorpResidentStatus)
	registerClientAPIPath("/cgi-bin/report/resident/get_order_info", "GetResidentOrderInfo", cRef.GetResidentOrderInfo)
	registerClientAPIPath("/cgi-bin/report/resident/get_order_list", "ListResidentOrder", cRef.ListResidentOrder)
	registerClientAPIPath("/cgi-bin/report/resident/get_user_status", "GetResidentUserStatus", cRef.GetResidentUserStatus)
}
