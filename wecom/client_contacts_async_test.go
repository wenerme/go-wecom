package wecom

func init() {
	registerClientAPIPath("/cgi-bin/batch/syncuser", "BatchSyncUser", cRef.BatchSyncUser)
	registerClientAPIPath("/cgi-bin/batch/replaceuser", "BatchReplaceUser", cRef.BatchReplaceUser)
	registerClientAPIPath("/cgi-bin/batch/replaceparty", "BatchReplaceParty", cRef.BatchReplaceParty)
	registerClientAPIPath("/cgi-bin/batch/getresult", "BatchGetResult", cRef.BatchGetResult)
}
