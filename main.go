package emdb

import (
	"embed"
	"encoding/json"
)

//go:embed speakers.json
var speakersJson []byte

//go:embed partners.json
var partnersJson []byte

//go:embed img
var imgs embed.FS

type speaker struct {
	Name []string `json:"name"`
	Img  []string `json:"img"`
	Desc []string `json:"desc"`
}

type partner struct {
	Name []string `json:"name"`
	Img  []string `json:"img"`
	Url  []string `json:"url"`
}

var (
	speakers map[string]speaker
	partners map[string]partner
)

func init() {
	speakers = make(map[string]speaker)
	partners = make(map[string]partner)
	json.Unmarshal(speakersJson, &speakers)
	json.Unmarshal(partnersJson, &partners)
}
