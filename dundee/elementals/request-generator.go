package elementals

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"path"
	"regexp"
	"strconv"
	"time"
)

var pattern_protocol *regexp.Regexp
var pattern_leadingSlash *regexp.Regexp

func init() {
	var err error
	pattern_protocol, err = regexp.Compile(`^(?:http(?P<secure>s)?:/+)?(?P<hostname>[^/]*)/*$`)
	if err != nil {
		log.Fatalln(err)
	}
	pattern_leadingSlash, err = regexp.Compile(`^/*(?P<path>.*)$`)
	if err != nil {
		log.Fatalln(err)
	}
}

func (this *ElementalServer) GenerateRequest(method, elementalPath string, body []byte) (*http.Request, error) {
	url := this.parseURL(this.Hostname, elementalPath)

	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	expires, authKey := this.generateAuthKey(elementalPath)

	req.Header.Add("X-Auth-User", this.Login)
	req.Header.Add("X-Auth-Expires", expires)
	req.Header.Add("X-Auth-Key", authKey)
	req.Header.Add("Accept", "application/xml")
	req.Header.Add("Content-type", "application/xml")

	return req, nil
}

func (this *ElementalServer) parseURL(host, p string) string {
	newHost := pattern_protocol.ReplaceAllString(host, `http$secure://$hostname`)
	newPath := pattern_leadingSlash.ReplaceAllString(path.Clean(p), `/$path`)

	return newHost + newPath
}

func (this *ElementalServer) generateAuthKey(elementalPath string) (string, string) {
	currTime := time.Now()
	expireTime := currTime.Add(time.Duration(10) * time.Second)
	expires := strconv.FormatInt(expireTime.Unix(), 10)

	h1 := md5.New()
	io.WriteString(h1, elementalPath)
	io.WriteString(h1, this.Login)
	io.WriteString(h1, this.ApiKey)
	io.WriteString(h1, expires)
	h1String := hex.EncodeToString(h1.Sum(nil))

	h2 := md5.New()
	io.WriteString(h2, this.ApiKey)
	io.WriteString(h2, h1String)
	h2String := hex.EncodeToString(h2.Sum(nil))

	return expires, h2String
}
