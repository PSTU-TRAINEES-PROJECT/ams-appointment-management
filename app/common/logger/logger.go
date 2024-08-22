package logger

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var loggerEntry *logrus.Entry

func NewLogger() {
	logger = &logrus.Logger{
		Out:       os.Stdout,
		Formatter: &logrus.TextFormatter{ForceColors: true},
		Level:     logrus.InfoLevel,
	}
	buildTag := "develop"
	loggerEntry = logrus.NewEntry(logger)
	loggerEntry = loggerEntry.WithField("build", buildTag)
}

func Info(args ...interface{}) {
	loggerEntry.WithField("file", fileInfo(2)).Info(args...)
}
func Error(args ...interface{}) {
	loggerEntry.WithField("file", fileInfo(2)).Error(args...)
}

func Debug(args ...interface{}) {
	loggerEntry.WithField("file", fileInfo(2)).Debug(args...)
}
func Fatal(args ...interface{}) {
	loggerEntry.WithField("file", fileInfo(2)).Fatal(args...)
}

func fileInfo(skip int) string {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		file = "<???>"
		line = 1
	}
	return fmt.Sprintf("%s:%d", file, line)
}
