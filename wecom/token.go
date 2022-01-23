package wecom

import "golang.org/x/oauth2"

type (
	Token       = oauth2.Token
	TokenSource = oauth2.TokenSource
)

var ReuseTokenSource = oauth2.ReuseTokenSource

type TokenSourceFunc func() (*Token, error)

func (f TokenSourceFunc) Token() (token *Token, err error) {
	return f()
}

func WrapTokenSource(src TokenSource, f func(TokenSource) (*Token, error)) TokenSource {
	return TokenSourceFunc(func() (*Token, error) {
		return f(src)
	})
}
