package log

import (
	"os"
	"path"
	"time"

	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func toFile(logger *logrus.Logger, logFilePath string) {

	// 日志文件
	fileName := logFilePath // path.Join(logFilePath, logFileName)

	logFileDir := path.Dir(logFilePath)
	if _, err := os.Stat(logFileDir); err != nil {
		err = os.MkdirAll(logFileDir, 0711)
	}

	// src, err := os.OpenFile(fileName+".txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	//
	// if _, errr := src.WriteString("aa"); errr != nil {
	// 	fmt.Println(errr)
	// }
	//
	// Logger.Out = src

	logWriter, err := rotatelogs.New(
		fileName+"_%Y%m%d.log",                    // 分割后的文件名称
		rotatelogs.WithLinkName(fileName),         // 生成软链，指向最新日志文件
		rotatelogs.WithRotationCount(365),         // 最多存365个文件
		rotatelogs.WithRotationTime(24*time.Hour), // 设置日志切割时间间隔(1天)
	)

	if err != nil {
		logger.Errorf("config local file system Logger error. %+v", errors.WithStack(err))
	}

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: logWriter,
		logrus.InfoLevel:  logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.FatalLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}, &logrus.TextFormatter{
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: false,
	})

	logger.AddHook(lfHook)

}
