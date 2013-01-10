package liveevents

import (
	"encoding/xml"
	"github.com/ign/ipl-dundee/dundee/elemental"
)

type Live_event_list struct {
	XMLName     xml.Name     `xml:"live_event_list"`
	Live_events []Live_event `xml:"live_event"`
}

type Live_event struct {
	Path      string                     `xml:"href,attr"`
	Name      string                     `xml:"name"`
	Status    string                     `xml:"status"`
	Elemental *elemental.ElementalServer `-`
}
