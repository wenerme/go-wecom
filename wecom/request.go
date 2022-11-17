package wecom

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/wenerme/go-req"
)

type middlewareOptions struct {
	WithoutAccessToken bool
	Debug              bool
	GetToken           func(c *Client, r *req.Request) (string, string, error)
}

const middlewareOptionsContextKey = contextKey("middlewareOptionsContextKey")

func getMiddlewareOptions(r *req.Request) *middlewareOptions {
	if r.Context == nil {
		r.Context = context.Background()
	}
	var o *middlewareOptions
	v := r.Context.Value(middlewareOptionsContextKey)
	if v == nil {
		o = &middlewareOptions{}
		r.Context = context.WithValue(r.Context, middlewareOptionsContextKey, o)
	} else {
		o = v.(*middlewareOptions)
	}
	return o
}

func getAccessToken(c *Client, r *req.Request) (key string, token string, err error) {
	key = "access_token"
	token, err = c.AccessToken()
	return
}

func requestMiddleware(r *req.Request) (err error) {
	if r.Context == nil {
		return nil
	}
	o := getMiddlewareOptions(r)
	client := FromContext(r.Context)
	if client == nil {
		return errors.New("no wecom.Client in context")
	}
	if o.GetToken == nil {
		o.GetToken = getAccessToken
	}
	var v url.Values
	v, err = req.ValuesOf(r.Query)
	if err != nil {
		return
	}
	if v == nil {
		v = url.Values{}
	}
	r.Query = v

	if o.Debug {
		v.Set("debug", "1")
	}
	if !o.WithoutAccessToken {
		key, val, err := o.GetToken(client, r)
		if err == nil {
			if _, found := v[key]; !found {
				v.Set(key, val)
			}
		}
	}
	return
}

var wecomeMiddleware = req.Hook{
	Name: "WecomMiddleware",
	HandleOption: func(r *req.Request, o interface{}) (bool, error) {
		switch v := o.(type) {
		case func(o *middlewareOptions):
			v(getMiddlewareOptions(r))
		default:
			return false, nil
		}
		return true, nil
	},
	OnResponse: func(r *http.Response) error {
		// todo detect json
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}
		r.Body = io.NopCloser(bytes.NewBuffer(body))
		er := GenericResponse{}
		// skip for now
		if err = json.Unmarshal(body, &er); err != nil {
			if r.StatusCode >= 400 {
				er.ErrorCode = r.StatusCode
				er.ErrorMessage = r.Status
				return er
			}
			return nil
		}
		return er.AsError()
	},
}
