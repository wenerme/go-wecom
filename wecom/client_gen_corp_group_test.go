package wecom

func init() {
	registerClientAPIPath("/cgi-bin/corpgroup/corp/get_chain_list", "GetChainList", cRef.GetChainList)
	registerClientAPIPath("/cgi-bin/corpgroup/corp/get_chain_user_custom_id", "GetChainUserCustomId", cRef.GetChainUserCustomId)
	registerClientAPIPath("/cgi-bin/corpgroup/corp/list_app_share_info", "ListCorpGroupAppShare", cRef.ListCorpGroupAppShare)
	registerClientAPIPath("/cgi-bin/corpgroup/corp/remove_corp", "RemoveCorpFromGroup", cRef.RemoveCorpFromGroup)
	registerClientAPIPath("/cgi-bin/corpgroup/get_corp_shared_chain_list", "GetCorpSharedChainList", cRef.GetCorpSharedChainList)
	registerClientAPIPath("/cgi-bin/corpgroup/getresult", "GetAsyncJobResult", cRef.GetAsyncJobResult)
	registerClientAPIPath("/cgi-bin/corpgroup/import_chain_contact", "ImportChainContact", cRef.ImportChainContact)
	registerClientAPIPath("/cgi-bin/corpgroup/rule/add_rule", "AddCorpGroupRule", cRef.AddCorpGroupRule)
	registerClientAPIPath("/cgi-bin/corpgroup/rule/delete_rule", "DeleteCorpGroupRule", cRef.DeleteCorpGroupRule)
	registerClientAPIPath("/cgi-bin/corpgroup/rule/get_rule_info", "GetRuleInfo", cRef.GetRuleInfo)
	registerClientAPIPath("/cgi-bin/corpgroup/rule/list_ids", "ListRuleId", cRef.ListRuleId)
	registerClientAPIPath("/cgi-bin/corpgroup/rule/modify_rule", "ModifyRule", cRef.ModifyRule)
	registerClientAPIPath("/cgi-bin/corpgroup/unionid_to_external_userid", "LinkUnionidToExternalUserid", cRef.LinkUnionidToExternalUserid)
	registerClientAPIPath("/cgi-bin/corpgroup/unionid_to_pending_id", "UnionidToPendingId", cRef.UnionidToPendingId)
}
