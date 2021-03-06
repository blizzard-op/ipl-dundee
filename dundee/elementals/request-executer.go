package elementals

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func ExecuteRequest(r *http.Request) (*http.Response, []byte, error) {
	client := new(http.Client)

	resp, err := client.Do(r)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, nil, errors.New("Request to elemental did not return 200 (ok).")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	return resp, body, nil
}
