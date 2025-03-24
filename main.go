package main

import (
	"pudding-server/jobs"
)

func main() {
	folder := "F:\\Data\\TwoYear"

	//err := jobs.ConstructTxoSpentTxi(folder)
	//err := jobs.ConstructTxoParentTrans(folder)
	err := jobs.RunServer(folder)

	if err != nil {
		println(err.Error())
	}
}
