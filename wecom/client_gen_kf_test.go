package wecom

func init() {
	registerClientAPIPath("/cgi-bin/kf/account/add", "AddKfAccount", cRef.AddKfAccount)
	registerClientAPIPath("/cgi-bin/kf/account/del", "DeleteKfAccount", cRef.DeleteKfAccount)
	registerClientAPIPath("/cgi-bin/kf/account/list", "ListKfAccount", cRef.ListKfAccount)
	registerClientAPIPath("/cgi-bin/kf/account/update", "UpdateKfAccount", cRef.UpdateKfAccount)
	registerClientAPIPath("/cgi-bin/kf/customer/batchget", "BatchGetCustomer", cRef.BatchGetCustomer)
	registerClientAPIPath("/cgi-bin/kf/customer/get_upgrade_service_config", "GetUpgradeServiceConfig", cRef.GetUpgradeServiceConfig)
	registerClientAPIPath("/cgi-bin/kf/get_corp_statistic", "GetCorpStatistic", cRef.GetCorpStatistic)
	registerClientAPIPath("/cgi-bin/kf/get_servicer_statistic", "GetServicerStatistic", cRef.GetServicerStatistic)
	registerClientAPIPath("/cgi-bin/kf/knowledge/add_group", "AddKnowledgeGroup", cRef.AddKnowledgeGroup)
	registerClientAPIPath("/cgi-bin/kf/knowledge/add_intent", "AddKnowledgeIntent", cRef.AddKnowledgeIntent)
	registerClientAPIPath("/cgi-bin/kf/send_msg", "SendKfMsg", cRef.SendKfMsg)
	registerClientAPIPath("/cgi-bin/kf/send_msg_on_event", "SendKfEventMsg", cRef.SendKfEventMsg)
	registerClientAPIPath("/cgi-bin/kf/service_state/trans", "TransferServiceState", cRef.TransferServiceState)
	registerClientAPIPath("/cgi-bin/kf/servicer/add", "AddServicer", cRef.AddServicer)
	registerClientAPIPath("/cgi-bin/kf/servicer/del", "DeleteServiceUser", cRef.DeleteServiceUser)
	registerClientAPIPath("/cgi-bin/kf/servicer/list", "ListServicer", cRef.ListServicer)
	registerClientAPIPath("/cgi-bin/kf/sync_msg", "SyncMsg", cRef.SyncMsg)
}
