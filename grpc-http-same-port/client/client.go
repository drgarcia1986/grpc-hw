package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	pb "github.com/drgarcia1986/grpc-hw/grpc-http-same-port/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	address := "localhost:8080"
	grpcReturn, err := callGrpcServer(address)
	if err != nil {
		fmt.Println("gRPC failed:", err)
	}
	httpReturn, err := callHttpServer(address)
	if err != nil {
		fmt.Println("http failed:", err)
	}
	fmt.Printf("grpc: %s, http: %s\n", grpcReturn, httpReturn)
}

func callGrpcServer(address string) (string, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	cli := pb.NewHelloWorldClient(conn)
	r, err := cli.Say(context.Background(), &pb.Req{Name: "World"})
	if err != nil {
		return "", err
	}
	return r.Msg, nil
}

func callHttpServer(address string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("http://%s/ping/", address))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
