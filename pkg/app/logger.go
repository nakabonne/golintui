package app

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/shibukawa/configdir"
	"github.com/sirupsen/logrus"

	"github.com/nakabonne/golintui/pkg/config"
)

func newLogger(conf *config.Config) *logrus.Entry {
	var log *logrus.Logger
	if conf.GetDebug() || os.Getenv("DEBUG") == "TRUE" {
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
	log.SetLevel(getLogLevel())
	file, err := os.OpenFile(filepath.Join(globalConfigDir(), "development.log"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic("unable to log to file") // TODO: don't panic (also, remove this call to the `panic` function)
	}
	log.SetOutput(file)
	return log
}

func newProductionLogger() *logrus.Logger {
	log := logrus.New()
	log.Out = ioutil.Discard
	log.SetLevel(logrus.ErrorLevel)
	return log
}

func getLogLevel() logrus.Level {
	strLevel := os.Getenv("LOG_LEVEL")
	level, err := logrus.ParseLevel(strLevel)
	if err != nil {
		return logrus.DebugLevel
	}
	return level
}

func globalConfigDir() string {
	configDirs := configdir.New("nakabonne", "golintui")
	configDir := configDirs.QueryFolders(configdir.Global)[0]
	return configDir.Path
}
