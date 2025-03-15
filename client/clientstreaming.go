package main

import (
	"context"
	"log"
	"time"

	pb "github.com/yobadagne/grpc-yt/proto"
)

func callClientStreaming(client pb.GreetServiceClient, Messages []string){
	ctx,cancel := context.WithTimeout(context.Background(),3* time.Second )
	defer cancel()
	stream, err := client.SayHelloClientStreaming(ctx)
	if err != nil {
		log.Fatalf("failed to start client stream: %v", err)
	}

	for _, v := range Messages {
		stream.Send(&pb.HelloRequest{
			Name: v,
		})
	}	
	message, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to start client stream: %v", err)
	}
	log.Printf("message is %v", message.Messages)

}