package elemental

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
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
	expireTime := currTime.Add(time.Duration(10) * time.Second)
	expires := strconv.FormatInt(expireTime.Unix(), 10)

	h1 := md5.New()
	io.WriteString(h1, elementalPath)
	io.WriteString(h1, elementalServer.Login)
	io.WriteString(h1, elementalServer.ApiKey)
	io.WriteString(h1, expires)
	h1String := hex.EncodeToString(h1.Sum(nil))

	h2 := md5.New()
	io.WriteString(h2, elementalServer.ApiKey)
	io.WriteString(h2, h1String)
	h2String := hex.EncodeToString(h2.Sum(nil))

	return expires, h2String
}
