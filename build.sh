#!/bin/sh
. /etc/profile
sudo curl http://54.215.230.116:9001/sendMsg?text=开始构建服务...
sudo docker stop gin_server
sudo docker rm gin_server
sudo docker rmi gin_server:latest
cd /home/admin/gin_server_v1

sudo curl http://54.215.230.116:9001/sendMsg?text=代码热更新中...
sudo git fetch --all && sudo git reset --hard origin/master && sudo git pull && sudo chmod +x ./build.sh ./build_ding.sh
sudo git pull

sudo curl http://54.215.230.116:9001/sendMsg?text=开始构建并运行中...
sudo docker build -t gin_server .
sudo docker run -d -p 9002:9002 --name gin_server gin_server
sudo docker rmi $(sudo docker images -f "dangling=true" -q)
sudo curl http://54.215.230.116:9001/sendMsg?text=构建成功并运行,服务端口开启:9002...