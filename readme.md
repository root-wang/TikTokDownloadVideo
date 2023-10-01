<!--
 * @Date: 2023-05-13 18:15:27
 * @LastEditors: root-wang && 276211640@qq.com
 * @LastEditTime: 2023-05-15 18:43:03
 * @FilePath: \TikTok\readme.md
 * @Description: Do not edit
-->
## TikTok 多线程 多用户所有视频批量 无水印下载

### 配置文件 config.yaml
```yaml
# 账户cookie
cookie: ""
# 下载视频保存路径(会根据作者名称在该目录下自动创建文件夹)
files-path: "f:/video/go_video/"

# 参考这个仓库 https://github.com/B1gM8c/X-Bogus
x-bogus-addrr: "http://127.0.0.1:8787/X-Bogus"

# 下载用户的主页视频
sec-user-id: []
user-total-video-nums: []
# 如果未指定某个用户的主页视频下载数量 则全部使用如下值
download-num: "50"

# 下载用户的喜欢视频
favorite-sec-id: "https://www.douyin.com/user/MS4wLjABAAAA_TSEtfu-jMuD8cUNzv_OmWSdm0_x4oe7lzqJTQrn5SHt1ttKe0APnfbEjsnRL-ZS?is_search=0&list_name=follow&nt=1&showTab=like"
total-video-num: "45"

# 下载用户关注的用户视频 默认下载每个用户的所有视频
# self就是下载token 所指向的用户 一般其他用户设置隐私不可见既不可下载其关注的用户视频
following-sec-id: "self"
# 下载关注用户的数量
following-user-num: 5
```

```bash
go build -o tiktok.exe
```

secUserId获取方式:
打开一个作者的 抖音主页 拿到地址栏的secUserId
https://www.douyin.com/user/MS4wLjABAAAA7xbdm1QfWD8Um6rFnrm0wVpnOI1uEHhbth1XDud_tWRxG5ZI6YUbNu9ES4uMjF0D?is_search=0&list_name=follow&nt=3
`MS4wLjABAAAA7xbdm1QfWD8Um6rFnrm0wVpnOI1uEHhbth1XDud_tWRxG5ZI6YUbNu9ES4uMjF0D` 为作者的secUserId

或者你真的很懒 那就把地址栏的整个地址复制到config.yaml的secUserId里面 有自动处理的
## 作者的secUserId 直接能用
```json
{
    "娜娜子~": "MS4wLjABAAAAlypPFLZEj21YnlMDhvYhp_xDBSYQO3qZhshy-0koHO8",
    "二丫": "MS4wLjABAAAAQP1C07RU9kTmLEyniA0DYMHZBJCGO9XS8HU1Vdag0nXBi4GaYtHNpZsz5EBfZPgB"
}
```
