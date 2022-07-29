package investingcom

import (
	"encoding/json"

	"github.com/scottjbarr/investingcom/internal"
)

type Pair struct {
	Name   string `json:"name"`
	PairID int    `json:"pair_id,string"`
}

func LoadPairs() ([]Pair, error) {
	b := []byte(internal.PairData)

	pairs := []Pair{}

	if err := json.Unmarshal(b, &pairs); err != nil {
		return nil, err
	}

	return pairs, nil
}

type Mapper struct {
	Pairs []Pair
}

func NewMapper(pairs []Pair) *Mapper {
	return &Mapper{
		Pairs: pairs,
	}
}

func LoadMapper() (*Mapper, error) {
	pairs, err := LoadPairs()
	if err != nil {
		return nil, err
	}

	return NewMapper(pairs), nil
}
