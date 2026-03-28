package wecom

func init() {
	registerClientAPIPath("/cgi-bin/get_launch_code", "GetLaunchCode", cRef.GetLaunchCode)
	registerClientAPIPath("/cgi-bin/ticket/get", "GetInvoiceTicket", cRef.GetInvoiceTicket)
}
