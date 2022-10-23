package variable

import (
	"bytes"
	"math/rand"
	"strconv"

	"github.com/wangtao334/ak47/element"
	"github.com/wangtao334/ak47/util"
)

type RandomText struct {
	*element.Parent
	Len string
	Set string
	l   int64
	s   []rune
	n   int
}

func (r *RandomText) Do(local map[string]string) error {
	buf := bytes.Buffer{}
	for i := int64(0); i < r.l; i++ {
		buf.WriteRune(r.s[rand.Intn(r.n)])
	}
	local[r.Name] = buf.String()
	return nil
}

func (r *RandomText) Replace(global map[string]string) {
	r.Len = util.Find(global, r.Len)
	r.Set = util.Find(global, r.Set)
}

func (r *RandomText) Check() (err error) {
	if r.l, err = strconv.ParseInt(r.Len, 10, 64); err != nil {
		return
	}
	r.s = []rune(r.Set)
	if len(r.s) == 0 {
		for c := 'a'; c <= 'z'; c++ {
			r.s = append(r.s, c)
		}
		for c := 'A'; c <= 'A'; c++ {
			r.s = append(r.s, c)
		}
		for c := '0'; c <= '9'; c++ {
			r.s = append(r.s, c)
		}
	}
	r.n = len(r.s)
	return nil
}

func (r *RandomText) Val(map[string]string) []byte {
	buf := bytes.Buffer{}
	for i := int64(0); i < r.l; i++ {
		buf.WriteRune(r.s[rand.Intn(r.n)])
	}
	return buf.Bytes()
}
