package wowapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

// WowAPI is the struct for the holder
type WowAPI struct {
	Locale string
	APIKey string
}

// CharacterProfile is the struct that will contain the json RESP from
// the API request
type CharacterProfile struct {
	LastModified      int    `json:"lastModified"`
	Name              string `json:"name"`
	Realm             string `json:"realm"`
	Battlegroup       string `json:"battlegroup"`
	Class             int    `json:"class"`
	Race              int    `json:"race"`
	Gender            int    `json:"gender"`
	Level             int    `json:"level"`
	AchievementPoints int    `json:"achievementPoints"`
	Thumbnail         string `json:"thumbnail"`
	CalcClass         string `json:"calcClass"`
	Faction           int    `json:"faction"`
	TotalHK           int    `json:"totalHonorableKills"`
}

type CharacterTalents struct {
	CharacterProfile
	Talents []Talent `json:"talents"`
}

type Talent struct {
	Selected   bool          `json:"selected,omitempty"`
	Talents    []TalentInner `json:"talents"`
	Glyphs     Glyphs        `json:"glyphs"`
	Spec       Spec          `json:"spec"`
	CalcTalent string        `json:"calcTalent"`
	CalcSpec   string        `json:"calcSpec"`
	CalcGlyph  string        `json:"calcGlyph"`
}

type Spec struct {
	Name            string `json:"name"`
	Role            string `json:"role"`
	BackgroundImage string `json:"backgroundImage"`
	Icon            string `json:"icon"`
	Description     string `json:"description"`
	Order           int    `json:"order"`
}

type Glyphs struct {
	Major []Glyph `json:"major"`
	Minor []Glyph `json:"minor"`
}

type Glyph struct {
	Glyph int    `json:"glyph"`
	Item  int    `json:"item"`
	Name  string `json:"name"`
	Icon  string `json:"icon"`
}

type TalentInner struct {
	Tier   int   `json:"tier"`
	Column int   `json:"tier"`
	Spell  Spell `json:"spell"`
}

type Spell struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
	CastTime    string `json:"castTime"`
	Cooldown    string `json:"cooldown"`
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
