package sand

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const CheckMemberStatusUrl = "https://cap.sandpay.com.cn/v4/elecaccount/ceas.elec.member.status.query"

// CheckMemberStatus 查询会员状态
func CheckMemberStatus(userNo string) {
	reqBody, err := ConstructQueryRequestParams(userNo)
	if err != nil {
		panic(err)
	}
	jsonStr, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}
	client := http.DefaultClient
	buff := bytes.NewBuffer(jsonStr)
	req, err := http.NewRequest("POST", CheckMemberStatusUrl, buff)
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
