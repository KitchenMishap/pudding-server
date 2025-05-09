package blockchain

import "github.com/KitchenMishap/pudding-server/multidag"

type chain interface {
	GetBlockchainVertex() multidag.Vertex
}
