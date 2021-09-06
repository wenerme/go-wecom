package wecom

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestMessageAuditMoreSerialization(t *testing.T) {
	ts := NewTestServer()
	handleTokens(ts)
	handleMockData(ts)
	defer ts.Start()()
	c := ts.Client
	validate := validator.New()

	{
		request := MessageAuditCheckRoomAgreeRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/msgaudit/check_room_agree.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		response := MessageAuditCheckRoomAgreeResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/msgaudit/check_room_agree.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.MessageAuditCheckRoomAgree(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := MessageAuditCheckSingleAgreeRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/msgaudit/check_single_agree.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		response := MessageAuditCheckSingleAgreeResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/msgaudit/check_single_agree.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.MessageAuditCheckSingleAgree(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
}
