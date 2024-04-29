package main

import (
	"context"
	"log"

	"github.com/kevalsabhani/go-protobuf/protofiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server at localhost:50051: %v", err)
	}
	defer conn.Close()
	client := protofiles.NewPersonServiceClient(conn)

	res, err := client.GetAddressBook(context.Background(), &protofiles.None{})
	if err != nil {
		log.Fatalf("error calling function SayHello: %v", err)
	}
	log.Printf("+%v", res)
}
