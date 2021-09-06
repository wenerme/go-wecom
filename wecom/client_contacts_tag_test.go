package wecom

func init() {
	registerClientAPIPath("/cgi-bin/tag/create", "CreateTag", cRef.CreateTag)
	registerClientAPIPath("/cgi-bin/tag/update", "UpdateTag", cRef.UpdateTag)
	registerClientAPIPath("/cgi-bin/tag/delete", "DeleteTag", cRef.DeleteTag)
	registerClientAPIPath("/cgi-bin/tag/get", "GetTag", cRef.GetTag)
	registerClientAPIPath("/cgi-bin/tag/addtagusers", "AddTagUsers", cRef.AddTagUsers)
	registerClientAPIPath("/cgi-bin/tag/deltagusers", "DeleteTagUsers", cRef.DeleteTagUsers)
	registerClientAPIPath("/cgi-bin/tag/list", "ListTag", cRef.ListTag)
}
