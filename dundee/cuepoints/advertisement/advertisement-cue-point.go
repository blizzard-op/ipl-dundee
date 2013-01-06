package advertisement

import (
	"encoding/xml"
)

func init() {
	RegisterCuePointType("advertisement", New)
}

type cuePoint struct {
	XMLName  xml.Name `xml:"data"`
	Category string   `xml:"category"`
	Action   string   `xml:"action"`
	Label    string   `xml:"label"`
	Value    string   `xml:"value"`
}

func New(r *http.Request) (*cuePoint, error) {
	return &cuePoint{Category: "Video Cue Point", Action: "13", Label: "Ad Name", Value: "30"}, nil
}
