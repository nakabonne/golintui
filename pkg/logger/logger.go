package logger

import (
	"io/ioutil"
	"os"

	"github.com/k0kubun/pp"

	"github.com/sirupsen/logrus"

	"github.com/nakabonne/golintui/pkg/config"
)

func NewLogger(conf *config.Config) *logrus.Entry {
	var log *logrus.Logger
	if conf.GetDebug() {
		log = newDevelopmentLogger()
	} else {
		log = newProductionLogger()
	}

	log.Formatter = &logrus.JSONFormatter{}

	return log.WithFields(logrus.Fields{
		"debug":     conf.GetDebug(),
		"version":   conf.GetVersion(),
		"commit":    conf.GetCommit(),
		"buildDate": conf.GetBuildDate(),
	})
}

func newDevelopmentLogger() *logrus.Logger {
	log := logrus.New()
	log.SetLevel(logrus.TraceLevel)
	// TODO: Use regular directories for configuration files by using https://github.com/shibukawa/configdir.
	file, err := os.OpenFile("development.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("unable to log to file")
	}
	log.SetOutput(file)
	pp.SetDefaultOutput(log.Out)
	return log
}

func newProductionLogger() *logrus.Logger {
	log := logrus.New()
	log.Out = ioutil.Discard
	log.SetLevel(logrus.ErrorLevel)
	return log
}
