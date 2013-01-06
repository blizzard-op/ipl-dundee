package cuepoints

import (
	"dundee/elemental"
)

func Inject(eventID string, elementalServer elemental.ElementalServer, cuePoint interface{}) error {
	body, err := xml.Marshal(cuePoint)
	if err != nil {
		return err
	}

	path := elementalServer.URL + "/live_events/" + eventID + "/stream_metadata"

	req, err := elemental.GenerateRequest(elementalServer, path, body)
	if err != nil {
		return err
	}

	resp, err = elemental.ExecuteRequest(req)
	if err != nil {
		return err
	}

	return nil
}
