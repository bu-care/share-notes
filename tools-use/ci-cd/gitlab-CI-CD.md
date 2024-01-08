## 1. 项目迁移

首先需要在新的服务服务器上新建一个项目

gitlab默认的master分支是受保护的，需要取消保护之后才能push，否则会报错，无法 push，在settings的repository中把 “Protected Branches”设置为unprotect

然后用 Git Bash 执行以下命令

```bat
git clone --mirror 项目原代码仓库地址 # 将原仓库远程项目克隆到本地会生成一个 .git 文件
cd 生成的.git文件 //cd 进入（xxx.git）文件
git push --mirror 新项目代码仓库地址  # 将本地克隆推送到新服务器
```

最后开发人员将本地的远程仓库地址修改成新服务器地址，在项目根目录下执行以下代码即可

```bash
git remote set-url origin 新地址  //将本地远程仓库地址改为新地址
```



## 2. 自建gitlab

### docker安装gitlab

首先设置数据卷位置，配置一个新的环境变量

```bat
export GITLAB_HOME=/srv/gitlab
```



安装

```bat
sudo docker run --detach \
  --hostname gitlab.xbu.com \
  --publish 443:443 --publish 80:80 --publish 220:22 \
  --name gitlab \
  --restart always \
  --volume $GITLAB_HOME/config:/etc/gitlab \
  --volume $GITLAB_HOME/logs:/var/log/gitlab \
  --volume $GITLAB_HOME/data:/var/opt/gitlab \
  --shm-size 256m \
  gitlab/gitlab-ee
# 网站访问是80端口，fetch、push代码是220端口

sudo docker run --detach \
  --hostname 10.103.12.238 \
  --publish 4430:443 --publish 800:80 --publish 2200:22 \
  --name gitlab02 \
  --restart always \
  --volume $GITLAB_HOME/config:/etc/gitlab \
  --volume $GITLAB_HOME/logs:/var/log/gitlab \
  --volume $GITLAB_HOME/data:/var/opt/gitlab \
  --shm-size 256m \
  xbu-gitlab:0.1
# 网站访问是800端口，fetch、push代码是2200端口
```

### 创建后修改登陆密码

```bat
# 执行gitlab-rails console 命令 开始初始化密码
# 通过 u=User.where(id:1).first 来查找与切换账号（User.all 可以查看所有用户）,也可以使用u=User.where(username:'root').first 来查找
# 通过u.password='wireless_dev'设置密码为 wireless_dev
# 通过u.password_confirmation='wireless_dev' 再次确认密码
# 通过 u.save! 进行保存
root@gitlab:/# gitlab-rails console
--------------------------------------------------------------------------------
 Ruby:         ruby 2.7.5p203 (2021-11-24 revision f69aeb8314) [x86_64-linux]
 GitLab:       14.7.2-ee (39a169b2f25) EE
 GitLab Shell: 13.22.2
 PostgreSQL:   12.7
--------------------------------------------------------------------------------
Loading production environment (Rails 6.1.4.4)
irb(main):001:0> user = User.where(username: ‘root’).first
Traceback (most recent call last):
        1: from (irb):1
NameError (undefined local variable or method `‘root’' for main:Object)
irb(main):002:0> u=User.where(id:1).first
=> #<User id:1 @root>
irb(main):003:0> u.password='wireless_dev'
=> "wireless_dev"
irb(main):004:0> u.password_confirmation='wireless_dev'
=> "wireless_dev"
irb(main):005:0> u.save!
=> true

```



## 3. 创建gitlab-runner

### 1. docker安装gitlab-runner

```bat
docker run -d --name gitlab-runner --restart always \
     -v /srv/gitlab-runner/config:/etc/gitlab-runner \
     -v /var/run/docker.sock:/var/run/docker.sock \
     -v /srv/gitlab-runner/files:/home/code \
     gitlab/gitlab-runner

docker run -d --name wnm-gitlab-runner --restart always \
     -v /srv/wnm-gitlab-runner/config:/etc/gitlab-runner \
     -v /var/run/docker.sock:/var/run/docker.sock \
     gitlab/gitlab-runner
