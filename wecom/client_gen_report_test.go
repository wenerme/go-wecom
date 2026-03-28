package wecom

func init() {
	registerClientAPIPath("/cgi-bin/oa/journal/download_wedrive_file", "DownloadWeDriveFile", cRef.DownloadWeDriveFile)
	registerClientAPIPath("/cgi-bin/oa/journal/get_record_detail", "GetJournalDetail", cRef.GetJournalDetail)
	registerClientAPIPath("/cgi-bin/oa/journal/get_record_list", "ListJournalRecord", cRef.ListJournalRecord)
	registerClientAPIPath("/cgi-bin/oa/journal/get_stat_list", "ListJournalStat", cRef.ListJournalStat)
}
