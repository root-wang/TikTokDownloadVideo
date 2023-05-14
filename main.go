/*
* @Date: 2023-05-13 15:45:47
  - @LastEditors: root-wang && 276211640@qq.com
  - @LastEditTime: 2023-05-14 18:44:33
  - @FilePath: \TikTok\main.go
  - @Description: Do not edit
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"tiktok/tiktok"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./config.yaml") // 指定配置文件路径
	viper.SetConfigName("config")        // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")          // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("./")            // 还可以在工作目录中查找配置
	err := viper.ReadInConfig()          // 查找并读取配置文件
	if err != nil {                      // 处理读取配置文件的错误
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
	secUserId := viper.Get("secUserId").([]interface{})
	videoCount := viper.Get("video-nums").([]interface{})
	if len(secUserId) != len(videoCount) {
		panic(fmt.Errorf("作者数目和每个作者作品下载数目长度不一致"))
	}
	cookie := viper.Get("cookie").(string)
	if cookie == "" {
		panic(fmt.Errorf("cookie不能为空"))
	}
	filesPath := viper.Get("filesPath").(string)
	if filesPath == "" {
		panic(fmt.Errorf("下载根目录路径不能为空"))
	}
	// 创建一个http客户端
	user_num := len(secUserId)

	for i := 0; i < user_num; i++ {
		url := tiktok.NewReq(secUserId[i].(string), videoCount[i].(string))
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set(
			"User-Agent",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36",
		)
		req.Header.Set("Sec-Ch-Ua-Platform", "Windows")
		req.Header.Set("Referer", "https://www.douyin.com/user/MS4wLjABAAAAfv7AYteDioF8Ts21H_GkcaXExqLqEa8l1ABiKIvX4oA?is_search=0&list_name=follow&nt=3")
		req.Header.Set("Cookie", cookie)

		h := &http.Client{}
		resp, _ := h.Do(req)
		respStruct := &tiktok.UserVideoResp{}
		_ = json.NewDecoder(resp.Body).Decode(respStruct)
		nameVideo := respStruct.GetAllVideoWithName()
		authorName := respStruct.AwemeList[0].Author.Nickname
		var wg = sync.WaitGroup{}
		fmt.Println("开始下载" + authorName + "的视频")
		for n, add := range nameVideo {
			wg.Add(1)
			go download(n, add, &wg, filesPath+authorName+"\\")
		}
		wg.Wait()
		fmt.Println("共下载" + authorName + "的" + fmt.Sprintf("%d", len(nameVideo)) + "个视频")
		fmt.Println("下载完成")
	}
	fmt.Println("全部下载完成")
}

var token = make(chan struct{}, 10)

func download(n string, add string, wg *sync.WaitGroup, filesPath string) {
	token <- struct{}{}
	defer func() { <-token }()
	resp, _ := http.Get(add + ".mp4")
	defer resp.Body.Close()
	// 检查filesPath是否存在
	_, err := os.Stat(filesPath)
	if err != nil {
		// 创建文件夹
		_ = os.Mkdir(filesPath, os.ModePerm)
	}
	out, _ := os.Create(filesPath + n + ".mp4")
	defer out.Close()
	io.Copy(out, resp.Body)
	wg.Done()
}
