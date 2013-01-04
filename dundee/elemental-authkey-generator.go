package dundee

import (
	"crypto/md5"
	"io"
	"strconv"
	"time"
)

const APP_KEY = "jaklsjdflkajsdflajsfkl"

func GenerateElementalAuthKey(endPoint string) (string, string) {
	currTime := time.Now()
	expireTime := currTime.Add(time.Duration(5) * time.Second)
	expires := strconv.FormatInt(expireTime.Unix(), 10)

	h := md5.New()
	io.WriteString(h, expires)
	io.WriteString(h, APP_KEY)
	io.WriteString(h, endPoint)

	return expires, string(h.Sum(nil))
}
