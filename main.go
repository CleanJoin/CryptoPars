package main

import (
	"cryptopars/internal"

	"google.golang.org/api/sheets/v4"
)

var idSheets = "1T-G5CjLd3bmtCHVjgMmAAyCcxLqIkFFjPuQVnWuaiIM"

func main() {
	spreadSheetService, client := internal.Connect()

	sheet, _ := sheets.New(client)
	spreadsheet, _ := spreadSheetService.FetchSpreadsheet(idSheets)
	internal.CreateSheet(sheet, idSheets, "Coingecko")
	internal.DeleteZeroSheet(sheet, idSheets, spreadsheet)
	internal.CreateSheet(sheet, idSheets, "Cryptorank")
	internal.GecParser(spreadSheetService, idSheets)
	internal.CryptoParser(spreadSheetService, idSheets, "")
}
