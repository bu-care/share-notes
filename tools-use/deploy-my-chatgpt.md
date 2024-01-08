获取api-key：https://platform.openai.com/api-keys

```
sk-MeNU1MaV82d7I8LgGODl
```

使用代理build镜像

```shell
docker build --build-arg http_proxy=http://10.50.128.110:3128 --build-arg https_proxy=http://10.50.128.110:3128 -t xbu/chatgpt .

# 启动容器
docker run -it --name chatgpt-php -p 38080:80 --restart=always xbu/chatgpt
```

