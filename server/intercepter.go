package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func Interceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	log.Print("pre handler interceptor")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("no metadata")
	} else {
		log.Printf("Metadata %v+\n", md)
	}
	res, err := handler(ctx, req)
	log.Print("post handler interceptor")
	return res, err
}
