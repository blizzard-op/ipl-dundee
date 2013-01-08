package liveevents

import (
	"dundee/elemental"
	"dundee/streams"
	"errors"
	"regexp"
)

func Find(stream Stream, liveEventLists []Live_event_list) (string, *elemental.ElementalServer, error) {
	for _, liveEventList := range liveEventLists {
		for _, liveEvent := range liveEventList.Live_events {
			if Match(stream, liveEvent) {
				return liveEvent.Path, liveEventList.Elemental, nil
			}
		}
	}
	return "", nil, errors.New("Unable to find a live event relating to the franchise \"" + franchise + "\"")
}
