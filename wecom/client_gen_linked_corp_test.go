package wecom

func init() {
	registerClientAPIPath("/cgi-bin/corpgroup/corp/gettoken", "GetCorpGroupToken", cRef.GetCorpGroupToken)
	registerClientAPIPath("/cgi-bin/miniprogram/transfer_session", "TransferMiniProgramSession", cRef.TransferMiniProgramSession)
}
