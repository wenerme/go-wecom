package wecom

// SuiteInstalledRequestParams suite installed callback params
type SuiteInstalledRequestParams struct {
	AuthCode  string `json:"auth_code"`
	ExpiresIn int    `json:"expires_in"` // 1200
	State     string `json:"state"`
}

func (r SuiteInstalledRequestParams) IsValid() bool {
	return r.AuthCode != "" && r.ExpiresIn > 0
}
