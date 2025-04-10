package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"time"

	pb "github.com/your_username/grpc_video_project/proto/videopb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedVideoServiceServer
}

func (s *server) ProcessVideo(stream pb.VideoService_ProcessVideoServer) error {
	// 버퍼링 임계치 설정 (예: 2MB)
	const bufferThreshold = 2 * 1024 * 1024
	preBuffer := bytes.NewBuffer(nil)
	playbackStarted := false

	// ffmpeg 변환 체인으로 데이터를 전달할 파이프 라이터
	var ffmpegInWriter *io.PipeWriter

	// ffmpeg 변환 체인을 시작하는 함수
	startConversionChain := func() {
		ffmpegInReader, writer := io.Pipe()
		ffmpegInWriter = writer

		// ffmpeg 명령어에 추가 옵션으로 -probesize, -analyzeduration, -fflags +genpts 추가
		ffmpegCmd := exec.Command("ffmpeg",
			"-probesize", "5000000",
			"-analyzeduration", "5000000",
			"-fflags", "+genpts",
			"-i", "pipe:0",
			"-movflags", "faststart",
			"-f", "mpegts", "pipe:1")
		ffmpegCmd.Stdin = ffmpegInReader

		ffmpegOutPipe, err := ffmpegCmd.StdoutPipe()
		if err != nil {
			log.Fatalf("ffmpeg stdout 파이프 생성 오류: %v", err)
		}

		if err := ffmpegCmd.Start(); err != nil {
			log.Fatalf("ffmpeg 실행 오류: %v", err)
		}

		go func() {
			cmd := exec.Command("ffplay", "-autoexit", "-")
			cmd.Stdin = ffmpegOutPipe
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Start(); err != nil {
				log.Printf("ffplay 실행 오류: %v", err)
			} else {
				log.Println("플레이어가 실행 중입니다...")
			}
			cmd.Wait()
			log.Println("플레이어 종료")
		}()

		log.Println("버퍼 임계치 도달 - 변환 체인 시작 (MP4 → MPEG‑TS)")
	}

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			if playbackStarted && ffmpegInWriter != nil {
				ffmpegInWriter.Close()
			}
			log.Println("모든 청크 수신 완료")
			break
		}
		if err != nil {
			return fmt.Errorf("청크 수신 오류: %v", err)
		}

		if !playbackStarted {
			_, err := preBuffer.Write(chunk.Data)
			if err != nil {
				return fmt.Errorf("버퍼 기록 오류: %v", err)
			}
			if preBuffer.Len() >= bufferThreshold {
				playbackStarted = true
				startConversionChain()
				_, err := ffmpegInWriter.Write(preBuffer.Bytes())
				if err != nil {
					return fmt.Errorf("ffmpeg로 preBuffer 데이터 기록 오류: %v", err)
				}
				preBuffer.Reset()
			}
		} else {
			_, err := ffmpegInWriter.Write(chunk.Data)
			if err != nil {
				return fmt.Errorf("ffmpeg로 데이터 기록 오류: %v", err)
			}
		}

		start := time.Now()
		time.Sleep(20 * time.Millisecond)
		processingTime := float32(time.Since(start).Seconds())

		status := &pb.ProcessingStatus{
			Sequence:       chunk.Sequence,
			Status:         "processed",
			ProcessingTime: processingTime,
		}
		if err := stream.Send(status); err != nil {
			return fmt.Errorf("상태 전송 오류: %v", err)
		}
		log.Printf("청크 %d 처리 및 전송됨", chunk.Sequence)
	}

	return nil
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
