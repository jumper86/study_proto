package main

import (
	"fmt"
	"google.golang.org/grpc"
	"person/client"
)

func main() {

	addr := fmt.Sprintf("%s:%d", "127.0.0.1", 16666)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}
	defer conn.Close()

	//只需要建立一个连接，从该连接上获取不同client，在各自类型client上获得各自流进行通信
	client.SimpleClient(conn)
	client.ClientStreamClient(conn)
	client.ServerStreamClient(conn)
	client.BidStreamClient(conn)
}
