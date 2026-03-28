package wecom

func init() {
	registerClientAPIPath("/cgi-bin/corp/getapprovaldata", "GetApprovalData", cRef.GetApprovalData)
	registerClientAPIPath("/cgi-bin/corp/getopenapprovaldata", "GetOpenApprovalData", cRef.GetOpenApprovalData)
	registerClientAPIPath("/cgi-bin/oa/applyevent", "SubmitApprovalEvent", cRef.SubmitApprovalEvent)
	registerClientAPIPath("/cgi-bin/oa/approval/create_template", "CreateApprovalTemplate", cRef.CreateApprovalTemplate)
	registerClientAPIPath("/cgi-bin/oa/approval/update_template", "UpdateApprovalTemplate", cRef.UpdateApprovalTemplate)
	registerClientAPIPath("/cgi-bin/oa/getapprovaldetail", "GetApprovalDetail", cRef.GetApprovalDetail)
	registerClientAPIPath("/cgi-bin/oa/getapprovalinfo", "GetApprovalInfo", cRef.GetApprovalInfo)
	registerClientAPIPath("/cgi-bin/oa/gettemplatedetail", "GetApprovalTemplateDetail", cRef.GetApprovalTemplateDetail)
	registerClientAPIPath("/cgi-bin/oa/vacation/getcorpconf", "GetCorpVacationConfig", cRef.GetCorpVacationConfig)
	registerClientAPIPath("/cgi-bin/oa/vacation/getuservacationquota", "GetUserVacationQuota", cRef.GetUserVacationQuota)
	registerClientAPIPath("/cgi-bin/oa/vacation/setoneuserquota", "SetUserQuota", cRef.SetUserQuota)
}
