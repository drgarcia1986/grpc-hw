package main

import (
	"fmt"
	"log"

	pb "github.com/drgarcia1986/grpc-hw/simple-auth/pb"
	"golang.org/x/net/context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type tokenAuth struct {
	token string
}

func (t *tokenAuth) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{"token": t.token}, nil
}

func (*tokenAuth) RequireTransportSecurity() bool { return false }

func main() {
	server := "localhost:50051"
	token, err := login(server)
	if err != nil {
		if grpc.Code(err) == codes.PermissionDenied {
			log.Fatal("Access Denied, try again")
		}
		log.Fatal("login: ", err)
	}

	msg, err := restrictedCall(server, token)
	if err != nil {
		log.Fatal("login: ", err)
	}
	fmt.Println(msg)
}

func restrictedCall(server, token string) (string, error) {
	conn, err := grpc.Dial(server, grpc.WithPerRPCCredentials(&tokenAuth{token}), grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	cli := pb.NewSimpleAuthClient(conn)
	res, err := cli.Restricted(context.Background(), &pb.Empty{})
	if err != nil {
		return "", err
	}
	return res.Message, nil
}

func login(server string) (string, error) {
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		return "", err
	}
	defer conn.Close()

	cli := pb.NewSimpleAuthClient(conn)
	r, err := cli.Login(context.Background(), &pb.LoginRequest{
		User:     "admin",
		Password: "ananas2000",
	})
	if err != nil {
		return "", err
	}
	return r.Token, nil
}
