package main

import (
	"context"
	"log"

	pb "github.com/weldonkipchirchir/go-grpc-demo/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Streaming started")

	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Printf("Error while calling SayHelloServerStreaming: %v", err)
	}
	for {
		res, err := stream.Recv()
		if err != nil {
			log.Printf("Error while receiving from stream: %v", err)
			break
		}
		log.Printf("Received from stream: %v", res.Message)
	}

	log.Printf("Streaming finished")
}
