package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"person/code/person"
)

func main() {
	per := &person.Person{
		Name:  "wang",
		Id:    1,
		Email: "wang@163.com",
		Phone: []*person.Person_PhoneNumber{
			&person.Person_PhoneNumber{
				Type:   person.Person_Home,
				Number: "111111",
			},
			&person.Person_PhoneNumber{
				Type:   person.Person_Mobile,
				Number: "666666",
			},
		},
	}

	fmt.Printf("person info: %v\n", per.String())

	code, err := proto.Marshal(per)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("person code: %v\n", code)

	var per2 person.Person
	err = proto.Unmarshal(code, &per2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("person2 info: %v\n", per2.String())
}
