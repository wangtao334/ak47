package sampler

import (
	"github.com/wangtao334/ak47/client"
	"github.com/wangtao334/ak47/data"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type HttpSampler struct {
	Name    string
	Enable  bool
	Url     string
	Method  string
	Queries []*data.Variable
	Headers []*data.Variable
	Body    *data.Variable
}

func (h *HttpSampler) Sample() *SampleResult {
	var err error
	result := AcquireSampleResult()
	result.StartTime = time.Now().UnixNano() / 1e6
	defer func() {
		if result.EndTime == 0 {
			result.EndTime = time.Now().UnixNano() / 1e6
		}
		PutSampleResult(result)
	}()
	req, err := http.NewRequest(h.Method, h.Url, strings.NewReader(h.Body.Value))
	if err != nil {
		result.Err = err
		return result
	}
	queries := req.URL.Query()
	for _, query := range h.Queries {
		queries.Add(query.Name, query.Value)
	}
	req.URL.RawQuery = queries.Encode()
	for _, header := range h.Headers {
		req.Header.Add(header.Name, header.Value)
	}
	res, err := client.AcquireHttpClient().Do(req)
	if err != nil {
		result.Err = err
		return result
	}
	defer func() {
		_ = res.Body.Close()
	}()
	result.StatusCode = res.StatusCode
	result.ResponseData, result.Err = ioutil.ReadAll(res.Body)
	result.EndTime = time.Now().UnixNano() / 1e6
	return result
}

func (h *HttpSampler) Enabled() bool {
	return h.Enable
}

func (h *HttpSampler) Parse(userVariables []*data.Variable) {
	log.Printf("parse variables and functions for %s", h.Name)
	m := map[string]string{}
	for _, v := range userVariables {
		m["${"+v.Name+"}"] = v.Value
	}
	for name, value := range m {
		for _, v := range h.Queries {
			v.Value = strings.ReplaceAll(v.Value, name, value)
		}
		for _, v := range h.Headers {
			v.Value = strings.ReplaceAll(v.Value, name, value)
		}
		if h.Body != nil {
			h.Body.Value = strings.ReplaceAll(h.Body.Value, name, value)
		}
	}
	if h.Body == nil {
		h.Body = &data.Variable{}
	}
}
