package wecom

func init() {
	registerClientAPIPath("/cgi-bin/user/create", "CreateUser", cRef.CreateUser)
	registerClientAPIPath("/cgi-bin/user/update", "UpdateUser", cRef.UpdateUser)
	registerClientAPIPath("/cgi-bin/user/delete", "DeleteUser", cRef.DeleteUser)
	registerClientAPIPath("/cgi-bin/user/batchdelete", "BatchDeleteUser", cRef.BatchDeleteUser)
	registerClientAPIPath("/cgi-bin/user/simplelist", "SimpleListUser", cRef.SimpleListUser)
	registerClientAPIPath("/cgi-bin/user/list", "ListUser", cRef.ListUser)
	registerClientAPIPath("/cgi-bin/user/convert_to_openid", "ConvertToOpenID", cRef.ConvertToOpenID)
	registerClientAPIPath("/cgi-bin/user/authsucc", "AuthSuccess", cRef.AuthSuccess)
	registerClientAPIPath("/cgi-bin/batch/invite", "BatchInvite", cRef.BatchInvite)
	registerClientAPIPath("/cgi-bin/corp/get_join_qrcode", "GetJoinQrcode", cRef.GetJoinQrcode)
	registerClientAPIPath("/cgi-bin/user/get_active_stat", "GetActiveStat", cRef.GetActiveStat)
}
