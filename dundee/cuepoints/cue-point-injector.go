package cuepoints

import (
	"dundee/elemental"
	"encoding/xml"
	"path"
)

func Inject(eventPath string, elementalServer *elemental.ElementalServer, cuePoint interface{}) error {
	body, err := xml.Marshal(cuePoint)
	if err != nil {
		return err
	}

	req, err := elemental.GenerateRequest("POST", elementalServer, path.Join(eventPath, "stream_metadata"), body)
	if err != nil {
		return err
	}

	_, err = elemental.ExecuteRequest(req)
	if err != nil {
		return err
	}

	return nil
}
