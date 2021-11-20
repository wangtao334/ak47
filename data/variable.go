package data

type Variable struct {
	Name  string
	Value string
}

func (v *Variable) V(times int64) string {
	return v.Value
}
