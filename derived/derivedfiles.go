package derived

import (
	"github.com/KitchenMishap/pudding-shed/chainstorage"
	"github.com/KitchenMishap/pudding-shed/wordfile"
)

type DerivedFiles struct {
	folder             string
	txoSpentTxiFactory wordfile.WordFileCreator
	txoSpentTxi        wordfile.ReadAtWordCounter
	privilegedFiles    chainstorage.IPrivilegedFiles
}

func NewDerivedFiles(folder string) (*DerivedFiles, error) {
	result := DerivedFiles{}

	result.folder = folder

	readCreator, err := chainstorage.NewConcreteAppendableChainCreator(folder, []string{}, []string{}, false)
	if err != nil {
		return nil, err
	}
	_, _, _, files, err := readCreator.OpenReadOnly()
	if err != nil {
		return nil, err
	}
	result.privilegedFiles = files

	// Want to use the same bytes per word as something that refers to txis
	var transFirstTxiWordFile wordfile.ReadAtWordCounter = files.TransFirstTxiFile()
	var wordSize int64 = transFirstTxiWordFile.WordSize()

	result.txoSpentTxiFactory = wordfile.NewConcreteWordFileCreator("txospenttxi", folder+"/derived", wordSize)

	return &result, nil
}

func (df *DerivedFiles) OpenReadOnly() error {
	var err error
	df.txoSpentTxi, err = df.txoSpentTxiFactory.OpenWordFileReadOnly()
	return err
}

func (df *DerivedFiles) Close() error {
	return df.txoSpentTxi.Close()
}

func (df *DerivedFiles) GetTxoSpentTxi(txo int64) (txi int64, unspent bool, er error) {
	val, err := df.txoSpentTxi.ReadWordAt(txo)
	if err != nil {
		return -1, true, err
	}
	// Now val is the txi plus one.
	// But if it is already zero, this represents unspent
	if val == 0 {
		return -1, true, nil
	} else {
		return val - 1, false, nil
	}
}
