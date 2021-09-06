package wecom

func init() {
	registerClientAPIPath("/cgi-bin/oa/calendar/add", "AddCalendar", cRef.AddCalendar)
	registerClientAPIPath("/cgi-bin/oa/calendar/update", "UpdateCalendar", cRef.UpdateCalendar)
	registerClientAPIPath("/cgi-bin/oa/calendar/get", "GetCalendar", cRef.GetCalendar)
	registerClientAPIPath("/cgi-bin/oa/calendar/del", "DeleteCalendar", cRef.DeleteCalendar)
	registerClientAPIPath("/cgi-bin/oa/schedule/add", "AddSchedule", cRef.AddSchedule)
	registerClientAPIPath("/cgi-bin/oa/schedule/update", "UpdateSchedule", cRef.UpdateSchedule)
	registerClientAPIPath("/cgi-bin/oa/schedule/get", "GetSchedule", cRef.GetSchedule)
	registerClientAPIPath("/cgi-bin/oa/schedule/del", "DeleteSchedule", cRef.DeleteSchedule)
	registerClientAPIPath("/cgi-bin/oa/schedule/get_by_calendar", "ScheduleGetByCalendar", cRef.ScheduleGetByCalendar)
}
