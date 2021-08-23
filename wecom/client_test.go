package wecom

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
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

	client := NewClient(Conf{
		CorpID:     mockCorpID,
		AgentID:    mockAgentID,
		CorpSecret: mockCorpSecret,
	})
	// client.Request.Options = append(client.Request.Options, req.DebugHook(nil))
	client.Request.BaseURL = server.URL
	client.OnTokenUpdate = func(c *Client) {
		assert.Equal(t, mockAccessToken, c.AccessTokenCache.AccessToken)
		fmt.Println(
			"token:", c.AccessTokenCache.AccessToken,
			"ticket:", c.JsAPITicketCache.Ticket,
			"agent-ticket:", c.AgentTicketCache.Ticket,
			"provider-token:", c.ProviderTokenCache.ProviderAccessToken,
		)
	}
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
	providerToken, err := client.ProviderToken()
	assert.NoError(t, err)
	assert.Equal(t, mockProviderToken, providerToken)

	cache := client.DumpCache()
	nc := NewClient(Conf{})
	nc.LoadCache(&cache)
	assert.Equal(t, nc.AccessTokenCache, client.AccessTokenCache)
	assert.Equal(t, nc.AgentTicketCache, client.AgentTicketCache)
	assert.Equal(t, nc.JsAPITicketCache, client.JsAPITicketCache)
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
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(m)
}
