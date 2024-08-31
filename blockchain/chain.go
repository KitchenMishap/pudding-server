package blockchain

import (
	"github.com/KitchenMishap/pudding-shed/chainreadinterface"
	"github.com/KitchenMishap/pudding-shed/chainstorage"
	"math"
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
	readableChain, handleCreator, parents, _ := creator.OpenReadOnly()
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
	return vertex
}
