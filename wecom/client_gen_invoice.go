package wecom

// Code generated from WeChat Work API documentation. DO NOT EDIT.

import (
	"github.com/wenerme/go-req"
)

// GetInvoiceInfo 查询电子发票
// 报销方在获得用户选择的电子发票标识参数后，可以通过该接口查询电子发票的结构化信息，并获取发票PDF文件。
//
// see https://developer.work.weixin.qq.com/document/path/90103
func (c *Client) GetInvoiceInfo(r *GetInvoiceInfoRequest, opts ...any) (out GetInvoiceInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/card/invoice/reimburse/getinvoiceinfo",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// GetInvoiceInfoRequest is request of Client.GetInvoiceInfo
type GetInvoiceInfoRequest struct {
	// CardID 发票id
	CardID string `json:"card_id" validate:"required"`
	// EncryptCode 加密code
	EncryptCode string `json:"encrypt_code" validate:"required"`
}

// GetInvoiceInfoResponse is response of Client.GetInvoiceInfo
type GetInvoiceInfoResponse struct {
	// CardID 发票id
	CardID string `json:"card_id"`
	// BeginTime 发票的有效期起始时间
	BeginTime int `json:"begin_time"`
	// EndTime 发票的有效期截止时间
	EndTime int `json:"end_time"`
	// OpenID 用户标识
	OpenID string `json:"openid"`
	// Type 发票类型，如广东增值税普通发票
	Type string `json:"type"`
	// Payee 发票的收款方
	Payee string `json:"payee"`
	// Detail 发票详情
	Detail string `json:"detail"`
	// UserInfo 发票的用户信息
	UserInfo any `json:"user_info"`
	// Info 商品信息结构
	Info []any `json:"info"`
}

// BatchGetInvoiceInfo 批量查询电子发票
// 报销方在获得用户选择的电子发票标识参数后，可以通过该接口批量查询电子发票的结构化信息。
//
// see https://developer.work.weixin.qq.com/document/path/90106
func (c *Client) BatchGetInvoiceInfo(r *BatchGetInvoiceInfoRequest, opts ...any) (out BatchGetInvoiceInfoResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/card/invoice/reimburse/getinvoiceinfobatch",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchGetInvoiceInfoRequest is request of Client.BatchGetInvoiceInfo
type BatchGetInvoiceInfoRequest struct {
	// ItemList 发票列表
	ItemList []map[string]any `json:"item_list" validate:"required"`
	// CardID 发票id
	CardID string `json:"card_id" validate:"required"`
	// EncryptCode 加密code
	EncryptCode string `json:"encrypt_code" validate:"required"`
}

// BatchGetInvoiceInfoResponse is response of Client.BatchGetInvoiceInfo
type BatchGetInvoiceInfoResponse struct {
	// ItemList 发票列表信息
	ItemList []map[string]any `json:"item_list"`
}

// UpdateInvoiceStatus 更新发票状态
// 报销企业和报销服务商可以通过该接口对某一张发票进行锁定、解锁和报销操作。
//
// see https://developer.work.weixin.qq.com/document/path/90104
func (c *Client) UpdateInvoiceStatus(r *UpdateInvoiceStatusRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/card/invoice/reimburse/updateinvoicestatus",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// UpdateInvoiceStatusRequest is request of Client.UpdateInvoiceStatus
type UpdateInvoiceStatusRequest struct {
	// CardID 发票id
	CardID string `json:"card_id" validate:"required"`
	// EncryptCode 加密code
	EncryptCode string `json:"encrypt_code" validate:"required"`
	// ReimburseStatus 发报销状态 INVOICE_REIMBURSE_INIT：发票初始状态，未锁定；INVOICE_REIMBURSE_LOCK：发票已锁定，无法重复提交报销;INVOICE_REIMBURSE_C...
	ReimburseStatus string `json:"reimburse_status" validate:"required"`
}

// BatchUpdateInvoiceStatus 批量更新发票状态
// 发票平台可以通过该接口对某个成员的一批发票进行锁定、解锁和报销操作。注意，报销状态为不可逆状态，请开发者慎重调用。
//
// see https://developer.work.weixin.qq.com/document/path/90105
func (c *Client) BatchUpdateInvoiceStatus(r *BatchUpdateInvoiceStatusRequest, opts ...any) (out GenericResponse, err error) {
	err = c.Request.With(req.Request{
		Method:  "POST",
		URL:     "/cgi-bin/card/invoice/reimburse/updatestatusbatch",
		Body:   r,
		Options: opts,
	}).Fetch(&out)
	return
}

// BatchUpdateInvoiceStatusRequest is request of Client.BatchUpdateInvoiceStatus
type BatchUpdateInvoiceStatusRequest struct {
	// OpenID 用户openid，可用userid与openid互换接口获取
	OpenID string `json:"openid" validate:"required"`
	// ReimburseStatus 发票报销状态 INVOICE_REIMBURSE_INIT：发票初始状态，未锁定；INVOICE_REIMBURSE_LOCK：发票已锁定，无法重复提交报销;INVOICE_REIMBURSE_...
	ReimburseStatus string `json:"reimburse_status" validate:"required"`
	// InvoiceList 发票列表，必须全部属于同一个openid
	InvoiceList []map[string]any `json:"invoice_list" validate:"required"`
}

