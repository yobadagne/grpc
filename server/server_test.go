package main

import (
	"context"
	"log"
	"net"
	"testing"

	pb "github.com/yobadagne/grpc-yt/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func TestGreetServiceSayHello(t *testing.T) {
	//set up the server
	lis := bufconn.Listen(1024 * 1024)
	grpcserver := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcserver, &helloServer{})
	go func() {
		if err := grpcserver.Serve(lis); err != nil {
			log.Fatalf("failed to start : %v", err)
		}
	}()

	// - Test
	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(context.Background(), "bufnet",
		grpc.WithContextDialer(dialer),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("Failed to dial grpc: %v", err)
	}
	t.Cleanup(func() {
		conn.Close()
	})
	if err != nil {
		t.Fatalf("grpc.NewClient %v", err)
	}
	client := pb.NewGreetServiceClient(conn)
	res, err := client.SayHello(context.Background(), &pb.NoParam{})
	if err != nil {
		t.Fatalf("client.SayHello %v", err)
	}

	if res.Message != "Hello" {
		t.Fatalf("Unexpected values %v", res.Message)
	}
}
