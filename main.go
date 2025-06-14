package main

import (
	"github.com/KitchenMishap/pudding-server/jobs"
)

func main() {
	//folder := "E:\\Data\\FleeSwallowImmune888888CswHashesDeleted"
	folder := "E:\\Data\\FleeSI_ReadOnly"

	//err := jobs.ConstructTxoSpentTxi(folder)
	//err := jobs.ConstructTxoParentTrans(folder)
	err := jobs.RunServer(folder)

	if err != nil {
		println(err.Error())
	}
}
