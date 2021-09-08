package wecom

// ExtAttrs extra attributes model
type ExtAttrs struct {
	Attrs []ExtAttr `json:"attrs"`
}

// ExtAttr extra attribute model
type ExtAttr struct {
	// Type 0 文本, 1 网页, 2 小程序
	Type int    `json:"type"`
	Name string `json:"name"`
	Text struct {
		Value string `json:"value"`
	} `json:"text,omitempty"`
	Web struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	} `json:"web,omitempty"`
	MiniProgram struct {
		AppID    string `json:"appid"`
		PagePath string `json:"pagepath"`
		Title    string `json:"title"`
	} `json:"miniprogram,omitempty"`
}
