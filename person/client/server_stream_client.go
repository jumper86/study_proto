package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"io"
	"person/code/person"
)

func ServerStreamClient(conn grpc.ClientConnInterface) error {

	client := person.NewServerStreamClient(conn)

	values := map[string]string{
		"server-stream-key": "server-stream-value",
	}
	md := metadata.New(values)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	stream, err := client.GetPersonServerStream(ctx, &person.Ids{Ids: []uint64{1, 2, 3}})
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	for {
		p, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		fmt.Printf("get person in ServerStreamClient : %v\n", p)
	}

	return nil
}
