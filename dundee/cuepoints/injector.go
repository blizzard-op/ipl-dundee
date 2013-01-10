package cuepoints

import (
	"encoding/xml"
	"github.com/ign/ipl-dundee/dundee/elementals"
	"github.com/ign/ipl-dundee/dundee/liveevents"
	"path"
)

func Inject(liveEvent *liveevents.Live_event, cuePoint interface{}) error {

	body, err := xml.Marshal(cuePoint)
	if err != nil {
		return err
	}

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
