package main

import "ceas-go-demo/sand"

// $env:GOPROXY = "https://proxy.golang.com.cn,direct"

func main() {
	sand.CancelAccount("5789984")
}
