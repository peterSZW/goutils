package goutils

import (
	_ "bytes"
	"fmt"
	_ "image"
	_ "image/color"
	_ "image/png"
	"io/ioutil"
	"peterSZW/goutils/logger"

	"github.com/RaymondChou/goqr/pkg"
)

func MakeQrcode(data string, pngfilename string) {

	c, err := qr.Encode(data, qr.L)
	if err != nil {
		logger.Error(err)
	}

	pngdat := c.PNG()

	if true {
		fmt.Println("***************************")
		err := ioutil.WriteFile(pngfilename, pngdat, 0666)
		fmt.Println("err=", err)
	}
}
