## 1、virtualenv 常用命令

### 1、virtualenv 创建虚拟环境

需要选择一个python解释器来创建虚拟化环境，命令则为：

```bat
virtualenv -p /usr/bin/python3.8 devtest
virtualenv /home/python_env/devtest --python=python3.9
```

激活虚拟化环境的命令为：

```bat
source devtest/bin/activate
# 退出当前的venv虚拟化环境
deactivate
```

其他命令

```bat
用法:

virtualenv [OPTIONS] DEST_DIR 
选项: 
–version 			  # 显示当前版本号。 
-h, –help 			  #显示帮助信息。 
-v, –verbose 		  #显示详细信息。 
-q, –quiet 			  #不显示详细信息。 
#指定所用的python解析器的版本，比如 –python=python2.5 就使用2.5版本的解析器创建新的隔离环境。 默认使用的是当前系统安装(/usr/bin/python)的python解析器 
-p PYTHON_EXE, –python=PYTHON_EXE 

–clear            	    #清空非root用户的安装，并重头开始创建隔离环境。 
–no-site-packages       #令隔离环境不能访问系统全局的site-packages目录。 
–system-site-packages   #令隔离环境可以访问系统全局的site-packages目录。 
–unzip-setuptools       #安装时解压Setuptools或Distribute 
–relocatable            #重定位某个已存在的隔离环境。使用该选项将修正脚本并令所有.pth文件使用相当路径。 
–distribute             #使用Distribute代替Setuptools，也可设置环境变量VIRTUALENV_DISTRIBUTE达到同样效要。 
–extra-search-dir=SEARCH_DIRS #用于查找setuptools/distribute/pip发布包的目录。可以添加任意数量的–extra-search-dir路径。 
–never-download         #禁止从网上下载任何数据。此时，如果在本地搜索发布包失败，virtualenv就会报错。 
–prompt==PROMPT         #定义隔离环境的命令行前缀。
```

安装virtualenv 遇到问题

```bat
# WARNING: The script virtualenv is installed in '/home/xbu/.local/bin' which is not on PATH.
# 原因：'/home/xbu/.local/bin' 未添加到环境变量
echo 'export PATH=/home/xbu/.local/bin:$PATH' >>~/.bashrc
source ~/.bashrc
```

