package sampler

import (
	"time"

	"github.com/wangtao334/ak47/constant"
	"github.com/wangtao334/ak47/element"
	"github.com/wangtao334/ak47/logger"
	"github.com/wangtao334/ak47/util"
)

type Demo struct {
	*element.Parent
	Method  string
	Url     element.Element
	Headers []element.Element
	Body    element.Element
}

func (d *Demo) Do(local map[string]string) error {
	workerName := local[constant.InnerWorker]
	logger.Info("%s : %v", workerName, local)
	if d.Url != nil {
		logger.Info("%s : %s", workerName, util.BytesToString(d.Url.Val(local)))
	}
	for _, h := range d.Headers {
		logger.Info("%s : [%q=%q]", workerName, h.GetName(), util.BytesToString(h.Val(local)))
	}
	if d.Body != nil {
		logger.Info("%s : %s", workerName, util.BytesToString(d.Body.Val(local)))
	}
	time.Sleep(time.Second)
	return nil
}

func (d *Demo) Replace(global map[string]string) {
	if d.Url != nil {
		d.Url.Replace(global)
	}
	for _, h := range d.Headers {
		h.Replace(global)
	}
	if d.Body != nil {
		d.Body.Replace(global)
	}
}

func (d *Demo) Check() error {
	if d.Url != nil {
		if err := d.Url.Check(); err != nil {
			return err
		}
	}
	for _, h := range d.Headers {
		if err := h.Check(); err != nil {
			return err
		}
	}
	if d.Body != nil {
		return d.Body.Check()
	}
	return nil
}

func (d *Demo) Parse() error {
	if d.Url != nil {
		if err := d.Url.Parse(); err != nil {
			return err
		}
	}
	for _, h := range d.Headers {
		if err := h.Parse(); err != nil {
			return err
		}
	}
	if d.Body != nil {
		return d.Body.Parse()
	}
	return nil
}
