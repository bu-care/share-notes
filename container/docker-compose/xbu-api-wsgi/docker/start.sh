#!/bin/bash


#LOG_DIR="/opt/xbuApiLog"
#mkdir -p $LOG_DIR/xbuApi

APP_DIR="/srv/xbuApi/app"

sudo rm -rf $APP_DIR
sudo mkdir -p $APP_DIR

#docker-compose up -d --remove-orphans xbu_api
docker-compose up -d xbu_api

#./run-agent.sh -s 2cb8ed694e20 -a http://10.103.12.69:8004 -n agent1
