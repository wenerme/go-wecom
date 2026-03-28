package wecom

func init() {
	registerClientAPIPath("/cgi-bin/card/invoice/reimburse/getinvoiceinfo", "GetInvoiceInfo", cRef.GetInvoiceInfo)
	registerClientAPIPath("/cgi-bin/card/invoice/reimburse/getinvoiceinfobatch", "BatchGetInvoiceInfo", cRef.BatchGetInvoiceInfo)
	registerClientAPIPath("/cgi-bin/card/invoice/reimburse/updateinvoicestatus", "UpdateInvoiceStatus", cRef.UpdateInvoiceStatus)
	registerClientAPIPath("/cgi-bin/card/invoice/reimburse/updatestatusbatch", "BatchUpdateInvoiceStatus", cRef.BatchUpdateInvoiceStatus)
}
