#!/bin/sh
sudo docker stop dingTalk && docker rm dingTalk
sudo docker build -t server .

cd /home/admin/gin_server_v1
sudo git pull

sudo docker run -d -p 9001:9001 --name dingTalk server