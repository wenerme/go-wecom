package wecom

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

//nolint:funlen
func TestExternalContactSerialization(t *testing.T) {
	{
		r := GetFollowUserListResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_follow_user_list.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetExternalContactRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetExternalContactResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := BatchGetByUserRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/batch/get_by_user.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := BatchGetByUserResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/batch/get_by_user.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := RemarkExternalContactRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/remark.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := UnionIDToExternalUseridRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/unionid_to_external_userid.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetCorpTagListRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_corp_tag_list.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetCorpTagListResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_corp_tag_list.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := AddCorpTagRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/add_corp_tag.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := AddCorpTagResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/add_corp_tag.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := EditCorpTagRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/edit_corp_tag.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := DelCorpTagRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/del_corp_tag.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := MarkTagExternalContactRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/mark_tag.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := TransferCustomerExternalContactRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/transfer_customer.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := TransferCustomerExternalContactResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/transfer_customer.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := TransferResultExternalContactRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/transfer_result.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := TransferResultExternalContactResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/transfer_result.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetUnassignedListResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_unassigned_list.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := TransferCustomerResignedRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/resigned/transfer_customer.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := TransferCustomerResignedResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/resigned/transfer_customer.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := TransferResultResignedRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/resigned/transfer_result.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := TransferResultResignedResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/resigned/transfer_result.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := TransferGroupChatRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/transfer.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := TransferGroupChatResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/transfer.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := ListGroupChatRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/list.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := ListGroupChatResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/list.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetGroupChatResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/get.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetMomentListRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_list.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetMomentListResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_list.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetMomentTaskRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_task.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetMomentTaskResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_task.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetMomentCustomerListRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_customer_list.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetMomentCustomerListResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_customer_list.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetMomentSendResultRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_send_result.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetMomentSendResultResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_send_result.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetMomentCommentsRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_comments.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetMomentCommentsResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_moment_comments.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := AddMessageTemplateRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/add_msg_template.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := AddMessageTemplateResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/add_msg_template.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetGroupMessageListV2Request{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_groupmsg_list_v2.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetGroupMessageListV2Response{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_groupmsg_list_v2.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetGroupMessageTaskRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_groupmsg_task.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetGroupMessageTaskResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_groupmsg_task.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetGroupMessageSendResultRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_groupmsg_send_result.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetGroupMessageSendResultResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_groupmsg_send_result.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := SendWelcomeMessageRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/send_welcome_msg.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := AddGroupWelcomeTemplateRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/group_welcome_template/add.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := AddGroupWelcomeTemplateResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/group_welcome_template/add.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := EditRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/group_welcome_template/edit.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetGroupWelcomeTemplateRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/group_welcome_template/get.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetGroupWelcomeTemplateResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/group_welcome_template/get.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := DelGroupWelcomeTemplateRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/group_welcome_template/del.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetUserBehaviorDataRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_user_behavior_data.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GetUserBehaviorDataResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/get_user_behavior_data.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GroupChatStatisticRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/statistic.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := GroupChatStatisticResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/statistic.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := StatisticGroupByDayRequest{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/statistic_group_by_day.request.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
	{
		r := StatisticGroupByDayResponse{}
		data, err := os.ReadFile("./testdata/cgi-bin/externalcontact/groupchat/statistic_group_by_day.response.json")
		assert.NoError(t, err)
		assert.NoError(t, json.Unmarshal(data, &r))
	}
}
