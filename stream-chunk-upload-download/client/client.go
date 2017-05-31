package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/context"

	pb "github.com/drgarcia1986/grpc-hw/stream-chunk-upload-download/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	cli := pb.NewUploaderClient(conn)

	stream, err := cli.Upload(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	msg := &pb.Msg{Value: &pb.Msg_Meta_{Meta: &pb.Msg_Meta{Filename: "foo.txt"}}}
	if err := stream.Send(msg); err != nil {
		panic(err)
	}

	f, err := os.Open("foo.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	r := bufio.NewReader(f)
	buf := make([]byte, 1024)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		bufMsg := &pb.Msg{Value: &pb.Msg_File_{File: &pb.Msg_File{Data: buf}}}
		if err := stream.Send(bufMsg); err != nil {
			panic(err)
		}
	}
	stream.CloseSend()

	// TODO: read file from stream
	fmt.Println("Done")
}
