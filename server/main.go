package main

import (
	"context"
	"fmt"
	"net"

	messages "github.com/bkpeh/protobuf_poly/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/structpb"
)

type server struct {
	messages.UnimplementedGetSystemEventsServer
}

func (s server) GetEvent(ctx context.Context, e *messages.Event) (*messages.Pid, error) {
	fmt.Println("Event Name:", e.Name)

	for _, v := range e.Details {
		fmt.Println("==============")
		m := &structpb.Struct{}
		v.UnmarshalTo(m)

		for ii, vv := range m.Fields {
			fmt.Println(ii, ":", vv)
		}
	}

	return &messages.Pid{Id: 1}, nil
}

func main() {
	svr, err := net.Listen("tcp", ":9000")

	if err != nil {
		fmt.Println("Error starting Server:", err)
	}

	s := grpc.NewServer()

	messages.RegisterGetSystemEventsServer(s, &server{})

	if err = s.Serve(svr); err != nil {
		fmt.Println("Error serving:", err)
	}
}
