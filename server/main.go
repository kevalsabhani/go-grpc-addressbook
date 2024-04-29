package main

import (
	"log"
	"net"

	"github.com/kevalsabhani/go-protobuf/protofiles"
	"github.com/kevalsabhani/go-protobuf/server/handlers"
	"google.golang.org/grpc"
)

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("tcp error: ", err)
	}

	s := grpc.NewServer()
	protofiles.RegisterPersonServiceServer(s, &handlers.Server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
