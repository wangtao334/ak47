package worker

import (
	"github.com/wangtao334/ak47/rate"
	"github.com/wangtao334/ak47/sampler"
	"log"
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
	Samplers []sampler.Sampler
}

func (w *Worker) Do() {
	log.Printf("worker : %d started", w.WorkerId+1)
	defer w.Wait.Done()
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
	for _, s := range w.Samplers {
		if s.Enabled() {
			s.Sample()
		}
	}
}
