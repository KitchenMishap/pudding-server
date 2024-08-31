package blockchain

import "pudding-server/multidag"

type chain interface {
	GetBlockchainVertex() multidag.Vertex
}
