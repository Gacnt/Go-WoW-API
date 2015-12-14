package wowapi

import (
	"encoding/json"
	"log"
)

func (w *WowAPI) GetTalents(username, server string) CharacterTalents {
	// client := http.Client{}
	characterTalents := CharacterTalents{}
	url := server + "/" + username
	resp := w.Req(url, "&fields=talents")
	defer resp.Close()
	err := json.NewDecoder(resp).Decode(&characterTalents)
	if err != nil {
		log.Println(err)
	}
	return characterTalents
}
