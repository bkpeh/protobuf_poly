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
	fmt.Println("GetEvent")
	fmt.Println(e.GetDetails())
	return &messages.Pid{Id: 1}, nil
}

func main() {
	svr, err := net.Listen("tcp", ":9000")

	if err != nil {
		fmt.Println("Error starting Server.", err)
	}

	s := grpc.NewServer()

	messages.RegisterGetSystemEventsServer(s, &server{})

	if err = s.Serve(svr); err != nil {
		fmt.Println("Error serving.", err)
	}
}
