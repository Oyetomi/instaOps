package api

import (
	"github.com/Oyetomi/instaOps/internal/errors"
	"github.com/go-resty/resty/v2"
	"log"
	"strconv"
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
	resp, err := client.R().SetFormData(
		map[string]string{
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

func SetSettings(settings, sessionid string) string {
	resp, err := client.R().SetFormData(
		map[string]string{
			"settings":  settings,
			"sessionid": sessionid,
		}).Post("/auth/settings/set")
	if err != nil {
		log.Println(errors.ErrCouldNotSetSettings)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

func GetTimelineFeed(sessionid string) string {
	resp, err := client.R().SetQueryParams(
		map[string]string{
			"sessionid": sessionid,
		}).Get("/auth/timeline_feed")
	if err != nil {
		log.Println(errors.ErrCouldNotGetFeed)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}

func GetMediaID(media_pk int) string {
	resp, err := client.R().SetQueryParams(
		map[string]string{
			"media_pk": strconv.Itoa(media_pk),
		}).Get("/media/id")
	if err != nil {
		log.Println(errors.ErrCouldNotGetMediaID)
	}
	if resp.StatusCode() != 200 {
		return resp.String()
	}
	return resp.String()
}
