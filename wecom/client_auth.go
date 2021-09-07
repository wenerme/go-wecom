package wecom

import (
	"github.com/pkg/errors"
)

func (c *Client) JsAPITicket() (string, error) {
	// depends on AccessToken
	return c.TokenProvider.Refresh(&GenericToken{
		OwnerID: c.Conf.CorpID + c.Conf.AuthCorpID,
		Depends: c.Conf.SuiteID,
		Name:    "JsAPITicket",
	}, func() (Token, error) {
		return c.GetJsAPITicket()
	})
}

func (c *Client) AgentTicket() (string, error) {
	// depends on AccessToken
	return c.TokenProvider.Refresh(&GenericToken{
		OwnerID: c.Conf.CorpID + c.Conf.AuthCorpID,
		Depends: c.Conf.SuiteID,
		Name:    "AgentTicket",
	}, func() (Token, error) {
		return c.GetAgentTicket()
	})
}

func (c *Client) AccessToken() (string, error) {
	switch {
	case c.Conf.CorpSecret != "":
		return c.TokenProvider.Refresh(&GenericToken{
			OwnerID: c.Conf.CorpID,
			Name:    "AccessToken",
		}, func() (Token, error) {
			return c.GetToken()
		})
	case c.Conf.PermanentCode != "":
		// depends on Suite
		return c.TokenProvider.Refresh(&GenericToken{
			OwnerID: c.Conf.AuthCorpID,
			Depends: c.Conf.SuiteID,
			Name:    "AuthCorpAccessToken",
		}, func() (Token, error) {
			return c.GetAuthCorpAccessToken()
		})
	default:
		return "", errors.New("unable to get access token: missing CorpSecret or PermanentCode")
	}
}

func (c *Client) ProviderAccessToken() (string, error) {
	return c.TokenProvider.Refresh(&GenericToken{
		OwnerID: c.Conf.CorpID,
		Name:    "ProviderAccessToken",
	}, func() (Token, error) {
		return c.GetProviderToken()
	})
}

func (c *Client) SuiteAccessToken() (string, error) {
	return c.TokenProvider.Refresh(&GenericToken{
		OwnerID: c.Conf.SuiteID,
		Name:    "SuiteAccessTokenCache",
	}, func() (Token, error) {
		return c.ProviderGetSuiteToken(&ProviderGetSuiteTokenRequest{
			SuiteID:     c.Conf.SuiteID,
			SuiteSecret: c.Conf.SuiteSecret,
			SuiteTicket: c.Conf.SuiteTicket,
		})
	})
}
