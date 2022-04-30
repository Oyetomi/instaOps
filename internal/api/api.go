package api

import (
	"github.com/Oyetomi/instaOps/internal/errors"
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
		log.Println(errors.ErrCouldNotGetAPIVersion)
	}
	return resp.String()
}

func Login(username, password string) string {
	resp, err := client.R().SetFormData(map[string]string{
		"username": username,
		"password": password,
	}).Post("/auth/login")
	if err != nil {
		log.Println(errors.ErrLoginFailed)
	}
	if resp.StatusCode() != 200 {
		log.Println(resp.String())
	}
	return resp.String()
}

func GetSettings(sessionid string) string {
	resp, err := client.R().SetQueryParams(
		map[string]string{
			"sessionid": sessionid,
		}).Get("/auth/settings/get")
	if err != nil {
		log.Println(errors.ErrCouldNotGetSettings)
	}
	if resp.StatusCode() != 200 {
		log.Println(resp.String())
	}
	return resp.String()
}
