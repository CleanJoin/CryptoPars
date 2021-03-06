package internal

import (
	"log"
	"net/http"
	"regexp"
	"time"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/Iwark/spreadsheet.v2"
)

type GecData struct {
	Coin  string
	Money string
	Time  string
}

func GecParser(spreadSheetService *spreadsheet.Service, idSheets string) error {

	resp, err := http.Get("https://www.coingecko.com/en")
	if err != nil {
		log.Println(err)
		time.Sleep(time.Minute + 1)
		GecParser(spreadSheetService, idSheets)
	}

	new := new(GecData)
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}
	dd2 := document.Find("tbody")
	ss := dd2.Find("tr")
	aa := ss.Find("span")
	index := 0
	for _, app := range aa.Nodes {
		if index == 7 {
			index = 0
		}
		if index == 0 {
			re := regexp.MustCompile(`[[:space:]]`)
			new.Coin = re.ReplaceAllString(app.FirstChild.Data, "")
			index++
		} else if index == 1 {
			re := regexp.MustCompile(`[[:space:]]`)
			new.Money = re.ReplaceAllString(app.FirstChild.Data, "")
			new.Time = time.Now().Format(time.RFC850)
			indexTable := WriteTable(new, spreadSheetService, idSheets)
			if indexTable == 65 {
				break
			}
			index++
		} else {
			index++
		}

	}

	return nil
}
