package jobs

import (
	"pudding-server/blockchain"
	"pudding-server/derived"
	"pudding-server/server"
)

func RunServer(folder string) error {
	println("Please wait... opening files")
	reader := blockchain.NewChainReader(folder)
	df, err := derived.NewDerivedFiles(folder)
	if err != nil {
		return err
	}

	err = df.OpenReadOnly()
	if err != nil {
		return err
	}

	server.Server(reader, df)
	err = df.Close()
	if err != nil {
		return err
	}
	return nil
}

func ConstructTxoSpentTxi(folder string) error {
	println("Please wait... Constructing TxoSpentTxi")
	derivedFiles, err := derived.NewDerivedFiles(folder)
	if err != nil {
		return err
	}

	err = derived.ConstructTxoSpentTxi(derivedFiles)
	if err != nil {
		return err
	}
	println("...Done")
	return nil
}

func ConstructTxoParentTrans(folder string) error {
	println("Please wait... Constructing TxoParentTrans")
	derivedFiles, err := derived.NewDerivedFiles(folder)
	if err != nil {
		return err
	}

	err = derived.ConstructTxoParentTrans(derivedFiles)
	if err != nil {
		return err
	}
	println("...Done")
	return nil
}
