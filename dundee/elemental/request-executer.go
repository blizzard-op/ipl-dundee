package elemental

import (
	"io/ioutil"
	"net/http"
)

func ExecuteRequest(r *http.Request) (*[]byte, error) {
	client := new(http.Client)

	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &respBody, nil
}
