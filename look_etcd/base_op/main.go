package main

import (
	"context"
	"fmt"
	etcd3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	cli, err := etcd3.New(etcd3.Config{
		Endpoints:   []string{"192.168.1.77:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	_, err = cli.Put(context.Background(), "test-key", "test-value")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	rsp, err := cli.Get(context.Background(), "test-", etcd3.WithPrefix())
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	if rsp.Count > 0 {
		for _, kv := range rsp.Kvs {
			fmt.Printf("key: %s, value: %s\n", kv.Key, kv.Value)
		}
	}

	//_, err = cli.Delete(context.Background(), "test-key")
	//if err != nil {
	//	fmt.Printf("err: %v\n", err)
	//	return
	//}
	//
	//grantRsp, err := cli.Grant(context.Background(), 10)
	//if err != nil {
	//	fmt.Printf("err: %v\n", err)
	//	return
	//}

	//fmt.Printf("lease id: %x, info: %v\n", grantRsp.ID, grantRsp)
	//
	//_, err = cli.Revoke(context.Background(), grantRsp.ID)
	//if err != nil {
	//	fmt.Printf("err: %v\n", err)
	//	return
	//}
	//
	//keepChan, err := cli.KeepAlive(context.Background(), grantRsp.ID)
	//if err != nil {
	//	fmt.Printf("err: %v\n", err)
	//	return
	//}
	//
	//for v := range keepChan {
	//	fmt.Printf("id: %d, ttl: %d\n", v.ID, v.TTL)
	//}

	watchChan := cli.Watch(context.Background(), "test-key")
	go func() {
		for rsp := range watchChan {
			for _, event := range rsp.Events {
				if event.Type == etcd3.EventTypePut {
					fmt.Printf("kv put, %s\n", event.Kv.Key)
				} else if event.Type == etcd3.EventTypeDelete {
					fmt.Printf("kv delete, %s\n", event.Kv.Key)
				}
			}
		}

		fmt.Println("退出监控")
	}()

	go func() {
		timer := time.NewTimer(30 * time.Second)
		select {
		case <-timer.C:
			cli.Watcher.Close()
		}
	}()

	select {}

	return
}
