package testplan

import (
	"github.com/wangtao334/ak47/data"
	"github.com/wangtao334/ak47/group"
	"github.com/wangtao334/ak47/sampler"
	"log"
	"testing"
)

func TestTestPlan_Do(t1 *testing.T) {
	tp := &TestPlan{
		Name: "tp1",
		Variables: []*data.Variable{
			{
				Name:  "name",
				Value: "wang tao",
			},
			{
				Name:  "age",
				Value: "18",
			},
			{
				Name:  "desc",
				Value: "${__time(2006-01-02 15:04:05)} : ${name} is ${age}",
			},
		},
		Groups: []*group.Group{
			{
				Name:      "g1",
				Enable:    true,
				WorkerNum: 5,
				Duration:  20,
				Loops:     5,
				Samplers: []sampler.Sampler{
					&sampler.HttpSampler{
						Name:   "http 1",
						Enable: true,
						Url:    "http://localhost:8080/hello",
						Method: "POST",
						Queries: []*data.Variable{
							{
								Name:  "q1",
								Value: "q - v1, ${name}",
							},
						},
						Headers: []*data.Variable{
							{
								Name:  "h1",
								Value: "h -v1, ${age}",
							},
						},
						Body: &data.Variable{
							Name:  "body",
							Value: "i am body, ${desc}",
						},
					},
				},
			},
		},
	}
	if err := tp.Do(); err != nil {
		log.Println(err)
	}
}
