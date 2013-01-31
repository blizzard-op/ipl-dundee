package cuepoints

import (
	"encoding/xml"
	"github.com/ign/ipl-dundee/dundee/elementals"
	"github.com/ign/ipl-dundee/dundee/liveevents"
	"log"
	"path"
)

func Inject(liveEvent *liveevents.LiveEvent, cp interface{}, tm *TimedMetadata) []error {

	pipe := make(chan error, 2)
	errors := make([]error, 0)

	go injectStreamMetadata(pipe, liveEvent, cp)
	go injectTimedMetadata(pipe, liveEvent, tm)

	for i := 0; i < 2; i++ {
		err := <-pipe
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

func injectStreamMetadata(pipe chan error, liveEvent *liveevents.LiveEvent, cp interface{}) {
	var err error

	defer func() {
		pipe <- err
	}()

	body, err := xml.Marshal(cp)
	if err != nil {
		return
	}

	body = append([]byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>"), body...)

	req, err := liveEvent.Elemental.GenerateRequest("POST", path.Join(liveEvent.Path, "stream_metadata"), body)
	if err != nil {
		return
	}

	_, _, err = elementals.ExecuteRequest(req)
	if err != nil {
		return
	}

	log.Println("Successfully injected cuepoint into event:", liveEvent.Name)

}

func injectTimedMetadata(pipe chan error, liveEvent *liveevents.LiveEvent, tm *TimedMetadata) {
	var err error

	defer func() {
		pipe <- err
	}()

	body := []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>" + "<timed_metadata>" + "<id3 encoding=\"base64\">" + generateId3(tm.Title, tm.Subtitle) + "</id3>" + "<cancel>false</cancel>" + "</timed_metadata>")

	req, err := liveEvent.Elemental.GenerateRequest("POST", path.Join(liveEvent.Path, "timed_metadata"), body)
	if err != nil {
		return
	}

	_, _, err = elementals.ExecuteRequest(req)
	if err != nil {
		return
	}

	log.Println("Successfully injected timed_metadata cuepoint into event:", liveEvent.Name)
}
