### DS918+7.1.0安装视频

视频链接：https://www.bilibili.com/video/BV18B4y1U7DK/?spm_id_from=333.337.search-card.all.click&vd_source=0d1b5c9c647732bc2bf2089f7476dfb0

### 1. 安装步骤

1. 查找局域网中群辉设备：https://finds.synology.com/

2. 查找到设备后直接在浏览器中输入 IP:5000 进入网页按照视频步骤安装即可。

### 2. 注意

安装前进行路由器断网(拔掉WAN口网线)，不然在联网的状态下安装会有错误（我们检测到您之前将硬盘移动到新的DS3617xs。如果您要现在还原数据和设置，请单击“还原”，参考链接：https://blog.csdn.net/qq_42966566/article/details/126229547）。

### 3. 出现的错误

#### 1. 群晖docker查询注册表失败

参考链接：https://post.smzdm.com/p/aqme9k5k/

我在网上看常用的解决办法就是填写国内加速镜像地址和修改DNS，具体步骤如下：

1、群晖docker——注册表——设置——选中Docker Hub——编辑——启用注册表镜像——在里面填写国内加速镜像地址：https://registry.docker-cn.com，然后重启docker

2、控制面板——网络——手动配置DNS[服务器](https://www.smzdm.com/fenlei/fuwuqi/)，在里面填写国内公用DNS，一般用阿里的或者114.114.114.114

如果这两步做完docker注册表就可以正常使用了的话，接下来就不用再看啦，如果这两步做完仍然解决不了问题，那接下来看我琢磨出来的方案。

方案二：（我用的这个方案）

1. 首先在控制面板——终端机和SNMP里面启用SSH功能
2. 使用终端工具ssh连上，用命令行docker pull



