package main

import (
	"dundee"
	"dundee/cuepoints"
	"dundee/liveevents"
	"dundee/streams"
	"fmt"
	"log"
	"net/http"
)

var config *dundee.Config

func main() {
	config = getConfig()

	fmt.Println("Dundee starting on port " + config.Port)

	//Set Routes
	http.HandleFunc("/ping", pingHandler)
	http.HandleFunc("/cuepoints", cuePointsHandler)

	//Start server
	err := http.ListenAndServe(config.Port, nil)
	if err != nil {
		log.Fatalf("Failed to start Dundee on port "+config.Port, err)
	}
}

func getConfig() *dundee.Config {
	con, err := dundee.GetConfig()
	if err != nil {
		log.Fatalf("Dundee failed to read config. Aborting.\nReason: " + err.Error())
	}
	fmt.Println("Config loaded successfully.")
	return con
}

func cuePointsHandler(w http.ResponseWriter, r *http.Request) {

	streamID := r.FormValue("streamid")
	if streamID == nil {
		w.WriteHeader(400)
		fmt.Fprint(w, errors.New("You must include a valid streamid."))
	}

	cuePoint, err := cuepoints.New(r)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err)
		return
	}

	streamData, err := streams.RetrieveData(config.Streams_url)
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

		liveEventResults, err := liveevents.Retrieve(config.Elementals)
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

func PingHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "must-revalidate, no-cache, no-store")
	w.Header().Set("Connection", "close")
	fmt.Fprint(w, "pong")
}
