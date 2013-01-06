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
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/injectcuepoint", injectCuePoint)

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

func injectCuePoint(w http.ResponseWriter, r *http.Request) {
	streamID, err := streams.ValidateID(r)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err)
		return
	}

	cuePoint, err := cuepoints.New(r)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err)
		return
	}

	streamSlice, err := streams.Retrieve(config)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, err)
		return
	}

	franchise, err := streams.Exists(streamID, streamSlice)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprint(w, err)
		return
	}

	//Beyond this point the client doesn't care - return 201
	w.WriteHeader(201)
	fmt.Fprint(w, franchise)

	liveEventResults, err := liveevents.Retrieve(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	eventID, elemental, err := liveevents.Find(franchise, liveEventResults)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = cuepoints.Inject(eventID, elemental, cuePoint)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "must-revalidate, no-cache, no-store")
	w.Header().Set("Connection", "close")
	fmt.Fprint(w, "pong")
}
