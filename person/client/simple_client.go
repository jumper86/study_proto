package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"person/code/person"
)

func SimpleClient(conn grpc.ClientConnInterface) error {

	client := person.NewSimpleClient(conn)

	values := map[string]string{
		"simple-key": "simple-value",
	}
	md := metadata.New(values)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	var ids person.Ids
	ids.Ids = []uint64{1, 2, 3, 4, 5, 6}
	person, err := client.GetPerson(ctx, &ids)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return err
	}
	fmt.Printf("get person in SimpleClient: %v\n", person)

	return nil
}
