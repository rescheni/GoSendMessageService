# ========== 前端构建阶段 ==========
FROM node:18-alpine AS frontend-builder

WORKDIR /app/frontend

# 复制前端项目文件
COPY frontEnd/my-next-app/package*.json ./
RUN npm install

COPY frontEnd/my-next-app/ .

# 构建前端项目
# RUN npm ci --omit=dev && npm run build
RUN npm run build

# ========== 后端构建阶段 ==========
FROM golang:1.23.2-alpine AS backend-builder
# 安装安装编译工具链
RUN apk add --no-cache build-base sqlite-dev
WORKDIR /app/backend

# 复制整个项目
# COPY . .
# 依赖清单
COPY go.mod go.sum main.go ./
COPY API/ ./API/
# 项目库
COPY database/ ./database/
COPY sendserver/ ./sendserver/
COPY services/ ./services/
COPY Basic/ ./Basic/
COPY log/ ./log/
# 插件库
COPY plug-in/ ./plug-in/
# 配置文件
COPY config/ ./config/
# 下载依赖
RUN go mod download -x
# 后端构建阶段
RUN CGO_ENABLED=1 GOOS=linux go build -trimpath -ldflags="-s -w" -o server .
# 确保编译后文件可执行
RUN chmod +x server


# ========== 最终运行阶段 ==========
FROM alpine:latest

# 安装 nodejs 和 npm
# RUN apk add --no-cache nodejs npm ca-certificates
RUN apk add --no-cache \
    nodejs \
    npm \
    dos2unix \
    && addgroup -S appgroup \
    && adduser -S appuser -G appgroup

# 创建工作目录
WORKDIR /app

# 复制前端文件
COPY --from=frontend-builder /app/frontend/.next ./.next
COPY --from=frontend-builder /app/frontend/node_modules ./node_modules
COPY --from=frontend-builder /app/frontend/package.json ./package.json
COPY --from=frontend-builder /app/frontend/public ./public

# 复制后端文件
COPY --from=backend-builder /app/backend/server ./
COPY --from=backend-builder /app/backend/config ./config

# 暴露端口
EXPOSE 3000 8080

# 持久化数据
VOLUME [ "/app/data" ]
VOLUME [ "/app/config" ]

# 启动脚本
COPY start.sh .
RUN dos2unix start.sh && chmod +x start.sh

# 启动服务
CMD ["./start.sh"]