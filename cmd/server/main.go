package main

import (
	"context"
	"github.com/ElDanto/ShortGen/internal/domain/model"
	pb "github.com/ElDanto/ShortGen/pkg/proto/hash"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedHashServer
}

func (s *server) Do(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v", in)
	hash := model.Hash{}
	return &pb.Response{
		Hash: hash.Generate(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterHashServer(grpcServer, &server{
		UnimplementedHashServer: pb.UnimplementedHashServer{},
	})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
