package blockchain

import (
	"github.com/KitchenMishap/pudding-shed/chainreadinterface"
	"github.com/KitchenMishap/pudding-shed/chainstorage"
	"github.com/KitchenMishap/pudding-shed/indexedhashes"
	"math"
	"pudding-server/derived"
	"pudding-server/multidag"
	"strconv"
)

type ChainReader struct {
	folder        string
	chainRead     chainreadinterface.IBlockChain
	handleCreator chainreadinterface.IHandleCreator
	parents       chainstorage.IParents
}

func NewChainReader(folder string) ChainReader {
	reader := ChainReader{}
	reader.folder = folder
	creator, _ := chainstorage.NewConcreteAppendableChainCreator(
		folder,
		[]string{"time", "mediantime", "difficulty", "strippedsize", "size", "weight"},
		[]string{"size", "vsize", "weight"},
		true)
	readableChain, handleCreator, parents, _, _ := creator.OpenReadOnly()
	reader.chainRead = readableChain
	reader.handleCreator = handleCreator
	reader.parents = parents
	return reader
}

func (cr *ChainReader) GetBlockchainVertex() multidag.Vertex {
	vertex := multidag.NewConcreteVertex()
	hLatestBlock, _ := cr.chainRead.LatestBlock()
	if hLatestBlock.HeightSpecified() {
		latestBlockHeight := hLatestBlock.Height()
		blocksToShow := math.Min(2, float64(latestBlockHeight)+1)
		blockHeights := []int64{}
		for i := 0; i < int(blocksToShow); i++ {
			blockHeights = append(blockHeights, int64(i))
		}
		vertex.AddMultiOutpoint("blocks", "block", latestBlockHeight+1, blockHeights)
	}
	return vertex
}

func (cr *ChainReader) GetBlockVertex(blockHeight int64) multidag.Vertex {
	vertex := multidag.NewConcreteVertex()
	handle, _ := cr.handleCreator.BlockHandleByHeight(blockHeight)
	block, _ := cr.chainRead.BlockInterface(handle)

	// Attributes
	nei, _ := block.NonEssentialInts()
	for k, v := range *nei {
		vertex.AddAttribute(k, strconv.Itoa(int(v)))
	}

	// Parent blockchain
	vertex.AddSingleInpoint("blockchain", "blockchain", 0, "blocks")

	// Children transactions
	transHeights := []int64{}
	transactionCount, _ := block.TransactionCount()
	toShow := math.Min(2, float64(transactionCount))
	for i := 0; i < int(toShow); i++ {
		hTrans, _ := block.NthTransaction(int64(i))
		if hTrans.HeightSpecified() {
			transHeights = append(transHeights, hTrans.Height())
		}
	}
	vertex.AddMultiOutpoint("transactions", "transaction", transactionCount, transHeights)

	return vertex
}

func (cr *ChainReader) GetTransactionVertex(transHeight int64) multidag.Vertex {
	vertex := multidag.NewConcreteVertex()
	handle, _ := cr.handleCreator.TransactionHandleByHeight(transHeight)
	trans, _ := cr.chainRead.TransInterface(handle)

	// Attributes
	nei, _ := trans.NonEssentialInts()
	for k, v := range *nei {
		vertex.AddAttribute(k, strconv.Itoa(int(v)))
	}

	// Parent block
	parentBlockHeight, _ := cr.parents.ParentBlockOfTrans(transHeight)
	vertex.AddSingleInpoint("block", "block", parentBlockHeight, "transactions")

	// Parent txos as txis
	txiTxoHeights := []int64{}
	txiCount, _ := trans.TxiCount()
	toShow := math.Min(2, float64(txiCount))
	for i := int64(0); i < int64(toShow); i++ {
		hTxi, _ := trans.NthTxi(i)
		txi, _ := cr.chainRead.TxiInterface(hTxi)
		hTxo, _ := txi.SourceTxo()
		if hTxo.TxoHeightSpecified() {
			txiTxoHeights = append(txiTxoHeights, hTxo.TxoHeight())
		}
	}
	vertex.AddMultiInpoint("txis", "txo", txiCount, txiTxoHeights)

	// Children Txos
	txoHeights := []int64{}
	txoCount, _ := trans.TxoCount()
	toShow = math.Min(2, float64(txoCount))
	for i := int64(0); i < int64(toShow); i++ {
		hTxo, _ := trans.NthTxo(i)
		if hTxo.TxoHeightSpecified() {
			txoHeights = append(txoHeights, hTxo.TxoHeight())
		}
	}
	vertex.AddMultiOutpoint("txos", "txo", txoCount, txoHeights)

	return vertex
}

