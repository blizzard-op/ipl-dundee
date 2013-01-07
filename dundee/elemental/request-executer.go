package elemental

import (
	"errors"
	"net/http"
)

func ExecuteRequest(r *http.Request) (*http.Response, error) {
	client := new(http.Client)

	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("Request to elemental did not return 200 (ok).")
	}

	return resp, nil
}
