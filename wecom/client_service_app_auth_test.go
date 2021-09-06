package wecom

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

//nolint:funlen,gocyclo
func TestServiceAppAuthSerialization(t *testing.T) {
	ts := NewTestServer()
	handleTokens(ts)
	handleMockData(ts)
	defer ts.Start()()
	c := ts.Client
	validate := validator.New()

	{
		request := ProviderGetSuiteTokenRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_suite_token.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		response := SuiteTokenResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_suite_token.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.ProviderGetSuiteToken(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := ProviderGetPreAuthCodeRequest{}

		response := PreAuthCodeResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_pre_auth_code.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.ProviderGetPreAuthCode(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := ProviderSetSessionInfoRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/set_session_info.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		res, err := c.ProviderSetSessionInfo(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := ProviderGetPermanentCodeRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_permanent_code.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		response := ProviderGetPermanentCodeResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_permanent_code.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.ProviderGetPermanentCode(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := ProviderGetAuthInfoRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_auth_info.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		response := ProviderGetAuthInfoResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_auth_info.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.ProviderGetAuthInfo(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := ProviderGetCorpTokenRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_corp_token.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		response := ProviderGetCorpTokenResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_corp_token.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.ProviderGetCorpToken(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := ProviderGetAdminListRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_admin_list.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		response := ProviderGetAdminListResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/get_admin_list.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.ProviderGetAdminList(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
}
