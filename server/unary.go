package main

import (
	"context"
	pb "github.com/yobadagne/grpc-yt/proto"

)
func (h *helloServer) SayHello(context.Context, *pb.NoParam) (*pb.HelloResponse, error){
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
