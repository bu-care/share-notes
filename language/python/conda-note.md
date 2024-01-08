通过vim /etc/profile在最后一行添加

```bat
export PATH="$PATH:/opt/anaconda/bin" 
source /etc/profile # 重启生效
```


共享环境

```bat
 groupadd anaconda  # 创建anaconda组
 chgrp -R anaconda /home/anaconda # 组内共享这个目录下的文件
 chmod 770 -R /home/anaconda # 设置权限
 adduser username anaconda # 添加用户进组
 source /etc/profile # 进组的用户需要自己刷新一下哦
```

用户将conda初始化添加进 .bashrc

```bat
/opt/anaconda3/bin/conda init bash.
source .bashrc # 用户需要自己刷新一下

# 关闭每次启动前面带有一个(bash)
conda config --set auto_activate_base false
# 重新开启
conda config --set auto_activate_base true
```



### 2、Anaconda创建新环境出错

在/opt目录下安装anaconda给所有用户使用，但有些用户没有权限，出现如下错误：

```bat
Collecting package metadata (current_repodata.json): failed

NotWritableError: The current user does not have write permissions to a required path.
  path: /opt/anaconda3/pkgs/cache/9e0f62c3.json
  uid: 1001
  gid: 1001

If you feel that permissions on this path are set incorrectly, you can manually
change them by executing

  $ sudo chown 1001:1001 /opt/anaconda3/pkgs/cache/9e0f62c3.json

In general, it's not advisable to use 'sudo conda'.
############# 对于单个用户
# 针对单个用户改变anaconda文件夹的权限，其中需要更改的地方，
# 1、1001:1001 第一个1001是uid, 第二个1001是gid, 这个两个数字根据报错的命令行提示更改
# 2、/opt/anaconda3/， 这是anaconda安装的路径，
sudo chown 1001:1001 /opt/anaconda3/

############## 对于多个用户
# 重看所有用户的uuid和gid
cat /etc/passwd 
root:x:0:0:root:/root:/bin/bash
uuidd:x:107:114::/run/uuidd:/usr/sbin/nologin
sonicwall:x:1000:1000:sonicwall:/home/sonicwall:/bin/bash
xbu:x:1001:1001::/home/xbu:/bin/bash
# 此时应该修改为：
sudo chown -R 1000:1001 /opt/anaconda3/

```

### 3. 将Conda Prompt Here添加到右键菜单

参考链接：https://blog.csdn.net/u012308586/article/details/129548115

1. 添加到空白处的右键菜单

```bash
REG ADD HKCR\Directory\Background\shell\Conda\ /ve /f /d "Conda Powershell Prompt Here"

REG ADD HKCR\Directory\Background\shell\Conda\ /v Icon /f /t REG_EXPAND_SZ /d C:\programs\anaconda3\Menu\Iconleak-Atrous-Console.ico

REG ADD HKCR\Directory\Background\shell\Conda\command /f /ve /t REG_EXPAND_SZ /d "%windir%\System32\cmd.exe "/K" C:\programs\anaconda3\Scripts\activate.bat
```

2. 添加到目录的右键菜单

   ```bash
   REG ADD HKCR\Directory\shell\Conda\ /ve /f /d "Conda Powershell Prompt Here"
   
   REG ADD HKCR\Directory\shell\Conda\ /v Icon /f /t REG_EXPAND_SZ /d C:\programs\anaconda3\Menu\Iconleak-Atrous-Console.ico
   
   REG ADD HKCR\Directory\shell\Conda\command /f /ve /t REG_EXPAND_SZ /d "%windir%\System32\cmd.exe "/K" C:\programs\anaconda3\Scripts\activate.bat
   ```

   

