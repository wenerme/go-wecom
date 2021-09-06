package wecom

func init() {
	registerClientAPIPath("/cgi-bin/linkedcorp/user/simplelist", "LinkSimpleListUser", cRef.LinkSimpleListUser)
	registerClientAPIPath("/cgi-bin/linkedcorp/user/list", "LinkListUser", cRef.LinkListUser)
	registerClientAPIPath("/cgi-bin/linkedcorp/department/list", "LinkListDepartment", cRef.LinkListDepartment)
	registerClientAPIPath("/cgi-bin/linkedcorp/user/get", "LinkGetUser", cRef.LinkGetUser)
	registerClientAPIPath("/cgi-bin/linkedcorp/agent/get_perm_list", "LinkGetPermList", cRef.LinkGetPermList)
}
