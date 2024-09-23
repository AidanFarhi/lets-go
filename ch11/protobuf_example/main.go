package main

import (
	"fmt"
	"protoexample/data"

	"google.golang.org/protobuf/proto"
)

//go:generate protoc -I=. --go_out=. --go_opt=module=protoexample --go_opt=Mperson.proto=protoexample/data person.proto
func main() {

	p := &data.Person{
		Name:  "Joe",
		Id:    1,
		Email: "joe@gmail.com",
	}
	fmt.Println(p)

	protoBytes, _ := proto.Marshal(p)
	fmt.Println(protoBytes)
	var p2 data.Person
	proto.Unmarshal(protoBytes, &p2)
	fmt.Println(&p2)
}
