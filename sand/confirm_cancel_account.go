package sand

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const ConfirmCancelAccountUrl = "https://cap.sandpay.com.cn/v4/elecaccount/ceas.elec.account.member.modify.confirm"

// ConfirmCancelAccount 确认销户 captcha 验证码 oriCustomerOrderNo 原订单号
func ConfirmCancelAccount(captcha, userNo, oriCustomerOrderNo string) {
	reqBody, err := ConstructConfirmRequestParams(captcha, userNo, oriCustomerOrderNo)
	if err != nil {
		panic(err)
	}
	jsonStr, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient
	buff := bytes.NewBuffer(jsonStr)
	req, err := http.NewRequest("POST", ConfirmCancelAccountUrl, buff)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			panic(err)
		}
	}()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("响应:", string(respBody))
	return
}
