package liveevents

import (
	"dundee/streams"
	"errors"
	"fmt"
)

func Find(stream *streams.Stream, liveEvents []Live_event) (*Live_event, error) {
	for i, _ := range liveEvents {
		liveEvent := &liveEvents[i]
		if Match(stream, liveEvent) {
			return liveEvent, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("%s%+v", "Unable to find a live event relating to the provided stream: ", stream))
}
