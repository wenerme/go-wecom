package wecom

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type OpaqueToken interface {
	GetAccessToken() string
	GetExpiresIn() int
}

func (r ProviderTokenResponse) GetAccessToken() string {
	return r.ProviderAccessToken
}

func (r ProviderTokenResponse) GetExpiresIn() int {
	return r.ExpiresIn
}

func (r TokenResponse) GetAccessToken() string {
	return r.AccessToken
}

func (r TokenResponse) GetExpiresIn() int {
	return r.ExpiresIn
}

func (r TicketResponse) GetAccessToken() string {
	return r.Ticket
}

func (r TicketResponse) GetExpiresIn() int {
	return r.ExpiresIn
}

func (r SuiteTokenResponse) GetAccessToken() string {
	return r.SuiteAccessToken
}

func (r SuiteTokenResponse) GetExpiresIn() int {
	return r.ExpiresIn
}

func (r PreAuthCodeResponse) GetAccessToken() string {
	return r.PreAuthCode
}

func (r PreAuthCodeResponse) GetExpiresIn() int {
	return r.ExpiresIn
}

func (r ProviderGetCorpTokenResponse) GetAccessToken() string {
	return r.AccessToken
}

func (r ProviderGetCorpTokenResponse) GetExpiresIn() int {
	return r.ExpiresIn
}

type GenericToken struct {
	Type      string `json:",omitempty"` //
	OwnerID   string `json:",omitempty"` // if owner change, this token become invalid
	Depends   string `json:",omitempty"` // extra depends
	Secret    string `json:",omitempty"` // secret content
	ExpiresIn int    `json:",omitempty"` // valid in seconds - 0 means no expire
	ExpiresAt int64  `json:",omitempty"` // expires unix timestamp
}

func (t *GenericToken) IsValid() bool {
	if t == nil || t.Secret == "" {
		return false
	}
	return t.ExpiresIn == 0 || t.ExpiresAt >= timeNow().Unix()
}

func (t GenericToken) GetAccessToken() string {
	return t.Secret
}

func (t GenericToken) GetExpiresAt() int64 {
	return t.ExpiresAt
}

func (t GenericToken) GetExpiresIn() int {
	return t.ExpiresIn
}

func (t *GenericToken) SetFromToken(token OpaqueToken) {
	if token == nil {
		return
	}
	t.Secret = token.GetAccessToken()
	t.ExpiresIn = token.GetExpiresIn()
	t.ExpiresAt = 0
	if i, ok := token.(interface {
		GetExpiresAt() int64
	}); ok {
		t.ExpiresAt = i.GetExpiresAt()
	}
	if t.ExpiresIn > 0 && t.ExpiresAt == 0 {
		t.ExpiresAt = int64(t.ExpiresIn) + timeNowUnix()
	}
}

var (
	timeNow     = time.Now
	timeNowUnix = func() int64 {
		return timeNow().Unix()
	}
)

// Refresh token, return changed and error
func (t *GenericToken) Refresh(exp *GenericToken, f func() (OpaqueToken, error)) (bool, error) {
	var changed bool
	if t.ShouldRefresh(exp) {
		token, err := f()
		// can keep using valid
		switch {
		case err != nil && !t.IsValid():
			return false, err
		case err != nil:
			// proper log error
			log.Printf("Refresh token %v ignoreable error: %v", exp.Type, err.Error())
		case err == nil && token != nil:
			t.SetFromToken(token)
			t.Type = exp.Type
			t.OwnerID = exp.OwnerID
			t.Depends = exp.Depends
			changed = true
		}
	}
	return changed, nil
}

func (t *GenericToken) ShouldRefresh(exp *GenericToken) bool {
	switch {
	case t == nil || t.Secret == "":
	case t.OwnerID != exp.OwnerID || t.Depends != exp.Depends:
	case t.ExpiresIn > 0 && t.ExpiresAt-int64(.2*float64(t.ExpiresIn)) <= timeNow().Unix():
		// 80% expires_in
	default:
		return false
	}
	return true
}

type TokenCache struct {
	Store TokenLoadStore
}

func keyOfGenericToken(t *GenericToken) (k string, err error) {
	if t.Type == "" {
		return "", errors.New("token need type")
	}
	k = t.Type
	if t.OwnerID != "" {
		k += "." + t.OwnerID
	}
	return
}

func (tc *TokenCache) Refresh(exp *GenericToken, f func() (OpaqueToken, error)) (string, error) {
	token := *exp
	_, err := tc.Store.Load(&token, func(last *GenericToken) (out *GenericToken, changed bool, err error) {
		if last == nil {
			last = &GenericToken{}
		}
		changed, err = last.Refresh(exp, f)
		return last, changed, err
	})
	return token.Secret, err
}

type TokenProvider interface {
	Refresh(exp *GenericToken, f func() (OpaqueToken, error)) (string, error)
}

type TokenLoadStore interface {
	Load(out *GenericToken, load func(last *GenericToken) (out *GenericToken, changed bool, err error)) (changed bool, err error)
}

type SyncMapStore struct {
	m        sync.Map
	OnChange func(s *SyncMapStore)
}

type cacheEntry struct {
	Key  string
	Data GenericToken
}

func (s *SyncMapStore) Restore(data []byte) error {
	var a []*cacheEntry
	if data != nil {
		err := json.Unmarshal(data, &a)
		if err != nil {
			return err
		}
	}
	for _, v := range a {
		if v.Key != "" && v.Data.IsValid() {
			s.m.Store(v.Key, v.Data)
		}
	}
	return nil
}

func (s *SyncMapStore) Dump() []byte {
	var a []*cacheEntry
	s.m.Range(func(key, value interface{}) bool {
		e := &cacheEntry{
			Key:  key.(string),
			Data: value.(GenericToken),
		}
		if e.Data.IsValid() {
			a = append(a, e)
		}
		return true
	})
	out, err := json.Marshal(a)
	if err != nil {
		log.Println("Marshal Dump failed:", err.Error())
	}
	return out
}

func (s *SyncMapStore) Set(in *GenericToken) (err error) {
	key, err := keyOfGenericToken(in)
	if err != nil {
		return err
	}
	if in.Secret == "" {
		s.m.Delete(key)
		return
	}
	if err == nil {
		s.m.Store(key, *in)
	}
	return
}

func (s *SyncMapStore) Load(out *GenericToken, loadFunc func(last *GenericToken) (out *GenericToken, changed bool, err error)) (changed bool, err error) {
	key, err := keyOfGenericToken(out)
	if err != nil {
		return false, err
	}

	// no exclusive lock here
	load, loaded := s.m.Load(key)
	var data *GenericToken
	if loaded {
		t := load.(GenericToken)
		data = &t
	}

	o, c, err := loadFunc(data)
	if err == nil && c {
		data = o
		s.m.Store(key, *o)
		if s.OnChange != nil {
			s.OnChange(s)
		}
	}
	if err == nil && data != nil {
		*out = *data
	}
	return c, err
}

func (s *SyncMapStore) Get(out *GenericToken) (found bool, err error) {
	key, err := keyOfGenericToken(out)
	if err != nil {
		return false, err
	}

	var load interface{}
	load, found = s.m.Load(key)
	if found {
		*out = load.(GenericToken)
	}
	return
}
