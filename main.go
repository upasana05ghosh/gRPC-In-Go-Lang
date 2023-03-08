package main

import (
	"context"
	"log"
	"net"

	hello "github.com/ughosh/grcp-learn/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type myHelloMsgServer struct {
	hello.UnimplementedHelloMsgServiceServer
}

func (s myHelloMsgServer) Hello(ctx context.Context, req *hello.HelloRequest) (*hello.HelloResponse, error) {
	return &hello.HelloResponse{
		Msg: "Hello " + req.GetName(),
	}, nil
}

func main() {
	// Start a tcp connection on port 8080
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("cannot create listnere: %s", err)
	}
	log.Print("server started")
	//Create new grpc server
	serverRegistrar := grpc.NewServer()
	service := &myHelloMsgServer{}
	// We've to register our endpoint with grpc
	//It takes 2 arg, one is serverRegistrar and other HelloMsgServiceServer interface
	//For this, need to create myHelloMsgServer and implement all methods
	hello.RegisterHelloMsgServiceServer(serverRegistrar, service)

	//Required for grpc curl testing . Not a must steps
	reflection.Register(serverRegistrar)

	//Start server to listen on our configured port
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatalf("not possible to server: %s", err)
	}
}
