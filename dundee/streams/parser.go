package streams

import (
	"encoding/json"
)

func Parse(b []byte) ([]Stream, error) {
	var streams []Stream

	err := json.Unmarshal(b, &streams)
	if err != nil {
		return nil, err
	}

	return streams, nil
}
