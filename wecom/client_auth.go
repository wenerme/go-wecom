package wecom

import (
	"github.com/pkg/errors"
)

type TokenType string

// nolint:gosec
const (
	TokenTypeJsAPITicket           = "JsAPITicket"
	TokenTypeAgentTicket           = "AgentTicket"
	TokenTypeSuiteTicket           = "SuiteTicket"
	TokenTypeAccessToken           = "AccessToken"
	TokenTypeAuthCorpAccessToken   = "AuthCorpAccessToken"
	TokenTypeAuthCorpPermanentCode = "AuthCorpPermanentCode"
	TokenTypeProviderAccessToken   = "ProviderAccessToken"
	TokenTypeSuiteAccessToken      = "SuiteAccessToken"
	TokenTypeSuitePreAuthCode      = "SuitePreAuthCode"
)

func joinIds(s ...string) string {
	o := ""
	for _, v := range s {
		if v != "" {
			if o != "" {
				o += "."
			}
			o += v
		}
	}
	return o
}

func (c *Client) JsAPITicket() (string, error) {
	// depends on AccessToken
	t := &GenericToken{
		OwnerID: c.Conf.CorpID,
		Type:    TokenTypeJsAPITicket,
	}
	if c.Conf.SuiteID != "" && c.Conf.AuthCorpID != "" {
		t.OwnerID = joinIds(c.Conf.SuiteID, c.Conf.AuthCorpID)
	}
	return c.TokenProvider.Refresh(t, func() (OpaqueToken, error) {
		return c.GetJsAPITicket()
	})
}

func (c *Client) AgentTicket() (string, error) {
	// depends on AccessToken
	t := &GenericToken{
		OwnerID: c.Conf.CorpID,
		Type:    TokenTypeAgentTicket,
	}
	if c.Conf.SuiteID != "" && c.Conf.AuthCorpID != "" {
		t.OwnerID = joinIds(c.Conf.SuiteID, c.Conf.AuthCorpID)
	}

	return c.TokenProvider.Refresh(t, func() (OpaqueToken, error) {
		return c.GetAgentTicket()
	})
}

func (c *Client) AuthCorpAccessToken() (string, error) {
	// depends on Suite
	return c.TokenProvider.Refresh(&GenericToken{
		OwnerID: joinIds(c.Conf.SuiteID, c.Conf.AuthCorpID),
		Type:    TokenTypeAuthCorpAccessToken,
	}, func() (o OpaqueToken, err error) {
		code := c.Conf.AuthCorpPermanentCode
		if code == "" {
			code, err = c.TokenProvider.Refresh(&GenericToken{
				OwnerID: joinIds(c.Conf.SuiteID, c.Conf.AuthCorpID),
				Type:    TokenTypeAuthCorpPermanentCode,
			}, func() (OpaqueToken, error) {
				return nil, errors.New("missing auth corp permanent code")
			})
		}
		if err == nil {
			o, err = c.ProviderGetCorpToken(&ProviderGetCorpTokenRequest{
				AuthCorpID:    c.Conf.AuthCorpID,
				PermanentCode: code,
			})
		}
		return
	})
}

func (c *Client) AccessToken() (string, error) {
	switch {
	case c.Conf.CorpID != "" && c.Conf.CorpSecret != "":
		return c.TokenProvider.Refresh(&GenericToken{
			OwnerID: c.Conf.CorpID,
			Type:    TokenTypeAccessToken,
		}, func() (OpaqueToken, error) {
			return c.GetToken()
		})
	case c.Conf.SuiteID != "" && c.Conf.AuthCorpID != "":
		return c.AuthCorpAccessToken()
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
