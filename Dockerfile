# 前端构建阶段
FROM node:18-alpine AS frontend-builder

WORKDIR /app/frontend

# 复制前端项目文件
COPY frontEnd/my-next-app/package*.json ./
RUN npm install

COPY frontEnd/my-next-app/ .
RUN npm run build

# 后端构建阶段
FROM golang:1.21-alpine AS backend-builder

WORKDIR /app/backend

# 复制后端项目文件
COPY API/go.mod API/go.sum ./
RUN go mod download

COPY API/ .
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# 最终运行阶段
FROM alpine:latest

# 安装 nodejs 和 npm
RUN apk add --no-cache nodejs npm ca-certificates

# 创建工作目录
WORKDIR /app

# 复制前端文件
COPY --from=frontend-builder /app/frontend/.next ./.next
COPY --from=frontend-builder /app/frontend/node_modules ./node_modules
COPY --from=frontend-builder /app/frontend/package.json ./package.json
COPY --from=frontend-builder /app/frontend/public ./public

# 复制后端文件
COPY --from=backend-builder /app/backend/main ./
COPY --from=backend-builder /app/backend/config ./config

# 暴露端口
EXPOSE 3000 8080

# 启动脚本
COPY start.sh .
RUN chmod +x start.sh

# 启动服务
CMD ["./start.sh"]