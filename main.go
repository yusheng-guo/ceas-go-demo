package main

import "ceas-go-demo/sand"

// $env:GOPROXY = "https://proxy.golang.com.cn,direct"

//const (
//	PrivateKeyFilepath = "./cert/"
//	PublicKeyFilepath  = "./cert/"
//)

func main() {
	sand.CancelAccount("5789984")
}
