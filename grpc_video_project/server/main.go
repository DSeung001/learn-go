package main

import (
	"fmt"
	pb "github.com/your_username/grpc_video_project/proto/videopb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type server struct {
	pb.UnimplementedVideoServiceServer
}

func (s *server) ProcessVideo(stream pb.VideoService_ProcessVideoServer) error {
	// 클라이언트로부터 청크를 수신하며, 각 청크마다 처리 후 상태 스트림으로 응답
	file, _ := os.Create("received_video.mp4")
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			// 클라이언트 전송 완료
			log.Printf("모든 청크 수신 완료")
			return nil
		}
		if err != nil {
			return fmt.Errorf("청크 수신 오류: %v", err)
		}

		if _, err := file.Write(chunk.Data); err != nil {
			return fmt.Errorf("파일 기록 오류: %v", err)
		}

		// 처리 시작 시각 기록 (여기서는 처리 시뮬레이션)
		start := time.Now()
		// 예시 처리 로직: 실제 처리(예: 디코딩, 필터 적용) 대신 간단한 sleep
		time.Sleep(20 * time.Millisecond)
		processingTime := float32(time.Since(start).Seconds())

		status := &pb.ProcessingStatus{
			Sequence:       chunk.Sequence,
			Status:         "processed",
			ProcessingTime: processingTime,
		}

		// 처리 결과를 실시간으로 클라이언트에 전송
		if err := stream.Send(status); err != nil {
			return fmt.Errorf("상태 전송 오류: %v", err)
		}
		log.Printf("청크 %d 처리 완료 (%.3fs)", chunk.Sequence, processingTime)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("포트 리스닝 오류: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterVideoServiceServer(grpcServer, &server{})
	log.Println("gRPC 서버 실행 중: :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("서버 실행 오류: %v", err)
	}
}
