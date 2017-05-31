package main

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/drgarcia1986/grpc-hw/stream-chunk-upload-download/pb"
)

type server struct{}

func (*server) Upload(stream pb.Uploader_UploadServer) error {
	var filename string
	content := new(bytes.Buffer)
	for {
		in, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if meta := in.GetMeta(); meta != nil {
			filename = meta.Filename
			continue
		}
		if data := in.GetFile(); data != nil {
			content.Write(data.Data)
		}
	}
	if filename == "" {
		return errors.New("Invalid filename")
	}

	// TODO: write file to stream
	return ioutil.WriteFile(filename, content.Bytes(), 0644)
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("failed to listen:", err)
	}

	s := grpc.NewServer()
	pb.RegisterUploaderServer(s, &server{})

	log.Println("start server")
	log.Println("gRPC server: ", s.Serve(listener))
}
