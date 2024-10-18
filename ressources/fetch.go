package groupie

import (
	"encoding/json"
	"net/http"
)


// fetch retrieves JSON data from a specified API endpoint and decodes it into the provided data structure.
func fetch(url string, endpoint string, data interface{}) error {
	res, err := http.Get(url + endpoint)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(data); err != nil {
		return err
	}
	return nil
}
