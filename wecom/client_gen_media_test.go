package wecom

func init() {
	registerClientAPIPath("/cgi-bin/media/get", "GetTempMedia", cRef.GetTempMedia)
	registerClientAPIPath("/cgi-bin/media/get/jssdk", "GetMediaVoice", cRef.GetMediaVoice)
	registerClientAPIPath("/cgi-bin/media/upload", "UploadMedia", cRef.UploadMedia)
	registerClientAPIPath("/cgi-bin/media/upload_by_url", "UploadMediaByUrl", cRef.UploadMediaByUrl)
	registerClientAPIPath("/cgi-bin/media/uploadimg", "UploadImage", cRef.UploadImage)
}
