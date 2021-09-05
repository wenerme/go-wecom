package wecom

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestContactsUserSerialization(t *testing.T) {
	ts := NewTestServer()
	handleTokens(ts)
	handleMockData(ts)
	defer ts.Start()()
	c := ts.Client

	{
		request := CreateUserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/user/create.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.CreateUser(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := UpdateUserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/user/update.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.UpdateUser(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := DeleteUserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/user/delete.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.DeleteUser(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := BatchDeleteUserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/user/batchdelete.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.BatchDeleteUser(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := SimpleListUserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/user/simplelist.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := SimpleListUserResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/user/simplelist.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.SimpleListUser(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := ListUserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/user/list.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := ListUserResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/user/list.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.ListUser(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := ConvertToOpenIDRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/user/convert_to_openid.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := ConvertToOpenIDResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/user/convert_to_openid.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.ConvertToOpenID(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := AuthSuccessRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/user/authsucc.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.AuthSuccess(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := GetJoinQrcodeRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/corp/get_join_qrcode.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := GetJoinQrcodeResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/corp/get_join_qrcode.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.GetJoinQrcode(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := GetActiveStatRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/user/get_active_stat.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		response := GetActiveStatResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/user/get_active_stat.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.GetActiveStat(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
}
