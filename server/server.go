package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/drgarcia1986/grpc-hw/pb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	rpcPort  = ":50051"
	httpPort = ":8080"
)

type server struct{}

func (*server) Say(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	return &pb.Response{Msg: fmt.Sprintf("Hello %s", in.Name)}, nil
}

func main() {
	go runRPCServer()
	runHTTPServer()
}

func runHTTPServer() {
	ctx := context.Background()

	gwmux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := pb.RegisterHelloWorldHandlerFromEndpoint(ctx, gwmux, rpcPort, opts); err != nil {
		panic(err)
	}

	fileServer := http.FileServer(http.Dir("swagger-ui"))

	mux := new(http.ServeMux)
	mux.Handle("/docs/", http.StripPrefix("/docs/", fileServer))
	mux.Handle("/", gwmux)

	log.Println("HTTP Server: ", http.ListenAndServe(httpPort, mux))
}

func runRPCServer() {
	listener, err := net.Listen("tcp", rpcPort)
	if err != nil {
		log.Fatal("failed to listen: ", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloWorldServer(s, &server{})

	log.Println("gRPC server: ", s.Serve(listener))
}
