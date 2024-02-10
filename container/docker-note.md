[TOC]

# 安装 docker

### 1、卸载旧版本

```bat
sudo apt-get remove docker docker-engine docker.io containerd runc

# centos 卸载
yum -y remove docker-ce docker-ce-cli containerd.io 
rm -rf /var/lib/docker 
```

### 2、安装依赖

vecnter 中的虚拟机需要改为阿里云的源，不然第三步的添加、验证密钥无法成功，后面的安装docker也就无法成功。这一步执行失败了也无所谓，接着下面的步骤继续执行。

```bat
# 先更新一下
sudo apt-get update 
# 安装相应的依赖
sudo apt-get install \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
```

### 3、添加、验证密钥

```bash
# 添加 Docker 的官方 GPG 密钥，如果公司网络无法下载，可以尝试命令中使用代理：curl  -x http://10.50.128.110:3128
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
# 验证密钥
sudo apt-key fingerprint 0EBFCD88

# chatgpt 步骤
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
```

### 4、设置一个稳定的仓库

```bat
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
# 再更新一次 apt   
  sudo apt-get update
```

### 5、配置镜像加速器

到目录中设置daemon.json文件，通过修改daemon配置文件/etc/docker/daemon.json来使用加速器

```bat
sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "insecure-registries": [
    "10.103.12.238:5000",
    "docker.xbu.io",
    "registry"
  ],
  "exec-opts": ["native.cgroupdriver=systemd"],
  "registry-mirrors": ["https://bjikbp76.mirror.aliyuncs.com"]
}
EOF

sudo systemctl daemon-reload
# sudo systemctl restart docker
```

### 6、安装最新的 docker

```shell
sudo apt-get install docker-ce docker-ce-cli containerd.io
```

### 7、把普通用户加入到docker组

```bat
sudo gpasswd -a $USER docker
newgrp docker
# 执行 docker info 查看

# 如果遇到报错，dial unix /var/run/docker.sock: connect: permission denied，对/var/run/docker.sock进行赋予666权限
ll /var/run/docker.sock
chmod 666 /var/run/docker.sock
```

## 遇到的docker错误

#### 启动docker遇到错误： * /usr/bin/dockerd not present or not executable

```bat
root@sxlin-OptiPlex-7050:/etc/docker# systemctl status docker.service
● docker.service - LSB: Create lightweight, portable, self-sufficient containers.
     Loaded: loaded (/etc/init.d/docker; generated)
     Active: failed (Result: exit-code) since Wed 2022-02-09 13:47:58 CST; 29s ago
       Docs: man:systemd-sysv-generator(8)
    Process: 922231 ExecStart=/etc/init.d/docker start (code=exited, status=1/FAILURE)

2月 09 13:47:58 sxlin-OptiPlex-7050 systemd[1]: Starting LSB: Create lightweight, portable, self-sufficient containers....
2月 09 13:47:58 sxlin-OptiPlex-7050 docker[922231]:  * /usr/bin/dockerd not present or not executable
2月 09 13:47:58 sxlin-OptiPlex-7050 systemd[1]: docker.service: Control process exited, code=exited, status=1/FAILURE
2月 09 13:47:58 sxlin-OptiPlex-7050 systemd[1]: docker.service: Failed with result 'exit-code'.
2月 09 13:47:58 sxlin-OptiPlex-7050 systemd[1]: Failed to start LSB: Create lightweight, portable, self-sufficient containers..
```

```bat
# 备份
cd /var/lib
mkdir /var/lib/docker1
cp -rf docker/* /var/lib/docker1/
# 重装
sudo apt install docker.io
apt-get install docker-ce docker-ce-cli containerd.io
# 启动
sudo -s
systemctl restart docker

Docker error: HTTP 408 response body: invalid character '<' looking for beginning of value
```

#### Docker error: HTTP 408 response body: invalid character '<' looking for beginning of value

docker pull镜像时遇到408错误，在命令行重新输入 docker login，登录自己的docker账号，则问题解决了。

```shell
docker login -u buxuehu --password-stdin </srv/mydocker
```

# 一、Docker 基础

## docker 常用命令

### 1、镜像命令

```bat
docker images      # 查看所有本地主机上的镜像 可以使用docker image ls代替
docker images -aq  # 显示所有镜像的id

docker search #搜索镜像
docker pull 镜像名[:tag]      # 下载镜像 docker image pull
docker rmi -f 镜像ID          # 删除镜像 docker image rm
```

### 2、容器命令

```bat
docker ps            # 列出所有运行的容器 docker container list
docker run 镜像id    # 新建容器并启动，(run是新建容器，搭配的是镜像ID)

docker rm 容器id       #删除指定容器
docker start 容器id    #启动容器（启动已有的容器）
docker restart 容器id  #重启容器
docker stop 容器id     #停止当前正在运行的容器
docker kill 容器id     #强制停止当前容器
```

#### 1、新建容器并启动（run）

```bat
docker run [可选参数] image名称[:tag] | docker container run [可选参数] image
#参书说明
--name="Name"         容器名字 tomcat01 tomcat02 用来区分容器
-d                     后台方式运行
-it                 使用交互方式运行，进入容器查看内容
-p                     指定容器的端口 -p 8080(宿主机):8080(容器)
            -p ip:主机端口:容器端口
            -p 主机端口:容器端口(常用)
            -p 容器端口
            容器端口
-P(大写)                 随机指定端口

# 使用 -d 后台启动时常见的坑
# docker容器使用后台运行，就必须要有要一个前台进程，docker发现没有应用，就会自动停止
# nginx，容器启动后，发现自己没有提供服务，就会立刻停止，就是没有程序了
```

#### 2、退出容器

```bat
exit         # 容器直接退出
ctrl +p +q   # 容器不停止退出
```

#### 3、进入正在运行的容器  （exec、attach）

```bat
# 方式一
docker exec -it 容器id bashshell   # 进入当前正在运行的容器
# 以  root  身份进入容器：
docker exec -it -u root 容器id /bin/bash
# 方式二
docker attach 容器id
# 区别
docker exec #进入当前容器后开启一个 新的终端 ，可以在里面操作。（常用）
docker attach # 进入容器 正在执行 的终端
```

#### 4、logs、top 、inspect、cp

```bat
docker logs -t --tail n 容器id  #查看n行日志, --tail number  需要显示日志条数
docker logs -ft 容器id          #跟着日志, -tf 显示日志信息（一直更新）

docker top 容器id             # 查看容器中进程信息

docker inspect 容器id         # 查看镜像的元数据

docker cp 容器id:容器内路径 主机目的路径  # 从容器内拷贝到主机上
```

#### 5、run、start、exec、attach的区别

```bat
run       # 新建容器并启动，搭配的是 镜像ID

# 下面搭配的都是 容器ID
start     # 启动已有的容器
exec      # 进入当前容器后开启一个 新的终端 ，可以在里面操作。（常用）
attach    # 进入容器 正在执行 的终端
```

### 3、删除命令

1. 删除镜像之前需要先停掉容器，并且要删掉容器。
2. 需要注意删除镜像和容器的命令不一样。 docker rmi ID ，其中 容器(rm) 和 镜像(rmi)
3. 顺序需要先删除容器

```bat
# 停用全部运行中的容器:
docker stop $(docker ps -q)

# 删除全部容器：
docker rm $(docker ps -aq)

# 删除全部image
docker rmi $(docker images -q)

# 删除id为的image
docker rmi $(docker images | grep "^<none>" | awk "{print $3}")
```

### 4、commit 镜像

```bat
# docker commit 提交容器成为一个新的副本， 命令和git原理类似
docker commit -m="描述信息" -a="作者" 容器id 目标镜像名:[TAG]

# 实战，提交了一个建了devtest虚拟环境的anaconda3
(base) [xbu@localhost ~]$ docker commit -m "Anaconda3 of devtest virtual environment is built" -a "xbu" df7e684a010e anaconda3:1.0
sha256:82f1256b38f73e96fb61b3c6de988b6aa2565b95feb7cdcf3689d3e5c815e249
# 查看
(base) [xbu@localhost ~]$ docker images
REPOSITORY              TAG       IMAGE ID       CREATED          SIZE
anaconda3               1.0       82f1256b38f7   10 seconds ago   3.65GB
tomcat_xbu              1.0       7276dcc23052   7 days ago       684MB
```

## docker数据卷

数据卷挂载分为两种情况，一是指定宿主机路径挂在(type: bind)，二是没有指定宿主机路径( type: volume)。

```bat
# 挂载方式： 指定路径挂载、匿名挂载、具名挂载
-v 宿主机路径：容器内路径 #指定路径挂载 docker volume ls 是查看不到的
-v 容器内路径 #匿名挂载
-v 卷名：容器内路径 #具名挂载
```

### 1. 指定宿主机路径

```bat
# -v, --volume list Bind mount a volume
# docker run -it -v 主机目录:容器内目录 -p 主机端口:容器内端口
docker run -it -v /home/ceshi:/home centos /bin/bash
#通过 docker inspect 容器id 查看，可以看到数据卷挂在宿主机 /home/ceshi 路径下
```

### 2. 没有指定宿主机路径

