package blockchain

import (
	"github.com/KitchenMishap/pudding-shed/chainreadinterface"
	"github.com/KitchenMishap/pudding-shed/chainstorage"
	"pudding-server/multidag"
)

type ChainReader struct {
	folder    string
	chainRead chainreadinterface.IBlockChain
}

func NewChainReader(folder string) ChainReader {
	reader := ChainReader{}
	reader.folder = folder
	creator, _ := chainstorage.NewConcreteAppendableChainCreator(folder)
	appendableChain, _, _ := creator.Open(false)
	reader.chainRead = appendableChain.GetAsChainReadInterface()
	return reader
}

func (cr *ChainReader) GetBlockchainVertex() multidag.Vertex {
	vertex := multidag.NewConcreteVertex()
	vertex.AddAttribute("Hello", "There")
	vertex.AddMultiOutpoint("blocks", "block", 800000, []int64{0, 1})
	return vertex
}

func (cr *ChainReader) GetBlockVertex(blockHeight int64) multidag.Vertex {
	vertex := multidag.NewConcreteVertex()
	vertex.AddAttribute("ThisIsA", "Block")
	vertex.AddSingleInpoint("blockchain", "blockchain", 0, "blocks")
	return vertex
}
