package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"person/code/person"
	ps "person/server"
)

func main() {

	lis, err := net.Listen("tcp", "192.168.1.188:26667")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	server := grpc.NewServer()
	person.RegisterSimpleServer(server, ps.NewSimpleServer())
	person.RegisterClientStreamServer(server, ps.NewClientStreamServer())
	person.RegisterServerStreamServer(server, ps.NewServerStreamServer())
	person.RegisterBidStreamServer(server, ps.NewBidStreamServer())
	server.Serve(lis)
}
