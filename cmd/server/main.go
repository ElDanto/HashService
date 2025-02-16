package main

import (
	"context"
	pb "github.com/ElDanto/HashService/pkg/pb/hash"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedReverseServer
}

func (s *server) Do(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	log.Printf("Received: %v", in)

	return &pb.Response{
		Hash: "Test",
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
	pb.RegisterReverseServer(grpcServer, &server{
		UnimplementedReverseServer: pb.UnimplementedReverseServer{},
	})

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
