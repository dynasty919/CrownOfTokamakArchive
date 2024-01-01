# 使用 Go 镜像
FROM golang:1.19.3 AS go-builder

WORKDIR /app

ENV GOPROXY=https://goproxy.cn

# 复制 go.mod 和 go.sum 文件并下载依赖
COPY go.mod go.sum ./
RUN go mod download

# 复制整个项目并构建
COPY . .

# 构建 Go 项目
RUN go build -o main ./server

# 最终镜像
FROM golang:1.19.3

WORKDIR /app

# 复制二进制文件
COPY --from=go-builder /app/main /app/main

# 复制其他需要的文件
COPY server/template /app/server/template
COPY server/pb /app/server/template
COPY persist /app/persist
COPY util /app/util

# 启动 Go 项目
CMD ["./main"]