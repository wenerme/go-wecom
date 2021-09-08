package wecom

type AddCalendar struct {
	// Organizer 指定的组织者userid。注意该字段指定后不可更新
	Organizer string `json:"organizer"  validate:"required"`
	// Readonly 日历组织者对日历是否只读权限（即不可编辑日历，不可在日历上添加日程，仅可作为组织者删除日历）。0-否；1-是。默认为1，即只读
	Readonly int `json:"readonly"  `
	// SetAsDefault 是否将该日历设置为组织者的默认日历。0-否；1-是。默认为0，即不设为默认日历第三方应用不支持使用该参数
	SetAsDefault int `json:"set_as_default"  `
	// Summary 日历标题。1 ~ 128 字符
	Summary string `json:"summary"  validate:"required"`
	// Color 日历在终端上显示的颜色，RGB颜色编码16进制表示，例如：”#0000FF” 表示纯蓝色
	Color string `json:"color"  validate:"required"`
	// Description 日历描述。0 ~ 512 字符
	Description string `json:"description"  `
	// Shares 日历共享成员列表。最多2000人
	Shares []AddCalendarRequestShares `json:"shares"  `
}

type UpdateCalendar struct {
	// CalID 日历ID
	CalID string `json:"cal_id"  validate:"required"`
	// Readonly 日历组织者对日历是否只读权限（即不可编辑日历，不可在日历上添加日程，仅可作为组织者删除日历）。0-否；1-是。默认为1，即只读
	Readonly int `json:"readonly"  `
	// Summary 日历标题。1 ~ 128 字符
	Summary string `json:"summary"  validate:"required"`
	// Color 日历颜色，RGB颜色编码16进制表示，例如：”#0000FF” 表示纯蓝色
	Color string `json:"color"  validate:"required"`
	// Description 日历描述。0 ~ 512 字符
	Description string `json:"description"  `
	// Shares 日历共享成员列表。最多2000人
	Shares []UpdateCalendarRequestShares `json:"shares"  `
}

// GetCalendarResponseItem is item model of GetCalendarResponse.CalendarList
type GetCalendarResponseItem struct {
	// CalID 日历ID
	CalID string `json:"cal_id"  `
	// Organizer 指定的组织者userid
	Organizer string `json:"organizer"  `
	// Readonly 日历组织者对日历是否只读权限。0-否；1-是；
	Readonly int `json:"readonly"  `
	// Summary 日历标题。1 ~ 128 字符
	Summary string `json:"summary"  `
	// Color 日历颜色，RGB颜色编码16进制表示，例如：”#0000FF” 表示纯蓝色
	Color string `json:"color"  `
	// Description 日历描述。0 ~ 512 字符
	Description string `json:"description"  `
	// Shares 日历共享成员列表。最多2000人
	Shares []GetCalendarResponseShares `json:"shares"  `
}
