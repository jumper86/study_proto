package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"person/code/person"
)

func ClientStreamClient(conn grpc.ClientConnInterface) error {

	client := person.NewClientStreamClient(conn)

	values := map[string]string{
		"client-stream-key": "client-stream-value",
	}
	md := metadata.New(values)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	stream, err := client.GetPersonClientStream(ctx)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	ids := []*person.Ids{
		&person.Ids{Ids: []uint64{1, 2, 3}},
		&person.Ids{Ids: []uint64{4, 5, 6}},
		&person.Ids{Ids: []uint64{7, 8, 9}},
	}
	for _, info := range ids {
		stream.Send(info)
	}
	p, err := stream.CloseAndRecv()
	fmt.Printf("get person in ClientStreamClient: %v\n", p)
	return nil
}
