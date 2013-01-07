package liveevents

import (
	"dundee"
	"dundee/elemental"
	"encoding/xml"
	"io/ioutil"
)

func Retrieve(config *dundee.Config) ([]LiveEventResult, error) {
	var results = make([]LiveEventResult, 3)

	for _, server := range config.Elementals {
		data, err := getXML(server, config.Live_events_path)
		if err != nil {
			continue
		}

		var tempLiveEvents []liveEvents
		var result = LiveEventResult{elemental: &server, liveEvents: &tempLiveEvents}

		err = xml.Unmarshal(data, &tempLiveEvents)
		if err != nil {
			continue
		}

		results.append(result)
	}

	return results, nil
}

func getXML(elementalServer elemental.ElementalServer, liveEventsPath string) ([]bytes, error) {
	req, err := elemental.GenerateRequest("GET", elementalServer, liveEventsPath)
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
