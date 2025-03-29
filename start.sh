#!/bin/sh

# 启动服务
./server &
SERVER_PID=$!

npm start &
FRONTEND_PID=$!

trap "kill -TERM $SERVER_PID $FRONTEND_PID" SIGTERM
wait $SERVER_PID $FRONTEND_PID