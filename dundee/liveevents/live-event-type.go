package liveevents

import (
	"dundee/elemental"
	"encoding/xml"
)

type Live_event_list struct {
	XMLName     xml.Name                   `xml:"live_event_list"`
	Live_events []Live_event               `xml:"live_event"`
	Elemental   *elemental.ElementalServer `-`
}

type Live_event struct {
	Path   string `xml:"href,attr"`
	Name   string `xml:"name"`
	Status string `xml:"status"`
}
