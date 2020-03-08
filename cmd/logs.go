package cmd

import (
	"github.com/dkbrummitt/go-build-template/pkg/version"
	"github.com/sirupsen/logrus"
)

var (
	Logger        *logrus.Logger
	ContextLogger *logrus.Entry
)

func NewLogger(asJson bool, reportCaller bool, logLevel string) (*logrus.Entry, *logrus.Logger) {
	log := logrus.New()
	var logger *logrus.Entry
	if asJson {
		log.SetFormatter(&logrus.JSONFormatter{})
	}
	if reportCaller {
		log.SetReportCaller(true)
	}

	switch logLevel {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	}
	logger = log.WithFields(logrus.Fields{
		"appName":    version.AppName,
		"appVersion": version.VERSION,
	})

	return logger, log

}

func SetAppLog(asJson bool, reportCaller bool, logLevel string) {
	if ContextLogger == nil {
		ContextLogger, Logger = NewLogger(asJson, reportCaller, logLevel)
	}

}
