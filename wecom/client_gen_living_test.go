package wecom

func init() {
	registerClientAPIPath("/cgi-bin/living/cancel", "CancelLiving", cRef.CancelLiving)
	registerClientAPIPath("/cgi-bin/living/create", "CreateLiving", cRef.CreateLiving)
	registerClientAPIPath("/cgi-bin/living/get_living_code", "GetLivingCode", cRef.GetLivingCode)
	registerClientAPIPath("/cgi-bin/living/get_living_share_info", "GetLivingShareInfo", cRef.GetLivingShareInfo)
	registerClientAPIPath("/cgi-bin/living/modify", "ModifyLiving", cRef.ModifyLiving)
}
