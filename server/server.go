package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/drgarcia1986/grpc-hw/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Say(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	return &pb.Response{Msg: fmt.Sprintf("Hello %s", in.Name)}, nil
}

func main() {
	port := ":50051"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloWorldServer(s, &server{})

	s.Serve(listener)
}
