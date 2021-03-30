package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"person/code/person"
)

type SimpleServerImpl struct {
	person.UnimplementedSimpleServer
}

func NewSimpleServer() *SimpleServerImpl {
	return &SimpleServerImpl{}
}

func (this *SimpleServerImpl) GetPerson(ctx context.Context, ids *person.Ids) (*person.Person, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if value, ok := md["simple-key"]; ok {
			fmt.Printf("get context from simple stream: %v\n", value)
		}
	}

	fmt.Printf("get ids: %v\n", ids.Ids)
	return &person.Person{
		Id:   999,
		Name: "yeah",
		Age:  999,
	}, nil
}
