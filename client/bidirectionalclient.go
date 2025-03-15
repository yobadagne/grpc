package main

import (
	"context"
	"io"
	"log"

	pb "github.com/yobadagne/grpc-yt/proto"
)

func sayHelloBiDirectional(client pb.GreetServiceClient, names []string) {
	stream, err := client.SayHelloBiDirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("failed to create bi directional stream : %v", err)
	}
	for _, v := range names {
		err := stream.Send(&pb.HelloRequest{
			Name: v,
		})

		if err != nil {
			log.Fatalf("failed to send bi directional stream ")
		}

	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to send bi directional stream ")
		}
		log.Printf("message is: %v", res.Message)
	}

}
