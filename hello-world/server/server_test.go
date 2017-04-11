package main

import (
	"testing"

	pb "github.com/drgarcia1986/grpc-hw/hello-world/pb"
	"golang.org/x/net/context"
)

func TestSay(t *testing.T) {
	s := new(server)
	r, err := s.Say(context.Background(), &pb.Request{Name: "World"})
	if err != nil {
		t.Fatal("err on call method say: ", err)
	}

	expected := "Hello World"
	if r.Msg != expected {
		t.Errorf("expected %s, got %s", expected, r.Msg)
	}
}
