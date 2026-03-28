package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// ListExternalPayBill 获取对外收款记录
// 企业和服务商可通过此接口获取企业的对外收款记录。
//
// see https://developer.work.weixin.qq.com/document/path/93667
func (c *Client) ListExternalPayBill(r *ListExternalPayBillRequest, opts ...any) (out ListExternalPayBillResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalpay/get_bill_list",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ListExternalPayBillRequest is request of Client.ListExternalPayBill
type ListExternalPayBillRequest struct {
	// BeginTime 收款记录开始时间戳，单位为秒
	BeginTime int `json:"begin_time" validate:"required"`
	// EndTime 收款记录结束时间戳，单位为秒
	EndTime int `json:"end_time" validate:"required"`
	// PayeeUserID 企业收款成员userid，不填则为全部成员
	PayeeUserID string `json:"payee_userid"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"`
	// Limit 返回的最大记录数，整型，最大值1000
	Limit int `json:"limit"`
}

// ListExternalPayBillResponse is response of Client.ListExternalPayBill
type ListExternalPayBillResponse struct {
	// NextCursor 分页游标，在下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则不返回该字段
	NextCursor string `json:"next_cursor"`
	// BillList 交易单详情列表
	BillList []map[string]any `json:"bill_list"`
}

// GetFundFlow 获取资金流水
// 企业可通过此接口获取企业的资金流水。
//
// see https://developer.work.weixin.qq.com/document/path/98100
func (c *Client) GetFundFlow(r *GetFundFlowRequest, opts ...any) (out GetFundFlowResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalpay/get_fund_flow",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetFundFlowRequest is request of Client.GetFundFlow
type GetFundFlowRequest struct {
	// BeginTime 资金流水记录开始时间
	BeginTime int `json:"begin_time" validate:"required"`
	// EndTime 资金流水记录结束时间
	EndTime int `json:"end_time" validate:"required"`
	// MchID 商户号ID，若不填写则拉取所有商户号的资金流水
	MchID string `json:"mch_id"`
	// Cursor 用于分页查询的游标，字符串类型，由上一次调用返回，首次调用可不填
	Cursor string `json:"cursor"`
	// Limit 返回的最大记录数，默认值100，最大不超过200
	Limit int `json:"limit"`
}

// GetFundFlowResponse is response of Client.GetFundFlow
type GetFundFlowResponse struct {
	// NextCursor 分页游标，在下次请求时填写以获取之后分页的记录，如果已经没有更多的数据则不返回该字段
	NextCursor string `json:"next_cursor"`
	// FundFlowList 资金流水记录列表
	FundFlowList []map[string]any `json:"fund_flow_list"`
}

// GetPaymentInfo 获取收款项目的商户单号
// 获取对外收款项目中每笔收款的商户单号
//
// see https://developer.work.weixin.qq.com/document/path/96076
func (c *Client) GetPaymentInfo(r *GetPaymentInfoRequest, opts ...any) (out GetPaymentInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalpay/get_payment_info",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetPaymentInfoRequest is request of Client.GetPaymentInfo
type GetPaymentInfoRequest struct {
	// PaymentID 收款项目单号。在发起对外收款中返回
	PaymentID string `json:"payment_id" validate:"required"`
}

// GetPaymentInfoResponse is response of Client.GetPaymentInfo
type GetPaymentInfoResponse struct {
	// OutTradeNoList 收款单号列表。每笔支付对应一个收款单号
	OutTradeNoList []map[string]any `json:"out_trade_no_list"`
}

// GetMerchantDetail 查询商户号详情
// 通过该接口可以查询已绑定商户号信息以及商户号使用范围
//
// see https://developer.work.weixin.qq.com/document/path/93666
func (c *Client) GetMerchantDetail(r *GetMerchantDetailRequest, opts ...any) (out GetMerchantDetailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/externalpay/getmerchant",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMerchantDetailRequest is request of Client.GetMerchantDetail
type GetMerchantDetailRequest struct {
	// MchID 微信支付商户号，不超过32字节
	MchID string `json:"mch_id" validate:"required"`
}

// GetMerchantDetailResponse is response of Client.GetMerchantDetail
type GetMerchantDetailResponse struct {
	// MchID 微信支付商户号，不超过32字节
	MchID string `json:"mch_id"`
	// MerchantName 微信支付商户号全称，不超过50字节
	MerchantName string `json:"merchant_name"`
	// BindStatus 商户号绑定状态。1:申请中 2:已绑定 3:已撤销
	BindStatus int `json:"bind_status"`
	// AllowUseScope 该商户号使用范围，仅当商户号已绑定时才返回
	AllowUseScope any `json:"allow_use_scope"`
}

// ApplyMch 提交创建对外收款账户的申请单
// 企业可以通过这个接口快速递交材料，用于创建对外收款账户申请单
//
// see https://developer.work.weixin.qq.com/document/path/99106
func (c *Client) ApplyMch(r *ApplyMchRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/miniapppay/apply_mch",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// ApplyMchRequest is request of Client.ApplyMch
type ApplyMchRequest struct {
	// OutRequestNo 业务申请编号，长度限制为1~32个字符
	OutRequestNo string `json:"out_request_no" validate:"required"`
	// OrganizationType 主体类型 (0:企业, 1:个体, 2:社会团体组织, 3:事业单位)
	OrganizationType int `json:"organization_type" validate:"required"`
	// BusinessLicenseInfo 营业执照/登记证书信息
	BusinessLicenseInfo any `json:"business_license_info" validate:"required"`
	// FinanceInstitutionInfo 金融机构许可证信息，主体是金融机构时必填
	FinanceInstitutionInfo any `json:"finance_institution_info"`
	// MerchantShortName 商户简称，UTF-8格式，最多21个汉字
	MerchantShortName string `json:"merchant_short_name" validate:"required"`
	// IDCardInfo 经营者/法人证件信息
	IDCardInfo any `json:"id_card_info" validate:"required"`
	// Owner 经营者/法人是否为受益人，主体为企业时必填
	Owner bool `json:"owner"`
	// UboInfo 受益人证件信息，企业且非受益人时必填
	UboInfo any `json:"ubo_info"`
	// ContactInfo 超级管理员信息
	ContactInfo any `json:"contact_info" validate:"required"`
	// AccountInfo 结算账户信息
	AccountInfo any `json:"account_info" validate:"required"`
	// SalesSceneInfo 经营场景证明
	SalesSceneInfo any `json:"sales_scene_info" validate:"required"`
	// BusinessID 经营范围
	BusinessID int `json:"business_id" validate:"required"`
	// Qualifications 特殊资质
	Qualifications any `json:"qualifications"`
	// BusinessAdditionPics 补充材料
	BusinessAdditionPics any `json:"business_addition_pics"`
	// UserID 提现人员，成员UserID
	UserID string `json:"userid" validate:"required"`
}

// GetApplymentStatus 查询申请单状态
// 通过这个接口查询已提交的申请单状态，只能使用该应用本身提交的申请单号。
//
// see https://developer.work.weixin.qq.com/document/path/98974
func (c *Client) GetApplymentStatus(r *GetApplymentStatusRequest, opts ...any) (out GetApplymentStatusResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/miniapppay/get_applyment_status",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetApplymentStatusRequest is request of Client.GetApplymentStatus
type GetApplymentStatusRequest struct {
	// OutRequestNo 业务申请编号，长度限制为1~32个字符
	OutRequestNo string `json:"out_request_no" validate:"required"`
}

// GetApplymentStatusResponse is response of Client.GetApplymentStatus
type GetApplymentStatusResponse struct {
	// Status 申请单的具体状态
	Status any `json:"status"`
	// ApplyState 申请单当前所处阶段
	ApplyState int `json:"apply_state"`
	// RealSignState 当前签约阶段
	RealSignState int `json:"real_sign_state"`
	// RejectReason 驳回理由
	RejectReason string `json:"reject_reason"`
}

// UploadMiniAppImage 提交图片
// 上传图片得到图片ID，该ID30天后过期。返回的图片ID用于提交创建对外收款账户的申请单中填写图片相关字段。
//
// see https://developer.work.weixin.qq.com/document/path/98972
func (c *Client) UploadMiniAppImage(r *UploadMiniAppImageRequest, opts ...any) (out UploadMiniAppImageResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/miniapppay/upload_image",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UploadMiniAppImageRequest is request of Client.UploadMiniAppImage
type UploadMiniAppImageRequest struct {}

// UploadMiniAppImageResponse is response of Client.UploadMiniAppImage
type UploadMiniAppImageResponse struct {
	// OpenWxPayMediaID 上传后得到的图片id，30天后过期
	OpenWxPayMediaID string `json:"open_wx_pay_media_id"`
}

// PayEmployee 向员工付款
// 通过企业微信向员工进行付款操作（基于微信支付接口）
//
// see https://developer.work.weixin.qq.com/document/path/90097
func (c *Client) PayEmployee(r *PayEmployeeRequest, opts ...any) (out PayEmployeeResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/mmpaymkttransfers/promotion/paywwsptrans2pocket",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// PayEmployeeRequest is request of Client.PayEmployee
type PayEmployeeRequest struct {
	// AppID 微信分配的公众账号ID
	AppID string `json:"appid" validate:"required"`
	// MchID 微信支付分配的商户号
	MchID string `json:"mch_id" validate:"required"`
	// DeviceInfo 微信支付分配的终端设备号
	DeviceInfo string `json:"device_info"`
	// NonceStr 随机字符串，不长于32位
	NonceStr string `json:"nonce_str" validate:"required"`
	// Sign 微信支付签名
	Sign string `json:"sign" validate:"required"`
	// PartnerTradeNo 商户订单号，需保持唯一性
	PartnerTradeNo string `json:"partner_trade_no" validate:"required"`
	// OpenID 用户openid
	OpenID string `json:"openid" validate:"required"`
	// CheckName 校验用户姓名选项 (NO_CHECK/FORCE_CHECK)
	CheckName string `json:"check_name" validate:"required"`
	// ReUserName 收款用户真实姓名
	ReUserName string `json:"re_user_name"`
	// Amount 企业微信企业付款金额，单位为分
	Amount int `json:"amount" validate:"required"`
	// Desc 向员工付款说明信息
	Desc string `json:"desc" validate:"required"`
	// SpbillCreateIP 调用接口的机器Ip地址
	SpbillCreateIP string `json:"spbill_create_ip" validate:"required"`
	// WorkwxSign 企业微信签名
	WorkwxSign string `json:"workwx_sign" validate:"required"`
	// WwMsgType 付款消息类型 (NORMAL_MSG/APPROVAL_MSG)
	WwMsgType string `json:"ww_msg_type" validate:"required"`
	// ApprovalNumber 审批单号
	ApprovalNumber string `json:"approval_number"`
	// ApprovalType 审批类型
	ApprovalType int `json:"approval_type"`
	// ActName 项目名称
	ActName string `json:"act_name" validate:"required"`
	// AgentID 付款的应用id
	AgentID int `json:"agentid"`
}

// PayEmployeeResponse is response of Client.PayEmployee
type PayEmployeeResponse struct {
	// ReturnCode 返回状态码 (SUCCESS/FAIL)
	ReturnCode string `json:"return_code"`
	// ReturnMsg 返回信息
	ReturnMsg string `json:"return_msg"`
	// AppID 商户appid
	AppID string `json:"appid"`
	// MchID 商户号
	MchID string `json:"mch_id"`
	// DeviceInfo 设备号
	DeviceInfo string `json:"device_info"`
	// NonceStr 随机字符串
	NonceStr string `json:"nonce_str"`
	// ResultCode 业务结果 (SUCCESS/FAIL)
	ResultCode string `json:"result_code"`
	// ErrCode 错误代码
	ErrCode string `json:"err_code"`
	// ErrCodeDes 错误代码描述
	ErrCodeDes string `json:"err_code_des"`
	// PartnerTradeNo 商户订单号
	PartnerTradeNo string `json:"partner_trade_no"`
	// PaymentNo 微信订单号
	PaymentNo string `json:"payment_no"`
	// PaymentTime 微信支付成功时间
	PaymentTime string `json:"payment_time"`
}

// QueryWxPaymentRecord 查询付款记录
// 用于商户的企业微信企业付款操作进行结果查询，返回付款操作详细结果。仅支持查询30天内的订单。
//
// see https://developer.work.weixin.qq.com/document/path/90098
func (c *Client) QueryWxPaymentRecord(r *QueryWxPaymentRecordRequest, opts ...any) (out QueryWxPaymentRecordResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/mmpaymkttransfers/promotion/querywwsptrans2pocket",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// QueryWxPaymentRecordRequest is request of Client.QueryWxPaymentRecord
type QueryWxPaymentRecordRequest struct {
	// NonceStr 随机字符串，不长于32位
	NonceStr string `json:"nonce_str" validate:"required"`
	// Sign 微信支付签名
	Sign string `json:"sign" validate:"required"`
	// PartnerTradeNo 商户订单号
	PartnerTradeNo string `json:"partner_trade_no" validate:"required"`
	// MchID 商户号
	MchID string `json:"mch_id" validate:"required"`
	// AppID Appid
	AppID string `json:"appid" validate:"required"`
}

// QueryWxPaymentRecordResponse is response of Client.QueryWxPaymentRecord
type QueryWxPaymentRecordResponse struct {
	// ReturnCode 返回状态码 (SUCCESS/FAIL)
	ReturnCode string `json:"return_code"`
	// ReturnMsg 返回信息
	ReturnMsg string `json:"return_msg"`
	// ResultCode 业务结果 (SUCCESS/FAIL)
	ResultCode string `json:"result_code"`
	// ErrCode 错误代码
	ErrCode string `json:"err_code"`
	// ErrCodeDes 错误代码描述
	ErrCodeDes string `json:"err_code_des"`
	// PartnerTradeNo 商户订单号
	PartnerTradeNo string `json:"partner_trade_no"`
	// MchID 商户号
	MchID string `json:"mch_id"`
	// DetailID 付款单号
	DetailID string `json:"detail_id"`
	// Status 转账状态 (SUCCESS/FAILED/PROCESSING)
	Status string `json:"status"`
	// Reason 失败原因
	Reason string `json:"reason"`
	// OpenID 收款用户openid
	OpenID string `json:"openid"`
	// TransferName 收款用户姓名
	TransferName string `json:"transfer_name"`
	// PaymentAmount 付款金额 (单位：分)
	PaymentAmount int `json:"payment_amount"`
	// TransferTime 转账时间
	TransferTime string `json:"transfer_time"`
	// Desc 付款描述
	Desc string `json:"desc"`
}

// QueryWorkWxRedpack 查询红包记录
// 查询企业微信红包发放记录的状态和详细信息
//
// see https://developer.work.weixin.qq.com/document/path/90095
func (c *Client) QueryWorkWxRedpack(r *QueryWorkWxRedpackRequest, opts ...any) (out QueryWorkWxRedpackResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/mmpaymkttransfers/queryworkwxredpack",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// QueryWorkWxRedpackRequest is request of Client.QueryWorkWxRedpack
type QueryWorkWxRedpackRequest struct {
	// NonceStr 随机字符串
	NonceStr string `json:"nonce_str" validate:"required"`
	// Sign 微信支付签名
	Sign string `json:"sign" validate:"required"`
	// MchBillno 商户订单号
	MchBillno string `json:"mch_billno" validate:"required"`
	// MchID 商户号
	MchID string `json:"mch_id" validate:"required"`
	// AppID Appid
	AppID string `json:"appid" validate:"required"`
}

// QueryWorkWxRedpackResponse is response of Client.QueryWorkWxRedpack
type QueryWorkWxRedpackResponse struct {
	// ReturnCode 返回状态码
	ReturnCode string `json:"return_code"`
	// ReturnMsg 返回信息
	ReturnMsg string `json:"return_msg"`
	// Sign 微信支付签名
	Sign string `json:"sign"`
	// ResultCode 业务结果
	ResultCode string `json:"result_code"`
	// ErrCode 错误代码
	ErrCode string `json:"err_code"`
	// ErrCodeDes 错误代码描述
	ErrCodeDes string `json:"err_code_des"`
	// MchBillno 商户订单号
	MchBillno string `json:"mch_billno"`
	// MchID 商户号
	MchID string `json:"mch_id"`
	// DetailID 红包单号
	DetailID string `json:"detail_id"`
	// Status 红包状态
	Status string `json:"status"`
	// SendType 发放类型
	SendType string `json:"send_type"`
	// TotalAmount 红包总金额
	TotalAmount int `json:"total_amount"`
	// Reason 失败原因
	Reason string `json:"reason"`
	// SendTime 红包发送时间
	SendTime string `json:"send_time"`
	// RefundTime 红包退款时间
	RefundTime string `json:"refund_time"`
	// RefundAmount 红包退款金额
	RefundAmount int `json:"refund_amount"`
	// Wishing 祝福语
	Wishing string `json:"wishing"`
	// Remark 活动描述
	Remark string `json:"remark"`
	// ActName 活动名称
	ActName string `json:"act_name"`
	// OpenID 领取红包的Openid
	OpenID string `json:"openid"`
	// Amount 金额
	Amount int `json:"amount"`
	// RcvTime 接收时间
	RcvTime string `json:"rcv_time"`
	// SenderName 发送者名称
	SenderName string `json:"sender_name"`
	// SenderHeaderMediaID 发送者头像
	SenderHeaderMediaID string `json:"sender_header_media_id"`
}

// SendWorkWxRedPacket 发放企业红包
// 向企业微信用户发放现金红包
//
// see https://developer.work.weixin.qq.com/document/path/90094
func (c *Client) SendWorkWxRedPacket(r *SendWorkWxRedPacketRequest, opts ...any) (out SendWorkWxRedPacketResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/mmpaymkttransfers/sendworkwxredpack",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// SendWorkWxRedPacketRequest is request of Client.SendWorkWxRedPacket
type SendWorkWxRedPacketRequest struct {
	// NonceStr 随机字符串，不长于32位
	NonceStr string `json:"nonce_str" validate:"required"`
	// Sign 微信支付签名
	Sign string `json:"sign" validate:"required"`
	// MchBillno 商户订单号（每个订单号必须唯一）
	MchBillno string `json:"mch_billno" validate:"required"`
	// MchID 微信支付分配的商户号
	MchID string `json:"mch_id" validate:"required"`
	// Wxappid 微信分配的公众账号ID（企业微信corpid）
	Wxappid string `json:"wxappid" validate:"required"`
	// SenderName 发送者名称，以个人名义发红包。与agentid互斥
	SenderName string `json:"sender_name"`
	// AgentID 发送红包的应用id，以企业应用的名义发红包。与sender_name互斥
	AgentID int `json:"agentid"`
	// SenderHeaderMediaID 发送者头像素材id
	SenderHeaderMediaID string `json:"sender_header_media_id"`
	// ReOpenID 接受红包的用户openid
	ReOpenID string `json:"re_openid" validate:"required"`
	// TotalAmount 金额，单位分
	TotalAmount int `json:"total_amount" validate:"required"`
	// Wishing 红包祝福语
	Wishing string `json:"wishing" validate:"required"`
	// ActName 项目名称
	ActName string `json:"act_name" validate:"required"`
	// Remark 备注信息
	Remark string `json:"remark" validate:"required"`
	// SceneID 发放红包使用场景，金额大于200或小于1元时必传
	SceneID string `json:"scene_id"`
	// WorkwxSign 企业微信签名
	WorkwxSign string `json:"workwx_sign" validate:"required"`
}

// SendWorkWxRedPacketResponse is response of Client.SendWorkWxRedPacket
type SendWorkWxRedPacketResponse struct {
	// ReturnCode 返回状态码
	ReturnCode string `json:"return_code"`
	// ReturnMsg 返回信息
	ReturnMsg string `json:"return_msg"`
	// Sign 微信支付签名
	Sign string `json:"sign"`
	// ResultCode 业务结果
	ResultCode string `json:"result_code"`
	// ErrCode 错误码信息
	ErrCode string `json:"err_code"`
	// ErrCodeDes 结果信息描述
	ErrCodeDes string `json:"err_code_des"`
	// MchBillno 商户订单号
	MchBillno string `json:"mch_billno"`
	// MchID 微信支付分配的商户号
	MchID string `json:"mch_id"`
	// Wxappid 商户appid
	Wxappid string `json:"wxappid"`
	// ReOpenID 接受收红包的用户openid
	ReOpenID string `json:"re_openid"`
	// TotalAmount 付款金额，单位分
	TotalAmount int `json:"total_amount"`
	// SendListid 红包订单的微信单号
	SendListid string `json:"send_listid"`
	// SenderName 红包发送者名称
	SenderName string `json:"sender_name"`
	// SenderHeaderMediaID 发送者头像素材id
	SenderHeaderMediaID string `json:"sender_header_media_id"`
}

