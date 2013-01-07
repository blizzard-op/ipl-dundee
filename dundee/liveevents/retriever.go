package liveevents

import (
	"dundee/elemental"
	"encoding/xml"
	"fmt"
)

const liveEventsPath = "/live_events"

func Retrieve(elementalServers []elemental.ElementalServer) ([]Live_event_list, error) {
	var results = make([]Live_event_list, 0)

	for _, server := range elementalServers {
		data, err := getXML(&server)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var liveEventList Live_event_list

		err = xml.Unmarshal(data, &liveEventList)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(len(liveEventList.Live_events))
		fmt.Println(liveEventList.Live_events)

		liveEventList.Elemental = &server

		results = append(results, liveEventList)
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
