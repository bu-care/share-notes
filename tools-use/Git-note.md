[toc]

---

狂神说Git：https://mp.weixin.qq.com/s/Bf7uVhGiu47uOELjmC5uXQ

博客园深入浅出Git教程：https://www.cnblogs.com/syp172654682/p/7689328.html

## 1、Git 相关配置

注册 ssh-key

```bat
ssh-keygen -t rsa -b 2048 -C "xbu@sonicwall.com"
```

设置签名

```bat
#项目级别/仓库级别： 仅在当前本地库范围内有效
git config user.name xbu
git config user.email xbu@sonicwall.com
#信息保存位置： ./.git/config 文件

#系统用户级别： 登录当前操作系统的用户范围
git config --global user.name xbu
git config --global user.email xbu@sonicwall.com
# 信息保存位置： ~/.gitconfig 文件
```

测试连接

```bat
ssh -vT git@gitlab.com

eval $(ssh-agent -s)
ssh-add ~/.ssh/id_rsa
```

## 2、Git 基本原理

### 1、工作区域

Git本地有三个工作区域：工作目录（Working Directory）、暂存区(Stage/Index)、资源库(Repository或Git Directory)。如果在加上远程的git仓库(Remote Directory)就可以分为四个工作区域。文件在这四个区域之间的转换关系如下：

