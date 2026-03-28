package wecom

func init() {
	registerClientAPIPath("/cgi-bin/health/get_health_report_stat", "GetHealthReportStat", cRef.GetHealthReportStat)
	registerClientAPIPath("/cgi-bin/health/get_report_answer", "GetReportAnswer", cRef.GetReportAnswer)
	registerClientAPIPath("/cgi-bin/health/get_report_job_info", "GetHealthReportJobInfo", cRef.GetHealthReportJobInfo)
	registerClientAPIPath("/cgi-bin/health/get_report_jobids", "GetHealthReportJobIds", cRef.GetHealthReportJobIds)
	registerClientAPIPath("/cgi-bin/living/delete_replay_data", "DeleteLivingReplayData", cRef.DeleteLivingReplayData)
	registerClientAPIPath("/cgi-bin/living/get_user_all_livingid", "ListUserLivingId", cRef.ListUserLivingId)
	registerClientAPIPath("/cgi-bin/school/get_payment_result", "GetPaymentResult", cRef.GetPaymentResult)
	registerClientAPIPath("/cgi-bin/school/get_trade", "GetTradeDetail", cRef.GetTradeDetail)
	registerClientAPIPath("/cgi-bin/school/living/get_living_info", "GetLivingInfo", cRef.GetLivingInfo)
	registerClientAPIPath("/cgi-bin/school/living/get_unwatch_stat", "GetUnwatchLivingStat", cRef.GetUnwatchLivingStat)
	registerClientAPIPath("/cgi-bin/school/living/get_unwatch_stat_v2", "GetUnwatchStatV2", cRef.GetUnwatchStatV2)
	registerClientAPIPath("/cgi-bin/school/living/get_watch_stat", "GetLivingWatchStat", cRef.GetLivingWatchStat)
	registerClientAPIPath("/cgi-bin/school/living/get_watch_stat_v2", "GetWatchStatV2", cRef.GetWatchStatV2)
}
