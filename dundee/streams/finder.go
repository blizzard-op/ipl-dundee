package streams

import (
	"dundee/config"
	"errors"
	"net/http"
)

var conf = config.Get()

func Find(streamID string, w http.ResponseWriter) (*Stream, error) {
	streamData, err := RetrieveData(conf.Streams_url)
	if err != nil {
		w.WriteHeader(500)
		return nil, err
	}

	streamList, err := ProcessData(streamData)
	if err != nil {
		w.WriteHeader(500)
		return nil, err
	}

	for _, stream := range streamList {
		if stream.Id == streamID {
			return &stream, nil
		}
	}

	w.WriteHeader(400)
	return nil, errors.New("No stream with an ID matching \"" + streamID + "\" was found.")
}
