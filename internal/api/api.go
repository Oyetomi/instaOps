package api

import (
	"github.com/go-resty/resty/v2"
	"log"
)

var client *resty.Client

func init() {
	client = resty.New()
	client.SetBaseURL("http://localhost:8000")
	client.SetContentLength(true)
}

func GetApiVersion() string {
	resp, err := client.R().Get("/version")
	if err != nil {
		log.Println("Error getting API version")
	}
	return resp.String()
}
