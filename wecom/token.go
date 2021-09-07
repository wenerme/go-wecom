package wecom

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/pkg/errors"
)

type Token interface {
	GetToken() string
	GetExpiresIn() int
}

func (r ProviderTokenResponse) GetToken() string {
	return r.ProviderAccessToken
}

func (r ProviderTokenResponse) GetExpiresIn() int {
	return r.ExpiresIn
}

func (r TokenResponse) GetToken() string {
	return r.AccessToken
}

func (r TokenResponse) GetExpiresIn() int {
	return r.ExpiresIn
}

func (r TicketResponse) GetToken() string {
	return r.Ticket
}

func (r TicketResponse) GetExpiresIn() int {
	return r.ExpiresIn
}

func (r SuiteTokenResponse) GetToken() string {
	return r.SuiteAccessToken
}

func (r SuiteTokenResponse) GetExpiresIn() int {
	return r.ExpiresIn
}

func (r PreAuthCodeResponse) GetToken() string {
	return r.PreAuthCode
}

func (r PreAuthCodeResponse) GetExpiresIn() int {
	return r.ExpiresIn
}

type GenericToken struct {
	Name      string `json:",omitempty"` // for match
	OwnerID   string `json:",omitempty"` // if owner change, this token become invalid
	Depends   string `json:",omitempty"` // extra depends
	Token     string `json:",omitempty"` // token content
	ExpiresIn int    `json:",omitempty"` // valid in seconds - 0 means no expire
	ExpiresAt int64  `json:",omitempty"` // expires unix timestamp
}

func (t *GenericToken) IsValid() bool {
	if t == nil || t.Token == "" {
		return false
	}
	return t.ExpiresIn == 0 || t.ExpiresAt >= timeNow().Unix()
}

func (t GenericToken) GetToken() string {
	return t.Token
}

func (t GenericToken) GetExpiresAt() int64 {
	return t.ExpiresAt
}

func (t GenericToken) GetExpiresIn() int {
	return t.ExpiresIn
}

func (t *GenericToken) SetFromToken(token Token) {
	if token == nil {
		return
	}
	t.Token = token.GetToken()
	t.ExpiresIn = token.GetExpiresIn()
	if i, ok := token.(interface {
		GetExpiresAt() int64
	}); ok {
		t.ExpiresAt = i.GetExpiresAt()
	}
	if t.ExpiresIn > 0 && t.ExpiresAt == 0 {
		t.ExpiresAt = int64(t.ExpiresIn) + timeNow().Unix()
	}
}

var timeNow = time.Now

// Refresh token, return changed and error
func (t *GenericToken) Refresh(exp *GenericToken, f func() (Token, error)) (bool, error) {
	var changed bool
	if t.ShouldRefresh(exp) {
		token, err := f()
		// can keep using valid
		switch {
		case err != nil && !t.IsValid():
			return false, err
		case err != nil:
			// proper log error
			log.Printf("Refresh token %v ignoreable error: %v", exp.Name, err.Error())
		case err == nil && token != nil:
			t.SetFromToken(token)
			t.Name = exp.Name
			t.OwnerID = exp.OwnerID
			t.Depends = exp.Depends
			changed = true
		}
	}
	return changed, nil
}

func (t *GenericToken) ShouldRefresh(exp *GenericToken) bool {
	switch {
	case t == nil || t.Token == "":
	case t.OwnerID != exp.OwnerID || t.Depends != exp.Depends:
	case t.ExpiresIn > 0 && t.ExpiresAt-int64(.2*float64(t.ExpiresIn)) <= timeNow().Unix():
		// 80% expires_in
	default:
		return false
	}
	return true
}

type TokenCache struct {
	Store LoadStore
}

func (tc *TokenCache) KeyOf(t *GenericToken) (k string, err error) {
	if t.Name == "" {
		return "", errors.New("get token need name")
	}
	k = t.Name
	if t.OwnerID != "" {
		k += "@" + t.OwnerID
	}
	return
}

func (tc *TokenCache) Refresh(exp *GenericToken, f func() (Token, error)) (string, error) {
	key, err := tc.KeyOf(exp)
	if err != nil {
		return "", err
	}
	var token GenericToken
	_, err = tc.Store.Load(key, &token, func(last []byte) (out interface{}, changed bool, err error) {
		l := GenericToken{}
		if last != nil {
			_ = json.Unmarshal(last, &l)
		}
		changed, err = l.Refresh(exp, f)
		return l, changed, err
	})
	return token.Token, err
}

type TokenProvider interface {
	Refresh(exp *GenericToken, f func() (Token, error)) (string, error)
}

type LoadStore interface {
	Load(key string, out interface{}, load func(last []byte) (out interface{}, changed bool, err error)) (changed bool, err error)
}

type SyncMapStore struct {
	m        sync.Map
	OnChange func(s *SyncMapStore)
}

type cacheEntry struct {
	Key  string
	Data []byte
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
		if v.Key != "" && v.Data != nil {
			s.m.Store(v.Key, v.Data)
		}
	}
	return nil
}

func (s *SyncMapStore) Dump() []byte {
	var a []*cacheEntry
	s.m.Range(func(key, value interface{}) bool {
		a = append(a, &cacheEntry{
			Key:  key.(string),
			Data: value.([]byte),
		})
		return false
	})
	out, err := json.Marshal(a)
	if err != nil {
		log.Println("Marshal Dump failed:", err.Error())
	}
	return out
}

func (s *SyncMapStore) Set(key string, in interface{}) (err error) {
	if in == nil {
		s.m.Delete(key)
		return
	}
	var data []byte
	data, err = json.Marshal(in)
	if err == nil {
		s.m.Store(key, data)
	}
	return
}

func (s *SyncMapStore) Load(key string, out interface{}, loadFunc func(last []byte) (out interface{}, changed bool, err error)) (changed bool, err error) {
	// no exclusive lock here

	load, loaded := s.m.Load(key)
	var data []byte
	if loaded {
		data = load.([]byte)
	}

	o, c, err := loadFunc(data)
	if err == nil && c {
		data, err = json.Marshal(o)
		s.m.Store(key, data)
		if s.OnChange != nil {
			s.OnChange(s)
		}
	}
	if err == nil && data != nil {
		err = json.Unmarshal(data, out)
	}
	return c, err
}

func (s *SyncMapStore) Get(key string, out interface{}) (found bool, err error) {
	var data []byte
	var load interface{}
	load, found = s.m.Load(key)
	if found {
		data = load.([]byte)
		err = json.Unmarshal(data, out)
	}
	return
}
