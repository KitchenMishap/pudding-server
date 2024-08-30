package multidag

import (
	"encoding/json"
	"pudding-server/jsonstuff"
	"strconv"
)

type concreteVertex struct {
	attributes Attributes
	inPoints   Endpoints
	outPoints  Endpoints
}

func NewConcreteVertex() Vertex {
	v := concreteVertex{}
	v.attributes = Attributes{}
	v.attributes.Attributes = make(map[string]string)
	v.inPoints = Endpoints{}
	v.inPoints.Single = make(map[string]SingleEndpoint)
	v.inPoints.Multi = make(map[string]MultiEndpoint)
	v.outPoints = Endpoints{}
	v.outPoints.Single = make(map[string]SingleEndpoint)
	v.outPoints.Multi = make(map[string]MultiEndpoint)
	return &v
}

type Attributes struct {
	Attributes map[string]string
}

func (a *Attributes) EncodeAsJson() ([]byte, error) {
	jsonBytes, err := json.Marshal(a.Attributes)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

type Endpoints struct {
	Single map[string]SingleEndpoint `json:"single"`
	Multi  map[string]MultiEndpoint  `json:"multi"`
}

func (e *Endpoints) EncodeAsJson() ([]byte, error) {
	jsonBytes, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

type SingleEndpoint struct {
	OtherVertex string `json:"otherVertex"`
	Otherlabel  string `json:"otherLabel"`
}

type MultiEndpoint struct {
	TotalCount      int             `json:"totalCount"`
	VertexSelection VertexSelection `json:"vertexSelection"`
}

type VertexSelection struct {
	FirstIndex    int      `json:"firstIndex"`
	OtherVertices []string `json:"otherVertices"`
}

func (cv *concreteVertex) GetAttributes() jsonstuff.Jsonable {
	return &cv.attributes
}
func (cv *concreteVertex) GetInEndpoints() jsonstuff.Jsonable {
	return &cv.inPoints
}
func (cv *concreteVertex) GetOutEndpoints() jsonstuff.Jsonable {
	return &cv.outPoints
}

func (cv *concreteVertex) AddAttribute(key string, value string) {
	cv.attributes.Attributes[key] = value
}
func (cv *concreteVertex) AddSingleInpoint(name string, sourceType string, sourceIndex int64, otherEndLabel string) {
	value := SingleEndpoint{}
	value.OtherVertex = sourceType + strconv.Itoa(int(sourceIndex))
	value.Otherlabel = otherEndLabel
	cv.inPoints.Single[name] = value
}
func (cv *concreteVertex) AddMultiInpoint(name string, sourceType string, totalCount int64, selectionIndices []int64) {
	value := MultiEndpoint{}
	value.TotalCount = int(totalCount)
	vertexSelection := VertexSelection{}
	vertexSelection.FirstIndex = 0
	for _, v := range selectionIndices {
		partialUrl := sourceType + strconv.Itoa(int(v))
		vertexSelection.OtherVertices = append(vertexSelection.OtherVertices, partialUrl)
	}
	value.VertexSelection = vertexSelection
	cv.inPoints.Multi[name] = value
}

func (cv *concreteVertex) AddSingleOutpoint(name string, targetType string, targetIndex int64, otherEndLabel string) {
	value := SingleEndpoint{}
	value.OtherVertex = targetType + strconv.Itoa(int(targetIndex))
	value.Otherlabel = otherEndLabel
	cv.inPoints.Single[name] = value
}
func (cv *concreteVertex) AddMultiOutpoint(name string, targetType string, totalCount int64, selectionIndices []int64) {
	value := MultiEndpoint{}
	value.TotalCount = int(totalCount)
	vertexSelection := VertexSelection{}
	vertexSelection.FirstIndex = 0
	for _, v := range selectionIndices {
		partialUrl := targetType + strconv.Itoa(int(v))
		vertexSelection.OtherVertices = append(vertexSelection.OtherVertices, partialUrl)
	}
	value.VertexSelection = vertexSelection
	cv.inPoints.Multi[name] = value
}
