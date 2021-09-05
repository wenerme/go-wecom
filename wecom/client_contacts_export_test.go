package wecom

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestContactsExportSerialization(t *testing.T) {
	ts := NewTestServer()
	handleTokens(ts)
	handleMockData(ts)
	defer ts.Start()()
	c := ts.Client

	{
		request := ExportSimpleUserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/export/simple_user.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := ExportSimpleUserResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/export/simple_user.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.ExportSimpleUser(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := ExportUserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/export/user.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.ExportUser(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := ExportDepartmentRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/export/department.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := ExportDepartmentResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/export/department.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.ExportDepartment(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := ExportTagUserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/export/taguser.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := ExportTagUserResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/export/taguser.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.ExportTagUser(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
}