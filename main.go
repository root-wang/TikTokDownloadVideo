/*
* @Date: 2023-05-13 15:45:47
  - @LastEditors: root-wang && 276211640@qq.com
  - @LastEditTime: 2023-05-15 23:31:33
  - @FilePath: \TikTok\main.go
  - @Description: Do not edit
*/
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"tiktok/tiktok"
	"tiktok/utils"

	"github.com/spf13/viper"
)

func init() {
	utils.PrintInitLogo()
	viper.SetConfigFile("config.yaml") // 指定配置文件路径
	viper.SetConfigName("config")      // 配置文件名称(无扩展名)
	viper.SetConfigType("yaml")        // 如果配置文件的名称中没有扩展名，则需要配置此项
	viper.AddConfigPath("./")          // 还可以在工作目录中查找配置
	err := viper.ReadInConfig()        // 查找并读取配置文件
	if err != nil {                    // 处理读取配置文件的错误
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func main() {

	// 初始化配置文件
	cookie := viper.Get("cookie").(string)
	if cookie == "" {
		panic(fmt.Errorf("cookie不能为空"))
	}
	filesPath := viper.Get("files-path").(string)
	if filesPath == "" {
		panic(fmt.Errorf("下载根目录路径不能为空"))
	} else if _, err := os.Stat(filesPath); err != nil {
		panic(fmt.Errorf("下载根目录路径%s不存在", filesPath))
	}

	// 初始化下载用户的主页视频配置
	secUserIds := viper.Get("sec-user-id").([]interface{})
	videoCounts := viper.Get("user-total-video-nums").([]interface{})
	secUserIdsStr := make([]string, len(secUserIds))
	videoCountsStr := make([]string, len(videoCounts))
	if len(secUserIds) != 0 {
		log.Println("下载用户视频配置正确, 开始下载指定用户的视频")
		if len(secUserIds) != len(videoCounts) {
			fmt.Printf("作者数目%d和每个作者作品下载数目%d长度不一致,将自动下载未指定数目作者的所有视频\n", len(secUserIds), len(videoCounts))
			rest := len(secUserIds) - len(videoCounts)
			downloadNum := viper.Get("download-num").(string)
			for i := 0; i < rest; i++ {
				videoCounts = append(videoCounts, downloadNum)
			}
		}
		// Convert secUserIds and videoCounts to []string type
		for i, v := range secUserIds {
			secUserIdsStr[i] = fmt.Sprintf("%v", v)
		}
		for i, v := range videoCounts {
			videoCountsStr[i] = fmt.Sprintf("%v", v)
		}
		// 下载指定用户的主页视频
		tiktok.DownloadUserVideos(secUserIdsStr, videoCountsStr, filesPath)
	}

	// 初始化下载用户的喜欢视频配置
	favoriteSceId := viper.Get("favorite-sec-id").(string)
	totalVideoNum := viper.Get("total-video-num").(string)
	if favoriteSceId != "" && totalVideoNum != "" {
		log.Println("下载用户的喜欢视频配置正确, 开始下载用户喜欢的视频")

		num, err := strconv.Atoi(totalVideoNum)
		if err != nil {
			panic("total-video-num参数不正确")
		}
		tiktok.DownloadFavoriteVideos(favoriteSceId, num, filesPath)
	}
	// 初始化下载用户关注的用户视频配置

	followingSceId := viper.Get("following-sec-id").(string)
	followingUserNum := viper.Get("following-user-num").(int)
	var followingUserId string
	if followingSceId != "" {
		log.Println("下载用户的关注的博主视频配置正确, 开始下载关注博主的全部视频")
		println("默认下载用户的所有主页视频")

		if followingSceId == "self" {
			followingUserId = tiktok.UserSelfUid()
			followingSceId = ""
		}

		var userIds []string
		var userVideoCounts []string
		userIds = tiktok.FollowingUsersSecId(followingSceId, followingUserId, followingUserNum)
		for _, _ = range userIds {
			userVideoCounts = append(userVideoCounts, "1000")
		}
		tiktok.DownloadUserVideos(userIds, userVideoCounts, filesPath)
	}

}
