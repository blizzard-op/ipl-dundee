package liveevents

import (
	"encoding/xml"
)

type LiveEvent struct {
	XMLName xml.Name `xml:"live_event"`
	Path string `xml:"href,attr"`
}