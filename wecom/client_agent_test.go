package wecom

func init() {
	registerClientAPIPath("/cgi-bin/agent/get", "GetAgent", cRef.GetAgent)
	registerClientAPIPath("/cgi-bin/agent/list", "ListAgent", cRef.ListAgent)
	registerClientAPIPath("/cgi-bin/agent/set_workbench_template", "SetWorkbenchTemplate", cRef.SetWorkbenchTemplate)
	registerClientAPIPath("/cgi-bin/agent/get_workbench_template", "GetWorkbenchTemplate", cRef.GetWorkbenchTemplate)
	registerClientAPIPath("/cgi-bin/agent/set_workbench_data", "SetWorkbenchData", cRef.SetWorkbenchData)
}
