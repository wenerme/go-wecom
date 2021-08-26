package wecom

import "time"

type Token interface {
	GetToken() string
	GetExpireAt() int64
}

func (r ProviderTokenResponse) GetToken() string {
	return r.ProviderAccessToken
}

func (r ProviderTokenResponse) GetExpireAt() int64 {
	return r.ExpireAt
}

func (r TokenResponse) GetToken() string {
	return r.AccessToken
}

func (r TokenResponse) GetExpireAt() int64 {
	return r.ExpireAt
}

func (r TicketResponse) GetToken() string {
	return r.Ticket
}

func (r TicketResponse) GetExpireAt() int64 {
	return r.ExpireAt
}

func (r SuiteTokenResponse) GetToken() string {
	return r.SuiteAccessToken
}

func (r SuiteTokenResponse) GetExpireAt() int64 {
	return r.ExpireAt
}

func (r GetPreAuthCodeResponse) GetToken() string {
	return r.PreAuthCode
}

func (r GetPreAuthCodeResponse) GetExpireAt() int64 {
	return r.ExpireAt
}

func IsExpired(token Token) bool {
	return token.GetExpireAt() >= time.Now().Unix()
}
