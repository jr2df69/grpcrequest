package grpcrequest

import "github.com/sirupsen/logrus"

var logger *logrus.Logger

func getLogger() *logrus.Logger {
	if logger == nil {
		return logrus.New()
	}

	return logger
}
