package internal

import (
	"time"

	"gopkg.in/Iwark/spreadsheet.v2"
)

func WriteTable(gecData *GecData, spreadSheetService *spreadsheet.Service, idSheets string) int {
	spreadsheet, _ := spreadSheetService.FetchSpreadsheet(idSheets)
	sheet, err := spreadsheet.SheetByIndex(0)
	CheckError(err)

	rowIndex := 0

	for range spreadsheet.Sheets[0].Rows {

		rowIndex++
	}

	if rowIndex <= 65 {

		sheet.Update(0, 0, "Coin Name")
		sheet.Update(0, 1, "USD")
		sheet.Update(0, 2, "Дата и Время")

		sheet.Update(rowIndex, 0, gecData.Coin)
		sheet.Update(rowIndex, 1, gecData.Money)
		sheet.Update(rowIndex, 2, gecData.Time)
		sheet.Synchronize()
	}

	return rowIndex

}

func WriteTableCry(gecData *CrypData, spreadSheetService *spreadsheet.Service, idSheets string) int {
	spreadsheet, _ := spreadSheetService.FetchSpreadsheet(idSheets)
	sheet, err := spreadsheet.SheetByIndex(1)
	CheckError(err)

	rowIndex := 0

	for range spreadsheet.Sheets[1].Rows {
		rowIndex++
	}

	if rowIndex <= 3 {
		time.Sleep(time.Millisecond + 2)
		sheet.Update(0, 0, "Coin Name")
		sheet.Update(0, 1, "Tag")
		sheet.Update(0, 2, "Дата и Время")

		sheet.Update(rowIndex, 0, gecData.Name)
		sheet.Update(rowIndex, 1, gecData.Tag)
		sheet.Update(rowIndex, 2, gecData.Time)
		sheet.Synchronize()
	}

	return rowIndex

}
