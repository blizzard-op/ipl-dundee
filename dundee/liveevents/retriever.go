package liveevents

import (
	"dundee/elemental"
	"encoding/xml"
)

const liveEventsPath = "/live_events"

func Retrieve(elementalServers []elemental.ElementalServer) []Live_event {
	var results = make([]Live_event, 0)
	var pipe = make(chan []Live_event, len(elementalServers))

	for i, _ := range elementalServers {
		server := &elementalServers[i]
		go retrieveServerEvents(pipe, server)
	}

	for _, _ = range elementalServers {
		result := <-pipe
		if result != nil {
			results = append(results, result...)
		}
	}

	return results
}

func retrieveServerEvents(pipe chan []Live_event, server *elemental.ElementalServer) {
	var liveEventList Live_event_list
	defer func() {
		pipe <- liveEventList.Live_events
	}()

	data, err := retrieveServerData(server)
	if err != nil {
		return
	}

	err = xml.Unmarshal(data, &liveEventList)
	if err != nil {
		return
	}

	//Bind reference to server in each live event
	for i, _ := range liveEventList.Live_events {
		liveEventList.Live_events[i].Elemental = server
	}
}

func retrieveServerData(elementalServer *elemental.ElementalServer) ([]byte, error) {
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
