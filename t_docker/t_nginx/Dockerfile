# 使用官方Nginx镜像为基础镜像
FROM nginx:latest

# 安装构建工具和Brotli模块的依赖
RUN apt-get update && apt-get install -y \
    build-essential \
    zlib1g-dev \
    libpcre3-dev \
    libssl-dev \
    git \
    wget \
    && rm -rf /var/lib/apt/lists/*

# 下载和安装Brotli模块
RUN git clone --recursive https://github.com/google/ngx_brotli.git /tmp/ngx_brotli \
    && wget http://nginx.org/download/nginx-1.23.0.tar.gz -P /tmp \
    && cd /tmp && tar -zxvf nginx-1.23.0.tar.gz \
    && cd nginx-1.23.0 \
    && ./configure --add-dynamic-module=/tmp/ngx_brotli \
    && make && make install

# 清理
RUN rm -rf /tmp/*
