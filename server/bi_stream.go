package main

import (
	"io"
	"log"

	pb "github.com/weldonkipchirchir/go-grpc-demo/proto"
)

func (s *helloServer) SayHelloBidiStreaming(stream pb.GreetService_SayHelloBidiStreamingServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		log.Printf("Received from client: %v", req.Name)
		res := &pb.HelloResponse{
			Message: "Hello " + req.Name,
		}
		if err := stream.Send(res); err != nil {
			panic(err)
		}
	}
	return nil
}
