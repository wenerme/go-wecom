package wecom

import (
	"github.com/pkg/errors"
)

// nolint:gosec
const (
	SecretTypeJsAPITicket           = "JsAPITicket"
	SecretTypeAgentTicket           = "AgentTicket"
	SecretTypeSuiteTicket           = "SuiteTicket"
	SecretTypeAccessToken           = "AccessToken"
	SecretTypeAuthCorpAccessToken   = "AuthCorpAccessToken"
	SecretTypeAuthCorpPermanentCode = "AuthCorpPermanentCode"
	SecretTypeProviderAccessToken   = "ProviderAccessToken"
	SecretTypeSuiteAccessToken      = "SuiteAccessToken"
	SecretTypeSuitePreAuthCode      = "SuitePreAuthCode"
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

// JsAPITicket request or return cached JsAPITicket
func (c *Client) JsAPITicket() (string, error) {
	// depends on AccessToken
	t := &GenericSecret{
		OwnerID: c.Conf.CorpID,
		Type:    SecretTypeJsAPITicket,
	}
	if c.Conf.SuiteID != "" && c.Conf.AuthCorpID != "" {
		t.OwnerID = joinIds(c.Conf.SuiteID, c.Conf.AuthCorpID)
	}
	return c.TokenProvider.Refresh(t, func() (OpaqueSecret, error) {
		return c.GetJsAPITicket()
	})
}

// AgentTicket request or return cached AgentTicket
func (c *Client) AgentTicket() (string, error) {
	// depends on AccessToken
	t := &GenericSecret{
		OwnerID: c.Conf.CorpID,
		Type:    SecretTypeAgentTicket,
	}
	if c.Conf.SuiteID != "" && c.Conf.AuthCorpID != "" {
		t.OwnerID = joinIds(c.Conf.SuiteID, c.Conf.AuthCorpID)
	}

	return c.TokenProvider.Refresh(t, func() (OpaqueSecret, error) {
		return c.GetAgentTicket()
	})
}

// AuthCorpAccessToken request or return cached AuthCorpAccessToken
func (c *Client) AuthCorpAccessToken() (string, error) {
	// depends on Suite
	return c.TokenProvider.Refresh(&GenericSecret{
		OwnerID: joinIds(c.Conf.SuiteID, c.Conf.AuthCorpID),
		Type:    SecretTypeAuthCorpAccessToken,
	}, func() (o OpaqueSecret, err error) {
		code := c.Conf.AuthCorpPermanentCode
		if code == "" {
			code, err = c.TokenProvider.Refresh(&GenericSecret{
				OwnerID: joinIds(c.Conf.SuiteID, c.Conf.AuthCorpID),
				Type:    SecretTypeAuthCorpPermanentCode,
			}, func() (OpaqueSecret, error) {
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

// AccessToken request or return cached AccessToken
func (c *Client) AccessToken() (string, error) {
	switch {
	case c.Conf.CorpID != "" && c.Conf.CorpSecret != "":
		return c.TokenProvider.Refresh(&GenericSecret{
			OwnerID: c.Conf.CorpID,
			Type:    SecretTypeAccessToken,
		}, func() (OpaqueSecret, error) {
			return c.GetToken()
		})
	case c.Conf.SuiteID != "" && c.Conf.AuthCorpID != "":
		return c.AuthCorpAccessToken()
	default:
		return "", errors.New("unable to get access token: missing CorpSecret or PermanentCode")
	}
}

// ProviderAccessToken request or return cached ProviderAccessToken
func (c *Client) ProviderAccessToken() (string, error) {
	return c.TokenProvider.Refresh(&GenericSecret{
		OwnerID: c.Conf.CorpID,
		Type:    SecretTypeProviderAccessToken,
	}, func() (OpaqueSecret, error) {
		return c.GetProviderToken()
	})
}

// SuiteAccessToken request or return cached SuiteAccessToken
func (c *Client) SuiteAccessToken() (string, error) {
	return c.TokenProvider.Refresh(&GenericSecret{
		OwnerID: c.Conf.SuiteID,
		Type:    SecretTypeSuiteAccessToken,
	}, func() (OpaqueSecret, error) {
		ticket := c.Conf.SuiteTicket
		if ticket == "" {
			var err error
			ticket, err = c.TokenProvider.Refresh(&GenericSecret{
				OwnerID: c.Conf.SuiteID,
				Type:    SecretTypeSuiteTicket,
			}, func() (OpaqueSecret, error) {
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
