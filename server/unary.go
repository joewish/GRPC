package main

import (
  "context"
  pb "github.com/GRPC/proto"
)
func (s *helloServer)  SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error){
  return &pb.HelloResponse{
    message : "hello"
  },nil
}