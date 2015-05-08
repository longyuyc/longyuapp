package exportExcel

import (
	"fmt"
	"github.com/go-martini/martini"
	"github.com/tealeg/xlsx"
)

//r.Get("/exportBom", WebBomData)
func ExportExcelApiRoute(m martini.Router) {
	m.Get("/", ExportDataToExcel)
	m.Get("/Get", GetExportData)
}

func GetExportData() {
	excelFileName := "Bom.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println(err)
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				fmt.Printf("%s\n", cell.String())
			}
		}
	}
}

func ExportDataToExcel() {
	//导出到excel文件中
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var sheet2 *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var errs error
	file = xlsx.NewFile()
	sheet = file.AddSheet("Sheet1")

	for i := 0; i < 5; i++ {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = "a"
		cell = row.AddCell()
		cell.Value = "b"
		cell = row.AddCell()
		cell.Value = "c"
	}

	sheet2 = file.AddSheet("Sheet2")
	for i := 0; i < 5; i++ {
		row = sheet2.AddRow()
		cell = row.AddCell()
		cell.Value = "a2"
		cell = row.AddCell()
		cell.Value = "b2"
		cell = row.AddCell()
		cell.Value = "c2"
	}

	errs = file.Save("Bom.xlsx")
	if errs != nil {
		fmt.Println(errs)
	}
}
