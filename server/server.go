package server

import (
	"fmt"
	"github.com/KitchenMishap/pudding-server/blockchain"
	"github.com/KitchenMishap/pudding-server/derived"
	"github.com/KitchenMishap/pudding-server/multidag"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var theChain *blockchain.ChainReader
var theDerived *derived.DerivedFiles

func Server(reader *blockchain.ChainReader, df *derived.DerivedFiles) {
	theChain = reader
	theDerived = df
	http.HandleFunc("/", handler) // each request calls handler
	// 14699 is a random port chosen by random.org as being between 1024 and 65535
	println("Listening on port 14699, server is now active (the wait is over)")
	// 0.0.0.0 allows any IP address to connect to us
	log.Fatal(http.ListenAndServe("0.0.0.0:14699", nil))
}

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
	if len(parts) < 2 {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	rootString := parts[0]
	if rootString != "" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	verticesString := parts[1]
	if verticesString == "fulldag-vertex" {
		handleFulldagVertex(w, r, parts)
	} else if verticesString == "minidag-vertex" {
		handleFulldagVertex(w, r, parts)
	} else if verticesString == "lookup" {
		handleLookups(w, r, parts)
	} else {
		http.Error(w, "404 not found.", http.StatusNotFound)
	}
	return
}
func handleFulldagVertex(w http.ResponseWriter, r *http.Request, parts []string) {
	vertexType := parts[2]
	vertexNumberString := parts[3]
	filename := parts[4]
	vertexNumber, _ := strconv.ParseInt(vertexNumberString, 10, 64)

	var vertex multidag.Vertex

	if vertexType == "blockchain" && vertexNumber == 0 {
		vertex = theChain.GetBlockchainVertex()
	} else if vertexType == "block" {
		vertex = theChain.GetBlockVertex(vertexNumber)
	} else if vertexType == "transaction" {
		vertex = theChain.GetTransactionVertex(vertexNumber)
	} else if vertexType == "txo" {
		vertex = theChain.GetTxoVertex(vertexNumber, theDerived)
	} else if vertexType == "address" {
		vertex = theChain.GetAddressVertex(vertexNumber)
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
}

func handleLookups(w http.ResponseWriter, r *http.Request, parts []string) {
	hashOrAddress := parts[2]
	// Is this string a known address, block hash, or transaction hash?
	addressHeight := theChain.LookupAddress(hashOrAddress)
	if addressHeight != -1 {
		fmt.Fprintf(w, "{\"partialUrl\":\"address/%d\"}", addressHeight)
	} else {
		transHeight := theChain.LookupTransaction(hashOrAddress)
		if transHeight != -1 {
			fmt.Fprintf(w, "{\"partialUrl\":\"transaction/%d\"}", transHeight)
		} else {
			blockHeight := theChain.LookupBlock(hashOrAddress)
			if blockHeight != -1 {
				fmt.Fprintf(w, "{\"partialUrl\":\"block/%d\"}", blockHeight)
			} else {
				http.Error(w, "404 not found.", http.StatusNotFound)
			}
		}
	}
}
