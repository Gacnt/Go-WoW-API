package wowapi

import (
	"io"
	"log"
	"net/http"
)

// WowAPI is the struct for the holder
type WowAPI struct {
	Locale string
	APIKey string
}

// NewAPI will create a new API helper
func NewAPI(apikey, locale string) *WowAPI {
	return &WowAPI{locale, apikey}
}

// Req will format a new request
func (w *WowAPI) Req(url, params string) io.ReadCloser {
	client := http.Client{}
	var requestURL string
	if params == "" {
		requestURL = "https://us.api.battle.net/wow/character/" + url + "?locale=" + w.Locale + "&apikey=" + w.APIKey
	} else {
		requestURL = "https://us.api.battle.net/wow/character/" + url + "?locale=" + w.Locale + "&apikey=" + w.APIKey + params
	}
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		log.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	body := resp.Body
	return body
}
