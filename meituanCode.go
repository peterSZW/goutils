package goutils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	ss := hex.EncodeToString(h.Sum(nil))
	ss = "0x" + ss[0:11]

	fmt.Println(strconv.ParseInt(ss, 0, 0))

	return hex.EncodeToString(h.Sum(nil))
}

func Fmt12code(s string) string {
	return s[0:4] + " " + s[4:8] + " " + s[8:]
}

func Get12Code() string {
	h := md5.New()

	var cc string

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	//    for i:=0; i<10; i++ {
	//        fmt.Println(r.Intn(100))
	//    }

	for {

		h.Write([]byte(time.Now().Format("2006-01-02 15:04:05.006") + strconv.Itoa(r.Intn(10000))))
		ss := hex.EncodeToString(h.Sum(nil))

		cc = "0x" + ss[0:12]

		i, _ := strconv.ParseInt(cc, 0, 0)

		cc = fmt.Sprintf("%d", i)

		l := len(cc)
		if l >= 12 {
			cc = cc[l-12 : l]
			break
		}

	}

	return cc
}
