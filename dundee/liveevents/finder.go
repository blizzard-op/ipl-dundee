package liveevents

import (
	"dundee/elemental"
	"errors"
	"regexp"
)

func Find(franchise string, liveEventLists []Live_event_list) (string, *elemental.ElementalServer, error) {
	for _, liveEventList := range liveEventLists {
		for _, liveEvent := range liveEventList.Live_events {
			found, _ := regexp.MatchString(franchise, liveEvent.Name)
			if found == true && liveEvent.Status == "running" {
				return liveEvent.Path, liveEventList.Elemental, nil
			}
		}
	}
	return "", nil, errors.New("Unable to find a live event relating to the franchise \"" + franchise + "\"")
}
