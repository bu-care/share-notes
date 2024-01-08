FROM ubuntu

# Install logrotate
RUN apt-get update && \
    apt-get -y install iputils-ping iproute2 curl && \
    apt-get -y install logrotate python3 python3-pip

#RUN pip3 config set global.index-url http://mirrors.aliyun.com/pypi/simple
#RUN pip3 config set install.trusted-host mirrors.aliyun.com


COPY ./requirements.txt /
#RUN pip3 install -r requirements.txt
## vcenter中的虚拟机使用阿里云源pip下载不了，使用这个清华源却可以
RUN pip3 install -r requirements.txt -i https://pypi.tuna.tsinghua.edu.cn/simple/

COPY ./runserver.sh /
COPY ./app  /app

# RUN mkdir /var/log/gunicorn
# #Copy logrotate gunicorn configuration
# COPY ./gunicorn.logrotate /etc/logrotate.d/gunicorn
# # change config file mode
# RUN chmod 644 /etc/logrotate.d/gunicorn
# # Delete su root syslog line since ubuntu doesn't support logrotate syslog
# RUN sed -i '/^su *root *syslog$/d' /etc/logrotate.conf

# # Start gunicorn nginx and cron as services
# CMD service cron start && /runserver.sh

# RUN sed -i '$a\if [ -f /runserver.sh ]; then ' /root/.bashrc \
# && sed -i '$a\     . /runserver.sh' /root/.bashrc \
# && sed -i '$a\fi' /root/.bashrc

CMD ["/bin/bash", "runserver.sh"]