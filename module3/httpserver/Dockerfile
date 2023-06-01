# 基于官方 Ubuntu 镜像构建
FROM ubuntu

# 安装必要的软件包和依赖项
RUN apt-get update && apt-get install -y curl git build-essential

# 下载并安装 Go 1.18.3
RUN curl -O https://dl.google.com/go/go1.18.3.linux-arm64.tar.gz \
    && tar -C /usr/local -xzf go1.18.3.linux-arm64.tar.gz \
    && rm go1.18.3.linux-arm64.tar.gz

# 设置 Go 环境变量
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/go

# 设置工作目录
WORKDIR /app

# 将宿主机当前目录下的文件都复制到容器的 /app 目录下
COPY . /app

# 编译 Go 应用
RUN go build -o httpserver .

# 暴露容器的端口
EXPOSE 80

# 启动应用
CMD ["./httpserver"]