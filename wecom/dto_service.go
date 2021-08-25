package wecom

/*
https://work.weixin.qq.com/api/doc/10975
*/

type GetSuiteTokenRequest struct {
	SuiteID     string `json:"suite_id"`
	SuiteSecret string `json:"suite_secret"`
	SuiteTicket string `json:"suite_ticket"`
}

type GetPreAuthCodeResponse struct {
	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int    `json:"expires_in"`
	ExpiresAt   int    `json:"expires_at"`
}
