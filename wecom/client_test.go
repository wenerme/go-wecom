package wecom

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/assert"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// nolint:funlen
func TestNewClient(t *testing.T) {
	mux := chi.NewMux()

	getTokenCount := 0
	mockAccessToken := "AccessToken" + createNonce()
	mockJsTicket := "JsTicket" + createNonce()
	mockAgentTicket := "AgentTicket" + createNonce()
	mockProviderToken := "ProviderToken" + createNonce()
	mockCorpID := "CorpID" + createNonce()
	mockAgentID := "AgentID" + createNonce()
	mockCorpSecret := "CorpSecret" + createNonce()

	requireAccessToken := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			if request.URL.Query().Get("access_token") != mockAccessToken {
				fmt.Printf("%v %v\n\tInvalid Token", request.Method, request.RequestURI)
				render.JSON(writer, request, GenericResponse{ErrorCode: 1, ErrorMessage: "invalid token"})
				return
			}
			next.ServeHTTP(writer, request)
		})
	}

	api := mux.With(requireAccessToken)
	mux.Get("/cgi-bin/gettoken", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Query().Get("corpsecret") != "CorpSecret" {
			render.JSON(writer, request, GenericResponse{ErrorCode: 1, ErrorMessage: "invalid secret"})
			return
		}
		render.JSON(writer, request, TokenResponse{AccessToken: mockAccessToken, ExpiresIn: 7200})
		getTokenCount++
		assert.Equal(t, 1, getTokenCount)
	})
	api.Get("/cgi-bin/get_jsapi_ticket", func(writer http.ResponseWriter, request *http.Request) {
		render.JSON(writer, request, TicketResponse{Ticket: mockJsTicket, ExpiresIn: 7200})
		getTokenCount++
		assert.Equal(t, 2, getTokenCount)
	})
	api.Get("/cgi-bin/ticket/get", func(writer http.ResponseWriter, request *http.Request) {
		render.JSON(writer, request, TicketResponse{Ticket: mockAgentTicket, ExpiresIn: 7200})
		getTokenCount++
		assert.Equal(t, 3, getTokenCount)
	})
	mux.Post("/cgi-bin/service/get_provider_token", func(writer http.ResponseWriter, request *http.Request) {
		r := GetProviderTokenRequest{}
		assert.NoError(t, render.Decode(request, &r))
		if r.CorpID != mockCorpID && r.ProviderSecret != mockProviderToken {
			render.JSON(writer, request, GenericResponse{ErrorCode: 1, ErrorMessage: "invalid secret"})
			return
		}
		render.JSON(writer, request, ProviderTokenResponse{ProviderAccessToken: mockProviderToken, ExpiresIn: 7200})
		getTokenCount++
		assert.Equal(t, 4, getTokenCount)
	})
	api.Get("/cgi-bin/get_api_domain_ip", func(writer http.ResponseWriter, request *http.Request) {
		// test error handler
		render.JSON(writer, request, GenericResponse{ErrorCode: 1, ErrorMessage: "test"})
	})
	mockIPList := IPListResponse{IPList: []string{"127.0.0.1"}}
	api.Get("/cgi-bin/getcallbackip", func(writer http.ResponseWriter, request *http.Request) {
		render.JSON(writer, request, mockIPList)
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	// token store - 可以使用数据库
	store := &SyncMapStore{}
	// 加载缓存 - 复用之前的 Token
	if bytes, err := os.ReadFile("wecom-cache.json"); err == nil {
		_ = store.Restore(bytes)
	}
	// 当 Token 变化时生成缓存
	store.OnChange = func(s *SyncMapStore) {
		_ = os.WriteFile("wecom-cache.json", s.Dump(), 0o600)
	}

	client := NewClient(Conf{
		CorpID:        mockCorpID,
		AgentID:       mockAgentID,
		CorpSecret:    mockCorpSecret,
		TokenProvider: &TokenCache{Store: store},
	})
	// client.Request.Options = append(client.Request.Options, req.DebugHook(nil))
	client.Request.BaseURL = server.URL

	{
		_, err := client.AccessToken()
		assert.Error(t, err)
	}
	client.Conf.CorpSecret = "CorpSecret"
	token, err := client.AccessToken()
	assert.NoError(t, err)
	assert.Equal(t, mockAccessToken, token)
	ticket, err := client.JsAPITicket()
	assert.NoError(t, err)
	assert.Equal(t, mockJsTicket, ticket)
	agentTicket, err := client.AgentTicket()
	assert.NoError(t, err)
	assert.Equal(t, mockAgentTicket, agentTicket)
	providerToken, err := client.ProviderAccessToken()
	assert.NoError(t, err)
	assert.Equal(t, mockProviderToken, providerToken)

	{
		config, err := client.SignConfig("http://test")
		assert.NoError(t, err)
		assert.NotEmpty(t, config.Signature)
	}
	{
		config, err := client.SignAgentConfig("http://test")
		assert.NoError(t, err)
		assert.NotEmpty(t, config.Signature)
	}

	_, err = client.GetAPIDomainIP()
	assert.Error(t, err)
	assert.EqualError(t, err, GenericResponse{ErrorCode: 1, ErrorMessage: "test"}.Error())

	{
		r, err := client.GetCallbackIP()
		assert.NoError(t, err)
		assert.Equal(t, mockIPList, r)
	}
}

func DecodeURLValues(values url.Values, out interface{}) error {
	m := map[string]interface{}{}
	for k, v := range values {
		if len(v) == 0 {
			continue
		}
		switch len(v) {
		case 0:
		case 1:
			m[k] = v[0]
		default:
			m[k] = v
		}
	}
	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           out,
		WeaklyTypedInput: true,
		TagName:          "json",
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(m)
}
