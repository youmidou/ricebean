# 使用多阶段构建
FROM golang:1.24.3 AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o gameserver ./examples/GameServer

# 使用轻量级的 alpine 作为基础镜像
FROM alpine:latest

# 安装必要的系统依赖
RUN apk --no-cache add ca-certificates tzdata

# 设置时区
ENV TZ=Asia/Shanghai

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译好的应用
COPY --from=builder /app/gameserver .

# 复制配置文件
COPY --from=builder /app/config ./config

# 暴露端口
EXPOSE 1250

# 运行应用
CMD ["./gameserver"] 