package mylog

import (
	"fmt"
	"os"
)

// FileLogger 文件中记录日志的结构体
type FileLogger struct {
	level       int
	logFilePath string
	logFileName string
	logFile     *os.File
}

// NewFileLogger 生成文件日志结构体实例的构造函数
func NewFileLogger(level int, logFilePath, logFileName string) *FileLogger {
	return &FileLogger{
		level:       level,
		logFilePath: logFilePath,
		logFileName: logFileName,
	}
}

// 专门用来初始化文件日志的文件句柄
func (f *FileLogger) initFileLogger() {
	filePath := fmt.Sprintf("%s/%s", f.logFilePath, f.logFileName)
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(fmt.Sprintf("open file:%s failed", filePath))
	}
	f.logFile = file

}

// Debug 记录日志
func (f *FileLogger) Debug(msg string) {
	// 往文件里写

}
