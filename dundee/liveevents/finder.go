package liveevents

import (
	"errors"
	"fmt"
	"github.com/ign/ipl-dundee/dundee/streams"
)

func (this LiveEvents) Find(stream *streams.Stream) (*LiveEvent, error) {
	for i, _ := range this {
		liveEvent := &this[i]
		if liveEvent.Match(stream) {
			return liveEvent, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("%s%+v", "Unable to find a live event relating to the provided stream: ", stream))
}
