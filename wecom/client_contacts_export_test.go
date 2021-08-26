package wecom

import (
	"encoding/json"
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
		request := SimpleUserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/export/simple_user.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.SimpleUser(&request)
		assert.NoError(t, err)
		_ = res

		response := SimpleUserResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/export/simple_user.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := UserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/export/user.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.User(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := DepartmentRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/export/department.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.Department(&request)
		assert.NoError(t, err)
		_ = res

		response := DepartmentResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/export/department.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := TaguserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/export/taguser.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.Taguser(&request)
		assert.NoError(t, err)
		_ = res

		response := TaguserResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/export/taguser.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
}
