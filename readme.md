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
# 作者id
secUserId: ["",""]
# 每个作者下载视频数量(不知道多少的话 可以随便设置 自动 会下载完所有视频)
video-nums: ["100","500"]
# 下载视频保存路径(会根据作者名称在该目录下自动创建文件夹)
filesPath: "D:\\"
# 参考这个仓库 https://github.com/B1gM8c/X-Bogus
x_bogus_addr: ""
```

```bash
go build -o tiktok.exe
```

secUserId获取方式:
打开一个作者的 抖音主页 拿到地址栏的secUserId
https://www.douyin.com/user/MS4wLjABAAAA7xbdm1QfWD8Um6rFnrm0wVpnOI1uEHhbth1XDud_tWRxG5ZI6YUbNu9ES4uMjF0D?is_search=0&list_name=follow&nt=3
`MS4wLjABAAAA7xbdm1QfWD8Um6rFnrm0wVpnOI1uEHhbth1XDud_tWRxG5ZI6YUbNu9ES4uMjF0D` 为作者的secUserId

或者你真的很懒 那就把地址栏的整个地址复制到config.yaml的secUserId里面 有自动处理的
## 擦边球作者的secUserId 直接能用
```json
{
    "娜娜子~": "MS4wLjABAAAAlypPFLZEj21YnlMDhvYhp_xDBSYQO3qZhshy-0koHO8",
    "二丫": "MS4wLjABAAAAQP1C07RU9kTmLEyniA0DYMHZBJCGO9XS8HU1Vdag0nXBi4GaYtHNpZsz5EBfZPgB",
    "和女儿跳舞的波斯猫妞": "MS4wLjABAAAAvcq7vaJZIHkLFM95W_5Pr8PpARQPpZrfSLd3_Yvw953IxdM24XWW7WGueCfSreAR",
    "A上去了": "MS4wLjABAAAA7I564aTdxYOoOZOlQffdQLTZ-XbsDOCohZB_QlySZ_EyfvOxOKdeaXbbU0hLWPsX",
    "🐟彤彤同学": "MS4wLjABAAAAg9BCvf0ikd_G76VxDf0nN-dxtafZDuoWFhe0BKThgNY",
    "八分饱": "MS4wLjABAAAAZXZWBUdn_JknDb_SNP6wJS6-WQ-0CqQUwOF6pWzls9A2HzLI7h3vWoZaIsO0MNn3",
    "小海盐🌙": "MS4wLjABAAAAmEtU1mPrmQU4mYOTgBrkNiB-7itLfB7XExkhfX3Y39HuHw0ZFkvjn2pXZj-uz60s",
    "欲美少女": "MS4wLjABAAAAZ9aDPtZ-hyxxKxudJzE6BGMVTtG0DO6mPmcHi0SzD90jv5uJSVkUvP94V-eGtKRk",
    "叶小喵啦": "MS4wLjABAAAAsNnkQ27HiYZRi8AgbABhW27zQi8FwK49bTlUrhQZop8",
    "_别管我了11": "MS4wLjABAAAAfv7AYteDioF8Ts21H_GkcaXExqLqEa8l1ABiKIvX4oA",
    "南恬": "MS4wLjABAAAAHuFui7x3Walw7nu61E_iQgXBU-23twsaX4FGWpb_bmo",
    "嗯嗯啊啊": "MS4wLjABAAAAjkxEqv9RZpMM0xjx47czGeyZUJiVlWlCCL2kXSUptaNSs4e8fjbL7jW1vnigGeC-",
    "𝐒":"MS4wLjABAAAA5kDzLLO5D3cVUTrHVIDnLDdaPy87-fvOM2wfXB5TFlg",
    "小羊肖恩.":"MS4wLjABAAAAHOHnSnqXeVhFZWnVMeOYDjMzxr1sxZj5vn0SNEkIEqw",
    "织梦映像":"MS4wLjABAAAA7xbdm1QfWD8Um6rFnrm0wVpnOI1uEHhbth1XDud_tWRxG5ZI6YUbNu9ES4uMjF0D"
}
```
