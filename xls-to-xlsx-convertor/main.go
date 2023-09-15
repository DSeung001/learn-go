package main

// go get github.com/360EntSecGroup-Skylar/excelize => xlsx 전용 라이브러리
// 이 친구로 변경 필요 => https://github.com/extrame/xls
import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/extrame/xls"
	"log"
	"os"
	"xls-to-xlsx-convertor.com/utils"
)

const xlsDirPath = "resources/xls"
const xlsxDirPath = "resources/xlsx"

func main() {
	// XLS 파일 경로
	fileNames := getXlsFileNames()

	for _, fileName := range fileNames {
		xlsFilePath := xlsDirPath + "/" + fileName
		var data [][]string
		var arr []string

		// Excel 파일 열기
		xlFile, err := xls.Open(xlsFilePath, "utf-8")
		if err != nil {
			log.Fatalf("엑셀 파일 열기 실패: %v", err)
		}

		for i := 0; i < xlFile.NumSheets(); i++ {
			sheet := xlFile.GetSheet(i)
			// 행 순회
			for rowIdx := 0; rowIdx <= int(sheet.MaxRow); rowIdx++ {
				row := sheet.Row(rowIdx)

				data = append(data, arr)

				if row != nil {
					// 열 순회
					for colIdx := 0; colIdx < row.LastCol(); colIdx++ {
						cell := row.Col(colIdx)
						data[rowIdx] = append(data[rowIdx], cell)

						fmt.Printf("row : %d col : %d length : %d value : %v \n", rowIdx, colIdx, len(data[rowIdx]), cell)
					}
				}
			}
			fmt.Println(data)
		}

		utils.HandleErr(createXLSXFile(data, xlsxDirPath+"/"+fileName+"x"))
	}

}

func createXLSXFile(data [][]string, fileName string) error {
	// 새로운 워크북 생성
	file := excelize.NewFile()

	for rowIndex, row := range data {
		for colIndex, cellValue := range row {
			cell := excelize.ToAlphaString(colIndex) + fmt.Sprintf("%d", rowIndex+1)
			file.SetCellValue("Sheet1", cell, cellValue)
		}
	}

	// 시트 저장
	if err := file.SaveAs(fileName); err != nil {
		return err
	}

	return nil
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
