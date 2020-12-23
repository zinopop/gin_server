#!/bin/sh
#sudo curl http://54.215.230.116:9001/sendMsg?text=开始构建服务...
sudo docker stop dingtalk
sudo docker rm dingtalk
sudo docker rmi dingtalk:latest

cd /home/admin/gin_server_v1
#sudo curl http://54.215.230.116:9001/sendMsg?text=最新代码拉取...
sudo git pull
sudo docker build -t dingtalk -f Dockerfile_ding

sudo docker run -d -p 9001:9001 --name dingtalk dingtalk