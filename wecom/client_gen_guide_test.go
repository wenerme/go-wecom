package wecom

func init() {
	registerClientAPIPath("/cgi-bin/get_api_domain_ip", "GetApiDomainIp", cRef.GetApiDomainIp)
	registerClientAPIPath("/cgi-bin/getcallbackip", "GetCallbackIp", cRef.GetCallbackIp)
}
