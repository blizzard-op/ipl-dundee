package liveevents

import (
	"dundee/elemental"
)

func Gather(elementalServers []elemental.ElementalServer) []Live_event {
	var results = make([]Live_event, 0)
	var pipe = make(chan []Live_event, len(elementalServers))

	for i, _ := range elementalServers {
		server := &elementalServers[i]
		go gatherFromServer(pipe, server)
	}

	for _, _ = range elementalServers {
		result := <-pipe
		if result != nil {
			results = append(results, result...)
		}
	}

	return results
}

func gatherFromServer(pipe chan []Live_event, server *elemental.ElementalServer) {
	var liveEvents []Live_event

	defer func() {
		pipe <- liveEvents
	}()

	data, err := Fetch(server)
	if err != nil {
		return
	}

	liveEvents, err = Parse(data, server)
	if err != nil {
		return
	}
}
