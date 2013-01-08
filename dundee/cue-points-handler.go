package dundee

import (
	"dundee/config"
	"dundee/cuepoints"
	"dundee/liveevents"
	"dundee/streams"
	"errors"
	"fmt"
	"net/http"
)

var conf = config.Get()

func CuePointsHandler(w http.ResponseWriter, r *http.Request) {

	streamID := r.FormValue("streamid")
	if streamID == "" {
		w.WriteHeader(400)
		fmt.Fprint(w, errors.New("You must include a valid streamid."))
		return
	}

	cuePoint, err := cuepoints.New(r)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err)
		return
	}

	streamData, err := streams.RetrieveData(conf.Streams_url)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	streamList, err := streams.ProcessData(streamData)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
	}

	franchise, err := streams.ValidateStreamID(streamID, streamList)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err)
		return
	}

	//Beyond this point the client doesn't care - return 201
	w.WriteHeader(201)
	fmt.Fprint(w, franchise)

	go func() {

		liveEventResults, err := liveevents.Retrieve(conf.Elementals)
		if err != nil {
			fmt.Println(err)
			return
		}

		eventPath, elemental, err := liveevents.Find(franchise, liveEventResults)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = cuepoints.Inject(eventPath, elemental, cuePoint)
		if err != nil {
			fmt.Println(err)
			return
		}

	}()
}
