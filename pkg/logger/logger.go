package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func GetLogger() (*logrus.Logger, error) {
	log := logrus.New()
	logPath := "./app.log"
	log.SetFormatter(&logrus.JSONFormatter{})
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}
	log.SetOutput(file)
	return log, nil
}
