package rate

import (
	"go.uber.org/ratelimit"
)

type Rate interface {
	Init()
	Take() bool
}

type FakeRateLimit struct {
}

func (r *FakeRateLimit) Init() {
}

func (r *FakeRateLimit) Take() bool {
	return true
}

type DefaultRateLimit struct {
	Limit     int
	rateLimit ratelimit.Limiter
}

func (r *DefaultRateLimit) Init() {
	r.rateLimit = ratelimit.New(r.Limit)
}

func (r *DefaultRateLimit) Take() bool {
	r.rateLimit.Take()
	return true
}
