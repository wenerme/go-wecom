package wecom

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestExternalContactSerialization(t *testing.T) {
	ts := NewTestServer()
	handleTokens(ts)
	handleMockData(ts)
	defer ts.Start()()
	c := ts.Client

	{

		res, err := c.GetFollowUserList()
		assert.NoError(t, err)
		_ = res

		response := GetFollowUserListResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_follow_user_list.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := AddContactWayRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/add_contact_way.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.AddContactWay(&request)
		assert.NoError(t, err)
		_ = res

		response := AddContactWayResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/add_contact_way.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := GetContactWayRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_contact_way.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.GetContactWay(&request)
		assert.NoError(t, err)
		_ = res

		response := GetContactWayResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_contact_way.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := UpdateContactWayRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/update_contact_way.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.UpdateContactWay(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := DeleteContactWayRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/del_contact_way.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.DeleteContactWay(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := CloseTempChatRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/close_temp_chat.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.CloseTempChat(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := ListExternalContactRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/list.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.ListExternalContact(&request)
		assert.NoError(t, err)
		_ = res

		response := ListExternalContactResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/list.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := GetExternalContactRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.GetExternalContact(&request)
		assert.NoError(t, err)
		_ = res

		response := GetExternalContactResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := BatchGetByUserRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/batch/get_by_user.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.BatchGetByUser(&request)
		assert.NoError(t, err)
		_ = res

		response := BatchGetByUserResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/batch/get_by_user.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := RemarkExternalContactRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/remark.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.RemarkExternalContact(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := UnionIDToExternalUserIDRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/unionid_to_external_userid.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.UnionIDToExternalUserID(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := GetCorpTagListRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_corp_tag_list.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.GetCorpTagList(&request)
		assert.NoError(t, err)
		_ = res

		response := GetCorpTagListResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_corp_tag_list.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := AddCorpTagRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/add_corp_tag.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.AddCorpTag(&request)
		assert.NoError(t, err)
		_ = res

		response := AddCorpTagResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/add_corp_tag.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := EditCorpTagRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/edit_corp_tag.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.EditCorpTag(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := DeleteCorpTagRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/del_corp_tag.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.DeleteCorpTag(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := MarkTagExternalContactRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/mark_tag.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.MarkTagExternalContact(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := TransferCustomerExternalContactRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/transfer_customer.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.TransferCustomerExternalContact(&request)
		assert.NoError(t, err)
		_ = res

		response := TransferCustomerExternalContactResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/transfer_customer.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := TransferResultExternalContactRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/transfer_result.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.TransferResultExternalContact(&request)
		assert.NoError(t, err)
		_ = res

		response := TransferResultExternalContactResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/transfer_result.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := GetUnassignedListRequest{}

		res, err := c.GetUnassignedList(&request)
		assert.NoError(t, err)
		_ = res

		response := GetUnassignedListResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_unassigned_list.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := TransferCustomerResignedRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/resigned/transfer_customer.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.TransferCustomerResigned(&request)
		assert.NoError(t, err)
		_ = res

		response := TransferCustomerResignedResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/resigned/transfer_customer.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := TransferResultResignedRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/resigned/transfer_result.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.TransferResultResigned(&request)
		assert.NoError(t, err)
		_ = res

		response := TransferResultResignedResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/resigned/transfer_result.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := TransferGroupChatRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/transfer.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.TransferGroupChat(&request)
		assert.NoError(t, err)
		_ = res

		response := TransferGroupChatResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/transfer.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := ListGroupChatRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/list.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.ListGroupChat(&request)
		assert.NoError(t, err)
		_ = res

		response := ListGroupChatResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/list.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{

		res, err := c.GetGroupChat()
		assert.NoError(t, err)
		_ = res

		response := GetGroupChatResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/get.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := GetMomentListRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_list.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.GetMomentList(&request)
		assert.NoError(t, err)
		_ = res

		response := GetMomentListResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_list.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := GetMomentTaskRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_task.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.GetMomentTask(&request)
		assert.NoError(t, err)
		_ = res

		response := GetMomentTaskResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_task.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := GetMomentCustomerListRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_customer_list.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.GetMomentCustomerList(&request)
		assert.NoError(t, err)
		_ = res

		response := GetMomentCustomerListResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_customer_list.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := GetMomentSendResultRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_send_result.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.GetMomentSendResult(&request)
		assert.NoError(t, err)
		_ = res

		response := GetMomentSendResultResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_send_result.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := GetMomentCommentsRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_comments.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.GetMomentComments(&request)
		assert.NoError(t, err)
		_ = res

		response := GetMomentCommentsResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_comments.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := AddMessageTemplateRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/add_msg_template.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.AddMessageTemplate(&request)
		assert.NoError(t, err)
		_ = res

		response := AddMessageTemplateResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/add_msg_template.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := GetGroupMessageListV2Request{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_groupmsg_list_v2.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.GetGroupMessageListV2(&request)
		assert.NoError(t, err)
		_ = res

		response := GetGroupMessageListV2Response{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_groupmsg_list_v2.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := GetGroupMessageTaskRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_groupmsg_task.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.GetGroupMessageTask(&request)
		assert.NoError(t, err)
		_ = res

		response := GetGroupMessageTaskResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_groupmsg_task.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := GetGroupMessageSendResultRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_groupmsg_send_result.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.GetGroupMessageSendResult(&request)
		assert.NoError(t, err)
		_ = res

		response := GetGroupMessageSendResultResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_groupmsg_send_result.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := SendWelcomeMessageRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/send_welcome_msg.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.SendWelcomeMessage(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := AddGroupWelcomeTemplateRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/group_welcome_template/add.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.AddGroupWelcomeTemplate(&request)
		assert.NoError(t, err)
		_ = res

		response := AddGroupWelcomeTemplateResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/group_welcome_template/add.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := EditGroupWelcomeTemplateRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/group_welcome_template/edit.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.EditGroupWelcomeTemplate(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := GetGroupWelcomeTemplateRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/group_welcome_template/get.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.GetGroupWelcomeTemplate(&request)
		assert.NoError(t, err)
		_ = res

		response := GetGroupWelcomeTemplateResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/group_welcome_template/get.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := DeleteGroupWelcomeTemplateRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/group_welcome_template/del.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.DeleteGroupWelcomeTemplate(&request)
		assert.NoError(t, err)
		_ = res

	}
	{
		request := GetUserBehaviorDataRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_user_behavior_data.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.GetUserBehaviorData(&request)
		assert.NoError(t, err)
		_ = res

		response := GetUserBehaviorDataResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_user_behavior_data.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := GroupChatStatisticRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/statistic.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.GroupChatStatistic(&request)
		assert.NoError(t, err)
		_ = res

		response := GroupChatStatisticResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/statistic.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
	{
		request := StatisticGroupByDayRequest{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/statistic_group_by_day.request.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &request))
		}

		res, err := c.StatisticGroupByDay(&request)
		assert.NoError(t, err)
		_ = res

		response := StatisticGroupByDayResponse{}
		{
			data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/statistic_group_by_day.response.json")
			assert.NoError(t, err)
			assert.NoError(t, json.Unmarshal(data, &response))
		}
		assert.Equal(t, response, res)
	}
}
