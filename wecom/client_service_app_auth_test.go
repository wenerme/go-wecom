package wecom

func init() {
	registerClientAPIPath("/cgi-bin/service/get_suite_token", "ProviderGetSuiteToken", cRef.ProviderGetSuiteToken)
	registerClientAPIPath("/cgi-bin/service/get_pre_auth_code", "ProviderGetPreAuthCode", cRef.ProviderGetPreAuthCode)
	registerClientAPIPath("/cgi-bin/service/set_session_info", "ProviderSetSessionInfo", cRef.ProviderSetSessionInfo)
	registerClientAPIPath("/cgi-bin/service/get_permanent_code", "ProviderGetPermanentCode", cRef.ProviderGetPermanentCode)
	registerClientAPIPath("/cgi-bin/service/get_auth_info", "ProviderGetAuthInfo", cRef.ProviderGetAuthInfo)
	registerClientAPIPath("/cgi-bin/service/get_corp_token", "ProviderGetCorpToken", cRef.ProviderGetCorpToken)
	registerClientAPIPath("/cgi-bin/service/get_admin_list", "ProviderGetAdminList", cRef.ProviderGetAdminList)
}