没有指定宿主机数据挂载路径时，数据都是挂载在 ```/var/lib/docker/volumes/xxxx/_data```路径下，可以使用```docker volume ls```命令查看，又分为匿名挂在与具名挂在。

#### 1. 匿名挂载

```bat
# 匿名挂载
# -v 容器内路径
docker run -d -P --name nginx01 -v /etc/nginx nginx
# 这里发现，我们在 -v只写了容器内的路径，没有写容器外的路径，这种就是匿名挂载。
# 使用 docker volume ls 查看数据卷
[root@localhost _data]# docker volume ls
DRIVER    VOLUME NAME
local     d2f68bfce5de6671a022ef0dbc1959879d8adddb8f850fbf3c8e20fd1b977e36
local     de3ff82855055c094a9f55c4993e83716e34f373f736004bcaf549a33aa66557
# 数据是存储在宿主机的 /var/lib/docker/volumes/d2f68bfce5d/_data 路径下
```

#### 2. 具名挂载

```bat
# 给数据卷的起了名字，
docker run -d -P --name nginx02 -v xbu-nginx:/etc/nginx nginx
# 使用 docker volume ls 查看数据卷
[root@localhost _data]# docker volume ls
DRIVER    VOLUME NAME
local     xbu-nginx
# 数据是存储在宿主机的 /var/lib/docker/volumes/xbu-nginx/_data 路径下
```

### 3. 拓展

```bat
# 通过 -v 容器内路径： ro rw 改变读写权限
ro #readonly 只读
rw #readwrite 可读可写
docker run -d -P --name nginx05 -v xbu:/etc/nginx:ro nginx
docker run -d -P --name nginx05 -v xbu:/etc/nginx:rw nginx
# ro 只要看到ro就说明这个路径只能通过宿主机来操作，容器内部是无法操作！
```

### 4. 初识 dockerfile

Dockerfile 就是用来构建docker镜像的命令脚本！通过这个脚本可以生成镜像 。

## Dockerfile

### 1.  DockerFile介绍

dockerfile 是用来构建docker镜像的文件！命令参数脚本！  

构建步骤：
1、 编写一个dockerfile文件
2、 docker build 构建称为一个镜像
3、 docker run运行镜像
4、 docker push发布镜像（DockerHub 、阿里云仓库)

看一下官方的 centos 怎么做的

![image-20220104162047049](images/docker学习笔记.assets/image-20220104162047049.png)

![image-20220104162645706](images/docker学习笔记.assets/image-20220104162645706.png)

### 2.  DockerFile构建过程

