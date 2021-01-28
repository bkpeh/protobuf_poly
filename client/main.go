package main

import (
	"context"
	"fmt"

	messages "github.com/bkpeh/protobuf_poly/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

type msg1 struct {
	evtname string
}

func (m msg1) ProtoReflect() protoreflect.Message {
	var a protoreflect.Message
	return a
}

type msg2 struct {
	evtname string
	evtid   int
}

func main() {

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Fail to dial.", err)
	}

	defer conn.Close()

	newmsg1 := msg1{
		evtname: "E1",
	}

	//any, _ := anypb.New(newmsg1)
	fmt.Println("xxxxx")
	//var anyarr []*anypb.Any
	//msgdata, err := proto.Marshal(newmsg1)
	any, err := anypb.New(newmsg1)
	if err != nil {
		fmt.Println("ANYPB:", err)
	}

	anyarr := []*anypb.Any{any}
	newevt := messages.Event{
		Name:    "EVENT",
		Details: anyarr,
	}

	client := messages.NewGetSystemEventsClient(conn)
	respond, err := client.GetEvent(context.Background(), &newevt)

	fmt.Println("Respond.", respond.Id)
}