```

>在注册 GitLab Runner 时挂载 `/var/run/docker.sock` 文件是为了让 GitLab Runner 在执行作业时能够与宿主机上的 Docker 引擎进行交互。这样，GitLab Runner 就能够使用 Docker 来创建和管理容器，并在容器中执行作业。在executor使用docker时会在docker容器中再创建docker容器(DinD模式)，此时可能会需要和宿主机docker引擎交互，

### 2. 使用自签名证书
因为自己的gitlab使用的是自签名证书，在注册时会报错证书为未知授权证书
```shell
ERROR: Registering runner... failed                 runner=WDPUeLf8 status=couldn't execute POST against https://10.103.49.239/api/v4/runners: Post "https://10.103.49.239/api/v4/runners": x509: certificate signed by unknown authority
PANIC: Failed to register the runner. You may be having network problems. 
```
所以在注册之前，先下载自签名证书，用来后面注册使用
官方文档教程：https://docs.gitlab.com/runner/configuration/tls-self-signed.html
```shell
# 您可以使用 openssl 客户端将 GitLab 实例的证书下载到 /etc/gitlab-runner/certs
openssl s_client -showcerts -connect gitlab.example.com:443 -servername gitlab.example.com < /dev/null 2>/dev/null | openssl x509 -outform PEM > /etc/gitlab-runner/certs/gitlab.example.com.crt

# 验证文件是否已正确安装
echo | openssl s_client -CAfile /etc/gitlab-runner/certs/gitlab.example.com.crt -connect gitlab.example.com:443 -servername gitlab.example.com
```
示例
```shell
mkdir -p /etc/gitlab-runner/certs/
openssl s_client -showcerts -connect 10.103.49.239:443 -servername 10.103.49.239 < /dev/null 2>/dev/null | openssl x509 -outform PEM > /etc/gitlab-runner/certs/gitlab.10.103.49.239.crt

echo | openssl s_client -CAfile /etc/gitlab-runner/certs/gitlab.10.103.49.239.crt -connect 10.103.49.239:443 -servername 10.103.49.239
```

### 3. 注册gitlab-runner
官方文档：https://docs.gitlab.com/runner/register/index.html#docker

```shell
docker exec -it gitlab-runner gitlab-runner register \
  --non-interactive \
  --url "https://10.103.49.239/" \
  --registration-token "XR6xa77vXxTS_cSuz1se" \
  --executor "docker" \
  --docker-image alpine:latest \
  --description "test zta project" \
  --tag-list "docker,zta" \
  --run-untagged="true" \
  --locked="false" \
  --access-level="not_protected"
  
  --tls-ca-file=/etc/ssl/certs/ca-certificates.crt  

# 示例，使用自签名证书
gitlab-runner register \
  --non-interactive \
  --url "https://10.103.49.239/" \
  --registration-token "WDPUeLf8rxv8V7NYzAdh" \
  --executor "docker" \
  --docker-image alpine:latest \
  --description "test wnm API" \
  --tag-list "xbutest" \
  --run-untagged="true" \
  --locked="false" \
  --access-level="not_protected" \
  --tls-ca-file=/etc/gitlab-runner/certs/gitlab.10.103.49.239.crt

# executor使用docker会在docker容器中再创建docker容器(DinD模式)，为了方便从别的仓库拉取代码，所以选择shell
gitlab-runner register \
  --non-interactive \
  --url "https://10.103.49.239/" \
  --registration-token "WDPUeLf8rxv8V7NYzAdh" \
  --executor "shell" \
  --description "test wnm API" \
  --tag-list "xbutest" \
  --run-untagged="true" \
  --locked="false" \
  --access-level="not_protected" \
  --tls-ca-file=/etc/gitlab-runner/certs/gitlab.10.103.49.239.crt
```

重启容器`docker restart gitlab-runner`

修改配置文件`docker exec -it name vim /etc/gitlab-runner/config.toml`

直接注册

```bat
docker run --rm -v /srv/gitlab-runner/config:/etc/gitlab-runner gitlab/gitlab-runner register \
  --non-interactive \
  --executor "docker" \
  --docker-image alpine:latest \
  --url "https://10.103.49.239/" \
  --registration-token "XR6xa77vXxTS_cSuz1se" \
  --description "test zta project" \
  --tag-list "docker,aws" \
  --run-untagged="true" \
  --locked="false" \
  --access-level="not_protected"
