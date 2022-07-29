package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/scottjbarr/investingcom"
)

const (
	PairEURUSD = int(1)
)

func main() {
	pairID := int(0)
	period := ""
	interval := ""
	pointsCount := int(0)

	flag.IntVar(&pairID, "pair", 0, "Pair ID to fetch history for. See investingcom-pairs")
	flag.StringVar(&period, "period", investingcom.PeriodMonth, "Period of time to query. e.g. P1D, P1W, P1M, P3M, P6M, P1Y, P5Y, MAX")
	flag.StringVar(&interval, "interval", investingcom.IntervalDay, "Interval of each candle. e.g. PT1M, PT5M, PT15M, PT30M, PT1H, PT5H, P1D, P1W, P1M")
	flag.IntVar(&pointsCount, "points", 120, "Number of points to return. Affected by period and interval. Default 120")
	flag.Parse()

	if pairID <= 0 || len(period) == 0 || len(interval) == 0 || pointsCount == 0 {
		flag.Usage()
		os.Exit(1)
	}

	c := investingcom.New()

	opts := investingcom.BuildHistoryOpts(period, interval, pointsCount)

	resp, err := c.History(pairID, opts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("time,open,high,low,close,volume\n")
	for _, c := range resp.Data {
		fmt.Printf("%v,%0.4f,%0.4f,%0.4f,%0.4f,%v\n", c.Timestamp, c.Open, c.High, c.Low, c.Close, c.Volume)
	}
}
