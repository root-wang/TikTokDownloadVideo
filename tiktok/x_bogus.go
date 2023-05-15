package tiktok

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

type XBogusReq struct {
	DevicePlatform string `json:"device_platform"`
	Aid            string `json:"aid"`
	SecUserId      string `json:"sec_user_id"`
	MaxCursor      string `json:"max_cursor"`
	Count          string `json:"count"`
	VersionName    string `json:"version_name"`
	OsVersion      string `json:"os_version"`
}

func NewXBogusReq(secUserId string, videoCount string) string {
	x_bogus_addr := viper.Get("x_bogus_addr").(string)
	req := &XBogusReq{
		DevicePlatform: "android",
		Aid:            "1128",
		SecUserId:      secUserId,
		MaxCursor:      "0",
		Count:          videoCount,
		VersionName:    "23.5.0",
		OsVersion:      "2333",
	}
	var url = strings.Builder{}
	url.WriteString("https://www.douyin.com/aweme/v1/web/aweme/post/?")
	typ := reflect.TypeOf(*req)
	val := reflect.ValueOf(*req)
	num := val.NumField()
	for i := 0; i < num; i++ {
		val := fmt.Sprint(val.Field(i))
		tagVal := typ.Field(i).Tag.Get("json")
		url.WriteString(tagVal + "=" + val + "&")
	}
	xbogusurl := url.String()[:len(url.String())-1]
	body := struct {
		URL       string `json:"url"`
		UserAgent string `json:"user_agent"`
	}{xbogusurl, "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"}

	reqbody, _ := json.Marshal(body)
	resp, err := http.Post(x_bogus_addr, "application/json", bytes.NewBuffer(reqbody))
	if err != nil {
		panic("请求浏览器指纹失败 请设置浏览器指纹处理地址" + err.Error())
	}

	respbody := struct {
		XBogus string `json:"X-Bogus"`
		Param  string `json:"param"`
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&respbody); err != nil {
		resp.Body.Close()
		return ""
	}

	resp.Body.Close()
	return respbody.Param
}
