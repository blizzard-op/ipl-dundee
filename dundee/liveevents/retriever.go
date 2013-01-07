package liveevents

import (
	"dundee/elemental"
	"encoding/xml"
)

const liveEventsPath = "/live_events"

func Retrieve(elementalServers []elemental.ElementalServer) ([]LiveEventResult, error) {
	var results = make([]LiveEventResult, 0)

	for _, server := range elementalServers {
		data, err := getXML(&server)
		if err != nil {
			continue
		}

		var tempLiveEvents []LiveEvent
		var result = LiveEventResult{Elemental: &server, LiveEvents: tempLiveEvents}

		err = xml.Unmarshal(data, &tempLiveEvents)
		if err != nil {
			continue
		}

		results = append(results, result)
	}

	return results, nil
}

func getXML(elementalServer *elemental.ElementalServer) ([]byte, error) {
	req, err := elemental.GenerateRequest("GET", elementalServer, liveEventsPath, nil)
	if err != nil {
		return nil, err
	}

	_, body, err := elemental.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	return body, nil
}
