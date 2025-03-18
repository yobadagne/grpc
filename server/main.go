package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/yobadagne/grpc-yt/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type helloServer struct {
	pb.UnimplementedGreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen on %v error : %v", port, err)
	}

	grpcserver := grpc.NewServer(
		grpc.UnaryInterceptor(Interceptor),
	)
	pb.RegisterGreetServiceServer(grpcserver, &helloServer{})
	reflection.Register(grpcserver)
	go func() {
		if err := grpcserver.Serve(lis); err != nil {
			log.Fatalf("failed to start : %v", err)
		}
	}()
	log.Printf("listening on port: %v", port)
	// gracefull shutdown
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, os.Interrupt)
	signal.Notify(sign, syscall.SIGTERM)

	<-sign
	log.Print("shutting server gracefully")
	grpcserver.GracefulStop()
}
