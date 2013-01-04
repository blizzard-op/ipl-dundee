package main

import (
	"dundee"
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8008"

func main() {

	fmt.Println("Dundee starting on " + PORT)

	http.HandleFunc("/ping", ping)
	http.HandleFunc("/injectcuepoint", injectCuePoint)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatalf("Failed to start Dundee on port "+PORT, err)
	}
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "must-revalidate, no-cache, no-store")
	w.Header().Set("Connection", "close")
	fmt.Fprint(w, "pong")
}

func injectCuePoint(w http.ResponseWriter, r *http.Request) {
	franchise, err := dundee.RetrieveFranchise(r)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, franchise)
}
