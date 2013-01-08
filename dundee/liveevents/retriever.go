package liveevents

import (
	"dundee/elemental"
	"encoding/xml"
)

const liveEventsPath = "/live_events"

func Retrieve(elementalServers []elemental.ElementalServer) ([]Live_event_list, error) {
	var results = make([]Live_event_list, 0)

	for _, server := range elementalServers {
		data, err := retrieveData(&server)
		if err != nil {
			continue
		}

		var liveEventList Live_event_list

		err = xml.Unmarshal(data, &liveEventList)
		if err != nil {
			continue
		}

		liveEventList.Elemental = &server

		results = append(results, liveEventList)
	}

	return results, nil
}

func retrieveData(elementalServer *elemental.ElementalServer) ([]byte, error) {
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
