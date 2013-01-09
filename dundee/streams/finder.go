package streams

import (
	"errors"
)

func Find(streamID string, streamList []Stream) (*Stream, error) {
	for i, _ := range streamList {
		stream := streamList[i]
		if stream.Id == streamID {
			return &stream, nil
		}
	}
	return nil, errors.New("No stream with an ID matching \"" + streamID + "\" was found.")
}
