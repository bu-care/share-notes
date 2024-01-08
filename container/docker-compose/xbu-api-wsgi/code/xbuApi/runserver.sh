#!/bin/bash


# cd app && \
# gunicorn -p /run/gunicorn.pid --log-level debug --log-file=/var/log/gunicorn/gunicorn.log --workers 2 --name app -b 0.0.0.0:80 --reload mian:app
# # gunicorn -p /run/gunicorn.pid --log-level debug --log-file=/dev/stdout --workers 4 --name app -b 0.0.0.0:80 --reload app:app --preload

LOGTIME=$(date "+%Y-%m-%d %H:%M:%S")
echo "[$LOGTIME] startup run..." >>/app/runserver.log

cd app && python3 main.py