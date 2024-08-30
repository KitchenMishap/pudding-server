package multidag

import "pudding-server/jsonstuff"

type Vertex interface {
	GetAttributes() jsonstuff.Jsonable
	GetInEndpoints() jsonstuff.Jsonable
	GetOutEndpoints() jsonstuff.Jsonable
	AddAttribute(name string, val string)
	AddSingleInpoint(name string, sourceType string, sourceIndex int64, otherEndLabel string)
	AddMultiInpoint(name string, sourceType string, totalCount int64, selectionIndices []int64)
	AddSingleOutpoint(name string, targetType string, targetIndex int64, otherEndLabel string)
	AddMultiOutpoint(name string, targetType string, totalCount int64, selectionIndices []int64)
}
