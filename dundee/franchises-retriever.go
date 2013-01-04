package dundee

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const STREAMS_URL = "http://esports.ign.com/content/v1/streams.json"

type stream struct {
	Id        string
	Franchise map[string]string
}

func RetrieveFranchise(request *http.Request) (string, error) {
	var streams []stream

	streamID, err := verifyPOSTData(request)
	if err != nil {
		return "", err
	}

	jsonBytes, err1 := getJSON()
	if err1 != nil {
		return "", err1
	}

	err2 := unmarshalJSON(jsonBytes, &streams)
	if err2 != nil {
		return "", err2
	}

	for i := 0; i < len(streams); i++ {
		if streams[i].Id == streamID {
			return streams[i].Franchise["name"], nil
		}
	}

	return "", errors.New("Did not find stream corresponding to given streamid.")
}

func unmarshalJSON(jsonBytes []byte, streams *[]stream) error {
	if err := json.Unmarshal(jsonBytes, streams); err != nil {
		return err
	}
	return nil
}

func verifyPOSTData(r *http.Request) (string, error) {
	streamid := r.FormValue("streamid")
	var err error
	if len(streamid) != 24 {
		err = errors.New("Invalid streamid.")
	}
	return streamid, err
}

func getJSON() ([]byte, error) {
	resp, err := http.Get(STREAMS_URL)
	if err != nil {
		return []byte(""), err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return body, nil
}
