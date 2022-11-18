package WeWorkFinanceSDK

import (
	"fmt"
)

type Error struct {
	Code    int    `json:"errcode,omitempty"`
	Message string `json:"errmsg,omitempty"`
	Detail  string `json:"-"`
}

func (err Error) IsRetryable() bool {
	return err.Code >= 10001 && err.Code <= 1003
}

func (err Error) Error() string {
	if err.Detail != "" {
		return fmt.Sprintf("%d: %s %s", err.Code, err.Message, err.Detail)
	}
	return fmt.Sprintf("%d: %s", err.Code, err.Message)
}

func (err Error) AsError() error {
	if err.Code == 0 {
		return nil
	}
	return err
}

func ErrorOfCode(code int, detail string) Error {
	msg := ""
	switch code {
	case 10000:
		msg = "参数错误，请求参数错误"
	case 10001:
		msg = "网络错误，网络请求错误"
	case 10002:
		msg = "数据解析失败"
	case 10003:
		msg = "系统失败"
	case 10004:
		msg = "密钥错误导致加密失败"
	case 10005:
		msg = "fileid错误"
	case 10006:
		msg = "解密失败"
	case 10007:
		msg = "找不到消息加密版本的私钥，需要重新传入私钥对"
	case 10008:
		msg = "解析encrypt_key出错"
	case 10009:
		msg = "ip非法"
	case 10010:
		msg = "数据过期"
	case 10011:
		msg = "证书错误"
	}
	return Error{Code: code, Message: msg, Detail: detail}
}
