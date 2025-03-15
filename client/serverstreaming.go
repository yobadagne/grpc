package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/yobadagne/grpc-yt/proto"
)

func sayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NameList) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	res, err := client.SayHelloServiceStreaming(ctx, names)
	if err != nil {
		log.Fatalf("failed to read server stream: %v", err)
		return err
	}
	for {
		message, err := res.Recv()
		if err != io.EOF && err != nil{
			log.Fatalf("error received: %v", err)
			return err 
		}else if err == io.EOF{
			log.Print("last message received ")
			return nil
		} 
		log.Printf("received message: %v", message)
	}
}
