FROM golang:1.15 as builder
# 配置代理
ENV GOPROXY https://goproxy.cn
# 设置go缓存
ENV GO111MODULE=on

# 设置编码格式
ENV LANG en_US.UTF-8
ENV LANGUAGE en_US:en
ENV LC_ALL en_US.UTF-8

# go缓存
WORKDIR /go/cache
ADD go.mod .
ADD go.sum .
RUN go mod download
# 项目工作路径
WORKDIR /go/release

ADD . .

# 运行命令
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go

# 调试模式开始 ---------------------------------------------------------------------------
#FROM alpine:3.12.3
## 安装 /bin/bash 方便调试
#RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.4/main/" > /etc/apk/repositories
#RUN apk update \
#        && apk upgrade \
#        && apk add --no-cache bash \
#        bash-doc \
#        bash-completion \
#        && rm -rf /var/cache/apk/* \
#        && /bin/bash
## 安装证书 解决http请求
#RUN apk add --update ca-certificates
#RUN update-ca-certificates
#
##安装curl
#
## 统一亚洲时区
#ENV TZ=Asia/Shanghai
#RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezon
#
## 设置目录
#WORKDIR /root/
#RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
##把app文件从“builder”中拷贝到本级的当前目录
#COPY --from=builder /go/release .
##把app文件从“builder”中拷贝到本级的当前目录
# 调试模式结束 ---------------------------------------------------------------------------


# 暴露端口
EXPOSE 9002
# 项目默认启动命令
ENTRYPOINT ./\app