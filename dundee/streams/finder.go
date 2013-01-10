package streams

import (
	"errors"
)

func (this Streams) Find(streamID string) (*Stream, error) {
	for i, _ := range this {
		stream := this[i]
		if stream.Id == streamID {
			return &stream, nil
		}
	}
	return nil, errors.New("No stream with an ID matching \"" + streamID + "\" was found.")
}
