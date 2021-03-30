package server

import (
	"fmt"
	"google.golang.org/grpc/metadata"
	"math/rand"
	"person/code/person"
)

type ServerStreamServerImpl struct {
	person.UnimplementedServerStreamServer
}

func NewServerStreamServer() *ServerStreamServerImpl {
	return &ServerStreamServerImpl{}
}

func (this *ServerStreamServerImpl) GetPersonServerStream(ids *person.Ids, stream person.ServerStream_GetPersonServerStreamServer) error {
	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		value := md.Get("server-stream-key")
		fmt.Printf("get context from server stream %v\n", value)
	}

	fmt.Printf("get ids: %v\n", ids)
	for count := 0; count < 1; count++ {
		stream.Send(&person.Person{
			Id:   rand.Uint64(),
			Name: "随机",
			Age:  99,
		})
	}

	return nil
}
