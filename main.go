package main

import (
	"pudding-server/jobs"
)

func main() {
	folder := "E:\\Data\\GaugeFlyDivorce888888_5digitsNewParams"

	err := jobs.ConstructTxoSpentTxi(folder)
	//err := jobs.ConstructTxoParentTrans(folder)
	//err := jobs.RunServer(folder)

	if err != nil {
		println(err.Error())
	}
}
