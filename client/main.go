package main

import (
	"log"

	pb "github.com/GRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":9000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials((insecure.NewCredentials())))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn) //

	// name := &pb.NamesList{
	// 	Names: []string{"Joewish","Alice","Guddi"}
	// }
	//callSayHello(client)
}
