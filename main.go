package main

import (
	"cryptopars/internal"

	"google.golang.org/api/sheets/v4"
)

var idSheets = "1T-G5CjLd3bmtCHVjgMmAAyCcxLqIkFFjPuQVnWuaiIM"

func main() {
	spreadSheetService, client := internal.Connect()

	sheet, _ := sheets.New(client)

	internal.CreateSheet(sheet, idSheets, "test")
	spreadsheet, _ := spreadSheetService.FetchSpreadsheet(idSheets)

	for _, r := range spreadsheet.Sheets {
		internal.DeleteAllSheet(sheet, idSheets, int64(r.Properties.ID))
	}
	spreadsheet, _ = spreadSheetService.FetchSpreadsheet(idSheets)
	internal.CreateSheet(sheet, idSheets, "Coingecko")
	internal.DeleteZeroSheet(sheet, idSheets, spreadsheet)
	internal.CreateSheet(sheet, idSheets, "Cryptorank")
	go internal.CryptoParser(spreadSheetService, idSheets)

	internal.GecParser(spreadSheetService, idSheets)

}
