package wecom

func init() {
	registerClientAPIPath("/cgi-bin/department/get", "GetDepartment", cRef.GetDepartment)
	registerClientAPIPath("/cgi-bin/department/simplelist", "SimpleListDepartment", cRef.SimpleListDepartment)
	registerClientAPIPath("/cgi-bin/user/get_userid_by_email", "GetUserIdByEmail", cRef.GetUserIdByEmail)
	registerClientAPIPath("/cgi-bin/user/getuserid", "GetUserIDByMobile", cRef.GetUserIDByMobile)
	registerClientAPIPath("/cgi-bin/user/list_id", "ListUserId", cRef.ListUserId)
}
