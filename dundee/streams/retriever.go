package streams

import (
	"dundee"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Retrieve(config *dundee.Config) ([]Stream, error) {
	var streams []Stream

	jsonBytes, err := getJSON(config)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonBytes, &streams)
	if err != nil {
		return nil, err
	}

	return streams, nil
}

func getJSON(config *dundee.Config) ([]byte, error) {
	resp, err := http.Get(config.Streams_url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
