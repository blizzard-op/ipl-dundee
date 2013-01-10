package liveevents

import (
	"github.com/ign/ipl-dundee/dundee/elementals"
)

type LiveEvent struct {
	Path      string                      `xml:"href,attr"`
	Name      string                      `xml:"name"`
	Status    string                      `xml:"status"`
	Elemental *elementals.ElementalServer `-`
}

type LiveEvents []LiveEvent

//Used for unmarshalling XML
type live_event_list struct {
	LiveEvents LiveEvents `xml:"live_event"`
}
