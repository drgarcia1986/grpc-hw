package main

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"golang.org/x/net/context"

	pb "github.com/drgarcia1986/grpc-hw/grpc-http-same-port/pb"
)

type grpcServer struct{}

func (*grpcServer) Say(ctx context.Context, in *pb.Req) (*pb.Res, error) {
	return &pb.Res{Msg: fmt.Sprintf("Hello %s", in.Name)}, nil
}

func grpcServe(l net.Listener) error {
	s := grpc.NewServer()
	pb.RegisterHelloWorldServer(s, &grpcServer{})

	return s.Serve(l)
}
