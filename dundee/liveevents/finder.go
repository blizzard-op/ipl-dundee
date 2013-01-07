package liveevents

import (
	"dundee/elemental"
	"errors"
	"regexp"
)

func Find(franchise string, liveEventResults []LiveEventResult) (string, *elemental.ElementalServer, error) {
	for _, liveEventResult := range liveEventResults {
		for _, liveEvent := range liveEventResult.LiveEvents {
			found, _ := regexp.MatchString(franchise, liveEvent.Title)
			if found == true {
				return liveEvent.Path, liveEventResult.Elemental, nil
			}
		}
	}
	return "", nil, errors.New("Unable to find a live event relating to the franchise \"" + franchise + "\"")
}
