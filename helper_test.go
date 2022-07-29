package investingcom

import (
	"fmt"
	"io/ioutil"
)

func loadFixture(name string) []byte {
	filename := fmt.Sprintf("fixtures/%s", name)

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return b
}
