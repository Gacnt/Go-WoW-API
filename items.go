package wowapi

import (
	"encoding/json"
	"log"
)

// GetItems returns the items the character is currently wearing
func (w *WowAPI) GetItems(username, server string) CharacterItems {
	// client := http.Client{}
	characterItems := CharacterItems{}
	url := server + "/" + username
	resp := w.Req(url, "&fields=items")
	defer resp.Close()
	err := json.NewDecoder(resp).Decode(&characterItems)
	if err != nil {
		log.Println(err)
	}
	return characterItems
}
