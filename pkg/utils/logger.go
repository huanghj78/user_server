package utils

import (
	"bytes"
	"fmt"
	"user_server/Config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"time"
)

var Logger *logrus.Logger

type MyFormatter struct {}
func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error){
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var newLog string

	//HasCaller()为true才会有调用信息
	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		newLog = fmt.Sprintf("[%s] [%s] [%s:%d %s] %s\n",
			timestamp, entry.Level, fName, entry.Caller.Line, entry.Caller.Function, entry.Message)
	} else{
		newLog = fmt.Sprintf("[%s] [%s] %s\n", timestamp, entry.Level, entry.Message)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}


func init() {
	fileName := path.Join(Config.LogConf.FilePath, Config.LogConf.FileName + ".log")
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}
	Logger = logrus.New()
	Logger.SetLevel(logrus.DebugLevel)
	// 设置输出
	Logger.Out = src

	// 设置rotate log
	logWriter, err := rotatelogs.New(
		fileName+".%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),
		// 最大保存时间
		rotatelogs.WithMaxAge(7*24*time.Hour),
		// 切割时间间隔
		rotatelogs.WithRotationTime(24*time.Hour),
		)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel: logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	Logger.SetReportCaller(true)
	lfHook := lfshook.NewHook(writeMap, &MyFormatter{})
	Logger.AddHook(lfHook)
}

