package group

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/wangtao334/ak47/constant"
	"github.com/wangtao334/ak47/element"
	"github.com/wangtao334/ak47/logger"
	"github.com/wangtao334/ak47/util"
	"github.com/wangtao334/ak47/worker"
)

type Group struct {
	*element.Parent
	WorkerNum string
	Mode      int
	Loops     string
	Duration  string
	workerNum int64
	loops     int64
	duration  int64
}

func (g *Group) Do(inner map[string]string) error {
	logger.Info("%s : start to run with loops mode - %d", g.Name, g.loops)
	wg := &sync.WaitGroup{}
	var endTime int64
	if g.Mode == constant.ModeDuration {
		endTime = time.Now().Add(time.Duration(g.duration) * time.Second).UnixNano()
	}
	for i := int64(1); i <= g.workerNum; i++ {
		w := worker.Worker{
			Name:    fmt.Sprintf("%s - %d", g.Name, i),
			Mode:    g.Mode,
			Loops:   g.loops,
			EndTime: endTime,
			Actions: g.Children,
			WG:      wg,
		}
		local := map[string]string{
			constant.InnerWorker: w.Name,
		}
		for k, v := range inner {
			local[k] = v
		}
		wg.Add(1)
		go w.Work(local)
	}
	wg.Wait()
	return nil
}

func (g *Group) Replace(global map[string]string) {
	g.WorkerNum = util.Find(global, g.WorkerNum)
	g.Loops = util.Find(global, g.Loops)
	g.Duration = util.Find(global, g.Duration)
	g.Parent.Replace(global)
}

func (g *Group) Check() (err error) {
	if g.workerNum, err = strconv.ParseInt(g.WorkerNum, 10, 64); err != nil {
		return
	}
	if g.Mode == constant.ModeLoops {
		if g.loops, err = strconv.ParseInt(g.Loops, 10, 64); err != nil {
			return
		}
		if g.loops <= 0 {
			return fmt.Errorf("%s : invalid loops - %d", g.Name, g.loops)
		}
	} else if g.Mode == constant.ModeDuration {
		if g.duration, err = strconv.ParseInt(g.Duration, 10, 64); err != nil {
			return
		}
		if g.duration <= 0 {
			return fmt.Errorf("%s : invalid duration - %d", g.Name, g.duration)
		}
	} else {
		return fmt.Errorf("%s : invalid mode - %d", g.Name, g.Mode)
	}
	return g.Parent.Check()
}

func IsGroup(e element.Element) bool {
	switch e.(type) {
	case *Group:
		return true
	}
	return false
}
