package variable

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/wangtao334/ak47/constant"
	"github.com/wangtao334/ak47/element"
	"github.com/wangtao334/ak47/util"
)

type Text struct {
	*element.Parent
	Value   string
	b       []byte
	indices [][]int
}

func (t *Text) Do(local map[string]string) error {
	if len(t.indices) == 0 {
		local[t.Name] = t.Value
	} else {
		local[t.Name] = util.BytesToString(t.Val(local))
	}
	return nil
}

func (t *Text) Replace(vars map[string]string) {
	for k, v := range vars {
		t.Value = strings.ReplaceAll(t.Value, fmt.Sprintf("${%s}", k), v)
	}
}

func (t *Text) Parse() error {
	t.indices = constant.FindReg.FindAllStringIndex(t.Value, -1)
	if len(t.indices) == 0 {
		t.b = util.StringToBytes(t.Value)
	}
	return nil
}

func (t *Text) Val(vars map[string]string) []byte {
	if len(t.indices) == 0 {
		return t.b
	}
	buf := bytes.Buffer{}
	var start int
	for _, i := range t.indices {
		buf.WriteString(t.Value[start:i[0]])
		buf.WriteString(util.Find(vars, t.Value[i[0]:i[1]]))
		start = i[1]
	}
	buf.WriteString(t.Value[start:])
	return buf.Bytes()
}

type CSV struct {
	*element.Parent
	FilePath string
	headers  []string
	records  [][]string
	n        int
}

func (c *CSV) Do(local map[string]string) error {
	row := rand.Intn(c.n)
	for i, h := range c.headers {
		local[h] = c.records[row][i]
	}
	return nil
}

func (c *CSV) Parse() error {
	file, err := os.Open(c.FilePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}
	if len(records) < 2 {
		return fmt.Errorf("%s : no data", c.FilePath)
	}
	c.headers = records[0]
	filename := filepath.Base(c.FilePath)
	for i := range c.headers {
		c.headers[i] = fmt.Sprintf("%s.%s", filename, c.headers[i])
	}
	c.records = records[1:]
	c.n = len(c.records)
	return nil
}

type File struct {
	*element.Parent
	FilePath string
	b        []byte
}

func (f *File) Do(local map[string]string) error {
	local[f.Name] = util.BytesToString(f.b)
	return nil
}

func (f *File) Parse() error {
	file, err := os.Open(f.FilePath)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	f.b = b
	return nil
}

func (f *File) Val(map[string]string) []byte {
	return f.b
}

func IsVariable(e element.Element) bool {
	switch e.(type) {
	case *Text, *CSV, *File:
		return true
	default:
		return false
	}
}

func IsCSV(e element.Element) bool {
	_, ok := e.(*CSV)
	return ok
}
