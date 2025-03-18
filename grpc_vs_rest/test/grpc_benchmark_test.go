package test

import (
	"context"
	"testing"
	"time"

	"google.golang.org/grpc"
	pb "grpc_vs_rest/proto/user"
)

func BenchmarkGrpcRegisterParallel(b *testing.B) {
	// 동시에 실행되는 고루틴 수를 제한 (예: 4)
	b.SetParallelism(4)

	// gRPC 연결 생성 (연결 재사용됨)
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		b.Logf("failed to connect: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	req := &pb.UserRequest{
		LastName:  "Kim",
		FirstName: "Minji",
		Phone:     "010-1234-5678",
		Email:     "minji@example.com",
		Gender:    "F",
		BirthDate: "1990-01-01",
		Username:  "minji90",
	}

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			_, err := client.RegisterUser(ctx, req)
			if err != nil {
				b.Logf("gRPC call error: %v", err)
				cancel()
				continue
			}
			cancel()
		}
	})
}
