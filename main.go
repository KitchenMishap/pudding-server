package main

import (
	"pudding-server/jobs"
)

func main() {
	folder := "F:/Data/858000AddressesCswParents"

	//err := jobs.ConstructTxoSpentTxi(folder)
	err := jobs.RunServer(folder)

	if err != nil {
		println(err.Error())
	}
}
