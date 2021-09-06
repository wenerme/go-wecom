package wecom

func init() {
	registerClientAPIPath("/cgi-bin/msgaudit/check_room_agree", "MessageAuditCheckRoomAgree", cRef.MessageAuditCheckRoomAgree)
	registerClientAPIPath("/cgi-bin/msgaudit/check_single_agree", "MessageAuditCheckSingleAgree", cRef.MessageAuditCheckSingleAgree)
}
