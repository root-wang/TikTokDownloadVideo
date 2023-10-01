package tiktok

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"net/http"
)

func NewXBogusReq(xBogusUrl string) string {
	xBogusAddr := viper.Get("x-bogus-addr")
	if xBogusAddr == nil {
		panic("请设置浏览器指纹处理地址\n")
	}

	body := struct {
		URL       string `json:"url"`
		UserAgent string `json:"user_agent"`
	}{xBogusUrl, "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36"}

	reqBody, _ := json.Marshal(body)
	resp, err := http.Post(xBogusAddr.(string), "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		panic("请求浏览器指纹失败 请设置浏览器指纹处理地址" + err.Error())
	}

	respBody := struct {
		XBogus string `json:"X-Bogus"`
		Param  string `json:"param"`
	}{}

	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		resp.Body.Close()
		return ""
	}

	resp.Body.Close()
	return respBody.Param
}