![图片](https://mmbiz.qpic.cn/mmbiz_png/uJDAUKrGC7Ksu8UlITwMlbX3kMGtZ9p0NJ4L9OPI9ia1MmibpvDd6cSddBdvrlbdEtyEOrh4CKnWVibyfCHa3lzXw/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

- Workspace：工作区，就是你平时存放项目代码的地方
- Index / Stage：暂存区，用于临时存放你的改动，事实上它<u>**只是一个文件**</u>，保存即将提交到文件列表信息
- Repository：仓库区（或本地仓库），就是安全存放数据的位置，这里面有你提交到所有版本的数据。其中HEAD指向最新放入仓库的版本
- Remote：远程仓库，托管代码的服务器，可以简单的认为是你项目组中的一台电脑用于远程数据交换

本地的三个区域确切的说应该是git仓库中HEAD指向的版本：

![图片](https://mmbiz.qpic.cn/mmbiz_png/uJDAUKrGC7Ksu8UlITwMlbX3kMGtZ9p0icz6X2aibIgUWzHxtwX8kicPCKpDrsiaPzZk04OlI2bzlydzicBuXTJvLEQ/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

- Directory：使用Git管理的一个目录，也就是一个仓库，包含我们的工作空间和Git的管理空间。
- WorkSpace：需要通过Git进行版本控制的目录和文件，这些目录和文件组成了工作空间。
- .git：存放Git管理信息的目录，初始化仓库的时候自动创建。
- Index/Stage：暂存区，或者叫待提交更新区，在提交进入repo之前，我们可以把所有的更新放在暂存区。
- Local Repo：本地仓库，一个存放在本地的版本库；HEAD会只是当前的开发分支（branch）。
- Stash：隐藏，是一个工作状态保存栈，用于保存/恢复WorkSpace中的临时状态。

### 2、工作流程

git的工作流程一般是这样的：

１、在工作目录中添加、修改文件；

２、将需要进行版本管理的文件放入暂存区域；

３、将暂存区域的文件提交到git仓库。

因此，git管理的文件有三种状态：已修改（modified）,已暂存（staged）,已提交(committed)

![图片](https://mmbiz.qpic.cn/mmbiz_png/uJDAUKrGC7Ksu8UlITwMlbX3kMGtZ9p09iaOhl0dACfLrMwNbDzucGQ30s3HnsiaczfcR6dC9OehicuwibKuHjRlzg/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

## 3、Git 项目搭建

### 1、创建工作目录与常用命令

工作目录（WorkSpace)一般就是你希望Git帮助你管理的文件夹，可以是你项目的目录，也可以是一个空目录，建议不要有中文。

日常使用只要记住下图6个命令：

![图片](https://mmbiz.qpic.cn/mmbiz_png/uJDAUKrGC7Ksu8UlITwMlbX3kMGtZ9p0AII6YVooUzibpibzJnoOHHXUsL3f9DqA4horUibfcpEZ88Oyf2gQQNR6w/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

>本地仓库搭建

创建本地仓库的方法有两种：一种是创建全新的仓库，另一种是克隆远程仓库。

1、创建全新的仓库，需要用GIT管理的项目的根目录执行：

```bat
# 在当前目录新建一个Git代码库
$ git init
```

2、执行后可以看到，仅仅在项目目录多出了一个.git目录，关于版本等的所有信息都在这个目录里面。

>克隆远程仓库

1、另一种方式是克隆远程目录，由于是将远程服务器上的仓库完全镜像一份至本地！

```bat
# 克隆一个项目和它的整个代码历史(版本信息)
$ git clone [url]  # https://gitee.com/kuangstudy/openclass.git
```

2、去 gitee 或者 github 上克隆一个测试！

### 2、Git 文件操作

>文件的四种状态

版本控制就是对文件的版本控制，要对文件进行修改、提交等操作，首先要知道文件当前在什么状态，不然可能会提交了现在还不想提交的文件，或者要提交的文件没提交上。

- **Untracked**: 未跟踪, 此文件在文件夹中, 但并没有加入到git库, 不参与版本控制. 通过`git add` 状态变为`Staged`.
- **Unmodify**: 文件已经入库, 未修改, 即版本库中的文件快照内容与文件夹中完全一致. 这种类型的文件有两种去处, 如果它被修改, 而变为`Modified`. 如果使用`git rm`移出版本库, 则成为`Untracked`文件
- **Modified**: 文件已修改, 仅仅是修改, 并没有进行其他的操作. 这个文件也有两个去处, 通过`git add`可进入暂存`staged`状态, 使用`git checkout` 则丢弃修改过, 返回到`unmodify`状态, 这个`git checkout`即从库中取出文件, 覆盖当前修改
- **Staged**: 暂存状态. 执行`git commit`则将修改同步到库中, 这时库中的文件和本地文件又变为一致, 文件为`Unmodify`状态. 执行`git reset HEAD filename`取消暂存, 文件状态为`Modified`

![img](https://images2017.cnblogs.com/blog/63651/201709/63651-20170909091456335-1787774607.jpg)



>查看文件状态

上面说文件有4种状态，通过如下命令可以查看到文件的状态：

```bat
#查看指定文件状态
git status [filename]
#查看所有文件状态
git status
# git add .                  添加所有文件到暂存区
# git commit -m "消息内容"    提交暂存区中的内容到本地仓库 -m 提交信息
```

>忽略文件

有些时候我们不想把某些文件纳入版本控制中，比如数据库文件，临时文件，设计文件等

在主目录下建立".gitignore"文件，此文件有如下规则：

1. 忽略文件中的空行或以井号（#）开始的行将会被忽略。
2. 可以使用Linux通配符。例如：星号（*）代表任意多个字符，问号（？）代表一个字符，方括号（[abc]）代表可选字符范围，大括号（{string1,string2,...}）代表可选的字符串等。
3. 如果名称的最前面有一个感叹号（!），表示例外规则，将不被忽略。
4. 如果名称的最前面是一个路径分隔符（/），表示要忽略的文件在此目录下，而子目录中的文件不忽略。
5. 如果名称的最后面是一个路径分隔符（/），表示要忽略的是此目录下该名称的子目录，而非文件（默认文件或目录都忽略）。

```shell
#为注释
*.txt        #忽略所有 .txt结尾的文件,这样的话上传就不会被选中！
!lib.txt     #但lib.txt除外
/temp        #仅忽略项目根目录下的TODO文件,不包括其它目录temp
build/       #忽略build/目录下的所有文件
doc/*.txt    #会忽略 doc/notes.txt 但不包括 doc/server/arch.txt
```

### 3、IDE 中集成Git

1、新建项目，绑定git。

![图片](https://mmbiz.qpic.cn/mmbiz_png/uJDAUKrGC7Ksu8UlITwMlbX3kMGtZ9p0D8LPGu2SNKXD01IMqDaSkBeP8ibtvnasBYiaReyuZWAl0EjEib8IYf7cQ/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

注意观察idea中的变化

![图片](https://mmbiz.qpic.cn/mmbiz_png/uJDAUKrGC7Ksu8UlITwMlbX3kMGtZ9p0Cs93BiaOia1Sdk8icdH7vQzPfzIjuoTNYquKzYtrEe5mklhg2b7KOYsow/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

2、修改文件，使用IDEA操作git。

- 添加到暂存区
- commit 提交
- push到远程仓库

3、提交测试

## 4、Git 分支

分支在GIT中相对较难，分支就是科幻电影里面的平行宇宙，如果两个平行宇宙互不干扰，那对现在的你也没啥影响。不过，在某个时间点，两个平行宇宙合并了，我们就需要处理一些问题了！

![图片](https://mmbiz.qpic.cn/mmbiz_png/uJDAUKrGC7Ksu8UlITwMlbX3kMGtZ9p0BOGzaG4QTc4JXO0hSlwcNtujNzAvxeibSrajLYLCT6otNnHDV9xYWwA/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

![图片](https://mmbiz.qpic.cn/mmbiz_png/uJDAUKrGC7Ksu8UlITwMlbX3kMGtZ9p0Ayn87woxfepOhSlUj4FQTFUsia4ic0j6aQy4Tz32PRuJ0HSVeGeUzURA/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)

git分支中常用指令：

```bat
# 列出所有本地分支
git branch

# 列出所有远程分支
git branch -r

# 新建一个分支，但依然停留在当前分支
git branch [branch-name]

# 新建一个分支，并切换到该分支
git checkout -b [branch]

# 合并指定分支到当前分支
$ git merge [branch]

# 删除分支
$ git branch -d [branch-name]

# 删除远程分支
$ git push origin --delete [branch-name]
$ git branch -dr [remote/branch]
```

IDEA中操作

![图片](https://mmbiz.qpic.cn/mmbiz_png/uJDAUKrGC7Ksu8UlITwMlbX3kMGtZ9p0wHNIYeTHC8aHGASoDyZO64QicslqiaMb1OJ1Z1LPoic3LBGyDIYBa7XXw/640?wx_fmt=png&wxfrom=5&wx_lazy=1&wx_co=1)







## Git 相关操作

### 1. 远程拉取代码

```bat
#初始化仓库
git init

# 可以直接克隆
git clone *****.com:***git

# 添加远程仓库
git remote add origin *****.com:***git

# 查看本地分支
git branch -a 或者 -v

git fetch origin master
git checkout <branch>

#在本地创建分支dev并切换到该分支
git checkout -b dev origin/dev
```

### 2. git更新远程仓库代码到本地

git fetch：这将更新git remote 中所有的远程仓库所包含分支的最新commit-id, 将其记录到.git/FETCH_HEAD文件中 

git fetch更新远程仓库的方式如下：
```bat
# 在本地新建一个temp分支，并将远程origin仓库的master分支代码下载到本地temp分支
git fetch origin master:temp 
 
# 来比较本地代码与刚刚从远程下载下来的代码的区别
git diff temp 
 
# 合并temp分支到本地的master分支
git merge temp
 
# 如果不想保留temp分支 可以用这步删除
git branch -d temp
```

也可以用以下指令：
```bat
# 将远程仓库的master分支下载到本地当前branch中
git fetch orgin master
 
# 比较本地的master分支和origin/master分支的差别
git log -p master  ..origin/master
 
# 进行合并
git merge origin/master
```

git pull : 首先，基于本地的FETCH_HEAD记录，比对本地的FETCH_HEAD记录与远程仓库的版本号，然后git fetch 获得当前指向的远程分支的后续版本的数据，然后再利用git merge将其与本地的当前分支合并。所以可以认为git pull是git fetch和git merge两个步骤的结合。 
git pull的用法如下：

```bat
# 取回远程主机某个分支的更新，再与本地的指定分支合并。
git pull <远程库名> <远程分支名>:<本地分支名>
 
# 取回远程库中的master分支，与本地的master分支进行合并更新，要写成：
git pull origin master:master
 
# 如果是要与本地当前分支合并更新，则冒号后面的<本地分支名>可以不写
git pull origin master
```

push代码时发现远端和本地不一致，用`git fetch && git rebase` 替换`git pull`命令，这样就不会产生这种没用的 `git commit` 了。



### 3、git stash

#### 1、git stash 的作用

链接：https://blog.csdn.net/andyzhaojianhui/article/details/80586695

`git stash`用于想要保存当前的修改,但是想回到之前最后一次提交的干净的工作仓库时进行的操作.`git stash`将本地的修改保存起来,并且将当前代码切换到`HEAD`提交上.

通过`git stash`存储的修改列表,可以通过`git stash list`查看.`git stash show`用于校验,`git stash apply`用于重新存储.直接执行`git stash`等同于`git stash save`.

最新的存储保存在`refs/stash`中.老的存储可以通过相关的参数获得,例如`stash@{0}`获取最新的存储,`stash@{1}`获取次新.`stash@{2.hour.ago}`获取两小时之前的.存储可以直接通过索引的位置来获得`stash@{n}`.

#### 2、命令详解

链接：https://blog.csdn.net/stone_yw/article/details/80795669

##### 1、git stash

能够将所有未提交的修改（工作区和暂存区）保存至堆栈中，用于后续恢复当前工作目录。

```shell
$ git status
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

        modified:   src/main/java/com/wy/CacheTest.java
        modified:   src/main/java/com/wy/StringTest.java

no changes added to commit (use "git add" and/or "git commit -a")

$ git stash
Saved working directory and index state WIP on master: b2f489c second

$ git status
On branch master
nothing to commit, working tree clean
```

##### 2、git stash save

作用等同于git stash，区别是可以加一些注释，如下：

```shell
# git stash的效果：
stash@{0}: WIP on master: b2f489c second

# git stash save “test1”的效果：
stash@{0}: On master: test1
```

##### 3. git stash list

查看当前stash中的内容

##### 4. git stash pop

将当前stash中的内容弹出，并应用到当前分支对应的工作目录上。注：该命令将堆栈中最近保存的内容删除（栈是先进后出）
顺序执行git stash save “test1”和git stash save “test2”命令，效果如下：

```shell
$ git stash list
stash@{0}: On master: test2
stash@{1}: On master: test1

$ git stash pop
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

        modified:   src/main/java/com/wy/StringTest.java

no changes added to commit (use "git add" and/or "git commit -a")
Dropped refs/stash@{0} (afc530377eacd4e80552d7ab1dad7234edf0145d)

$ git stash list
stash@{0}: On master: test1
```

可见，test2的stash是首先pop出来的。
如果从stash中恢复的内容和当前目录中的内容发生了冲突，也就是说，恢复的内容和当前目录修改了同一行的数据，那么会提示报错，需要解决冲突，可以通过**创建新的分支来解决冲突**。

##### 5. git stash apply

将堆栈中的内容应用到当前目录，不同于git stash pop，该命令**不会将内容从堆栈中删除**，也就说该命令能够将堆栈的内容多次应用到工作目录中，适应于多个分支的情况。

```shell
$ git stash apply
On branch master
Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git checkout -- <file>..." to discard changes in working directory)

        modified:   src/main/java/com/wy/StringTest.java

no changes added to commit (use "git add" and/or "git commit -a")

$ git stash list
stash@{0}: On master: test2
stash@{1}: On master: test1
```

**堆栈中的内容并没有删除**。可以使用git stash apply + stash名字（如stash@{1}）指定恢复哪个stash到当前的工作目录。

##### 6 . git stash drop + 名称

从堆栈中移除某个指定的stash

##### 7. git stash clear

清除堆栈中的所有 内容

##### 8. git stash show

查看堆栈中最新保存的stash和当前目录的差异。

```shell
$ git stash show
 src/main/java/com/wy/StringTest.java | 2 +-
 1 file changed, 1 insertion(+), 1 deletion(-)
```

git stash show stash@{1}查看指定的stash和当前目录差异。
通过 git stash show -p 查看详细的不同：

##### 9. git stash branch

从最新的stash创建分支。
应用场景：当储藏了部分工作，暂时不去理会，继续在当前分支进行开发，后续想将stash中的内容恢复到当前工作目录时，如果是针对同一个文件的修改（即便不是同行数据），那么可能会发生冲突，恢复失败，这里通过创建新的分支来解决。可以用于解决stash中的内容和当前目录的内容发生冲突的情景。


## 遇到的错误
### 1.  git@github.com Permission denied (publickey)
ssh -T git@github.com   测试公钥是否添加成功(只有域名，没有后面的仓库)

解决办法： https://blog.csdn.net/qq_32786873/article/details/80947195

一： 使用ssh-agent代理管理git私钥
```shell
# 1）启动agent： 
eval $(ssh-agent) 
# 或者  (注意这里是反引号)
eval `ssh-agent -s`
#（2）添加本地的私钥： ssh-add 本地私钥路径
ssh-add /root/.ssh/id_rsa
```











