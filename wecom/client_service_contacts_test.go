package wecom

func init() {
	registerClientAPIPath("/cgi-bin/service/media/upload", "ProviderUploadMedia", cRef.ProviderUploadMedia)
	registerClientAPIPath("/cgi-bin/service/contact/id_translate", "ProviderIDTranslateContact", cRef.ProviderIDTranslateContact)
	registerClientAPIPath("/cgi-bin/service/batch/getresult", "ProviderBatchGetResult", cRef.ProviderBatchGetResult)
	registerClientAPIPath("/cgi-bin/service/contact/sort", "ProviderSortContact", cRef.ProviderSortContact)
}
