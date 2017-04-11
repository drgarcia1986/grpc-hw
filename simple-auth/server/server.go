package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"golang.org/x/net/context"

	pb "github.com/drgarcia1986/grpc-hw/simple-auth/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var ErrPermissionDenied = grpc.Errorf(codes.PermissionDenied, "Permission Denied")

type server struct{}

func (*server) Login(ctx context.Context, request *pb.LoginRequest) (*pb.LoginResponse, error) {
	if request.User != "admin" || request.Password != "ananas2000" {
		return nil, ErrPermissionDenied
	}
	return &pb.LoginResponse{Token: "FOO"}, nil
}

func (*server) Restricted(ctx context.Context, _ *pb.Empty) (*pb.RestrictedResponse, error) {
	user := ctx.Value("user").(string)
	return &pb.RestrictedResponse{Message: fmt.Sprintf("Hi %s", user)}, nil
}

func authorize(ctx context.Context) error {
	if md, ok := metadata.FromContext(ctx); ok {
		if len(md["token"]) == 1 && md["token"][0] == "FOO" {
			return nil
		}
	}

	return ErrPermissionDenied
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if strings.HasSuffix(info.FullMethod, "Login") {
		return handler(ctx, req)
	}

	if err := authorize(ctx); err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, "user", "admin")
	return handler(ctx, req)
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(unaryInterceptor))
	pb.RegisterSimpleAuthServer(s, &server{})

	log.Println("gRPC server: ", s.Serve(listener))
}
