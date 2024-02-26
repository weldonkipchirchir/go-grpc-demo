package main

import (
	"log"
	"time"

	pb "github.com/weldonkipchirchir/go-grpc-demo/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Received request from client: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello, " + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(time.Second * 4)
	}
	return nil
}
