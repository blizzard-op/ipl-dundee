package cuepoints

import (
	"encoding/xml"
	"github.com/ign/ipl-dundee/dundee/elementals"
	"github.com/ign/ipl-dundee/dundee/liveevents"
	"path"
)

func Inject(liveEvent *liveevents.LiveEvent, cuePoint interface{}) error {

	body, err := xml.MarshalIndent(cuePoint, "", "    ")
	if err != nil {
		return err
	}

	body = append([]byte(xml.Header), body...)

	req, err := liveEvent.Elemental.GenerateRequest("POST", path.Join(liveEvent.Path, "stream_metadata"), body)
	if err != nil {
		return err
	}

	_, _, err = elementals.ExecuteRequest(req)
	if err != nil {
		return err
	}

	return nil
}
