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
func TestMessageAuditSerialization(t *testing.T) {
	ts := NewTestServer()
	handleTokens(ts)
	handleMockData(ts)
	defer ts.Start()()
	c := ts.Client
	validate := validator.New()

	{
		request := MessageAuditGetPermitUserListRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/msgaudit/get_permit_user_list.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		response := MessageAuditGetPermitUserListResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/msgaudit/get_permit_user_list.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.MessageAuditGetPermitUserList(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := MessageAuditGetGroupChatRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/msgaudit/groupchat/get.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		response := MessageAuditGetGroupChatResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/msgaudit/groupchat/get.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.MessageAuditGetGroupChat(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
}
