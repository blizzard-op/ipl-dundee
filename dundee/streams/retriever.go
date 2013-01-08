package streams

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func RetrieveData(url string) ([]byte, error) {
	resp, err := http.Get(url)
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

func ProcessData(b []byte) ([]Stream, error) {
	var streams []Stream

	err := json.Unmarshal(b, &streams)
	if err != nil {
		return nil, err
	}

	return streams, nil
}
