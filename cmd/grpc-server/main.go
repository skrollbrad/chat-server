package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
	"net"

	desc "github.com/skrollbrad/microservices/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatV1Server
}

// Create ...
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	name := req.Username
	log.Printf("Names is: %v", name)

	return &desc.CreateResponse{Id: 123}, nil
}
func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*empty.Empty, error) {
	id := req.Id
	log.Printf("This id: %d", id)

	return &empty.Empty{}, nil
}
func (s *server) Send(ctx context.Context, req *desc.SendMessageRequest) (*empty.Empty, error) {
	text := req.Text
	log.Printf("This id: %v", text)

	return &empty.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
