package sand

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Resp 响应
type Resp struct {
	Data       string `json:"data"`
	Sign       string `json:"sign"`
	EncryptKey string `json:"encryptKey"`
	Response   struct {
		ResponseDesc    string `json:"responseDesc"`
		ResponseTime    string `json:"responseTime"`
		Mid             string `json:"mid"`
		SandSerialNo    string `json:"sandSerialNo"`
		ResponseStatus  string `json:"responseStatus"`
		CustomerOrderNo string `json:"customerOrderNo"`
		Version         string `json:"version"`
		ResponseCode    string `json:"responseCode"`
	} `json:"response"`
	SignType    string `json:"signType"`
	EncryptType string `json:"encryptType"`
}

// SandHttp 发送杉德请求
func SandHttp(url string, reqBody *Req) (*Resp, error) {
	jsonStr, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal req body, err: %w", err)
	}
	client := http.DefaultClient
	buff := bytes.NewBuffer(jsonStr)
	req, err := http.NewRequest("POST", url, buff)
	if err != nil {
		return nil, fmt.Errorf("post request %s, err: %w", url, err)
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("do, err: %w", err)
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {

		}
	}()
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body, err: %w", err)
	}
	var respData = new(Resp)
	err = json.Unmarshal(respBody, respData)
	if err != nil {
		return nil, fmt.Errorf("unmarshal response, err: %w", err)
	}
	return respData, nil
}
