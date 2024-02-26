package main

import (
	"log"
	"net"

	pb "github.com/weldonkipchirchir/go-grpc-demo/proto"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	log.Printf("Server listening on port %s", port)
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
