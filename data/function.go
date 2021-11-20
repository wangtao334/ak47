package data

import (
	"strconv"
	"strings"
	"time"
)

var (
	fnM = map[string]Data{}
)

type DateTimeFormat struct {
	Layout string
}

func (d *DateTimeFormat) Value(_ ...interface{}) string {
	if d.Layout == "" {
		return strconv.FormatInt(time.Now().Unix(), 10)
	}
	return time.Now().Format(d.Layout)
}

func Fn(exp string) Data {
	if d, ok := fnM[exp]; ok {
		return d
	}
	if strings.HasPrefix(exp, `${__time(`) {
		layout := exp[9 : len(exp)-2]
		fnM[exp] = &DateTimeFormat{
			Layout: layout,
		}
	}
	return fnM[exp]
}

func AddFn(exp string, fn Data) {
	fnM[exp] = fn
}
