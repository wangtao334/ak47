package data

type Data interface {
	Value(...interface{}) string
}
