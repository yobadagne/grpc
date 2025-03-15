package main

import (
	"io"
	"log"

	pb "github.com/yobadagne/grpc-yt/proto"
	"google.golang.org/grpc"
)

func (h *helloServer) SayHelloBiDirectionalStreaming(stream grpc.BidiStreamingServer[pb.HelloRequest, pb.HelloResponse]) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("failed to receive the request: %v", err)
			return err
		}

		if err := stream.Send(&pb.HelloResponse{
			Message: "Hello " + req.Name,
		}); err != nil {
			log.Fatalf("failed to send the response: %v", err)
			return err
		}
	}
	return nil

}
