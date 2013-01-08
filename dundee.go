package main

import (
	"dundee"
	"dundee/config"
	"fmt"
	"log"
	"net/http"
)

var conf = config.Get()

func main() {
	fmt.Println("Dundee starting on port " + conf.Port)

	http.HandleFunc("/ping", dundee.PingHandler)
	http.HandleFunc("/cuepoints", dundee.CuePointsHandler)

	err := http.ListenAndServe(conf.Port, nil)
	if err != nil {
		log.Fatalf("Failed to start Dundee on port "+conf.Port, err)
	}
}
