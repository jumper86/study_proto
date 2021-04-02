package main

import (
	"fmt"
	etcd3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"look_etcd"
	"person/client"
	"time"
)

const (
	defaultScheme string = "testetcd"
)

func main() {
	ec, err := etcd3.New(etcd3.Config{
		Endpoints: []string{"192.168.1.77:2379"},
	})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	rb := look_etcd.NewResolverBuilder(defaultScheme, ec)
	resolver.Register(rb)

	cc, err := grpc.Dial("testetcd://testes/testservice",
		grpc.WithInsecure(),
		grpc.WithBlock(), grpc.WithTimeout(3*time.Second),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)

	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}
	defer cc.Close()

	for {

		//todo: 开始之前就要获取所有node的连接，看看room NewEnviron srv.watch()
		client.ClientStreamClient(cc)
		time.Sleep(1 * time.Second)

	}

	select {}
}
