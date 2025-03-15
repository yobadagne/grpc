package main

import (
	"log"

	pb "github.com/yobadagne/grpc-yt/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial grpc: %v", err)
	}
	defer conn.Close()
	names := pb.NameList{
		Names: []string{"Eyob", "Dagne", "Taye"},
	}
	name := []string{"Eyob", "Dagne", "Taye"}

	messges := []string{
		"Hi", "Eyob",
	}
	client := pb.NewGreetServiceClient(conn)
	callSayHello(client)
	sayHelloServerStreaming(client, &names)
	callClientStreaming(client, messges)
	sayHelloBiDirectional(client, name)

}
