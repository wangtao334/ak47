package testplan

import (
	"errors"
	"fmt"
	"github.com/wangtao334/ak47/constant"
	"github.com/wangtao334/ak47/data"
	"github.com/wangtao334/ak47/group"
	"log"
	"strings"
)

type TestPlan struct {
	Name      string
	CSVFiles  []string
	Variables []*data.Variable
	Groups    []*group.Group
}

func (t *TestPlan) Do() error {
	log.Printf("test plan : %s started", t.Name)

	log.Println("load csv files")
	for _, path := range t.CSVFiles {
		if err := data.LoadCSV(path); err != nil {
			return err
		}
	}

	log.Println("replace user variables and functions")
	if err := t.replaceVariables(); err != nil {
		return err
	}

	log.Println("parse sampler variables and functions")
	t.parseSamplers()

	log.Println("test started")
	for _, g := range t.Groups {
		if g.Enable {
			g.Do()
		}
	}

	return nil
}

func (t *TestPlan) replaceVariables() error {
	if loopCount := len(t.Variables); loopCount > 0 {
		// replace functions
		for _, v := range t.Variables {
			exps := constant.RegFunction.FindAllString(v.Value, -1)
			if len(exps) > 0 {
				for _, exp := range exps {
					fn := data.Fn(exp)
					if fn == nil {
						return errors.New(fmt.Sprintf("can not replace function - %s", exp))
					}
					v.Value = strings.Replace(v.Value, exp, fn.V(0, nil), 1)
				}
			}
		}

		// replace variables
		m := map[string]string{}
		for i := 0; i < loopCount && len(m) < loopCount; i++ {
			for _, v := range t.Variables {
				exps := constant.RegVariable.FindAllString(v.Value, -1)
				if len(exps) == 0 {
					m[v.Name] = v.Value
				} else {
					for _, exp := range exps {
						name := exp[2 : len(exp)-1]
						if value, ok := m[name]; ok {
							v.Value = strings.Replace(v.Value, exp, value, 1)
						}
					}
				}
			}
		}
		if len(m) < len(t.Variables) {
			for _, v := range t.Variables {
				exps := constant.RegVariable.FindAllString(v.Value, -1)
				if len(exps) != 0 {
					return errors.New(fmt.Sprintf("can not replace variable - %s", strings.Join(exps, ",")))
				}
			}
		}
	}
	return nil
}

func (t *TestPlan) parseSamplers() {
	for _, g := range t.Groups {
		if !g.Enable {
			continue
		}
		for _, s := range g.Samplers {
			if !s.Enabled() {
				continue
			}
			s.Parse(t.Variables)
		}
	}
}
