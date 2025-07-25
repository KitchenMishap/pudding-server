package main

import (
	"flag"
	"github.com/KitchenMishap/pudding-server/jobs"
)

func main() {
	var sDirFlag = flag.String("Dir", "", "Directory to serve data from")
	var bConstructFlag = flag.Bool("Construct", false, "Construct Txo Spent Txi data")
	flag.Parse()

	//folder := "/mnt/FleeSI"

	var err error
	if *bConstructFlag {
		err = jobs.ConstructTxoSpentTxi(*sDirFlag)
	} else {
		err = jobs.RunServer(*sDirFlag)
	}

	//err := jobs.ConstructTxoParentTrans(folder)

	if err != nil {
		println(err.Error())
	}
}
