package wecom

func init() {
	registerClientAPIPath("/cgi-bin/externalcontact/add_intercept_rule", "AddInterceptRule", cRef.AddInterceptRule)
	registerClientAPIPath("/cgi-bin/externalcontact/add_moment_task", "AddMomentTask", cRef.AddMomentTask)
	registerClientAPIPath("/cgi-bin/externalcontact/add_product_album", "AddProductAlbum", cRef.AddProductAlbum)
	registerClientAPIPath("/cgi-bin/externalcontact/cancel_groupmsg_send", "CancelGroupMsgSend", cRef.CancelGroupMsgSend)
	registerClientAPIPath("/cgi-bin/externalcontact/cancel_moment_task", "CancelMomentTask", cRef.CancelMomentTask)
	registerClientAPIPath("/cgi-bin/externalcontact/customer_acquisition_quota", "GetCustomerAcquisitionQuota", cRef.GetCustomerAcquisitionQuota)
	registerClientAPIPath("/cgi-bin/externalcontact/customer_acquisition/customer", "ListCustomerAcquisition", cRef.ListCustomerAcquisition)
	registerClientAPIPath("/cgi-bin/externalcontact/customer_acquisition/get_chat_info", "GetCustomerAcquisitionChatInfo", cRef.GetCustomerAcquisitionChatInfo)
	registerClientAPIPath("/cgi-bin/externalcontact/customer_acquisition/list_link", "ListCustomerAcquisitionLinks", cRef.ListCustomerAcquisitionLinks)
	registerClientAPIPath("/cgi-bin/externalcontact/customer_strategy/list", "ListCustomerStrategy", cRef.ListCustomerStrategy)
	registerClientAPIPath("/cgi-bin/externalcontact/get_strategy_tag_list", "GetStrategyTagList", cRef.GetStrategyTagList)
	registerClientAPIPath("/cgi-bin/externalcontact/groupchat/add_join_way", "AddJoinWay", cRef.AddJoinWay)
	registerClientAPIPath("/cgi-bin/externalcontact/groupchat/onjob_transfer", "TransferGroupChatOwner", cRef.TransferGroupChatOwner)
	registerClientAPIPath("/cgi-bin/externalcontact/moment_strategy/list", "ListMomentStrategy", cRef.ListMomentStrategy)
	registerClientAPIPath("/cgi-bin/externalcontact/opengid_to_chatid", "OpenGidToChatId", cRef.OpenGidToChatId)
	registerClientAPIPath("/cgi-bin/externalcontact/remind_groupmsg_send", "RemindGroupMsgSend", cRef.RemindGroupMsgSend)
	registerClientAPIPath("/cgi-bin/media/upload_attachment", "UploadAttachment", cRef.UploadAttachment)
}
