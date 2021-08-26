package wecom

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/stretchr/testify/assert"
)

type TestServer struct {
	mockAccessToken   string
	mockJsTicket      string
	mockAgentTicket   string
	mockProviderToken string
	mockCorpID        string
	mockAgentID       string
	mockCorpSecret    string
	Mux               *chi.Mux
	Server            *httptest.Server
	Client            *Client
}

func NewTestServer() *TestServer {
	ts := &TestServer{
		mockAccessToken:   "AccessToken" + createNonce(),
		mockJsTicket:      "JsTicket" + createNonce(),
		mockAgentTicket:   "AgentTicket" + createNonce(),
		mockProviderToken: "ProviderToken" + createNonce(),
		mockCorpID:        "CorpID" + createNonce(),
		mockAgentID:       "AgentID" + createNonce(),
		mockCorpSecret:    "CorpSecret" + createNonce(),
		Mux:               chi.NewMux(),
	}
	ts.Client = NewClient(Conf{
		CorpID:     ts.mockCorpID,
		AgentID:    ts.mockAgentID,
		CorpSecret: ts.mockCorpSecret,
	})
	return ts
}

func (ts *TestServer) RequestAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Query().Get("access_token") != ts.mockAccessToken {
			fmt.Printf("%v %v\n\tInvalid Token", request.Method, request.RequestURI)
			render.JSON(writer, request, GenericResponse{ErrorCode: 1, ErrorMessage: "invalid token"})
			return
		}
		next.ServeHTTP(writer, request)
	})
}

func (ts *TestServer) Start() func() {
	server := httptest.NewServer(ts.Mux)
	ts.Server = server
	ts.Client.Request.BaseURL = server.URL
	return func() {
		server.Close()
	}
}

func handleMockData(ts *TestServer) {
	// ts.Mux.
	ts.Mux.HandleFunc("/cgi-bin/*", func(writer http.ResponseWriter, request *http.Request) {
		param := chi.URLParam(request, "*")

		d, err := os.ReadFile("./testdata/cgi-bin/" + param + ".response.json")
		if err != nil {
			// no response
			if _, err := os.Stat("./testdata/cgi-bin/" + param + ".request.json"); err == nil {
				render.JSON(writer, request, GenericResponse{ErrorMessage: "success"})
				return
			}

			render.JSON(writer, request, GenericResponse{ErrorMessage: err.Error(), ErrorCode: -1})
			return
		}
		render.JSON(writer, request, json.RawMessage(d))
	})
}

func TestMock(t *testing.T) {
	ts := NewTestServer()
	handleTokens(ts)
	handleMockData(ts)
	defer ts.Start()()
	c := ts.Client
	_, err := c.GetExternalContact(&GetExternalContactRequest{})
	assert.NoError(t, err)
}

func handleTokens(ts *TestServer) {
	mux := ts.Mux
	api := mux.With(ts.RequestAccessToken)
	mux.Get("/cgi-bin/gettoken", func(writer http.ResponseWriter, request *http.Request) {
		if request.URL.Query().Get("corpsecret") != ts.mockCorpSecret {
			render.JSON(writer, request, GenericResponse{ErrorCode: 1, ErrorMessage: "invalid secret"})
			return
		}
		render.JSON(writer, request, TokenResponse{AccessToken: ts.mockAccessToken, ExpiresIn: 7200})
	})
	api.Get("/cgi-bin/get_jsapi_ticket", func(writer http.ResponseWriter, request *http.Request) {
		render.JSON(writer, request, TicketResponse{Ticket: ts.mockJsTicket, ExpiresIn: 7200})
	})
	api.Get("/cgi-bin/ticket/get", func(writer http.ResponseWriter, request *http.Request) {
		render.JSON(writer, request, TicketResponse{Ticket: ts.mockAgentTicket, ExpiresIn: 7200})
	})
	mux.Post("/cgi-bin/service/get_provider_token", func(writer http.ResponseWriter, request *http.Request) {
		r := GetProviderTokenRequest{}
		if r.CorpID != ts.mockCorpID && r.ProviderSecret != ts.mockProviderToken {
			render.JSON(writer, request, GenericResponse{ErrorCode: 1, ErrorMessage: "invalid secret"})
			return
		}
		render.JSON(writer, request, ProviderTokenResponse{ProviderAccessToken: ts.mockProviderToken, ExpiresIn: 7200})
	})
}
