package main

import (
	"pudding-server/blockchain"
	"pudding-server/server"
)

func main() {
	println("Please wait... opening files")
	reader := blockchain.NewChainReader("G:/Data/858000AddressesCswParents")
	server.Server(reader)
}
