package goutils

import (
	"fmt"
	"github.com/peterSZW/goutils/logger"

	//"weixin/utils"
	"encoding/json"
	//"github.com/astaxie/beego"
	"io/ioutil"
	"net/http"
	//"strings"
	//"time"
)

type GoogleAddr struct {
	Formatted_address string `json:"formatted_address"`
}
type GoogleAddrs struct {
	Results []GoogleAddr `json:"results"`
}

func GetAddressByGPS(x float32, y float32) string {
	s := fmt.Sprintf("http://maps.google.com/maps/api/geocode/json?latlng=%f,%f&language=zh-CN&sensor=false", x, y)

	resp, err := http.Get(s)
	if err != nil {
		logger.Error(err)
		return ""
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
	}
	defer resp.Body.Close()

	var googleAddrs GoogleAddrs
	err2 := json.Unmarshal(body, &googleAddrs)
	if err2 != nil {
		logger.Error(err2)
	}
	if len(googleAddrs.Results) > 0 {
		return googleAddrs.Results[0].Formatted_address
	} else {
		return ""
	}
}
