package main

import (
	"context"
	"time"

	pb "github.com/weldonkipchirchir/go-grpc-demo/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		panic(err)
	}

	println(res.Message)
}
