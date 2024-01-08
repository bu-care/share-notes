## 1. tmux 的常用命令



```bat
# 在ubuntu下使用命令安装tmux
sudo apt-get install tmux
# centos安装tmux
sudo yum install tux
 
# 显示所有会话 
tmux ls

# 运行tmux并开启一个新的会话 
tmux 
# 新建会话（不指定会话名称）
tmux new

# 新建会话并指定会话名称（建议制定会话名称，以便了解该会话用途） 
tmux new -s <session-name>
 
# 接入上一个会话
tmux a #
接入指定名称的会话
tmux a -t <session-name>
 
# 断开当前会话（还可以使用快捷键：control+b，再按d）
tmux detach
 
# 关闭指定会话 
tmux kill-session -t session-name
 
# 关闭除指定会话外的所有会话
tmux kill-session -a -t session-name
 
# 在会话中切换 
control+b，再按s 显示会话列表，再进行会话切换
 
# 销毁所有会话并停止
tmux tmux kill-server
 
# 重命名会话
tmux rename-session -t abc cba
```



## 2. screen 的常用命令



```bat
# 在ubuntu下使用命令安装tmux
sudo apt-get install screen
 
# 显示所有会话 
screen -ls

# 新建会话并指定会话名称（建议制定会话名称，以便了解该会话用途） 
screen -S <session-name>

# 接入指定名称的会话
screen -r <session-name>
 
# 断开当前会话dettach（还可以使用快捷键：control+a，再按d）
screen -d
 
# kill掉一个screen
#（1）、使用screen名字，kill掉。
 screen -S session_name -X quit
#（2）、激活screen,然后在会话中输出exit直接退出，也会关闭会话。
screen -r session_name
exit # 提示：[screen is terminating]，表示已经成功退出screen会话。
 
# 设置缓存：输入 ctr + a、: 再输入 scrollback 1234，代表设置窗口缓存为1234行。
#查看历史信息：Ctrl + a、Esc，进入 “copy mode” ，然后就可以查看历史信息，甚至可以使用vim命令。按 esc 退出。
```

使用screen难免遇到这种情况，当screen输出太长时屏幕滚动,不能看到全部信息。

```bat
# 第一种方法：
# 启动时添加选项-L（Turn on output logging.），会在当前目录下生成screenlog.0文件。
screen -L -dmS zta #启动一个开始就处于断开模式的会话，会话的名称是 zta
screen -r zta #连接该会话，在会话中的所有屏幕输出都会记录到screenlog.0文件。
# 注：如果执行 -L 命令后看不到新建的日志，可能是文件读写权限不够，e.g. sudo chmod 777 filename 可以修改文件读写权限。

# 第二种方法：
# 不加选项-L，启动后，在 screen session 下按 ctrl+a H，同样会在当前目录下生成 screenlog.0 文件。
# 第一次按下ctrl+a H，屏幕左下角会提示Creating logfile “screenlog.0”.，开始记录日志。
# 再次按下ctrl+a H，屏幕左下角会提示Logfile “screenlog.0” closed.，停止记录日志。
# 注：上面两个方法有个缺点：当创建多个screen会话的时候，每个会话都会记录日志到screenlog.0文件。screenlog.0中的内容就比较混乱了。

# 第三种方法（推荐）：
# 在 /tmp 目录下创建文件夹，添加权限
mkdir /tmp/screen_log -p
chmod 777 screen_log -R
# 在 screen 配置文件 /etc/screenrc 最后添加下面一行：
logfile /tmp/screenlog_%t.log
# 注意，如果写成：logfile ./screenlog_%t.log 则是把日志文件记录到当前目录下
# %t 是指window窗口的名称，对应 screen 的 -t 参数。所以我们启动 screen 的时候要指定窗口的名称，例如：
screen -L -t zta-window -S zta
# 意思是启动 zta 会话，窗口名称为 zta-window 。屏幕日志记录在/tmp/screenlog_zta-window.log。
# 如果启动的时候不加 -L 参数，在screen session下按 ctrl+a H，日志也会记录在/tmp/screenlog_zta-window.log。
```

