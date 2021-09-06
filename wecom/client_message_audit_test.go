package wecom

func init() {
	registerClientAPIPath("/cgi-bin/msgaudit/get_permit_user_list", "MessageAuditGetPermitUserList", cRef.MessageAuditGetPermitUserList)
	registerClientAPIPath("/cgi-bin/msgaudit/groupchat/get", "MessageAuditGetGroupChat", cRef.MessageAuditGetGroupChat)
}
