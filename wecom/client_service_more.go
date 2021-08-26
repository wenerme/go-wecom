package wecom

/*
https://work.weixin.qq.com/api/doc/10975
*/

type SetSessionInfo struct {
	AppID    []int `json:"appid"`
	AuthType int   `json:"auth_type"`
}

type GetAdminListItem struct {
	// UserID 管理员的userid
	UserID string `json:"userid"`
	// OpenUserID 管理员的open_userid
	OpenUserID string `json:"open_userid"`
	// AuthType 该管理员对应用的权限：0&#x3D;发消息权限，1&#x3D;管理权限
	AuthType int `json:"auth_type"`
}
