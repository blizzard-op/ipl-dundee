package elemental

import (
	"bytes"
	"crypto/md5"
	"io"
	"net/http"
	"strconv"
	"time"
	"path"
)

func GenerateRequest(method string, elementalServer *ElementalServer, path string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, path.Join(elementalServer.URL, path), bytes.NewReader(body))
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
	//hashed_key = Digest::MD5.hexdigest("#{key}#{Digest::MD5.hexdigest("#{url.path}#{login}#{key}#{expires}")}")
	currTime := time.Now()
	expireTime := currTime.Add(time.Duration(5) * time.Second)
	expires := strconv.FormatInt(expireTime.Unix(), 10)

	h2 := md5.New()
	io.WriteString(h2, path)
	io.WriteString(h2, elementalServer.AuthUser)
	io.WriteString(h2, elementalServer.AppKey)
	io.WriteString(h2, expires)
	h2String := string(h2.Sum(nil))

	h1 := md5.New()
	io.WriteString(h1, elementalServer.AppKey)
	io.WriteString(h1, h2String)

	return expires, string(h2.Sum(nil))
}
