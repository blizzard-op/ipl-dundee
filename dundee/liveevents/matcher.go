package liveevents

import (
	"dundee/streams"
)

func Match(stream *Stream, liveEvent Live_event) bool {
	found, _ := regexp.MatchString(stream.Franchise, liveEvent.Name)
	if found == true && liveEvent.Status == "running" {
		return true
	}
	return false
}
