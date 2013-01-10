package liveevents

import (
	"encoding/xml"
	"github.com/ign/ipl-dundee/dundee/elementals"
)

func Parse(data []byte, server *elementals.ElementalServer) (LiveEvents, error) {
	var liveEventList live_event_list

	err := xml.Unmarshal(data, &liveEventList)
	if err != nil {
		return nil, err
	}

	//Bind reference to server in each live event
	for i, _ := range liveEventList.LiveEvents {
		liveEventList.LiveEvents[i].Elemental = server
	}

	return liveEventList.LiveEvents, nil
}
