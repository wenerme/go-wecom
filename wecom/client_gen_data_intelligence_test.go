package wecom

func init() {
	registerClientAPIPath("/cgi-bin/chatdata/async_program_task", "CreateProgramTask", cRef.CreateProgramTask)
	registerClientAPIPath("/cgi-bin/chatdata/check_debug_mode", "CheckDebugMode", cRef.CheckDebugMode)
	registerClientAPIPath("/cgi-bin/chatdata/close_debug_mode", "CloseDebugMode", cRef.CloseDebugMode)
	registerClientAPIPath("/cgi-bin/chatdata/get_auth_user_list", "GetAuthUserList", cRef.GetAuthUserList)
	registerClientAPIPath("/cgi-bin/chatdata/open_debug_mode", "OpenChatDataDebugMode", cRef.OpenChatDataDebugMode)
	registerClientAPIPath("/cgi-bin/chatdata/set_hide_sensitiveinfo_config", "SetChatdataHideSensitiveInfoConfig", cRef.SetChatdataHideSensitiveInfoConfig)
	registerClientAPIPath("/cgi-bin/chatdata/set_log_level", "SetLogLevel", cRef.SetLogLevel)
	registerClientAPIPath("/cgi-bin/chatdata/set_public_key", "SetPublicKey", cRef.SetPublicKey)
	registerClientAPIPath("/cgi-bin/chatdata/set_receive_callback", "SetReceiveCallback", cRef.SetReceiveCallback)
	registerClientAPIPath("/cgi-bin/chatdata/sync_call_program", "SyncCallProgram", cRef.SyncCallProgram)
	registerClientAPIPath("/cgi-bin/chatdata/upload_media", "UploadChatDataMedia", cRef.UploadChatDataMedia)
}
