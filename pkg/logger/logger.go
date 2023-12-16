package logger

import (
	"bytes"
	"fmt"
	"github.com/mangohow/gin-template/conf"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

const TimeFormat = "2006-01-02 15:04:05"

var logger *logrus.Logger

func Logger() *logrus.Logger {
	return logger
}

func InitLogger() error {
	logger = logrus.New()
	logger.SetOutput(os.Stdout)

	formatterName := strings.ToLower(conf.Config().Logger.Formatter)
	var formatter logrus.Formatter
	switch formatterName {
	case "json":
		formatter = &logrus.JSONFormatter{}
	case "text":
		formatter = &logrus.TextFormatter{}
	default:
		formatter = &LogFormatter{}
	}
	logger.SetFormatter(formatter)

	level, err := logrus.ParseLevel(conf.Config().Logger.Level)
	if err != nil {
		fmt.Printf("parse log level error:%v", err)
		level = logrus.InfoLevel
	}
	logger.SetLevel(level)
	logger.SetReportCaller(conf.Config().Logger.Caller)

	if conf.Config().Logger.ToFile {
		w, err := OpenLogFile(conf.Config().Logger.FilePath, conf.Config().Logger.FileName)
		if err != nil {
			return err
		}
		logger.SetOutput(w)
	}

	return nil
}

func OpenLogFile(filePath, fileName string) (io.Writer, error) {
	if !strings.HasSuffix(fileName, ".log") {
		fileName = fileName + ".log"
	}
	filep := path.Join(filePath, fileName)

	err := os.Mkdir(filePath, 0644)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}
	file, err := os.OpenFile(filep, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString(entry.Time.Format(TimeFormat))
	buf.WriteString(" ")
	buf.WriteString(strings.ToUpper(entry.Level.String()))
	buf.WriteByte(' ')

	if entry.HasCaller() {
		buf.WriteString(path.Base(entry.Caller.File))
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(entry.Caller.Line))
		buf.WriteByte(' ')
	}

	buf.WriteString(entry.Message)
	buf.WriteByte('\n')

	return buf.Bytes(), nil
}
