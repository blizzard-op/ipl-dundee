package main

import (
	"dundee"
	"dundee/streams"
	"dundee/types"
	"fmt"
	"log"
	"net/http"
	"os"
)

var config *types.Config

func main() {
	//Get config_path
	var config_path string
	if len(os.Args) > 1 {
		config_path = os.Args[1]
	}

	//Get config
	var err error
	config, err = dundee.GetConfig(config_path)
	if err != nil {
		log.Fatalf("Dundee failed to read config at " + config_path + ". Aborting.\nReason: " + err.Error())
	}
	fmt.Println("Config loaded successfully. Dundee starting on port " + config.Port)

	//Set Routes
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/injectcuepoint", injectCuePoint)

	//Start server
	err = http.ListenAndServe(config.Port, nil)
	if err != nil {
		log.Fatalf("Failed to start Dundee on port "+config.Port, err)
	}
}

func injectCuePoint(w http.ResponseWriter, r *http.Request) {
	streamID, err := streams.ValidateID(r)
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
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "must-revalidate, no-cache, no-store")
	w.Header().Set("Connection", "close")
	fmt.Fprint(w, "pong")
}
