package sand

import (
	"errors"
)

const ConfirmCancelAccountUrl = "https://cap.sandpay.com.cn/v4/elecaccount/ceas.elec.account.member.modify.confirm"

// ConfirmCancelAccount 确认销户 captcha 验证码 oriCustomerOrderNo 原订单号
func ConfirmCancelAccount(captcha, userNo, oriCustomerOrderNo string) error {
	reqBody, err := ConstructConfirmRequestParams(captcha, userNo, oriCustomerOrderNo)
	if err != nil {
		panic(err)
	}
	respBody, err := SandHttp(ConfirmCancelAccountUrl, reqBody)
	if err != nil {
		return err
	}
	if respBody.Response.ResponseCode != "00000" {
		return errors.New(respBody.Response.ResponseDesc)
	}
	return nil
}
