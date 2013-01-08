package streams

import (
	"errors"
	"net/http"
)

func ValidateStreamID(streamID string, streams []Stream) (string, error) {
	for _, value := range streams {
		if value.Id == streamID {
			return value.Franchise.Name, nil
		}
	}
	return "", errors.New("No stream with an ID matching \"" + streamID + "\" was found.")
}
