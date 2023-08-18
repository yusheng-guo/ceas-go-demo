package sand

import "errors"

const CheckMemberStatusUrl = "https://cap.sandpay.com.cn/v4/elecaccount/ceas.elec.member.status.query"

// CheckMemberStatus 查询会员状态
func CheckMemberStatus(userNo string) error {
	reqBody, err := ConstructQueryRequestParams(userNo)
	if err != nil {
		return err
	}
	respBody, err := SandHttp(CheckMemberStatusUrl, reqBody)
	if err != nil {
		return err
	}
	if respBody.Response.ResponseCode != "00000" {
		return errors.New(respBody.Response.ResponseDesc)
	}
	return nil
}
