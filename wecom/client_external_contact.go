package wecom

import (
	"github.com/wenerme/go-req"
)

// GetFollowUserList 获取配置了客户联系功能的成员列表
// 企业和第三方服务商可通过此接口获取配置了客户联系功能的成员列表。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92576
func (c *Client) GetFollowUserList(opts ...interface{}) (out GetFollowUserListResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/externalcontact/get_follow_user_list",
		Options: opts,
	}).Fetch(&out)
	return
}

// AddContactWay 配置客户联系「联系我」方式
// 企业可以在管理后台-客户联系-加客户中配置成员的「联系我」的二维码或者小程序按钮，客户通过扫描二维码或点击小程序上的按钮，即可获取成员联系方式，主动联系到成员。企业可通过此接口为具有客户联系功能的成员生成专属的「联系我」二维码或者「联系我」按钮。如果配置的是「联系我」按钮，需要开发者的小程序接入小程序插件。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92577
func (c *Client) AddContactWay(r *AddContactWayRequest, opts ...interface{}) (out AddContactWayResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/add_contact_way",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetContactWay 获取企业已配置的「联系我」方式
// 获取企业配置的「联系我」二维码和「联系我」小程序按钮。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92577
func (c *Client) GetContactWay(r *GetContactWayRequest, opts ...interface{}) (out GetContactWayResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/get_contact_way",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateContactWay 更新企业已配置的「联系我」方式
// 更新企业配置的「联系我」二维码和「联系我」小程序按钮中的信息，如使用人员和备注等。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92577
func (c *Client) UpdateContactWay(r *UpdateContactWayRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/update_contact_way",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteContactWay 删除企业已配置的「联系我」方式
// 删除一个已配置的「联系我」二维码或者「联系我」小程序按钮。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92577
func (c *Client) DeleteContactWay(r *DeleteContactWayRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/del_contact_way",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CloseTempChat 结束临时会话
// 将指定的企业成员和客户之前的临时会话断开，断开前会自动下发已配置的结束语。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92577
func (c *Client) CloseTempChat(r *CloseTempChatRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/close_temp_chat",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListExternalContact 获取客户列表
// 企业可通过此接口获取指定成员添加的客户列表。客户是指配置了客户联系功能的成员所添加的外部联系人。没有配置客户联系功能的成员，所添加的外部联系人将不会作为客户返回。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92264
func (c *Client) ListExternalContact(r *ListExternalContactRequest, opts ...interface{}) (out ListExternalContactResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/externalcontact/list",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetExternalContact 获取客户详情
// 企业可通过此接口，根据外部联系人的userid（如何获取?），拉取客户详情。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92265
func (c *Client) GetExternalContact(r *GetExternalContactRequest, opts ...interface{}) (out GetExternalContactResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "GET",
		URL:     "/cgi-bin/externalcontact/get",
		Query:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchGetByUser 批量获取客户详情
// 企业/第三方可通过此接口获取指定成员添加的客户信息列表。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93010
func (c *Client) BatchGetByUser(r *BatchGetByUserRequest, opts ...interface{}) (out BatchGetByUserResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/batch/get_by_user",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// RemarkExternalContact 修改客户备注信息
// 企业可通过此接口修改指定用户添加的客户的备注信息。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92694
func (c *Client) RemarkExternalContact(r *RemarkExternalContactRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/remark",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UnionIDToExternalUserID 外部联系人unionid转换
// 服务商为企业代开发微信小程序的场景，服务商可通过此接口，将微信客户的unionid转为external_userid（如何获取?）。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93274
func (c *Client) UnionIDToExternalUserID(r *UnionIDToExternalUserIDRequest, opts ...interface{}) (out UnionIDToExternalUserIDResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/unionid_to_external_userid",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetCorpTagList 获取企业标签库
// 企业可通过此接口获取企业客户标签详情。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92696
func (c *Client) GetCorpTagList(r *GetCorpTagListRequest, opts ...interface{}) (out GetCorpTagListResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/get_corp_tag_list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddCorpTag 添加企业客户标签
// 企业可通过此接口向客户标签库中添加新的标签组和标签，每个企业最多可配置3000个企业标签。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92696
func (c *Client) AddCorpTag(r *AddCorpTagRequest, opts ...interface{}) (out AddCorpTagResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/add_corp_tag",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// EditCorpTag 编辑企业客户标签
// 企业可通过此接口编辑客户标签/标签组的名称或次序值。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92696
func (c *Client) EditCorpTag(r *EditCorpTagRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/edit_corp_tag",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteCorpTag 删除企业客户标签
// 企业可通过此接口删除客户标签库中的标签，或删除整个标签组。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92696
func (c *Client) DeleteCorpTag(r *DeleteCorpTagRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/del_corp_tag",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// MarkTagExternalContact 编辑客户企业标签
// 企业可通过此接口为指定成员的客户添加上由企业统一配置的标签。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92697
func (c *Client) MarkTagExternalContact(r *MarkTagExternalContactRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/mark_tag",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// TransferCustomerExternalContact 分配在职成员的客户
// 企业可通过此接口，转接在职成员的客户给其他成员。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/94096
func (c *Client) TransferCustomerExternalContact(r *TransferCustomerExternalContactRequest, opts ...interface{}) (out TransferCustomerExternalContactResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/transfer_customer",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// TransferResultExternalContact 查询客户接替状态
// 企业和第三方可通过此接口查询在职成员的客户转接情况。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/94097
func (c *Client) TransferResultExternalContact(r *TransferResultExternalContactRequest, opts ...interface{}) (out TransferResultExternalContactResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/transfer_result",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetUnassignedList 获取待分配的离职成员列表
// 企业和第三方可通过此接口，获取所有离职成员的客户列表，并可进一步调用分配离职成员的客户接口将这些客户重新分配给其他企业成员。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92273
func (c *Client) GetUnassignedList(r *GetUnassignedListRequest, opts ...interface{}) (out GetUnassignedListResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/get_unassigned_list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// TransferCustomerResigned 分配离职成员的客户
// 企业可通过此接口，分配离职成员的客户给其他成员。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/94100
func (c *Client) TransferCustomerResigned(r *TransferCustomerResignedRequest, opts ...interface{}) (out TransferCustomerResignedResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/resigned/transfer_customer",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// TransferResultResigned 查询客户接替状态
// 企业和第三方可通过此接口查询离职成员的客户分配情况。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/94101
func (c *Client) TransferResultResigned(r *TransferResultResignedRequest, opts ...interface{}) (out TransferResultResignedResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/resigned/transfer_result",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// TransferGroupChat 分配离职成员的客户群
// 企业可通过此接口，将已离职成员为群主的群，分配给另一个客服成员。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93242
func (c *Client) TransferGroupChat(r *TransferGroupChatRequest, opts ...interface{}) (out TransferGroupChatResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/groupchat/transfer",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListGroupChat 获取客户群列表
// 该接口用于获取配置过客户群管理的客户群列表。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93414
func (c *Client) ListGroupChat(r *ListGroupChatRequest, opts ...interface{}) (out ListGroupChatResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/groupchat/list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetGroupChat 获取客户群详情
// 通过客户群ID，获取详情。包括群名、群成员列表、群成员入群时间、入群方式。（客户群是由具有客户群使用权限的成员创建的外部群）
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92707
func (c *Client) GetGroupChat(r *GetGroupChatRequest, opts ...interface{}) (out GetGroupChatResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/groupchat/get",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMomentList 获取企业全部的发表列表
// 企业和第三方应用可通过该接口获取企业全部的发表内容。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93443
func (c *Client) GetMomentList(r *GetMomentListRequest, opts ...interface{}) (out GetMomentListResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/get_moment_list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMomentTask 获取客户朋友圈企业发表的列表
// 企业和第三方应用可通过该接口获取企业发表的朋友圈成员执行情况
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93443
func (c *Client) GetMomentTask(r *GetMomentTaskRequest, opts ...interface{}) (out GetMomentTaskResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/get_moment_task",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMomentCustomerList 获取客户朋友圈发表时选择的可见范围
// 企业和第三方应用可通过该接口获取客户朋友圈创建时，选择的客户可见范围
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93443
func (c *Client) GetMomentCustomerList(r *GetMomentCustomerListRequest, opts ...interface{}) (out GetMomentCustomerListResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/get_moment_customer_list",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMomentSendResult 获取客户朋友圈发表后的可见客户列表
// 企业和第三方应用可通过该接口获取客户朋友圈发表后，可在微信朋友圈中查看的客户列表
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93443
func (c *Client) GetMomentSendResult(r *GetMomentSendResultRequest, opts ...interface{}) (out GetMomentSendResultResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/get_moment_send_result",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMomentComments 获取客户朋友圈的互动数据
// 企业和第三方应用可通过此接口获取客户朋友圈的互动数据。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93443
func (c *Client) GetMomentComments(r *GetMomentCommentsRequest, opts ...interface{}) (out GetMomentCommentsResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/get_moment_comments",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddMessageTemplate 创建企业群发
// 企业跟第三方应用可通过此接口添加企业群发消息的任务并通知成员发送给相关客户或客户群。（注：企业微信终端需升级到2.7.5版本及以上）注意：调用该接口并不会直接发送消息给客户/客户群，需要成员确认后才会执行发送（客服人员的企业微信需要升级到2.7.5及以上版本）旧接口创建企业群发已经废弃，接口升级后支持发送视频文件，并且支持最多同时发送9个附件。同一个企业每个自然月内仅可针对一个客户/客户群发送4条消息，超过接收上限的客户将无法再收到群发消息。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92698
func (c *Client) AddMessageTemplate(r *AddMessageTemplateRequest, opts ...interface{}) (out AddMessageTemplateResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/add_msg_template",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetGroupMessageListV2 获取群发记录列表
// 企业和第三方应用可通过此接口获取企业与成员的群发记录。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93439
func (c *Client) GetGroupMessageListV2(r *GetGroupMessageListV2Request, opts ...interface{}) (out GetGroupMessageListV2Response, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/get_groupmsg_list_v2",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetGroupMessageTask 获取群发成员发送任务列表
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93439
func (c *Client) GetGroupMessageTask(r *GetGroupMessageTaskRequest, opts ...interface{}) (out GetGroupMessageTaskResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/get_groupmsg_task",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetGroupMessageSendResult 获取企业群发成员执行结果
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93439
func (c *Client) GetGroupMessageSendResult(r *GetGroupMessageSendResultRequest, opts ...interface{}) (out GetGroupMessageSendResultResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/get_groupmsg_send_result",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SendWelcomeMessage 发送新客户欢迎语
// 企业微信在向企业推送添加外部联系人事件时，会额外返回一个welcome_code，企业以此为凭据调用接口，即可通过成员向新添加的客户发送个性化的欢迎语。为了保证用户体验以及避免滥用，企业仅可在收到相关事件后20秒内调用，且只可调用一次。如果企业已经在管理端为相关成员配置了可用的欢迎语，则推送添加外部联系人事件时不会返回welcome_code。每次添加新客户时可能有多个企业自建应用/第三方应用收到带有welcome_code的回调事件，但仅有最先调用的可以发送成功。后续调用将返回41051（externaluser has started chatting）错误，请用户根据实际使用需求，合理设置应用可见范围，避免冲突。旧接口发送新客户欢迎语已经废弃，接口升级后支持发送视频文件，并且最多支持同时发送9个附件
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92599
func (c *Client) SendWelcomeMessage(r *SendWelcomeMessageRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/send_welcome_msg",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// AddGroupWelcomeTemplate 添加入群欢迎语素材
// 企业可通过此API向企业的入群欢迎语素材库中添加素材。每个企业的入群欢迎语素材库中，最多容纳100个素材。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93438
func (c *Client) AddGroupWelcomeTemplate(r *AddGroupWelcomeTemplateRequest, opts ...interface{}) (out AddGroupWelcomeTemplateResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/group_welcome_template/add",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// EditGroupWelcomeTemplate 编辑入群欢迎语素材
// 企业可通过此API编辑入群欢迎语素材库中的素材，且仅能够编辑调用方自己创建的入群欢迎语素材。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93438
func (c *Client) EditGroupWelcomeTemplate(r *EditGroupWelcomeTemplateRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/group_welcome_template/edit",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetGroupWelcomeTemplate 获取入群欢迎语素材
// 企业可通过此API获取入群欢迎语素材。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93438
func (c *Client) GetGroupWelcomeTemplate(r *GetGroupWelcomeTemplateRequest, opts ...interface{}) (out GetGroupWelcomeTemplateResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/group_welcome_template/get",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// DeleteGroupWelcomeTemplate 删除入群欢迎语素材
// 企业可通过此API删除入群欢迎语素材，且仅能删除调用方自己创建的入群欢迎语素材。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93438
func (c *Client) DeleteGroupWelcomeTemplate(r *DeleteGroupWelcomeTemplateRequest, opts ...interface{}) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/group_welcome_template/del",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetUserBehaviorData 获取「联系客户统计」数据
// 企业可通过此接口获取成员联系客户的数据，包括发起申请数、新增客户数、聊天数、发送消息数和删除/拉黑成员的客户数等指标。
//
// see https://work.weixin.qq.com/api/doc/90001/90143/92275
func (c *Client) GetUserBehaviorData(r *GetUserBehaviorDataRequest, opts ...interface{}) (out GetUserBehaviorDataResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/get_user_behavior_data",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GroupChatStatistic 按群主聚合的方式
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93476
func (c *Client) GroupChatStatistic(r *GroupChatStatisticRequest, opts ...interface{}) (out GroupChatStatisticResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/groupchat/statistic",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// StatisticGroupByDay 按自然日聚合的方式
//
// see https://work.weixin.qq.com/api/doc/90001/90143/93476
func (c *Client) StatisticGroupByDay(r *StatisticGroupByDayRequest, opts ...interface{}) (out StatisticGroupByDayResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalcontact/groupchat/statistic_group_by_day",
		Body:    r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetFollowUserListResponse is response of Client.GetFollowUserList
type GetFollowUserListResponse struct {
	// FollowUser 配置了客户联系功能的成员userid列表
	FollowUser []string `json:"follow_user"  `
}

// AddContactWayRequest is request of Client.AddContactWay
type AddContactWayRequest struct {
	// Type 联系方式类型,1-单人, 2-多人
	Type int `json:"type"  validate:"required"`
	// Scene 场景，1-在小程序中联系，2-通过二维码联系
	Scene int `json:"scene"  validate:"required"`
	// Style 在小程序中联系时使用的控件样式，详见附表
	Style int `json:"style"  `
	// Remark 联系方式的备注信息，用于助记，不超过30个字符
	Remark string `json:"remark"  `
	// SkipVerify 外部客户添加时是否无需验证，默认为true
	SkipVerify bool `json:"skip_verify"  `
	// State 企业自定义的state参数，用于区分不同的添加渠道，在调用“获取外部联系人详情”时会返回该参数值，不超过30个字符
	State string `json:"state"  `
	// User 使用该联系方式的用户userID列表，在type为1时为必填，且只能有一个
	User []string `json:"user"  `
	// Party 使用该联系方式的部门id列表，只在type为2时有效
	Party []int `json:"party"  `
	// IsTemp 是否临时会话模式，true表示使用临时会话模式，默认为false
	IsTemp bool `json:"is_temp"  `
	// ExpiresIn 临时会话二维码有效期，以秒为单位。该参数仅在is_temp为true时有效，默认7天
	ExpiresIn int `json:"expires_in"  `
	// ChatExpiresIn 临时会话有效期，以秒为单位。该参数仅在is_temp为true时有效，默认为添加好友后24小时
	ChatExpiresIn int `json:"chat_expires_in"  `
	// UnionID 可进行临时会话的客户unionid，该参数仅在is_temp为true时有效，如不指定则不进行限制
	UnionID string `json:"unionid"  `
	// Conclusions 结束语，会话结束时自动发送给客户，可参考“结束语定义”，仅在is_temp为true时有效
	Conclusions *Conclusions `json:"conclusions"  `
}

// AddContactWayResponse is response of Client.AddContactWay
type AddContactWayResponse struct {
	// ConfigID 新增联系方式的配置id
	ConfigID string `json:"config_id"  `
	// QrCode 联系我二维码链接，仅在scene为2时返回
	QrCode string `json:"qr_code"  `
}

// GetContactWayRequest is request of Client.GetContactWay
type GetContactWayRequest struct {
	// ConfigID 联系方式的配置id
	ConfigID string `json:"config_id"  validate:"required"`
}

// GetContactWayResponse is response of Client.GetContactWay
type GetContactWayResponse struct {
	// ConfigID 新增联系方式的配置id
	ConfigID string `json:"config_id"  `
	// Type 联系方式类型，1-单人，2-多人
	Type string `json:"type"  `
	// Scene 场景，1-在小程序中联系，2-通过二维码联系
	Scene string `json:"scene"  `
	// IsTemp 是否临时会话模式，默认为false，true表示使用临时会话模式
	IsTemp string `json:"is_temp"  `
	// Remark 联系方式的备注信息，用于助记
	Remark string `json:"remark"  `
	// SkipVerify 外部客户添加时是否无需验证
	SkipVerify string `json:"skip_verify"  `
	// State 企业自定义的state参数，用于区分不同的添加渠道，在调用“获取外部联系人详情”时会返回该参数值
	State string `json:"state"  `
	// Style 小程序中联系按钮的样式，仅在scene为1时返回，详见附录
	Style string `json:"style"  `
	// QrCode 联系二维码的URL，仅在scene为2时返回
	QrCode string `json:"qr_code"  `
	// User 使用该联系方式的用户userID列表
	User string `json:"user"  `
	// Party 使用该联系方式的部门id列表
	Party string `json:"party"  `
	// ExpiresIn 临时会话二维码有效期，以秒为单位
	ExpiresIn string `json:"expires_in"  `
	// ChatExpiresIn 临时会话有效期，以秒为单位
	ChatExpiresIn string `json:"chat_expires_in"  `
	// UnionID 可进行临时会话的客户unionid
	UnionID string `json:"unionid"  `
	// Conclusions 结束语，可参考“结束语定义”
	Conclusions string `json:"conclusions"  `
}

// UpdateContactWayRequest is request of Client.UpdateContactWay
type UpdateContactWayRequest struct {
	// ConfigID 企业联系方式的配置id
	ConfigID string `json:"config_id"  validate:"required"`
	// Remark 联系方式的备注信息，不超过30个字符，将覆盖之前的备注
	Remark string `json:"remark"  `
	// SkipVerify 外部客户添加时是否无需验证
	SkipVerify bool `json:"skip_verify"  `
	// Style 样式，只针对“在小程序中联系”的配置生效
	Style int `json:"style"  `
	// State 企业自定义的state参数，用于区分不同的添加渠道，在调用“获取外部联系人详情”时会返回该参数值
	State string `json:"state"  `
	// User 使用该联系方式的用户列表，将覆盖原有用户列表
	User []string `json:"user"  `
	// Party 使用该联系方式的部门列表，将覆盖原有部门列表，只在配置的type为2时有效
	Party []int `json:"party"  `
	// ExpiresIn 临时会话二维码有效期，以秒为单位，该参数仅在临时会话模式下有效
	ExpiresIn int `json:"expires_in"  `
	// ChatExpiresIn 临时会话有效期，以秒为单位，该参数仅在临时会话模式下有效
	ChatExpiresIn int `json:"chat_expires_in"  `
	// UnionID 可进行临时会话的客户unionid，该参数仅在临时会话模式有效，如不指定则不进行限制
	UnionID string `json:"unionid"  `
	// Conclusions 结束语，会话结束时自动发送给客户，可参考“结束语定义”，仅临时会话模式（is_temp为true）可设置
	Conclusions *Conclusions `json:"conclusions"  `
}

// DeleteContactWayRequest is request of Client.DeleteContactWay
type DeleteContactWayRequest struct {
	// ConfigID 企业联系方式的配置id
	ConfigID string `json:"config_id"  validate:"required"`
}

// CloseTempChatRequest is request of Client.CloseTempChat
type CloseTempChatRequest struct {
	// UserID 企业成员的userid
	UserID string `json:"userid"  validate:"required"`
	// ExternalUserID 客户的外部联系人userid
	ExternalUserID string `json:"external_userid"  validate:"required"`
}

// ListExternalContactRequest is request of Client.ListExternalContact
type ListExternalContactRequest struct {
	// UserID 企业成员的userid
	UserID string `json:"userid"  validate:"required"`
}

// ListExternalContactResponse is response of Client.ListExternalContact
type ListExternalContactResponse struct {
	// ExternalUserID 外部联系人的userid列表
	ExternalUserID []string `json:"external_userid"  `
}

// GetExternalContactRequest is request of Client.GetExternalContact
type GetExternalContactRequest struct {
	// ExternalUserID 外部联系人的userid，注意不是企业成员的帐号
	ExternalUserID string `json:"external_userid"  validate:"required"`
	// Cursor 上次请求返回的next_cursor
	Cursor string `json:"cursor"  `
}

// GetExternalContactResponse is response of Client.GetExternalContact
type GetExternalContactResponse struct {
	// NextCursor 分页的cursor，当跟进人多于500人时返回
	NextCursor string `json:"next_cursor"  `

	FollowUser []GetExternalContactResponseFollowUser `json:"follow_user"  `

	ExternalContact ExternalContactResponse `json:"external_contact"  `
}

// GetExternalContactResponseFollowUser is model of GetExternalContactResponse.FollowUser
type GetExternalContactResponseFollowUser struct {
	// UserID 添加了此外部联系人的企业成员userid
	UserID string `json:"userid"  `
	// Remark 该成员对此外部联系人的备注
	Remark string `json:"remark"  `
	// Description 该成员对此外部联系人的描述
	Description string `json:"description"  `
	// CreateTime 该成员添加此外部联系人的时间
	CreateTime int `json:"createtime"  `
	// RemarkCorpName 该成员对此客户备注的企业名称
	RemarkCorpName string `json:"remark_corp_name"  `
	// RemarkMobiles 该成员对此客户备注的手机号码，代开发自建应用需要管理员授权才可以获取，第三方不可获取
	RemarkMobiles []string `json:"remark_mobiles"  `
	// AddWay 该成员添加此客户的来源，具体含义详见来源定义
	AddWay int `json:"add_way"  `
	// OperUserID 发起添加的userid，如果成员主动添加，为成员的userid；如果是客户主动添加，则为客户的外部联系人userid；如果是内部成员共享/管理员分配，则为对应的成员/管理员userid
	OperUserID string `json:"oper_userid"  `
	// State 企业自定义的state参数，用于区分客户具体是通过哪个「联系我」添加，由企业通过创建「联系我」方式指定
	State string `json:"state"  `

	Tags []GetExternalContactResponseFollowUserTags `json:"tags"  `
}

// GetExternalContactResponseFollowUserTags is model of GetExternalContactResponseFollowUser.Tags
type GetExternalContactResponseFollowUserTags struct {
	// GroupName 该成员添加此外部联系人所打标签的分组名称（标签功能需要企业微信升级到2.7.5及以上版本）
	GroupName string `json:"group_name"  `
	// TagName 该成员添加此外部联系人所打标签名称
	TagName string `json:"tag_name"  `
	// Type 该成员添加此外部联系人所打标签类型, 1-企业设置，2-用户自定义，3-规则组标签（仅系统应用返回）
	Type int `json:"type"  `
	// TagID 该成员添加此外部联系人所打企业标签的id，用户自定义类型标签（type&#x3D;2）不返回
	TagID string `json:"tag_id"  `
}

// BatchGetByUserRequest is request of Client.BatchGetByUser
type BatchGetByUserRequest struct {
	// UserIDList 企业成员的userid列表，字符串类型，最多支持100个
	UserIDList []string `json:"userid_list"  validate:"required"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"  `
	// Limit 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取最大值
	Limit int `json:"limit"  `
}

// BatchGetByUserResponse is response of Client.BatchGetByUser
type BatchGetByUserResponse struct {
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"  `

	ExternalContactList []BatchGetByUserResponseExternalContactList `json:"external_contact_list"  `
}

// BatchGetByUserResponseExternalContactList is model of BatchGetByUserResponse.ExternalContactList
type BatchGetByUserResponseExternalContactList struct {
	// ExternalContact 客户的基本信息，可以参考获取客户详情
	ExternalContact ExternalContactResponse `json:"external_contact"  `
	// FollowInfo 企业成员客户跟进信息，可以参考获取客户详情，但标签信息只会返回企业标签和规则组标签的tag_id，个人标签将不再返回
	FollowInfo GetExternalContactResponseFollowUser `json:"follow_info"  `
}

// RemarkExternalContactRequest is request of Client.RemarkExternalContact
type RemarkExternalContactRequest struct {
	// UserID 企业成员的userid
	UserID string `json:"userid"  validate:"required"`
	// ExternalUserID 外部联系人userid
	ExternalUserID string `json:"external_userid"  validate:"required"`
	// Remark 此用户对外部联系人的备注，最多20个字符
	Remark string `json:"remark"  `
	// Description 此用户对外部联系人的描述，最多150个字符
	Description string `json:"description"  `
	// RemarkCompany 此用户对外部联系人备注的所属公司名称，最多20个字符
	RemarkCompany string `json:"remark_company"  `
	// RemarkMobiles 此用户对外部联系人备注的手机号
	RemarkMobiles []string `json:"remark_mobiles"  `
	// RemarkPicMediaID 备注图片的mediaid，
	RemarkPicMediaID string `json:"remark_pic_mediaid"  `
}

// UnionIDToExternalUserIDRequest is request of Client.UnionIDToExternalUserID
type UnionIDToExternalUserIDRequest struct {
	// UnionID 微信客户的unionid
	UnionID string `json:"unionid"  validate:"required"`
}

// UnionIDToExternalUserIDResponse is response of Client.UnionIDToExternalUserID
type UnionIDToExternalUserIDResponse struct {
	// ExternalUserID 该企业的外部联系人ID
	ExternalUserID string `json:"external_userid"  `
}

// GetCorpTagListRequest is request of Client.GetCorpTagList
type GetCorpTagListRequest struct {
	// TagID 要查询的标签id
	TagID []string `json:"tag_id"  `
	// GroupID 要查询的标签组id，返回该标签组以及其下的所有标签信息
	GroupID []string `json:"group_id"  `
}

// GetCorpTagListResponse is response of Client.GetCorpTagList
type GetCorpTagListResponse struct {
	// TagGroup 标签组列表
	TagGroup []GetCorpTagListResponseTagGroup `json:"tag_group"  `
}

// GetCorpTagListResponseTagGroup is model of GetCorpTagListResponse.TagGroup
type GetCorpTagListResponseTagGroup struct {
	// GroupID 标签组id
	GroupID string `json:"group_id"  `
	// GroupName 标签组名称
	GroupName string `json:"group_name"  `
	// CreateTime 标签组创建时间
	CreateTime int `json:"create_time"  `
	// Order 标签组排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
	Order int `json:"order"  `
	// Deleted 标签组是否已经被删除，只在指定tag_id进行查询时返回
	Deleted bool `json:"deleted"  `
	// Tag 标签组内的标签列表
	Tag []GetCorpTagListResponseTagGroupTag `json:"tag"  `
}

// GetCorpTagListResponseTagGroupTag is model of GetCorpTagListResponseTagGroup.Tag
type GetCorpTagListResponseTagGroupTag struct {
	// ID 标签id
	ID string `json:"id"  `
	// Name 标签名称
	Name string `json:"name"  `
	// CreateTime 标签创建时间
	CreateTime int `json:"create_time"  `
	// Order 标签排序的次序值，order值大的排序靠前。有效的值范围是[0, 2^32)
	Order int `json:"order"  `
	// Deleted 标签是否已经被删除，只在指定tag_id/group_id进行查询时返回
	Deleted bool `json:"deleted"  `
}

// AddCorpTagRequest is request of Client.AddCorpTag
type AddCorpTagRequest struct {
	// GroupID 标签组id
	GroupID string `json:"group_id"  `
	// GroupName 标签组名称，最长为30个字符
	GroupName string `json:"group_name"  `
	// Order 标签组次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
	Order int `json:"order"  `
	// AgentID 授权方安装的应用agentid。仅旧的第三方多应用套件需要填此参数
	AgentID int `json:"agentid"  `

	Tag []AddCorpTagRequestTag `json:"tag"  `
}

// AddCorpTagResponse is response of Client.AddCorpTag
type AddCorpTagResponse struct {
	TagGroup AddCorpTagResponseTagGroup `json:"tag_group"  `
}

// AddCorpTagRequestTag is model of AddCorpTagRequest.Tag
type AddCorpTagRequestTag struct {
	// Name 添加的标签名称，最长为30个字符
	Name string `json:"name"  validate:"required"`
	// Order 标签次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
	Order int `json:"order"  `
}

// AddCorpTagResponseTagGroup is model of AddCorpTagResponse.TagGroup
type AddCorpTagResponseTagGroup struct {
	// GroupID 标签组id
	GroupID string `json:"group_id"  `
	// GroupName 标签组名称
	GroupName string `json:"group_name"  `
	// CreateTime 标签组创建时间
	CreateTime int `json:"create_time"  `
	// Order 标签组次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
	Order int `json:"order"  `
	// Tag 标签组内的标签列表
	Tag []AddCorpTagResponseTagGroupTag `json:"tag"  `
}

// AddCorpTagResponseTagGroupTag is model of AddCorpTagResponseTagGroup.Tag
type AddCorpTagResponseTagGroupTag struct {
	// ID 新建标签id
	ID string `json:"id"  `
	// Name 新建标签名称
	Name string `json:"name"  `
	// CreateTime 标签创建时间
	CreateTime int `json:"create_time"  `
	// Order 标签次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
	Order int `json:"order"  `
}

// EditCorpTagRequest is request of Client.EditCorpTag
type EditCorpTagRequest struct {
	// ID 标签或标签组的id
	ID string `json:"id"  validate:"required"`
	// Name 新的标签或标签组名称，最长为30个字符
	Name string `json:"name"  `
	// Order 标签/标签组的次序值。order值大的排序靠前。有效的值范围是[0, 2^32)
	Order int `json:"order"  `
	// AgentID 授权方安装的应用agentid。仅旧的第三方多应用套件需要填此参数
	AgentID int `json:"agentid"  `
}

// DeleteCorpTagRequest is request of Client.DeleteCorpTag
type DeleteCorpTagRequest struct {
	// TagID 标签的id列表
	TagID []string `json:"tag_id"  `
	// GroupID 标签组的id列表
	GroupID []string `json:"group_id"  `
	// AgentID 授权方安装的应用agentid。仅旧的第三方多应用套件需要填此参数
	AgentID int `json:"agentid"  `
}

// MarkTagExternalContactRequest is request of Client.MarkTagExternalContact
type MarkTagExternalContactRequest struct {
	// UserID 添加外部联系人的userid
	UserID string `json:"userid"  validate:"required"`
	// ExternalUserID 外部联系人userid
	ExternalUserID string `json:"external_userid"  validate:"required"`
	// AddTag 要标记的标签列表
	AddTag []string `json:"add_tag"  `
	// RemoveTag 要移除的标签列表
	RemoveTag []string `json:"remove_tag"  `
}

// TransferCustomerExternalContactRequest is request of Client.TransferCustomerExternalContact
type TransferCustomerExternalContactRequest struct {
	// HandoverUserID 原跟进成员的userid
	HandoverUserID string `json:"handover_userid"  validate:"required"`
	// TakeoverUserID 接替成员的userid
	TakeoverUserID string `json:"takeover_userid"  validate:"required"`
	// ExternalUserID 客户的external_userid列表，每次最多分配100个客户
	ExternalUserID []string `json:"external_userid"  validate:"required"`
	// TransferSuccessMessage 转移成功后发给客户的消息，最多200个字符，不填则使用默认文案
	TransferSuccessMessage string `json:"transfer_success_msg"  `
}

// TransferCustomerExternalContactResponse is response of Client.TransferCustomerExternalContact
type TransferCustomerExternalContactResponse struct {
	Customer []TransferCustomerExternalContactResponseCustomer `json:"customer"  `
}

// TransferCustomerExternalContactResponseCustomer is model of TransferCustomerExternalContactResponse.Customer
type TransferCustomerExternalContactResponseCustomer struct {
	// ExternalUserID 客户的external_userid
	ExternalUserID string `json:"external_userid"  `
	// ErrorCode 对此客户进行分配的结果, 具体可参考全局错误码, 0表示成功发起接替,待24小时后自动接替,并不代表最终接替成功
	ErrorCode int `json:"errcode"  `
}

// TransferResultExternalContactRequest is request of Client.TransferResultExternalContact
type TransferResultExternalContactRequest struct {
	// HandoverUserID 原添加成员的userid
	HandoverUserID string `json:"handover_userid"  validate:"required"`
	// TakeoverUserID 接替成员的userid
	TakeoverUserID string `json:"takeover_userid"  validate:"required"`
	// Cursor 分页查询的cursor，每个分页返回的数据不会超过1000条；不填或为空表示获取第一个分页；
	Cursor string `json:"cursor"  `
}

// TransferResultExternalContactResponse is response of Client.TransferResultExternalContact
type TransferResultExternalContactResponse struct {
	// NextCursor 下个分页的起始cursor
	NextCursor string `json:"next_cursor"  `

	Customer []TransferResultExternalContactResponseCustomer `json:"customer"  `
}

// TransferResultExternalContactResponseCustomer is model of TransferResultExternalContactResponse.Customer
type TransferResultExternalContactResponseCustomer struct {
	// ExternalUserID 转接客户的外部联系人userid
	ExternalUserID string `json:"external_userid"  `
	// Status 接替状态， 1-接替完毕 2-等待接替 3-客户拒绝 4-接替成员客户达到上限 5-无接替记录
	Status int `json:"status"  `
	// TakeoverTime 接替客户的时间，如果是等待接替状态，则为未来的自动接替时间
	TakeoverTime int `json:"takeover_time"  `
}

// GetUnassignedListRequest is request of Client.GetUnassignedList
type GetUnassignedListRequest struct {
	// PageID 分页查询，要查询页号，从0开始
	PageID string `json:"page_id"  `
	// PageSize 每次返回的最大记录数，默认为1000，最大值为1000
	PageSize string `json:"page_size"  `
	// Cursor 分页查询游标，字符串类型，适用于数据量较大的情况，如果使用该参数则无需填写page_id，该参数由上一次调用返回
	Cursor string `json:"cursor"  `
}

// GetUnassignedListResponse is response of Client.GetUnassignedList
type GetUnassignedListResponse struct {
	// IsLast 是否是最后一条记录
	IsLast bool `json:"is_last"  `
	// NextCursor 分页查询游标,已经查完则返回空(“”)
	NextCursor string `json:"next_cursor"  `

	Info []GetUnassignedListResponseInfo `json:"info"  `
}

// GetUnassignedListResponseInfo is model of GetUnassignedListResponse.Info
type GetUnassignedListResponseInfo struct {
	// HandoverUserID 离职成员的userid
	HandoverUserID string `json:"handover_userid"  `
	// ExternalUserID 外部联系人userid
	ExternalUserID string `json:"external_userid"  `
	// DismissionTime 成员离职时间
	DismissionTime int `json:"dimission_time"  `
}

// TransferCustomerResignedRequest is request of Client.TransferCustomerResigned
type TransferCustomerResignedRequest struct {
	// HandoverUserID 原跟进成员的userid
	HandoverUserID string `json:"handover_userid"  validate:"required"`
	// TakeoverUserID 接替成员的userid
	TakeoverUserID string `json:"takeover_userid"  validate:"required"`
	// ExternalUserID 客户的external_userid列表，最多一次转移100个客户
	ExternalUserID []string `json:"external_userid"  validate:"required"`
}

// TransferCustomerResignedResponse is response of Client.TransferCustomerResigned
type TransferCustomerResignedResponse struct {
	Customer []TransferCustomerResignedResponseCustomer `json:"customer"  `
}

// TransferCustomerResignedResponseCustomer is model of TransferCustomerResignedResponse.Customer
type TransferCustomerResignedResponseCustomer struct {
	// ExternalUserID 客户的external_userid
	ExternalUserID string `json:"external_userid"  `
	// ErrorCode 对此客户进行分配的结果, 具体可参考全局错误码, 0表示开始分配流程,待24小时后自动接替,并不代表最终分配成功
	ErrorCode int `json:"errcode"  `
}

// TransferResultResignedRequest is request of Client.TransferResultResigned
type TransferResultResignedRequest struct {
	// HandoverUserID 原添加成员的userid
	HandoverUserID string `json:"handover_userid"  validate:"required"`
	// TakeoverUserID 接替成员的userid
	TakeoverUserID string `json:"takeover_userid"  validate:"required"`
	// Cursor 分页查询的cursor，每个分页返回的数据不会超过1000条；不填或为空表示获取第一个分页
	Cursor string `json:"cursor"  `
}

// TransferResultResignedResponse is response of Client.TransferResultResigned
type TransferResultResignedResponse struct {
	// NextCursor 下个分页的起始cursor
	NextCursor string `json:"next_cursor"  `

	Customer []TransferResultResignedResponseCustomer `json:"customer"  `
}

// TransferResultResignedResponseCustomer is model of TransferResultResignedResponse.Customer
type TransferResultResignedResponseCustomer struct {
	// ExternalUserID 转接客户的外部联系人userid
	ExternalUserID string `json:"external_userid"  `
	// Status 接替状态， 1-接替完毕 2-等待接替 3-客户拒绝 4-接替成员客户达到上限
	Status int `json:"status"  `
	// TakeoverTime 接替客户的时间，如果是等待接替状态，则为未来的自动接替时间
	TakeoverTime int `json:"takeover_time"  `
}

// TransferGroupChatRequest is request of Client.TransferGroupChat
type TransferGroupChatRequest struct {
	// ChatIDList 需要转群主的客户群ID列表。取值范围： 1 ~ 100
	ChatIDList []string `json:"chat_id_list"  validate:"required"`
	// NewOwner 新群主ID
	NewOwner string `json:"new_owner"  validate:"required"`
}

// TransferGroupChatResponse is response of Client.TransferGroupChat
type TransferGroupChatResponse struct {
	// FailedChatList 没能成功继承的群
	FailedChatList []TransferGroupChatResponseFailedChatList `json:"failed_chat_list"  `
}

// TransferGroupChatResponseFailedChatList is model of TransferGroupChatResponse.FailedChatList
type TransferGroupChatResponseFailedChatList struct {
	// ChatID 没能成功继承的群ID
	ChatID string `json:"chat_id"  `
	// ErrorCode 没能成功继承的群，错误码
	ErrorCode int `json:"errcode"  `
	// ErrorMessage 没能成功继承的群，错误描述
	ErrorMessage string `json:"errmsg"  `
}

// ListGroupChatRequest is request of Client.ListGroupChat
type ListGroupChatRequest struct {
	// StatusFilter 客户群跟进状态过滤。0 - 所有列表(即不过滤)1 - 离职待继承2 - 离职继承中3 - 离职继承完成默认为0
	StatusFilter int `json:"status_filter"  `
	// OwnerFilter 群主过滤。如果不填，表示获取应用可见范围内全部群主的数据（但是不建议这么用，如果可见范围人数超过1000人，为了防止数据包过大，会报错 81017）
	OwnerFilter ListGroupChatRequestOwnerFilter `json:"owner_filter"  `
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用不填
	Cursor string `json:"cursor"  `
	// Limit 分页，预期请求的数据量，取值范围 1 ~ 1000
	Limit int `json:"limit"  validate:"required"`
}

// ListGroupChatResponse is response of Client.ListGroupChat
type ListGroupChatResponse struct {
	// GroupChatList 客户群列表
	GroupChatList []ListGroupChatResponseGroupChatList `json:"group_chat_list"  `
	// NextCursor 分页游标，下次请求时填写以获取之后分页的记录。如果该字段返回空则表示已没有更多数据
	NextCursor string `json:"next_cursor"  `
}

// ListGroupChatRequestOwnerFilter is model of ListGroupChatRequest.OwnerFilter
type ListGroupChatRequestOwnerFilter struct {
	// UserIDList 用户ID列表。最多100个
	UserIDList []string `json:"userid_list"  `
}

// ListGroupChatResponseGroupChatList is model of ListGroupChatResponse.GroupChatList
type ListGroupChatResponseGroupChatList struct {
	// ChatID 客户群ID
	ChatID string `json:"chat_id"  `
	// Status 客户群跟进状态。0 - 跟进人正常1 - 跟进人离职2 - 离职继承中3 - 离职继承完成
	Status int `json:"status"  `
}

// GetGroupChatRequest is request of Client.GetGroupChat
type GetGroupChatRequest struct {
	// ChatID 客户群ID
	ChatID string `json:"chat_id"  validate:"required"`
	// NeedName 是否需要返回群成员的名字group_chat.member_list.name。0-不返回；1-返回。默认不返回
	NeedName int `json:"need_name"  `
}

// GetGroupChatResponse is response of Client.GetGroupChat
type GetGroupChatResponse struct {
	// GroupChat 客户群详情
	GroupChat GetGroupChatResponseGroupChat `json:"group_chat"  `
}

// GetGroupChatResponseGroupChat is model of GetGroupChatResponse.GroupChat
type GetGroupChatResponseGroupChat struct {
	// ChatID 客户群ID
	ChatID string `json:"chat_id"  `
	// Name 群名
	Name string `json:"name"  `
	// Owner 群主ID
	Owner string `json:"owner"  `
	// CreateTime 群的创建时间
	CreateTime int `json:"create_time"  `
	// Notice 群公告
	Notice string `json:"notice"  `
	// MemberList 群成员列表
	MemberList []GetGroupChatResponseGroupChatMemberList `json:"member_list"  `
	// AdminList 群管理员列表
	AdminList []GetGroupChatResponseGroupChatAdminList `json:"admin_list"  `
}

// GetGroupChatResponseGroupChatMemberList is model of GetGroupChatResponseGroupChat.MemberList
type GetGroupChatResponseGroupChatMemberList struct {
	// UserID 群成员id
	UserID string `json:"userid"  `
	// Type 成员类型。1 - 企业成员2 - 外部联系人
	Type int `json:"type"  `
	// UnionID 外部联系人在微信开放平台的唯一身份标识（微信unionid），通过此字段企业可将外部联系人与公众号/小程序用户关联起来。仅当群成员类型是微信用户（包括企业成员未添加好友），且企业或第三方服务商绑定了微信开发者ID有此字段。查看绑定方法
	UnionID string `json:"unionid"  `
	// JoinTime 入群时间
	JoinTime int `json:"join_time"  `
	// JoinScene 入群方式。1 - 由群成员邀请入群（直接邀请入群）2 - 由群成员邀请入群（通过邀请链接入群）3 - 通过扫描群二维码入群
	JoinScene int `json:"join_scene"  `
	// Invitor 邀请者。目前仅当是由本企业内部成员邀请入群时会返回该值
	Invitor GetGroupChatResponseGroupChatMemberListInvitor `json:"invitor"  `
	// GroupNickname 在群里的昵称
	GroupNickname string `json:"group_nickname"  `
	// Name 名字。仅当 need_name &#x3D; 1 时返回如果是微信用户，则返回其在微信中设置的名字如果是企业微信联系人，则返回其设置对外展示的别名或实名
	Name string `json:"name"  `
}

// GetGroupChatResponseGroupChatMemberListInvitor is model of GetGroupChatResponseGroupChatMemberList.Invitor
type GetGroupChatResponseGroupChatMemberListInvitor struct {
	// UserID 邀请者的userid
	UserID string `json:"userid"  `
}

// GetGroupChatResponseGroupChatAdminList is model of GetGroupChatResponseGroupChat.AdminList
type GetGroupChatResponseGroupChatAdminList struct {
	// UserID 群管理员userid
	UserID string `json:"userid"  `
}

// GetMomentListRequest is request of Client.GetMomentList
type GetMomentListRequest struct {
	// StartTime 朋友圈记录开始时间。Unix时间戳
	StartTime int `json:"start_time"  validate:"required"`
	// EndTime 朋友圈记录结束时间。Unix时间戳
	EndTime int `json:"end_time"  validate:"required"`
	// Creator 朋友圈创建人的userid
	Creator string `json:"creator"  `
	// FilterType 朋友圈类型。0：企业发表 1：个人发表 2：所有，包括个人创建以及企业创建，默认情况下为所有类型
	FilterType int `json:"filter_type"  `
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"  `
	// Limit 返回的最大记录数，整型，最大值100，默认值100，超过最大值时取默认值
	Limit int `json:"limit"  `
}

// GetMomentListResponse is response of Client.GetMomentList
type GetMomentListResponse struct {
	// NextCursor 分页游标，下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"  `
	// MomentList 朋友圈列表
	MomentList []GetMomentListResponseMomentList `json:"moment_list"  `
}

// GetMomentListResponseMomentList is model of GetMomentListResponse.MomentList
type GetMomentListResponseMomentList struct {
	// MomentID 朋友圈id
	MomentID string `json:"moment_id"  `
	// Creator 朋友圈创建者userid
	Creator string `json:"creator"  `
	// CreateTime 创建时间
	CreateTime string `json:"create_time"  `
	// CreateType 朋友圈创建来源。0：企业 1：个人
	CreateType int `json:"create_type"  `
	// VisibleType 可见范围类型。0：部分可见 1：公开
	VisibleType int `json:"visible_type"  `

	Text GetMomentListResponseMomentListText `json:"text"  `

	Image []GetMomentListResponseMomentListImage `json:"image"  `

	Video GetMomentListResponseMomentListVideo `json:"video"  `

	Link GetMomentListResponseMomentListLink `json:"link"  `

	Location GetMomentListResponseMomentListLocation `json:"location"  `
}

// GetMomentListResponseMomentListText is model of GetMomentListResponseMomentList.Text
type GetMomentListResponseMomentListText struct {
	// Content 文本消息结构
	Content string `json:"content"  `
}

// GetMomentListResponseMomentListImage is model of GetMomentListResponseMomentList.Image
type GetMomentListResponseMomentListImage struct {
	// MediaID 图片的media_id列表，可以通过获取临时素材下载资源
	MediaID string `json:"media_id"  `
}

// GetMomentListResponseMomentListVideo is model of GetMomentListResponseMomentList.Video
type GetMomentListResponseMomentListVideo struct {
	// MediaID 视频media_id，可以通过获取临时素材下载资源
	MediaID string `json:"media_id"  `
	// ThumbMediaID 视频封面media_id，可以通过获取临时素材下载资源
	ThumbMediaID string `json:"thumb_media_id"  `
}

// GetMomentListResponseMomentListLink is model of GetMomentListResponseMomentList.Link
type GetMomentListResponseMomentListLink struct {
	// Title 网页链接标题
	Title string `json:"title"  `
	// URL 网页链接url
	URL string `json:"url"  `
}

// GetMomentListResponseMomentListLocation is model of GetMomentListResponseMomentList.Location
type GetMomentListResponseMomentListLocation struct {
	// Latitude 地理位置纬度
	Latitude string `json:"latitude"  `
	// Longitude 地理位置经度
	Longitude string `json:"longitude"  `
	// Name 地理位置名称
	Name string `json:"name"  `
}

// GetMomentTaskRequest is request of Client.GetMomentTask
type GetMomentTaskRequest struct {
	// MomentID 朋友圈id,仅支持企业发表的朋友圈id
	MomentID string `json:"moment_id"  validate:"required"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"  `
	// Limit 返回的最大记录数，整型，最大值1000，默认值500，超过最大值时取默认值
	Limit int `json:"limit"  `
}

// GetMomentTaskResponse is response of Client.GetMomentTask
type GetMomentTaskResponse struct {
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"  `
	// TaskList 发表任务列表
	TaskList []GetMomentTaskResponseTaskList `json:"task_list"  `
}

// GetMomentTaskResponseTaskList is model of GetMomentTaskResponse.TaskList
type GetMomentTaskResponseTaskList struct {
	// UserID 发表成员用户userid
	UserID string `json:"userid"  `
	// PublishStatus 成员发表状态。0:未发表 1：已发表
	PublishStatus int `json:"publish_status"  `
}

// GetMomentCustomerListRequest is request of Client.GetMomentCustomerList
type GetMomentCustomerListRequest struct {
	// MomentID 朋友圈id
	MomentID string `json:"moment_id"  validate:"required"`
	// UserID 企业发表成员userid，如果是企业创建的朋友圈，可以通过获取客户朋友圈企业发表的列表获取已发表成员userid，如果是个人创建的朋友圈，创建人userid就是企业发表成员userid
	UserID string `json:"userid"  validate:"required"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"  `
	// Limit 返回的最大记录数，整型，最大值1000，默认值500，超过最大值时取默认值
	Limit int `json:"limit"  `
}

// GetMomentCustomerListResponse is response of Client.GetMomentCustomerList
type GetMomentCustomerListResponse struct {
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"  `
	// CustomerList 成员可见客户列表
	CustomerList []GetMomentCustomerListResponseCustomerList `json:"customer_list"  `
}

// GetMomentCustomerListResponseCustomerList is model of GetMomentCustomerListResponse.CustomerList
type GetMomentCustomerListResponseCustomerList struct {
	// UserID 发表成员用户userid
	UserID string `json:"userid"  `
	// ExternalUserID 发送成功的外部联系人userid
	ExternalUserID string `json:"external_userid"  `
}

// GetMomentSendResultRequest is request of Client.GetMomentSendResult
type GetMomentSendResultRequest struct {
	// MomentID 朋友圈id
	MomentID string `json:"moment_id"  validate:"required"`
	// UserID 企业发表成员userid，如果是企业创建的朋友圈，可以通过获取客户朋友圈企业发表的列表获取已发表成员userid，如果是个人创建的朋友圈，创建人userid就是企业发表成员userid
	UserID string `json:"userid"  validate:"required"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"  `
	// Limit 返回的最大记录数，整型，最大值5000，默认值3000，超过最大值时取默认值
	Limit int `json:"limit"  `
}

// GetMomentSendResultResponse is response of Client.GetMomentSendResult
type GetMomentSendResultResponse struct {
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"  `
	// CustomerList 成员发送成功客户列表
	CustomerList []GetMomentSendResultResponseCustomerList `json:"customer_list"  `
}

// GetMomentSendResultResponseCustomerList is model of GetMomentSendResultResponse.CustomerList
type GetMomentSendResultResponseCustomerList struct {
	// ExternalUserID 成员发送成功的外部联系人userid
	ExternalUserID string `json:"external_userid"  `
}

// GetMomentCommentsRequest is request of Client.GetMomentComments
type GetMomentCommentsRequest struct {
	// MomentID 朋友圈id
	MomentID string `json:"moment_id"  validate:"required"`
	// UserID 企业发表成员userid，如果是企业创建的朋友圈，可以通过获取客户朋友圈企业发表的列表获取已发表成员userid，如果是个人创建的朋友圈，创建人userid就是企业发表成员userid
	UserID string `json:"userid"  validate:"required"`
}

// GetMomentCommentsResponse is response of Client.GetMomentComments
type GetMomentCommentsResponse struct {
	// CommentList 评论列表
	CommentList []GetMomentCommentsResponseCommentList `json:"comment_list"  `
	// LikeList 点赞列表
	LikeList []GetMomentCommentsResponseLikeList `json:"like_list"  `
}

// GetMomentCommentsResponseCommentList is model of GetMomentCommentsResponse.CommentList
type GetMomentCommentsResponseCommentList struct {
	// ExternalUserID 评论的外部联系人userid
	ExternalUserID string `json:"external_userid"  `
	// UserID 评论的企业成员userid，userid与external_userid不会同时出现
	UserID string `json:"userid"  `
	// CreateTime 评论时间
	CreateTime int `json:"create_time"  `
}

// GetMomentCommentsResponseLikeList is model of GetMomentCommentsResponse.LikeList
type GetMomentCommentsResponseLikeList struct {
	// ExternalUserID 点赞的外部联系人userid
	ExternalUserID string `json:"external_userid"  `
	// UserID 点赞的企业成员userid，userid与external_userid不会同时出现
	UserID string `json:"userid"  `
	// CreateTime 点赞时间
	CreateTime int `json:"create_time"  `
}

// AddMessageTemplateRequest is request of Client.AddMessageTemplate
type AddMessageTemplateRequest struct {
	// ChatType 群发任务的类型，默认为single，表示发送给客户，group表示发送给客户群
	ChatType string `json:"chat_type"  `
	// ExternalUserID 客户的外部联系人id列表，仅在chat_type为single时有效，不可与sender同时为空，最多可传入1万个客户
	ExternalUserID []string `json:"external_userid"  `
	// Sender 发送企业群发消息的成员userid，当类型为发送给客户群时必填
	Sender string `json:"sender"  `
	// Attachments 附件，最多支持添加9个附件
	Attachments []AddMessageTemplateRequestAttachments `json:"attachments"  `

	Text AddMessageTemplateRequestText `json:"text"  `

	Image AddMessageTemplateRequestImage `json:"image"  `

	Link AddMessageTemplateRequestLink `json:"link"  `

	MiniProgram AddMessageTemplateRequestMiniProgram `json:"miniprogram"  `

	Video AddMessageTemplateRequestVideo `json:"video"  `

	File AddMessageTemplateRequestFile `json:"file"  `
}

// AddMessageTemplateResponse is response of Client.AddMessageTemplate
type AddMessageTemplateResponse struct {
	// FailList 无效或无法发送的external_userid列表
	FailList []string `json:"fail_list"  `
	// MessageID 企业群发消息的id，可用于获取群发消息发送结果
	MessageID string `json:"msgid"  `
}

// AddMessageTemplateRequestText is model of AddMessageTemplateRequest.Text
type AddMessageTemplateRequestText struct {
	// Content 消息文本内容，最多4000个字节
	Content string `json:"content"  `
}

// AddMessageTemplateRequestAttachments is model of AddMessageTemplateRequest.Attachments
type AddMessageTemplateRequestAttachments struct {
	// MessageType 附件类型，可选image、link、miniprogram或者video
	MessageType string `json:"msgtype"  validate:"required"`
}

// AddMessageTemplateRequestImage is model of AddMessageTemplateRequest.Image
type AddMessageTemplateRequestImage struct {
	// MediaID 图片的media_id，可以通过素材管理接口获得
	MediaID string `json:"media_id"  `
	// PicURL 图片的链接，仅可使用上传图片接口得到的链接
	PicURL string `json:"pic_url"  `
}

// AddMessageTemplateRequestLink is model of AddMessageTemplateRequest.Link
type AddMessageTemplateRequestLink struct {
	// Title 图文消息标题，最长128个字节
	Title string `json:"title"  validate:"required"`
	// PicURL 图文消息封面的url，最长2048个字节
	PicURL string `json:"picurl"  `
	// Desc 图文消息的描述，最多512个字节
	Desc string `json:"desc"  `
	// URL 图文消息的链接，最长2048个字节
	URL string `json:"url"  validate:"required"`
}

// AddMessageTemplateRequestMiniProgram is model of AddMessageTemplateRequest.MiniProgram
type AddMessageTemplateRequestMiniProgram struct {
	// Title 小程序消息标题，最多64个字节
	Title string `json:"title"  validate:"required"`
	// PicMediaID 小程序消息封面的mediaid，封面图建议尺寸为520*416
	PicMediaID string `json:"pic_media_id"  validate:"required"`
	// AppID 小程序appid（可以在微信公众平台上查询），必须是关联到企业的小程序应用
	AppID string `json:"appid"  validate:"required"`
	// Page 小程序page路径
	Page string `json:"page"  validate:"required"`
}

// AddMessageTemplateRequestVideo is model of AddMessageTemplateRequest.Video
type AddMessageTemplateRequestVideo struct {
	// MediaID 视频的media_id，可以通过素材管理接口获得
	MediaID string `json:"media_id"  validate:"required"`
}

// AddMessageTemplateRequestFile is model of AddMessageTemplateRequest.File
type AddMessageTemplateRequestFile struct {
	// MediaID 文件的media_id，可以通过素材管理接口获得
	MediaID string `json:"media_id"  validate:"required"`
}

// GetGroupMessageListV2Request is request of Client.GetGroupMessageListV2
type GetGroupMessageListV2Request struct {
	// ChatType 群发任务的类型，默认为single，表示发送给客户，group表示发送给客户群
	ChatType string `json:"chat_type"  validate:"required"`
	// StartTime 群发任务记录开始时间
	StartTime int `json:"start_time"  validate:"required"`
	// EndTime 群发任务记录结束时间
	EndTime int `json:"end_time"  validate:"required"`
	// Creator 群发任务创建人企业账号id
	Creator string `json:"creator"  `
	// FilterType 创建人类型。0：企业发表 1：个人发表 2：所有，包括个人创建以及企业创建，默认情况下为所有类型
	FilterType int `json:"filter_type"  `
	// Limit 返回的最大记录数，整型，最大值100，默认值50，超过最大值时取默认值
	Limit int `json:"limit"  `
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"  `
}

// GetGroupMessageListV2Response is response of Client.GetGroupMessageListV2
type GetGroupMessageListV2Response struct {
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"  `
	// GroupMessageList 群发记录列表
	GroupMessageList []GetGroupMessageListV2ResponseGroupMessageList `json:"group_msg_list"  `
}

// GetGroupMessageListV2ResponseGroupMessageList is model of GetGroupMessageListV2Response.GroupMessageList
type GetGroupMessageListV2ResponseGroupMessageList struct {
	// MessageID 企业群发消息的id，可用于获取企业群发成员执行结果
	MessageID string `json:"msgid"  `
	// Creator 群发消息创建者userid，API接口创建的群发消息不返回该字段
	Creator string `json:"creator"  `
	// CreateTime 创建时间
	CreateTime string `json:"create_time"  `
	// CreateType 群发消息创建来源。0：企业 1：个人
	CreateType int `json:"create_type"  `

	Text GetGroupMessageListV2ResponseGroupMessageListText `json:"text"  `

	Attachments []GetGroupMessageListV2ResponseGroupMessageListAttachments `json:"attachments"  `
}

// GetGroupMessageListV2ResponseGroupMessageListText is model of GetGroupMessageListV2ResponseGroupMessageList.Text
type GetGroupMessageListV2ResponseGroupMessageListText struct {
	// Content 消息文本内容，最多4000个字节
	Content string `json:"content"  `
}

// GetGroupMessageListV2ResponseGroupMessageListAttachments is model of GetGroupMessageListV2ResponseGroupMessageList.Attachments
type GetGroupMessageListV2ResponseGroupMessageListAttachments struct {
	// MessageType 值必须是image
	MessageType string `json:"msgtype"  `

	Image GetGroupMessageListV2ResponseGroupMessageListAttachmentsImage `json:"image"  `

	Link GetGroupMessageListV2ResponseGroupMessageListAttachmentsLink `json:"link"  `

	MiniProgram GetGroupMessageListV2ResponseGroupMessageListAttachmentsMiniProgram `json:"miniprogram"  `

	Video GetGroupMessageListV2ResponseGroupMessageListAttachmentsVideo `json:"video"  `

	File GetGroupMessageListV2ResponseGroupMessageListAttachmentsFile `json:"file"  `
}

// GetGroupMessageListV2ResponseGroupMessageListAttachmentsImage is model of GetGroupMessageListV2ResponseGroupMessageListAttachments.Image
type GetGroupMessageListV2ResponseGroupMessageListAttachmentsImage struct {
	// MediaID 图片的media_id，可以通过获取临时素材下载资源
	MediaID string `json:"media_id"  `
	// PicURL 图片的url，与图片的media_id不能共存优先吐出media_id
	PicURL string `json:"pic_url"  `
}

// GetGroupMessageListV2ResponseGroupMessageListAttachmentsLink is model of GetGroupMessageListV2ResponseGroupMessageListAttachments.Link
type GetGroupMessageListV2ResponseGroupMessageListAttachmentsLink struct {
	// Title 图文消息标题
	Title string `json:"title"  `
	// PicURL 图文消息封面的url
	PicURL string `json:"picurl"  `
	// Desc 图文消息的描述，最多512个字节
	Desc string `json:"desc"  `
	// URL 图文消息的链接
	URL string `json:"url"  `
}

// GetGroupMessageListV2ResponseGroupMessageListAttachmentsMiniProgram is model of GetGroupMessageListV2ResponseGroupMessageListAttachments.MiniProgram
type GetGroupMessageListV2ResponseGroupMessageListAttachmentsMiniProgram struct {
	// Title 小程序消息标题，最多64个字节
	Title string `json:"title"  `
	// AppID 小程序appid，必须是关联到企业的小程序应用
	AppID string `json:"appid"  `
	// Page 小程序page路径
	Page string `json:"page"  `
}

// GetGroupMessageListV2ResponseGroupMessageListAttachmentsVideo is model of GetGroupMessageListV2ResponseGroupMessageListAttachments.Video
type GetGroupMessageListV2ResponseGroupMessageListAttachmentsVideo struct {
	// MediaID 视频的media_id，可以通过获取临时素材下载资源
	MediaID string `json:"media_id"  `
}

// GetGroupMessageListV2ResponseGroupMessageListAttachmentsFile is model of GetGroupMessageListV2ResponseGroupMessageListAttachments.File
type GetGroupMessageListV2ResponseGroupMessageListAttachmentsFile struct {
	// MediaID 文件的media_id，可以通过获取临时素材下载资源
	MediaID string `json:"media_id"  `
}

// GetGroupMessageTaskRequest is request of Client.GetGroupMessageTask
type GetGroupMessageTaskRequest struct {
	// MessageID 群发消息的id，通过获取群发记录列表接口返回
	MessageID string `json:"msgid"  validate:"required"`
	// Limit 返回的最大记录数，整型，最大值1000，默认值500，超过最大值时取默认值
	Limit int `json:"limit"  `
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"  `
}

// GetGroupMessageTaskResponse is response of Client.GetGroupMessageTask
type GetGroupMessageTaskResponse struct {
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"  `
	// TaskList 群发成员发送任务列表
	TaskList []GetGroupMessageTaskResponseTaskList `json:"task_list"  `
}

// GetGroupMessageTaskResponseTaskList is model of GetGroupMessageTaskResponse.TaskList
type GetGroupMessageTaskResponseTaskList struct {
	// UserID 企业服务人员的userid
	UserID string `json:"userid"  `
	// Status 发送状态：0-未发送 2-已发送
	Status int `json:"status"  `
	// SendTime 发送时间，未发送时不返回
	SendTime int `json:"send_time"  `
}

// GetGroupMessageSendResultRequest is request of Client.GetGroupMessageSendResult
type GetGroupMessageSendResultRequest struct {
	// MessageID 群发消息的id，通过获取群发记录列表接口返回
	MessageID string `json:"msgid"  validate:"required"`
	// UserID 发送成员userid，通过获取群发成员发送任务列表接口返回
	UserID string `json:"userid"  validate:"required"`
	// Limit 返回的最大记录数，整型，最大值1000，默认值500，超过最大值时取默认值
	Limit int `json:"limit"  `
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"  `
}

// GetGroupMessageSendResultResponse is response of Client.GetGroupMessageSendResult
type GetGroupMessageSendResultResponse struct {
	// NextCursor 分页游标，再下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则返回空
	NextCursor string `json:"next_cursor"  `
	// SendList 群成员发送结果列表
	SendList []GetGroupMessageSendResultResponseSendList `json:"send_list"  `
}

// GetGroupMessageSendResultResponseSendList is model of GetGroupMessageSendResultResponse.SendList
type GetGroupMessageSendResultResponseSendList struct {
	// ExternalUserID 外部联系人userid，群发消息到企业的客户群不返回该字段
	ExternalUserID string `json:"external_userid"  `
	// ChatID 外部客户群id，群发消息到客户不返回该字段
	ChatID string `json:"chat_id"  `
	// UserID 企业服务人员的userid
	UserID string `json:"userid"  `
	// Status 发送状态：0-未发送 1-已发送 2-因客户不是好友导致发送失败 3-因客户已经收到其他群发消息导致发送失败
	Status int `json:"status"  `
	// SendTime 发送时间，发送状态为1时返回
	SendTime int `json:"send_time"  `
}

// SendWelcomeMessageRequest is request of Client.SendWelcomeMessage
type SendWelcomeMessageRequest struct {
	// WelcomeCode 通过添加外部联系人事件推送给企业的发送欢迎语的凭证，有效期为20秒
	WelcomeCode string `json:"welcome_code"  validate:"required"`
	// Attachments 附件，最多可添加9个附件
	Attachments []SendWelcomeMessageRequestAttachments `json:"attachments"  `

	Text SendWelcomeMessageRequestText `json:"text"  `

	Image SendWelcomeMessageRequestImage `json:"image"  `

	Link SendWelcomeMessageRequestLink `json:"link"  `

	MiniProgram SendWelcomeMessageRequestMiniProgram `json:"miniprogram"  `

	Video SendWelcomeMessageRequestVideo `json:"video"  `

	File SendWelcomeMessageRequestFile `json:"file"  `
}

// SendWelcomeMessageRequestText is model of SendWelcomeMessageRequest.Text
type SendWelcomeMessageRequestText struct {
	// Content 消息文本内容,最长为4000字节
	Content string `json:"content"  `
}

// SendWelcomeMessageRequestAttachments is model of SendWelcomeMessageRequest.Attachments
type SendWelcomeMessageRequestAttachments struct {
	// MessageType 附件类型，可选image、link、miniprogram或者video
	MessageType string `json:"msgtype"  validate:"required"`
}

// SendWelcomeMessageRequestImage is model of SendWelcomeMessageRequest.Image
type SendWelcomeMessageRequestImage struct {
	// MediaID 图片的media_id，可以通过素材管理接口获得
	MediaID string `json:"media_id"  `
	// PicURL 图片的链接，仅可使用上传图片接口得到的链接
	PicURL string `json:"pic_url"  `
}

// SendWelcomeMessageRequestLink is model of SendWelcomeMessageRequest.Link
type SendWelcomeMessageRequestLink struct {
	// Title 图文消息标题，最长为128字节
	Title string `json:"title"  validate:"required"`
	// PicURL 图文消息封面的url
	PicURL string `json:"picurl"  `
	// Desc 图文消息的描述，最长为512字节
	Desc string `json:"desc"  `
	// URL 图文消息的链接
	URL string `json:"url"  validate:"required"`
}

// SendWelcomeMessageRequestMiniProgram is model of SendWelcomeMessageRequest.MiniProgram
type SendWelcomeMessageRequestMiniProgram struct {
	// Title 小程序消息标题，最长为64字节
	Title string `json:"title"  validate:"required"`
	// PicMediaID 小程序消息封面的mediaid，封面图建议尺寸为520*416
	PicMediaID string `json:"pic_media_id"  validate:"required"`
	// AppID 小程序appid，必须是关联到企业的小程序应用
	AppID string `json:"appid"  validate:"required"`
	// Page 小程序page路径
	Page string `json:"page"  validate:"required"`
}

// SendWelcomeMessageRequestVideo is model of SendWelcomeMessageRequest.Video
type SendWelcomeMessageRequestVideo struct {
	// MediaID 视频的media_id，可以通过素材管理接口获得
	MediaID string `json:"media_id"  validate:"required"`
}

// SendWelcomeMessageRequestFile is model of SendWelcomeMessageRequest.File
type SendWelcomeMessageRequestFile struct {
	// MediaID 文件的media_id,  可以通过素材管理接口获得
	MediaID string `json:"media_id"  validate:"required"`
}

// AddGroupWelcomeTemplateRequest is request of Client.AddGroupWelcomeTemplate
type AddGroupWelcomeTemplateRequest struct {
	// AgentID 授权方安装的应用agentid。仅旧的第三方多应用套件需要填此参数
	AgentID int `json:"agentid"  `
	// Notify 是否通知成员将这条入群欢迎语应用到客户群中，0-不通知，1-通知， 不填则通知
	Notify int `json:"notify"  `

	Text AddGroupWelcomeTemplateRequestText `json:"text"  `

	Image AddGroupWelcomeTemplateRequestImage `json:"image"  `

	Link AddGroupWelcomeTemplateRequestLink `json:"link"  `

	MiniProgram AddGroupWelcomeTemplateRequestMiniProgram `json:"miniprogram"  `

	File AddGroupWelcomeTemplateRequestFile `json:"file"  `

	Video AddGroupWelcomeTemplateRequestVideo `json:"video"  `
}

// AddGroupWelcomeTemplateResponse is response of Client.AddGroupWelcomeTemplate
type AddGroupWelcomeTemplateResponse struct {
	// TemplateID 欢迎语素材id
	TemplateID string `json:"template_id"  `
}

// AddGroupWelcomeTemplateRequestText is model of AddGroupWelcomeTemplateRequest.Text
type AddGroupWelcomeTemplateRequestText struct {
	// Content 消息文本内容,最长为3000字节
	Content string `json:"content"  `
}

// AddGroupWelcomeTemplateRequestImage is model of AddGroupWelcomeTemplateRequest.Image
type AddGroupWelcomeTemplateRequestImage struct {
	// MediaID 图片的media_id，可以通过素材管理接口获得
	MediaID string `json:"media_id"  `
	// PicURL 图片的链接，仅可使用上传图片接口得到的链接
	PicURL string `json:"pic_url"  `
}

// AddGroupWelcomeTemplateRequestLink is model of AddGroupWelcomeTemplateRequest.Link
type AddGroupWelcomeTemplateRequestLink struct {
	// Title 图文消息标题，最长为128字节
	Title string `json:"title"  validate:"required"`
	// PicURL 图文消息封面的url
	PicURL string `json:"picurl"  `
	// Desc 图文消息的描述，最长为512字节
	Desc string `json:"desc"  `
	// URL 图文消息的链接
	URL string `json:"url"  validate:"required"`
}

// AddGroupWelcomeTemplateRequestMiniProgram is model of AddGroupWelcomeTemplateRequest.MiniProgram
type AddGroupWelcomeTemplateRequestMiniProgram struct {
	// Title 小程序消息标题，最长为64字节
	Title string `json:"title"  validate:"required"`
	// PicMediaID 小程序消息封面的mediaid，封面图建议尺寸为520*416
	PicMediaID string `json:"pic_media_id"  validate:"required"`
	// AppID 小程序appid，必须是关联到企业的小程序应用
	AppID string `json:"appid"  validate:"required"`
	// Page 小程序page路径
	Page string `json:"page"  validate:"required"`
}

// AddGroupWelcomeTemplateRequestFile is model of AddGroupWelcomeTemplateRequest.File
type AddGroupWelcomeTemplateRequestFile struct {
	// MediaID 文件id，可以通过素材管理接口获得
	MediaID string `json:"media_id"  validate:"required"`
}

// AddGroupWelcomeTemplateRequestVideo is model of AddGroupWelcomeTemplateRequest.Video
type AddGroupWelcomeTemplateRequestVideo struct {
	// MediaID 视频媒体文件id，可以通过素材管理接口获得
	MediaID string `json:"media_id"  validate:"required"`
}

// EditGroupWelcomeTemplateRequest is request of Client.EditGroupWelcomeTemplate
type EditGroupWelcomeTemplateRequest struct {
	// TemplateID 欢迎语素材id
	TemplateID string `json:"template_id"  validate:"required"`
	// AgentID 授权方安装的应用agentid。仅旧的第三方多应用套件需要填此参数
	AgentID int `json:"agentid"  `

	Text EditGroupWelcomeTemplateRequestText `json:"text"  `

	Image EditGroupWelcomeTemplateRequestImage `json:"image"  `

	Link EditGroupWelcomeTemplateRequestLink `json:"link"  `

	MiniProgram EditGroupWelcomeTemplateRequestMiniProgram `json:"miniprogram"  `

	File EditGroupWelcomeTemplateRequestFile `json:"file"  `

	Video EditGroupWelcomeTemplateRequestVideo `json:"video"  `
}

// EditGroupWelcomeTemplateRequestText is model of EditGroupWelcomeTemplateRequest.Text
type EditGroupWelcomeTemplateRequestText struct {
	// Content 消息文本内容,最长为4000字节
	Content string `json:"content"  `
}

// EditGroupWelcomeTemplateRequestImage is model of EditGroupWelcomeTemplateRequest.Image
type EditGroupWelcomeTemplateRequestImage struct {
	// MediaID 图片的media_id，可以通过素材管理接口获得
	MediaID string `json:"media_id"  `
	// PicURL 图片的链接，仅可使用上传图片接口得到的链接
	PicURL string `json:"pic_url"  `
}

// EditGroupWelcomeTemplateRequestLink is model of EditGroupWelcomeTemplateRequest.Link
type EditGroupWelcomeTemplateRequestLink struct {
	// Title 图文消息标题，最长为128字节
	Title string `json:"title"  validate:"required"`
	// PicURL 图文消息封面的url
	PicURL string `json:"picurl"  `
	// Desc 图文消息的描述，最长为512字节
	Desc string `json:"desc"  `
	// URL 图文消息的链接
	URL string `json:"url"  validate:"required"`
}

// EditGroupWelcomeTemplateRequestMiniProgram is model of EditGroupWelcomeTemplateRequest.MiniProgram
type EditGroupWelcomeTemplateRequestMiniProgram struct {
	// Title 小程序消息标题，最长为64字节
	Title string `json:"title"  validate:"required"`
	// PicMediaID 小程序消息封面的mediaid，封面图建议尺寸为520*416
	PicMediaID string `json:"pic_media_id"  validate:"required"`
	// AppID 小程序appid，必须是关联到企业的小程序应用
	AppID string `json:"appid"  validate:"required"`
	// Page 小程序page路径
	Page string `json:"page"  validate:"required"`
}

// EditGroupWelcomeTemplateRequestFile is model of EditGroupWelcomeTemplateRequest.File
type EditGroupWelcomeTemplateRequestFile struct {
	// MediaID 文件id，可以通过素材管理接口获得
	MediaID string `json:"media_id"  validate:"required"`
}

// EditGroupWelcomeTemplateRequestVideo is model of EditGroupWelcomeTemplateRequest.Video
type EditGroupWelcomeTemplateRequestVideo struct {
	// MediaID 视频媒体文件id，可以通过素材管理接口获得
	MediaID string `json:"media_id"  validate:"required"`
}

// GetGroupWelcomeTemplateRequest is request of Client.GetGroupWelcomeTemplate
type GetGroupWelcomeTemplateRequest struct {
	// TemplateID 群欢迎语的素材id
	TemplateID string `json:"template_id"  validate:"required"`
}

// GetGroupWelcomeTemplateResponse is response of Client.GetGroupWelcomeTemplate
type GetGroupWelcomeTemplateResponse struct {
	Text GetGroupWelcomeTemplateResponseText `json:"text"  `

	Image GetGroupWelcomeTemplateResponseImage `json:"image"  `

	Link GetGroupWelcomeTemplateResponseLink `json:"link"  `

	MiniProgram GetGroupWelcomeTemplateResponseMiniProgram `json:"miniprogram"  `

	File GetGroupWelcomeTemplateResponseFile `json:"file"  `

	Video GetGroupWelcomeTemplateResponseVideo `json:"video"  `
}

// GetGroupWelcomeTemplateResponseText is model of GetGroupWelcomeTemplateResponse.Text
type GetGroupWelcomeTemplateResponseText struct {
	// Content 消息文本内容
	Content string `json:"content"  `
}

// GetGroupWelcomeTemplateResponseImage is model of GetGroupWelcomeTemplateResponse.Image
type GetGroupWelcomeTemplateResponseImage struct {
	// PicURL 图片的url
	PicURL string `json:"pic_url"  `
}

// GetGroupWelcomeTemplateResponseLink is model of GetGroupWelcomeTemplateResponse.Link
type GetGroupWelcomeTemplateResponseLink struct {
	// Title 图文消息标题
	Title string `json:"title"  `
	// PicURL 图文消息封面的url
	PicURL string `json:"picurl"  `
	// Desc 图文消息的描述
	Desc string `json:"desc"  `
	// URL 图文消息的链接
	URL string `json:"url"  `
}

// GetGroupWelcomeTemplateResponseMiniProgram is model of GetGroupWelcomeTemplateResponse.MiniProgram
type GetGroupWelcomeTemplateResponseMiniProgram struct {
	// Title 小程序消息标题
	Title string `json:"title"  `
	// PicMediaID 小程序消息封面的mediaid
	PicMediaID string `json:"pic_media_id"  `
	// AppID 小程序appid
	AppID string `json:"appid"  `
	// Page 小程序page路径
	Page string `json:"page"  `
}

// GetGroupWelcomeTemplateResponseFile is model of GetGroupWelcomeTemplateResponse.File
type GetGroupWelcomeTemplateResponseFile struct {
	// MediaID 文件id，可以通过素材管理接口获得
	MediaID string `json:"media_id"  `
}

// GetGroupWelcomeTemplateResponseVideo is model of GetGroupWelcomeTemplateResponse.Video
type GetGroupWelcomeTemplateResponseVideo struct {
	// MediaID 视频媒体文件id，可以通过素材管理接口获得
	MediaID string `json:"media_id"  `
}

// DeleteGroupWelcomeTemplateRequest is request of Client.DeleteGroupWelcomeTemplate
type DeleteGroupWelcomeTemplateRequest struct {
	// TemplateID 群欢迎语的素材id
	TemplateID string `json:"template_id"  validate:"required"`
	// AgentID 授权方安装的应用agentid。仅旧的第三方多应用套件需要填此参数
	AgentID int `json:"agentid"  `
}

// GetUserBehaviorDataRequest is request of Client.GetUserBehaviorData
type GetUserBehaviorDataRequest struct {
	// UserID 成员ID列表，最多100个
	UserID []string `json:"userid"  `
	// PartyID 部门ID列表，最多100个
	PartyID []int `json:"partyid"  `
	// StartTime 数据起始时间
	StartTime int `json:"start_time"  validate:"required"`
	// EndTime 数据结束时间
	EndTime int `json:"end_time"  validate:"required"`
}

// GetUserBehaviorDataResponse is response of Client.GetUserBehaviorData
type GetUserBehaviorDataResponse struct {
	BehaviorData []GetUserBehaviorDataResponseBehaviorData `json:"behavior_data"  `
}

// GetUserBehaviorDataResponseBehaviorData is model of GetUserBehaviorDataResponse.BehaviorData
type GetUserBehaviorDataResponseBehaviorData struct {
	// StatTime 数据日期，为当日0点的时间戳
	StatTime int `json:"stat_time"  `
	// NewApplyCount 发起申请数，成员通过「搜索手机号」、「扫一扫」、「从微信好友中添加」、「从群聊中添加」、「添加共享、分配给我的客户」、「添加单向、双向删除好友关系的好友」、「从新的联系人推荐中添加」等渠道主动向客户发起的好友申请数量。
	NewApplyCount int `json:"new_apply_cnt"  `
	// NewContactCount 新增客户数，成员新添加的客户数量。
	NewContactCount int `json:"new_contact_cnt"  `
	// ChatCount 聊天总数， 成员有主动发送过消息的单聊总数。
	ChatCount int `json:"chat_cnt"  `
	// MessageCount 发送消息数，成员在单聊中发送的消息总数。
	MessageCount int `json:"message_cnt"  `
	// ReplyPercentage 已回复聊天占比，浮点型，客户主动发起聊天后，成员在一个自然日内有回复过消息的聊天数/客户主动发起的聊天数比例，不包括群聊，仅在确有聊天时返回。
	ReplyPercentage float64 `json:"reply_percentage"  `
	// AvgReplyTime 平均首次回复时长，单位为分钟，即客户主动发起聊天后，成员在一个自然日内首次回复的时长间隔为首次回复时长，所有聊天的首次回复总时长/已回复的聊天总数即为平均首次回复时长，不包括群聊，仅在确有聊天时返回。
	AvgReplyTime int `json:"avg_reply_time"  `
	// NegativeFeedbackCount 删除/拉黑成员的客户数，即将成员删除或加入黑名单的客户数。
	NegativeFeedbackCount int `json:"negative_feedback_cnt"  `
}

// GroupChatStatisticRequest is request of Client.GroupChatStatistic
type GroupChatStatisticRequest struct {
	// DayBeginTime 起始日期的时间戳，填当天的0时0分0秒（否则系统自动处理为当天的0分0秒）。取值范围：昨天至前180天。
	DayBeginTime int `json:"day_begin_time"  validate:"required"`
	// DayEndTime 结束日期的时间戳，填当天的0时0分0秒（否则系统自动处理为当天的0分0秒）。取值范围：昨天至前180天。如果不填，默认同 day_begin_time（即默认取一天的数据）
	DayEndTime int `json:"day_end_time"  `
	// OwnerFilter 群主过滤。如果不填，表示获取应用可见范围内全部群主的数据（但是不建议这么用，如果可见范围人数超过1000人，为了防止数据包过大，会报错 81017）
	OwnerFilter GroupChatStatisticRequestOwnerFilter `json:"owner_filter"  validate:"required"`
	// OrderBy 排序方式。1 - 新增群的数量2 - 群总数3 - 新增群人数4 - 群总人数默认为1
	OrderBy int `json:"order_by"  `
	// OrderAsc 是否升序。0-否；1-是。默认降序
	OrderAsc int `json:"order_asc"  `
	// Offset 分页，偏移量, 默认为0
	Offset int `json:"offset"  `
	// Limit 分页，预期请求的数据量，默认为500，取值范围 1 ~ 1000
	Limit int `json:"limit"  `
}

// GroupChatStatisticResponse is response of Client.GroupChatStatistic
type GroupChatStatisticResponse struct {
	// Total 命中过滤条件的记录总个数
	Total int `json:"total"  `
	// NextOffset 当前分页的下一个offset。当next_offset和total相等时，说明已经取完所有
	NextOffset int `json:"next_offset"  `
	// Items 记录列表。表示某个群主所拥有的客户群的统计数据
	Items []GroupChatStatisticResponseItems `json:"items"  `
}

// GroupChatStatisticRequestOwnerFilter is model of GroupChatStatisticRequest.OwnerFilter
type GroupChatStatisticRequestOwnerFilter struct {
	// UserIDList 群主ID列表。最多100个
	UserIDList []string `json:"userid_list"  validate:"required"`
}

// GroupChatStatisticResponseItems is model of GroupChatStatisticResponse.Items
type GroupChatStatisticResponseItems struct {
	// Owner 群主ID
	Owner string `json:"owner"  `
	// Data 详情
	Data GroupChatStatisticResponseItemsData `json:"data"  `
}

// GroupChatStatisticResponseItemsData is model of GroupChatStatisticResponseItems.Data
type GroupChatStatisticResponseItemsData struct {
	// NewChatCount 新增客户群数量
	NewChatCount int `json:"new_chat_cnt"  `
	// ChatTotal 截至当天客户群总数量
	ChatTotal int `json:"chat_total"  `
	// ChatHasMessage 截至当天有发过消息的客户群数量
	ChatHasMessage int `json:"chat_has_msg"  `
	// NewMemberCount 客户群新增群人数。
	NewMemberCount int `json:"new_member_cnt"  `
	// MemberTotal 截至当天客户群总人数
	MemberTotal int `json:"member_total"  `
	// MemberHasMessage 截至当天有发过消息的群成员数
	MemberHasMessage int `json:"member_has_msg"  `
	// MessageTotal 截至当天客户群消息总数
	MessageTotal int `json:"msg_total"  `
	// MigrateTraineeChatCount 截至当天新增迁移群数(仅教培行业返回)
	MigrateTraineeChatCount int `json:"migrate_trainee_chat_cnt"  `
}

// StatisticGroupByDayRequest is request of Client.StatisticGroupByDay
type StatisticGroupByDayRequest struct {
	// DayBeginTime 起始日期的时间戳，填当天的0时0分0秒（否则系统自动处理为当天的0分0秒）。取值范围：昨天至前180天。
	DayBeginTime int `json:"day_begin_time"  validate:"required"`
	// DayEndTime 结束日期的时间戳，填当天的0时0分0秒（否则系统自动处理为当天的0分0秒）。取值范围：昨天至前180天。如果不填，默认同 day_begin_time（即默认取一天的数据）
	DayEndTime int `json:"day_end_time"  `
	// OwnerFilter 群主过滤。如果不填，表示获取应用可见范围内全部群主的数据（但是不建议这么用，如果可见范围人数超过1000人，为了防止数据包过大，会报错 81017）
	OwnerFilter StatisticGroupByDayRequestOwnerFilter `json:"owner_filter"  validate:"required"`
}

// StatisticGroupByDayResponse is response of Client.StatisticGroupByDay
type StatisticGroupByDayResponse struct {
	// Items 记录列表。表示某个自然日客户群的统计数据
	Items []StatisticGroupByDayResponseItems `json:"items"  `
}

// StatisticGroupByDayRequestOwnerFilter is model of StatisticGroupByDayRequest.OwnerFilter
type StatisticGroupByDayRequestOwnerFilter struct {
	// UserIDList 群主ID列表。最多100个
	UserIDList []string `json:"userid_list"  validate:"required"`
}

// StatisticGroupByDayResponseItems is model of StatisticGroupByDayResponse.Items
type StatisticGroupByDayResponseItems struct {
	// StatTime 数据日期，为当日0点的时间戳
	StatTime int `json:"stat_time"  `
	// Data 详情
	Data StatisticGroupByDayResponseItemsData `json:"data"  `
}

// StatisticGroupByDayResponseItemsData is model of StatisticGroupByDayResponseItems.Data
type StatisticGroupByDayResponseItemsData struct {
	// NewChatCount 新增客户群数量
	NewChatCount int `json:"new_chat_cnt"  `
	// ChatTotal 截至当天客户群总数量
	ChatTotal int `json:"chat_total"  `
	// ChatHasMessage 截至当天有发过消息的客户群数量
	ChatHasMessage int `json:"chat_has_msg"  `
	// NewMemberCount 客户群新增群人数。
	NewMemberCount int `json:"new_member_cnt"  `
	// MemberTotal 截至当天客户群总人数
	MemberTotal int `json:"member_total"  `
	// MemberHasMessage 截至当天有发过消息的群成员数
	MemberHasMessage int `json:"member_has_msg"  `
	// MessageTotal 截至当天客户群消息总数
	MessageTotal int `json:"msg_total"  `
	// MigrateTraineeChatCount 截至当天新增迁移群数(仅教培行业返回)
	MigrateTraineeChatCount int `json:"migrate_trainee_chat_cnt"  `
}
