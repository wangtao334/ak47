package data

import (
	"bytes"
	"github.com/wangtao334/ak47/constant"
)

type Variable struct {
	Name     string
	Value    string
	segments []Data
	buf      *bytes.Buffer
}

func (v *Variable) V(times int64, m map[string]*Variable) string {
	if len(v.segments) == 0 {
		return v.Value
	}
	v.buf.Reset()
	for _, d := range v.segments {
		v.buf.WriteString(d.V(times, m))
	}
	return v.buf.String()
}

func (v *Variable) Parse() {
	indices := constant.RegVariable.FindAllStringIndex(v.Value, -1)
	if len(indices) > 0 {
		v.buf = bytes.NewBufferString("")
		var pre int
		for _, index := range indices {
			v.segments = append(v.segments, &Variable{
				Value: v.Value[pre:index[0]],
			})
			pre = index[1]
			exp := v.Value[index[0]:index[1]]
			// csv or function
			d := Fn(exp)
			if d != nil {
				v.segments = append(v.segments, d)
				continue
			}
			// goroutine variable
			v.segments = append(v.segments, &GoroutineVariable{
				Name: exp[2 : len(exp)-1],
			})
		}

		// tail
		v.segments = append(v.segments, &Variable{
			Value: v.Value[pre:],
		})

		v.Value = ""
	}
}

type GoroutineVariable struct {
	Name string
}

func (v *GoroutineVariable) V(_ int64, m map[string]*Variable) string {
	if variable, ok := m[v.Name]; ok {
		return variable.Value
	}
	return ""
}
