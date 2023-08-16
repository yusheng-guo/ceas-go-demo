package sand

import (
	"bytes"
	"ceas-go-demo/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	MID                   = "6888802117311"       // 商户编号
	SignType              = "SHA1WithRSA"         // 签名方式
	EncryptType           = "AES"                 // 加密方式
	Version               = "1.0.0"               // API接口版本
	Layout                = "2006-01-02 15:04:05" // 时间戳格式
	ManageMemberStatusUrl = "https://cap.sandpay.com.cn/v4/elecaccount/ceas.elec.account.member.status.modify"
)

// CancelAccountReq 销户请求
//type CancelAccountReq struct {
//	Mid             string `json:"mid"`             // 杉德支付分配给接入商户的商户编号
//	Sign            string `json:"sign"`            // (SHA1WithRSA)，签名结果采用base64编码
//	Timestamp       string `json:"timestamp"`       // 发送请求的时间
//	CustomerOrderNo string `json:"customerOrderNo"` // 商户订单号 商户号下每次请求的唯一流水号
//	SignType        string `json:"signType"`        // 签名方式 SHA1WithRSA
//	EncryptType     string `json:"encryptType"`     // 加密方式 AES
//	EncryptKey      string `json:"encryptKey"`      // 加密key 	使用杉德公钥对16位随机数进行RSA加密(RSA/ECB/PKCS1Padding)，加密结果采用base64编码
//	Data            string `json:"data"`            // 使用16位随机数对明文参数进行AES加密(AES/ECB/PKCS5Padding) 加密结果采用base64编码
//	Version         string `json:"version"`         // 接口版本
//}

// Data 请求数据
type Data struct {
	Mid             string `json:"mid"`             // 商户号
	CustomerOrderNo string `json:"customerOrderNo"` // 商户订单号
	BizUserNo       string `json:"bizUserNo"`       // 会员编号 需要注销的会员
	BizType         string `json:"bizType"`         // 操作类型 "CLOSE" 销户
	NotifyUrl       string `json:"notifyUrl"`       // 异步通知地址
}

// Req 请求报文 公共报文头
type Req struct {
	Mid             string `json:"mid"`             // 商户号
	Sign            string `json:"sign"`            // 签名 使用商户私钥对data参数进行RSA签名(SHA1WithRSA)，签名结果采用base64编码
	Timestamp       string `json:"timestamp"`       // 格式 时间戳 2021-02-21 20:28:10
	Version         string `json:"version"`         // 版本号 1.0.0
	CustomerOrderNo string `json:"customerOrderNo"` // 商户订单号
	SignType        string `json:"signType"`        // 签名方式 SHA1WithRSA
	EncryptType     string `json:"encryptType"`     // 加密方式 AES
	EncryptKey      string `json:"encryptKey"`      // 加密 Key 使用杉德公钥对16位随机数进行RSA加密(RSA/ECB/PKCS1Padding) 加密结果采用base64编码
	Data            string `json:"data"`            // 使用16位随机数对明文参数进行AES加密(AES/ECB/PKCS5Padding) 加密结果采用base64编码
}

// ConstructRequestParams 构造请求参数
func ConstructRequestParams(userNo string) (req *Req, err error) {
	orderNo := utils.GenerateOrderNo()
	var data = &Data{
		Mid:             MID,
		CustomerOrderNo: orderNo,
		BizUserNo:       userNo,
		BizType:         "CLOSE",
		NotifyUrl:       "",
	}
	// 对 data 数据进行加密
	marshal, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	fmt.Println("data: ", string(marshal))
	//var src map[string]any
	//err = json.Unmarshal(marshal, &src)
	//if err != nil {
	//	return nil, err
	//}
	value, key, err := AESEncrypt(marshal) // AES 加密
	if err != nil {
		return nil, err
	}
	fmt.Println("value => ", value)
	fmt.Println("key => ", key)

	// 对数据进行 签名
	sign, err := Sign([]byte(value))
	if err != nil {
		return nil, err
	}
	return &Req{
		Mid:             MID,
		Sign:            sign,
		Timestamp:       time.Now().Format(Layout),
		Version:         Version,
		CustomerOrderNo: orderNo,
		SignType:        SignType,
		EncryptType:     EncryptType,
		EncryptKey:      key,
		Data:            value,
	}, nil
}

// CancelAccount 杉德云账户注销 销户
func CancelAccount(userNo string) {
	fmt.Println(ManageMemberStatusUrl)
	reqBody, err := ConstructRequestParams(userNo)
	if err != nil {
		panic(err)
	}
	jsonStr, err := json.Marshal(reqBody)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonStr))
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