```

在gitlab网站的项目中设置 **.gitlab-ci.yml**文件后，gitlab-runner服务器中会自动生成gitlab-runner目录（/home/gitlab-runner/builds/sQEK3YjX/0/xbu/test_pipeline/.git/），并从gitlab网站拉取最新代码，然后执行 **.gitlab-ci.yml**文件中的shell脚本，如下：

```bat
Running with gitlab-runner 14.7.0 (98daeee0)
  on test_zta sQEK3YjX
Preparing the "shell" executor
00:00
Using Shell executor...
Preparing environment
00:00
Running on f6a6f54f2ede...
Getting source from Git repository
00:01
Fetching changes with git depth set to 20...
Reinitialized existing Git repository in /home/gitlab-runner/builds/sQEK3YjX/0/xbu/test_pipeline/.git/
Checking out 325ab1a8 as master...
Skipping object checkout, Git LFS is not installed.
Skipping Git submodules setup
Executing "step_script" stage of the job script
$ cd ztna_local_environment
$ pip install -r requirements.txt
```

### 4. 遇见的问题

在使用docker in docker的方式 "docker" executor 起不来，报错：exec /usr/bin/dumb-init: argument list too long 或者standard_init_linux.go:228: exec user process caused: argument list too long

```
Running with gitlab-runner 14.7.0 (98daeee0)
  on docker build migration app QUVgF7a1
Preparing the "docker" executor
Using Docker executor with image golang:1.20-bullseye ...
Pulling docker image golang:1.20-bullseye ...
Using docker image sha256:608bff313ea66e415713957d9b840176de7114d87a7c2ea6970fa17b4d226392 for golang:1.20-bullseye with digest golang@sha256:c4c28e1d66ff6e8bfa51543940688d512fc1c128325149021e123bd642a47e55 ...
Preparing environment
exec /usr/bin/dumb-init: argument list too long
ERROR: Job failed: prepare environment: exit code 1. Check https://docs.gitlab.com/runner/shells/index.html#shell-profile-loading for more information
```

解决方法：

 把这个当域名是使用了自签名证书，按照使用自签名证书的步骤去进行注册gitlab-runner可以解决问题：

```shell
mkdir -p /etc/gitlab-runner/certs/
openssl s_client -showcerts -connect example.com:443 -servername example.com < /dev/null 2>/dev/null | openssl x509 -outform PEM > /etc/gitlab-runner/certs/example.com.crt

echo | openssl s_client -CAfile /etc/gitlab-runner/certs/example.com.crt -connect example.com:443 -servername example.com

gitlab-runner register \
  --non-interactive \
  --url "https://example.com/" \
  --registration-token "******" \
  --executor "docker" \
  --docker-image golang:1.20-bullseye \
  --docker-pull-policy "if-not-present" \
  --description "docker build app" \
  --tag-list "builder" \
  --run-untagged="true" \
  --locked="true" \
  --access-level="not_protected" \
  --tls-ca-file=/etc/gitlab-runner/certs/example.com.crt
