package liveevents

import (
	"github.com/ign/ipl-dundee/dundee/elementals"
)

const liveEventsPath = "/live_events/"

func Fetch(elementalServer *elementals.ElementalServer) ([]byte, error) {
	req, err := elementalServer.GenerateRequest("GET", liveEventsPath, nil)
	if err != nil {
		return nil, err
	}

	_, body, err := elementals.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	return body, nil
}
