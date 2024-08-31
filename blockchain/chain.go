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
}

func NewChainReader(folder string) ChainReader {
	reader := ChainReader{}
	reader.folder = folder
	creator, _ := chainstorage.NewConcreteAppendableChainCreator(
		folder,
		[]string{"time", "mediantime", "difficulty", "strippedsize", "size", "weight"},
		[]string{"size", "vsize", "weight"},
		true)
	readableChain, handleCreator, _ := creator.OpenReadOnly()
	reader.chainRead = readableChain
	reader.handleCreator = handleCreator
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
	nei, _ := block.NonEssentialInts()
	for k, v := range *nei {
		vertex.AddAttribute(k, strconv.Itoa(int(v)))
	}
	vertex.AddSingleInpoint("blockchain", "blockchain", 0, "blocks")
	return vertex
}
