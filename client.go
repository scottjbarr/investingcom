package investingcom

import (
	"fmt"
	"net/http"
	"time"
)

const (
	host = "https://api.investing.com"

	pathHistory = "/api/financialdata/%v/historical/chart/?%s"
)

type Client struct {
	HTTPClient *http.Client
}

func New() *Client {
	return &Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
	}
}

// History fetches historical OHLC data for a pair.
//
// Candles are returned in order of oldest to most recent.
func (c *Client) History(pairID int, opts *HistoryOpts) (*HistoryResponse, error) {
	path := fmt.Sprintf(pathHistory, pairID, opts.Params())

	r := HistoryResponse{}

	if err := c.get(path, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

// get a response from the API.
func (c *Client) get(path string, in interface{}) error {
	url := host + path
	return get(c.HTTPClient, url, in)
}
