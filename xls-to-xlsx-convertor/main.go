package main

// go get github.com/360EntSecGroup-Skylar/excelize => xlsx 전용 라이브러리
// 이 친구로 변경 필요 => https://github.com/extrame/xls
import (
	"fmt"
	"github.com/extrame/xls"
	"log"
	"os"
	"path/filepath"
	"xls-to-xlsx-convertor.com/utils"
)

const xlsDirPath = "resources/xls"
const xlsxDirPath = "resources/xlsx"

func main() {
	// XLS 파일 경로
	fileNames := getXlsFileNames()

	for _, fileName := range fileNames {
		xlsFilePath := xlsDirPath + "/" + fileName

		// Excel 파일 열기
		xlFile, err := xls.Open(xlsFilePath, "utf-8")
		if err != nil {
			log.Fatalf("엑셀 파일 열기 실패: %v", err)
		}

		for i := 0; i < xlFile.NumSheets(); i++ {
			sheet := xlFile.GetSheet(i)
			// 행 순회
			for rowIdx := 0; rowIdx <= int(sheet.MaxRow); rowIdx++ {

				// 행이 값이 없을 경우 에러가 발생 => 어떻게 해결하지
				row := sheet.Row(rowIdx)

				if row != nil {
					// 열 순회
					for colIdx := 0; colIdx < row.LastCol(); colIdx++ {
						cell := row.Col(colIdx)
						fmt.Printf("행: %d, 열: %d, 값: %s\n", rowIdx, colIdx, cell)
					}
				}
			}
		}
	}
}

// changeFileExtension : 파일 확장자 변경
func changeFileExtension(filePath, newExtension string) string {
	fileBase := filepath.Base(filePath)
	fileDir := filepath.Dir(filePath)
	fileName := fileBase[:len(fileBase)-len(filepath.Ext(fileBase))]
	return filepath.Join(fileDir, fileName+"."+newExtension)
}

// getXlsFileNames : xls 파일 이름 가져오기
func getXlsFileNames() []string {
	var fileNames []string

	// 폴더 열기
	dir, err := os.Open(xlsDirPath)
	utils.HandleErr(err)
	defer func(dir *os.File) {
		err := dir.Close()
		utils.HandleErr(err)
	}(dir)

	// 디렉토리 내 파일 목록 읽기
	fileInfos, err := dir.Readdir(-1)
	utils.HandleErr(err)

	// 파일 이름을 리스트에 넣기
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			fileNames = append(fileNames, fileInfo.Name())
		}
	}

	return fileNames
}
