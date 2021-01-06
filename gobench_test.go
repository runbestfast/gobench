package main

import (
	"fmt"
	"github.com/rcrowley/go-metrics"
	"log"
	"os"
	"testing"
	"time"
)

func Test_metrics(t *testing.T) {
	// create metrics
	total := 10000
	s := metrics.NewExpDecaySample(total, 0.015) // or metrics.NewUniformSample(1028)
	h := metrics.NewHistogram(s)
	meter := metrics.NewMeter()
	timer := metrics.NewCustomTimer(h, meter)

	metrics.Register("metrics", timer)

	var j int64
	j = 0
	for j < 50 {
		now := time.Now()
		// ra := rand.Intn(2) + 1
		time.Sleep(time.Millisecond * 100)
		j++
		timer.UpdateSince(now)
	}

	ps := timer.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
	du := float64(time.Millisecond)
	duSuffix := time.Millisecond.String()[1:]
	l := log.New(os.Stdout, "", log.LstdFlags)

	l.Printf("==========================================")
	l.Printf("timer\n")
	l.Printf("  count:       %9d\n", timer.Count())
	l.Printf("  min:         %12.2f%s\n", float64(timer.Min())/du, duSuffix)
	l.Printf("  max:         %12.2f%s\n", float64(timer.Max())/du, duSuffix)
	l.Printf("  mean:        %12.2f%s\n", timer.Mean()/du, duSuffix)
	l.Printf("  stddev:      %12.2f%s\n", timer.StdDev()/du, duSuffix)
	l.Printf("  median:      %12.2f%s\n", ps[0]/du, duSuffix)
	l.Printf("  75%%:         %12.2f%s\n", ps[1]/du, duSuffix)
	l.Printf("  95%%:         %12.2f%s\n", ps[2]/du, duSuffix)
	l.Printf("  99%%:         %12.2f%s\n", ps[3]/du, duSuffix)
	l.Printf("  99.9%%:       %12.2f%s\n", ps[4]/du, duSuffix)
	l.Printf("  1-min rate:  %12.2f\n", timer.Rate1())
	l.Printf("  5-min rate:  %12.2f\n", timer.Rate5())
	l.Printf("  15-min rate: %12.2f\n", timer.Rate15())
	l.Printf("  mean rate:   %12.2f\n", timer.RateMean())
}

func Test_DateFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05.000"))
}
