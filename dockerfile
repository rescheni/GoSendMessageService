FROM golang:1.21-alpine AS builder

# 设置工作目录
WORKDIR /app

# 复制 go.mod 和 go.sum
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制源代码
COPY . .

# 构建应用
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 运行阶段
FROM alpine:latest

# 安装 CA 证书
RUN apk --no-cache add ca-certificates

WORKDIR /app

# 从构建阶段复制二进制文件
COPY --from=builder /app/main .
COPY --from=builder /app/config.json .

# 暴露端口（如果需要的话）
EXPOSE 8080

# 运行应用
CMD ["./main"]
