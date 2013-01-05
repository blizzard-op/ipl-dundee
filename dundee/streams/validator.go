package streams

import (
	"dundee/types"
	"errors"
	"net/http"
)

//ValidateID takes the "streamID" param from either the query string or POST body and
//checks if the string looks to be of an appropriate format. The "streamID" is returned
func ValidateID(r *http.Request) (string, error) {
	streamID := r.FormValue("streamid")
	if len(streamID) != 24 {
		return "", errors.New("The streamID provided was not 24 chars long.")
	}
	return streamID, nil
}

//Exists takes a streamID and an array of streams and searches
//for a stream matching the given streamID - returning the franchise
func Exists(streamID string, streams []types.Stream) (string, error) {
	for _, value := range streams {
		if value.Id == streamID {
			return value.Franchise.Name, nil
		}
	}
	return "", errors.New("No stream with an ID matching \"" + streamID + "\" was found.")
}
