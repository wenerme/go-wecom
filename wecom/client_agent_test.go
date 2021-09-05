package wecom

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestAgentSerialization(t *testing.T) {
	ts := NewTestServer()
	handleTokens(ts)
	handleMockData(ts)
	defer ts.Start()()
	c := ts.Client

	{
		request := GetAgentRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/agent/get.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := GetAgentResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/agent/get.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.GetAgent(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{

		response := ListAgentResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/agent/list.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.ListAgent()
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{

		response := SetWorkbenchTemplateResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/agent/set_workbench_template.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.SetWorkbenchTemplate()
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{

		res, err := c.GetWorkbenchTemplate()
		assert.NoError(t, err)
		_ = res

	}
	{

		res, err := c.SetWorkbenchData()
		assert.NoError(t, err)
		_ = res

	}
}