基础知识：
1、每个保留关键字(指令）都是必须是大写字母
2、执行从上到下顺序
3、#表示注释
4、每一个指令都会创建提交一个新的镜像层，并提交！  

DockerFile：构建文件，定义了一切的步骤，源代码
DockerImages：通过DockerFile构建生成的镜像，最终发布和运行产品。
Docker容器：容器就是镜像运行起来提供服务。  

#### 1、DockerFile常用指令

```bat
FROM         # 基础镜像，一切从这里开始构建
MAINTAINER   # maintainer 镜像是谁写的， 姓名+邮箱
RUN          # 镜像构建的时候需要运行的命令
ADD         # 步骤，tomcat镜像，这个tomcat压缩包！添加内容 添加同目录
WORKDIR     # 镜像的工作目录
VOLUME         # 挂载的目录
EXPOSE         # 保留端口配置
ONBUILD     # on build 当构建一个被继承 DockerFile 这个时候就会运行ONBUILD的指令，触发指令。
COPY        # 类似ADD，将我们文件拷贝到镜像中
ENV         # 构建的时候设置环境变量！

ENTRYPOINT     # entry point 指定这个容器启动的时候要运行的命令，可以追加命令
USER         # 设置构建用户
```

#### 2、CMD与ENTRYPOINT

```bat
# CMD和RUN命令相似，CMD可以用于执行特定的命令。CMD每次启动容器时运行，RUN在创建镜像时执行一次，固化在image        中,RUN命令先于CMD和ENTRYPOINT
# Dockerfile只允许使用一次CMD指令，只有最后一个会生效，一般都是脚本中最后一条指令。
# 如果docker run后面出现与CMD指定的相同的命令，那么CMD就会被覆盖。而ENTRYPOINT会把容器名后面的所有内容都当成参    数传递给其指定的命令
CMD ["executable","param1","param2"]  #CMD 的推荐格式。
CMD ["param1","param2"] #为 ENTRYPOINT 提供额外的参数，此时 ENTRYPOINT 必须使用 Exec 格式。 
CMD command param1 param2  #Shell 格式

# 类似CMD.配置容器启动后执行的命令，但是它不可被 docker run 提供的参数覆盖
# 每个 Dockerfile 中只能有一个 ENTRYPOINT，当指定多个时，只有最后一个起效。
ENTRYPOINT     # entry point 指定这个容器启动的时候要运行的命令，可以追加命令
ENTRYPOINT ["executable", "param1", "param2"]# exec 这是 ENTRYPOINT 的推荐格式
ENTRYPOINT command param1 param2  #shell格式

# 备注
#Exec 格式 ：ENTRYPOINT 中的参数始终会被使用，而 CMD 的额外参数可以在容器启动时动态替换掉。
#ENTRYPOINT 的 Exec 格式用于设置要执行的命令及其参数，同时可通过 CMD 提供额外的参数。  
#举例： Dockerfile 片段： 
ENTRYPOINT ["/bin/echo", "Hello"]   
CMD ["world"] 
#当容器通过 docker run -it [image] 启动时，输出为： 
Hello world 
#而如果通过 docker run -it [image] haha启动，则输出为： 
Hello haha
#Shell 格式： ENTRYPOINT 的 Shell 格式会忽略任何 CMD 或 docker run 提供的参数。 
```

#### 3、VOLUME

```bat
# 向镜像创建的容器中添加数据卷，数据卷可以在容器之间共享和重用。
# 数据卷的修改是立即生效的。数据卷的修改会对更新镜像产生影响。数据卷会一直存在，直到没有任何容器使用它。
VOLUME ["/root/data1", "/root/data2"] # 匿名挂在，指定了容器内的两个挂载点 /root/data1 和 /root/data2
VOLUME /data
# 容器共享卷（挂载点）
--volumes-from 容器id或容器名
例如：docker run --name 容器2 -it --volumes-from 容器1  ubuntu  /bin/bash
```

### 3.  实战

#### 1. 创建一个自己的centos

##### 1、编写 Dockerfile 文件

```bat
# 创建 dockerfile 文档
vim mydockerfile-centos
```

```shell
# Dockerfile 的内容
FROM centos
MAINTAINER xbu<xbu@sonicwall.com>
# 环境目录设为 /usr/local
ENV MYPATH /usr/local
# 工作目录也为 /usr/local
WORKDIR $MYPATH
# 安装 vim、net-tools（可以使用ifconfig）
RUN yum -y install vim
RUN yum -y install net-tools
# 暴露端口 80
EXPOSE 80  
CMD echo $MYPATH
CMD echo "-----end----"
CMD /bin/bash
```

##### 2、通过这个文件构建镜像

```shell
# 命令 docker build -f 文件路径 -t 镜像名:[tag] .
docker build -f mydockerfile-centos -t mycentos:0.1 . # ‘.’ 表示在当前目录
```

##### 3、查看镜像制作过程

```shell
# 使用 docker history 镜像id 查看
(base) [xbu@localhost dockerfile_test]$ docker history a389a73653b0
IMAGE          CREATED         CREATED BY                                      SIZE      COMMENT
a389a73653b0   6 minutes ago   /bin/sh -c #(nop)  CMD ["/bin/sh" "-c" "/bin…   0B
0dcba12cd4c0   6 minutes ago   /bin/sh -c #(nop)  CMD ["/bin/sh" "-c" "echo…   0B
664da0cc91a3   6 minutes ago   /bin/sh -c #(nop)  CMD ["/bin/sh" "-c" "echo…   0B
612fd43926c1   6 minutes ago   /bin/sh -c #(nop)  EXPOSE 80                    0B
608c31487a05   6 minutes ago   /bin/sh -c yum -y install net-tools             28.4MB
321552774697   6 minutes ago   /bin/sh -c yum -y install vim                   66.3MB
7d152764648f   6 minutes ago   /bin/sh -c #(nop) WORKDIR /usr/local            0B
7fd901655801   6 minutes ago   /bin/sh -c #(nop)  ENV MYPATH=/usr/local        0B
18126e1e1649   6 minutes ago   /bin/sh -c #(nop)  MAINTAINER xbu<xbu@sonicw…   0B
5d0da3dc9764   3 months ago    /bin/sh -c #(nop)  CMD ["/bin/bash"]            0B
<missing>      3 months ago    /bin/sh -c #(nop)  LABEL org.label-schema.sc…   0B
<missing>      3 months ago    /bin/sh -c #(nop) ADD file:805cb5e15fb6e0bb0…   231MB
```

#### 2、Tomcat 镜像

##### 1、准备镜像文件

```bat
# 准备好 tomcat 和 jdk 压缩文件，编写 README
(base) [xbu@localhost tomcat_test]$ ls
apache-tomcat-9.0.56.tar.gz  jdk-11.0.12_linux-x64_bin.tar.gz  README
```

##### 2、编写 dockerfile

```bat
FROM centos
MAINTAINER xbu<xbu@sonicwall.com>
# 复制文件
COPY README /usr/local/README
# 复制解压
ADD jdk-11.0.12_linux-x64_bin.tar.gz /usr/local
ADD apache-tomcat-9.0.56.tar.gz /usr/local
RUN yum -y install vim
#设置环境变量
ENV MYPATH /usr/local
#设置工作目录
WORKDIR $MYPATH 
#设置环境变量
ENV JAVA_HOME /usr/local/jdk-11.0.12
ENV CLASSPATH $JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar 
ENV CATALINA_HOME /usr/local/apache-tomcat-9.0.56 
ENV CATALINA_BASH /usr/local/apache-tomcat-9.0.56 
#设置环境变量 分隔符是：
ENV PATH $PATH:$JAVA_HOME/bin:$CATALINA_HOME/lib:$CATALINA_HOME/bin 
#设置暴露的端口
EXPOSE 8080 
# 设置默认命令
CMD /usr/local/apache-tomcat-9.0.56/bin/startup.sh && tail -F /usr/local/apachetomcat-9.0.56/logs/catalina.out 
```

##### 3、构建镜像

```bat
# 因为dockerfile命名使用默认命名 因此不用使用-f 指定文件
$ docker build -t mytomcat:0.1 .
```

##### 4、run镜像

```bat
$ docker run -d -p 8080:8080 --name xbu-tomcat01 -v
/home/xbu/files/docker_test/build/tomcat/test:/usr/local/apache-tomcat-9.0.56/webapps/test -
v /home/xbu/files/docker_test/build/tomcat/tomcatlogs/:/usr/local/apache-tomcat-9.0.56/logs
mytomcat:0.1
```

##### 5、测试

在宿主机目录```/home/xbu/files/docker_test/build/tomcat/test```下编写 web.xml 和 index.jsp 文件

```xml
<web-app id="MyStrutsApp" version="2.4" 
         xmlns="http://java.sun.com/xml/ns/j2ee" 
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
         xsi:schemaLocation="http://java.sun.com/xml/ns/j2ee http://java.sun.com/xml/ns/j2ee/web-app_2_4.xsd">

    <filter>
        <filter-name>struts2</filter-name>
        <filter-class>org.apache.struts2.dispatcher.filter.StrutsPrepareAndExecuteFilter</filter-class>
    </filter>

    <filter-mapping>
        <filter-name>struts2</filter-name>
        <url-pattern>/*</url-pattern>
    </filter-mapping>

    <!-- ... -->

</web-app>
```

```xml
<?xml version="1.0" encoding="UTF-8"?>
<web-app xmlns="http://java.sun.com/xml/ns/javaee"
         xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
         xsi:schemaLocation="http://java.sun.com/xml/ns/javaee
                             http://java.sun.com/xml/ns/javaee/web-app_2_5.xsd"
         version="2.5">

</web-app>
```

```jsp
<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>hello, xbu</title>
</head>
<body>
Hello World!<br/>
<%
System.out.println("-------xbu tomcat test web logs--------");
%>
</body>
</html>
```

##### 6、项目发布

#### 3、构建自己的测试环境

##### 1、文件准备

```bat
# 包括 allure、jdk、README、Dockerfile
(base) sonicwall@sonicwall-virtual-machine:~/files/docker/devtest_env/dockerfile$ ls
allure-commandline-2.13.8.tgz  jdk-11.0.12_linux-x64_bin.tar.gz  README
```

##### 2、编写Dockerfile

###### 1、构建 pip 环境

```bat
FROM ubuntu
MAINTAINER xbu<xbu@sonicwall.com>
# 复制文件
COPY README /usr/local/README
COPY get-pip.py /root/
# 复制解压
ADD jdk-11.0.12_linux-x64_bin.tar.gz /usr/local
ADD allure-commandline-2.13.8.tgz /usr/local
ADD Python-3.9.6.tgz /usr/local
# 备份镜像源，更换镜像源
RUN mv /etc/apt/sources.list /etc/apt/sources_init.list
COPY sources.list /etc/apt/
RUN apt-get clean && apt-get update && apt-get upgrade

RUN apt-get install -y net-tools

RUN python3 /root/get-pip.py
#RUN apt-get install -y python-pip-whl
# RUN apt-get install -y python3
#RUN apt-get install -y python3-pip
#设置环境变量
ENV MYPATH /usr/local
#设置工作目录
WORKDIR $MYPATH 
#设置环境变量
ENV JAVA_HOME /usr/local/jdk-11.0.12
ENV CLASSPATH $JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar 
ENV ALLURE /usr/local/allure-2.13.8/bin/ 

#设置环境变量 分隔符是：
ENV PATH $PATH:$JAVA_HOME/bin:$ALLURE
#设置暴露的端口
EXPOSE 8081 
# 设置默认命令
CMD /bin/bash
```

###### 2、构建 pipenv 环境

```bat
FROM ubuntu
MAINTAINER xbu<xbu@sonicwall.com>
# 复制文件
COPY README /usr/local/README
COPY Pipfile /root/Pipfile

# 备份镜像源，更换镜像源
RUN mv /etc/apt/sources.list /etc/apt/sources_init.list
COPY sources.list /etc/apt/
RUN apt-get clean && apt-get update && apt-get upgrade

RUN apt-get install -y net-tools
RUN apt-get install -y vim

#设置环境变量
ENV MYPATH /usr/local
#设置工作目录
WORKDIR $MYPATH 

#设置暴露的端口
EXPOSE 8081 
# 设置默认命令
CMD /bin/bash
```

###### 3、构建 anaconda3 环境

```bat
FROM continuumio/anaconda3
MAINTAINER xbu<xbu@sonicwall.com>
# 复制文件
COPY README /usr/local/README

# 复制解压
ADD jdk-11.0.12_linux-x64_bin.tar.gz /usr/local
ADD allure-commandline-2.13.8.tgz /usr/local

# 备份镜像源，更换镜像源
RUN mv /etc/apt/sources.list /etc/apt/sources_init.list
COPY sources.list /etc/apt/
RUN apt-get clean && apt-get update && apt-get upgrade

RUN apt-get install -y net-tools

#设置环境变量
ENV MYPATH /usr/local
#设置工作目录
WORKDIR $MYPATH 
#设置环境变量
ENV JAVA_HOME /usr/local/jdk-11.0.12
ENV CLASSPATH $JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar 
ENV ALLURE /usr/local/allure-2.13.8/bin/ 

#设置环境变量 分隔符是：
ENV PATH $PATH:$JAVA_HOME/bin:$ALLURE
#设置暴露的端口
EXPOSE 8081 
# 设置默认命令
CMD /bin/bash
```

###### 4、构建 Python

```bat
FROM python:3.8
MAINTAINER xbu<xbu@sonicwall.com>
# 复制文件
COPY README /usr/local/README
COPY requirements.txt /root/requirements.txt
COPY pip/ /root/

# 备份镜像源，更换镜像源
RUN mv /etc/apt/sources.list /etc/apt/sources_init.list
COPY sources.list /etc/apt/
RUN apt-get clean && apt-get update && apt-get upgrade

RUN apt-get install -y vim
RUN pip install virtualenv
RUN pip install --user -r /root/requirements.txt

#设置环境变量
ENV MYPATH /usr/local
#设置工作目录
WORKDIR $MYPATH 

# 设置默认命令
CMD ["/bin/bash", "python", "server.py"]
```

###### 5、构建jenkins

创建本地数据卷

```bat
注意：这一步必须做
mkdir jenkins/jenkins_home
# 需要修改下目录权限，因为本地jenkins目录的拥有者为root用户，而容器中jenkins用户的 uid 为 1000
sudo chown -R 1000:1000 jenkins_home
```

```bat
FROM jenkins/jenkins:lts-jdk11
MAINTAINER xbu
# 如果不设置 USER root 就需要修改 jenkins_home 目录权限
USER root

# 复制解压
ADD allure-commandline-2.13.8.tgz /usr/local

# 将 debian 源更换为 阿里云源
RUN cp /etc/apt/sources.list /etc/apt/sources.list.init
RUN sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list
RUN apt-get update --fix-missing && apt-get install -y vim --fix-missing

RUN apt-get install -y python3-pip sudo openssh-server openssh-clients curl lsof
# 更换 pip 源，一旦RUN命令完成，这个目录就消失了
RUN mkdir -p /home/.pip && touch /home/.pip/pip.conf
# dockerfile 中要加 \ 不然无法识别为一行命令
RUN echo $'[global]\n \
\nindex-url = http://mirrors.aliyun.com/pypi/simple/\n \
\n[install]\n \
\ntrusted-host=mirrors.aliyun.com\n' \
>> /home/.pip/pip.conf
# 安装 python 环境
RUN pip install virtualenv
RUN mkdir -p /home/python_env && virtualenv /home/python_env/devtest

#设置环境变量
#ENV JAVA_HOME /opt/java/openjdk 
ENV ALLURE /usr/local/allure-2.13.8/bin/ 
#设置环境变量 分隔符是：
ENV PATH $PATH:$ALLURE

# 更换插件源
RUN sed -i 's/https:\/\/updates.jenkins.io\/download/http:\/\/mirrors.tuna.tsinghua.edu.cn\/jenkins/g' /var/jenkins_home/updates/default.json && sed -i 's/http:\/\/www.google.com/https:\/\/www.baidu.com/g' /var/jenkins_home/updates/default.json

# drop back to the regular jenkins user - good practice
#USER jenkins
#设置环境变量
ENV MYPATH /usr/local
#设置工作目录
WORKDIR /home 

# 设置默认命令
# CMD /bin/bash 不能用 /bin/bash，不然网页访问不了 jenkins
```

启动容器

```bat
# docker run 的时候不能用 /bin/bash，不然网页访问不了 jenkins
docker run -it --name myjenkins03 -p 7080:8080 -p 50000:50000 -p 7081:8081 -v /home/xbu/files/docker/myjenkins03/jenkins_home:/var/jenkins_home -v /home/xbu/files/docker/myjenkins03/files:/home --env JENKINS_SLAVE_AGENT_PORT=50001 jenkins/jenkins:lts-jdk11

docker run -d -it --name myjenkins01 -p 7080:8080 -p 50000:50000 -p 7081:8081 -v /home/xbu/files/docker/jenkins01/jenkins_home:/var/jenkins_home -v /home/xbu/files/docker/jenkins01/files:/home --env JENKINS_SLAVE_AGENT_PORT=50001 jenkins-xbu:0.1 
```

更换插件源

```bat
sed -i 's/https:\/\/updates.jenkins.io\/download/http:\/\/mirrors.tuna.tsinghua.edu.cn\/jenkins/g' /var/jenkins_home/updates/default.json && sed -i 's/http:\/\/www.google.com/https:\/\/www.baidu.com/g' /var/jenkins_home/updates/default.json
```

更换 pip 源

```bat
mkdir -p /home/xbu/.pip
tee /home/xbu/.pip/pip.conf <<-'EOF'
[global]
index-url = http://mirrors.aliyun.com/pypi/simple/
[install]
trusted-host=mirrors.aliyun.com
EOF
```

myjenkins01的dockerfile

```bat
FROM jenkins/jenkins
MAINTAINER xbu

USER root
# 将 debian 源更换为 阿里云源
RUN cp /etc/apt/sources.list /etc/apt/sources.list.init
RUN sed -i 's/deb.debian.org/mirrors.aliyun.com/g' /etc/apt/sources.list
RUN apt-get update --fix-missing && apt-get install -y vim --fix-missing

RUN apt-get install -y python3-pip
# 更换 pip 源
RUN mkdir -p /home/.pip
RUN touch /home/.pip/pip.conf
# dockerfile 中要加 \ 不然无法识别为一行命令
RUN echo $'[global]\n \
index-url = http://mirrors.aliyun.com/pypi/simple/\n \
[install]\n \
trusted-host=mirrors.aliyun.com\n' \
>> /home/.pip/pip.conf

RUN pip install virtualenv

#设置环境变量
#ENV MYPATH /usr/local
#设置工作目录
WORKDIR /home
```

###### 6、更换源时提示GPG error缺少公钥

```bat
# 出现错误
W: GPG error: http://mirrors.aliyun.com/ubuntu trusty-security InRelease: The following signatures couldn't be verified because the public key is not available: NO_PUBKEY 40976EAF437D05B5 NO_PUBKEY 3B4FE6ACC0B21F32
# 解决方法一：https://blog.csdn.net/qq_38889662/article/details/108205364
apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 3B4FE6ACC0B21F32（这个码是上面的错误提示的码）
# 网上解决方法一：https://blog.csdn.net/zs15yy/article/details/62892366
gpg --keyserver pgpkeys.mit.edu --recv-keys C857C906
gpg --export --armor C857C906 | sudo apt-key add -

gpg --keyserver keyserver.ubuntu.com --recv-keys 40976EAF437D05B5
gpg --export --armor 40976EAF437D05B5 | apt-key add -
```

#### 4、发布自己的镜像

##### 1、登陆

```bat
docker login --help
Usage: docker login [OPTIONS] [SERVER]
Log in to a Docker registry.
If no server is specified, the default is defined by the daemon.
Options:
    -p, --password string Password
    --password-stdin Take the password from stdin
    -u, --username string Username
```

##### 2、push 提交镜像

发现问题：push不上去？   因为如果没有前缀的话默认是push到 官方的library  

解决方法：

```bat
# 方法一
# 从新build， build的时候添加自己的dockerhub用户名，然后在push就可以放到自己的仓库了
docker build -t xbu/mytomcat:0.1 .
docker push xbu/mytomcat:0.1
# 方法二
# 使用docker tag 修改自己制作的镜像的标签，然后再次push
docker tag 容器id xbu/mytomcat:1.0
docker push xbu/mytomcat:0.1
```

##### 3、发布到阿里云服务器

###### 1、登陆阿里云

2、找到容器镜像服务

3、创建命名空间（一个账号只能创建 3 个命名空间）

4、创建镜像仓库

参考官方文档

```bat
# 看官网 很详细https://cr.console.aliyun.com/repository/
$ sudo docker login --username=zchengx registry.cn-shenzhen.aliyuncs.com
$ sudo docker tag [ImageId] registry.cn-shenzhen.aliyuncs.com/dsadxzc/cheng:[镜像版本号]
# 修改id 和 版本
sudo docker tag a5ef1f32aaae registry.cn-shenzhen.aliyuncs.com/dsadxzc/cheng:1.0
# 修改版本
$ sudo docker push registry.cn-shenzhen.aliyuncs.com/dsadxzc/cheng:[镜像版本号]
```

## Docker 网络

### 1、理解docker0

```bat
(base) xbu@sxlin-OptiPlex-7050:~$ ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
    inet6 ::1/128 scope host 
       valid_lft forever preferred_lft forever
2: enp1s0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc mq state UP group default qlen 1000
    link/ether b4:96:91:33:d8:96 brd ff:ff:ff:ff:ff:ff
    inet 192.168.168.67/24 brd 192.168.168.255 scope global noprefixroute enp1s0
       valid_lft forever preferred_lft forever
    inet6 fe80::371b:5e38:bc14:8c8e/64 scope link noprefixroute 
       valid_lft forever preferred_lft forever
3: enp0s31f6: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc fq_codel state UP group default qlen 1000
    link/ether 54:bf:64:90:93:33 brd ff:ff:ff:ff:ff:ff
    inet 10.103.12.238/23 brd 10.103.13.255 scope global noprefixroute enp0s31f6
       valid_lft forever preferred_lft forever
    inet6 2001:470:80b7:670c:a0aa:b616:2235:2915/64 scope global temporary dynamic 
       valid_lft 596603sec preferred_lft 77711sec
    inet6 2001:470:80b7:670c:44e7:9a39:c173:a82/64 scope global temporary deprecated dynamic 
       valid_lft 424792sec preferred_lft 0sec
    inet6 2001:470:80b7:670c:c0a9:ae70:4ccc:f377/64 scope global dynamic mngtmpaddr noprefixroute 
       valid_lft 2591914sec preferred_lft 604714sec
    inet6 fe80::fdcc:61eb:981a:910d/64 scope link noprefixroute 
       valid_lft forever preferred_lft forever
4: docker0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:8b:74:1c:e6 brd ff:ff:ff:ff:ff:ff
    inet 172.17.0.1/16 brd 172.17.255.255 scope global docker0
       valid_lft forever preferred_lft forever
    inet6 fe80::42:8bff:fe74:1ce6/64 scope link 
       valid_lft forever preferred_lft forever
14: veth1521c3d@if13: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default 
    link/ether 52:94:49:9e:cb:55 brd ff:ff:ff:ff:ff:ff link-netnsid 1
    inet6 fe80::5094:49ff:fe9e:cb55/64 scope link 
       valid_lft forever preferred_lft forever
16: veth1d5bc2e@if15: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue master docker0 state UP group default 
    link/ether 16:c1:73:07:0e:48 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet6 fe80::14c1:73ff:fe07:e48/64 scope link 
       valid_lft forever preferred_lft forever
```

本机回环地址

![image-20220424124938784](images/docker学习笔记.assets/image-20220424124938784.png)

内网地址

![image-20220424125307369](images/docker学习笔记.assets/image-20220424125307369.png)

Docker0 地址

每启动一个docker容器，docker就会给docker容器分配一个ip，我们只要安装了docker， 就会有一个网卡docker0 桥接模式，使用的技术是veth-pair技术！veth-pair 就是一对的虚拟设备接口，他们都是成对出现的，一端连着协议，一端彼此相连 。

![image-20220424234657102](images/docker学习笔记.assets/image-20220424234657102.png)

 Docker使用的是Linux的桥接，宿主机是一个Docker容器的网桥（docker0）。

![image-20220424233746377](images/docker学习笔记.assets/image-20220424233746377.png)

Docker中所有网络接口都是虚拟的，虚拟的转发效率高（内网传递文件）。

只要容器删除，对应的网桥一对就没了！

![image-20220425002426694](images/docker学习笔记.assets/image-20220425002426694.png)

### 2、-link

思考一个场景：我们编写了一个微服务，database url=ip: ，数据ip换了，我们希望项目不重启，可以通过名字来进行访问容器？

```bat
$ docker exec -it tomcat02 ping tomca01# ping不通 
ping: tomca01: Name or service not known 

# 运行一个tomcat03 --link tomcat02 
$ docker run -d -P --name tomcat03 --link tomcat02 tomcat 5f9331566980a9e92bc54681caaac14e9fc993f14ad13d98534026c08c0a9aef 
# 用tomcat03 ping tomcat02 可以ping通 
$ docker exec -it tomcat03 ping tomcat02 
PING tomcat02 (172.17.0.3) 56(84) bytes of data. 
64 bytes from tomcat02 (172.17.0.3): icmp_seq=1 ttl=64 time=0.115 ms 
64 bytes from tomcat02 (172.17.0.3): icmp_seq=2 ttl=64 time=0.080 ms 
# 发现用tomcat02 ping tomcat03 ping不通，因为没有配置 
```

```bat
(base) wirelessdev@wirelesss-MacBook-Pro ~ % docker network --help
Usage:  docker network COMMAND
Manage networks
Commands:
  connect     Connect a container to a network
  create      Create a network
  disconnect  Disconnect a container from a network
  inspect     Display detailed information on one or more networks
  ls          List networks
  prune       Remove all unused networks
  rm          Remove one or more networks
```

**探究：** docker network inspect 网络id 网段相同

![image-20220425002305067](images/docker学习笔记.assets/image-20220425002305067.png)

 docker inspect tomcat03

![image-20220425002931888](images/docker学习笔记.assets/image-20220425002931888.png)

查看 tomcat03 里面的 /etc/hosts 发现有 tomcat02 的配置

![image-20220425003049555](images/docker学习笔记.assets/image-20220425003049555.png)

**–link 本质就是在hosts配置中添加映射**

现在使用Docker已经不建议使用–link了！自定义网络，不使用docker0！

docker0问题：不支持容器名连接访问！

### 3、自定义网络（docker network create）

#### 1、查看所有的 docker 网络

```bat
# docker network ls 查看所有的网络
(base) wirelessdev@wirelesss-MacBook-Pro ~ % docker network ls    
NETWORK ID     NAME      DRIVER    SCOPE
a2a3e350af08   bridge    bridge    local
3e98cb5a2419   host      host      local
c657fc71e864   none      null      local
```

#### 2、网络模式

bridge ：桥接 docker（默认，自己创建也是用bridge模式）

none ：不配置网络，一般不用

host ：和所主宿主机共享网络

container ：容器网络连通（用得少！局限很大）

#### 3、测试几种网络模式

```bat
# 我们直接启动的命令默认 --net bridge,而这个就是我们得docker0 
# bridge就是docker0 
$ docker run -d -P --name tomcat01 tomcat 等价于 => 
docker run -d -P --name tomcat01 --net bridge tomcat 
# docker0，特点：默认，域名不能访问。 --link可以打通连接，但是很麻烦！ 

# 我们可以 自定义一个网络 
# --driver bridge
# --subnet 192.168.0.0/16（可以支持从192.168.0.2 到 192.168.255.255）
# --gateway 192.168.0.1
$ docker network create --driver bridge --subnet 192.168.0.0/16 --gateway 192.168.0.1 mynet
docker network create --driver bridge --subnet 192.168.100.0/24 --gateway 192.168.100.1 tele-net
```

```bat
(base) wirelessdev@wirelesss-MacBook-Pro ~ % docker network create --driver bridge --subnet 192.168.0.0/16 --gateway 192.168.0.1 mynet
7ce9e7d4a3954dbf527f05965fe00b2d271a233c7b1ce1ae6fa18550a678adc2
(base) wirelessdev@wirelesss-MacBook-Pro ~ % docker network ls
NETWORK ID     NAME      DRIVER    SCOPE
82386a4f8fca   bridge    bridge    local
3e98cb5a2419   host      host      local
7ce9e7d4a395   mynet     bridge    local
c657fc71e864   none      null      local
```

**docker network inspect mynet 查看创建的网络**

![image-20220425142407026](images/docker学习笔记.assets/image-20220425142407026.png)

启动两个tomcat,再次查看网络情况

![image-20220425143543150](images/docker学习笔记.assets/image-20220425143543150.png)

![image-20220425143714873](images/docker学习笔记.assets/image-20220425143714873.png)

**在自定义的网络下，服务可以通过容器名互相ping通，不用使用–link**

![image-20220425143825539](images/docker学习笔记.assets/image-20220425143825539.png)

我们自定义的网络docker都已经帮我们维护好了对应的关系，推荐我们平时这样使用网络！

好处：

redis、mysql  -不同的集群使用不同的网络，保证集群是安全和健康的

![image-20220425144120376](images/docker学习笔记.assets/image-20220425144120376.png)

### 4、网络连通

两个不同网段的网络，是不能相互 ping 通的。如何将docker0的容器连接到自己创建的网络？

![image-20220425144855474](images/docker学习笔记.assets/image-20220425144855474.png)

![image-20220425145058849](images/docker学习笔记.assets/image-20220425145058849.png)

![image-20220425145158369](images/docker学习笔记.assets/image-20220425145158369.png)

要将tomcat01 连通 tomcat-net-01 ，就是将 tomcat01加到 mynet网络下，此时 tomcat01又多了一个192.168.0.4/16的 IP，实现一个容器两个ip地址（172.18.0.1、192.168.0.4/16）

例如阿里云服务也是两个IP：公网IP、私网IP

![image-20220425145619584](images/docker学习笔记.assets/image-20220425145619584.png)

### 5、实战：部署Redis集群

![image-20220508160943051](images/docker学习笔记.assets/image-20220508160943051.png)

```shell
# 创建网卡 
docker network create redis --subnet 172.38.0.0/16 

# 通过脚本创建六个redis配置 
for port in $(seq 1 6);\ 
do \ 
mkdir -p /mydata/redis/node-${port}/conf 
touch /mydata/redis/node-${port}/conf/redis.conf 
cat << EOF >> /mydata/redis/node-${port}/conf/redis.conf 
port 6379 
bind 0.0.0.0 
cluster-enabled yes 
cluster-config-file nodes.conf 
cluster-node-timeout 5000 
cluster-announce-ip 172.38.0.1${port} 
cluster-announce-port 6379 
cluster-announce-bus-port 16379 
appendonly yes
EOF 
done

# 通过脚本运行六个redis 
for port in $(seq 1 6);\ 
docker run -p 637${port}:6379 -p 1667${port}:16379 --name redis-${port} \ 
-v /mydata/redis/node-${port}/data:/data \ 
-v /mydata/redis/node-${port}/conf/redis.conf:/etc/redis/redis.conf \ 
-d --net redis --ip 172.38.0.1${port} redis:5.0.9-alpine3.11 redis-server /etc/redis/redis.conf 
docker exec -it redis-1 /bin/sh #redis默认没有bash 

# 创建集群
redis-cli --cluster create 172.38.0.11:6379 172.38.0.12:6379 172.38.0.13:6379 172.38.0.14:6379 172.38.0.15:6379 172.38.0.16:6379 --cluster-replicas 1
```

# 二、Docker进阶

## 1、docker compose

### 1、简介

docker compose 轻松高效的管理容器，定义运行多个容器。

> 官方介绍

Compose is a tool for defining and running multi-container Docker applications. With Compose, you use a YAML file to configure your application’s services. Then, with a single command, you create and start all the services from your configuration. To learn more about all the features of Compose, see [the list of features](https://docs.docker.com/compose/#features).

Compose works in all environments: production, staging, development, testing, as well as CI workflows. You can learn more about each case in [Common Use Cases](https://docs.docker.com/compose/#common-use-cases).

Using Compose is basically a three-step process:

1. Define your app’s environment with a `Dockerfile` so it can be reproduced anywhere.
2. Define the services that make up your app in `docker-compose.yml` so they can be run together in an isolated environment.
3. Run `docker compose up` and the [Docker compose command](https://docs.docker.com/compose/#compose-v2-and-the-new-docker-compose-command) starts and runs your entire app. You can alternatively run `docker-compose up` using the docker-compose binary.

> 理解

compose 是 docker 官方的开源项目，需要安装。

Dockerfile 让程序在任何地方运行

A `docker-compose.yml` looks like this:

```yaml
version: "3.9"  # optional since v1.27.0
services:
  web:
    build: .
    ports:
      - "8000:5000"
    volumes:
      - .:/code
      - logvolume01:/var/log
    links:
      - redis # web服务连接到redis服务上
  redis:
    image: redis
volumes:
  logvolume01: {}
```

compose 的重要概念：

- 服务、容器、应用。
- 项目：一组关联的容器。例如一个简单的博客，有web、mysql服务

### 2、安装

官网链接：https://docs.docker.com/compose/install/compose-plugin/#install-using-the-repository

您可以将Compose作为独立二进制文件使用，而无需安装Docker CLI。

**1- Run this command to download the current stable release of Docker Compose:**

```shell
 # vcenter 中的虚拟机只能使用官网的这个地址慢慢下载,公司的网络还需要加上代理才行
 $ sudo curl -SL https://github.com/docker/compose/releases/download/v2.5.0/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose -x http://10.50.128.110:3128

  # 使用国内镜像下载
curl -L https://get.daocloud.io/docker/compose/releases/download/1.25.5/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose

curl -L https://get.daocloud.io/docker/compose/releases/download/1.29.2/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose

curl -L "https://get.daocloud.io/docker/compose/releases/download/2.5.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```

> 若要安装其他版本的Compose，请将“v2.5.0”替换为要使用的Compose版本。

**2- Apply executable permissions to the binary:**

```
  $ sudo chmod +x /usr/local/bin/docker-compose
```

> **Note**:
> 
> If the command `docker-compose` fails after installation, check your path. You can also create a symbolic link to `/usr/bin` or any other directory in your path.
> 
> For example:
> 
> ```
> $ sudo ln -s /usr/local/bin/docker-compose /usr/bin/docker-compose
> ```

**3- 测试安装结果**

```
 $ docker-compose --version
 Docker Compose version v2.5.0
```

### 3、体验

#### 具体步骤：

1. 写应用app.py。
2. Dockerfile 将应用打包为镜像。
3. 写一个Docker-compose yaml文件，文件定义整个服务需要的环境。
4. 启动 compose 项目 （docker-compose up）

执行`docker-compose up` 命令后操作流程：

1- 创建网络。

2- 执行docker-compose.yml 文件

3- 启动服务。

官网链接：https://docs.docker.com/compose/gettingstarted/

在此页面上，您构建了一个在Docker Compose上运行的简单Python web应用程序。该应用程序使用Flask框架，并在Redis中维护一个命中计数器。

#### 1、定义应用依赖

1. Create a directory for the project:
   
   ```
   $ mkdir composetest
   $ cd composetest
   ```

2. Create a file called `app.py` in your project directory and paste this in:
   
   ```python
   import time
   
   import redis
   from flask import Flask
   
   app = Flask(__name__)
   cache = redis.Redis(host='redis', port=6379)
   
   def get_hit_count():
       retries = 5
       while True:
           try:
               return cache.incr('hits')
           except redis.exceptions.ConnectionError as exc:
               if retries == 0:
                   raise exc
               retries -= 1
               time.sleep(0.5)
   
   @app.route('/')
   def hello():
       count = get_hit_count()
       return 'Hello World! I have been seen {} times.\n'.format(count)
   ```
   
   在本例中，“redis”是应用程序网络上redis容器的主机名。我们使用Redis的默认端口“6379”。
   
   > **Handling transient errors**
   > 
   > 请注意 get_hit_count 函数的编写方式。如果redis服务不可用，此基本重试循环允许我们多次尝试请求。这在应用程序联机时启动时很有用，但如果Redis服务在应用程序的生命周期内需要随时重新启动，也会使我们的应用程序更有弹性。在集群中，这也有助于处理节点之间的瞬时连接中断。

3、Create another file called `requirements.txt` in your project directory and paste this in:

```
flask
redis
```

#### 2、创建 Dockerfile

在这一步中，您将编写一个Docker文件来构建Docker映像。该映像包含Python应用程序所需的所有依赖项，包括Python本身。

在项目目录中，创建一个名为“Dockerfile”的文件并粘贴以下内容：

```shell
# syntax=docker/dockerfile:1
FROM python:3.7-alpine
WORKDIR /code
ENV FLASK_APP=app.py
ENV FLASK_RUN_HOST=0.0.0.0
RUN apk add --no-cache gcc musl-dev linux-headers
COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt
EXPOSE 5000
COPY . .
CMD ["flask", "run"]
```

This tells Docker to:

- Build an image starting with the Python 3.7 image.
- Set the working directory to `/code`.
- Set environment variables used by the `flask` command.
- Install gcc and other dependencies
- Copy `requirements.txt` and install the Python dependencies.
- Add metadata to the image to describe that the container is listening on port 5000
- Copy the current directory `.` in the project to the workdir `.` in the image.
- Set the default command for the container to `flask run`.

#### 3、在docker compose file 中定义服务

在项目目录中创建一个名为 docker compose.yml 的文件，并粘贴以下内容：

```yaml
version: "3.9"
services:
  web:
    build: .
    ports:
      - "8000:5000"
  redis:
    image: "redis:alpine"
```

This Compose file defines two services: `web` and `redis`.

Web service[🔗](https://docs.docker.com/compose/gettingstarted/#web-service)

“web”服务使用从当前目录中的“Dockerfile”生成的镜像。然后，它将容器和主机绑定到公开的端口“8000”。此示例服务使用Flask web服务器的默认端口“5000”。

Redis service

The `redis` service uses a public [Redis](https://registry.hub.docker.com/_/redis/) image pulled from the Docker Hub registry.

#### 4. 使用Compose构建并运行应用程序

从项目目录中，通过运行 `docker-compose up`  启动应用程序.

```sell
$ docker-compose up

Creating network "composetest_default" with the default driver
Creating composetest_web_1 ...
Creating composetest_redis_1 ...
Creating composetest_web_1
Creating composetest_redis_1 ... done
Attaching to composetest_web_1, composetest_redis_1
web_1    |  * Running on http://0.0.0.0:5000/ (Press CTRL+C to quit)
redis_1  | 1:C 17 Aug 22:11:10.480 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
redis_1  | 1:C 17 Aug 22:11:10.480 # Redis version=4.0.1, bits=64, commit=00000000, modified=0, pid=1, just started
redis_1  | 1:C 17 Aug 22:11:10.480 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
web_1    |  * Restarting with stat
redis_1  | 1:M 17 Aug 22:11:10.483 * Running mode=standalone, port=6379.
redis_1  | 1:M 17 Aug 22:11:10.483 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.
web_1    |  * Debugger is active!
redis_1  | 1:M 17 Aug 22:11:10.483 # Server initialized
redis_1  | 1:M 17 Aug 22:11:10.483 # WARNING you have Transparent Huge Pages (THP) support enabled in your kernel. This will create latency and memory usage issues with Redis. To fix this issue run the command 'echo never > /sys/kernel/mm/transparent_hugepage/enabled' as root, and add it to your /etc/rc.local in order to retain the setting after a reboot. Redis must be restarted after THP is disabled.
web_1    |  * Debugger PIN: 330-787-903
```

#### 5. 启动服务后的内容详解

```shell
# 依次创建了网络和两个服务
Use 'docker scan' to run Snyk tests against images to find vulnerabilities and learn how to fix them
[+] Running 3/3
 ⠿ Network composetest_default    Created                                               0.2s
 ⠿ Container composetest-redis-1  Created                                               0.2s
 ⠿ Container composetest-web-1    Created                                               0.2s
Attaching to composetest-redis-1, composetest-web-1


# 服务已经运行起来了
composetest-redis-1  | 1:M 26 May 2022 07:42:53.955 * Ready to accept connections
composetest-web-1    |  * Serving Flask app 'app.py' (lazy loading)
composetest-web-1    |  * Environment: production
composetest-web-1    |    WARNING: This is a development server. Do not use it in a production deployment.
composetest-web-1    |    Use a production WSGI server instead.
composetest-web-1    |  * Debug mode: off
composetest-web-1    |  * Running on all addresses (0.0.0.0)
composetest-web-1    |    WARNING: This is a development server. Do not use it in a production deployment.
composetest-web-1    |  * Running on http://127.0.0.1:5000
composetest-web-1    |  * Running on http://172.18.0.2:5000 (Press CTRL+C to quit)


# 访问一下 web 的地址，注意容器内的端口是5000，映射到主机的端口是8000
xbu@sonicwall-virtual-machine:~$ curl http://127.0.0.1:5000
curl: (7) Failed to connect to 127.0.0.1 port 5000: Connection refused
xbu@sonicwall-virtual-machine:~$ curl http://172.18.0.2:5000
Hello World! I have been seen 1 times.

xbu@sonicwall-virtual-machine:~$ curl http://0.0.0.0:8000
Hello World! I have been seen 6 times.
xbu@sonicwall-virtual-machine:~$ curl localhost:8000
Hello World! I have been seen 7 times.
```

查看镜像，镜像中多了 redis 和运行 app.py 的 composetest_web 镜像

![image-20220526173604902](images/docker学习笔记.assets/image-20220526173604902.png)

使用 `docker service ls` 查看服务，提示此节点不是swarm 管理。

```shell
xbu@sonicwall-virtual-machine:~$ docker service ls
Error response from daemon: This node is not a swarm manager. Use "docker swarm init" or "docker swarm join" to connect this node to swarm and try again.
```

默认的服务名是 **文件名-服务名-num**，-num 表示的是副本数量。集群状态下，服务会有多个运行实例。

使用 `docker network ls` 查看网络，多了一个 composetest_default 网络。只要是使用 docker compose 启动，就会自动给这个应用创建一个网络。项目中的服务都会在这个网络下，这些服务可以通过域名互相访问。

![image-20220526173739070](images/docker学习笔记.assets/image-20220526173739070.png)

通过 `docker network inspect composetest_default` 查看这个网络，就可以看到两个服务都是在网络中。

![image-20220526174428477](images/docker学习笔记.assets/image-20220526174428477.png)

web服务在写应用的代码中 host 就写为了 redis 域名，只有两个服务在同一个网络下，app.py的服务才能通过**redis** 域名去访问另一个Redis服务，不然是ping 不通的，也就访问不了。

![image-20220526174826373](images/docker学习笔记.assets/image-20220526174826373.png)

#### 6.  停止应用程序

通过在第二个终端的 *<u>项目目录</u>* 中运行“`docker compose down`”，或在启动应用程序的原始终端中按CTRL+C，停止应用程序。

![image-20220526184412440](images/docker学习笔记.assets/image-20220526184412440.png)

#### 7. 编辑Compose文件以添加绑定挂载

Edit `docker-compose.yml` in your project directory to add a [bind mount](https://docs.docker.com/storage/bind-mounts/) for the `web` service:

```
version: "3.9"
services:
  web:
    build: .
    ports:
      - "8000:5000"
    volumes:
      - .:/code
    environment:
      FLASK_ENV: development
  redis:
    image: "redis:alpine"
```

新的“volumes”键将主机上的项目目录（当前目录）装载到容器内的“/code”，允许您动态修改代码，而无需重建映像。“environment”键设置“FLASK_ENV”环境变量，该变量告诉“`flask run`”在开发模式下运行，并在更改时重新加载代码。此模式只能用于开发。

#### 8. 使用Compose重新构建并运行应用程序

在项目目录中，键入“`docker compose up`”以使用更新的 compose 文件构建应用程序，然后运行它。

```
$ docker-compose up

Creating network "composetest_default" with the default driver
Creating composetest_web_1 ...
Creating composetest_redis_1 ...
Creating composetest_web_1
Creating composetest_redis_1 ... done
Attaching to composetest_web_1, composetest_redis_1
web_1    |  * Running on http://0.0.0.0:5000/ (Press CTRL+C to quit)
...
```

再次在web浏览器中检查“Hello World”消息，然后刷新以查看计数增量。

#### 9. 其他命令

如果要在后台运行服务，可以使用 “`docker compose up -d`” （用于“分离”模式），并使用“docker compose ps”查看当前正在运行的内容：

```shell
$ docker-compose up -d

Starting composetest_redis_1...
Starting composetest_web_1...

$ docker-compose ps

       Name                      Command               State           Ports         
-------------------------------------------------------------------------------------
composetest_redis_1   docker-entrypoint.sh redis ...   Up      6379/tcp              
composetest_web_1     flask run                        Up      0.0.0.0:8000->5000/tcp
```

“`docker compose run`”命令允许您为您的服务运行一次性命令。例如，要查看哪些环境变量可用于“web”服务，请执行以下操作：

```
$ docker-compose run web env
```

See `docker-compose --help` to see other available commands.

If you started Compose with `docker-compose up -d`, stop your services once you’ve finished with them:

```
$ docker-compose stop
```

You can bring everything down, removing the containers entirely, with the `down` command. Pass `--volumes` to also remove the data volume used by the Redis container:

```
$ docker-compose down --volumes
```

#### 10.  docker-compose.yaml 编写规则

>  docker-compose.yaml 规则

官网文档链接：https://docs.docker.com/compose/compose-file/compose-file-v3/

**一共只有三层**

```yaml
# 第一层：版本
version: '3.0'
# 第二层：服务
service:
    服务1: web
        # 服务配置
        images
        build
        network
        ......
    服务2: redis
        ......
    服务3: redis
        ......
# 第三层：其他配置：网络、数据集、全局规则
volumes:
networks:
configs:
```

`depends_on` 依赖：

![image-20220530115126035](images/docker学习笔记.assets/image-20220530115126035.png)

**开源项目，搭建博客：Quickstart: Compose and WordPress**

官方文档：https://docs.docker.com/samples/wordpress/

You can name the directory something easy for you to remember. This directory is the context for your application image. 目录应只包含用于构建该映像的资源。

1. Change into your project directory.
   
   For example, if you named your directory `my_wordpress`:
   
   ```
   $ cd my_wordpress/
   ```

2. Create a `docker-compose.yml` file that starts your `WordPress` blog and a separate `MySQL` instance with volume mounts for data persistence:
   
   ```yaml
   version: "3.9"
   
   services:
     db:
       image: mysql:5.7
       volumes:
         - db_data:/var/lib/mysql
       restart: always
       environment:
         MYSQL_ROOT_PASSWORD: somewordpress
         MYSQL_DATABASE: wordpress
         MYSQL_USER: wordpress
         MYSQL_PASSWORD: wordpress
   
     wordpress:
       depends_on:
         - db
       image: wordpress:latest
       volumes:
         - wordpress_data:/var/www/html
       ports:
         - "8000:80"
       restart: always
       environment:
         WORDPRESS_DB_HOST: db
         WORDPRESS_DB_USER: wordpress
         WORDPRESS_DB_PASSWORD: wordpress
         WORDPRESS_DB_NAME: wordpress
   volumes:
     db_data: {}
     wordpress_data: {}
   ```

后台启动：run `docker-compose up -d` from your project directory. 这将以分离模式运行docker compose，提取所需的docker镜像，并启动wordpress和数据库容器。

## 2、docker swarm

### 1. 购买服务器

在阿里云购买4台服务器

### 2. 安装docker

### 3. 工作模式

官方文档链接：https://docs.docker.com/engine/swarm/how-swarm-mode-works/nodes/

Docker Engine 1.12 introduces swarm mode that enables you to create a cluster of one or more Docker Engines called a swarm. A swarm consists of one or more nodes: physical or virtual machines running Docker Engine 1.12 or later in swarm mode.

There are two types of nodes: [**managers**](https://docs.docker.com/engine/swarm/how-swarm-mode-works/nodes/#manager-nodes) and [**workers**](https://docs.docker.com/engine/swarm/how-swarm-mode-works/nodes/#worker-nodes).

![image-20220531152749071](images/docker学习笔记.assets/image-20220531152749071.png)

## docker stack

## docker secret

## docker config

## k8s

# docker技巧

### 1、给已经启动的容器增加暴露端口

#### 1.停止容器（特别重要）

在执行以下步骤之前 一定要先停止容器

```bat
docker stop 容器id/容器名称
```

#### 2.在宿主机修改容器配置文件

配置文件的位置如下，一共要修改两个配置文件

```bat
/var/lib/docker/containers/[hash_of_the_container]/hostconfig.json  
/var/lib/docker/containers/[hash_of_the_container]/config.v2.json 
```

路径里面的[hash_of_the_container]通过下面命令查看

```
docker inspect 容器ID/容器名称
```

首先修改hostconfig.json如我的地址

```bat
vim /var/lib/docker/containers/[hash_of_the_container]/hostconfig.json
# 如下：
{"Binds":["/srv/gitlab-runner/config:/etc/gitlab-runner","/var/run/docker.sock:/var/run/docker.sock"],"ContainerIDFile":"","LogConfig":{"Type":"json-file","Config":{}},"NetworkMode":"default","PortBindings":{},"RestartPolicy":{"Name":"always","MaximumRetryCount":0},"AutoRemove":false,"VolumeDriver":"","VolumesFrom":null,"CapAdd":null,"CapDrop":null,"CgroupnsMode":"host","Dns":[],"DnsOptions":[],"DnsSearch":[],"ExtraHosts":null,"GroupAdd":null,"IpcMode":"private","Cgroup":"","Links":null,"OomScoreAdj":0,"PidMode":"","Privileged":false,"PublishAllPorts":false,"ReadonlyRootfs":false,"SecurityOpt":null,"UTSMode":"","UsernsMode":"","ShmSize":67108864,"Runtime":"runc","ConsoleSize":[0,0],"Isolation":"","CpuShares":0,"Memory":0,"NanoCpus":0,"CgroupParent":"","BlkioWeight":0,"BlkioWeightDevice":[],"BlkioDeviceReadBps":null,"BlkioDeviceWriteBps":null,"BlkioDeviceReadIOps":null,"BlkioDeviceWriteIOps":null,"CpuPeriod":0,"CpuQuota":0,"CpuRealtimePeriod":0,"CpuRealtimeRuntime":0,"CpusetCpus":"","CpusetMems":"","Devices":[],"DeviceCgroupRules":null,"DeviceRequests":null,"KernelMemory":0,"KernelMemoryTCP":0,"MemoryReservation":0,"MemorySwap":0,"MemorySwappiness":null,"OomKillDisable":false,"PidsLimit":null,"Ulimits":null,"CpuCount":0,"CpuPercent":0,"IOMaximumIOps":0,"IOMaximumBandwidth":0,"MaskedPaths":["/proc/asound","/proc/acpi","/proc/kcore","/proc/keys","/proc/latency_stats","/proc/timer_list","/proc/timer_stats","/proc/sched_debug","/proc/scsi","/sys/firmware"],"ReadonlyPaths":["/proc/bus","/proc/fs","/proc/irq","/proc/sys","/proc/sysrq-trigger"]}
```

我们找到**PortBindings**关键字（如果没有就增加），在该节点里面仿照形式再添加一组端口如我添加的

```bat
"9000/tcp":[{"HostIp":"","HostPort":"9002"}],"8080/tcp":[{"HostIp":"","HostPort":"8088"}]
```

再修改配置文件config.v2.json暴露端口

```
vim /var/lib/docker/containers/[hash_of_the_container]/config.v2.json
```

在config节点里面增加一个 （如果没有ExposedPorts的话），如果存在了，直接在ExposedPorts里面仿照已有格式暴露需要暴露的端口

```bat
### 增加需要额外暴露的 9000 和 8080 端口 
"ExposedPorts":{"9000/tcp":{},"8080/tcp":{}}
```

增加后的效果

![image-20220216143535929](images/docker学习笔记.assets/image-20220216143535929.png)

#### 3.重启docker服务（在这之前不要启动容器）

如果改动文件后启动了容器，那么你的配置又会被刷新为之前的配置了（你只好从第一步再来一遍）

```undefined
service docker restart
```

#### 4.启动容器

注意：此时虽然在 **docker ps -a** 看不到新增的暴露端口，但是通过配置文件和实际访问均可验证配置成功了

为什么我的这个启动后可以看到新增的暴露端口？而且重启docker服务后自动就启动了这个容器？

```bat
root@sxlin-OptiPlex-7050:~# service docker restart
root@sxlin-OptiPlex-7050:~# docker ps
CONTAINER ID   IMAGE                  COMMAND                  CREATED      STATUS              PORTS                                                                                  NAMES
f6a6f54f2ede   gitlab/gitlab-runner   "/usr/bin/dumb-init …"   7 days ago   Up About a minute   0.0.0.0:8088->8080/tcp, :::8088->8080/tcp, 0.0.0.0:9002->9000/tcp, :::9002->9000/tcp   gitlab-runner
```

### 2、将一台服务器上的docker容器打包发到另一台服务器

#### 步骤 1 ：打包

##### 方法 1 ：直接打包容器（import/export）

使用 docker export 命令根据容器 ID 将镜像导出成一个文件

```powershell
docker export 容器ID > 文件名.tar
```

 使用 docker import 命令导入

```bat
docker import 文件名.tar 目标镜像名:[TAG]
```

##### 方法 2 ：将容器制作为镜像后打包

使用 docker commit 命令制作镜像

```bat
docker commit -m="描述信息" -a="作者" 容器id 目标镜像名:[TAG]
```

使用 docker save 命令打包

```bat
docker save -o 镜像文件.tar  目标镜像名:[TAG]
```

使用 docker load 命令导入

```bat
docker load < 镜像文件.tar
```

##### 两种方法区别

```bat
1，文件大小不同
export 导出的镜像文件体积小于 save 保存的镜像

2，是否可以对镜像重命名
docker import 可以为镜像指定新名称
docker load 不能对载入的镜像重命名

3，是否可以同时将多个镜像打包到一个文件中
docker export 不支持
docker save 支持

4，是否包含镜像历史
export 导出（import 导入）是根据容器拿到的镜像，再导入时会丢失镜像所有的历史记录和元数据信息（即仅保存容器当时的快照状态），所以无法进行回滚操作。
而 save 保存（load 加载）的镜像，没有丢失镜像的历史，可以回滚到之前的层（layer）。

5，应用场景不同
docker export 的应用场景：主要用来制作基础镜像，比如我们从一个 ubuntu 镜像启动一个容器，然后安装一些软件和进行一些设置后，使用 docker export 保存为一个基础镜像。然后，把这个镜像分发给其他人使用，比如作为基础的开发环境。
docker save 的应用场景：如果我们的应用是使用 docker-compose.yml 编排的多个镜像组合，但我们要部署的客户服务器并不能连外网。这时就可以使用 docker save 将用到的镜像打个包，然后拷贝到客户服务器上使用 docker load 载入。
```

#### 步骤 2：发送

MobaXterm 本地中转，使用scp或curl命令

```bat
python3 -m http.server 端口 #（在原镜像文件服务器上先起一个server）
curl http://ip:port/原文件路径/文件名 -o /目标路径/文件名 # 在目标服务器上发送下载请求，可以修改文件名

scp 用户名@ip:/原文件路径/文件名 /目标路径 # 直接传输到目标目录下，不能修改文件名
```

### 3. 设置docker开机自启动，并设置容器自动重启

链接：https://blog.csdn.net/chj_1224365967/article/details/109029856

#### 1、设置[docker](https://so.csdn.net/so/search?q=docker&spm=1001.2101.3001.7020)开机启动

```bash
systemctl enable docker
```

#### 2、设置容器自动重启

##### 1）创建容器时设置

```shell
docker run -d --restart=always --name 设置容器名 使用的镜像
（上面命令  --name后面两个参数根据实际情况自行修改）

# Docker 容器的重启策略如下：
 --restart具体参数值详细信息：
       no　　　　　　　 // 默认策略,容器退出时不重启容器；
       on-failure　　  // 在容器非正常退出时（退出状态非0）才重新启动容器；
       on-failure:3    // 在容器非正常退出时重启容器，最多重启3次；
       always　　　　  // 无论退出状态是如何，都重启容器；
       unless-stopped  // 在容器退出时总是重启容器，但是不考虑在 Docker 守护进程启动时就已经停止了的容器。
```

##### 2）修改已有容器，使用update

如果创建时未指定 --restart=always，可通过update 命令设置

```shell
docker update --restart=always 容器ID(或者容器名)
```

### 4. Docker磁盘占用与清理问题

docker system prune后可以加额外的参数，如：

```shell
docker system prune -a # 一并清除所有未被使用的镜像和悬空镜像。
docker system prune -f # 用以强制删除，不提示信息。
```

对于悬空镜像和未使用镜像可以使用手动进行个别删除：
1、删除所有悬空镜像，不删除未使用镜像：
`docker rmi $(docker images -f “dangling=true” -q)`
2、删除所有未使用镜像和悬空镜像
`docker rmi $(docker images -q)`
3、清理卷
如果卷占用空间过高，可以清除一些不使用的卷，包括一些未被任何容器调用的卷（-v 详细信息中若显示 LINKS = 0，则是未被调用）：

```shell
# 删除所有未被容器引用的卷：
docker volume rm $(docker volume ls -qf dangling=true)
```

4、容器清理
如果发现是容器占用过高的空间，可以手动删除一些：

```shell
# 删除所有已退出的容器：
docker rm -v $(docker ps -aq -f status=exited)

# 删除所有状态为dead的容器
docker rm -v $(docker ps -aq -f status=dead)
```

### 5. Docker-compose 重建的容器内的代码没有更新

主要是挂载文件的时候把容器内的文件挂载在宿主机，而重新建容器的时候挂载的是之前的目录，目录内的内容不会被新建容器重写，反而是`挂载目录内的文件内容会覆盖容器内被挂载的文件`，所以会导致新建容器内容还是旧的。

```dockerfile
# docker-compose.yml
version: '3.0'
services:

    xbu_api:
        restart: "always"
        image: xbu_api
        container_name: xbuApi
        # tty: true     # 给容器设置一个伪终端防止进程结束容器退出
        build:
            context: ../code/xbuApi
            dockerfile: $PWD/xbuApi/Dockerfile
        volumes:
            - xbu_api_data:/app
        ports:
            - "1993:80"
#        networks:
#            - xbu_api_net
        environment:
            FLASK_ENV: development

volumes:
    xbu_api_data:
        driver: local
        driver_opts:
            type: local
            device: /srv/xbuApi/app # 也可以挂载目录只挂载log文件，这样不会影响新建容器的文件
            o: bind

#networks:
#  xbu_api_net:
#    driver: bridge
#    ipam:
#      config:
#        - subnet: 169.254.100.0/24
```

解决方法是每次`docker-compose up -d xbu_api`新建容器前，把之前挂载目录删除，然后重建挂载目录。

```shell
#!/bin/bash


#LOG_DIR="/opt/xbuApiLog"
#mkdir -p $LOG_DIR/xbuApi

APP_DIR="/srv/xbuApi/app"

sudo rm -rf $APP_DIR  # 删除挂载目录
sudo mkdir -p $APP_DIR # 新建挂载目录

#docker-compose up -d --remove-orphans xbu_api
docker-compose up -d xbu_api

#./run-agent.sh -s 2cb8ed694e20 -a http://10.103.12.69:8004 -n agent1
```

### 6. Error response from daemon

遇到的错误：Error response from daemon: Get "https://registry-1.docker.io/v2/": net/http: request canceled while waiting for connection (Client.Timeout exceeded while awaiting headers)

参考链接：https://askubuntu.com/questions/1400464/docker-error-response-from-daemon-get-https-registry-1-docker-io-v2-net-ht

```bash
sudo mkdir -p /etc/systemd/system/docker.service.d
# 创建代理配置
sudo vim /etc/systemd/system/docker.service.d/http-proxy.conf

# 代理文件内容
[Service]
Environment="HTTP_PROXY=http://10.50.128.110:3128/"
Environment="HTTPS_PROXY=http://10.50.128.110:3128/"
Environment="NO_PROXY=localhost,127.0.0.1,registry.xbu.dev"

# 重启
sudo systemctl daemon-reload
sudo  systemctl restart docker
```

还有些其他方法，没有尝试：

```bash
cat ~/.docker/config.json 
{
 "proxies":
 {
   "default":
   {
     "httpProxy": "http://webproxy.bu.edu:8900",
     "httpsProxy": "http://webproxy.bu.edu:8900",
     "noProxy": "localhost,127.0.0.1,.bu.edu,.ad.bu.edu,128.197.,10."
   }
 }
}
```

