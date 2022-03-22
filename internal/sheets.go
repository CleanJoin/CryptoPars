package internal

import (
	"context"

	"google.golang.org/api/sheets/v4"
	"gopkg.in/Iwark/spreadsheet.v2"
)

func CreateSheet(srv *sheets.Service, spreadsheetId string, nameSheet string) (*sheets.BatchUpdateSpreadsheetResponse, error) {
	req := sheets.Request{
		AddSheet: &sheets.AddSheetRequest{
			Properties: &sheets.SheetProperties{
				Title: nameSheet,
			},
		},
	}

	rbb := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{&req},
	}

	resp, err := srv.Spreadsheets.BatchUpdate(spreadsheetId, rbb).Context(context.Background()).Do()
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func DeleteZeroSheet(srv *sheets.Service, spreadsheetId string, spreadsheet spreadsheet.Spreadsheet) error {
	req := sheets.Request{
		DeleteSheet: &sheets.DeleteSheetRequest{
			SheetId: int64(spreadsheet.Sheets[0].Properties.ID),
		},
	}
	rbb := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{&req},
	}
	_, err := srv.Spreadsheets.BatchUpdate(spreadsheetId, rbb).Context(context.Background()).Do()
	if err != nil {
		return err
	}
	return nil
}
