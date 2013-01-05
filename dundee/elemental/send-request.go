package elemental

import (
	"crypto/md5"
	"dundee/types"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func Send(elementalRequest *types.ElementalRequest) (*[]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", elementalRequest.Path, nil)
	if err != nil {
		return nil, err
	}

	expires, authkey := generateAuthkey(elementalRequest)

	req.Header.Add("X-Auth-User", elementalRequest.Elemental.Authuser)
	req.Header.Add("X-Auth-Expires", expires)
	req.Header.Add("X-Auth-Key", authkey)
	req.Header.Add("Accept", "application/xml")
	req.Header.Add("Content-type", "application/xml")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBody, nil
}

func generateAuthkey(endPoint string) (string, string) {
	currTime := time.Now()
	expireTime := currTime.Add(time.Duration(5) * time.Second)
	expires := strconv.FormatInt(expireTime.Unix(), 10)

	h := md5.New()
	io.WriteString(h, expires)
	io.WriteString(h, APP_KEY)
	io.WriteString(h, endPoint)

	return expires, string(h.Sum(nil))
}
