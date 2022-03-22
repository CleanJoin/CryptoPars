package internal

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

func CheckError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}

func Connect() (*spreadsheet.Service, *http.Client) {
	data, err := ioutil.ReadFile("./src/cred.json")
	CheckError(err)
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	CheckError(err)
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)
	return service, client
}
