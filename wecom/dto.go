package wecom

import "fmt"

type GenericResponse struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

func (r GenericResponse) AsError() error {
	if r.ErrorCode != 0 {
		return fmt.Errorf("wecom api(%v): %v", r.ErrorCode, r.ErrorMessage)
	}
	return nil
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ExpireAt    int64  `json:"expire_at"` //  extra
}

type TicketResponse struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
	ExpireAt  int64  `json:"expire_at"` //  extra
}

type GetApiDomainIpResponse struct {
	IpList []string `json:"ip_list"`
}
