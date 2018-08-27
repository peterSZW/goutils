package goutils

import (
	//"crypto/rand"
	"crypto/md5"
	"encoding/hex"
)

func Md5string(str string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(str))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
