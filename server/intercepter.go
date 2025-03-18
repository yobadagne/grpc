package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
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
func ServerStreamIterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Print("pre server stream interceptor")
	err := handler(srv, ss)
	if err != nil {
		log.Printf("failed server stream: %v", err)
		return err
	}

	log.Print("post server stream interceptor")
	return nil
}
