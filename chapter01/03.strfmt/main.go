package main

import "fmt"

func main() {
	fmt.Println("中文， abc, 1234, ❤😍💕")
	fmt.Printf("我的名字是%s\n", "小强")        //%s 字符串
	fmt.Printf("我的名字是%q，%q\n", 78, "小强") //%q 带引号，且是字节类型的
	fmt.Printf("我的名字是%x，%x\n", "小强", 17) //十六进制,且英文小写
	fmt.Printf("我的名字是%X，%X\n", "小强", 17) //十六进制，英文大写
	//更多的去看占位符
}
