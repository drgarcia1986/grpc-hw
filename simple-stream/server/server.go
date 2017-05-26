package main

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc"

	pb "github.com/drgarcia1986/grpc-hw/simple-stream/pb"
)

type server struct{}

func (*server) Echo(req *pb.Msg, stream pb.Messages_EchoServer) error {
	ctx := stream.Context()

	for _, s := range req.Msg {
		select {
		case <-ctx.Done():
			log.Println("context Done")
			return nil
		case <-time.After(1000 * time.Millisecond):
			if err := stream.Send(&pb.Msg{Msg: string(s)}); err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	s := grpc.NewServer()
	pb.RegisterMessagesServer(s, &server{})

	log.Println("start server")
	log.Println("gRPC server: ", s.Serve(listener))
}