```

去chatgpt的问答分析：

1. gitlab-runner注册的时候参数用--tls-ca-file=/etc/ssl/certs/ca-certificates.crt，这个参数需要修改为特定的证书么？
   
   --tls-ca-file 参数用于在 GitLab Runner 注册时指定用于验证与 GitLab CI/CD 服务器之间的连接的 **CA 证书**文件。在一般情况下，你可以使用**系统的根证书文件**，通常是 /etc/ssl/certs/ca-certificates.crt，因为它包含了**一组受信任的 CA 证书**。
   
   如果你的 GitLab CI/CD 服务器使用的是**自签名的证书**，或者是一些**自定义的 CA 证书**，你可能需要将 --tls-ca-file 参数修改为指向相应证书的路径。这个证书文件应该是一个包含 CA 证书链的 PEM 文件。
   
   例如，如果你的 GitLab CI/CD 服务器使用自签名证书，你可以将证书导出为一个文件（比如 custom-ca.crt），然后使用这个文件的路径来注册 Runner：
   
   ```bash
   gitlab-runner register --tls-ca-file=/path/to/custom-ca.crt
   ```
   
   总的来说，如果你使用的是由信任的 CA 签发的证书，并且系统根证书文件是最新的，那么默认使用**系统的根证书文件是安全的**。只有在使用自定义或自签名证书时，你才需要修改 --tls-ca-file 参数。
   
2. 注册gitlab-runner的时候使用docker in docker的方式，在环境准备的时候报错exec /usr/bin/dumb-init: argument list too long，但是将参数更换为到处的证书--tls-ca-file=/path/to/custom-ca.crt，就没有报错，这是为什么？是因为加载的证书太长么？

   错误信息 "exec /usr/bin/dumb-init: argument list too long" 表示参数列表太长，这可能与你在使用 Docker-in-Docker（DinD）时传递的参数数量有关。

   在使用 DinD 时，因为是从新再注册一个容器去执行pipeline任务，所以 **gitlab-runner register 命令的参数**会被传递到运行在 Docker 容器内的 GitLab Runner 进程。如果传递的参数列表过长，可能会导致这个错误。

   使用 --tls-ca-file=/path/to/custom-ca.crt 替代 --tls-ca-file=/etc/ssl/certs/ca-certificates.crt 可能导致参数列表变短，从而解决了这个问题。原因在于 **/etc/ssl/certs/ca-certificates.crt 文件通常包含大量的 CA 证书**，如果使用这个文件，参数列表可能会很长。



## 4. 编写gitlab-ci-yaml

参考链接：https://fennay.github.io/gitlab-ci-cn/gitlab-ci-yaml.html

#### 1. 发布 release

参考链接：



## 5. 自建docker仓库

### 1. 服务器节点：

```shell
docker run -itd -v /data/registry:/var/lib/registry -p 5000:5000 --restart always --name registry registry:2
```

使用自建仓库的节点：

在使用时，docker pull默认是https去请求仓库地址，因为上面的自建仓库没有使用nginx反向代理以及设置证书，所以使用仓库的节点要设置daemon.json，如下：

```
{
  "insecure-registries": [
    "10.103.12.238:5000",
    "docker.xbu.io"
  ],
  "exec-opts": ["native.cgroupdriver=systemd"],
  "registry-mirrors": ["https://bjikbp76.mirror.aliyuncs.com"]
}
```



也可以使用nginx反向代理，配置证书

新建docker.ken.io访问配置文件

```
vi /etc/nginx/conf.d/docker_registry.conf
```

写入Nginx配置内容，将访问域名docker.ken.io的HTTP以及HTTPS请求都转到Registry

```
server {
    listen 443;          #监听443端口
    server_name  docker.ken.io;    #监听的域名
    ssl on; #开启SSL
    ssl_certificate     /var/cert/docker.ken.io_bundle.crt;    #证书文件
    ssl_certificate_key /var/cert/docker.ken.io.key;    #私钥文件
    location / {                #转发或处理
        proxy_pass http://127.0.0.1:5000;
    }
}

server {
    listen 80;        #监听80端口
    server_name  docker.ken.io; #监听的域名
    location / {                #转发或处理
        proxy_pass http://127.0.0.1:5000;
    }
}
```

配置传输内容大小
Nginx默认只允许客户端传出1MB的内容，这不满足镜像的提交需求，可以修改Nginx的配置放开限制

```
# 修改配置文件
vi /etc/nginx/nginx.conf

# 在http配置项增加以下配置
http {

        ##省略其他配置##
        client_max_body_size 4096M;
        ##省略其他配置##

}
```

重新加载Nginx配置

```bash
nginx -s reload
```



### 2. 删除私有仓库镜像

获取镜像的 ETag 值

```bash
curl -I -XGET --header "Accept:application/vnd.docker.distribution.manifest.v2+json" \
-u admin:admin http://192.168.91.18:5000/v2/centos/manifests/latest

# 这里必须带上--header，不然获得的hash值不对，无法删除镜像
curl -I -XGET --header "Accept:application/vnd.docker.distribution.manifest.v2+json" http://10.103.12.238:5000/v2/migration-app-build/manifests/v0.0.1
# 返回结果
HTTP/1.1 200 OK
Content-Length: 11744
Content-Type: application/vnd.docker.distribution.manifest.v1+prettyjws
Docker-Content-Digest: sha256:af0507631bb023ff6e1075a30402fd2d7ac88771efbef07e0bc48bfd8f404cc9
Docker-Distribution-Api-Version: registry/2.0
ETag: "sha256:af0507631bb023ff6e1075a30402fd2d7ac88771efbef07e0bc48bfd8f404cc9"
X-Content-Type-Options: nosniff
Date: Fri, 08 Dec 2023 12:24:46 GMT
X-Cache: MISS from STC4PXY00
X-Cache-Lookup: MISS from STC4PXY00:3128
Via: 1.1 STC4PXY00 (squid/3.5.12)
Connection: keep-alive
```

使用ETag删除镜像

```bash
curl -I -XDELETE -u admin:admin \
http://192.168.91.18:5000/v2/centos/manifests/sha256:a1801b843b1bfaf77c501e7a6d3f709401a1e0c83863037fa3aab063a7fdb9dc

