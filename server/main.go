package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/grpc-testing/pkg/proto/credit"
	"google.golang.org/grpc"
)

type server struct {
	credit.UnimplementedCreditServiceServer
}

func main() {
	log.Println("Server is running ...")

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	srv := grpc.NewServer()
	credit.RegisterCreditServiceServer(srv, &server{})

	log.Fatal(srv.Serve(lis))
}

func (s *server) Credit(ctx context.Context, request *credit.CreditRequest) (*credit.CreditResponse, error) {
	log.Println(fmt.Sprintf("Request : %g", request.GetAmount()))
	return &credit.CreditResponse{Confirmation: fmt.Sprintf("Credited %g", request.GetAmount())}, nil
}
