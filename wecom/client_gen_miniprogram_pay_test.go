package wecom

func init() {
	registerClientAPIPath("/cgi-bin/miniapppay/close_order", "CloseMiniAppPayOrder", cRef.CloseMiniAppPayOrder)
	registerClientAPIPath("/cgi-bin/miniapppay/create_order", "CreateMiniAppOrder", cRef.CreateMiniAppOrder)
	registerClientAPIPath("/cgi-bin/miniapppay/get_order", "GetOrder", cRef.GetOrder)
	registerClientAPIPath("/cgi-bin/miniapppay/get_refund_detail", "GetRefundDetail", cRef.GetRefundDetail)
	registerClientAPIPath("/cgi-bin/miniapppay/get_sign", "GetMiniAppPaySign", cRef.GetMiniAppPaySign)
	registerClientAPIPath("/cgi-bin/miniapppay/refund", "Refund", cRef.Refund)
}
