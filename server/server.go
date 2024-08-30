package server

import (
	"fmt"
	"log"
	"net/http"
	"pudding-server/multidag"
	"strconv"
	"strings"
)

func Server() {
	http.HandleFunc("/", handler) // each request calls handler
	println("Listening on port 8000, ctrl-c to end")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	// Get the request origin
	origin := r.Header.Get("Origin")
	// The following tells the browser to allow requests from 127.0.0.1
	// This helps with CORS restrictions - Cross-Origin Resource Sharing
	if origin == "http://127.0.0.1" {
		w.Header().Add("Access-Control-Allow-Origin", origin)
	}
	if origin == "http://localhost" {
		w.Header().Add("Access-Control-Allow-Origin", origin)
	}

	parts := strings.Split(r.URL.Path+"//", "/")
	vertexType := parts[1]
	vertexNumberString := parts[2]
	filename := parts[3]
	vertexNumber, _ := strconv.ParseInt(vertexNumberString, 10, 64)

	vertex := multidag.NewConcreteVertex()

	if vertexType == "blockchain" && vertexNumber == 0 {
		vertex.AddAttribute("technology", "Bitcoin")
	}

	if filename == "attributes.json" {
		byts, _ := vertex.GetAttributes().EncodeAsJson()
		fmt.Fprintf(w, "%s", byts)
	}
}
