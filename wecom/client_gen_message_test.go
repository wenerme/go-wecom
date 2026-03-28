package wecom

func init() {
	registerClientAPIPath("/cgi-bin/appchat/create", "CreateGroupChat", cRef.CreateGroupChat)
	registerClientAPIPath("/cgi-bin/appchat/get", "GetAppChat", cRef.GetAppChat)
	registerClientAPIPath("/cgi-bin/appchat/send", "SendAppChatMessage", cRef.SendAppChatMessage)
	registerClientAPIPath("/cgi-bin/appchat/update", "UpdateAppChat", cRef.UpdateAppChat)
	registerClientAPIPath("/cgi-bin/externalcontact/getjoinqrcode", "VerifyJoinQrCodeUrl", cRef.VerifyJoinQrCodeUrl)
	registerClientAPIPath("/cgi-bin/externalcontact/message/send", "SendExternalContactMessage", cRef.SendExternalContactMessage)
	registerClientAPIPath("/cgi-bin/message/recall", "RecallAppMessage", cRef.RecallAppMessage)
	registerClientAPIPath("/cgi-bin/message/send", "SendMessage", cRef.SendMessage)
	registerClientAPIPath("/cgi-bin/message/update_template_card", "UpdateTemplateCard", cRef.UpdateTemplateCard)
	registerClientAPIPath("/cgi-bin/wedoc/smartsheet/groupchat/update", "UpdateGroupChat", cRef.UpdateGroupChat)
}
