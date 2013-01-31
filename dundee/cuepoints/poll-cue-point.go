package cuepoints

import (
	"encoding/xml"
	"errors"
	"log"
	"net/http"
)

type poll_cuepoint struct {
	XMLName  xml.Name `xml:"data"`
	Category string   `xml:"category"`
	PollType string   `xml:"polltype"`
	PollId   string   `xml:"pollid"`
}

func init() {
	err := RegisterCuePointType("poll", func(r *http.Request) (interface{}, *TimedMetadata, error) {
		pollType := r.FormValue("poll-type")
		pollId := r.FormValue("poll-id")

		if pollType == "" || pollId == "" {
			return nil, nil, errors.New("A poll-type and poll-id must be included.")
		}

		cp := &poll_cuepoint{Category: "Poll Cue Point", PollType: pollType, PollId: pollId}
		tm := &TimedMetadata{Title: "poll" + pollType, Subtitle: pollId}

		return cp, tm, nil
	})
	if err != nil {
		log.Println("Failed to register cuepoint type.\n", err)
	}
}
