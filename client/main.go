package main

import (
	"context"
	"fmt"

	messages "github.com/bkpeh/protobuf_poly/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

/*
var msg1 = &structpb.Struct{
	Fields: map[string]*structpb.Value{
		"subevtname": &structpb.Value{
			Kind: &structpb.Value_StringValue{
				StringValue: "E1",
			},
		},
	},
}

var msg2 = &structpb.Struct{
	Fields: map[string]*structpb.Value{
		"subevtname": &structpb.Value{
			Kind: &structpb.Value_StringValue{
				StringValue: "E2",
			},
		},
		"evtid": &structpb.Value{
			Kind: &structpb.Value_StringValue{
				StringValue: "22",
			},
		},
	},
}
*/
type newmsg struct {
	evtname string
}

func (x newmsg) ProtoReflect() protoreflect.Message {
	return
}

func main() {

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Fail to dial.", err)
	}

	defer conn.Close()

	var n newmsg
	anym1, _ := anypb.New(n)

	/*
		anym1, _ := anypb.New(msg1)
		anym2, _ := anypb.New(msg2)
		anym3, _ := anypb.New(msg3)
	*/
	anyarr := []*anypb.Any{anym1}

	newevt := messages.Event{
		Name:    "EVENT",
		Details: anyarr,
	}

	client := messages.NewGetSystemEventsClient(conn)
	respond, err := client.GetEvent(context.Background(), &newevt)

	fmt.Println("Respond:", respond.Id)
}
