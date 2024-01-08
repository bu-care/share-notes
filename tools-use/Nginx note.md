## 1 什么是NGINX
Nginx (engine x) 是一个高性能的HTTP和反向代理web服务器，同时也提供了IMAP/POP3/SMTP服务。其特点是占有内存少，并发能力强，
Nginx 是一个安装非常的简单、配置文件非常简洁（还能够支持perl语法）、Bug非常少的服务。Nginx 启动特别容易，并且几乎可以做到==7 * 24不间断运行，即使运行数个月也不需要重新启动。你还能够不间断服务的情况下进行软件版本的升级。==
Nginx代码完全用C语言从头写成。官方数据测试表明能够支持高达 50,000 个并发连接数的响应。

[Nginxopen in new window](https://nginx.org/) 是一个高性能的 HTTP 和反向代理 Web 服务器。
经常用于：
- 反向代理，将客户端的 HTTP 请求转发到后端服务器进行处理
- 负载均衡，将客户端的请求分配到多个后端服务器上进行处理
- Web 服务器，处理静态文件

正向代理是代理客户端的，反向代理是部署在服务端，代理服务端的，请求都是从客户端向服务端发送的请求。
正向代理：
![](image-20231017200038548.png)

反向代理：
![](image-20231017200104969.png)

动静分离，在我们的软件开发中，有些请求是需要后台处理的，有些请求是不需要经过后台处理的（如：css、html、jpg、js等等文件），这些不需要经过后台处理的文件称为静态文件。让动态网站里的动态网页根据一定规则把不变的资源和经常变的资源区分开来，动静资源做好了拆分以后，我们就可以根据静态资源的特点将其做缓存操作。提高资源响应的速度。
![](image-20231017200417074.png)

4、配置监听

nginx的配置文件是conf目录下的nginx.conf，默认配置的nginx监听的端口为80，如果80端口被占用可以修改为未被占用的端口即可。
![](image-20231017200644306.png)
linux下配置，使用默认配置，在nginx根目录下执行

```shell
./configure
make
make install
```
## 2- Nginx常用命令

```shell
cd /usr/local/nginx/sbin/
./nginx                  # 启动
./nginx -s stop          # 停止
./nginx -s quit          # 安全退出
./nginx -s reload        # 重新加载配置文件
ps aux|grep nginx        # 查看nginx进程
```
注意：如何连接不上，检查阿里云安全组是否开放端口，或者服务器防火墙是否开放端口！

## 3- 实战演示
### 1 config 文件讲解
```nginx
# 全局配置

events {
	work_connections   1024;
}

http (
	http 配置
	
	upstream kuangstudy (
		// 负载均衡配置，权重为1，交替执行请求
		server 127.0.0.1:8080 weight=1;
	    server 127.0.0.1:8081 weight=1;
	)

	server (
		listen    80;
		server_name localhost;
		// 代理

		location / (
			// xxx 10.103.230.21 # 根目录请求访问这个服务器
			root  html;
			index index.html index.htm;
			proxy_pass http://kuangstudy; # 这里是上面的 upstream 配置
		)
		location /admin (
			// xxx 10.103.230.18 # /admin 的请求访问这个服务器
		)
	)

	server (
		listen    443;
		server_name localhost;
		// 代理
	)
)
```

本地启动了两个服务，
```ini
upstream lb{
    server 127.0.0.1:8080 weight=1;
    server 127.0.0.1:8081 weight=1;
}

location / {
    proxy_pass http://lb;
}
```

### 2. 文件存储中心

网上参考链接：https://pythonjishu.com/xmeesttywrwapbr/

```bash
# 在本地创建一个名为mydata的数据卷
docker volume create migration-app-assets

# 运行一个新的容器并将该数据卷与容器内部的目录进行关联
docker run -d --name=migration-app-assets -p 18080:80 -v migration-app-assets:/var/www/html/assets nginx
```

chatgpt方法：容器中配置

```bash
# Create directories for uploads and downloads
mkdir -p /var/www/html/assets

# Set permissions
chown -R nginx:nginx /var/www/html
```

只用修改 /etc/nginx/conf.d/default.conf 配置文件即可

```nginx
server {
    listen 80;
    server_name localhost;
  
    location /uploads {
        client_max_body_size 100m;
        dav_methods PUT;
        create_full_put_path on;
        dav_access user:rw group:rw all:r;
        alias /var/www/html/assets;
    }

    location /downloads {
        alias /var/www/html/assets;
        autoindex on;
        autoindex_exact_size off;
        autoindex_localtime on;
    }
}
```

如果要在浏览器页面上添加文件上传按钮并启用选择本地文件进行上传，需要添加配置以及使用HTML和一些前端脚本。

```nginx
server {
    listen 80;
    server_name localhost;
  
    location / {
        root /usr/share/nginx/html;  # 设置网页根目录
        index index.html;         # 设置默认索引文件
    }
  
    location /uploads {
        client_max_body_size 100m;
        dav_methods PUT;
        create_full_put_path on;
        dav_access user:rw group:rw all:r;
        alias /var/www/html/uploads;

        # 允许 POST 请求
        allow all;
        deny all;
        
        # 用于传递必要的参数
        include fastcgi_params;
    }

    location /downloads {
        alias /var/www/html/downloads;
        autoindex on;
        autoindex_exact_size off;
        autoindex_localtime on;
    }
}
```



以下是一个基本的HTML页面示例，其中包含一个上传按钮：

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>File Upload</title>
</head>
<body>

<!-- 文件上传表单 -->
<form id="uploadForm" action="http://localhost/uploads" method="post" enctype="multipart/form-data">
    <input type="file" name="file" id="fileInput" />
    <button type="submit">上传文件</button>
</form>

<script>
    // 使用JavaScript添加事件监听器以处理上传成功或失败的情况
    document.getElementById('uploadForm').addEventListener('submit', function (event) {
        event.preventDefault();
        var fileInput = document.getElementById('fileInput');
        var file = fileInput.files[0];

        if (file) {
            // 处理文件上传
            var formData = new FormData();
            formData.append('file', file);

            fetch('http://localhost/uploads', {
                method: 'POST',
                body: formData,
            })
            .then(response => {
                if (response.ok) {
                    alert('文件上传成功！');
                } else {
                    alert('文件上传失败，请检查服务器配置和文件大小限制。');
                }
            })
            .catch(error => {
                console.error('文件上传失败:', error);
            });
        } else {
            alert('请选择要上传的文件。');
        }
    });
</script>

</body>
</html>
```

百度文言一心方法：

```nginx

user nginx;
worker_processes auto;
error_log /var/log/nginx/error.log;
pid /var/run/nginx.pid;

events {
    worker_connections 1024;
}

http {
    include /etc/nginx/mime.types;
    default_type application/octet-stream;
    log_format main '$remote_addr - $remote_user [$time_local] "$request" '
                    '$status $body_bytes_sent "$http_referer" '
                    '"$http_user_agent" "$http_x_forwarded_for"';
    access_log /var/log/nginx/access.log main;

    server {
        listen 80;
        server_name localhost; # 替换为你的域名或IP地址
        root /var/www/html; # 替换为你的网站根目录路径

        location /upload {
            alias /var/www/html/upload/; # 替换为你的上传目录路径
            upload_pass @pass; # 指定上传处理的后端位置，例如代理到后端服务器进行处理
            upload_store /var/www/html/upload/; # 指定上传文件的存储位置
        }

        location /download {
            alias /var/www/html/upload/; # 替换为你的下载目录路径
            # 其他下载相关的配置...
        }
    }
}
```



- 使用以下Dockerfile内容构建镜像：

```dockerfile
FROM nginx:latest

COPY nginx.conf /etc/nginx/nginx.conf

RUN mkdir -p /var/www/html/upload/

EXPOSE 80
```

构建镜像：

```bash
docker build -t xbu/software_nginx
```



### 3. 搭建nginx反向代理

1. 创建自签名证书

   参考链接：https://github.com/michaelliao/itranswarp.js/blob/master/conf/ssl/gencert.sh

   这个上面说“创建的签名请求的CN必须与域名完全一致，否则无法通过浏览器验证”，是这样么？

   ```bash
   #!/bin/sh
   
   # create self-signed server certificate:
   
   read -p "Enter your domain [www.example.com]: " DOMAIN
   
   echo "Create server key..."
   openssl genrsa -des3 -out $DOMAIN.key 1024
   
   echo "Create server certificate signing request..."
   SUBJECT="/C=US/ST=Mars/L=iTranswarp/O=iTranswarp/OU=iTranswarp/CN=$DOMAIN"
   openssl req -new -subj $SUBJECT -key $DOMAIN.key -out $DOMAIN.csr
   
   echo "Remove password..."
   mv $DOMAIN.key $DOMAIN.origin.key
   openssl rsa -in $DOMAIN.origin.key -out $DOMAIN.key
   
   echo "Sign SSL certificate..."
   openssl x509 -req -days 3650 -in $DOMAIN.csr -signkey $DOMAIN.key -out $DOMAIN.crt
   
   echo "TODO:"
   echo "Copy $DOMAIN.crt to /etc/nginx/ssl/$DOMAIN.crt"
   echo "Copy $DOMAIN.key to /etc/nginx/ssl/$DOMAIN.key"
   echo "Add configuration in nginx:"
   echo "server {"
   echo "    ..."
   echo "    listen 443 ssl;"
   echo "    ssl_certificate     /etc/nginx/ssl/$DOMAIN.crt;"
   echo "    ssl_certificate_key /etc/nginx/ssl/$DOMAIN.key;"
   echo "}"
   ```

   

2. 编辑nginx配置文件 /etc/nginx/conf.d/default.con，注意这里在容器中proxy_pass不能写127.0.0.1，不然它请求的是容器内的服务，需要写宿主机的IP

   ```ini
   server {
       listen 443 ssl;
       server_name  registry.xbu.dev;
       
       ssl_certificate     /var/cert/xbu.dev.crt;
       ssl_certificate_key /var/cert/xbu.dev.key;
   
       # root       /srv/itranswarp/www;
       access_log /var/log/nginx/access.log;
       error_log  /var/log/nginx/error.log;
   
       location / {
           proxy_pass http://10.103.12.238:5000;
           proxy_set_header host $host;
           proxy_set_header X-real-ip $remote_addr;
           proxy_set_header X-forwarded-for $proxy_add_x_forwarded_for;
       }
   }
   
   server {
       listen 80;
       server_name  registry.xbu.dev;
   
       # return 301 https://$host$request_uri;
   
       location / {
           proxy_pass http://10.103.12.238:5000;
       }
   }
   
   server {
       listen 443 ssl;        
       server_name  teleport.xbu.dev;
       
       ssl_certificate     /var/cert/xbu.dev.crt;
       ssl_certificate_key /var/cert/xbu.dev.key;
   
       location / {               
           proxy_pass http://10.103.12.238:3080;
           proxy_set_header host $host;
           real_ip_header X-Forwarded-For;
           proxy_set_header X-REAL-IP $remote_addr;
           proxy_set_header X-forwarded-for $proxy_add_x_forwarded_for;
       }
   }
   ```

   编写/etc/nginx/nginx.conf文件，修改上传文件大小的限制client_max_body_size 4096M;

   ```ini
   
   user  nginx;
   worker_processes  auto;
   
   error_log  /var/log/nginx/error.log notice;
   pid        /var/run/nginx.pid;
   
   
   events {
       worker_connections  1024;
   }
   
   
   http {
       include       /etc/nginx/mime.types;
       default_type  application/octet-stream;
   
       log_format  main  '$remote_addr - $remote_user [$time_local] "$request" '
                         '$status $body_bytes_sent "$http_referer" '
                         '"$http_user_agent" "$http_x_forwarded_for"';
   
       access_log  /var/log/nginx/access.log  main;
   
       sendfile        on;
       #tcp_nopush     on;
   
       keepalive_timeout  65;
   
       #gzip  on;
   
       client_max_body_size 4096M;
   
       add_header 'Access-Control-Allow-Origin' '*';
       add_header 'Access-Control-Allow-Credentials' 'true';
       add_header 'Access-Control-Allow-Headers' 'Authorization,Accept,Origin,DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Content-Range,Range';
       add_header 'Access-Control-Allow-Methods' 'GET,POST,OPTIONS,PUT,DELETE,PATCH';
   
       include /etc/nginx/conf.d/*.conf;
   }
   ```

   

3. 编写dockerfile

   ```dockerfile
   FROM nginx:latest
   
   LABEL maintainer="xbu"
   
   ENV DEBIAN_FRONTEND noninteractive
   
   RUN apt-get -o Acquire::http::Proxy="http://10.50.128.110:3128" update --fix-missing && \
       apt-get -y install vim
   
   RUN DEBIAN_FRONTEND=noninteractive apt-get install -y tzdata && \
      cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
      echo 'Asia/Shanghai' > /etc/timezone
   
   # Copy Nginx configuration file
   COPY default.conf /etc/nginx/conf.d/default.conf
   COPY myCert/xbu.dev.crt myCert/xbu.dev.key /var/cert/
   
   # # Create directories for uploads and downloads
   # RUN mkdir -p /var/www/html/uploads /var/www/html/downloads
   
   # # Set permissions
   # RUN chown -R nginx:nginx /var/www/html
   
   # # Expose ports
   # EXPOSE 80
   
   # Start Nginx
   CMD ["nginx", "-g", "daemon off;"]
   ```

   