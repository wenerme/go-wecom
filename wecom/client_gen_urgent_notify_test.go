package wecom

func init() {
	registerClientAPIPath("/cgi-bin/pstncc/call", "CallPstnCc", cRef.CallPstnCc)
	registerClientAPIPath("/cgi-bin/pstncc/getstates", "GetPstnccState", cRef.GetPstnccState)
}
