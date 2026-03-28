package wecom

func init() {
	registerClientAPIPath("/cgi-bin/batch/openuserid_to_userid", "BatchOpenUserIdToUserId", cRef.BatchOpenUserIdToUserId)
	registerClientAPIPath("/cgi-bin/idconvert/convert_tmp_external_userid", "ConvertTmpExternalUserid", cRef.ConvertTmpExternalUserid)
}
