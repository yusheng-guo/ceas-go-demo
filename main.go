package main

import (
	"bufio"
	"ceas-go-demo/sand"
	"fmt"
	"log"
	"os"
	"strings"
)

// $env:GOPROXY = "https://proxy.golang.com.cn,direct"

//const (
//	PrivateKeyFilepath = "./cert/"
//	PublicKeyFilepath  = "./cert/"
//)

func main() {
	var err error
	var userNo, captcha string
	var reader = bufio.NewReader(os.Stdin)
	for {
		// 1.查询会员状态
		fmt.Print("请求如需要销户的用户ID: ")
		userNo, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		userNo = strings.TrimSpace(userNo)
		fmt.Printf("ID为[%s]的用户正在进行销户操作...", userNo)
		err = sand.CheckMemberStatus(userNo)
		if err != nil {
			log.Println(err)
			continue
		}
		// 2.请求销户
		customerOrderNo, err := sand.CancelAccount(userNo)
		if err != nil {
			log.Println(err)
			continue
		}

		// 3.确认销户
		fmt.Print("请输入收到的验证码: ")
		captcha, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		captcha = strings.TrimSpace(captcha)
		fmt.Printf("ID为[%s]用户输入的验证码为[%s],验证码长度为[%d],正在进行销户确认...", userNo, captcha, len(captcha))

		err = sand.ConfirmCancelAccount(customerOrderNo, userNo, captcha)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
