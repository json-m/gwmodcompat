package main

import (
	"encoding/json"
	"fmt"
	"gwmodcompat/pkg"
	"net/http"
	"time"
)

// handles starting up the local api server
func startHttpServer() {
	http.HandleFunc("/", siteIndex)
	http.HandleFunc("/data", getData)
	fs := http.FileServer(http.Dir("mods"))
	http.Handle("/mods/", http.StripPrefix("/mods", fs))

	// start server
	if err := http.ListenAndServe(fmt.Sprintf(":%d", 44444), nil); err != nil {
		go startHttpServer()
		time.Sleep(1 * time.Second)
	}
}

// just quick status output
func getData(w http.ResponseWriter, req *http.Request) {
	// marshal to actual json bytes
	h, err := json.Marshal(pkg.RetrieveData())
	if err != nil {
		w.WriteHeader(500)
		_, _ = fmt.Fprint(w, "problem in marshal health:", err)
		return
	}

	// set proper header and return the data
	w.Header().Set("Content-Type", "application/json")
	_, _ = fmt.Fprintf(w, string(h))
}

// main page
func siteIndex(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprint(w, "modcompat server running")
}
