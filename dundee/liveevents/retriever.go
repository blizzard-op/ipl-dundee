package liveevents

import (
	"dundee"
	"dundee/elemental"
	"encoding/xml"
	"io/ioutil"
)

func Retrieve(config *dundee.Config) ([]LiveEvent, error) {
	var liveEvents = make([]LiveEvent, 3)

	for _, server := range config.Elementals {
		data, err := getXML(server, config.Live_events_path)
		if err != nil {
			continue
		}

		var tempLiveEvents []liveEvents
		err = xml.Unmarshal(data, &tempLiveEvents)
		if err != nil {
			continue
		}
		liveEvents.append(tempLiveEvents)
	}

	if len(liveEvents) == 0 {
		return nil, errors.New("No Live Events were found.")
	}

	return liveEvents, nil
}

func getXML(elementalServer elemental.ElementalServer, liveEventsPath string) ([]bytes, error) {
	req, err := elemental.GenerateRequest(elementalServer, liveEventsPath)
	if err != nil {
		return nil, err
	}

	resp, err = elemental.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
