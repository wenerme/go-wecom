package wecom

func init() {
	registerClientAPIPath("/cgi-bin/externalpay/get_bill_list", "ListExternalPayBill", cRef.ListExternalPayBill)
	registerClientAPIPath("/cgi-bin/externalpay/get_fund_flow", "GetFundFlow", cRef.GetFundFlow)
	registerClientAPIPath("/cgi-bin/externalpay/get_payment_info", "GetPaymentInfo", cRef.GetPaymentInfo)
	registerClientAPIPath("/cgi-bin/externalpay/getmerchant", "GetMerchantDetail", cRef.GetMerchantDetail)
	registerClientAPIPath("/cgi-bin/miniapppay/apply_mch", "ApplyMch", cRef.ApplyMch)
	registerClientAPIPath("/cgi-bin/miniapppay/get_applyment_status", "GetApplymentStatus", cRef.GetApplymentStatus)
	registerClientAPIPath("/cgi-bin/miniapppay/upload_image", "UploadMiniAppImage", cRef.UploadMiniAppImage)
	registerClientAPIPath("/mmpaymkttransfers/promotion/paywwsptrans2pocket", "PayEmployee", cRef.PayEmployee)
	registerClientAPIPath("/mmpaymkttransfers/promotion/querywwsptrans2pocket", "QueryWxPaymentRecord", cRef.QueryWxPaymentRecord)
	registerClientAPIPath("/mmpaymkttransfers/queryworkwxredpack", "QueryWorkWxRedpack", cRef.QueryWorkWxRedpack)
	registerClientAPIPath("/mmpaymkttransfers/sendworkwxredpack", "SendWorkWxRedPacket", cRef.SendWorkWxRedPacket)
}
