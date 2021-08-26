package wecom

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestServiceAppAuthSerialization(t *testing.T) {
	ts := NewTestServer()
	handleTokens(ts)
	handleMockData(ts)
	defer ts.Start()()
	c := ts.Client

	{
		request := GetSuiteTokenRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_suite_token.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := SuiteTokenResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_suite_token.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.GetSuiteToken(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := GetPreAuthCodeRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_pre_auth_code.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := GetPreAuthCodeResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_pre_auth_code.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.GetPreAuthCode(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := SetSessionInfoRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/set_session_info.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.SetSessionInfo(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := GetPermanentCodeRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_permanent_code.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := GetPermanentCodeResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_permanent_code.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.GetPermanentCode(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := GetAuthInfoRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_auth_info.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := GetAuthInfoResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_auth_info.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.GetAuthInfo(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := GetCorpTokenRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_corp_token.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := GetCorpTokenResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_corp_token.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.GetCorpToken(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := GetAdminListRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_admin_list.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := GetAdminListResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_admin_list.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.GetAdminList(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
}
