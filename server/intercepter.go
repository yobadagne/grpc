package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	log.Printf("pre handler interceptor")
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Println("no metadata %v", info.FullMethod)
	} else {
		log.Printf("Metadata %v+\n", md)
	}
	res, err := handler(ctx, req)
	log.Printf("post handler interceptor %v", info.FullMethod)
	return res, err
}



type wrappedStream struct {
	grpc.ServerStream
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	log.Printf("Received %T - %v\n", m, m)

	return w.ServerStream.RecvMsg(m)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	log.Printf("Sent %T - %v\n", m, m)

	return w.ServerStream.SendMsg(m)
}

func newWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &wrappedStream{s}
}

func ServerStreamIterceptor(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	log.Printf("pre server stream interceptor %v", info.FullMethod)

	err := handler(srv, newWrappedStream(ss))	
	if err != nil {
		log.Printf("failed server stream: %v", err)
		return err
	}

	log.Printf("post server stream interceptor %v", info.FullMethod)
	return nil
}
