package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// CloseMiniAppPayOrder 关闭订单
// 关闭订单，用于支付失败或超时等情况
//
// see https://developer.work.weixin.qq.com/document/path/97324
func (c *Client) CloseMiniAppPayOrder(r *CloseMiniAppPayOrderRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/miniapppay/close_order",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CloseMiniAppPayOrderRequest is request of Client.CloseMiniAppPayOrder
type CloseMiniAppPayOrderRequest struct {
	// Mchid 二级商户号，由企业微信生成并下发
	Mchid string `json:"mchid" validate:"required"`
	// OutTradeNo 商户系统内部订单号
	OutTradeNo string `json:"out_trade_no" validate:"required"`
}

// CreateMiniAppOrder 小程序下单
// 商户系统先调用该接口通过企微后台生成预支付交易单，返回正确的预支付交易会话标识后再按小程序场景生成交易串调起支付。
//
// see https://developer.work.weixin.qq.com/document/path/97322
func (c *Client) CreateMiniAppOrder(r *CreateMiniAppOrderRequest, opts ...any) (out CreateMiniAppOrderResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/miniapppay/create_order",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// CreateMiniAppOrderRequest is request of Client.CreateMiniAppOrder
type CreateMiniAppOrderRequest struct {
	// AppID 二级商户申请的公众号或移动应用appid
	AppID string `json:"appid" validate:"required"`
	// Mchid 二级商户号，由企业微信生成并下发
	Mchid string `json:"mchid" validate:"required"`
	// OutTradeNo 商户系统内部订单号
	OutTradeNo string `json:"out_trade_no" validate:"required"`
	// Description 商品描述
	Description string `json:"description" validate:"required"`
	// Scenekey 用来统计企微成员发出小程序的交易业绩
	Scenekey string `json:"scenekey"`
	// Amount 订单金额信息
	Amount any `json:"amount" validate:"required"`
	// Payer 支付者标识
	Payer any `json:"payer" validate:"required"`
	// TimeExpire 订单失效时间，遵循rfc3339标准
	TimeExpire string `json:"time_expire"`
	// Attach 附加数据
	Attach string `json:"attach"`
	// GoodsTag 订单优惠标记
	GoodsTag string `json:"goods_tag"`
	// Detail 商品明细信息
	Detail any `json:"detail"`
	// SceneInfo 下单场景信息
	SceneInfo any `json:"scene_info" validate:"required"`
}

// CreateMiniAppOrderResponse is response of Client.CreateMiniAppOrder
type CreateMiniAppOrderResponse struct {
	// PrepayID 预支付交易会话标识
	PrepayID string `json:"prepay_id"`
}

// GetOrder 查询订单
// 商户可以通过商户订单号查询订单，获取订单状态等信息，完成下一步的业务逻辑。
//
// see https://developer.work.weixin.qq.com/document/path/97323
func (c *Client) GetOrder(r *GetOrderRequest, opts ...any) (out GetOrderResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/miniapppay/get_order",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetOrderRequest is request of Client.GetOrder
type GetOrderRequest struct {
	// Mchid 二级商户号，由企业微信生成并下发
	Mchid string `json:"mchid" validate:"required"`
	// OutTradeNo 商户系统内部订单号，只能是数字、大小写字母_-*且在同一个商户号下唯一
	OutTradeNo string `json:"out_trade_no" validate:"required"`
}

// GetOrderResponse is response of Client.GetOrder
type GetOrderResponse struct {
	// Mchid 二级商户号
	Mchid string `json:"mchid"`
	// OutTradeNo 商户系统内部订单号
	OutTradeNo string `json:"out_trade_no"`
	// TradeState 交易状态，枚举值：SUCCESS, REFUND, NOTPAY等
	TradeState string `json:"trade_state"`
	// TradeStateDesc 交易状态描述
	TradeStateDesc string `json:"trade_state_desc"`
	// Payer 支付者信息
	Payer any `json:"payer"`
	// TransactionID 微信支付系统生成的订单号
	TransactionID string `json:"transaction_id"`
	// BankType 银行类型
	BankType string `json:"bank_type"`
	// Attach 附加数据
	Attach string `json:"attach"`
	// SuccessTime 支付完成时间，遵循rfc3339标准格式
	SuccessTime string `json:"success_time"`
	// Amount 金额信息
	Amount any `json:"amount"`
	// SceneInfo 商户端设备号等信息
	SceneInfo any `json:"scene_info"`
	// PromotionDetail 优惠详情信息
	PromotionDetail any `json:"promotion_detail"`
}

// GetRefundDetail 查询退款
// 提交退款申请后，通过调用该接口查询退款状态。
//
// see https://developer.work.weixin.qq.com/document/path/97352
func (c *Client) GetRefundDetail(r *GetRefundDetailRequest, opts ...any) (out GetRefundDetailResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/miniapppay/get_refund_detail",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetRefundDetailRequest is request of Client.GetRefundDetail
type GetRefundDetailRequest struct {
	// Mchid 企业微信分配的商户号
	Mchid string `json:"mchid" validate:"required"`
	// OutRefundNo 商户系统内部的退款单号，商户系统内部唯一
	OutRefundNo string `json:"out_refund_no" validate:"required"`
}

// GetRefundDetailResponse is response of Client.GetRefundDetail
type GetRefundDetailResponse struct {
	// RefundID 微信支付退款订单号
	RefundID string `json:"refund_id"`
	// OutRefundNo 商户系统内部的退款单号
	OutRefundNo string `json:"out_refund_no"`
	// TransactionID 微信支付交易订单号
	TransactionID string `json:"transaction_id"`
	// OutTradeNo 返回的原交易订单号
	OutTradeNo string `json:"out_trade_no"`
	// Channel 退款渠道：ORIGINAL（原路退款）、BALANCE（退回到余额）等
	Channel string `json:"channel"`
	// UserReceivedAccount 取当前退款单的退款入账方
	UserReceivedAccount string `json:"user_received_account"`
	// SuccessTime 退款成功时间，RFC3339格式
	SuccessTime string `json:"success_time"`
	// CreateTime 退款受理时间，RFC3339格式
	CreateTime string `json:"create_time"`
	// Status 退款状态：SUCCESS、CLOSE、PROCESSING、ABNORMAL
	Status string `json:"status"`
	// Amount 订单退款金额信息
	Amount any `json:"amount"`
	// PromotionDetail 优惠退款信息
	PromotionDetail []any `json:"promotion_detail"`
}

// GetMiniAppPaySign 获取支付签名
// 商户系统获取完预支付交易会话标识后，在小程序端调用wxsdk前，需要调用本接口获取到必要的签名字段，再调起微信支付。
//
// see https://developer.work.weixin.qq.com/document/path/98130
func (c *Client) GetMiniAppPaySign(r *GetMiniAppPaySignRequest, opts ...any) (out GetMiniAppPaySignResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/miniapppay/get_sign",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetMiniAppPaySignRequest is request of Client.GetMiniAppPaySign
type GetMiniAppPaySignRequest struct {
	// AppID 二级商户申请的公众号或移动应用appid
	AppID string `json:"appid" validate:"required"`
	// PrepayID 小程序下单接口返回的prepay_id参数值，仅支持下单两小时内的prepay_id
	PrepayID string `json:"prepay_id" validate:"required"`
	// SignType 签名类型，默认为RSA，仅支持RSA
	SignType string `json:"sign_type"`
	// Nonce 随机字符串，不长于32位，内容仅支持数字、大小写字母
	Nonce string `json:"nonce" validate:"required"`
	// Timestamp 当前的秒级时间戳
	Timestamp int `json:"timestamp" validate:"required"`
}

// GetMiniAppPaySignResponse is response of Client.GetMiniAppPaySign
type GetMiniAppPaySignResponse struct {
	// PaySign 签名，使用字段appid、timestamp、nonce、prepay_id计算得出的签名值
	PaySign string `json:"pay_sign"`
}

// Refund 申请退款
// 当交易发生之后一段时间内，由于买家或者卖家的原因需要退款时，卖家可以通过退款接口将支付款退还给买家。
//
// see https://developer.work.weixin.qq.com/document/path/97333
func (c *Client) Refund(r *RefundRequest, opts ...any) (out RefundResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/miniapppay/refund",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// RefundRequest is request of Client.Refund
type RefundRequest struct {
	// Mchid 企业微信分配商户号
	Mchid string `json:"mchid" validate:"required"`
	// AppID 小程序appid
	AppID string `json:"appid" validate:"required"`
	// OutTradeNo 原支付交易对应的商户订单号
	OutTradeNo string `json:"out_trade_no" validate:"required"`
	// OutRefundNo 商户系统内部的退款单号，唯一
	OutRefundNo string `json:"out_refund_no" validate:"required"`
	// Reason 退款原因
	Reason string `json:"reason"`
	// FundsAccount 资金账户，如AVAILABLE
	FundsAccount string `json:"funds_account"`
	// Amount 订单金额信息
	Amount any `json:"amount" validate:"required"`
}

// RefundResponse is response of Client.Refund
type RefundResponse struct {
	// OutRefundNo 商户系统内部的退款单号
	OutRefundNo string `json:"out_refund_no"`
	// Amount 订单金额信息
	Amount any `json:"amount"`
	// PromotionDetail 优惠退款功能信息
	PromotionDetail []any `json:"promotion_detail"`
}

