package liveevents

import (
	"dundee/streams"
	"errors"
	"fmt"
)

func Find(stream *streams.Stream, liveEvents []Live_event) (*Live_event, error) {

	fmt.Printf("%s%d%s", "Found ", len(liveEvents), " events.")

	for _, liveEvent := range liveEvents {
		if Match(stream, &liveEvent) {
			return &liveEvent, nil
		}
	}
	return nil, errors.New("Unable to find a live event relating to the provided stream.")
}
