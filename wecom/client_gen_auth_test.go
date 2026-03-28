package wecom

func init() {
	registerClientAPIPath("/cgi-bin/auth/get_tfa_info", "GetTfaInfo", cRef.GetTfaInfo)
	registerClientAPIPath("/cgi-bin/auth/getuserdetail", "GetAuthUserDetail", cRef.GetAuthUserDetail)
	registerClientAPIPath("/cgi-bin/user/tfa_succ", "SubmitTfaSuccess", cRef.SubmitTfaSuccess)
}
