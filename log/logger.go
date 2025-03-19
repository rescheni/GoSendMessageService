package log

import (
	basic "GoMessageService/Basic"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 全局 Logger
var Logger *logrus.Logger

func init() {

	cfg := basic.LoadConfig()

	Logger = logrus.New()

	// 设置日志格式（JSON / 文本）
	Logger.SetFormatter(&logrus.JSONFormatter{}) // JSON 格式
	// Logger.SetFormatter(&logrus.TextFormatter{}) // 文本格式

	// 设置日志级别
	Logger.SetLevel(logrus.InfoLevel)

	// 设置日志输出：文件 + 控制台
	fileOutput := &lumberjack.Logger{
		Filename:   cfg.Log.LogPath, // 日志文件路径
		MaxSize:    10,              // 最大 10MB
		MaxBackups: 5,               // 最多保留 5 个日志文件
		MaxAge:     30,              // 保留 30 天
		Compress:   true,            // 是否压缩
	}

	// 使用 MultiWriter 同时输出到文件和控制台
	Logger.SetOutput(io.MultiWriter(fileOutput, os.Stdout))
}
