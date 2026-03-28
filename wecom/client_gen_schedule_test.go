package wecom

func init() {
	registerClientAPIPath("/cgi-bin/oa/schedule/add_attendees", "AddScheduleAttendees", cRef.AddScheduleAttendees)
	registerClientAPIPath("/cgi-bin/oa/schedule/del_attendees", "DeleteScheduleAttendees", cRef.DeleteScheduleAttendees)
}
