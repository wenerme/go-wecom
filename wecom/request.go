package wecom

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/wenerme/go-req"
)

type middlewareOptions struct {
	WithoutAccessToken bool
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

func requestMiddleware(r *req.Request) (err error) {
	if r.Context == nil {
		return nil
	}
	o := getMiddlewareOptions(r)
	client := FromContext(r.Context)
	if client == nil {
		return errors.New("no wecom.Client in context")
	}
	if !o.WithoutAccessToken {
		var v url.Values
		v, err = req.ValuesOf(r.Query)
		if v == nil {
			v = url.Values{}
		}
		r.Query = v

		_, found := v["access_token"]
		if err == nil && !found {
			var token string
			token, err = client.AccessToken()
			if err == nil {
				v.Set("access_token", token)
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
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return err
		}
		r.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		er := GenericResponse{}
		// skip for now
		if err = json.Unmarshal(body, &er); err != nil {
			return nil
		}
		return er.AsError()
	},
}
