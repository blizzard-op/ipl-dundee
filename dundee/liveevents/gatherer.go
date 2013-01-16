package liveevents

import (
	"github.com/ign/ipl-dundee/dundee/elementals"
	"log"
)

func Gather(elementalServers []elementals.ElementalServer) LiveEvents {
	var results = make(LiveEvents, 0)
	var pipe = make(chan LiveEvents, len(elementalServers))

	for i, _ := range elementalServers {
		server := &elementalServers[i]
		go gatherFromServer(pipe, server)
	}

	for _, _ = range elementalServers {
		results = append(results, <-pipe...)
	}

	if len(results) == 0 {
		log.Println("Did not find any live events.")
	}

	return results
}

func gatherFromServer(pipe chan LiveEvents, server *elementals.ElementalServer) {
	var eventList LiveEvents

	defer func() {
		pipe <- eventList
	}()

	data, err := Fetch(server)
	if err != nil {
		return
	}

	eventList, err = Parse(data, server)
	if err != nil {
		return
	}
}
