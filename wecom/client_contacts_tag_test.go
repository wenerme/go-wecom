package wecom

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestContactsTagSerialization(t *testing.T) {
	ts := NewTestServer()
	handleTokens(ts)
	handleMockData(ts)
	defer ts.Start()()
	c := ts.Client

	{
		request := CreateTagRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/tag/create.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.CreateTag(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := UpdateTagRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/tag/update.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.UpdateTag(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := DeleteTagRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/tag/delete.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.DeleteTag(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := GetTagRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/tag/get.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.GetTag(&request)
		assert.NoError(t, err)
		_ = res

		response := GetTagResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/tag/get.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := AddTagUsersRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/tag/addtagusers.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.AddTagUsers(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := DeleteTagUsersRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/tag/deltagusers.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.DeleteTagUsers(&request)
		assert.NoError(t, err)
		_ = res

	}
	{

		res, err := c.ListTag()
		assert.NoError(t, err)
		_ = res

		response := ListTagResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/tag/list.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
}
