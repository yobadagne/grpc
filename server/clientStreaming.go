package main

import (
	"io"
	"log"

	pb "github.com/yobadagne/grpc-yt/proto"
	"google.golang.org/grpc"
)

func (h *helloServer) SayHelloClientStreaming(req grpc.ClientStreamingServer[pb.HelloRequest, pb.MessageList]) error {
	Messages := []string{}
	for {
		mess, err := req.Recv()
		if err == io.EOF {
			return req.SendMsg(&pb.MessageList {
				Messages: Messages})
		}
		if err != nil {
			log.Fatalf("failed to read client stream: %v", err)
			return err
		}

		Messages = append(Messages, mess.Name)
	}
}
