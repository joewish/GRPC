package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":9000"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	if err2 := grpcServer.Serve(lis); err2 != nil {
		log.Fatal(err2)
	}
}
