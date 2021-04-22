package cfg

import (
	"github.com/sirupsen/logrus"
)

//singleton logger
var Logger *logrus.Entry

func init() {
	Logger = logrus.WithFields(logrus.Fields{})
}

func NewLogger() {
	Logger = logrus.WithFields(logrus.Fields{})
}

func BindFields(fields logrus.Fields) {
	Logger = Logger.WithFields(fields)
}
