package liveevents

import (
	"encoding/xml"
)

type LiveEvent struct {
	XMLName xml.Name `xml:live_event`
	Path    string   `xml:"href,attr"`
	User_data
}

type User_data struct {
	Franchise string `xml:franchise`
}
