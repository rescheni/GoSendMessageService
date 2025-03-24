curl -X POST http://localhost:8080/send_wxpusher_datetime \
-H "Content-Type: application/json" \
-d '{
    "year": 2023,
    "month": 12,
    "day": 25,
    "hour": 10,
    "minute": 30,
    "second": 0,
    "title": "测试标题",
    "message": "测试消息内容",
    "user": "用户ID"
}'
