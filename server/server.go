package main

import (
	"context"
	"log"
	"net"
	"rpc/proto"

	pb "rpc/proto"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	proto.UnimplementedAddServiceServer
}

func (s *server) Add(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	result := in.GetA() + in.GetB()
	return &pb.Response{Result: result}, nil
}
func (s *server) Subtract(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	result := in.GetA() - in.GetB()
	return &pb.Response{Result: result}, nil
}
func (s *server) Multiply(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	result := in.GetA() * in.GetB()
	return &pb.Response{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAddServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
