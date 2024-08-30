package jsonstuff

type Jsonable interface {
	EncodeAsJson() ([]byte, error)
}
