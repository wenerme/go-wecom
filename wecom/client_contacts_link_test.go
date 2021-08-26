package wecom

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestContactsLinkSerialization(t *testing.T) {
	ts := NewTestServer()
	handleTokens(ts)
	handleMockData(ts)
	defer ts.Start()()
	c := ts.Client

	{
		request := LinkSimpleListUserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/linkedcorp/user/simplelist.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.LinkSimpleListUser(&request)
		assert.NoError(t, err)
		_ = res

		response := LinkSimpleListUserResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/linkedcorp/user/simplelist.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := LinkListUserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/linkedcorp/user/list.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.LinkListUser(&request)
		assert.NoError(t, err)
		_ = res

		response := LinkListUserResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/linkedcorp/user/list.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := LinkListDepartmentRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/linkedcorp/department/list.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.LinkListDepartment(&request)
		assert.NoError(t, err)
		_ = res

		response := LinkListDepartmentResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/linkedcorp/department/list.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := LinkGetUserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/linkedcorp/user/get.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.LinkGetUser(&request)
		assert.NoError(t, err)
		_ = res

		response := LinkGetUserResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/linkedcorp/user/get.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{

		res, err := c.LinkGetPermList()
		assert.NoError(t, err)
		_ = res

		response := LinkGetPermListResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/linkedcorp/agent/get_perm_list.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
}
