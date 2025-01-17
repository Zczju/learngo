package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"strconv"
	"time"
)

func main() {
	etcdGetDemo()

}

func etcdGetDemo() {
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"http://localhost:2379"}})
	if err != nil {
		log.Fatal(err)
	}
	resp, err := cli.Get(context.TODO(), "a")
	if err != nil {
		log.Fatal(err)
	}
	for _, kv := range resp.Kvs {
		log.Printf("key: %s, value: %s", kv.Key, kv.Value)
	}
}

func etcdWatchAndReactDemo() {
	cli, err := clientv3.New(clientv3.Config{Endpoints: []string{"http://localhost:2379"}})
	if err != nil {
		log.Fatal(err)
	}

	dataCh := make(chan int)
	go func() {
		watcher := cli.Watch(context.TODO(), "a") //etcd 的watch
		for respData := range watcher {
			evs := respData.Events //evt是事件   事件驱动 模型驱动
			whetherBreak := false  //结束双层循环
			for _, ev := range evs {
				i, err := strconv.Atoi(string(ev.Kv.Value))
				if err != nil {
					fmt.Println("不是数字，结束")
					whetherBreak = true
					break
				}
				dataCh <- i
			}
			if whetherBreak {
				break
			}
		}
	}()

	go func() {
		for i := range dataCh {
			_, err := cli.Put(context.TODO(), "a", fmt.Sprintf("%d", i+1))
			if err != nil {
				fmt.Println("WARNING:更新失败")
			}
		}
	}()

	time.Sleep(2 * time.Second)
}
