package main

import (
	"pudding-server/jobs"
)

func main() {
	folder := "F:\\Data\\BuzzSpinCredit 10Feb 16years"

	err := jobs.ConstructTxoSpentTxi(folder)
	//err := jobs.ConstructTxoParentTrans(folder)
	//err := jobs.RunServer(folder)

	if err != nil {
		println(err.Error())
	}
}
