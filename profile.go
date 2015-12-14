package wowapi

import (
	"encoding/json"
	"log"
)

// GetProfile will get the players profile
func (w *WowAPI) GetProfile(username, server string) CharacterProfile {
	url := server + "/" + username
	resp := w.Req(url, "")
	defer resp.Close()

	character := CharacterProfile{}
	err := json.NewDecoder(resp).Decode(&character)
	if err != nil {
		log.Println(err)
	}
	return character
}
