package wecom

func init() {
	registerClientAPIPath("/cgi-bin/oa/meetingroom/add", "AddMeetingRoom", cRef.AddMeetingRoom)
	registerClientAPIPath("/cgi-bin/oa/meetingroom/get_booking_info", "GetMeetingRoomBookingInfo", cRef.GetMeetingRoomBookingInfo)
}
