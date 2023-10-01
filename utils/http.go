package utils

import (
	"io"
	"net/http"

	"github.com/spf13/viper"
)

func HttpNewRequest(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Referer", "https://www.douyin.com")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")

	cookie := viper.Get("cookie").(string)
	req.Header.Add("Cookie", cookie)
	return req, nil
}
