package wecom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func TestNewClient(t *testing.T) {
	mux := chi.NewMux()
	mux.Get("/cgi-bin/gettoken", func(writer http.ResponseWriter, request *http.Request) {
		render.JSON(writer, request, TokenResponse{AccessToken: "token", ExpiresIn: 300})
	})
	mux.Get("/cgi-bin/get_jsapi_ticket", func(writer http.ResponseWriter, request *http.Request) {
		render.JSON(writer, request, TicketResponse{Ticket: "ticket", ExpiresIn: 300})
	})
	mux.Get("/cgi-bin/ticket/get", func(writer http.ResponseWriter, request *http.Request) {
		render.JSON(writer, request, TicketResponse{Ticket: "agent-ticket", ExpiresIn: 300})
	})
	server := httptest.NewServer(mux)
	defer server.Close()

	client := NewClient(Conf{
		CorpID:     "CorpID",
		AgentID:    "AgentID",
		CorpSecret: "CorpSecret",
	})
	// client.Request.Options = append(client.Request.Options, req.DebugHook(nil))
	client.Request.BaseURL = server.URL
	token, err := client.AccessToken()
	assert.NoError(t, err)
	assert.Equal(t, "token", token)
	ticket, err := client.JsAPITicket()
	assert.NoError(t, err)
	assert.Equal(t, "ticket", ticket)
	agentTicket, err := client.AgentTicket()
	assert.NoError(t, err)
	assert.Equal(t, "agent-ticket", agentTicket)

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
}
