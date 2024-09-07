package main

import (
	"github.com/KitchenMishap/pudding-shed/chainstorage"
	"github.com/KitchenMishap/pudding-shed/wordfile"
	"testing"
)

func TestTxoParent(t *testing.T) {
	folder := "F:/Data/858000AddressesCswParents"
	println("Please wait... opening files")
	creator, _ := chainstorage.NewConcreteAppendableChainCreator(
		folder,
		[]string{},
		[]string{},
		true)
	_, _, ip, ipf, err := creator.OpenReadOnly()
	if err != nil {
		t.Fail()
	}

	for trans := int64(1); trans < 44000000; trans++ {
		if trans%1000000 == 0 {
			println(trans)
		}
		firstTxo, _ := ipf.TransFirstTxoFile().ReadWordAt(trans)
		prevTransFirstTxo, _ := ipf.TransFirstTxoFile().ReadWordAt(trans - 1)
		if firstTxo == prevTransFirstTxo {
			println("Transaction ", trans-1, " has no txos!")
		}
		for txo := prevTransFirstTxo; txo < firstTxo; txo++ {
			parentTrans, _ := ip.ParentTransOfTxo(txo)
			if parentTrans != trans-1 {
				t.Fail()
			}
		}
	}
}

func TestTxoParentDerived(t *testing.T) {
	folder := "F:/Data/858000AddressesCswParents"
	println("Please wait... opening files")
	creator, _ := chainstorage.NewConcreteAppendableChainCreator(
		folder,
		[]string{},
		[]string{},
		true)
	_, _, _, ipf, err := creator.OpenReadOnly()
	if err != nil {
		t.Fail()
	}

	// Work out the word size
	wordSize := ipf.TxiTxFile().WordSize()
	fileCreator := wordfile.NewConcreteWordFileCreator("txoparenttrans", folder+"/derived", wordSize)
	txoParentTransWordFile, err := fileCreator.OpenWordFileReadOnly()
	if err != nil {
		t.Fail()
	}

	for trans := int64(1); trans < 44000000; trans++ {
		if trans%1000000 == 0 {
			println(trans)
		}
		firstTxo, _ := ipf.TransFirstTxoFile().ReadWordAt(trans)
		prevTransFirstTxo, _ := ipf.TransFirstTxoFile().ReadWordAt(trans - 1)
		if firstTxo == prevTransFirstTxo {
			println("Transaction ", trans-1, " has no txos!")
		}
		for txo := prevTransFirstTxo; txo < firstTxo; txo++ {
			parentTrans, _ := txoParentTransWordFile.ReadWordAt(txo)
			if parentTrans != trans-1 {
				t.Fail()
			}
		}
	}
}

func TestFirstTxoInSequence(t *testing.T) {
	folder := "F:/Data/CurrentJob"
	println("Please wait... opening files")
	creator, _ := chainstorage.NewConcreteAppendableChainCreator(
		folder,
		[]string{},
		[]string{},
		true)
	_, _, _, ipf, err := creator.OpenReadOnly()
	if err != nil {
		t.Fail()
	}

	for trans := int64(1); trans < 212979; trans++ {
		if trans%100000 == 0 {
			println(trans)
		}
		firstTxo, _ := ipf.TransFirstTxoFile().ReadWordAt(trans)
		prevTransFirstTxo, _ := ipf.TransFirstTxoFile().ReadWordAt(trans - 1)
		if firstTxo < prevTransFirstTxo {
			t.Error()
		}
	}
}
