package data

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

var (
	csvM = map[string]CSVColData{}
)

type CSVColData []string

func (d CSVColData) V(times int64, _ map[string]*Variable) string {
	if len(d) == 0 {
		return ""
	}
	times--
	index := int(times) % len(d)
	return d[index]
}

func LoadCSV(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	fileName := filepath.Base(path)
	reader := csv.NewReader(f)
	reader.ReuseRecord = true
	var record []string
	var headers []string
	record, err = reader.Read()
	if err != nil {
		return err
	}
	headers = append(headers, record...)
	var key string
	for _, header := range headers {
		key = fmt.Sprintf("${%s,%s}", fileName, header)
		csvM[key] = CSVColData{}
	}

	for {
		record, err = reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		for i, value := range record {
			key = fmt.Sprintf("${%s,%s}", fileName, headers[i])
			csvM[key] = append(csvM[key], value)
		}
	}
	for k, v := range csvM {
		AddFn(k, v)
	}
	return nil
}
