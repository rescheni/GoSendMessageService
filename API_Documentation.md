# GoMessageService API 文档

## 基础信息
- **主机地址**: `http://<api_host>:<api_port>`  
  （默认值为 `http://0.0.0.0:8080`，可在 `Message_main_config.yaml` 中配置）
- **API Key**: 每个请求需要携带 `api_key` 参数，用于验证身份。

---

## 1. 消息发送接口

### 1.1 发送到 WxPusher
- **URL**: `/send/wxpusher`
- **方法**: `POST`
- **请求体**:
  ```json
  {
    "api_key": "1234567890",
    "message": "消息内容",
    "title": "消息标题"
  }
  ```
- **响应**:
  ```json
  {
    "message": "Message sent successfully"
  }
  ```

---

### 1.2 发送到钉钉
- **URL**: `/send/dingding`
- **方法**: `POST`
- **请求体**:
  ```json
  {
    "api_key": "1234567890",
    "message": "消息内容"
  }
  ```
- **响应**:
  ```json
  {
    "message": "Message sent successfully"
  }
  ```

---

### 1.3 发送到 Server 酱
- **URL**: `/send/server_jiang`
- **方法**: `POST`
- **请求体**:
  ```json
  {
    "api_key": "1234567890",
    "message": "消息内容",
    "title": "消息标题"
  }
  ```
- **响应**:
  ```json
  {
    "message": "Message sent successfully"
  }
  ```

---

### 1.4 发送邮件
- **URL**: `/send/email`
- **方法**: `POST`
- **请求体**:
  ```json
  {
    "api_key": "1234567890",
    "message": "邮件内容",
    "title": "邮件标题"
  }
  ```
- **响应**:
  ```json
  {
    "message": "Message sent successfully"
  }
  ```

---

### 1.5 发送到飞书
- **URL**: `/send/feishu`
- **方法**: `POST`
- **请求体**:
  ```json
  {
    "api_key": "1234567890",
    "message": "消息内容",
    "title": "消息标题"
  }
  ```
- **响应**:
  ```json
  {
    "message": "Message sent successfully"
  }
  ```

---

### 1.6 发送到 Napcat QQ
- **URL**: `/send/napcat_qq`
- **方法**: `POST`
- **请求体**:
  ```json
  {
    "api_key": "1234567890",
    "message": "消息内容"
  }
  ```
- **响应**:
  ```json
  {
    "message": "Message sent successfully"
  }
  ```

---

## 2. 定时任务接口

### 2.1 设置定时任务
- **URL**: `/cron/set`
- **方法**: `POST`
- **请求体**:
  ```json
  {
    "api_key": "1234567890",
    "cron_expr": "0 0 * * * *",
    "message": "定时任务消息内容",
    "title": "定时任务标题",
    "task_type": "wxpusher",
    "is_open": true
  }
  ```
  - `cron_expr`: 定时任务的 Cron 表达式。
  - `task_type`: 任务类型，可选值为：
    - `wxpusher`
    - `dingding`
    - `server_jiang`
    - `email`
    - `feishu`
    - `napcat_qq`
- **响应**:
  ```json
  {
    "message": "Cron task set successfully",
    "cron_expr": "0 0 * * * *"
  }
  ```

---

### 2.2 删除定时任务
- **URL**: `/cron/delete`
- **方法**: `GET`
- **查询参数**:
  - `api_key`: API Key
  - `entryid`: 定时任务的 EntryID
- **响应**:
  ```json
  {
    "message": "Cron task deleted successfully"
  }
  ```

---

### 2.3 获取定时任务列表
- **URL**: `/cron/list`
- **方法**: `GET`
- **查询参数**:
  - `api_key`: API Key
- **响应**:
  ```json
  {
    "tasks": [
      {
        "EntryID": 1,
        "Expr": "0 0 * * * *"
      },
      {
        "EntryID": 2,
        "Expr": "0 30 * * * *"
      }
    ]
  }
  ```

---

## 错误响应
所有接口的错误响应格式如下：
```json
{
  "error": "错误描述"
}
```

---
