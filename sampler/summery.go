package sampler

import (
	"log"
	"time"
)

var (
	results    chan *SampleResult
	statistics map[string]*Statistics
	stopped    chan struct{}
	strAll     = "ALL"
)

type Statistics struct {
	StartTime        int64
	EndTime          int64
	Name             string
	TotalCount       int64
	FailureCount     int64
	TotalElapsedTime int64
}

func PutSampleResult(result *SampleResult) {
	results <- result
}

func StartSummery() {
	results = make(chan *SampleResult, 10000)
	statistics = make(map[string]*Statistics)
	statistics[strAll] = &Statistics{
		Name: strAll,
	}
	stopped = make(chan struct{})
	pre := time.Now().Unix()
	var s *Statistics
	for result := range results {
		s = statistics[strAll]
		if s.StartTime == 0 {
			s.StartTime = result.StartTime
		}
		s.EndTime = result.EndTime
		s.TotalCount++
		if result.Err != nil {
			statistics[strAll].FailureCount++
		}
		elapsedTime := result.EndTime - result.StartTime
		statistics[strAll].TotalElapsedTime += elapsedTime
		ReleaseSampleResult(result)
		now := time.Now().Unix()
		if now-pre >= 5 {
			outputSummery()
			pre = now
		}
	}
	outputSummery()
	stopped <- struct{}{}
}

func StopSummery() {
	close(results)
	<-stopped
}

func outputSummery() {
	statistics := statistics[strAll]
	log.Printf("[%s] Total Count : %d\t\t\tErr count : %d\t\t\tAvg Count/Second : %.2f\t\t\tAvg Elapsed Time(ms) : %.2f", statistics.Name,
		statistics.TotalCount, statistics.FailureCount, float64(statistics.TotalCount)*1000/float64(statistics.EndTime-statistics.StartTime),
		float64(statistics.TotalElapsedTime)/float64(statistics.TotalCount))
}
