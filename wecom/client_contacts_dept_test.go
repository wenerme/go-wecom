package wecom

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestContactsDeptSerialization(t *testing.T) {
	ts := NewTestServer()
	handleTokens(ts)
	handleMockData(ts)
	defer ts.Start()()
	c := ts.Client

	{
		request := CreateDepartmentRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/department/create.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.CreateDepartment(&request)
		assert.NoError(t, err)
		_ = res

		response := CreateDepartmentResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/department/create.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := UpdateDepartmentRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/department/update.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.UpdateDepartment(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := DeleteDepartmentRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/department/delete.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.DeleteDepartment(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := ListDepartmentRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/department/list.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.ListDepartment(&request)
		assert.NoError(t, err)
		_ = res

		response := ListDepartmentResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/department/list.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
}
