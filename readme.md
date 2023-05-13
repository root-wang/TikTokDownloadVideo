## TikTok 多线程 多用户所有视频批量 下载

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
```

```bash
go build -o tiktok.exe
```