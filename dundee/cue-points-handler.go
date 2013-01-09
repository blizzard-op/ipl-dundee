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

	streamID, cuePointType, err := fetchParams(r)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err)
		return
	}

	cuePoint, err := cuepoints.New(cuePointType, r)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err)
		return
	}

	stream, err := resolveStream(streamID, w)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	//Beyond this point the client doesn't care - return 201
	w.WriteHeader(201)
	fmt.Fprint(w, stream.Franchise.Name)

	go injectCuePoint(stream, cuePoint)
}

func injectCuePoint(stream *streams.Stream, cuePoint interface{}) {
	liveEvents := liveevents.Gather(conf.Elementals)

	liveEvent, err := liveevents.Find(stream, liveEvents)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = cuepoints.Inject(liveEvent, cuePoint)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func resolveStream(streamID string, w http.ResponseWriter) (*streams.Stream, error) {
	streamData, err := streams.Fetch(conf.Streams_url)
	if err != nil {
		w.WriteHeader(500)
		return nil, err
	}

	streamList, err := streams.Parse(streamData)
	if err != nil {
		w.WriteHeader(500)
		return nil, err
	}

	stream, err := streams.Find(streamID, streamList)
	if err != nil {
		w.WriteHeader(400)
		return nil, err
	}

	return stream, nil
}

func fetchParams(r *http.Request) (string, string, error) {
	streamID := r.FormValue("streamid")
	if streamID == "" {
		return "", "", errors.New("A streamid must be included.")
	}

	cuePointType := r.FormValue("cue-point-type")
	if cuePointType == "" {
		return "", "", errors.New("A cue-point-type must be included.")
	}

	return streamID, cuePointType, nil
}
