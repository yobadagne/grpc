package main

import (
	"time"

	pb "github.com/yobadagne/grpc-yt/proto"
	"google.golang.org/grpc"
)

func (h *helloServer) SayHelloServiceStreaming(in *pb.NameList, stream grpc.ServerStreamingServer[pb.HelloResponse]) error {
	for _, v := range in.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + v,
		}
		if err := stream.Send(res); err != nil{
			return err
		}
		time.Sleep(time.Second)
	}
	return nil 
}
