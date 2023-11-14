package logger

import (
	"io"
	internalLog "log"
	"os"

	"github.com/manicar2093/stori-challenge/pkg/config"

	"github.com/sirupsen/logrus"
)

var log *Logger

type (
	Logger struct {
		*logrus.Logger
	}
	Config struct {
		Environment string
	}
)

func GetLogger() *Logger {
	return log
}

func init() {
	logger := logrus.New()
	loggerOutput := io.MultiWriter(os.Stdout)
	logger.SetOutput(loggerOutput)
	internalLog.SetOutput(loggerOutput)
	logger.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint: true,
	})
	logger.SetReportCaller(true)
	if config.Instance.Environment != "prod" {
		logger.SetLevel(logrus.DebugLevel)
	}
	log = &Logger{Logger: logger}
}