func (cr *ChainReader) GetTxoVertex(txoHeight int64, df *derived.DerivedFiles) multidag.Vertex {
	vertex := multidag.NewConcreteVertex()
	handle, err := cr.handleCreator.TxoHandleByHeight(txoHeight)
	if err != nil {
		println(err.Error())
	}
	txo, _ := cr.chainRead.TxoInterface(handle)

	// Attributes
	sats, _ := txo.Satoshis()
	vertex.AddAttribute("satoshis", strconv.Itoa(int(sats)))

	// Parent transaction
	parentTransHeight, _ := cr.parents.ParentTransOfTxo(txoHeight)
	vertex.AddSingleInpoint("transaction", "transaction", parentTransHeight, "txos")

	// Parent Address
	hAddress, _ := txo.Address()
	if hAddress.HeightSpecified() {
		addressHeight := hAddress.Height()
		vertex.AddSingleInpoint("address", "address", addressHeight, "txos")
	}

	// Child Txi
	txi, spent, _ := df.GetTxoSpentTxi(txoHeight)
	if spent {
		vertex.AddSingleOutpoint("spent", "transaction", txi, "txis")
	}

	return vertex
}

func (cr *ChainReader) GetAddressVertex(addrHeight int64) multidag.Vertex {
	vertex := multidag.NewConcreteVertex()
	handle, _ := cr.handleCreator.AddressHandleByHeight(addrHeight)
	addr, _ := cr.chainRead.AddressInterface(handle)

	// Attributes

	// Child txos
	txoSelection := []int64{}
	txoCount, _ := addr.TxoCount()
	toShow := math.Min(2, float64(txoCount))
	for i := int64(0); i < int64(toShow); i++ {
		hTxo, _ := addr.NthTxo(i)
		if hTxo.TxoHeightSpecified() {
			txoSelection = append(txoSelection, hTxo.TxoHeight())
		}
	}
	vertex.AddMultiOutpoint("txos", "txo", txoCount, txoSelection)

	return vertex
}

func (cr *ChainReader) LookupAddress(address string) int64 {
	hAddress, _ := cr.handleCreator.AddressHandleByString(address)
	if hAddress == nil {
		return -1
	}
	if hAddress.HeightSpecified() {
		return hAddress.Height()
	} else {
		return -1
	}
}

func (cr *ChainReader) LookupTransaction(hashAscii string) int64 {
	hash := indexedhashes.Sha256{}
	err := indexedhashes.HashHexToSha256(hashAscii, &hash)
	if err != nil {
		return -1
	}
	hTrans, _ := cr.handleCreator.TransactionHandleByHash(hash)
	if hTrans == nil {
		return -1
	}
	if hTrans.HeightSpecified() {
		return hTrans.Height()
	} else {
		return -1
	}
}

func (cr *ChainReader) LookupBlock(hashAscii string) int64 {
	hash := indexedhashes.Sha256{}
	err := indexedhashes.HashHexToSha256(hashAscii, &hash)
	if err != nil {
		return -1
	}
	hBlock, _ := cr.handleCreator.BlockHandleByHash(hash)
	if hBlock == nil {
		return -1
	}
	if hBlock.HeightSpecified() {
		return hBlock.Height()
	} else {
		return -1
	}
}
