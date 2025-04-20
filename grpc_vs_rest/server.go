package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc_vs_rest/proto/user"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) RegisterUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("gRPC - Received user: %+v", req)
	// 실제 회원가입 로직 (예: DB 저장)
	return &pb.UserResponse{Status: "success"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{})
	log.Println("gRPC server is running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
