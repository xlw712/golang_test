package main

import (
	"fmt"
	qr "github.com/skip2/go-qrcode"
)

func main() {
	// 将内容生产一个二维码,并且保存成图片
	qr.WriteFile("陈辰大宝贝",qr.Medium,256,"/Users/mac/levin/1.png")
	// 生产一个二维码字节流
	if b1,error:=qr.Encode("http://www.baidu.com",qr.Medium,256); error ==nil{
		fmt.Print(b1)
	}
}
