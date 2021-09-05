package wecom

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestContactsAsyncSerialization(t *testing.T) {
	ts := NewTestServer()
	handleTokens(ts)
	handleMockData(ts)
	defer ts.Start()()
	c := ts.Client

	{
		request := BatchSyncUserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/batch/syncuser.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := BatchSyncUserResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/batch/syncuser.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.BatchSyncUser(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := BatchReplaceUserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/batch/replaceuser.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := BatchReplaceUserResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/batch/replaceuser.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.BatchReplaceUser(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := BatchReplacePartyRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/batch/replaceparty.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := BatchReplacePartyResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/batch/replaceparty.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.BatchReplaceParty(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := BatchGetResultRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/batch/getresult.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := BatchGetResultResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/batch/getresult.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.BatchGetResult(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
}
