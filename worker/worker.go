package worker

import (
	"github.com/wangtao334/ak47/data"
	"github.com/wangtao334/ak47/rate"
	"github.com/wangtao334/ak47/sampler"
	"go.uber.org/atomic"
	"log"
	"strconv"
	"sync"
	"time"
)

type Worker struct {
	Wait     *sync.WaitGroup
	WorkerId int
	Loops    int
	Duration int64
	EndTime  int64
	Rate     rate.Rate
	Times    *atomic.Int64
	Samplers []sampler.Sampler
	t        int64
	m        map[string]*data.Variable
}

func (w *Worker) Do() {
	log.Printf("worker : %d started", w.WorkerId+1)
	defer w.Wait.Done()
	w.m = make(map[string]*data.Variable)

	// inner variables
	w.m["__workerId"] = &data.Variable{
		Value: strconv.Itoa(w.WorkerId),
	}

	if w.EndTime != 0 {
		w.duration()
	} else {
		w.loops()
	}
}

func (w *Worker) loops() {
	for i := 0; w.Rate.Take() && i < w.Loops; i++ {
		w.test()
	}
}

func (w *Worker) duration() {
	for w.Rate.Take() && time.Now().UnixNano() < w.EndTime {
		w.test()
	}
}

func (w *Worker) test() {
	w.t = w.Times.Add(1)
	for _, s := range w.Samplers {
		if s.Enabled() {
			s.Sample(w.t, w.m)
		}
	}
}
