package derived

import (
	"github.com/KitchenMishap/pudding-shed/wordfile"
)

func ConstructTxoSpentTxi(derived *DerivedFiles) error {
	// Here are the files we wish to directly read from
	var txoSatsWordFile wordfile.ReadAtWordCounter = derived.privilegedFiles.TxoSatsFile()
	var txiTxWordFile wordfile.ReadAtWordCounter = derived.privilegedFiles.TxiTxFile()
	var txiVoutWordFile wordfile.ReadAtWordCounter = derived.privilegedFiles.TxiVoutFile()
	var transFirstTxoWordFile wordfile.ReadAtWordCounter = derived.privilegedFiles.TransFirstTxoFile()

	// How many txos?
	numTxos, err := txoSatsWordFile.CountWords()
	if err != nil {
		return err
	}

	// Want to use the same bytes per word as something that refers to txis
	var transFirstTxiWordFile wordfile.ReadAtWordCounter = derived.privilegedFiles.TransFirstTxiFile()
	var wordSize int64 = transFirstTxiWordFile.WordSize()

	// Create file right size & open it
	fileCreator := wordfile.NewConcreteWordFileCreator("txospenttxi", derived.folder+"/derived", wordSize)
	err = fileCreator.CreateWordFileFilledZeros(numTxos)
	if err != nil {
		return err
	}
	txoSpentTxiWordFile, err := fileCreator.OpenWordFile()
	if err != nil {
		return err
	}

	// Go through each txi
	numTxis, err := txiTxWordFile.CountWords()
	if err != nil {
		return err
	}
	for txi := int64(0); txi < numTxis; txi++ {
		// Get the txi's transaction and txi's vout
		trans, err := txiTxWordFile.ReadWordAt(txi)
		if err != nil {
			return err
		}
		vout, err := txiVoutWordFile.ReadWordAt(txi)
		if err != nil {
			return err
		}

		// Get the first txo of the trans
		firstTxo, err := transFirstTxoWordFile.ReadWordAt(trans)
		if err != nil {
			return err
		}

		// Find the txo
		txo := firstTxo + vout

		// Tell the txo that it was spent to txi
		// WE ADD ONE because zero is a valid txi but WE WANT IT TO MEAN UNSPENT
		err = txoSpentTxiWordFile.WriteWordAt(txi+1, txo)
	}
	err = txoSpentTxiWordFile.Close()
	if err != nil {
		return err
	}
	txoSatsWordFile.Close()
	txiTxWordFile.Close()
	txiVoutWordFile.Close()
	transFirstTxoWordFile.Close()
	transFirstTxiWordFile.Close()
	return nil
}
