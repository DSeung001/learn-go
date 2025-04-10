package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	pb "github.com/your_username/grpc_video_project/proto/videopb"
	"google.golang.org/grpc"
)

// 32KB 버퍼 풀 생성 (메모리 최적화)
var bufferPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 32*1024)
	},
}

func processVideo(client pb.VideoServiceClient, filePath string) error {
	// 5분 타임아웃 설정
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	stream, err := client.ProcessVideo(ctx)
	if err != nil {
		return fmt.Errorf("스트림 생성 오류: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	// 1. 청크 전송 고루틴 (동영상 파일을 읽어 청크 전송)
	go func() {
		defer wg.Done()
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("파일 열기 오류: %v\n", err)
			return
		}
		defer file.Close()

		sequence := int32(1)
		for {
			buf := bufferPool.Get().([]byte)
			n, err := file.Read(buf)
			if err != nil {
				if err == io.EOF {
					bufferPool.Put(buf)
					break
				}
				fmt.Printf("파일 읽기 오류: %v\n", err)
				bufferPool.Put(buf)
				break
			}

			req := &pb.VideoChunk{
				Data:     buf[:n],
				Sequence: sequence,
			}
			sequence++
			if err := stream.Send(req); err != nil {
				fmt.Printf("청크 전송 오류: %v\n", err)
				bufferPool.Put(buf)
				break
			}
			bufferPool.Put(buf)
		}
		// 전송 완료 시 스트림 종료 알림
		stream.CloseSend()
	}()

	// 2. 처리 상태 수신 고루틴 (서버의 실시간 응답을 콘솔에 출력)
	go func() {
		defer wg.Done()
		for {
			status, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Printf("상태 수신 오류: %v\n", err)
				break
			}
			// 실시간 처리 결과 출력
			fmt.Printf("청크 %d: %s (%.3fs)\n", status.Sequence, status.Status, status.ProcessingTime)
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(10*1024*1024),
			grpc.MaxCallSendMsgSize(10*1024*1024),
		),
	)
	if err != nil {
		fmt.Printf("연결 실패: %v\n", err)
		return
	}
	defer conn.Close()

	client := pb.NewVideoServiceClient(conn)
	if err := processVideo(client, "../video.mp4"); err != nil {
		fmt.Printf("프로세싱 오류: %v\n", err)
	}
}
