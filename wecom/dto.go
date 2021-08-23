package wecom

import (
	"fmt"
)

type GenericResponse struct {
	ErrorCode    int    `json:"errcode"`
	ErrorMessage string `json:"errmsg"`
}

func (r GenericResponse) AsError() error {
	if r.ErrorCode != 0 {
		return &r
	}
	return nil
}

func (r GenericResponse) Error() string {
	return fmt.Sprintf("%v: %v", r.ErrorCode, r.ErrorMessage)
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	ExpireAt    int64  `json:"expire_at"` //  extra
}

type GetProviderTokenRequest struct {
	CorpID         string `json:"corpid"`
	ProviderSecret string `json:"provider_secret"`
}

type ProviderTokenResponse struct {
	ProviderAccessToken string `json:"provider_access_token"`
	ExpiresIn           int    `json:"expires_in"`
	ExpireAt            int64  `json:"expire_at"` //  extra
}

type SuiteTokenResponse struct {
	SuiteAccessToken string `json:"suite_access_token"`
	ExpiresIn        int    `json:"expires_in"`
	ExpireAt         int64  `json:"expire_at"` //  extra
}

type TicketResponse struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
	ExpireAt  int64  `json:"expire_at"` //  extra
}

type IPListResponse struct {
	IPList []string `json:"ip_list"`
}
