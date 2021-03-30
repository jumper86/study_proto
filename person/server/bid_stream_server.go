package server

import (
	"fmt"
	"google.golang.org/grpc/metadata"
	"io"
	"math/rand"
	"person/code/person"
)

type BidStreamServerImpl struct {
	person.UnimplementedBidStreamServer
}

func NewBidStreamServer() *BidStreamServerImpl {
	return &BidStreamServerImpl{}
}

func (this *BidStreamServerImpl) GetPersonBidStream(stream person.BidStream_GetPersonBidStreamServer) error {
	ctx := stream.Context()
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		value := md.Get("bid-stream-key")
		fmt.Printf("get context from bid stream %v\n", value)
	}

	for count := 1; count < 10; count++ {
		ids, err := stream.Recv()
		if err == io.EOF {
			return nil
		} else if err != nil {
			return err
		}
		fmt.Printf("get ids: %v\n", ids)
		stream.Send(&person.Person{
			Id:   rand.Uint64(),
			Name: "随机",
			Age:  999,
		})
	}

	return nil
}
