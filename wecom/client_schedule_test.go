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
func TestScheduleSerialization(t *testing.T) {
	ts := NewTestServer()
	handleTokens(ts)
	handleMockData(ts)
	defer ts.Start()()
	c := ts.Client
	validate := validator.New()

	{
		request := AddCalendarRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/oa/calendar/add.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		response := AddCalendarResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/oa/calendar/add.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.AddCalendar(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := UpdateCalendarRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/oa/calendar/update.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		res, err := c.UpdateCalendar(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := GetCalendarRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/oa/calendar/get.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		response := GetCalendarResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/oa/calendar/get.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.GetCalendar(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := DeleteCalendarRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/oa/calendar/del.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		res, err := c.DeleteCalendar(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := AddScheduleRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/oa/schedule/add.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		response := AddScheduleResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/oa/schedule/add.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.AddSchedule(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := UpdateScheduleRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/oa/schedule/update.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		res, err := c.UpdateSchedule(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := GetScheduleRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/oa/schedule/get.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		response := GetScheduleResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/oa/schedule/get.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.GetSchedule(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
	{
		request := DeleteScheduleRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/oa/schedule/del.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		res, err := c.DeleteSchedule(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := ScheduleGetByCalendarRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/oa/schedule/get_by_calendar.request.json")
			assert.NoError(t, err)
			if assert.NoError(t, json.Unmarshal(data, &request)) {
				assert.NoError(t, validate.Struct(request))
			}
		}

		response := ScheduleGetByCalendarResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/oa/schedule/get_by_calendar.response.json")
			assert.NoError(t, err)
			if !assert.NoError(t, json.Unmarshal(data, &response)) {
				fmt.Println(string(data))
			}
		}

		res, err := c.ScheduleGetByCalendar(&request)
		assert.NoError(t, err)
		_ = res

		assert.Equal(t, response, res)
	}
}
