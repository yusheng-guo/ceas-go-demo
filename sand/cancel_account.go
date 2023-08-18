package sand

import "errors"

const ManageMemberStatusUrl = "https://cap.sandpay.com.cn/v4/elecaccount/ceas.elec.account.member.status.modify"

// CancelAccount 杉德云账户注销 销户
func CancelAccount(userNo string) (string, error) {
	reqBody, err := ConstructCancelAccountRequestParams(userNo)
	if err != nil {
		return "", err
	}
	respBody, err := SandHttp(ManageMemberStatusUrl, reqBody)
	if err != nil {
		return "", err
	}
	if respBody.Response.ResponseCode != "00000" {
		return "", errors.New(respBody.Response.ResponseDesc)
	}
	return reqBody.CustomerOrderNo, nil
}
