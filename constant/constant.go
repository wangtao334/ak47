package constant

import "regexp"

const (
	ModeLoops = 1
	ModeDuration
)

var (
	FindReg  = regexp.MustCompile(`(?U)\${.+}`)
)

var (
	InnerGroup  = "__group"
	InnerWorker = "__worker"
)
