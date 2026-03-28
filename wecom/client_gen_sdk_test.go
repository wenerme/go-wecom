package wecom

func init() {
	registerClientAPIPath("/cgi-bin/user/getuserinfo", "GetUserInfoByCode", cRef.GetUserInfoByCode)
}
