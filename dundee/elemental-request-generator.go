package dundee

import (
	"io"
	"io/ioutil"
	"net/http"
)

const ELEMENTAL_HOST = "http://10.129.232.160"
const AUTH_USER = "ipl"

func GenerateElementalRequest(expires string, authkey string, path string, body []byte) ([]byte, error) {
	client := &http.Client{}
	url := ELEMENTAL_HOST + path

	var reader io.Reader

	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return []byte(""), err
	}

	req.Header.Add("X-Auth-User", AUTH_USER)
	req.Header.Add("X-Auth-Expires", expires)
	req.Header.Add("X-Auth-Key", authkey)
	req.Header.Add("Accept", "application/xml")
	req.Header.Add("Content-type", "application/xml")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte(""), err
	}

	return respBody, nil
}
