package main

import (
	"flag"
	"fmt"
	"github.com/ign/ipl-dundee/dundee"
	"log"
	"net/http"
)

const (
	defaultStreamsURL = "http://esports.ign.com/content/v1/streams.json"
	routePrefix       = "/dundee/v1"
)

var config dundee.Config
var port int

func init() {
	var configPath string
	var streamsURL string

	//Read flags
	flag.StringVar(&configPath, "config-path", "config.json", "Config filepath. Default is ./config.json")
	flag.StringVar(&streamsURL, "streams-url", defaultStreamsURL, "Streams Endpoint URL. Default is "+defaultStreamsURL)
	flag.IntVar(&port, "port", 80, "Port number. Default is 80")
	flag.Parse()

	config = dundee.Config{Streams_url: &streamsURL}
	err := config.Parse(&configPath)
	if err != nil {
		log.Fatalln("Failed to load config at:", configPath, "\nReason:", err)
	}
}

func main() {
	log.Println("Dundee started on port:", port)

	http.HandleFunc(routePrefix+"/ping", dundee.PingHandler)
	http.HandleFunc(routePrefix+"/cuepoints", delegateTo(dundee.CuePointsHandler))

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalln("Failed to start Dundee on port:", port, "\nReason:", err)
	}
}

func delegateTo(handler func(http.ResponseWriter, *http.Request, *dundee.Config)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, &config)
	}
}
