package elementals

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func (this *ElementalServer) GenerateRequest(method, path string, body []byte) (*http.Request, error) {
	url := this.parseURL(path)

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	expires, authKey := this.generateAuthKey(path)

	req.Header.Add("X-Auth-User", this.Login)
	req.Header.Add("X-Auth-Expires", expires)
	req.Header.Add("X-Auth-Key", authKey)
	req.Header.Add("Accept", "application/xml")
	req.Header.Add("Content-type", "application/xml")

	//fmt.Println(req)

	return req, nil
}

func (this *ElementalServer) generateAuthKey(path string) (string, string) {
	currTime := time.Now()
	expireTime := currTime.Add(time.Duration(10) * time.Minute)
	expires := strconv.FormatInt(expireTime.Unix(), 10)

	h1 := md5.New()
	io.WriteString(h1, path)
	io.WriteString(h1, this.Login)
	io.WriteString(h1, this.ApiKey)
	io.WriteString(h1, expires)
	h1String := fmt.Sprintf("%x", h1.Sum(nil))

	h2 := md5.New()
	io.WriteString(h2, this.ApiKey)
	io.WriteString(h2, h1String)
	h2String := fmt.Sprintf("%x", h2.Sum(nil))

	return expires, h2String
}
