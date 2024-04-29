package main

import (
	"context"
	"log"
	"net"

	"github.com/kevalsabhani/go-protobuf/protofiles"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type server struct {
	protofiles.UnimplementedPersonServiceServer
}

func (s *server) GetAddressBook(context.Context, *protofiles.None) (*protofiles.AddressBook, error) {
	person := protofiles.Person{
		Id:    8325923,
		Name:  "Foo bar",
		Email: "foo@bar.com",
		Phones: []*protofiles.PhoneNumber{
			{
				Number: "9876543210",
				Type:   protofiles.PhoneType_WORK,
			},
			{
				Number: "1234567890",
				Type:   protofiles.PhoneType_HOME,
			},
		},
		LastUpdated: timestamppb.Now(),
	}
	addressBook := &protofiles.AddressBook{
		People: []*protofiles.Person{
			&person,
		},
	}
	return addressBook, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("tcp error: ", err)
	}

	s := grpc.NewServer()
	protofiles.RegisterPersonServiceServer(s, &server{})
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
