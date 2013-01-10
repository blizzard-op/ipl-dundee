package cuepoints

import (
	"encoding/xml"
	"log"
	"net/http"
)

type inert_cuepoint struct {
	XMLName  xml.Name `xml:"data"`
	Category string   `xml:"category"`
	Action   string   `xml:"action"`
	Label    string   `xml:"label"`
	Value    string   `xml:"value"`
}

func init() {
	err := RegisterCuePointType("inert", func(r *http.Request) (interface{}, error) {
		return &inert_cuepoint{Category: "Video Cue Point", Action: "13", Label: "Ad Name", Value: "0"}, nil
	})
	if err != nil {
		log.Println("Failed to register cuepoint type.\n", err)
	}
}
