package server

import (
	"fmt"
	"log"
	"net/http"
	"pudding-server/blockchain"
	"pudding-server/multidag"
	"strconv"
	"strings"
)

var theChain blockchain.ChainReader

func Server(reader blockchain.ChainReader) {
	theChain = reader
	http.HandleFunc("/", handler) // each request calls handler
	println("Listening on port 8000, server is now active (the wait is over)")
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

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 5 {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	rootString := parts[0]
	if rootString != "" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	verticesString := parts[1]
	if verticesString != "vertices" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	vertexType := parts[2]
	vertexNumberString := parts[3]
	filename := parts[4]
	vertexNumber, _ := strconv.ParseInt(vertexNumberString, 10, 64)

	var vertex multidag.Vertex

	if vertexType == "blockchain" && vertexNumber == 0 {
		vertex = theChain.GetBlockchainVertex()
	} else if vertexType == "block" {
		vertex = theChain.GetBlockVertex(vertexNumber)
	}

	if filename == "attributes.json" {
		byts, _ := vertex.GetAttributes().EncodeAsJson()
		fmt.Fprintf(w, "%s", byts)
	}

	if filename == "in.json" {
		byts, _ := vertex.GetInEndpoints().EncodeAsJson()
		fmt.Fprintf(w, "%s", byts)
	}

	if filename == "out.json" {
		byts, _ := vertex.GetOutEndpoints().EncodeAsJson()
		fmt.Fprintf(w, "%s", byts)
	}
	fmt.Print()
}
