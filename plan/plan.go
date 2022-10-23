package plan

import (
	"math/rand"
	"time"

	"github.com/wangtao334/ak47/constant"
	"github.com/wangtao334/ak47/element"
	"github.com/wangtao334/ak47/group"
	"github.com/wangtao334/ak47/logger"
	"github.com/wangtao334/ak47/variable"
)

type TestPlan struct {
	Name      string
	Children  []element.Element
	variables []element.Element
	groups    []element.Element
}

func (t *TestPlan) Run() (err error) {
	logger.Info("%s : start to run", t.Name)
	// find enable variables
	t.findVariables()

	// parse global variables
	for _, v := range t.variables {
		if err = v.Parse(); err != nil {
			return err
		}
	}

	// generate global variable values
	global := map[string]string{}
	for _, v := range t.variables {
		if variable.IsCSV(v) {
			continue
		}
		_ = v.Do(global)
	}
	logger.Info("global variables : %v", global)

	// find enable groups
	t.findGroups()

	// remove disable test elements
	for _, g := range t.groups {
		g.Remove()
	}

	// replace with global variable values
	t.replace(global)

	// check variables and groups
	if err = t.check(); err != nil {
		return
	}

	// parse files/variables
	if err = t.parse(); err != nil {
		return
	}

	// find csv files
	var elements []element.Element
	for _, v := range t.variables {
		if variable.IsCSV(v) {
			elements = append(elements, v)
		}
	}

	// add csv files to groups
	for _, g := range t.groups {
		g.(*group.Group).Parent.Children = append(append([]element.Element{}, elements...), g.(*group.Group).Parent.Children...)
	}

	// init rand seed
	rand.Seed(time.Now().UnixNano())

	// run groups sequentially
	for _, g := range t.groups {
		inner := map[string]string{constant.InnerGroup: g.(*group.Group).Name}
		if err = g.Do(inner); err != nil {
			return err
		}
	}

	return nil
}

func (t *TestPlan) findVariables() {
	for _, child := range t.Children {
		if variable.IsVariable(child) && !child.CanRemove() {
			t.variables = append(t.variables, child)
		}
	}
}

func (t *TestPlan) findGroups() {
	for _, child := range t.Children {
		if group.IsGroup(child) && !child.CanRemove() {
			t.groups = append(t.groups, child)
		}
	}
}

func (t *TestPlan) replace(global map[string]string) {
	for _, child := range t.groups {
		child.Replace(global)
	}
}

func (t *TestPlan) check() (err error) {
	for _, g := range t.groups {
		if err = g.Check(); err != nil {
			return
		}
	}
	return
}

func (t *TestPlan) parse() (err error) {
	for _, g := range t.groups {
		if err = g.Parse(); err != nil {
			return
		}
	}
	return
}
