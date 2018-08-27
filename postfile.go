package utils

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func PostFile(filename string, targetUrl string) (error, string) {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("uploadfile", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err, "error writing to buffer"
	}

	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err, "error opening file"
	}

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err, "iocopy"
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err, "post fail"
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, string(resp_body)
	}
	//	fmt.Println(resp.Status)
	//	fmt.Println(string(resp_body))
	return nil, string(resp_body)
}
