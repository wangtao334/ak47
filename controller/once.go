package controller

import (
	"sync"

	"github.com/wangtao334/ak47/constant"
	"github.com/wangtao334/ak47/element"
)

type Once struct {
	*element.Parent
	m *sync.Map
}

func (o *Once) Do(local map[string]string) error {
	if _, ok := o.m.Load(local[constant.InnerGroup]); ok {
		return nil
	}
	o.m.Store(local[constant.InnerGroup], struct {}{})
	for _, child := range o.Children {
		if err := child.Do(local); err != nil {
			return err
		}
	}
	return nil
}

func (o *Once) Replace(global map[string]string) {
	o.m = &sync.Map{}
	o.Parent.Replace(global)
}
