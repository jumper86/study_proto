package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"person/code/person"
)

func BidStreamClient(conn grpc.ClientConnInterface) error {

	client := person.NewBidStreamClient(conn)

	values := map[string]string{
		"bid-stream-key": "bid-stream-value",
	}
	md := metadata.New(values)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	stream, err := client.GetPersonBidStream(ctx)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	err = stream.Send(&person.Ids{Ids: []uint64{1, 2, 3}})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	p, err := stream.Recv()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}
	fmt.Printf("get person in BidStreamClient : %v\n", p)

	stream.CloseSend()

	return nil
}
