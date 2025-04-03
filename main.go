package main

import (
	"pudding-server/jobs"
)

func main() {
	folder := "F:\\Data\\PerfectRemindSolution888888Blocks1Apr"

	//err := jobs.ConstructTxoSpentTxi(folder)
	//err := jobs.ConstructTxoParentTrans(folder)
	err := jobs.RunServer(folder)

	if err != nil {
		println(err.Error())
	}
}
