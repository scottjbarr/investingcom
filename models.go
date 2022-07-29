package investingcom

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const (
	PeriodP1D     = "P1D"
	PeriodP1W     = "P1W"
	PeriodMonth   = "P1M"
	PeriodMonths3 = "P3M"
	PeriodMonths6 = "P6M"
	PeriodYear    = "P1Y"
	PeriodYears5  = "P5Y"
	PeriodMax     = "MAX"

	IntervalMinute    = "PT1M"
	IntervalMinutes5  = "PT5M"
	IntervalMinutes15 = "PT15M"
	IntervalMinutes30 = "PT30M"
	InternvalHour     = "PT1H"
	IntervaHours5     = "PT5H"
	IntervalDay       = "P1D"
	IntervalWeek      = "P1W"
	IntervalMonth     = "P1M"

	PointsCount60  = int(60)
	PointsCount70  = int(60)
	PointsCount120 = int(120)
)

type HistoryOpts struct {
	Period      string
	Interval    string
	PointsCount int
}

// NewHistoryOpts returns a new set of options for a history http request.
//
// Period defaults to 1 month.  Interval defaults to 1 day.  PointsCount defaults to 120.
func NewHistoryOpts() *HistoryOpts {
	return &HistoryOpts{
		Period:      PeriodMonth,
		Interval:    IntervalDay,
		PointsCount: PointsCount120,
	}
}

func BuildHistoryOpts(period, interval string, pointsCount int) *HistoryOpts {
	return &HistoryOpts{
		Period:      period,
		Interval:    interval,
		PointsCount: pointsCount,
	}
}

func (o *HistoryOpts) Params() string {
	if o == nil {
		return NewHistoryOpts().Params()
	}

	params := []string{
		fmt.Sprintf("period=%s", o.Period),
		fmt.Sprintf("interval=%s", o.Interval),
		fmt.Sprintf("pointscount=%v", o.PointsCount),
	}

	return strings.Join(params, "&")
}

type HistoryResponse struct {
	Data   []OHLC  `json:"data"`
	Events []Event `json:"events"`
}

type OHLC struct {
	Timestamp int64
	Open      float64
	High      float64
	Low       float64
	Close     float64
	Volume    int64
	Currency  int64
}

func (c *OHLC) UnmarshalJSON(p []byte) error {
	var ary []interface{}
	if err := json.Unmarshal(p, &ary); err != nil {
		return err
	}

	i, err := parseInt(ary[0])
	if err != nil {
		return err
	}

	c.Timestamp = i

	var f float64

	f, err = parseFloat(ary[1])
	if err != nil {
		return err
	}

	c.Open = f

	f, err = parseFloat(ary[2])
	if err != nil {
		return err
	}

	c.High = f

	f, err = parseFloat(ary[3])
	if err != nil {
		return err
	}

	c.Low = f

	f, err = parseFloat(ary[4])
	if err != nil {
		return err
	}

	c.Close = f

	i, err = parseInt(ary[5])
	if err != nil {
		return err
	}

	c.Volume = i

	i, err = parseInt(ary[6])
	if err != nil {
		return err
	}

	c.Currency = i

	return nil
}

type Event struct {
	Text  string  `json:"text"`
	Time  int64   `json:"time"`
	Title string  `json:"title"`
	Type  string  `json:"news"`
	X     float64 `json:"x"`
	Y     float64 `json:"y"`
}

func parseInt(in interface{}) (int64, error) {
	switch v := in.(type) {
	case int64:
		return v, nil

	case int:
		return int64(v), nil

	case float64:
		f, _ := in.(float64)
		return int64(f), nil

	case string:
		f, _ := strconv.ParseFloat(v, 64)
		return int64(f), nil
	}

	return 0, fmt.Errorf("could not parse %v as int64", in)
}

func parseFloat(in interface{}) (float64, error) {
	s := fmt.Sprintf("%v", in)

	return strconv.ParseFloat(s, 64)
}
