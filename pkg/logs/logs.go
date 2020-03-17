package cmd

import (
	"reflect"
	"strings"

	"github.com/dkbrummitt/go-build-template/pkg/version"
	"github.com/sirupsen/logrus"
)

// NewLogger returns a new Logger(for Server logs) and LogEntry(for App Context logs)
func NewLogger(asJSON bool, reportCaller bool, logLevel string) (*logrus.Entry, *logrus.Logger) {
	log := logrus.New()
	var logger *logrus.Entry

	if asJSON {
		log.SetFormatter(&logrus.JSONFormatter{})
	}

	logger = log.WithFields(logrus.Fields{
		"appName":    version.AppName,
		"appVersion": version.VERSION,
	})

	if reportCaller {
		log.SetReportCaller(true)
		log.Warn("report caller is on this will cause slower performance")
	}

	switch logLevel {
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	} // of switch

	return logger, log

}

// GetPackage returns the associated package for the thing provided
func GetPackage(thing interface{}) string {
	var pkg string

	pkg = reflect.TypeOf(thing).PkgPath()
	pkg = strings.ReplaceAll(pkg, "/", ".")

	return pkg
}
