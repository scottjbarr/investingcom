package investingcom

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func get(httpClient *http.Client, url string, in interface{}) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// unmarshal into the given input struct
	if in == nil {
		return nil
	}

	return json.Unmarshal(body, in)
}
