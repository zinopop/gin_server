#!/bin/sh
sudo curl http://54.215.230.116:9001/sendMsg?text=开始构建服务...
sudo docker stop gin_server
sudo docker rm gin_server
sudo docker rmi gin_server:latest

cd /home/admin/gin_server_v1
#sudo curl http://54.215.230.116:9001/sendMsg?text=最新代码拉取...
sudo git pull
sudo docker build -t gin_server .

sudo docker run -d -p 9002:9002 --name gin_server gin_server