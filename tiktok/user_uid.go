package tiktok

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"tiktok/utils"
)

type UserUidResp struct {
	Id                 string `json:"id"`
	CreateTime         string `json:"create_time"`
	LastTime           string `json:"last_time"`
	UserUid            string `json:"user_uid"`
	UserUidType        int    `json:"user_uid_type"`
	FirebaseInstanceId string `json:"firebase_instance_id"`
	UserAgent          string `json:"user_agent"`
	BrowserName        string `json:"browser_name"`
}

func UserSelfUid() string {
	const URL = "https://www.douyin.com/aweme/v1/web/query/user/?device_platform=webapp"
	req, _ := utils.HttpNewRequest("GET", URL, nil)
	req.Header.Set("Referer", "https://www.douyin.com/user/self?showTab=favorite_collection")

	h := &http.Client{}
	respStruct := &UserUidResp{}

	resp, err := h.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("获取用户自身的UID失败: %s", err)
	}
	buff := bytes.NewBuffer(make([]byte, 0, 512))
	_, _ = io.Copy(buff, resp.Body)
	if err = json.NewDecoder(buff).Decode(respStruct); err != nil && err != io.EOF {
		log.Fatalf("解析用户自身的UID失败: %s 原始信息为%s", err, buff.String())
	}
	return respStruct.UserUid
}
