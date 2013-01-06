package elemental

import (
	"bytes"
	"crypto/md5"
	"io"
	"net/http"
	"strconv"
	"time"
)

func GenerateRequest(elementalServer *ElementalServer, path string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", elementalServer.URL+path, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	expires, authKey := generateAuthKey(elementalServer, path)

	req.Header.Add("X-Auth-User", elementalServer.AuthUser)
	req.Header.Add("X-Auth-Expires", expires)
	req.Header.Add("X-Auth-Key", authKey)
	req.Header.Add("Accept", "application/xml")
	req.Header.Add("Content-type", "application/xml")

	return req, nil
}

func generateAuthKey(elementalServer *ElementalServer, path string) (string, string) {
	currTime := time.Now()
	expireTime := currTime.Add(time.Duration(5) * time.Second)
	expires := strconv.FormatInt(expireTime.Unix(), 10)

	h := md5.New()
	io.WriteString(h, path)

	return expires, string(h.Sum(nil))
}
