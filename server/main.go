package main

import (
	"context"
	"fmt"
	"net"

	messages "github.com/bkpeh/protobuf_poly/proto"
	"google.golang.org/grpc"
)

type server struct {
	messages.UnimplementedGetSystemEventsServer
}

func (s server) GetEvent(ctx context.Context, e *messages.Event) (*messages.Pid, error) {
	fmt.Println("Event Name:", e.Name)

	for _, v := range e.Details {
		fmt.Println("==============")

		if v.MessageIs(&messages.EventMsg1{}) {
			msg := new(messages.EventMsg1)

			if err := v.UnmarshalTo(msg); err != nil {
				fmt.Println(v.MessageName(), err)
			}

			fmt.Println(v.MessageName())
			fmt.Println(msg.GetName())
			fmt.Println(msg.GetId())
		}
		if v.MessageIs(&messages.EventMsg2{}) {
			msg := new(messages.EventMsg2)

			if err := v.UnmarshalTo(msg); err != nil {
				fmt.Println(v.MessageName(), err)
			}

			fmt.Println(v.MessageName())
			fmt.Println(msg.GetName())
			fmt.Println(msg.GetText())
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
