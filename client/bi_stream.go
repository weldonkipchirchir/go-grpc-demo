package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/weldonkipchirchir/go-grpc-demo/proto"
)

func CallHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
	log.Printf("Streaming started")

	stream, err := client.SayHelloBidiStreaming(context.Background())

	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}

	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()

			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			log.Printf("Received from client: %v", message)
		}
	}()

	for _, name := range names.Names {
		if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		time.Sleep(time.Second * 4)
	}

	stream.CloseSend()
	<-waitc
	log.Printf("Streaming ended")
}
