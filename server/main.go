package main

import (
	"context"
	"fmt"
	"net"

	messages "github.com/bkpeh/protobuf_poly/proto"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type server struct {
	messages.UnimplementedGetSystemEventsServer
}

type evtmsg struct {
	evt string
}

func (x evtmsg) ProtoReflect() protoreflect.Message {
	var a protoreflect.Message
	return a
}

func (s server) GetEvent(ctx context.Context, e *messages.Event) (*messages.Pid, error) {
	fmt.Println("Event Name:", e.Name)

	for _, v := range e.Details {
		fmt.Println("==============")
		var m evtmsg

		v.UnmarshalTo(m)
		fmt.Println("EVT:", m.evt)

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
