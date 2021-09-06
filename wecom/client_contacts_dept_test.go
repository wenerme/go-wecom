package wecom

func init() {
	registerClientAPIPath("/cgi-bin/department/create", "CreateDepartment", cRef.CreateDepartment)
	registerClientAPIPath("/cgi-bin/department/update", "UpdateDepartment", cRef.UpdateDepartment)
	registerClientAPIPath("/cgi-bin/department/delete", "DeleteDepartment", cRef.DeleteDepartment)
	registerClientAPIPath("/cgi-bin/department/list", "ListDepartment", cRef.ListDepartment)
}
