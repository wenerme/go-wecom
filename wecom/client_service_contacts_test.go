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
func TestServiceContactsSerialization(t *testing.T) {
	ts := NewTestServer()
	handleTokens(ts)
	handleMockData(ts)
	defer ts.Start()()
	c := ts.Client
	validate := validator.New()

	{
		request := ProviderIDTranslateContactRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/contact/id_translate.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		response := ProviderIDTranslateContactResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/contact/id_translate.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.ProviderIDTranslateContact(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := ProviderBatchGetResultRequest{}

		response := ProviderBatchGetResultResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/service/batch/getresult.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.ProviderBatchGetResult(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
}
