package sand

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const ManageMemberStatusUrl = "https://cap.sandpay.com.cn/v4/elecaccount/ceas.elec.account.member.status.modify"

// CancelAccount 杉德云账户注销 销户
func CancelAccount(userNo string) {
	reqBody, err := ConstructRequestParams(userNo)
	if err != nil {
		panic(err)
	}
	jsonStr, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}
	fmt.Println("请求:", string(jsonStr))
	client := http.DefaultClient
	buff := bytes.NewBuffer(jsonStr)
	req, err := http.NewRequest("POST", ManageMemberStatusUrl, buff)
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
}
