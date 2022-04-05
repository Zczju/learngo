package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"learn.go/chapter13/02.4.grpc_fatr_shizhan/apis"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure()) //这里WithInsecure是必需的  grpc是基于https的，必须是以https访问的 ，
	// 如果强制不用https，就必须要WithInsecure
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := apis.NewRankServiceClient(conn)
	ret, err := c.Register(context.TODO(), &apis.PersonalInformation{
		Name:   "Tom",
		Sex:    "男",
		Tall:   1.8,
		Weight: 75.0,
		Age:    30,
	})
	if err != nil {
		log.Fatal("注册失败：", err)
	}
	log.Println("注册成功", ret)

	{ // 榜单
		top, err := c.GetTop(context.TODO(), &apis.Null{})
		if err != nil {
			log.Fatal("获取榜单失败：", err)
		}
		log.Println("榜单：", top.String())
	}

	log.Println("开始批量注册")
	regCli, err := c.RegisterPersons(context.TODO())
	if err != nil {
		log.Fatal("获取批量注册客户端失败：", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{
		Name:   fmt.Sprintf("tom-%d", time.Now().Nanosecond()),
		Sex:    "男",
		Tall:   1.8,
		Weight: 71,
		Age:    28,
	}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{Name: fmt.Sprintf("tom-%d", time.Now().Nanosecond())}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{Name: fmt.Sprintf("tom-%d", time.Now().Nanosecond())}); err != nil {
		log.Fatal("注册时失败：", err)
	}
	if err := regCli.Send(&apis.PersonalInformation{Name: fmt.Sprintf("tom-%d", time.Now().Nanosecond())}); err != nil {
		log.Fatal("注册时失败：", err)
	}

	resp, err := regCli.CloseAndRecv()
	if err != nil {
		log.Fatal("无法接收结果：", err)
	}
	log.Println("批量注册结果：", resp.String())

	{ // 榜单
		top, err := c.GetTop(context.TODO(), &apis.Null{})
		if err != nil {
			log.Fatal("获取榜单失败：", err)
		}
		log.Println("榜单：", top.String())
	}
}
