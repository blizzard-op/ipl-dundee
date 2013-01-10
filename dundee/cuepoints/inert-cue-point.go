package cuepoints

import (
	"log"
	"net/http"
)

func init() {
	err := RegisterCuePointType("inert", func(r *http.Request) (interface{}, error) {
		return &advertisement_cuepoint{Category: "Video Cue Point", Action: "13", Label: "Ad Name", Value: "0"}, nil
	})
	if err != nil {
		log.Println("Failed to register cuepoint type.\n", err)
	}
}
