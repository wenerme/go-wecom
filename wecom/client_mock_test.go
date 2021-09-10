package wecom

import (
	"encoding/json"
	"fmt"
	mathrand "math/rand"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/go-playground/validator/v10"

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
	mockAgentID       int
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
		mockAgentID:       mathrand.Int(), //nolint:gosec
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
			render.JSON(writer, request, GenericResponse{ErrorCode: 1, ErrorMessage: "invalid secret " + request.RequestURI})
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
			render.JSON(writer, request, GenericResponse{ErrorCode: 1, ErrorMessage: "invalid secret " + request.RequestURI})
			return
		}
		render.JSON(writer, request, ProviderTokenResponse{ProviderAccessToken: ts.mockProviderToken, ExpiresIn: 7200})
	})
}

var _pathToAPIName = map[string]*clientAPI{}

type clientAPI struct {
	Path              string
	Name              string
	Method            reflect.Method
	RequestType       reflect.Type
	ResponseType      reflect.Type
	IsGenericResponse bool
}

var (
	cRef = &Client{}
	cv   = reflect.ValueOf(cRef)
	ct   = cv.Type()
)

func registerClientAPIPath(path string, mn string, method interface{}) {
	if _, ok := _pathToAPIName[path]; ok {
		panic("path already exists: " + path)
	}
	a := &clientAPI{
		Path: path,
		Name: mn,
	}
	m, found := ct.MethodByName(mn)
	if !found {
		panic("method not found: " + mn + " for " + path)
	}
	a.Method = m

	mt := m.Type
	a.ResponseType = mt.Out(0)
	if mt.NumIn() > 1 && mt.In(1).Kind() != reflect.Slice {
		a.RequestType = mt.In(1)
	}
	a.IsGenericResponse = a.ResponseType.Name() == "GenericResponse"

	_pathToAPIName[path] = a
}

func TestGenericSerialization(t *testing.T) {
	ts := NewTestServer()
	handleMockData(ts)
	defer ts.Start()()
	c := ts.Client
	validate := validator.New()

	for _, v := range _pathToAPIName {
		testSerialize(t, c, validate, v)
	}
}

func testSerialize(t *testing.T, c *Client, validate *validator.Validate, a *clientAPI) {
	var request reflect.Value
	var response reflect.Value
	if a.RequestType != nil {
		request = reflect.New(a.RequestType.Elem())
		fn := fmt.Sprintf("./testdata%s.request.json", a.Path)
		data, err := os.ReadFile(fn)
		if err == nil {
			if assert.NoError(t, json.Unmarshal(data, request.Interface())) {
				assert.NoError(t, validate.Struct(request))
			} else {
				fmt.Println("failed request", a.Path, a.Name, "\n", string(data))
				return
			}
		}
	}

	response = reflect.New(a.ResponseType)
	var hasResponse bool
	{
		fn := fmt.Sprintf("./testdata%s.response.json", a.Path)
		data, err := os.ReadFile(fn)
		if err == nil {
			hasResponse = true
			if !assert.NoError(t, json.Unmarshal(data, response.Interface())) {
				fmt.Println("failed response", a.Path, a.Name, "\n", string(data))
				return
			}
		}
	}

	if hasResponse {
		in := []reflect.Value{reflect.ValueOf(c)}
		if a.RequestType != nil {
			in = append(in, request)
		}
		out := a.Method.Func.Call(in)
		res := out[0]
		var err error
		if !out[1].IsNil() {
			err = out[1].Interface().(error)
		}
		if assert.NoError(t, err) {
			assert.Equal(t, response.Elem().Interface(), res.Interface())
		}
	}
}

func TestStats(t *testing.T) {
	n := len(_eventModels)
	for _, v := range _eventChangeModels {
		n += len(v)
	}
	fmt.Printf(` Client stats
apis: %v
events: %v
`,
		len(_pathToAPIName),
		n,
	)
}
