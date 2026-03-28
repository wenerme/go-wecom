package wecom

func init() {
	registerClientAPIPath("/cgi-bin/external-contact/getchatmsg", "GetChatMsg", cRef.GetChatMsg)
}
