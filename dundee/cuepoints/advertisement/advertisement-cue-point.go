package advertisement

import (
	"dundee/cuepoints"
	"encoding/xml"
	"fmt"
	"net/http"
)

func init() {
	err := cuepoints.RegisterCuePointType("advertisement", New)
	if err != nil {
		fmt.Println("Failed to register cuepoint type.\n" + err.Error())
	}
}

type cuePoint struct {
	XMLName  xml.Name `xml:"data"`
	Category string   `xml:"category"`
	Action   string   `xml:"action"`
	Label    string   `xml:"label"`
	Value    string   `xml:"value"`
}

func New(r *http.Request) (interface{}, error) {
	return &cuePoint{Category: "Video Cue Point", Action: "13", Label: "Ad Name", Value: "30"}, nil
}
