package main

import (
	"io"
	"log"

	pb "github.com/weldonkipchirchir/go-grpc-demo/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		log.Printf("Received from client: %v", message.Name)

		messages = append(messages, message.Name)

	}
	return nil
}
