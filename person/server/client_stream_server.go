package server

import (
	"fmt"
	"google.golang.org/grpc/metadata"
	"io"
	"math/rand"
	"person/code/person"
)

type ClientStreamServerImpl struct {
	person.UnimplementedClientStreamServer
}

func NewClientStreamServer() *ClientStreamServerImpl {
	return &ClientStreamServerImpl{}
}

func (this *ClientStreamServerImpl) GetPersonClientStream(stream person.ClientStream_GetPersonClientStreamServer) error {

	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		value := md.Get("client-stream-key")
		fmt.Printf("get context from client stream %v\n", value)
	}

	for {
		ids, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		fmt.Printf("get ids: %v\n", ids.Ids)
	}
	err := stream.SendAndClose(&person.Person{
		Id:   rand.Uint64(),
		Name: "随机",
		Age:  999,
	})

	if err != nil {
		fmt.Printf("send and close failed, err: %v\n", err)
		return err
	}

	return nil
}
