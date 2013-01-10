package streams

import (
	"encoding/json"
)

func Parse(b []byte) (Streams, error) {
	var streamList Streams

	err := json.Unmarshal(b, &streamList)
	if err != nil {
		return nil, err
	}

	return streamList, nil
}
