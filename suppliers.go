package cache

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"

	"gopkg.in/kataras/iris.v6"
)

// Converts the request's path to a md5 hash
func RequestPathToMD5(ctx *iris.Context) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(ctx.Request.RequestURI)))
}

// Converts the request's path to a sha1 hash
func RequestPathToSha1(ctx *iris.Context) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(ctx.Request.RequestURI)))
}
