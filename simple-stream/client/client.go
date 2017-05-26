package main

import (
	"io"
	"log"

	"golang.org/x/net/context"

	pb "github.com/drgarcia1986/grpc-hw/simple-stream/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cli := pb.NewMessagesClient(conn)

	ctx := context.Background()
	stream, err := cli.Echo(ctx, &pb.Msg{Msg: "hello world"})
	if err != nil {
		log.Fatal(err)
	}

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(in.Msg))
	}
}
