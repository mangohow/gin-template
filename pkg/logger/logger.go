package logger

import (
	"bytes"
	"gin-template/conf"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

var TimeFormat = "2006-01-02 15:04:05"

var logger *logrus.Logger

func Logger() *logrus.Logger {
	return logger
}

func InitLogger() error {
	logger = logrus.New()
	logger.SetFormatter(&LogFormatter{})
	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetReportCaller(true)

	if conf.LoggerConfig.ToFile {
		w, err := OpenLogFile(conf.LoggerConfig.FilePath, conf.LoggerConfig.FileName)
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

	file, err := os.OpenFile(filep, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}

type LogFormatter struct{}

func (t *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteByte('[')
	buf.WriteString(entry.Time.Format(TimeFormat))
	buf.WriteString("][")
	buf.WriteString(strings.ToUpper(entry.Level.String()))
	buf.WriteByte(']')

	if entry.HasCaller() {
		buf.WriteByte('[')
		buf.WriteString(path.Base(entry.Caller.File))
		buf.WriteByte(':')
		buf.WriteString(strings.Split(entry.Caller.Function, ".")[1])
		buf.WriteByte(':')
		buf.WriteString(strconv.Itoa(entry.Caller.Line))
		buf.WriteByte(']')
	}
	buf.WriteByte(' ')
	buf.WriteString(entry.Message)
	buf.WriteByte('\n')

	return buf.Bytes(), nil
}
