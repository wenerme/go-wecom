package wecom

func init() {
	registerClientAPIPath("/cgi-bin/checkin/add_checkin_option", "AddCheckinOption", cRef.AddCheckinOption)
	registerClientAPIPath("/cgi-bin/checkin/add_checkin_record", "AddCheckinRecord", cRef.AddCheckinRecord)
	registerClientAPIPath("/cgi-bin/checkin/addcheckinuserface", "AddCheckinUserFace", cRef.AddCheckinUserFace)
	registerClientAPIPath("/cgi-bin/checkin/getcheckin_daydata", "GetCheckinDayData", cRef.GetCheckinDayData)
	registerClientAPIPath("/cgi-bin/checkin/getcheckin_monthdata", "GetCheckinMonthData", cRef.GetCheckinMonthData)
	registerClientAPIPath("/cgi-bin/checkin/getcheckindata", "GetCheckinData", cRef.GetCheckinData)
	registerClientAPIPath("/cgi-bin/checkin/getcheckinoption", "GetCheckinOption", cRef.GetCheckinOption)
	registerClientAPIPath("/cgi-bin/checkin/getcheckinschedulist", "GetCheckInScheduleList", cRef.GetCheckInScheduleList)
	registerClientAPIPath("/cgi-bin/checkin/getcorpcheckinoption", "GetCorpCheckinOption", cRef.GetCorpCheckinOption)
	registerClientAPIPath("/cgi-bin/checkin/punch_correction", "PunchCorrection", cRef.PunchCorrection)
	registerClientAPIPath("/cgi-bin/checkin/setcheckinschedulist", "SetCheckinScheduleList", cRef.SetCheckinScheduleList)
	registerClientAPIPath("/cgi-bin/hardware/get_hardware_checkin_data", "GetHardwareCheckinData", cRef.GetHardwareCheckinData)
}
