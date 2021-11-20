package constant

import "regexp"

var (
	RegVariable = regexp.MustCompile(`\${(?U).*}`)
	RegFunction = regexp.MustCompile(`\${__(?U).*\((?U).*\)}`)
)
