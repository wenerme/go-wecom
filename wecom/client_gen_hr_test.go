package wecom

func init() {
	registerClientAPIPath("/cgi-bin/hr/get_fields", "GetEmployeeFieldConfig", cRef.GetEmployeeFieldConfig)
	registerClientAPIPath("/cgi-bin/hr/get_staff_info", "GetStaffInfo", cRef.GetStaffInfo)
	registerClientAPIPath("/cgi-bin/hr/update_staff_info", "UpdateStaffInfo", cRef.UpdateStaffInfo)
}
