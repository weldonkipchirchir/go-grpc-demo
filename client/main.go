package main

import (
	"io"
	"log"

	pb "github.com/weldonkipchirchir/go-grpc-demo/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = "localhost:8080"
)

func main() {
	connection, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err == io.EOF {
		log.Fatalf("did not connect: %v", err)
	}
	if err != nil {
		panic(err)
	}

	defer connection.Close()

	client := pb.NewGreetServiceClient(connection)

	names := &pb.NameList{
		Names: []string{"Weldon", "Kipchirchir", "job", "dorsal", "sacral"},
	}
	//unary api
	// callSayHello(client)

	// callSayHelloServerStream(client, names)

	// callSayHelloClientStream(client, names)

	CallHelloBidirectionalStream(client, names)

}
