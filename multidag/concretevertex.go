package multidag

import (
	"encoding/json"
	"pudding-server/jsonstuff"
	"strconv"
)

type concreteVertex struct {
	attributes attributes
	inPoints   endpoints
	outPoints  endpoints
}

func NewConcreteVertex() Vertex {
	v := concreteVertex{}
	v.attributes = attributes{}
	v.attributes.attributes = make(map[string]string)
	return &v
}

type attributes struct {
	attributes map[string]string
}

func (a *attributes) EncodeAsJson() ([]byte, error) {
	jsonBytes, err := json.Marshal(a.attributes)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

type endpoints struct {
	single map[string]singleEndpoint
	multi  map[string]multiEndpoint
}

func (e *endpoints) EncodeAsJson() ([]byte, error) {
	jsonBytes, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return jsonBytes, nil
}

type singleEndpoint struct {
	otherVertex string
	otherlabel  string
}

type multiEndpoint struct {
	totalCount      int
	vertexSelection vertexSelection
}

type vertexSelection struct {
	firstIndex    int
	otherVertices []string
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
	cv.attributes.attributes[key] = value
}
func (cv *concreteVertex) AddSingleInpoint(name string, sourceType string, sourceIndex int64, otherEndLabel string) {
	value := singleEndpoint{}
	value.otherVertex = sourceType + strconv.Itoa(int(sourceIndex))
	value.otherlabel = otherEndLabel
	cv.inPoints.single[name] = value
}
func (cv *concreteVertex) AddMultiInpoint(name string, sourceType string, totalCount int64, selectionIndices []int64) {
	value := multiEndpoint{}
	value.totalCount = int(totalCount)
	vertexSelection := vertexSelection{}
	vertexSelection.firstIndex = 0
	for _, v := range selectionIndices {
		partialUrl := sourceType + strconv.Itoa(int(v))
		vertexSelection.otherVertices = append(vertexSelection.otherVertices, partialUrl)
	}
	value.vertexSelection = vertexSelection
	cv.inPoints.multi[name] = value
}

func (cv *concreteVertex) AddSingleOutpoint(name string, targetType string, targetIndex int64, otherEndLabel string) {
	value := singleEndpoint{}
	value.otherVertex = targetType + strconv.Itoa(int(targetIndex))
	value.otherlabel = otherEndLabel
	cv.inPoints.single[name] = value
}
func (cv *concreteVertex) AddMultiOutpoint(name string, targetType string, totalCount int64, selectionIndices []int64) {
	value := multiEndpoint{}
	value.totalCount = int(totalCount)
	vertexSelection := vertexSelection{}
	vertexSelection.firstIndex = 0
	for _, v := range selectionIndices {
		partialUrl := targetType + strconv.Itoa(int(v))
		vertexSelection.otherVertices = append(vertexSelection.otherVertices, partialUrl)
	}
	value.vertexSelection = vertexSelection
	cv.inPoints.multi[name] = value
}
