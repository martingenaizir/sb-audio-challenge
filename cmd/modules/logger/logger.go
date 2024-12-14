package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

// cutting corners here:
// this is a very simple logger.

var Log *logrus.Logger

func init() {
	Log = &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.TextFormatter{
			FullTimestamp: true,
		},
		Level: logrus.ErrorLevel,
	}
}

func SetLogLevel(level string) {
	if l, err := logrus.ParseLevel(level); err != nil {
		fmt.Println(err)
	} else {
		Log.SetLevel(l)
	}
}

func Error(err error, message string, args ...any) {
	if Log.Level >= logrus.ErrorLevel {
		if len(args) > 0 {
			message = fmt.Sprintf(message, args...)
		}

		Log.Errorf("%s - ERROR: %s", message, err)
	}
}

func Info(message string, args ...any) {
	if Log.Level >= logrus.InfoLevel {
		if len(args) > 0 {
			message = fmt.Sprintf(message, args...)
		}

		Log.Info(message)
	}
}

func Debug(message string, args ...any) {
	if Log.Level >= logrus.DebugLevel {
		if len(args) > 0 {
			message = fmt.Sprintf(message, args...)
		}

		Log.Debug(message)
	}
}
