package wecom

import (
	"github.com/pkg/errors"
)

type TokenType string

const (
	TokenTypeJsAPITicket           = "JsAPITicket"
	TokenTypeAgentTicket           = "AgentTicket"
	TokenTypeSuiteTicket           = "SuiteTicket"
	TokenTypeAccessToken           = "AccessToken"
	TokenTypeAuthCorpAccessToken   = "AuthCorpAccessToken"
	TokenTypeAuthCorpPermanentCode = "AuthCorpPermanentCode" // nolint:gosec
	TokenTypeProviderAccessToken   = "ProviderAccessToken"
	TokenTypeSuiteAccessToken      = "SuiteAccessToken"
)

func (c *Client) JsAPITicket() (string, error) {
	// depends on AccessToken
	return c.TokenProvider.Refresh(&GenericToken{
		OwnerID: c.Conf.CorpID + c.Conf.AuthCorpID,
		Depends: c.Conf.SuiteID,
		Type:    TokenTypeJsAPITicket,
	}, func() (OpaqueToken, error) {
		return c.GetJsAPITicket()
	})
}

func (c *Client) AgentTicket() (string, error) {
	// depends on AccessToken
	return c.TokenProvider.Refresh(&GenericToken{
		OwnerID: c.Conf.CorpID + c.Conf.AuthCorpID,
		Depends: c.Conf.SuiteID,
		Type:    TokenTypeAgentTicket,
	}, func() (OpaqueToken, error) {
		return c.GetAgentTicket()
	})
}

func (c *Client) AccessToken() (string, error) {
	switch {
	case c.Conf.CorpID != "":
		return c.TokenProvider.Refresh(&GenericToken{
			OwnerID: c.Conf.CorpID,
			Type:    TokenTypeAccessToken,
		}, func() (OpaqueToken, error) {
			return c.GetToken()
		})
	case c.Conf.AuthCorpID != "":
		// depends on Suite
		return c.TokenProvider.Refresh(&GenericToken{
			OwnerID: c.Conf.AuthCorpID,
			Depends: c.Conf.SuiteID,
			Type:    TokenTypeAuthCorpAccessToken,
		}, func() (OpaqueToken, error) {
			code := c.Conf.PermanentCode
			if code == "" {
				var err error
				code, err = c.TokenProvider.Refresh(&GenericToken{
					OwnerID: c.Conf.AuthCorpID,
					Depends: c.Conf.SuiteID,
					Type:    TokenTypeAuthCorpPermanentCode,
				}, func() (OpaqueToken, error) {
					return nil, errors.New("missing auth corp permanent code")
				})
				if err != nil {
					return nil, err
				}
			}
			o, err := c.ProviderGetCorpToken(&ProviderGetCorpTokenRequest{
				AuthCorpID:    c.Conf.AuthCorpID,
				PermanentCode: code,
			})
			return TokenResponse(o), err
		})
	default:
		return "", errors.New("unable to get access token: missing CorpSecret or PermanentCode")
	}
}

func (c *Client) ProviderAccessToken() (string, error) {
	return c.TokenProvider.Refresh(&GenericToken{
		OwnerID: c.Conf.CorpID,
		Type:    TokenTypeProviderAccessToken,
	}, func() (OpaqueToken, error) {
		return c.GetProviderToken()
	})
}

func (c *Client) SuiteAccessToken() (string, error) {
	return c.TokenProvider.Refresh(&GenericToken{
		OwnerID: c.Conf.SuiteID,
		Type:    TokenTypeSuiteAccessToken,
	}, func() (OpaqueToken, error) {
		ticket := c.Conf.SuiteTicket
		if ticket == "" {
			var err error
			ticket, err = c.TokenProvider.Refresh(&GenericToken{
				OwnerID: c.Conf.SuiteID,
				Type:    TokenTypeSuiteTicket,
			}, func() (OpaqueToken, error) {
				return nil, errors.New("missing suite ticket")
			})
			if err != nil {
				return nil, err
			}
		}
		return c.ProviderGetSuiteToken(&ProviderGetSuiteTokenRequest{
			SuiteID:     c.Conf.SuiteID,
			SuiteSecret: c.Conf.SuiteSecret,
			SuiteTicket: ticket,
		})
	})
}
