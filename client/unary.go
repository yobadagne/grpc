package main

import (
	"context"
	"log"
	"time"

	pb "github.com/yobadagne/grpc-yt/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := client.SayHello(ctx, &pb.NoParam{})
	if err != nil {
		log.Printf("error while calling say hello rpc: %v", err)
	}
	log.Print(res.Message)
}
