package main

import (
	"xls-to-xlsx-convertor.com/xlsToXlsx"
)

const xlsDirPath = "resources/xls"
const xlsxDirPath = "resources/xlsx"

func main() {
	// xls 경로랑 xlsx 경로 전달
	xlsToXlsx.FileConversion(xlsDirPath, xlsxDirPath)
}
