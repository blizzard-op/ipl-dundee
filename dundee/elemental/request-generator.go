package elemental

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
	"time"
)

func GenerateRequest(method string, elementalServer *ElementalServer, elementalPath string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest(method, path.Join(elementalServer.URL, elementalPath), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	expires, authKey := generateAuthKey(elementalServer, elementalPath)

	req.Header.Add("X-Auth-User", elementalServer.Login)
	req.Header.Add("X-Auth-Expires", expires)
	req.Header.Add("X-Auth-Key", authKey)
	req.Header.Add("Accept", "application/xml")
	req.Header.Add("Content-type", "application/xml")

	return req, nil
}

func generateAuthKey(elementalServer *ElementalServer, elementalPath string) (string, string) {
	//hashed_key = Digest::MD5.hexdigest("#{key}#{Digest::MD5.hexdigest("#{url.path}#{login}#{key}#{expires}")}")
	currTime := time.Now()
	expireTime := currTime.Add(time.Duration(5) * time.Second)
	expires := strconv.FormatInt(expireTime.Unix(), 10)

	h2 := md5.New()
	io.WriteString(h2, elementalPath)
	io.WriteString(h2, elementalServer.Login)
	io.WriteString(h2, elementalServer.ApiKey)
	io.WriteString(h2, expires)
	h2String := string(h2.Sum(nil))

	h1 := md5.New()
	io.WriteString(h1, elementalServer.ApiKey)
	io.WriteString(h1, h2String)

	//TESTING DIFFERENCE
	fmt.Println(string(h1.Sum(nil)))
	test := md5.New()
	io.WriteString(test, elementalServer.ApiKey)
	io.WriteString(test, elementalPath+elementalServer.Login+elementalServer.ApiKey+expires)
	fmt.Println(string(test.Sum(nil)))
	//TESTING

	return expires, string(h1.Sum(nil))
}
