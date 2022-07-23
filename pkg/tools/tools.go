package tools

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go_project/pkg/global"
	"path"
)

var Logger *logrus.Logger

func NewLogger() *logrus.Logger {
	logFilePath := global.Conf.Log.FilePath
	logFileName := global.Conf.Log.FileName
	// 日志文件
	fileName := path.Join(logFilePath, logFileName)
	//// 创建文件
	//err := os.MkdirAll(logFilePath, os.ModeDir|os.ModePerm)
	//if err != nil {
	//	fmt.Println("Create Dir err:", err)
	//}
	//// 写入文件
	//src, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
	//if err != nil {
	//	fmt.Println("Open log file err:", err)
	//}
	//writers := []io.Writer{
	//	src,
	//	os.Stdout,
	//}
	//fileAndStdoutWriter := io.MultiWriter(writers...)
	// 实例化
	logger := logrus.New()
	//// 设置输出
	//if err == nil {
	//	logger.SetOutput(fileAndStdoutWriter)
	//} else {
	//	logger.Info("failed to log to file.")
	//}
	// 日志级别
	logger.SetLevel(logrus.InfoLevel)

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  fileName,
		logrus.ErrorLevel: fileName,
	}
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	})
	// 新增 Hook
	logger.Hooks.Add(lfshook.NewHook(pathMap, &logrus.JSONFormatter{
		TimestampFormat:"2006-01-02 15:04:05",
	}))
	logger.Info("logger init")
	return logger
}
