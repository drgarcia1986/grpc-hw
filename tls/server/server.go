package main

import (
	"fmt"
	"log"
	"net"

	"golang.org/x/net/context"

	pb "github.com/drgarcia1986/grpc-hw/tls/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type server struct{}

func (*server) Say(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	return &pb.Response{Msg: fmt.Sprintf("Hello %s", in.Name)}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	creds, err := credentials.NewServerTLSFromFile("../cert.crt", "../cert.key")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterHelloWorldServer(s, &server{})

	log.Println("gRPC server: ", s.Serve(listener))
}
