package investingcom

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestHistoricalResponse(t *testing.T) {
	b := loadFixture("historical-eur-usd.json")

	got := HistoryResponse{}

	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatal(err)
	}

	want := HistoryResponse{
		Data: []OHLC{
			{
				Timestamp: 1652788800000,
				Open:      1.054,
				High:      1.0547,
				Low:       1.0537,
				Close:     1.0541,
				Volume:    1021,
				Currency:  41378,
			},
			{
				Timestamp: 1652789700000,
				Open:      1.054,
				High:      1.054,
				Low:       1.0532,
				Close:     1.0537,
				Volume:    1018,
				Currency:  42399,
			},
		},
		Events: []Event{
			{
				Text:  `<div class="event-container"><a href="/news/arcelormittal-adr-earnings-revenue-beat-in-q2-2856026" target="_blank">ArcelorMittal ADR Earnings, Revenue Beat in Q2</a><div class="event-credits arial_10">VAR_TIMESTAMP</div></div>`,
				Time:  1658985283000,
				Title: "N",
				Type:  "",
				X:     1.6589853e+12,
				Y:     1.0207,
			},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got\n%+v\nwant\n%+v", got, want)
	}
}
