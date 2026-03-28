package wecom

func init() {
	registerClientAPIPath("/cgi-bin/agent/set", "SetAgent", cRef.SetAgent)
	registerClientAPIPath("/cgi-bin/menu/create", "CreateMenu", cRef.CreateMenu)
	registerClientAPIPath("/cgi-bin/menu/delete", "DeleteMenu", cRef.DeleteMenu)
	registerClientAPIPath("/cgi-bin/menu/get", "GetMenu", cRef.GetMenu)
}
