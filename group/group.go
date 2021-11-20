package group

import (
	"github.com/wangtao334/ak47/client"
	"github.com/wangtao334/ak47/sampler"
	"github.com/wangtao334/ak47/worker"
	"log"
	"sync"
	"time"
)

type Group struct {
	Name      string
	Enable    bool
	WorkerNum int
	Loops     int
	Duration  int64
	Samplers  []sampler.Sampler
}

func (g *Group) Do() {
	if g.Duration <= 0 && g.Loops <= 0 {
		return
	}
	log.Printf("group : %s started", g.Name)
	go sampler.StartSummery()
	defer sampler.StopSummery()
	var endTime int64
	if g.Duration != 0 {
		endTime = time.Now().UnixNano() + g.Duration*1e9
	}
	client.InitClient()
	defer client.CloseClient()
	wg := &sync.WaitGroup{}
	wg.Add(g.WorkerNum)
	for i := 0; i < g.WorkerNum; i++ {
		wk := &worker.Worker{
			Wait:     wg,
			WorkerId: i,
			Loops:    g.Loops,
			Duration: g.Duration,
			EndTime:  endTime,
			Samplers: g.Samplers,
		}
		go wk.Do()
	}
	wg.Wait()
}
