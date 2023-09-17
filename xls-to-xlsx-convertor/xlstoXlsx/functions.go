package xlstoXlsx

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/extrame/xls"
	"log"
	"os"
	"xls-to-xlsx-convertor.com/utils"
)

// FileConversion : xlsDirPath 경로에 파일들을 xlsx 로 바꿔서 xlsxDirPath 에 저장
func FileConversion(xlsDirPath string, xlsxDirPath string) {
	fileNames := getFileNames(xlsDirPath)

	for _, fileName := range fileNames {
		var data [][]string
		var arr []string

		// Excel 파일 열기
		xlsFile, err := xls.Open(fileName, "utf-8")
		if err != nil {
			log.Fatalf("엑셀 파일 열기 실패: %v", err)
		}

		for sheet := 0; sheet < xlsFile.NumSheets(); sheet++ {
			sheet := xlsFile.GetSheet(sheet)
			// 행 순회
			for rowIdx := 0; rowIdx <= int(sheet.MaxRow); rowIdx++ {
				row := sheet.Row(rowIdx)
				data = append(data, arr)
				if row != nil {
					// 열 순회
					for colIdx := 0; colIdx < row.LastCol(); colIdx++ {
						cell := row.Col(colIdx)
						data[rowIdx] = append(data[rowIdx], cell)
					}
				}
			}
		}
		utils.HandleErr(createXLSXFile(data, xlsxDirPath+"/"+fileName+"x"))
	}
}

// getFileNames : dirPath 에 있는 파일 가져오기
func getFileNames(dirPath string) []string {
	var fileNames []string

	// 폴더 열기
	dir, err := os.Open(dirPath)
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
			fileNames = append(fileNames, dirPath+"/"+fileInfo.Name())
		}
	}

	return fileNames
}

// createXLSXFile : XLS 파일 만들기
func createXLSXFile(data [][]string, fileName string) error {
	// 파일 생성
	file := excelize.NewFile()

	for rowIndex, row := range data {
		for colIndex, cellValue := range row {
			// A1 형태로 위치 값 생성
			cellPosition := excelize.ToAlphaString(colIndex) + fmt.Sprintf("%d", rowIndex+1)
			file.SetCellValue("Sheet1", cellPosition, cellValue)
		}
	}

	// 시트 저장
	if err := file.SaveAs(fileName); err != nil {
		return err
	}
	return nil
}
