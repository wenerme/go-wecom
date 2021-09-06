package wecom

import (
	"sync"
	"time"
)

type Token interface {
	GetToken() string
	// GetExpireAt return unix timestamp
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

func (r PreAuthCodeResponse) GetToken() string {
	return r.PreAuthCode
}

func (r PreAuthCodeResponse) GetExpireAt() int64 {
	return r.ExpireAt
}

func IsExpired(token Token) bool {
	return token.GetExpireAt() >= time.Now().Unix()
}

type GenericToken struct {
	Token    string
	ExpireIn int
	ExpireAt int64
}

func (r GenericToken) GetToken() string {
	return r.Token
}

func (r GenericToken) GetExpireAt() int64 {
	return r.ExpireAt
}

func BuildTokenProvider(getter func() (string, int64, error)) func() (string, error) {
	cache := GenericToken{}
	next := GenericToken{}
	m := sync.Mutex{}
	return func() (string, error) {
		var err error
		if IsExpired(cache) {
			m.Lock()
			defer m.Unlock()
			if IsExpired(cache) {
				next.Token, next.ExpireAt, err = getter()
				if err == nil {
					cache = next
				}
			}
		}
		return validToken(cache, "token", err)
	}
}
