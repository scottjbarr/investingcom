package main

import (
	"fmt"

	"github.com/scottjbarr/investingcom"
)

func main() {
	m, err := investingcom.LoadMapper()
	if err != nil {
		panic(err)
	}

	for _, p := range m.Pairs {
		fmt.Printf("%v,%s\n", p.PairID, p.Name)
	}
}
