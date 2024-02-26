package main

import (
	"context"
	"log"
	"time"

	pb "github.com/weldonkipchirchir/go-grpc-demo/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		panic(err)
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{Name: name}
		err = stream.Send(req)

		if err != nil {
			panic(err)
		}
		log.Printf("Sent to client: %v", name)
		time.Sleep(time.Second * 4)
	}
	res, err := stream.CloseAndRecv()

	if err != nil {
		panic(err)
	}
	log.Printf("Received from server: %v", res.Messages)

	log.Printf("Streaming finished")
}
