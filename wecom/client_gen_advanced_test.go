package wecom

func init() {
	registerClientAPIPath("/cgi-bin/advanced_feature/get_apply_id_list", "ListApplyId", cRef.ListApplyId)
	registerClientAPIPath("/cgi-bin/advanced_feature/set_approval_detail", "SetApprovalDetail", cRef.SetApprovalDetail)
}
