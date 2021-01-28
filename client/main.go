package main

import (
	"context"
	"fmt"

	messages "github.com/bkpeh/protobuf_poly/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

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

func main() {

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		fmt.Println("Fail to dial.", err)
	}

	defer conn.Close()

	anym1, _ := anypb.New(msg1)
	anym2, _ := anypb.New(msg2)

	anyarr := []*anypb.Any{anym1, anym2}

	newevt := messages.Event{
		Name:    "EVENT",
		Details: anyarr,
	}

	client := messages.NewGetSystemEventsClient(conn)
	respond, err := client.GetEvent(context.Background(), &newevt)

	fmt.Println("Respond:", respond.Id)
}
