package cuepoints

import (
	"dundee/elemental"
	"path"
)

func Inject(eventPath string, elementalServer elemental.ElementalServer, cuePoint interface{}) error {
	body, err := xml.Marshal(cuePoint)
	if err != nil {
		return err
	}

	path := path.Join(elementalServer.URL, eventPath, "stream_metadata")

	req, err := elemental.GenerateRequest("POST", elementalServer, path, body)
	if err != nil {
		return err
	}

	resp, err = elemental.ExecuteRequest(req)
	if err != nil {
		return err
	}

	return nil
}
