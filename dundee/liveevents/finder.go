package liveevents

import (
	"dundee/streams"
	"errors"
)

func Find(stream *streams.Stream, liveEvents []Live_event) (*Live_event, error) {
	for i, _ := range liveEvents {
		liveEvent := &liveEvents[i]
		if Match(stream, liveEvent) {
			return liveEvent, nil
		}
	}
	return nil, errors.New("Unable to find a live event relating to the provided stream.")
}
