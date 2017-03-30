package main

import (
	"fmt"
	"log"

	pb "github.com/drgarcia1986/grpc-hw/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	address := "localhost:50051"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("failed to connect: ", err)
	}
	defer conn.Close()

	cli := pb.NewHelloWorldClient(conn)
	r, err := cli.Say(context.Background(), &pb.Request{Name: "World"})
	if err != nil {
		log.Fatal("error on request to server: ", err)
	}

	fmt.Println(r.Msg)
}
