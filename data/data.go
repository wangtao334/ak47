package data

type Data interface {
	V(int64, map[string]*Variable) string
}
