package liveevents

import (
	"encoding/xml"
	"github.com/ign/ipl-dundee/dundee/elementals"
)

func Parse(data []byte, server *elementals.ElementalServer) ([]Live_event, error) {
	var liveEventList Live_event_list

	err := xml.Unmarshal(data, &liveEventList)
	if err != nil {
		return nil, err
	}

	//Bind reference to server in each live event
	for i, _ := range liveEventList.Live_events {
		liveEventList.Live_events[i].Elemental = server
	}

	return liveEventList.Live_events, nil
}
