package main

import (
	"context"
	"fmt"
	"net"

	proto "github.com/smmile18/service-proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	listener, err := net.Listen("tcp", ":3000") // 4040u
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)
	fmt.Println("server 2 start....")
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) Add(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a - b
	fmt.Println(result)

	return &proto.Response{Result: result}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	a, b := request.GetA(), request.GetB()

	result := a / b
	fmt.Println(result)

	// talk with server1
	conn, err := grpc.Dial("lab-server.default.svc.cluster.local:3000", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(conn)
	defer conn.Close()

	req := &proto.Request{A: a, B: b}
	if response, err := client.Add(ctx, req); err == nil {
		return &proto.Response{Result: response.Result}, nil
	} else {
		return nil, err
	}

}
