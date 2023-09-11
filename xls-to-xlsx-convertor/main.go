package main

// go get github.com/360EntSecGroup-Skylar/excelize 사용
import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"os"
	"path/filepath"
	"xls-to-xlsx-convertor.com/utils"
)

const xlsDirPath = "xls"
const xlsxDirPath = "xlsx"

func main() {
	// XLS 파일 경로
	xlsFilePath := "example.xls"
	fileNames := getStationNames()

	for _, fileName := range fileNames {
		xlsFile, err := excelize.OpenFile(xlsDirPath + "/" + fileName)
		xlsxFilePath := changeFileExtension(xlsxDirPath+"/"+fileName, "xlsx")

		utils.HandleErr(err)
		utils.HandleErr(xlsFile.SaveAs(xlsxFilePath))

		fmt.Printf("성공적으로 %s 파일을 %s 파일로 변환했습니다.\n", xlsFilePath, xlsxFilePath)
	}
}

// 파일 확장자 변경 함수
func changeFileExtension(filePath, newExtension string) string {
	fileBase := filepath.Base(filePath)
	fileDir := filepath.Dir(filePath)
	fileName := fileBase[:len(fileBase)-len(filepath.Ext(fileBase))]
	return filepath.Join(fileDir, fileName+"."+newExtension)
}

func getStationNames() []string {
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
