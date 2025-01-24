package main

import (
	"image"
	"image/color"
	"image/png"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	inputDir := "./png"         // PNG 이미지가 저장된 폴더 경로
	outputDir := "./result_png" // 결과를 저장할 폴더 경로

	// 결과 저장 폴더 생성
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("결과 폴더 생성 실패: %v", err)
	}

	// 입력 폴더의 파일 탐색
	err = filepath.Walk(inputDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// PNG 파일만 처리
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".png") {
			processImage(path, filepath.Join(outputDir, info.Name()))
		}
		return nil
	})
	if err != nil {
		log.Fatalf("파일 탐색 실패: %v", err)
	}

	log.Println("처리 완료")
}

func processImage(inputPath, outputPath string) {
	// 이미지 파일 열기
	file, err := os.Open(inputPath)
	if err != nil {
		log.Printf("이미지 열기 실패 (%s): %v", inputPath, err)
		return
	}
	defer file.Close()

	// 이미지 디코딩
	img, err := png.Decode(file)
	if err != nil {
		log.Printf("이미지 디코딩 실패 (%s): %v", inputPath, err)
		return
	}

	// 새로운 RGBA 이미지 생성
	bounds := img.Bounds()
	outputImg := image.NewRGBA(bounds)

	// 픽셀 순회하여 흰색을 투명으로 변경
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			originalColor := img.At(x, y)
			r, g, b, _ := originalColor.RGBA()

			// 흰색 픽셀(255, 255, 255)을 투명 처리
			if r>>8 == 255 && g>>8 == 255 && b>>8 == 255 {
				outputImg.Set(x, y, color.RGBA{0, 0, 0, 0}) // 투명
			} else {
				outputImg.Set(x, y, originalColor)
			}
		}
	}

	// 결과 이미지 저장
	outputFile, err := os.Create(outputPath)
	if err != nil {
		log.Printf("결과 저장 실패 (%s): %v", outputPath, err)
		return
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, outputImg)
	if err != nil {
		log.Printf("PNG 인코딩 실패 (%s): %v", outputPath, err)
	}
}
