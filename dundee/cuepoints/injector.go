package cuepoints

import (
	"encoding/xml"
	"github.com/ign/ipl-dundee/dundee/elementals"
	"github.com/ign/ipl-dundee/dundee/liveevents"
	"path"
)

func Inject(liveEvent *liveevents.LiveEvent, cuePoint interface{}) error {

	body, err := xml.Marshal(cuePoint)
	if err != nil {
		return err
	}

	body = append([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>"), body...)

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
