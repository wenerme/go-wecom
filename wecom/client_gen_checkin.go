package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// AddCheckinOption 创建打卡规则
// 企业可通过自建应用或授权的代开发应用，为企业添加打卡规则。
//
// see https://developer.work.weixin.qq.com/document/path/98057
func (c *Client) AddCheckinOption(r *AddCheckinOptionRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/checkin/add_checkin_option",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddCheckinOptionRequest is request of Client.AddCheckinOption
type AddCheckinOptionRequest struct {
	// EffectiveNow 是否立即生效，默认为false
	EffectiveNow bool `json:"effective_now"`
	// Group 打卡规则详细定义
	Group any `json:"group" validate:"required"`
}

// AddCheckinRecord 添加打卡记录
// 可通过接口写入打卡记录，匹配打卡规则后可在企业微信打卡明细、统计中参与展示。
//
// see https://developer.work.weixin.qq.com/document/path/99647
func (c *Client) AddCheckinRecord(r *AddCheckinRecordRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/checkin/add_checkin_record",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddCheckinRecordRequest is request of Client.AddCheckinRecord
type AddCheckinRecordRequest struct {
	// Records 打卡记录，一批最多200个
	Records []map[string]any `json:"records" validate:"required"`
	// UserID 用户id
	UserID string `json:"userid" validate:"required"`
	// CheckinTime 打卡时间。Unix时间戳
	CheckinTime int `json:"checkin_time" validate:"required"`
	// LocationTitle 打卡地点title，限制1024字符
	LocationTitle string `json:"location_title" validate:"required"`
	// LocationDetail 打卡地点详情限制1024字符
	LocationDetail string `json:"location_detail" validate:"required"`
	// Notes 打卡备注限制1024字符
	Notes string `json:"notes"`
	// Wifiname 打卡wifi名称限制1024字符
	Wifiname string `json:"wifiname"`
	// Wifimac 打卡的MAC地址/bssid 满足正则表达式^[A-Fa-f0-9]{2}:[A-Fa-f0-9]{2}:[A-Fa-f0-9]{2}:[A-Fa-f0-9]{2}:[A-Fa-f0-9]{2}...
	Wifimac string `json:"wifimac"`
	// Mediaids 打卡的附件media_id，可使用media/upload上传附件。当前最多只允许传1个
	Mediaids []any `json:"mediaids"`
	// Lat 位置打卡地点纬度，是实际纬度的1000000倍，与腾讯地图一致采用GCJ-02坐标系统标准 范围 -90000000,90000000
	Lat int `json:"lat"`
	// Lng 位置打卡地点经度，是实际纬度的1000000倍，与腾讯地图一致采用GCJ-02坐标系统标准 范围-180000000,180000000
	Lng int `json:"lng"`
	// DeviceType 打卡设备类型：1、门禁 2、考勤机（人脸识别、指纹识别）3、其他；
	DeviceType int `json:"device_type" validate:"required"`
	// DeviceDetail 打卡设备品牌：字符串写入（限制40个字符内）
	DeviceDetail string `json:"device_detail" validate:"required"`
}

// AddCheckinUserFace 录入打卡人员人脸信息
// 企业可通过自建应用，为企业打卡人员录入人脸信息，人脸信息仅用于人脸打卡。
//
// see https://developer.work.weixin.qq.com/document/path/93456
func (c *Client) AddCheckinUserFace(r *AddCheckinUserFaceRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/checkin/addcheckinuserface",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddCheckinUserFaceRequest is request of Client.AddCheckinUserFace
type AddCheckinUserFaceRequest struct {
	// UserID 需要录入的用户id
	UserID string `json:"userid"`
	// Userface 需要录入的人脸图片数据，需要将图片数据base64处理后填入
	Userface string `json:"userface"`
}

// GetCheckinDayData 获取打卡日报数据
// 企业可通过具有调用权限的应用，获取应用可见范围内指定员工指定日期内的打卡日报统计数据。
//
// see https://developer.work.weixin.qq.com/document/path/93451
func (c *Client) GetCheckinDayData(r *GetCheckinDayDataRequest, opts ...any) (out GetCheckinDayDataResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/checkin/getcheckin_daydata",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCheckinDayDataRequest is request of Client.GetCheckinDayData
type GetCheckinDayDataRequest struct {
	// Starttime 获取日报的开始时间。0点Unix时间戳
	Starttime int `json:"starttime" validate:"required"`
	// Endtime 获取日报的结束时间。0点Unix时间戳
	Endtime int `json:"endtime" validate:"required"`
	// Useridlist 获取日报的userid列表。单个userid不少于1字节，不多于64字节。可填充个数：1 ~ 100
	Useridlist []string `json:"useridlist" validate:"required"`
}

// GetCheckinDayDataResponse is response of Client.GetCheckinDayData
type GetCheckinDayDataResponse struct {
	// Datas 日报数据列表
	Datas []map[string]any `json:"datas"`
}

// GetCheckinMonthData 获取打卡月报数据
// 企业可通过具有调用权限的应用，获取应用可见范围内指定员工指定日期内的打卡月报统计数据。
//
// see https://developer.work.weixin.qq.com/document/path/93452
func (c *Client) GetCheckinMonthData(r *GetCheckinMonthDataRequest, opts ...any) (out GetCheckinMonthDataResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/checkin/getcheckin_monthdata",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCheckinMonthDataRequest is request of Client.GetCheckinMonthData
type GetCheckinMonthDataRequest struct {
	// Starttime 获取月报的开始时间。0点Unix时间戳
	Starttime int `json:"starttime" validate:"required"`
	// Endtime 获取月报的结束时间。0点Unix时间戳
	Endtime int `json:"endtime" validate:"required"`
	// Useridlist 员工UserID列表，不少于1字节，不多于64字节，可填充个数：1 ~ 100
	Useridlist []string `json:"useridlist" validate:"required"`
}

// GetCheckinMonthDataResponse is response of Client.GetCheckinMonthData
type GetCheckinMonthDataResponse struct {
	// Datas 月报数据列表
	Datas []map[string]any `json:"datas"`
}

// GetCheckinData 获取打卡记录数据
// 应用可通过本接口，获取可见范围内员工指定时间段内的打卡记录数据。
//
// see https://developer.work.weixin.qq.com/document/path/93450
func (c *Client) GetCheckinData(r *GetCheckinDataRequest, opts ...any) (out GetCheckinDataResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/checkin/getcheckindata",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCheckinDataRequest is request of Client.GetCheckinData
type GetCheckinDataRequest struct {
	// Opencheckindatatype 打卡类型。1：上下班打卡；2：外出打卡；3：全部打卡
	Opencheckindatatype int `json:"opencheckindatatype" validate:"required"`
	// Starttime 获取打卡记录的开始时间。Unix时间戳
	Starttime int `json:"starttime" validate:"required"`
	// Endtime 获取打卡记录的结束时间。Unix时间戳
	Endtime int `json:"endtime" validate:"required"`
	// Useridlist 需要获取打卡记录的用户列表
	Useridlist []string `json:"useridlist" validate:"required"`
}

// GetCheckinDataResponse is response of Client.GetCheckinData
type GetCheckinDataResponse struct {
	// Checkindata 打卡记录列表
	Checkindata []map[string]any `json:"checkindata"`
	// UserID 用户id
	UserID string `json:"userid"`
	// Groupname 打卡规则名称
	Groupname string `json:"groupname"`
	// CheckinType 打卡类型。字符串，目前有：上班打卡，下班打卡，外出打卡，仅记录打卡时间和位置
	CheckinType string `json:"checkin_type"`
	// ExceptionType 异常类型，字符串，包括：时间异常，地点异常，未打卡，wifi异常，非常用设备。如果有多个异常，以分号间隔
	ExceptionType string `json:"exception_type"`
	// CheckinTime 打卡时间。Unix时间戳
	CheckinTime int `json:"checkin_time"`
	// LocationTitle 打卡地点title
	LocationTitle string `json:"location_title"`
	// LocationDetail 打卡地点详情
	LocationDetail string `json:"location_detail"`
	// Wifiname 打卡wifi名称
	Wifiname string `json:"wifiname"`
	// Notes 打卡备注
	Notes string `json:"notes"`
	// Wifimac 打卡的MAC地址/bssid
	Wifimac string `json:"wifimac"`
	// Mediaids 打卡的附件media_id，可使用media/get获取附件
	Mediaids []any `json:"mediaids"`
	// Lat 位置打卡地点纬度，是实际纬度的1000000倍，与腾讯地图一致采用GCJ-02坐标系统标准
	Lat int `json:"lat"`
	// Lng 位置打卡地点经度，是实际经度的1000000倍，与腾讯地图一致采用GCJ-02坐标系统标准
	Lng int `json:"lng"`
	// Deviceid 打卡设备id
	Deviceid string `json:"deviceid"`
	// SchCheckinTime 标准打卡时间，指此次打卡时间对应的标准上班时间或标准下班时间
	SchCheckinTime int `json:"sch_checkin_time"`
	// GroupID 规则id，表示打卡记录所属规则的id
	GroupID int `json:"groupid"`
	// ScheduleID 班次id，表示打卡记录所属规则中，所属班次的id
	ScheduleID int `json:"schedule_id"`
	// TimelineID 时段id，表示打卡记录所属规则中，某一班次中的某一时段的id
	TimelineID int `json:"timeline_id"`
}

// GetCheckinOption 获取员工打卡规则
// 自建应用、第三方应用和代开发应用可使用此接口，获取可见范围内指定员工指定日期的打卡规则。
//
// see https://developer.work.weixin.qq.com/document/path/93449
func (c *Client) GetCheckinOption(r *GetCheckinOptionRequest, opts ...any) (out GetCheckinOptionResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/checkin/getcheckinoption",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCheckinOptionRequest is request of Client.GetCheckinOption
type GetCheckinOptionRequest struct {
	// Datetime 需要获取规则的日期当天0点的Unix时间戳
	Datetime int `json:"datetime" validate:"required"`
	// Useridlist 需要获取打卡规则的用户列表
	Useridlist []string `json:"useridlist" validate:"required"`
}

// GetCheckinOptionResponse is response of Client.GetCheckinOption
type GetCheckinOptionResponse struct {
	// Info 返回的打卡规则列表
	Info []map[string]any `json:"info"`
	// UserID 打卡人员userid
	UserID string `json:"userid"`
	// Group 打卡规则相关信息
	Group any `json:"group"`
}

// GetCheckInScheduleList 获取打卡人员排班信息
// 获取应用可见范围内、打卡规则为按班次上下班规则的指定员工指定时间段内的排班信息
//
// see https://developer.work.weixin.qq.com/document/path/93453
func (c *Client) GetCheckInScheduleList(r *GetCheckInScheduleListRequest, opts ...any) (out GetCheckInScheduleListResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/checkin/getcheckinschedulist",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCheckInScheduleListRequest is request of Client.GetCheckInScheduleList
type GetCheckInScheduleListRequest struct {
	// Useridlist 需要获取排班信息的用户列表（不超过100个）
	Useridlist []string `json:"useridlist" validate:"required"`
	// Starttime 获取排班信息的开始时间。Unix时间戳
	Starttime int `json:"starttime" validate:"required"`
	// Endtime 获取排班信息的结束时间。Unix时间戳（与starttime跨度不超过一个月）
	Endtime int `json:"endtime" validate:"required"`
}

// GetCheckInScheduleListResponse is response of Client.GetCheckInScheduleList
type GetCheckInScheduleListResponse struct {
	// ScheduleList 排班表信息列表
	ScheduleList []map[string]any `json:"schedule_list"`
}

// GetCorpCheckinOption 获取企业所有打卡规则
// 自建应用、代开发应用可用此接口，获取企业内所有打卡规则。
//
// see https://developer.work.weixin.qq.com/document/path/93448
func (c *Client) GetCorpCheckinOption(r *GetCorpCheckinOptionRequest, opts ...any) (out GetCorpCheckinOptionResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/checkin/getcorpcheckinoption",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCorpCheckinOptionRequest is request of Client.GetCorpCheckinOption
type GetCorpCheckinOptionRequest struct {}

// GetCorpCheckinOptionResponse is response of Client.GetCorpCheckinOption
type GetCorpCheckinOptionResponse struct {
	// Group 企业规则信息列表
	Group []map[string]any `json:"group"`
}

// PunchCorrection 为打卡人员补卡
// 为打卡人员补录打卡记录
//
// see https://developer.work.weixin.qq.com/document/path/95803
func (c *Client) PunchCorrection(r *PunchCorrectionRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/checkin/punch_correction",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// PunchCorrectionRequest is request of Client.PunchCorrection
type PunchCorrectionRequest struct {
	// UserID 需要补卡的成员userid
	UserID string `json:"userid" validate:"required"`
	// ScheduleDateTime 应打卡日期，为当天0点的Unix时间戳
	ScheduleDateTime int `json:"schedule_date_time" validate:"required"`
	// ScheduleCheckinTime 应打卡时间点，相对应打卡日期0点的偏移秒数
	ScheduleCheckinTime int `json:"schedule_checkin_time"`
	// CheckinTime 实际打卡时间，Unix时间戳
	CheckinTime int `json:"checkin_time" validate:"required"`
	// Remark 备注信息
	Remark string `json:"remark"`
}

// SetCheckinScheduleList 为打卡人员排班
// 企业可通过具有调用权限的应用，为打卡规则为“按班次上下班”规则的指定员工排班。
//
// see https://developer.work.weixin.qq.com/document/path/93454
func (c *Client) SetCheckinScheduleList(r *SetCheckinScheduleListRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/checkin/setcheckinschedulist",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SetCheckinScheduleListRequest is request of Client.SetCheckinScheduleList
type SetCheckinScheduleListRequest struct {
	// Items 排班表信息
	Items []map[string]any `json:"items" validate:"required"`
	// GroupID 打卡规则的规则id，可通过“获取打卡规则”、“获取打卡数据”、“获取打卡人员排班信息”等相关接口获取
	GroupID int `json:"groupid" validate:"required"`
	// Yearmonth 排班表月份，格式为年月，如202011
	Yearmonth int `json:"yearmonth" validate:"required"`
}

// GetHardwareCheckinData 获取设备打卡数据
// 获取可见范围内成员在考勤设备上产生的原始打卡记录，包括未被打卡应用记录的不符合打卡规则的记录。
//
// see https://developer.work.weixin.qq.com/document/path/94126
func (c *Client) GetHardwareCheckinData(r *GetHardwareCheckinDataRequest, opts ...any) (out GetHardwareCheckinDataResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/hardware/get_hardware_checkin_data",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetHardwareCheckinDataRequest is request of Client.GetHardwareCheckinData
type GetHardwareCheckinDataRequest struct {
	// FilterType 过滤类型，1表示按打卡时间过滤，2表示按设备上传打卡记录的时间过滤，默认值是1
	FilterType int `json:"filter_type"`
	// Starttime Unix时间戳，当filter_type为1时，表示打卡的开始时间；当filter_type为2时，表示设备上传记录的开始时间
	Starttime int `json:"starttime" validate:"required"`
	// Endtime Unix时间戳，当filter_type为1时，表示打卡的结束时间；当filter_type为2时，表示设备上传记录的结束时间
	Endtime int `json:"endtime" validate:"required"`
	// Useridlist 需要获取打卡记录的用户列表
	Useridlist []string `json:"useridlist" validate:"required"`
}

// GetHardwareCheckinDataResponse is response of Client.GetHardwareCheckinData
type GetHardwareCheckinDataResponse struct {
	// Checkindata 打卡记录列表
	Checkindata []map[string]any `json:"checkindata"`
	// UserID 用户id
	UserID string `json:"userid"`
	// CheckinTime 打卡时间。Unix时间戳
	CheckinTime int `json:"checkin_time"`
	// DeviceSn 打卡设备的sn
	DeviceSn string `json:"device_sn"`
	// DeviceName 打卡设备名
	DeviceName string `json:"device_name"`
}

