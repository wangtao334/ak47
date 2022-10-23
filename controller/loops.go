package controller

import (
	"strconv"

	"github.com/wangtao334/ak47/element"
	"github.com/wangtao334/ak47/util"
)

type Loops struct {
	*element.Parent
	Loops string
	loops int64
}

func (l *Loops) Do(local map[string]string) error {
	var count int64
	for count < l.loops {
		for _, child := range l.Children {
			_ = child.Do(local)
		}
		count++
	}
	return nil
}

func (l *Loops) Replace(global map[string]string) {
	l.Loops = util.Find(global, l.Loops)
	l.Parent.Replace(global)
}

func (l *Loops) Check() (err error) {
	if l.loops, err = strconv.ParseInt(l.Loops, 10, 64); err != nil {
		return
	}
	return l.Parent.Check()
}
