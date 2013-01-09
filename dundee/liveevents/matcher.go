package liveevents

import (
	"dundee/streams"
	"regexp"
)

func Match(stream *streams.Stream, liveEvent *Live_event) bool {
	found, _ := regexp.MatchString(stream.Name, liveEvent.Name)
	if found == true && liveEvent.Status == "running" {
		return true
	}
	return false
}
