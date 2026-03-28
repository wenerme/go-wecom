package wecom

func init() {
	registerClientAPIPath("/cgi-bin/security/admin_oper_log/list", "ListAdminOperLog", cRef.ListAdminOperLog)
	registerClientAPIPath("/cgi-bin/security/get_file_oper_record", "GetFileOperRecord", cRef.GetFileOperRecord)
	registerClientAPIPath("/cgi-bin/security/get_screen_oper_record", "GetScreenOperRecord", cRef.GetScreenOperRecord)
	registerClientAPIPath("/cgi-bin/security/get_server_domain_ip", "GetServerDomainIp", cRef.GetServerDomainIp)
	registerClientAPIPath("/cgi-bin/security/member_oper_log/list", "ListMemberOperLog", cRef.ListMemberOperLog)
	registerClientAPIPath("/cgi-bin/security/trustdevice/import", "ImportTrustDevice", cRef.ImportTrustDevice)
	registerClientAPIPath("/cgi-bin/security/vip/list", "ListSecurityVip", cRef.ListSecurityVip)
	registerClientAPIPath("/cgi-bin/security/vip/submit_batch_add_job", "SubmitBatchAddJob", cRef.SubmitBatchAddJob)
	registerClientAPIPath("/cgi-bin/security/vip/submit_batch_del_job", "BatchDelSecurityVip", cRef.BatchDelSecurityVip)
}