curl -I -XDELETE http://10.103.12.238:5000/v2/migration-app-build/manifests/sha256:a3ec6cc821da2d3c7c7c9d6f21a0e08e33a8bfb23fda38f6f4286d0ae7610c83
# 返回结果
HTTP/1.1 202 Accepted
Docker-Distribution-Api-Version: registry/2.0
X-Content-Type-Options: nosniff
Date: Fri, 08 Dec 2023 12:32:35 GMT
Content-Length: 0
X-Cache: MISS from STC4PXY00
X-Cache-Lookup: MISS from STC4PXY00:3128
Via: 1.1 STC4PXY00 (squid/3.5.12)
Connection: keep-alive
```

查看镜像列表

```bash
curl -XGET http://10.103.12.238:5000/v2/magration-app-build/tags/list
```



### 3. CICD推送镜像失败

1. 使用shell方式推送

   参考链接：https://blog.csdn.net/qq_41967899/article/details/103283002

   shell方式执行CICD，每次执行任务的用户是gitlab-runner不是root，不能对docker.sock文件进行操作，所以执行docker没有权限。

   解决方法：将容器内docker组的id修改为和宿主机一样的id，然后将gitlab-runner用户加入到docker组。

   ```bash
   # 容器内创建docker组，指定id
   groupadd -g 997 docker
   # 将gitlab-runner用户加入docker组
   usermod -aG docker gitlab-runner
   root@33f9c34b2e61:/# cat /etc/group
   root:x:0:
   gitlab-runner:x:999:
   docker:x:997:gitlab-runner
   ```

   容器内安装docker

   ```shell
   apt update && apt install docker.io
   # 验证结果
   gitlab-runner@33f9c34b2e61:/$ docker ps
   CONTAINER ID   IMAGE                                COMMAND                  CREATED         STATUS       PORTS                                                                                                                         NAMES
   dabe91a64a97   golang:1.20-bullseye                 "bash"                   2 hours ago     Up 2 hours                                                                                                                                 go-build-test
   ```

2. 使用dind方式推送

   遇到的错误：

   ```shell
   error during connect: Post "http://docker:2375/v1.24/build?buildargs=%7B%7D&cachefrom=%5B%5D&cgroupparent=&cpuperiod=0&cpuquota=0&cpusetcpus=&cpusetmems=&cpushares=0&dockerfile=devtest%2Fpipeline%2FDockerfile&labels=%7B%7D&memory=0&memswap=0&networkmode=default&rm=1&shmsize=0&t=10.103.12.238%3A5000%2Fmigration-app-build%3Av0.0.1&target=&ulimits=null&version=1": dial tcp: lookup docker on 10.190.202.200:53: no such host
   ```

   看网上提供的方案，但没有试过，链接：https://forum.gitlab.com/t/error-during-connect-post-http-docker-2375-v1-40-auth-dial-tcp-lookup-docker-on-169-254-169-254-no-such-host/28678/4

   使用DOCKER_HOST，将请求由 docker:2375 换成 localhost:2375

   ```yml
   prepare_image:
     image: desmart/dind:latest
     stage: prepare
     variables:
       # DOCKER_DRIVER: overlay2
       DOCKER_HOST: tcp://localhost:2375
     rules:
       - if: '$CI_COMMIT_BRANCH == "master"'
       # - if: $CI_COMMIT_TAG
       - exists:
           - build
           # - release
     tags:
       # - prepare
       - builder
     retry: 1
     script:
       # - echo "$CI_REGISTRY"
       # - echo "$CI_REGISTRY_USER"
       # - echo "$CI_REGISTRY_PASSWORD"
       - whoami
       - bash devtest/pipeline/prepare_build_image.sh
   ```

   



## 6. 自建软件源



```bash
docker run -itd -v /srv/software_src:/usr/local/apache2/htdocs -p 18080:80 --restart always --name software_src httpd:latest
```

