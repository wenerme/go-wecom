package wecom

// ExtAttrs extra attributes model
type ExtAttrs struct {
	Attrs []ExtAttr `json:"attrs"`
}

// ExtAttr extra attribute model
type ExtAttr struct {
	// Type 0 文本, 1 网页, 2 小程序
	Type int    `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
	Text struct {
		Value string `json:"value,omitempty"`
	} `json:"text,omitempty"`
	Web struct {
		URL   string `json:"url,omitempty"`
		Title string `json:"title,omitempty"`
	} `json:"web,omitempty"`
	MiniProgram struct {
		AppID    string `json:"appid,omitempty"`
		PagePath string `json:"pagepath,omitempty"`
		Title    string `json:"title,omitempty"`
	} `json:"miniprogram,omitempty"`
}
