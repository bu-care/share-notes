### 1. 安装embyserver

使用官方镜像：https://hub.docker.com/r/emby/embyserver

参考链接：https://post.smzdm.com/p/ax0r7pe9/

```shell
chmod a+x /dev/dri
docker create
--name=emby
--device /dev/dri:/dev/dri
emby/embyserver:latest
```

这里出现问题，在群辉系统下没有/dev/dri目录，这个文件好像是用来进行硬编码的，不执行这一步好像也可以播放视频。





使用别人创建的镜像：https://hub.docker.com/r/lovechen/embyserver

```shell
docker run \
--network=bridge \
-p '8096:8096' \
-p '8920:8920' \
-p '1900:1900/udp' \
-p '7359:7359/udp' \
-v /data/emby:/config \
-v /data/downloads/:/data \
-e TZ="Asia/Shanghai" \
--device /dev/dri:/dev/dri \
-e UID=0 \
-e GID=0 \
-e GIDLIST=0 \
--restart always \
--name emby \
-d lovechen/embyserver:版本号
```

# 