package cuepoints

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type cuePoint struct {
	XMLName  xml.Name `xml:"data"`
	Category string   `xml:"category"`
	Action   string   `xml:"action"`
	Label    string   `xml:"label"`
	Value    string   `xml:"value"`
}

func init() {
	err := RegisterCuePointType("advertisement", new)
	if err != nil {
		fmt.Println("Failed to register cuepoint type.\n" + err.Error())
	}
}

func new(r *http.Request) (interface{}, error) {
	return &cuePoint{Category: "Video Cue Point", Action: "13", Label: "Ad Name", Value: "30"}, nil
}
