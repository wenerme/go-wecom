package wecom

func init() {
	registerClientAPIPath("/cgi-bin/dial/get_dial_record", "GetDialRecord", cRef.GetDialRecord)
}
