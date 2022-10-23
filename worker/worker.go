package worker

import (
	"sync"
	"time"

	"github.com/wangtao334/ak47/constant"
	"github.com/wangtao334/ak47/element"
	"github.com/wangtao334/ak47/logger"
)

type Worker struct {
	Name    string
	Mode    int
	Loops   int64
	EndTime int64
	Actions []element.Element
	WG      *sync.WaitGroup
}

func (w *Worker) Work(local map[string]string) {
	logger.Info("%s : start to work", w.Name)
	if w.Mode == constant.ModeLoops {
		var count int64
		for count < w.Loops {
			_ = w.do(local)
			count++
		}
	} else {
		for time.Now().UnixNano() < w.EndTime {
			_ = w.do(local)
		}
	}
	w.WG.Done()
	logger.Info("%s : finished", w.Name)
}

func (w *Worker) do(local map[string]string) error {
	for _, action := range w.Actions {
		_ = action.Do(local)
	}
	return nil
}
