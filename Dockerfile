# 导入其他的镜像
FROM golang:alpine

# 为我们的镜像设置必要的环境变量
# GO111MODULE=on \ 开启GO111MODULE
# CGO_ENABLED=0 \ 禁用CGO_ENABLED
# GOOS=linux \ 编译的操作系统
# GOARCH=amd64 \ 64位
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.io

# 移动到工作目录：/build
WORKDIR /build

# 将代码复制到容器中
COPY . .

# 将我们的代码编译成二进制可执行文件app
# -o  ： 指定编译后的二进制文件名称
# .  :编译当下目录下所有文件
RUN go build -o app .


# 声明服务端口
EXPOSE 8888

# 启动容器时运行的命令
CMD ["/build/app"]

#
#FROM golang:alpine AS builder
#
## 为我们的镜像设置必要的环境变量
#ENV GO111MODULE=on \
#    CGO_ENABLED=0 \
#    GOOS=linux \
#    GOARCH=amd64
#
## 移动到工作目录：/build
#WORKDIR /build
#
## 将代码复制到容器中
#COPY . .
#
## 将我们的代码编译成二进制可执行文件 app
#RUN go build -o app .
#
####################
## 接下来创建一个小镜像
####################
#FROM scratch
#
## 从builder镜像中把/dist/app 拷贝到当前目录
#COPY --from=builder /build/app /
#
## 需要运行的命令
#ENTRYPOINT ["/app"]