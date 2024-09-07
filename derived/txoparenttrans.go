package derived

import "github.com/KitchenMishap/pudding-shed/wordfile"

func ConstructTxoParentTrans(derived *DerivedFiles) error {
	println("Please wait, calculating parents file...")
	// Here are the files we wish to directly read from
	var txoSatsWordFile wordfile.ReadAtWordCounter = derived.privilegedFiles.TxoSatsFile()
	var transFirstTxoWordFile wordfile.ReadAtWordCounter = derived.privilegedFiles.TransFirstTxoFile()
	var txiTxWordFile = derived.privilegedFiles.TxiTxFile()

	// How many txos?
	numTxos, err := txoSatsWordFile.CountWords()
	if err != nil {
		return err
	}

	// Want to use the same bytes per word as something that refers to trans
	var wordSize int64 = txiTxWordFile.WordSize()

	// Create file & open it
	fileCreator := wordfile.NewConcreteWordFileCreator("txoparenttrans", derived.folder+"/derived", wordSize)
	err = fileCreator.CreateWordFile()
	if err != nil {
		return err
	}
	txoParentTransFile, err := fileCreator.OpenWordFile()
	if err != nil {
		return err
	}

	// Go through each trans
	for trans := int64(0); true; trans++ {
		firstTxo, err := transFirstTxoWordFile.ReadWordAt(trans)
		if err != nil {
			return err
		}
		lastTxoPlusOne, err := transFirstTxoWordFile.ReadWordAt(trans + 1)
		if err != nil {
			// Probably past the end of the file
			lastTxoPlusOne = numTxos
		}
		for txo := firstTxo; txo < lastTxoPlusOne; txo++ {
			txoParentTransFile.WriteWordAt(trans, txo)
		}
	}
	err = txoParentTransFile.Close()
	if err != nil {
		return err
	}
	txoSatsWordFile.Close()
	txiTxWordFile.Close()
	transFirstTxoWordFile.Close()
	println("...Done")
	return nil
}
