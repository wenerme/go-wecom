package wecom

func init() {
	registerClientAPIPath("/cgi-bin/export/simple_user", "ExportSimpleUser", cRef.ExportSimpleUser)
	registerClientAPIPath("/cgi-bin/export/user", "ExportUser", cRef.ExportUser)
	registerClientAPIPath("/cgi-bin/export/department", "ExportDepartment", cRef.ExportDepartment)
	registerClientAPIPath("/cgi-bin/export/taguser", "ExportTagUser", cRef.ExportTagUser)
	registerClientAPIPath("/cgi-bin/export/get_result", "ExportGetResult", cRef.ExportGetResult)
}
