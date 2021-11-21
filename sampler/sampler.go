package sampler

import (
	"github.com/wangtao334/ak47/data"
	"sync"
)

var (
	sampleResultPool = sync.Pool{
		New: func() interface{} {
			return &SampleResult{}
		},
	}
)

type Sampler interface {
	Sample(int64, map[string]*data.Variable) *SampleResult
	Enabled() bool
	Parse(userVariables []*data.Variable)
}

type SampleResult struct {
	Name         string
	StartTime    int64
	EndTime      int64
	StatusCode   int
	ResponseData []byte
	Err          error
}

func (s *SampleResult) Reset() {
	s.Name = ""
	s.StartTime = 0
	s.EndTime = 0
	s.StatusCode = 0
	s.ResponseData = s.ResponseData[:0]
	s.Err = nil
}

func AcquireSampleResult() *SampleResult {
	return sampleResultPool.Get().(*SampleResult)
}

func ReleaseSampleResult(result *SampleResult) {
	result.Reset()
	sampleResultPool.Put(result)
}
