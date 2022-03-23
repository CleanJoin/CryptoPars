package internal

import (
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/Iwark/spreadsheet.v2"
)

type CrypData struct {
	Name string
	Tag  string
	Time string
}

func CryptoParser(spreadSheetService *spreadsheet.Service, idSheets string) error {

	resp, err := http.Get("https://cryptorank.io/")
	if err != nil {
		log.Println(err)
		time.Sleep(time.Minute + 1)
		CryptoParser(spreadSheetService, idSheets)
	}

	parseHtml(resp, spreadSheetService, idSheets)

	return nil
}

func parseHtml(resp *http.Response, spreadSheetService *spreadsheet.Service, idSheets string) {

	new := new(CrypData)
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	dd2 := document.Find("tbody")
	ss := dd2.Find("tr")
	aa := ss.Find("span")
	index := 0
	for _, app := range aa.Nodes {
		if index == 3 {
			index = 0
		}
		if index == 0 {
			re := regexp.MustCompile(`[[:space:]]`)
			new.Name = re.ReplaceAllString(app.FirstChild.Data, "")
			index++
		} else if index == 1 {
			re := regexp.MustCompile(`[[:space:]]`)
			new.Tag = re.ReplaceAllString(app.FirstChild.Data, "")
			new.Time = time.Now().Format(time.RFC850)
			indexTable := WriteTableCry(new, spreadSheetService, idSheets)

			if indexTable == 3 {
				break
			}
			index++
		} else {
			index++
		}

	}
}
