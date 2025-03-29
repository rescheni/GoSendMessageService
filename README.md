# GoSendMessageService

- 信息发送服务
- 发送邮件、短信、微信、钉钉等

# 开发环境

- Golang 1.23.2
- Node.js 18

# 安装部署

## 克隆项目

```bash
git clone https://github.com/rescheni/GoSendMessageService.git
```

## 进入项目目录

```bash
cd GoSendMessageService
```

## 设置环境变量

- 在项目目录下的`.env`文件中编辑 docker 目录路径和一些其他配置[目前暂无]

```
CONFIG_DIR =
DATA_DIR =
```

## 部署 Docker 镜像

- 生产环境

```bash
docker-compose -f docker-compose.yml -f docker-compose.yml up
```

- 开发环境

```bash
docker-compose -f docker-compose.yml -f docker-compose.dev.yml up
```
