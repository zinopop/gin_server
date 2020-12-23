#!/bin/sh
sudo git pull https://github.com/zinopop/gin_server_v1.git
sudo docker build -t server:v2 .
sudo docker stop dingTalk
sudo docker run -d -p 9001:9001 --name dingTalk server:v2