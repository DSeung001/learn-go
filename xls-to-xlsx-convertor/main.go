package main

import (
	"xls-to-xlsx-convertor.com/xlstoXlsx"
)

const xlsDirPath = "resources/xls"
const xlsxDirPath = "resources/xlsx"

func main() {
	// xls 경로랑 xlsx 경로 전달
	xlstoXlsx.FileConversion(xlsDirPath, xlsxDirPath)

}
