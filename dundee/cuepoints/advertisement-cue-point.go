package cuepoints

import (
	"encoding/xml"
	"log"
	"net/http"
)

type advertisement_cuepoint struct {
	XMLName  xml.Name `xml:"data"`
	Category string   `xml:"category"`
	Action   string   `xml:"action"`
	Label    string   `xml:"label"`
	Value    string   `xml:"value"`
}

func init() {
	err := RegisterCuePointType("advertisement", func(r *http.Request) (interface{}, error) {
		return &advertisement_cuepoint{Category: "Video Cue Point", Action: "13", Label: "Ad Name", Value: "30"}, nil
	})
	if err != nil {
		log.Println("Failed to register cuepoint type.\n", err)
	}
}
