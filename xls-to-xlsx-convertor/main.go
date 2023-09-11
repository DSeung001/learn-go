package main

// go get github.com/360EntSecGroup-Skylar/excelize => xlsx 전용 라이브러리
// 이 친구로 변경 필요 => https://github.com/extrame/xls
import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
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
		xlsFile, err := excelize.OpenFile(xlsFilePath)
		utils.HandleErr(err)

		xlsxFilePath := changeFileExtension(xlsxDirPath+"/"+fileName, "xlsx")
		utils.HandleErr(xlsFile.SaveAs(xlsxFilePath))

		fmt.Printf("성공적으로 %s 파일을 %s 파일로 변환했습니다.\n", xlsFilePath, xlsxFilePath)
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
